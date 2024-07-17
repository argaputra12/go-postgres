package handlers

import (
	"encoding/json"
	"net/http"
)

/* array = [“Bandung”,”Cimahi”,”Ambon”,”Jayapura”,”Makasar”] .Buatlah fungsi yang
menerima masukan sebuah string nama kota dengan ketentuan berikut :
a. Mengembalikan Boolean TRUE jika kota tersebut ada di dalam array.
b. Mengembalikan Boolean FALSE , dan String “Saran Kota” jika kota tersebut
tidak ada di dalam array serta terdapat field. Ketentuan value saran kota
adalah dengan ketentuan berikut :
a. Kota dalam array dimunculkan sesuai dengan abjad pertama yang
sama. Contoh : Input “Bogor” hasil yang di keluarkan “Saran Kota =
Bandung”
b. Kota dalam array dimunculkan sesuai dengan abjad terakhir yang
sama. Contoh : Input “Bogor” hasil yang dikeluarkan “Saran Kota =
Makasar”
c. Maka Ketentuan Poin A dan B digabung dapat menghasilkan
keluaran seperti berikut. Contoh : Input “Bogor” hasil yang
dikeluarkan “Saran Kota = Bandung , Makasar”
*/

var cities = []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}

func (h handler) GetCity(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	var response struct {
		IsExist    bool   `json:"is_exist"`
		Suggestion string `json:"suggestion"`
	}

	for _, c := range cities {
		if c == city {
			response.IsExist = true
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// suggestion
	var suggestion string
	for _, c := range cities {
		if c[0] == city[0] {
			suggestion += c + ", "
		}
		if c[len(c)-1] == city[len(city)-1] {
			suggestion += c + ", "
		}
	}

	if suggestion == "" {
		response.IsExist = false
		response.Suggestion = "Saran Kota tidak ditemukan"
	} else {
		response.IsExist = false
		response.Suggestion = "Saran Kota = " + suggestion[:len(suggestion)-2]
	}

	json.NewEncoder(w).Encode(response)
}
