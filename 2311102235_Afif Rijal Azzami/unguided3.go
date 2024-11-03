package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencari faktor dari N
func cariFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	cariFaktor(n, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)
	fmt.Print("Faktor dari ", N, " adalah: ")
	cariFaktor(N, 1)
	fmt.Println()
}
