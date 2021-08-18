package data

type Customer struct {
	CustomerID uint64 `gorm:"primaryKey"`
	Name       string
	Email      string
}
