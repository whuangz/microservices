package repository

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
)

type AuthorRepo struct{}

func NewAuthorRepo() *AuthorRepo {
	return &AuthorRepo{}
}

func (a *AuthorRepo) CreateAuthor(c echo.Context) error {
	tx := c.Get("Tx").(*dbr.Tx)

	builder := tx.InsertBySql("INSERT author SET name=?", "Not Super User")
	_, err := builder.Exec()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
