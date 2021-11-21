package repository

import (
	"errors"
	"go-micro/entity"

	"gorm.io/gorm"
)

type IArticle interface {
	Get(id uint, fetchs []string) *entity.Article
	GetAll(fetchs []string) []entity.Article
	Add(a *entity.Article) (*entity.Article, error)
	Delete(id uint) *entity.Article
}

type Article struct {
	DB *gorm.DB
}

func (this *Article) Get(id uint, fetchs []string) *entity.Article {
	article := entity.Article{}
	db := this.DB

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	rows := db.Find(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}

func (this *Article) GetAll(fetchs []string) []entity.Article {
	articles := []entity.Article{}
	db := this.DB

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func (this *Article) Add(a *entity.Article) (*entity.Article, error) {
	if this.DB.Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func (this *Article) Delete(id uint) *entity.Article {
	article := entity.Article{}
	rows := this.DB.Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
