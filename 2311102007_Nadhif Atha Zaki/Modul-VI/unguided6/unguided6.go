package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung nilai pangkat
func power(base int, exponent int) int {
	// Basis rekursi: jika eksponen adalah 0, kembalikan 1
	if exponent == 0 {
		return 1
	}

	// Hitung pangkat dengan mengalikan base dengan hasil rekursif
	return base * power(base, exponent-1)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah kasus uji: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan nilai base dan eksponen setiap input
	baseInputs := make([]int, jumlahTest)
	exponentInputs := make([]int, jumlahTest)

	// Mengambil input dari pengguna untuk setiap kasus uji
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d (x y): ", i+1)
		fmt.Scan(&baseInputs[i], &exponentInputs[i])
	}

	// Memproses dan menampilkan hasil untuk setiap kasus uji
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
