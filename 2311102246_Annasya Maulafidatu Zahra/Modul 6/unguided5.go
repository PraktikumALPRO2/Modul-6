package main

import "fmt"

func cetakGanjil(n int) {
	if n < 1 {
		return
	}
	cetakGanjil(n - 1)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	cetakGanjil(n)
	fmt.Println()
}
