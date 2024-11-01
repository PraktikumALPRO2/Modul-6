package main

import "fmt"

// fungsi rekursif untuk penjumlahan
func penjumlahan (n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}

func main() {

	// Inputan user untuk bilangan
	var n int
	fmt.Scan(&n)
	// Tampilkan hasil penjumlahan
	fmt.Println(penjumlahan(n))
	
}