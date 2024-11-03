package main

import "fmt"

var n int
func pola153(n int) {
	if n <= 0 {
		return
	}
	pola153(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	fmt.Println("POLA BINTANG * ")
	fmt.Print("INPUT N: ")
	fmt.Scan(&n)
	pola153(n)
}
