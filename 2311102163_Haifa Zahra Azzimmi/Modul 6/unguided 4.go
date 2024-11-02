package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan barisan dari N hingga 1
func turunNaik(n, current int) {
	// Cetak nilai saat ini
	fmt.Print(current, " ")

	// Jika current mencapai 1, mulai mencetak kembali ke N
	if current == 1 {
		return
	}

	// Rekursi untuk turun ke 1
	turunNaik(n, current-1)

	// Cetak kembali saat naik lagi dari 2 hingga N
	fmt.Print(current, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali ke %d: ", N, N)
	turunNaik(N, N)
	fmt.Println()
}
