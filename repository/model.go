package repository

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name          string
	Price         int16
	ArticleTypeID uint
	ArticleType   *ArticleType `json:",omitempty"`
}

func (Article) TableName() string {
	return "Article"
}

type ArticleType struct {
	ID      uint `gorm:"primarykey"`
	Name    string
	Article []Article `json:",omitempty"`
}

func (ArticleType) TableName() string {
	return "ArticleType"
}

func Migrate() {
	getDB().AutoMigrate(&ArticleType{}, &Article{})
}
