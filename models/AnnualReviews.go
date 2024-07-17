package models

// AnnualReviews represents the annual review model
type AnnualReviews struct {
	ID         int `gorm:"primaryKey"`
	EmployeeID int
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	ReviewDate string
}

