package data

type Order struct {
	OrderID    uint64      `gorm:"primaryKey"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	Customer   Customer    `gorm:"references:CustomerID"`
	CustomerID uint64
	Confirmed  bool `gorm:"default:false"`
}
