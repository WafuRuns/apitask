package data

type Customer struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Email string
}
