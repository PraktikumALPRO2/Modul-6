package main

import (
	"fmt"
)

func cetakBintang(n int, current int) {
	if current > n {
		return
	}
	
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	cetakBintang(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&N)

	cetakBintang(N, 1)
}