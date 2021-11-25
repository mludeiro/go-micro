package entity

type Client struct {
	ID      uint   `gorm:"primarykey"`
	Name    string `gorm:"not null;default:null"`
	Address string ``
}

func (Client) TableName() string {
	return "Client"
}
