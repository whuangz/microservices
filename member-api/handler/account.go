package handler

import (
	"log"
	"net/http"

	"github.com/whuangz/microservices/member-api/usecase"

	"github.com/labstack/echo/v4"
)

type Account struct {
	l       *log.Logger
	usecase *usecase.AccountUsecase
}

func NewAccountHandler(l *log.Logger, au *usecase.AccountUsecase) *Account {
	return &Account{l, au}
}

func (a *Account) CreateAccount(c echo.Context) (err error) {
	return c.JSON(http.StatusBadRequest, err.Error())
}
