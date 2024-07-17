package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/argaputra12/go-postgres/pkg/models"
)

/*
 Tampilkan dalam bentuk response JSON untuk menampilkan perkiraan kenaikan gaji
 semua karyawan jika setiap tahunya kenaikan gaji adalah 15% dari 2009 hingga 2016,
 pada kolom salary di anggap sebagai gaji awal perekrutan, serta tampilkan berapa
 total ulasan yang dimiliki setiap karyawan, urutkan berdasarkan gaji karyawan paling
 terbesar pada tahun 2016 dan ulasan yang paling sedikit.
*/

type ExpectedSalary struct {
	Employee    models.Employee `json:"employee"`
	TotalReview int             `json:"total_review"`
	FinalSalary int             `json:"final_salary"`
}

// add expected salary to employee
func (h handler) GetExpectedSalary(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	var expectedSalaries []ExpectedSalary

	h.DB.Order("salary").Find(&employees)

	for _, employee := range employees {
		// get total review
		var totalReview int64
		h.DB.Model(&models.AnnualReviews{}).Where("employee_id = ?", employee.ID).Count(&totalReview)

		// calculate expected salary
		expectedSalary := float64(employee.Salary)
		for i := employee.HireDate.Year(); i < 2016; i++ {
			expectedSalary += expectedSalary * 0.15
		}

		expectedSalaries = append(expectedSalaries, ExpectedSalary{
			Employee:    employee,
			TotalReview: int(totalReview),
			FinalSalary: int(expectedSalary), // Convert float64 back to int
		})
	}

	// Sort by final salary in descending order and by total reviews in ascending order
	sort.Slice(expectedSalaries, func(i, j int) bool {
		if expectedSalaries[i].FinalSalary == expectedSalaries[j].FinalSalary {
			return expectedSalaries[i].TotalReview < expectedSalaries[j].TotalReview
		}
		return expectedSalaries[i].FinalSalary > expectedSalaries[j].FinalSalary
	})

		// save to .txt
		byteSlice, err := json.Marshal(expectedSalaries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		saveJSONtoText(byteSlice, "contoh5.txt")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expectedSalaries)
}
