package data

type OrderItem struct {
	ID      uint64 `gorm:"primaryKey"`
	Product Product
	Amount  uint64
}
