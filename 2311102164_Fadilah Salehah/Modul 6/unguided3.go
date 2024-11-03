package main

import (
	"fmt"
)

func faktorisasi(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorisasi(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah : ")
	faktorisasi(n, 1)
	fmt.Println()
}