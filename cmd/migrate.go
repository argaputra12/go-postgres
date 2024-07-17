package main

import (
	"fmt"
	"log"
	"github.com/argaputra12/go-postgres/pkg/db"
	"github.com/argaputra12/go-postgres/pkg/models"
	"github.com/argaputra12/go-postgres/pkg/db/seeder"
)

func main() {
	db := db.Init()

	// delete tables if exists
	db.Migrator().DropTable(&models.Employee{}, &models.AnnualReviews{})
	fmt.Println("Tables dropped")

	// migrate tables
	err := db.AutoMigrate(&models.Employee{}, &models.AnnualReviews{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	fmt.Println("Database migrated")

	// seeder
	seeder.SeedDatabase()
	fmt.Println("Seeding completed")
}
