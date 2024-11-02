package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung deret Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("n\tSn")
	// Cetak tabel deret Fibonacci hingga suku ke-10
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t%d\n", i, fibonacci(i))
	}
}
