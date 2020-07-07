package errors

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetResponseError(err error, msg string) (int, *ResponseError) {
	logrus.Error(err)

	r := &ResponseError{msg}
	switch err {
	case ErrBadParamInput:
		return http.StatusBadRequest, r
	case ErrNotFound:
		return http.StatusNotFound, r
	case ErrConflict:
		return http.StatusConflict, r
	case ErrUnprocessableEntity:
		return http.StatusUnprocessableEntity, &ResponseError{ErrUnprocessableEntity.Error()}
	default:
		return http.StatusInternalServerError, &ResponseError{ErrInternalServerError.Error()}
	}
}
