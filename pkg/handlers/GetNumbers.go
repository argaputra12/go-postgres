package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

/*
array = [9,1,6,4,8,6,6,3,8,2,3,3,4,1,8,2]

	a. Buatlah fungsi untuk mengurutkan data dalam array dari terkecil ke terbesar
	dengan kondisi angka yang ditampilkan tidak boleh duplikat, ex : 1,2,3,4,6,8,9
	b. Buatlah fungsi untuk menampilkan total duplikat di setiap,
	ex : 1[2],2[2],3[3],4[2],6[3],8[2],9[1]
	c. Buatlah fungsi untuk menghapus value di dalam array
	sesuai inputan, ex : input = 9,1,6,4 maka hasil yang
	diharapkan adalah 8,3,8,2,3,3,8,2
	d. Buatlah fungsi untuk menginputkan dan menjumlah kan
	sesuai inputan di setiap value dalam array maksimal
	menjadi 10, ex : input = 15 maka hasil yang diharpakan
	adalah 10,10,10,5,8,6,6,3,8,2,3,3,4,1,8,2
*/
var numbers = []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}

type Numbers struct {
	SortedArray []int       `json:"sorted_array"`
	TotalDup    map[int]int `json:"total_dup"`
	Deleted     []int       `json:"deleted"`
	Sum         []int       `json:"sum"`
}

func (h handler) GetNumbers(w http.ResponseWriter, r *http.Request) {
	var response Numbers

	// a. sort array
	willSortNumbers := make([]int, len(numbers))
	copy(willSortNumbers, numbers)
	sorted := sortArray(willSortNumbers)
	response.SortedArray = sorted

	// b. total dup
	willDupNumbers := make([]int, len(numbers))
	copy(willDupNumbers, numbers)
	totalDup := countDuplicates(willDupNumbers)
	response.TotalDup = totalDup

	// c. delete value
	deleteValue := r.URL.Query().Get("delete")
	willRemoveNumbers := make([]int, len(numbers))
	copy(willRemoveNumbers, numbers)
	valuesToRemove, err := stringToIntArray(deleteValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if deleteValue != "" {
		removed := removeValues(willRemoveNumbers, valuesToRemove)
		response.Deleted = removed
	} else {
		response.Deleted = nil
	}

	// d. sum
	sumValue := r.URL.Query().Get("sum")
	willSumNumbers := make([]int, len(numbers))
	copy(willSumNumbers, numbers)
	parsedSumValue, err := stringToInt(sumValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if sumValue != "" {
		result := addToArrayAndLimitMax(willSumNumbers, parsedSumValue, 10)
		response.Sum = result
	} else {
		response.Sum = nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sortArray(arr []int) []int {
	// Sort array
	sort.Ints(arr)

	// Remove duplicate
	unique := map[int]bool{}
	result := []int{}
	for i := range arr {
		if !unique[arr[i]] {
			unique[arr[i]] = true
			result = append(result, arr[i])
		}
	}
	return result
}

func countDuplicates(arr []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	return counts
}

func removeValues(arr []int, valuesToRemove []int) []int {
	// Buat map untuk nilai yang ingin dihapus agar pencarian lebih cepat
	removeMap := make(map[int]bool)
	for _, val := range valuesToRemove {
		removeMap[val] = true
	}

	// Buat slice baru untuk menampung nilai-nilai yang tidak dihapus
	var result []int
	for _, val := range arr {
		if !removeMap[val] {
			result = append(result, val)
		}
	}

	return result
}

func addToArrayAndLimitMax(arr []int, valueToAdd int, max int) []int {

	// Tambahkan value ke setiap elemen array
	for i := range arr {
		diff := difference(arr[i], max)
		if valueToAdd-diff > 0 {
			arr[i] = max
			valueToAdd -= diff
		} else {
			arr[i] += valueToAdd
			valueToAdd = 0
		}

	}
	return arr
}

func stringToIntArray(input string) ([]int, error) {
	// Pisahkan string berdasarkan koma
	stringArray := strings.Split(input, ",")

	// Buat array untuk menampung hasil konversi
	intArray := make([]int, len(stringArray))

	// Konversi setiap elemen string ke integer
	for i, str := range stringArray {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		intArray[i] = num
	}

	return intArray, nil
}

func stringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func difference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
