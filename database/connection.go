package migration

import (
	"fmt"
	"log"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "interview-nayaka"
)

// Employee represents the employee model
type Employee struct {
	ID              int `gorm:"primaryKey"`
	FirstName       string
	LastName        string
	HireDate        string
	TerminationDate *string
	Salary          int
}

// AnnualReviews represents the annual review model
type AnnualReviews struct {
	ID         int `gorm:"primaryKey"`
	EmployeeID int
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	ReviewDate string
}


func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	// Automigrate the schema
	err = db.AutoMigrate(&Employee{}, &AnnualReviews{})
	if err != nil {
		log.Fatalf("Error during migration: %v\n", err)
	}

	fmt.Println("Database migrated successfully")
}

// insert employees data from data.go to database
func InsertEmployeesData() {
	
