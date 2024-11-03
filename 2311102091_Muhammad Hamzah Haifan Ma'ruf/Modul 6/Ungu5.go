package main

import (
	"fmt"
)

func tampilkanganjil(n int, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	tampilkanganjil(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n : ")
	fmt.Scan(&n)
	fmt.Print("Barisan bilangan ganjil dari 1 - ", n, " adalah : ")
	tampilkanganjil(n, 1)
	fmt.Println()
}