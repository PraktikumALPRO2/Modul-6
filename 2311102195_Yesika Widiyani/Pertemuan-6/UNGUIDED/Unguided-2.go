package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bintang per baris
func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		fmt.Print("*")
	}
}

// Fungsi rekursif untuk menampilkan pola bintang
func printPattern(n, current int) {
	if current > n {
		return
	}
	printStars(current)
	fmt.Println()
	printPattern(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Println("Pola Bintang:")
	printPattern(n, 1)
}