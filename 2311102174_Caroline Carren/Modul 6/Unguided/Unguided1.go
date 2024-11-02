// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai deret Fibonacci ke-n
func fibonacci(n int) int {
	// Jika n adalah 0 atau 1, kembalikan nilai n (basis kasus rekursi)
	if n <= 1 {
		return n
	}
	// Jika n > 1, panggil fungsi fibonacci secara rekursif
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Menampilkan header tabel dengan dekorasi ASCII
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     Tabel Deret Fibonacci                          ║")
	fmt.Println("╠═════════╦══════════════════════════════════════════════════════════╣")

	// Menampilkan baris pertama dengan indeks N (0 hingga 10)
	fmt.Printf("║  %-5s  ║", "N")
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %-3d ", i) // Menampilkan indeks
	}
	fmt.Println("║") // Akhir dari baris indeks
	fmt.Println("╠═════════╬══════════════════════════════════════════════════════════╣")

	// Menampilkan baris kedua dengan nilai deret Fibonacci untuk setiap N
	fmt.Printf("║  %-5s  ║", "Sn")
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %-3d ", fibonacci(i)) // Menampilkan nilai deret Fibonacci
	}
	fmt.Println("║") // Akhir dari baris nilai Fibonacci
	fmt.Println("╚═════════╩══════════════════════════════════════════════════════════╝")
}
