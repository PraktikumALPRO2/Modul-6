package main

import (
	"fmt"
)

// Fungsi rekursif untuk menurun
func printDescending(n int, current int) {
	if current < 1 {
		return
	}

	fmt.Printf("%d ", current)

	printDescending(n, current-1)
}

// Fungsi rekursif untuk menaik
func printAscending(n int, current int) {
	if current > n {
		return
	}

	fmt.Printf("%d ", current)

	printAscending(n, current+1)
}

// Fungsi untuk mencetak pola lengkap
func printPattern(n int) {
	printDescending(n, n)

	printAscending(n, 2)
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
		printPattern(inputs[i])
		fmt.Println()
	}
}
