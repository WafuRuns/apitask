package data

type OrderItem struct {
	OrderItemID uint64  `gorm:"primaryKey"`
	Product     Product `gorm:"references:ProductID"`
	ProductID   uint64
	OrderID     uint64
	Amount      uint64
}
