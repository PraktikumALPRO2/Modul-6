// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi untuk menghitung pangkat x^y secara rekursif
func power(x, y int) int {
	// Basis rekursi: jika y == 0, maka hasilnya adalah 1
	if y == 0 {
		return 1
	}
	// Menghitung x * power(x, y-1)
	return x * power(x, y-1)
}

// Fungsi utama
func main() {
	var x, y int

	// Header Program
	fmt.Println("=========================================")
	fmt.Println("          PROGRAM KALKULATOR PANGKAT     ")
	fmt.Println("=========================================")

	// Input pengguna untuk bilangan x
	fmt.Print("Masukkan bilangan bulat (x): ")
	fmt.Scanln(&x)

	// Input pengguna untuk pangkat y
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scanln(&y)
	fmt.Println("-----------------------------------------")

	// Kalkulasi hasil pangkat
	result := power(x, y)

	// Menampilkan hasil
	fmt.Printf("Hasil dari %d ^ %d adalah: %d\n", x, y, result)
	fmt.Println("=========================================")
}
