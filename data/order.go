package data

import "time"

type Order struct {
	OrderID    uint64      `gorm:"primaryKey" json:"orderID"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Customer   Customer    `gorm:"references:CustomerID" json:"customer"`
	CustomerID uint64      `json:"-"`
	CreatedAt  time.Time   `json:"createdAt"`
	Confirmed  bool        `gorm:"default:false" json:"confirmed"`
	Reminded   bool        `gorm:"default:false" json:"reminded"`
}
