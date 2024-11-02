// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai 2^n
func pangkat(n int) int {
	if n == 0 {
		// Kondisi dasar: jika n == 0, kembalikan 1
		return 1
	} else {
		// Jika n > 0, kembalikan 2 * pangkat(n-1)
		return 2 * pangkat(n-1)
	}
}

// Fungsi utama
func main() {
	// Deklarasi variabel n untuk menyimpan input pengguna
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	// Mencetak hasil dari 2 pangkat n
	fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))
}
