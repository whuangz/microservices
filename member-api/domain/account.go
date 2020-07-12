package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	ID        int64     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
