package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan ganjil
func printOddNumbers(n int, current int) {
	// Basis rekursi: berhenti jika current lebih besar dari n
	if current > n {
		return
	}

	// Jika current adalah bilangan ganjil, cetak nilai current
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	// Panggilan rekursif untuk memeriksa bilangan berikutnya
	printOddNumbers(n, current+1)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah kasus uji: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan setiap nilai input
	inputs := make([]int, jumlahTest)

	// Mengambil input dari pengguna untuk setiap kasus uji
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d: ", i+1)
		fmt.Scan(&inputs[i])
	}

	// Memproses dan menampilkan hasil untuk setiap kasus uji
	fmt.Println("\nNo\tMasukan\tKeluaran")
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("%d\t%d\t", i+1, inputs[i])
		printOddNumbers(inputs[i], 1) // Mulai dari 1 untuk mencetak bilangan ganjil
		fmt.Println()                 // Pindah baris untuk kasus uji berikutnya
	}
}
