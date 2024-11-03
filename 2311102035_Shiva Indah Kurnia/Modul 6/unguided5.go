package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan ganjil
func printOddNumbers(n int, current int) {
	// Basis rekursi: berhenti jika current > n
	if current > n {
		return
	}
	// Jika current adalah bilangan ganjil, cetak
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	// Panggil rekursif dengan bilangan selanjutnya
	printOddNumbers(n, current+1)
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
		printOddNumbers(inputs[i], 1)
		// Mulai dari 1
		fmt.Println()
		// Pindah baris untuk test case berikutnya
	}
}
