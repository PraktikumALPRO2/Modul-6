package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan faktor-faktor dari N
func cariFaktor(n, i int) {
	// Jika i lebih besar dari n, maka rekursi berhenti
	if i > n {
		return
	}
	// Jika n habis dibagi i, maka i adalah faktor dari n
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	// Panggil fungsi rekursif untuk nilai i berikutnya
	cariFaktor(n, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor-faktor dari %d adalah: ", N)
	cariFaktor(N, 1)
	fmt.Println()
}
