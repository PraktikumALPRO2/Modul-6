package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan ganjil hingga N
func tampilGanjil(n, current int) {
	// Jika current melebihi n, hentikan rekursi
	if current > n {
		return
	}
	// Cetak nilai current jika ganjil
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	// Panggil fungsi untuk bilangan berikutnya
	tampilGanjil(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", N)
	tampilGanjil(N, 1)
	fmt.Println()
}
