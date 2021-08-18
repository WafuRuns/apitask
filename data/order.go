package data

type Order struct {
	OrderID   uint64      `gorm:"primaryKey"`
	Items     []OrderItem `gorm:"foreignKey:OrderItemID"`
	Customer  Customer    `gorm:"foreignKey:CustomerID;unique:false"`
	Confirmed bool        `gorm:"default:false"`
}
