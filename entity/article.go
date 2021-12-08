package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID            uint           `gorm:"primarykey"`
	CreatedAt     time.Time      ``
	UpdatedAt     time.Time      ``
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Name          string         `gorm:"not null;default:null"`
	Price         int16          ``
	ArticleTypeID uint           ``
	ArticleType   *ArticleType   `json:",omitempty"`
}

type ArticleResultSet struct {
	ResultSet
	Query
	Data []Article
}

func (Article) TableName() string {
	return "Article"
}
