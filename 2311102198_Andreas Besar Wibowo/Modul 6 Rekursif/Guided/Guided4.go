package main

import "fmt"

var n int
// Fungsi untuk menghitung faktorial
func faktorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main () {

	// Inputan user untuk bilangan
	fmt.Scan(&n)
	// Tampilkan nilai faktorial
	fmt.Println(faktorial(n))
}