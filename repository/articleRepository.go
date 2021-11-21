package repository

import (
	"errors"
	"go-micro/entity"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	Get(id uint, fetchs []string) *entity.Article
	GetAll(fetchs []string) []entity.Article
	Add(a *entity.Article) (*entity.Article, error)
	Delete(id uint) *entity.Article
}

type ArticleRepository struct {
	DB *gorm.DB
}

func (this ArticleRepository) Get(id uint, fetchs []string) *entity.Article {
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

func (this ArticleRepository) GetAll(fetchs []string) []entity.Article {
	articles := []entity.Article{}
	db := this.DB

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func (this ArticleRepository) Add(a *entity.Article) (*entity.Article, error) {
	if this.DB.Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func (this ArticleRepository) Delete(id uint) *entity.Article {
	article := entity.Article{}
	rows := this.DB.Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
