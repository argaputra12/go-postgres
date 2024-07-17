package seeder

import (
	"fmt"
	
	"github.com/argaputra12/go-postgres/pkg/db"
)

func SeedDatabase() {
	db := db.Init()
	SeedEmployees(db)
	SeedAnnualReviews(db)
	fmt.Println("Seeding completed")
}
