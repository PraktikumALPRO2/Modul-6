package main

import (
	"fmt"
)

func printOddNumbers(n int, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printOddNumbers(n, current+2)
}

func main() {
	var N int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&N)

	if N > 0 {
		printOddNumbers(N, 1)
		fmt.Println()
	} else {
		fmt.Println("N harus bilangan bulat positif.")
	}
}
