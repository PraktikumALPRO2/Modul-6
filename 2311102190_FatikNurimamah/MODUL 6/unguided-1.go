package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung suku ke-x dalam deret Fibonacci
func MenghitungFibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return MenghitungFibonacci(x-1) + MenghitungFibonacci(x-2)
}

func main() {
	var jumlahSuku int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scan(&jumlahSuku)
	fmt.Printf("n   |")
	for i := 0; i <= jumlahSuku; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()

	fmt.Printf("Sn  |")
	for i := 0; i <= jumlahSuku; i++ {
		fmt.Printf(" %d ", MenghitungFibonacci(i))
	}
	fmt.Println()
}