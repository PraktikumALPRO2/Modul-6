package main

import "fmt"

var n int
func fibo153(n int) int {
	if n <= 1 {
		return n
	}
	return fibo153(n-1) + fibo153(n-2)
}

func main() {
	fmt.Println("FIBONACCI")
	fmt.Print("INPUT N: ")
	fmt.Scan(&n)

	fmt.Print("Sn = ")
	for i := 0; i <= n; i++ {
		fmt.Print(" ", fibo153(i))
	}
	fmt.Println()
}
