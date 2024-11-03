package main

import "fmt"

func cetakGanjil(n, current int) {
	if current <= n {
		if current%2 != 0 {
			fmt.Print(current, " ")
		}
		cetakGanjil(n, current+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	fmt.Print("Bilangan ganjil dari 1 hingga ", n, " adalah: ")
	cetakGanjil(n, 1)
	fmt.Println()
}
