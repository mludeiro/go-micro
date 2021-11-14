package repository

import (
	"go-micro/model"

	"gorm.io/gorm"
)

type ArticlesDto struct {
	gorm.Model
	model.Article
}

func GetArticles() []ArticlesDto {
	return []ArticlesDto{}
}
