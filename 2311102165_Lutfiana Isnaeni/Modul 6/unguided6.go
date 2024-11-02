// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var base, exponent int
	fmt.Print("Masukkan basis (x): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scan(&exponent)
	result := power(base, exponent)
	fmt.Printf("%d pangkat %d = %d\n", base, exponent, result)
}
