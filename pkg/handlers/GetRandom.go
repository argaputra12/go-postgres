package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sort"
)

/*
	Buatlah fungsi generate string random huruf dan angka sebanyak 100 dengan
	ketentuan (50 huruf dan 50 angka)
	a. Tampilkan total dari hasil generate string diatas dengan ketentuan :
	i. Berapa banyak total huruf
	ii. Berapa banyak total huruf vokal
	iii. Berapa banyak total angka
	iv. Berapa banyak total angka genap

b. Urutkan hasil generate random tadi dari angka terbesar ke terkecil , dan

	huruf terkecil ke terbesar dengan ketentuan duplikasi dihapus angka atau
	abjad, ex : 5,4,3,2,1,0,a,b,c,d,e
	c. Urutkan hasil generate random tadi dari angka terbesar ke terkecil , diikuti
	dengan huruf terkecil ke terbesar, ex : 9a,9b,8c,7c,7c,6c,d,d
*/
type Total struct {
	TotalHuruf      int `json:"total_huruf"`
	TotalHurufVokal int `json:"total_huruf_vokal"`
	TotalAngka      int `json:"total_angka"`
	TotalAngkaGenap int `json:"total_angka_genap"`
}

type Random struct {
	Total   Total  `json:"total"`
	Sorted  string `json:"sorted"`
	Sorted2 string `json:"sorted2"`
}

func (h handler) GetRandom(w http.ResponseWriter, r *http.Request) {
	// generate random string
	randomString := generateRandomString()

	// a. count total huruf, huruf vokal, angka, angka genap
	total := countTotal(randomString)

	// b. sort
	sorted := sortString(randomString)

	// c. sort
	sorted2 := sortString2(randomString)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Random{Total: total, Sorted: sorted, Sorted2: sorted2})
	// w.Write([]byte(`{"total":` + total + `,"sorted":` + sorted + `,"sorted2":` + sorted2 + `}`))
}

// generate random string with 50 letters and 50 numbers
func generateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")
	b := make([]rune, 100)
	for i := range b {
		if i < 50 {
			b[i] = letters[rand.Intn(len(letters))]
		} else {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(b)
}

func countTotal(s string) Total {
	var total Total
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			total.TotalHuruf++
			if r == 'a' || r == 'i' || r == 'u' || r == 'e' || r == 'o' || r == 'A' || r == 'I' || r == 'U' || r == 'E' || r == 'O' {
				total.TotalHurufVokal++
			}
		} else {
			total.TotalAngka++
			if int(r)%2 == 0 {
				total.TotalAngkaGenap++
			}
		}
	}
	total = Total{
		TotalHuruf:      total.TotalHuruf,
		TotalHurufVokal: total.TotalHurufVokal,
		TotalAngka:      total.TotalAngka,
		TotalAngkaGenap: total.TotalAngkaGenap,
	}

	return total

}

func sortString(s string) string {
	var letters []rune
	var numbers []rune

	// sort letters and numbers
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			letters = append(letters, r)
		} else {
			numbers = append(numbers, r)
		}
	}

	// sort letters
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	// sort numbers
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// remove duplicate
	letters = removeDuplicate(letters)
	numbers = removeDuplicate(numbers)

	return string(numbers) + string(letters)
}

func sortString2(s string) string {
	var letters []rune
	var numbers []rune

	// sort letters and numbers
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			letters = append(letters, r)
		} else {
			numbers = append(numbers, r)
		}
	}

	// sort letters
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	// sort numbers
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// change format to 9a, 9b, 8c, 7c, 7c, 6c
	var result []rune
	for i := 0; i < len(numbers); i++ {
		result = append(result, numbers[i])
		if i < len(letters) {
			result = append(result, letters[i])
		}
	}

	return string(result)
}

func removeDuplicate(s []rune) []rune {
	unique := map[rune]bool{}
	result := []rune{}
	for i := range s {
		if !unique[s[i]] {
			unique[s[i]] = true
			result = append(result, s[i])
		}
	}
	return result
}
