package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan pola bintang
func printStars(n int, current int) {
	// Basis rekursi: berhenti jika current lebih besar dari n
	if current > n {
		return
	}

	// Cetak bintang sebanyak nilai current
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	// Panggilan rekursif untuk mencetak baris berikutnya
	printStars(n, current+1)
}

func main() {
	var jumlahTest int
	fmt.Print("Masukkan jumlah kasus uji: ")
	fmt.Scan(&jumlahTest)

	// Array untuk menyimpan setiap nilai input
	inputs := make([]int, jumlahTest)

	// Mengambil input untuk setiap kasus uji
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("Masukan #%d: ", i+1)
		fmt.Scan(&inputs[i])
	}

	// Memproses dan menampilkan hasil untuk setiap kasus uji
	fmt.Println("\nKeluaran:")
	fmt.Printf("No\tMasukan\tKeluaran\n")
	for i := 0; i < jumlahTest; i++ {
		fmt.Printf("%d\t%d\n", i+1, inputs[i])
		printStars(inputs[i], 1)
		fmt.Println()
	}
}
