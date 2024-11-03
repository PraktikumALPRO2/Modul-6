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
	var n int
	fmt.Println("Masukkan nilai n:")
	fmt.Scan(&n)

	fmt.Printf("Deret Fibonacci hingga suku ke-%d adalah:\n", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}
}
