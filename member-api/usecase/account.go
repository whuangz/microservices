package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/whuangz/microservices/member-api/domain"
	"github.com/whuangz/microservices/member-api/repository"
)

type AccountUsecase struct {
	accountRepo *repository.AccountRepo
}

func NewAccountUsecase(accRepo *repository.AccountRepo) *AccountUsecase {
	return &AccountUsecase{accRepo}
}

func (au *AccountUsecase) CreateAccount(c echo.Context, da *domain.Account) (err error) {
	return err
}
