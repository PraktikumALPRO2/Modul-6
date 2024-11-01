// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Inputan user untuk jumlah suku fibonacci
	var n int
	fmt.Print("	Jumlah suku Fibonacci : ")
	fmt.Scan(&n)

	// Cetak tabel deret fibonacci
	fmt.Println("\nDeret Fibonacci")
	fmt.Println("n   |   Sn")
	fmt.Println("------------")
	for a := 0; a <= n; a++ {
		fmt.Printf("%-3d | %3d\n", a, fibonacci(a))
	}
}
