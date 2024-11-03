package main

import (
	"fmt"
)

func faktorbilangan(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorbilangan(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah : ")
	faktorbilangan(n, 1)
	fmt.Println()
}