package data

type Product struct {
	ProductID uint64 `gorm:"primaryKey;foreignKey:ProductID"`
	Name      string
	PriceCZK  float64
}
