package main

import "fmt"

var n int
func billganjil153(n int) {
	if n < 1 {
		return
	}
	billganjil153(n - 1)
	if n % 2 != 0 {
		fmt.Print(" ", n)
	}
}

func main() {
	fmt.Print("INPUT N: ")
	fmt.Scan(&n)
	billganjil153(n)
}
