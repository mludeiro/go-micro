package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

var GetArticle func(uint, []string) *entity.Article = repository.GetArticle

var GetArticles func([]string) []entity.Article = repository.GetArticles

var AddArticle func(*entity.Article) (*entity.Article, error) = repository.AddArticle

var DeleteArticle func(uint) *entity.Article = repository.DeleteArticle
