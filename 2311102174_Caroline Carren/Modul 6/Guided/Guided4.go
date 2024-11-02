// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Deklarasi variabel global n untuk menyimpan input
var n int

// Fungsi rekursif untuk menghitung faktorial dari n
func faktorial(n int) int {
	if n == 0 || n == 1 {
		// Kondisi dasar: jika n == 0 atau n == 1, kembalikan 1
		return 1
	} else {
		// Jika n > 1, kembalikan n * faktorial(n-1)
		return n * faktorial(n-1)
	}
}

// Fungsi utama
func main() {
	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	// Mencetak hasil faktorial dari n
	fmt.Println("Hasil faktorial dari", n, "adalah", faktorial(n))
}
