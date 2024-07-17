package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/argaputra12/go-postgres/pkg/models"
)

/*
	Tampilkan dalam bentuk response JSON nama semua karyawan yang masih bekerja
 untuk perusahaan dengan nama belakang dimulai dengan "Smith" diurutkan
 berdasarkan nama belakang kemudian nama depan.
*/

func (h handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	h.DB.Where("termination_date IS NULL").Where("last_name LIKE ?", "Smith%").Order("last_name, first_name").Find(&employees)

	// save to .txt
	byteSlice, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saveJSONtoText(byteSlice, "contoh2.txt")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
