package main

import (
	"fmt"
)

// Fungsi untuk mencetak sejumlah bintang
func printStars(n int) {
	if n > 0 {
		fmt.Print("*")
		printStars(n - 1)
	}
}

// Fungsi rekursif untuk mencetak pola bintang
func printPattern(n, current int) {
	if current <= n {
		printStars(current)
		fmt.Println()
		printPattern(n, current+1)
	}
}

func main() {
	var n int
	fmt.Println("Masukkan nilai N:")
	fmt.Scan(&n)

	fmt.Printf("Pola bintang hingga %d baris:\n", n)
	printPattern(n, 1)
}
