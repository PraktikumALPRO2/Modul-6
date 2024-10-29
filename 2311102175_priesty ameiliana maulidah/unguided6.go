package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung x^y
func power(x int, y int) int {
	if y == 0 {
		return 1 // Basis: x^0 = 1
	}
	return x * power(x, y-1) // Rekurens: x^y = x * x^(y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	result := power(x, y)
	fmt.Printf("Hasil dari %d dipangkatkan %d adalah %d\n", x, y, result)
}
