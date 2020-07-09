package handlers

import (
	"log"

	_ "github.com/whuangz/microservices/product-api/docs"

	mysql "github.com/whuangz/microservices/helpers/database"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Products API
// @version 1.0
// @description This is Product list API.

// @contact.name Product API Support
// @contact.email n.hua.drt@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func Router(e *echo.Echo, l *log.Logger) {
	product := NewProducts(l)
	v1 := e.Group("/api/v1")
	{
		v1.GET("/products", product.getProducts)
		v1.GET("/products/:id", product.getProductDetail)
		v1.POST("/products", product.addProduct, product.middlewareProductValidation)
		v1.PUT("/products/:id", product.updateProduct, product.middlewareProductValidation)

	}

	test := NewTest(l)
	e.GET("/", test.TestHello, mysql.Init())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
