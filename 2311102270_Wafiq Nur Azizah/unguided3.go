package main

import (
	"fmt"
)

func cariFaktor(n int, current int) {
	if current > n {
		return
	}
	if n%current == 0 { 
		fmt.Print(current, " ")
	}
	cariFaktor(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&N)

	fmt.Print("Faktor dari ", N, ": ")
	cariFaktor(N, 1)
	fmt.Println()
}
