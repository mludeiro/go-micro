package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"

	"gorm.io/gorm"
)

type IArticleType interface {
	Get(id uint, fetchs []string) *entity.ArticleType
	GetAll(fetchs []string) []entity.ArticleType
	Add(a *entity.ArticleType) (*entity.ArticleType, error)
	Delete(id uint) *entity.ArticleType
}

type ArticleType struct {
	DB *gorm.DB
}

func (this ArticleType) Get(id uint, fetchs []string) *entity.ArticleType {
	articleType := entity.ArticleType{}

	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	rows := db.Find(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}

func (this ArticleType) GetAll(fetchs []string) []entity.ArticleType {
	articleTypes := []entity.ArticleType{}

	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articleTypes)
	return articleTypes
}

func (this ArticleType) Add(at *entity.ArticleType) (*entity.ArticleType, error) {
	if database.GetDB().Create(at).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return at, nil
}

func (this ArticleType) Delete(id uint) *entity.ArticleType {
	articleType := entity.ArticleType{}
	rows := database.GetDB().Delete(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}
