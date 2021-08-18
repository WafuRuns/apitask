package data

type Product struct {
	ProductID uint64 `gorm:"primaryKey"`
	Name      string
	PriceCZK  float64
}
