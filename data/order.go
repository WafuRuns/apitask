package data

import "time"

type Order struct {
	OrderID    uint64      `gorm:"primaryKey"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	Customer   Customer    `gorm:"references:CustomerID"`
	CustomerID uint64
	CreatedAt  time.Time
	Confirmed  bool `gorm:"default:false"`
	Reminded   bool `gorm:"default:false"`
}
