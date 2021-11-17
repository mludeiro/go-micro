package repository

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
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
