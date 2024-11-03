package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung pangkat
func pangkat(x, y int) int {
	if y == 0 {
		return 1 // Base case: jika pangkatnya 0, hasilnya 1
	} else {
		return x * pangkat(x, y-1) // Kalikan x dengan hasil pangkat(x, y-1)
	}
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scan(&y)
	fmt.Println("Hasil:", pangkat(x, y))
}
