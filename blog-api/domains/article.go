package domains

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Article) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
