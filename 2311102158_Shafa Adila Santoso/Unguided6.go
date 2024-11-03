package main

import "fmt"


// Fungsi rekursif untuk menghitung x pangkat y
func pangkat_158(x_158, y_158 int) int {
	if y_158 == 0 {
		return 1
	}

	return x_158 * pangkat_158(x_158, y_158-1)
}

func main() {
	var x_158, y_158 int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x_158)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y_158)

	// Menampilkan hasil pangkat_158 x_158^y_158
	fmt.Print("Hasil Perpangkatan: ", pangkat_158(x_158, y_158))
}
