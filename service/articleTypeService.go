package service

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
)

func GetArticleTypes(fetchs []string) []entity.ArticleType {
	repo := repository.ArticleType{DB: database.GetDB()}
	return repo.GetAll(fetchs)
}
