package main

import "fmt"

func fibonacci(n int) int {
	// Basis rekursi: jika n <= 1, kembalikan n
	if n <= 1 {
		return n
	}
	// Rekursi: Sn = Sn-1 + Sn-2
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Menampilkan deret Fibonacci hingga suku ke-10
	fmt.Println("Deret Fibonacci hingga suku ke-10:")
	fmt.Println("n\tSn")
	fmt.Println("-\t--")

	// Loop untuk menghitung dan menampilkan setiap suku
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t%d\n", i, fibonacci(i))
	}
}
