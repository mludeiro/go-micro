package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

func GetArticleType(id uint, fetchs []string) *entity.ArticleType {
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

func GetArticleTypes(fetchs []string) []entity.ArticleType {
	articleTypes := []entity.ArticleType{}

	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articleTypes)
	return articleTypes
}

func AddArticleType(at *entity.ArticleType) (*entity.ArticleType, error) {
	if database.GetDB().Create(at).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return at, nil
}

func DeleteArticleType(id uint) *entity.ArticleType {
	articleType := entity.ArticleType{}
	rows := database.GetDB().Where("deleted_at is NULL").Delete(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}
