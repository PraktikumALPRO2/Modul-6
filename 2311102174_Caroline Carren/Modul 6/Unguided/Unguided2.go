// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama
func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)
	printStar(n) // Memanggil fungsi printStar untuk mencetak pola bintang
}

// Fungsi rekursif untuk mencetak pola bintang
func printStar(n int) {
	// Basis rekursi: jika n == 0, hentikan rekursi
	if n == 0 {
		return
	}

	// Memanggil printStar dengan nilai n-1 (rekursi menurun)
	printStar(n - 1)

	// Setelah kembali dari rekursi, cetak bintang sebanyak nilai n
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println() // Pindah ke baris baru setelah mencetak bintang
}
