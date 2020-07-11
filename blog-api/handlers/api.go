package handlers

import (
	"fmt"
	"log"

	"github.com/whuangz/microservices/blog-api/repository"
	"github.com/whuangz/microservices/blog-api/usecase"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	mysql "github.com/whuangz/microservices/helpers/database"
	"github.com/whuangz/microservices/helpers/middlewares"
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

func databaseConnection() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	return src
}

func Router(e *echo.Echo, l *log.Logger) {

	dbConn := mysql.Init(databaseConnection())

	//Repo Definition
	articleRepo := repository.NewArticleRepo()
	authorRepo := repository.NewAuthorRepo()

	//Usecase Definition
	au := usecase.NewArticleUseCase(articleRepo, authorRepo)
	//Article Handler
	article := NewArticles(l, au)
	v1 := e.Group("/api/v1")
	{
		v1.GET("/articles", article.FetchArticles, mysql.TransactionHandler(dbConn))
		v1.POST("/articles", article.CreateArticle, mysql.TransactionHandler(dbConn))
	}

	//Global Middlewares
	cors := middlewares.InitMiddleware()
	e.Use(cors.CORS)
}
