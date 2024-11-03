package main

import (
	"fmt"
)

// Fungsi rekursif untuk menemukan dan mencetak faktor dari bilangan
func findFactors(n int, current int) {
	// Basis rekursi: berhenti jika current lebih besar dari n
	if current > n {
		return
	}

	// Jika current adalah faktor dari n, cetak nilai current
	if n%current == 0 {
		fmt.Printf("%d ", current)
	}

	// Panggilan rekursif untuk mengecek faktor berikutnya
	findFactors(n, current+1)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah kasus uji: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan nilai input
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
		findFactors(inputs[i], 1)
		fmt.Println()
	}
}
