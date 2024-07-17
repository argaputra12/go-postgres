package models

import (
	"time"
)

// Employee represents the employee model
type Employee struct {
	ID              int `gorm:"primaryKey"` 
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	HireDate        time.Time `json:"hire_date"`
	TerminationDate *time.Time `json:"termination_date"`
	Salary          int `json:"salary"`
}
