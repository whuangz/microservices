package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/whuangz/microservices/blog-api/domains"
	"github.com/whuangz/microservices/blog-api/repository"
)

type Articles struct {
	l    *log.Logger
	repo *repository.Article
}

func NewArticles(l *log.Logger) *Articles {
	return &Articles{l, &repository.Article{}}
}

func (a *Articles) FetchArticles(c echo.Context) error {
	articles, err := a.repo.GetArticles(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, articles)
}

func (a *Articles) CreateArticle(c echo.Context) (err error) {
	var article domains.Article
	err = c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = article.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.repo.CreateArticle(c, &article)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, article)
}
