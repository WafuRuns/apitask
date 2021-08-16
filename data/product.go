package data

type Product struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string
	PriceCZK float64
}
