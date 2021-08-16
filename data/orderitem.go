package data

type OrderItem struct {
	ID      uint64  `gorm:"primaryKey"`
	Product Product `gorm:"foreignKey:ID"`
	Amount  uint64
}
