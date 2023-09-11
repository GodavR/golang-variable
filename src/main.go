// variable
// Berikut adalah contoh cara pembuatan variabel yang tipe datanya harus ditulis.
// Silakan tulis pada project baru atau pada project yang sudah ada, bebas.
// Pastikan saja untuk setiap project baru untuk tidak lupa inisialisasi project
// menggunakan command go mod init <nama-project> .
package main

import "fmt"

func main() {
	var firstName string = "David"
	var lastname string = "Rinaldi"

	fmt.Printf("Halo %s %s! \n", firstName, lastname)
}

// Keyword var di atas digunakan untuk deklarasi variabel, contohnya bisa dilihat
// pada firstName dan lastName .
// Nilai variabel firstName diisi langsung ketika deklarasi, berbeda dibanding
// lastName yang nilainya diisi setelah baris kode deklarasi, hal seperti ini
// diperbolehkan di Go.

//dapat juga deklarasi variabel dengan menggunakan var
// var firstName string = "john"
// lastName := "wick"
// fmt.Printf("halo %s %s!\n", firstName, lastName)

// Variabel lastName dideklarasikan dengan menggunakan metode type inference.
// Penandanya tipe data tidak dituliskan pada saat deklarasi. Pada penggunaan
// metode ini, operand = harus diganti dengan := dan keyword var dihilangkan.
// Tipe data lastName secara otomatis akan ditentukan menyesuaikan value atau
// nilai-nya. Jika nilainya adalah berupa string maka tipe data variabel adalah
// string . Pada contoh di atas, nilainya adalah string "wick" .
// Diperbolehkan untuk tetap menggunakan keyword var pada saat deklarasi
// meskipun tanpa menuliskan tipe data, dengan ketentuan tidak menggunakan
// tanda := , melainkan tetap menggunakan = .

// // menggunakan var, tanpa tipe data, menggunakan perantara "="
// var firstName = "john"
// // tanpa var, tanpa tipe data, menggunakan perantara ":="
// lastName := "wick"

// Kedua deklarasi di atas maksudnya sama. Silakan pilih yang nyaman di hati.
// Tanda := hanya digunakan sekali di awal pada saat deklarasi. Untuk
// assignment nilai selanjutnya harus menggunakan tanda = , contoh:
// lastName := "wick"
// lastName = "ethan"
// lastName = "bourne"
// Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di
// dalam blok fungsi.
