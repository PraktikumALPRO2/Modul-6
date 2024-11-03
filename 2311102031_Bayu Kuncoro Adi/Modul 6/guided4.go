package main

import "fmt"

// Fungsi rekursif untuk menghitung faktorial dari n
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	if n >= 0 {
		fmt.Println("Hasil faktorial dari", n, "adalah", faktorial(n))
	} else {
		fmt.Println("Masukkan bilangan bulat non-negatif")
	}
}
