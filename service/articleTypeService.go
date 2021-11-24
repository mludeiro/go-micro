package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type ArticleType struct {
	Repository *repository.ArticleType
}

func (this *ArticleType) GetArticleTypes(fetchs []string) []entity.ArticleType {
	return this.Repository.GetAll(fetchs)
}
