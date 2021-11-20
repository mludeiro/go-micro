package entity

type ArticleType struct {
	ID      uint      `gorm:"primarykey"`
	Name    string    `gorm:"not null"`
	Article []Article `json:",omitempty"`
}

func (ArticleType) TableName() string {
	return "ArticleType"
}
