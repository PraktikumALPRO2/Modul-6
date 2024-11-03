package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencari faktor
func findFactors(n int, current int) {
	// Basis rekursi: jika current > n, berhenti
	if current > n {
		return
	}
	// Jika current adalah faktor dari n
	if n%current == 0 {
		fmt.Printf("%d ", current)
	}
	// Panggil rekursif untuk bilangan selanjutnya
	findFactors(n, current+1)
}
func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah test case: ")
	fmt.Scan(&jumlahTest)
	// Array untuk menyimpan input
	inputs := make([]int, jumlahTest)
	// Membaca input
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d: ", i+1)
		fmt.Scan(&inputs[i])
	}
	// Memproses dan menampilkan hasil
	fmt.Println("\nNo\tMasukan\tKeluaran")
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("%d\t%d\t", i+1, inputs[i])
		findFactors(inputs[i], 1)
		fmt.Println()
	}
}
