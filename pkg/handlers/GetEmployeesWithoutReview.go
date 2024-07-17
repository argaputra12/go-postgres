package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/argaputra12/go-postgres/pkg/models"
)

/*
Tampilkan dalam bentuk response JSON nama semua karyawan yang tidak pernah
 memiliki ulasan yang disortir oleh HireDate
*/

func (h handler) GetEmployeesWithoutReview(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	h.DB.Where("id NOT IN (SELECT employee_id FROM annual_reviews)").Order("hire_date").Find(&employees)

	// save to .txt
	byteSlice, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saveJSONtoText(byteSlice, "contoh3.txt")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
