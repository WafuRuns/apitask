package data

type OrderItem struct {
	OrderItemID uint64  `gorm:"primaryKey"`
	Product     Product `gorm:"foreignKey:ProductID"`
	Amount      uint64
}
