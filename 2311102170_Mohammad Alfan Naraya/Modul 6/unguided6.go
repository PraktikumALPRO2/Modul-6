package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung pangkat
func power(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}

	return base * power(base, exponent-1)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah test case: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan input
	baseInputs := make([]int, jumlahTest)
	exponentInputs := make([]int, jumlahTest)

	// Membaca input
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d (x y): ", i+1)
		fmt.Scan(&baseInputs[i], &exponentInputs[i])
	}

	// Memproses dan menampilkan hasil
	fmt.Println("\nNo\tMasukan\tKeluaran")
	for i := 0; i < jumlahTest; i++ {
		result := power(baseInputs[i], exponentInputs[i])
		fmt.Printf("%d\t%d %d\t%d\n",
			i+1,
			baseInputs[i],
			exponentInputs[i],
			result)
	}
}
