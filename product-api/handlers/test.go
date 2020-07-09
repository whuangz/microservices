package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	l *log.Logger
}

func NewTest(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) TestHello(c echo.Context) error {
	return nil
}
