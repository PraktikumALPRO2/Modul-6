package main

import "fmt"

func cetakGanjil(n int) {
	if n < 1 {
		return
	}

	cetakGanjil(n - 2)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)
	cetakGanjil(N)
	fmt.Println()
}
