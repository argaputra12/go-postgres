package seeder

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SeedAnnualReviews(db *gorm.DB) {
	// drop existing data
	db.Exec("DELETE FROM annual_reviews")

	// seed data with disabled foreign key constraint
	db.Exec("SET session_replication_role = 'replica'")
	
	for _, annualReview := range annualReviews {
		if err := db.Create(&annualReview).Error; err != nil {
			log.Fatalf("Error seeding annual reviews: %v", err)
		}
	}

	// enable foreign key constraint
	db.Exec("SET session_replication_role = 'origin'")

	fmt.Println("Annual reviews seeded")
}
