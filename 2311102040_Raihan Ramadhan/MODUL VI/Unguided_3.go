package main

import (
	"fmt"
)

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
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor-faktor dari ", n, " adalah: ")
	cariFaktor(n, 1)
	fmt.Println()
}
