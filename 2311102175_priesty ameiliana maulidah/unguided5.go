package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan ganjil
func printOddNumbers(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printOddNumbers(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Println("Barisan bilangan ganjil dari 1 hingga", N, ":")
	printOddNumbers(N, 1)
}
