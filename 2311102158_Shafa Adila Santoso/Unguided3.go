package main

import "fmt"

// Fungsi rekursif untuk menemukan dan mencetak faktor dari N
func faktorBilangan(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	faktorBilangan(n, i+1)
}

func main() {
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	fmt.Printf("Faktor dari %d adalah: ", n)
	faktorBilangan(n, 1)
	fmt.Println()
}
