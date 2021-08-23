package data

type Product struct {
	ProductID uint64  `gorm:"primaryKey;foreignKey:ProductID" json:"productID"`
	Name      string  `json:"name"`
	PriceCZK  float64 `json:"priceCZK"`
}
