package models

import	(
	"time"
)

// AnnualReviews represents the annual review model
type AnnualReviews struct {
	ID         int `gorm:"primaryKey"`
	EmployeeID int `json:"employee_id"`
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	ReviewDate time.Time `json:"review_date"`
}

