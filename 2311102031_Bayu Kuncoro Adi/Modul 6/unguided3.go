package main

import (
	"fmt"
)

// Fungsi rekursif untuk menemukan dan mencetak faktor-faktor dari N
func findFactors(N, current int) {
	if current > N {
		return
	}
	if N%current == 0 {
		fmt.Print(current, " ")
	}
	findFactors(N, current+1)
}

// Fungsi untuk memulai pencarian faktor dari 1 hingga N
func printFactors(N int) {
	findFactors(N, 1)
	fmt.Println()
}

func main() {
	var N int
	fmt.Println("Masukkan nilai N:")
	fmt.Scan(&N)

	fmt.Printf("Faktor-faktor dari %d:\n", N)
	printFactors(N)
}
