package handlers

import (
	"io"
	"net/http"
	"os"
)

/*
 Buatlah fungsi yang menerima masukan sebuah string nama file untuk membuka dan
 menampilkan file pada soal point 6 ke dalam bentuk response JSON. ex : masukan
 string “contoh3.txt”, keluaran adalah data json dari file contoh3.txt
*/

func (h handler) GetTextFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")

	// read file
	file, err := os.Open("assets/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// read file
	byteSlice, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteSlice)
}
