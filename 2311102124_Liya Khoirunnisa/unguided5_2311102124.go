/* Liya Khoirunnisa - 2311102124*/
package main

import (
	"fmt"
)

// Fungsi untuk menampilkan bilangan ganjil
func bilanganGanjil(n int) {
	// Jika n kurang dari 1, hentikan rekursif
	if n < 1 {
		return
	}
	// Panggil fungsi untuk menampilkan bilangan ganjil sebelumnya
	bilanganGanjil(n - 1)
	
	// Jika n adalah bilangan ganjil, cetak n
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	// Deklarasi variabel
	var N int

	// Input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Memanggil fungsi untuk menampilkan bilangan ganjil
	fmt.Print("Bilangan ganjil: ")
	bilanganGanjil(N)
}