package repository

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type AuthorRepo struct{}

func NewAuthorRepo() *AuthorRepo {
	return &AuthorRepo{}
}

func (a *AuthorRepo) CreateAuthor(c echo.Context) error {

	tx := c.Get("Tx").(*sql.Tx)

	rawQuery := "INSERT author SET name=?"
	_, err := tx.Exec(rawQuery, "Not Super User")

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
