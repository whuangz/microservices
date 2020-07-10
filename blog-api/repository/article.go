package repository

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/whuangz/microservices/blog-api/domains"
)

type Article struct {
}

func (a *Article) GetArticles(c echo.Context) ([]*domains.Article, error) {
	tx := c.Get("Tx").(*dbr.Tx)

	articles := make([]*domains.Article, 0)
	builder := tx.SelectBySql("SELECT id,title,content, author_id, updated_at, created_at FROM article")
	_, err := builder.Load(&articles)

	logrus.Error(err)

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) CreateArticle(c echo.Context, da *domains.Article) (err error) {
	tx := c.Get("Tx").(*dbr.Tx)

	err = InsertArticle(tx, da)
	err = CreateAuthor(tx)

	return err
}

func CreateAuthor(tx *dbr.Tx) error {

	builder := tx.InsertBySql("INSERT author SET name=?", "Not Super User")
	_, err := builder.Exec()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func InsertArticle(tx *dbr.Tx, da *domains.Article) (err error) {
	builder := tx.InsertBySql("INSERT  article SET title=? , content=?", da.Title, da.Content)
	_, err = builder.Exec()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}