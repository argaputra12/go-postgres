package handlers

import (
	"encoding/json"
	"net/http"

	// "gorm.io/datatypes"
	"github.com/argaputra12/go-postgres/pkg/models"
)

/*
Tampilkan dalam bentuk response JSON untuk menghitung selisih (dalam hari)
 antara karyawan perekrutan tanggal paling awal dan paling perekrutan tanggal
 paling akhir yang bekerja di perusahaan
*/

type HireDateDifference struct {
	EarliestEmployee models.Employee `json:"earliest_employee"`
	LatestEmployee   models.Employee `json:"latest_employee"`
	Difference       int             `json:"difference"`
}

func (h handler) GetHireDate(w http.ResponseWriter, r *http.Request) {
	var earliestEmployee models.Employee
	var latestEmployee models.Employee

	// get employee with earliest hire date
	if err := h.DB.Order("hire_date").First(&earliestEmployee).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get employee with latest hire date
	if err := h.DB.Order("hire_date desc").First(&latestEmployee).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// calculate difference
	difference := int(latestEmployee.HireDate.Sub(earliestEmployee.HireDate).Hours() / 24)

	response := HireDateDifference{
		EarliestEmployee: earliestEmployee,
		LatestEmployee:   latestEmployee,
		Difference:       difference,
	}

	// save to .txt
	byteSlice, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saveJSONtoText(byteSlice, "contoh4.txt")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func formatDate(d datatypes.Date) string {
// 	return time.Time(d).Format("2006-01-02")
// }
