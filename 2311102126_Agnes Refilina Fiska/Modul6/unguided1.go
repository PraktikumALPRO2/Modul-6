package main

import "fmt"

func fibonacci(n int) int {
	// Basis case
	if n <= 1 {
		return n
	}
	// Recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Print header
	fmt.Printf("n  : ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%-4d", i)
	}
	fmt.Println()

	// Print Sn values
	fmt.Printf("Sn : ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%-4d", fibonacci(i))
	}
	fmt.Println()
}
