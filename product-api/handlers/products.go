package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/whuangz/microservices/product-api/domains"
)

type Products struct {
	l *log.Logger
}

func NewProdutcs(e *echo.Echo, l *log.Logger) {
	p := &Products{l}
	e.GET("/products", p.getProducts)
	e.GET("/products/:id", p.getProductDetail)
	e.POST("/products", p.addProduct)
	e.PUT("/products/:id", p.updateProduct)
	e.Use(p.middlewareProductValidation)
}

func (p *Products) getProducts(c echo.Context) error {
	lp := domains.GetProducts()
	return c.JSON(http.StatusOK, lp)
}

func (p *Products) getProductDetail(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domains.ErrNotFound.Error())
	}

	product, _, err := domains.FindProduct(idP)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (p *Products) addProduct(c echo.Context) (err error) {

	prod := c.Get("keyProduct").(domains.Product)
	domains.AddProduct(&prod)

	return c.JSON(http.StatusCreated, prod)
}

func (p *Products) updateProduct(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, domains.ErrNotFound.Error())
	}

	prod := c.Get("keyProduct").(domains.Product)

	err = domains.UpdateProduct(idP, &prod)
	if err != nil {
		return c.JSON(http.StatusNotFound, domains.ErrNotFound.Error())
	}

	return c.JSON(http.StatusCreated, prod)
}

func (p *Products) middlewareProductValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var prod domains.Product
		err := c.Bind(&prod)

		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		//validation
		err = prod.Validate()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		c.Set("keyProduct", prod)
		return next(c)
	}
}

// Move to Helpers

type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domains.ErrInternalServerError:
		return http.StatusInternalServerError
	case domains.ErrNotFound:
		return http.StatusNotFound
	case domains.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
