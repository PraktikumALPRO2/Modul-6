package main

import "fmt"

// Fungsi rekursif untuk mencetak barisan turun
func printDescending(n int, current int) {
	// Basis case
	if current < 1 {
		return
	}

	// Cetak angka current
	if current == 1 {
		fmt.Printf("%d ", current)
	} else {
		fmt.Printf("%d ", current)
	}

	// Panggil rekursif untuk angka selanjutnya
	printDescending(n, current-1)
}

// Fungsi rekursif untuk mencetak barisan naik
func printAscending(n int, current int) {
	// Basis case
	if current > n {
		return
	}

	// Cetak angka current
	fmt.Printf("%d ", current)

	// Panggil rekursif untuk angka selanjutnya
	printAscending(n, current+1)
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

		fmt.Printf("Keluaran: ")
		// Cetak barisan turun dari N ke 1
		printDescending(n, n)
		// Cetak barisan naik dari 2 ke N
		printAscending(n, 2)
		fmt.Println()
	}
}
