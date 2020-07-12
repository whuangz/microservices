package handler

import (
	"log"

	"github.com/whuangz/microservices/member-api/repository"
	"github.com/whuangz/microservices/member-api/usecase"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	db "github.com/whuangz/microservices/helpers/database"
	"github.com/whuangz/microservices/helpers/middlewares"
)

func init() {
	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func Router(e *echo.Echo, l *log.Logger) {

	dbConn := db.Init(viper.GetString(`db.uri`))

	//Global Middlewares
	cors := middlewares.InitMiddleware()
	txMiddleware := db.TransactionHandler(dbConn)
	e.Use(cors.CORS)

	//Repo Definition
	accountRepo := repository.NewAccountRepo()

	//Usecase Definition
	au := usecase.NewAccountUsecase(accountRepo)

	//Article Handler
	account := NewAccountHandler(l, au)
	v1 := e.Group("/api/v1")
	{
		v1.POST("/accounts", account.CreateAccount, txMiddleware)
	}

}
