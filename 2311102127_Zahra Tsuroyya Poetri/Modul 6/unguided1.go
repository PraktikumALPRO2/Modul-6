package main

import "fmt"

// Fungsi rekursif
func S(n int) int {
	// Base case: Jika n adalah sama dengan 1, maka akan mengembalikan nilai n
	if n <= 1 { 
		return n
	}
	// Rekursif case:
	return S(n-1) + S(n-2) // Fungsi S melakukan rekursif untuk menghitung S(n-1) dan S(n-2)
}

func main() {
	var n int = 10 // Inisialisasi variabel untuk menyimpan jumlah suku

	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", n) // Menampilkan pesan
	for i := 0; i <= n; i++ {                             // Loop untuk mencetak setiap suku dari 0 hingga n
		fmt.Printf("%d\t", i) // Mencetak suku ke-i dan nilai menggunakan fungsi S(i)
	}

	fmt.Println()

	for i:= 0; i <= n; i++{
		fmt.Printf("%d\t", S(i))
	}
	fmt.Println()
}
