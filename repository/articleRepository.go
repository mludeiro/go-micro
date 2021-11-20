package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	GetArticle(id uint, fetchs []string) *entity.Article
	GetArticles(fetchs []string) []entity.Article
	AddArticle(a *entity.Article) (*entity.Article, error)
	DeleteArticle(id uint) *entity.Article
}

type ArticleRepository struct {
	DB *gorm.DB
}

func (ArticleRepository) GetArticle(id uint, fetchs []string) *entity.Article {
	article := entity.Article{}
	db := database.GetDB()

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

func (ArticleRepository) GetArticles(fetchs []string) []entity.Article {
	articles := []entity.Article{}
	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func (ArticleRepository) AddArticle(a *entity.Article) (*entity.Article, error) {
	if database.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func (ArticleRepository) DeleteArticle(id uint) *entity.Article {
	article := entity.Article{}
	rows := database.GetDB().Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
