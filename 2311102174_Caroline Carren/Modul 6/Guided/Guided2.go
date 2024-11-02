// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi rekursif untuk menghitung penjumlahan dari n hingga 1
func penjumlahan(n int) int {
	if n == 1 {
		// Kondisi dasar: jika n == 1, kembalikan 1
		return 1
	} else {
		// Jika n > 1, kembalikan n + penjumlahan(n-1)
		return n + penjumlahan(n-1)
	}
}

// Fungsi utama
func main() {
	// Deklarasi variabel n untuk menyimpan input pengguna
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	// Mencetak hasil penjumlahan dari n hingga 1
	fmt.Println(penjumlahan(n))
}
