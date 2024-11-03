package main

import "fmt"

func power(x, y int) int {
	// Basis case: jika pangkat 0, return 1
	if y == 0 {
		return 1
	}
	// Basis case: jika pangkat 1, return x
	if y == 1 {
		return x
	}
	// Basis case: jika pangkat negatif, return 0
	if y < 0 {
		return 0
	}

	// Recursive case: pangkat genap
	if y%2 == 0 {
		temp := power(x, y/2)
		return temp * temp
	}
	// Recursive case: pangkat ganjil
	return x * power(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat y: ")
	fmt.Scan(&y)

	hasil := power(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasil)
}
