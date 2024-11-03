/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk menampilkan bilangan dari N hingga 1
func cetakBilangan(n int) {
	if n < 1 {
		return
	}
	fmt.Print(n, " ") // Menampilkan bilangan N
	cetakBilangan(n - 1) // Panggil fungsi secara rekursif
	fmt.Print(n, " ") // Menampilkan bilangan N setelah kembali dari rekursif
}

func main() {
	// Deklarasi variabel
	var N int

	// Input N dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Jika N <= 0
	if N <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif.")
		return
	}

	// Jika N sudah sesuai maka tampilkan barisan bilangan 
	fmt.Print("Barisan bilangan: ")
	cetakBilangan(N)
}