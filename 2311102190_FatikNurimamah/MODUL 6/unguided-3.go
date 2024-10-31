package main

import (
	"fmt"
)

func cariFaktor(n, x int, daftarfaktor []int) []int {
	// Kondisi berhenti: jika x melebihi n, kembalikan daftar faktor
	if x > n {
		return daftarfaktor
	}
	// Cek apakah x adalah faktor dari n
	if n%x == 0 {
		daftarfaktor = append(daftarfaktor, x)
	}
	// Panggilan rekursif untuk memeriksa angka berikutnya
	return cariFaktor(n, x+1, daftarfaktor)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&n)

	daftarfaktor := cariFaktor(n, 1, []int{})
	fmt.Println("Faktor-faktor dari", n, "adalah:", daftarfaktor)
}
