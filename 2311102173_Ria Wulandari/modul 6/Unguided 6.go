package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai a pangkat b
func hitungPangkat(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * hitungPangkat(a, b-1)
}

func main() {
	var dasar, pangkat int
	fmt.Print("Masukkan dasar (a): ")
	fmt.Scan(&dasar)
	fmt.Print("Masukkan pangkat (b): ")
	fmt.Scan(&pangkat)
	hasil := hitungPangkat(dasar, pangkat)
	fmt.Printf("%d pangkat %d = %d\n", dasar, pangkat, hasil)
}
