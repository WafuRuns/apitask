package data

type OrderItem struct {
	OrderItemID uint64  `gorm:"primaryKey" json:"orderItemID"`
	Product     Product `gorm:"references:ProductID" json:"product"`
	ProductID   uint64  `json:"-"`
	OrderID     uint64  `json:"-"`
	Amount      uint64  `json:"amount"`
}
