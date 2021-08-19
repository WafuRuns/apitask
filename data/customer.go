package data

type Customer struct {
	CustomerID uint64 `gorm:"primaryKey;foreignKey:CustomerID"`
	Name       string
	Email      string
}
