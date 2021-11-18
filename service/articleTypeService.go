package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

var GetArticleTypes func(uint, []string) *entity.ArticleType = repository.GetArticleType
