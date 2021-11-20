package entity

type InvoiceLine struct {
	ID        uint     `gorm:"primarykey"`
	InvoiceID uint     `gorm:"index"`
	Invoice   *Invoice `json:",omitempty"`
	ArticleID uint     `gorm:"index"`
	Article   *Article `json:",omitempty"`
	Quantity  uint     `gorm:"not null"`
	UnitPrice uint     `gorm:"not null"`
}

func (InvoiceLine) TableName() string {
	return "InvoiceLine"
}
