/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk menghitung x dipangkatkan y
func pangkat(x int, y int) int {
	// Jika y = 0, maka hasil 1
	if y == 0 {
		return 1
	}
	// Rekursif
	return x * pangkat(x, y-1)
}

func main() {
	// Deklarasi variabel
	var x, y int

	// Input dari pengguna
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	// Menghitung hasil
	hasil := pangkat(x, y)

	// Menampilkan hasil
	fmt.Printf("Hasil %d pangkat %d adalah %d\n", x, y, hasil)
}
