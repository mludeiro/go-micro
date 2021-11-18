package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

func GetArticle(id uint, fetchs []string) *entity.Article {
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

func GetArticles(fetchs []string) []entity.Article {
	articles := []entity.Article{}
	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func AddArticle(a *entity.Article) (*entity.Article, error) {
	if database.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func DeleteArticle(id uint) *entity.Article {
	article := entity.Article{}
	rows := database.GetDB().Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
