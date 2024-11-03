package main

import (
	"fmt"
)

func displaySequence(n int) {
	if n == 1 {
		fmt.Print("1 ")
	} else {
		fmt.Printf("%d ", n)
		displaySequence(n - 1)
		fmt.Printf("%d ", n)
	}
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&N)
	displaySequence(N)
}
