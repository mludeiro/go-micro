package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IActicle interface {
	Get(uint, []string) *entity.Article
	GetAll([]string) []entity.Article
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) *entity.Article
}

type Article struct {
	repo repository.IArticle
}

func (this Article) Get(id uint, fetchs []string) *entity.Article {
	return this.repo.Get(id, fetchs)
}

func (this Article) GetAll(fetchs []string) []entity.Article {
	return this.repo.GetAll(fetchs)
}

func (this Article) Add(dto *entity.Article) (*entity.Article, error) {
	return this.repo.Add(dto)
}

func (this Article) Delete(id uint) *entity.Article {
	return this.repo.Delete(id)
}
