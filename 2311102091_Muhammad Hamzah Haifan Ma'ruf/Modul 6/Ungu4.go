package main

import (
	"fmt"
)

func turun(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
		} else {
			fmt.Printf("%d ", n)
			turun(n - 1)
		}
}

func naik(n, i int) {
	if i == n {
		fmt.Printf("%d ", i)
		} else {
			fmt.Printf("%d ", i)
			naik(n, i+1)
		}
}

func tampilkanbarisan(n int) {
	turun(n)
	naik(n, 2)
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n : ")
	fmt.Scanln(&n)
	fmt.Print("Hasil barisan bilangan : ")
	tampilkanbarisan(n)
}