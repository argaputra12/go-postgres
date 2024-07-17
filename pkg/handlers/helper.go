package handlers

import (
	"log"
	"os"
)

/*
Buatlah fungsi untuk menyimpan hasil format JSON setiap nomor 2– 5 ke dalam bentuk
 file txt dengan penamaan contoh sebagai berikut contoh2.txt, contoh3.txt,dst
*/

func saveJSONtoText(jsonResponse []byte, fileName string) {
	file, err := os.Create("assets/" + fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonResponse)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}
