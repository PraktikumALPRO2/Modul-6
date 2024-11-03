package main

import (
	"fmt"
)

func printOddSequence(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printOddSequence(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	printOddSequence(n, 1)
	fmt.Println()
}