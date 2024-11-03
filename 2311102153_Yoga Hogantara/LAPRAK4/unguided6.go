package main

import "fmt"

var x, y int
func pangkat153(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat153(x, y-1)
}

func main() {
	fmt.Println("INPUT X dan Y")
	fmt.Scan(&x, &y) 
	fmt.Println(pangkat153(x, y)) 
}
