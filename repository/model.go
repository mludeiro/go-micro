package repository

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name  string
	Price int16
}

func (Article) TableName() string {
	return "article"
}

func Migrate() {
	getDB().AutoMigrate(&Article{})
}
