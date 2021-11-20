package entity

import "time"

type Invoice struct {
	ID          uint          `gorm:"primarykey"`
	CreatedAt   time.Time     ``
	ClientID    uint          `gorm:"index"`
	Client      *Client       `json:",omitempty"`
	Amount      uint          `gorm:"not null;default:0"`
	Closed      bool          `gorm:"not null;default:false"`
	InvoiceLine []InvoiceLine ``
}

func (Invoice) TableName() string {
	return "Invoice"
}
