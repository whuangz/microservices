package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/whuangz/microservices/blog-api/domains"
	"github.com/whuangz/microservices/blog-api/repository"
)

type ArticleUseCase struct {
	articleRepo *repository.ArticleRepo
	authorRepo  *repository.AuthorRepo
}

func NewArticleUseCase(article *repository.ArticleRepo, author *repository.AuthorRepo) *ArticleUseCase {
	return &ArticleUseCase{article, author}
}

func (au *ArticleUseCase) GetArticles(c echo.Context) ([]*domains.Article, error) {
	articles, err := au.articleRepo.GetArticles(c)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (au *ArticleUseCase) CreateArticle(c echo.Context, da *domains.Article) (err error) {

	err = au.articleRepo.InsertArticle(c, da)
	err = au.authorRepo.CreateAuthor(c)

	return err
}
