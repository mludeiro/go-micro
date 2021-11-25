package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type ArticleType struct {
	Repository repository.IArticleType
}

func (this *ArticleType) GetArticleTypes(fetchs []string) []entity.ArticleType {
	return this.Repository.GetAll(fetchs)
}
