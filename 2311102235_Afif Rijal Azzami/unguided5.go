package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan ganjil dari 1 hingga n
func cetakGanjil(n, i int) {
	if i > n {
		return // Base case: jika i lebih besar dari n, hentikan rekursi
	}

	if i%2 != 0 {
		fmt.Print(i, " ") // Cetak i jika ganjil
	}
	cetakGanjil(n, i+1) // Rekursif untuk mengecek bilangan berikutnya
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)
	cetakGanjil(N, 1)
	fmt.Println()
}
