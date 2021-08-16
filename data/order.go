package data

type Order struct {
	ID       uint64      `gorm:"primaryKey"`
	Items    []OrderItem `gorm:"foreignKey:ID"`
	Customer Customer    `gorm:"foreignKey:ID"`
}
