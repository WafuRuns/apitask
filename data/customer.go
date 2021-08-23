package data

type Customer struct {
	CustomerID uint64 `gorm:"primaryKey;foreignKey:CustomerID" json:"customerID"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}
