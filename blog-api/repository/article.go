package repository

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/whuangz/microservices/blog-api/domains"
)

type ArticleRepo struct {
}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

func (a *ArticleRepo) GetArticles(c echo.Context) ([]*domains.Article, error) {
	tx := c.Get("Tx").(*sql.Tx)

	rawQuery := "SELECT id,title,content, author_id, updated_at, created_at FROM article"

	rows, err := tx.Query(rawQuery)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	articles := make([]*domains.Article, 0)

	for rows.Next() {
		t := &domains.Article{}
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = domains.Author{
			ID: authorID,
		}
		articles = append(articles, t)
	}

	return articles, nil
}

// func (a *ArticleRepo) InsertArticle(c echo.Context, da *domains.Article) (err error) {
// 	tx := c.Get("Tx").(*dbr.Tx)

// 	builder := tx.InsertBySql("INSERT  article SET title=? , content=?", da.Title, da.Content)
// 	_, err = builder.Exec()

// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return nil
// }
