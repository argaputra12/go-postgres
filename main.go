package main

import (
	"log"
	"net/http"

	"github.com/argaputra12/go-postgres/pkg/db"
	"github.com/argaputra12/go-postgres/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	// nomor 2
	router.HandleFunc("/employees", h.GetAllEmployees).Methods("GET")

	// nomor 3
	router.HandleFunc("/employees/without-review", h.GetEmployeesWithoutReview).Methods("GET")

	// nomor 4
	router.HandleFunc("/employees/hire-date", h.GetHireDate).Methods("GET")

	// nomor 5
	router.HandleFunc("/employees/expected-salary", h.GetExpectedSalary).Methods("GET")

	// nomor 7
	router.HandleFunc("/employees/text-file", h.GetTextFile).Methods("GET")

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
