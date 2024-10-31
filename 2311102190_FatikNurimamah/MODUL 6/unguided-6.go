package main

import (
	"fmt"
)

func HitungPangkat(x, y int) int {
	// Kondisi dasar: jika pangkat (y) adalah 0, hasilnya 1
	if y == 0 {
		return 1
	}
	return x * HitungPangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	// Panggil fungsi rekursif untuk menghitung pangkat
	hasilAkhir := HitungPangkat(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasilAkhir)
}
