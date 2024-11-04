package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	hasil := pangkat(x, y)
	fmt.Println(hasil)
}