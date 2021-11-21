package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"

	"gorm.io/gorm"
)

type IArticleTypeRepository interface {
	Get(id uint, fetchs []string) *entity.ArticleType
	GetAll(fetchs []string) []entity.ArticleType
	Add(a *entity.ArticleType) (*entity.ArticleType, error)
	Delete(id uint) *entity.ArticleType
}

type ArticleTypeRepository struct {
	DB *gorm.DB
}

func (this ArticleTypeRepository) Get(id uint, fetchs []string) *entity.ArticleType {
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

func (this ArticleTypeRepository) GetAll(fetchs []string) []entity.ArticleType {
	articleTypes := []entity.ArticleType{}

	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articleTypes)
	return articleTypes
}

func (this ArticleTypeRepository) Add(at *entity.ArticleType) (*entity.ArticleType, error) {
	if database.GetDB().Create(at).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return at, nil
}

func (this ArticleTypeRepository) Delete(id uint) *entity.ArticleType {
	articleType := entity.ArticleType{}
	rows := database.GetDB().Delete(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}
