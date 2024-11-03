package main

import "fmt"

var n int
func faktor153(n int, pembagi int) {
	if pembagi > n {
		return
	}
	if n % pembagi == 0 {
		fmt.Print(" ",pembagi)
	}
	faktor153(n, pembagi+1)
}

func main() {
	fmt.Println("FAKTOR")
	fmt.Print("INPUT: ")
	fmt.Scan(&n)
	faktor153(n, 1)
}
