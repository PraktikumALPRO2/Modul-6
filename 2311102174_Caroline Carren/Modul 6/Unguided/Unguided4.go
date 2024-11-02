// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama
func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	// Memanggil fungsi rekursif untuk mencetak pola
	printRecursive(n)
}

// Fungsi rekursif untuk mencetak angka dari n hingga 1, lalu kembali dari 1 hingga n
func printRecursive(n int) {
	// Basis rekursi: jika n mencapai 1, cetak 1 dan hentikan rekursi
	if n == 1 {
		fmt.Print(n, " ")
		return
	}

	// Mencetak nilai n dalam urutan menurun
	fmt.Print(n, " ")

	// Memanggil fungsi secara rekursif dengan n-1
	printRecursive(n - 1)

	// Mencetak nilai n kembali dalam urutan menaik
	fmt.Print(n, " ")
}
