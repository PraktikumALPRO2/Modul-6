package main

import "fmt"

func printStars(n int, current int) {
	if current > n {
		return
	}

	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printStars(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printStars(n, 1)
}
