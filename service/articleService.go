package service

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
)

func GetArticle(id uint, fetchs []string) *entity.Article {
	return getRepo().GetArticle(id, fetchs)
}

func GetArticles(fetchs []string) []entity.Article {
	return getRepo().GetArticles(fetchs)
}

func AddArticle(dto *entity.Article) (*entity.Article, error) {
	return getRepo().AddArticle(dto)
}

func DeleteArticle(id uint) *entity.Article {
	return getRepo().DeleteArticle(id)
}

func getRepo() repository.IArticleRepository {
	return &repository.ArticleRepository{DB: database.GetDB()}
}
