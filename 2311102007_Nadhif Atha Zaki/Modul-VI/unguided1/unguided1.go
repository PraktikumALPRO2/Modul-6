package main

import "fmt"

// Fungsi fibonacci menghitung nilai suku ke-n dari deret Fibonacci
func fibonacci(n int) int {
	// Basis rekursi: jika n kurang dari atau sama dengan 1, kembalikan nilai n
	if n <= 1 {
		return n
	}
	// Rekursi: hitung nilai Sn sebagai penjumlahan dari dua suku sebelumnya (Sn-1 + Sn-2)
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Menampilkan header untuk deret Fibonacci hingga suku ke-10
	fmt.Println("Deret Fibonacci hingga suku ke-10:")
	fmt.Println("n\tSn")
	fmt.Println("-\t--")

	// Loop untuk menghitung dan menampilkan setiap suku dari 0 hingga 10
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t%d\n", i, fibonacci(i))
	}
}
