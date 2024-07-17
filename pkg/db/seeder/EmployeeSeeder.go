package seeder

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SeedEmployees(db *gorm.DB) {
	// drop existing data
	db.Exec("DELETE FROM employees")

	// seed data
	for _, employee := range employees {
		if err := db.Create(&employee).Error; err != nil {
			log.Fatalf("Error seeding employees: %v", err)
		}
	}

	fmt.Println("Employees seeded")
}



