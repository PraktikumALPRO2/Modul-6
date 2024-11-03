package main

import "fmt"

var n int

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}
func main() {
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}
