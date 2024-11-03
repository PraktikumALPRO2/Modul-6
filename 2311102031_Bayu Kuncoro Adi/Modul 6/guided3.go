package main

import "fmt"

// Fungsi rekursif untuk menghitung 2^n
func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n (bilangan bulat non-negatif): ")
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))
	} else {
		fmt.Println("Masukkan bilangan bulat non-negatif")
	}
}
