package models

// Employee represents the employee model
type Employee struct {
	ID              int `gorm:"primaryKey"`
	FirstName       string
	LastName        string
	HireDate        string
	TerminationDate *string
	Salary          int
}

