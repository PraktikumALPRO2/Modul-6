package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 10
	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("Suku ke-%d: %d\n", i, fibonacci(i))
	}
}
