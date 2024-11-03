package main

import "fmt"

var n int
func barisbilangan153(n int) {
	if n < 1 {
		return
	}
	fmt.Print(" ", n)
	if n > 1 {
		barisbilangan153(n - 1)
		fmt.Print(" ",n)
	}
}

func main() {
	fmt.Print("INPUT N: ")
	fmt.Scan(&n)
	barisbilangan153(n)
}
