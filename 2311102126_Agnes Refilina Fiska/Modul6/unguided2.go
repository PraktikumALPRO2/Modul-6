package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang
func printStars(n int, current int) {
	// Basis case
	if n == 0 {
		return
	}

	// Cetak bintang sebanyak current
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	// Panggil rekursif untuk baris selanjutnya
	printStars(n-1, current+1)
}

func main() {
	var testCases int
	fmt.Print("Masukkan jumlah test cases: ")
	fmt.Scan(&testCases)

	for i := 1; i <= testCases; i++ {
		var n int
		fmt.Printf("\nTest Case %d\n", i)
		fmt.Print("Masukkan N: ")
		fmt.Scan(&n)

		fmt.Printf("Keluaran:\n")
		printStars(n, 1)
	}
}
