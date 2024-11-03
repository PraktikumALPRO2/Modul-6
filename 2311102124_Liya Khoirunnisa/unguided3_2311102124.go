/* Liya Khoirunnisa - 2311102124 */

package main

import (
	"fmt"
)

// Fungsi untuk menampilkan faktor dari N
func faktor(n int, i int) {
	// Jika i lebih beras dari n, maka keluar dari fungsi
	if i > n {
		return
	}

	// Jika n habis dibagi i, maka cetak faktor
	if n%i == 0 {
		fmt.Print(i)
	}

	// Memanggil fungsi rekursif
	faktor(n, i+1)
}

func main() {
	// Deklarasi variabel
	var N int

	// Input n
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&N)

	// Jika n negatif
	if N <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif.")
		return
	}

	// Mencetak faktor ke layar
	fmt.Printf("Faktor dari %d adalah: ", N)
	faktor(N, 1)
}