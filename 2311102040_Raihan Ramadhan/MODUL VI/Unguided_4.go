package main

import "fmt"

func tampilkanBilangan(n int) {
	if n == 1 {
		fmt.Print("1 ")
		return
	}

	fmt.Printf("%d ", n)

	tampilkanBilangan(n - 1)

	fmt.Printf("%d ", n)
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	tampilkanBilangan(n)
	fmt.Println()
}
