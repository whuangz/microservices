package domains

import (
	"github.com/go-playground/validator/v10"
	"github.com/gocraft/dbr/v2"
)

type Article struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title" validate:"required"`
	Content   string         `json:"content" validate:"required"`
	Author    Author         `json:"author"`
	UpdatedAt dbr.NullString `json:"updated_at"`
	CreatedAt dbr.NullString `json:"created_at"`
}

func (a *Article) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
