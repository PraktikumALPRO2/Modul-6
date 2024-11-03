package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga N
func tmplkanBlgnGanjil_158(n_158, current int) {
	// Basis rekursi: jika current melebihi n_158, maka hentikan fungsi
	if current > n_158 {
		return
	}

	// Jika current adalah bilangan ganjil, cetak current
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	// Panggil fungsi rekursif berikutnya dengan current + 1
	tmplkanBlgnGanjil_158(n_158, current+1)
}

func main() {
	var n_158 int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n_158)

	if n_158 <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	fmt.Print("Barisan bilangan ganjil: ")
	tmplkanBlgnGanjil_158(n_158, 1)
	fmt.Println()
}
