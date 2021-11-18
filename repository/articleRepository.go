package repository

import (
	"errors"
)

func GetArticle(id uint, fetchs []string) *Article {
	article := Article{}
	db := getDB()

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

func GetArticles(fetchs []string) []Article {
	articles := []Article{}
	db := getDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func AddArticle(a *Article) (*Article, error) {
	if getDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func DeleteArticle(id uint) *Article {
	article := Article{}
	rows := getDB().Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
