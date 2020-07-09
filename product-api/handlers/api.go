package handlers

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	_ "github.com/whuangz/microservices/product-api/docs"

	"github.com/labstack/echo/v4"
	mysql "github.com/whuangz/microservices/engine-helper/database"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

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

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn := mysql.Init(src)

	product := NewProducts(l)
	v1 := e.Group("/api/v1")
	{
		v1.GET("/products", product.getProducts)
		v1.GET("/products/:id", product.getProductDetail)
		v1.POST("/products", product.addProduct, product.middlewareProductValidation)
		v1.PUT("/products/:id", product.updateProduct, product.middlewareProductValidation)

	}

	test := NewTest(l)
	e.GET("/", test.TestHello, mysql.TransactionHandler(dbConn))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
