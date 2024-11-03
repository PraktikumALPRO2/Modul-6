package main

import (
	"fmt"
)

func displayOddSequence(n, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", current)
	displayOddSequence(n, current+2)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&N)
	displayOddSequence(N, 1)
}
