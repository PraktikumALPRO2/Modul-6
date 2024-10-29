package main

import (
	"fmt"
)

func printSequence(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printSequence(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)
	printSequence(N)
	fmt.Println()
}
