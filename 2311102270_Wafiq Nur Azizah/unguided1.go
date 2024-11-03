package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Print("n      ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()


	fmt.Print("Sn    ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
