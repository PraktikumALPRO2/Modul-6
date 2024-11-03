/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk mencetak pola bintang
func cetakBintang(n int, i int) {
	if i > n {
		return
	}

	// Perulangan mencetak bintang sebanyak i
	for j := 0; j < i; j++ {
		fmt.Print("* ")
	}
	fmt.Println()

	// Panggil fungsi rekursif untuk baris berikutnya
	cetakBintang(n, i+1)
}

func main() {
	// Deklarasi variabel
	var n int

	// Input n 
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	// Cetak pola bintang
	fmt.Println("Pola bintang:")
	cetakBintang(n, 1)
}