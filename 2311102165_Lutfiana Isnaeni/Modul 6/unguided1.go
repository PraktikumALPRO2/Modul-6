// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scanln(&n)

	// Menampilkan baris pertama untuk nilai n
	fmt.Print("n  ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Menampilkan baris kedua untuk nilai S_n
	fmt.Print("S_n ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}
