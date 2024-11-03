package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bintang
func printStars(n int, current int) {
	// Basis rekursi: jika current > n, berhenti
	if current > n {
		return
	}
	// Cetak bintang sebanyak current
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	// Panggil rekursif untuk baris selanjutnya
	printStars(n, current+1)
}
func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah test case: ")
	fmt.Scan(&jumlahTest)
	// Array untuk menyimpan hasil input
	inputs := make([]int, jumlahTest)
	// Membaca input
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d: ", i+1)
		fmt.Scan(&inputs[i])
	}
	// Memproses setiap test case
	fmt.Println("\nKeluaran:")
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("No\tMasukan\tKeluaran\n")
		fmt.Printf("%d\t%d\t\n", i+1, inputs[i])
		printStars(inputs[i], 1)
		fmt.Println()
	}
}
