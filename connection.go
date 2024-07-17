package main

import	(
	"fmt"
	"log"

	"github.com/argaputra12/go-postgres/models"
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

// // Employee represents the employee model
// type Employee struct {
// 	ID              int `gorm:"primaryKey"`
// 	FirstName       string
// 	LastName        string
// 	HireDate        string
// 	TerminationDate *string
// 	Salary          int
// }

// // AnnualReviews represents the annual review model
// type AnnualReviews struct {
// 	ID         int `gorm:"primaryKey"`
// 	EmployeeID int
// 	Employee   Employee `gorm:"foreignKey:EmployeeID"`
// 	ReviewDate string
// }


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

	// Insert employees data
	InsertEmployees(db)
	fmt.Println("Employees data inserted successfully")

	// Insert annual reviews data
	InsertAnnualReviews(db)
	fmt.Println("Annual reviews data inserted successfully")
	
}

// insert employees data from data.go to database
func InsertEmployees(db *gorm.DB) {
	employees := []models.Employee{
		{
			FirstName:       "Bob",
			LastName:        "Smith",
			HireDate:        "2009-06-20",
			TerminationDate: "2016-01-01",
			Salary:          10000,
		},
		{
			FirstName:       "Joe",
			LastName:        "Jarrod",
			HireDate:        "2010-02-12",
			TerminationDate: nil,
			Salary:          20000,
		},
		{
			FirstName:       "Nancy",
			LastName:        "Soley",
			HireDate:        "2012-03-14",
			TerminationDate: nil,
			Salary:          30000,
		},
		{
			FirstName:       "Keith",
			LastName:        "Widjaja",
			HireDate:        "2013-09-10",
			TerminationDate: "2014-01-01",
			Salary:          20000,
		},
		{
			FirstName:       "Kelly",
			LastName:        "Smalls",
			HireDate:        "2013-09-10",
			TerminationDate: nil,
			Salary:          20000,
		},
		{
			FirstName:       "Frank",
			LastName:        "Nguyen",
			HireDate:        "2015-04-10",
			TerminationDate: "2015-05-01",
			Salary:          60000,
		},
	}

	for _, employee := range employees {
		db.Create(&employee)
	}
}

// insert annual reviews data from data.go to database
func InsertAnnualReviews(db *gorm.DB) {
	annualReviews := []models.AnnualReviews{
		{
			ID:         10,
			EmployeeID: 1,
			ReviewDate: "2016-01-01",
		},
		{
			ID:         20,
			EmployeeID: 2,
			ReviewDate: "2016-04-12",
		},
		{
			ID:         30,
			EmployeeID: 10,
			ReviewDate: "2016-05-01",
		},
	}

	for _, review := range annualReviews {
		db.Create(&review)
	}
}

