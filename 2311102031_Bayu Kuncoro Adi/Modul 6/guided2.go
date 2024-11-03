package main

import "fmt"

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	if n > 0 {
		fmt.Println("Hasil penjumlahan:", penjumlahan(n))
	} else {
		fmt.Println("Masukkan bilangan positif")
	}
}
