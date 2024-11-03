package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak angka menurun
func printDescending(n int, current int) {
	if current < 1 {
		return
	}
	fmt.Printf("%d ", current)
	printDescending(n, current-1)
}

// Fungsi rekursif untuk mencetak angka menaik
func printAscending(n int, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", current)
	printAscending(n, current+1)
}

// Fungsi untuk mencetak pola lengkap dari angka menurun ke menaik
func printPattern(n int) {
	printDescending(n, n)
	printAscending(n, 2)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah kasus uji: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan nilai input
	inputs := make([]int, jumlahTest)

	// Mengambil input untuk setiap kasus uji
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d: ", i+1)
		fmt.Scan(&inputs[i])
	}

	// Memproses dan menampilkan hasil untuk setiap kasus uji
	fmt.Println("\nNo\tMasukan\tKeluaran")
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("%d\t%d\t", i+1, inputs[i])
		printPattern(inputs[i])
		fmt.Println()
	}
}
