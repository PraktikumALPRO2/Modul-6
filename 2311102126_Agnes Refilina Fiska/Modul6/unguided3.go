package main

import "fmt"

// Fungsi rekursif untuk mencari faktor
func findFactors(n int, current int) {
	// Basis case
	if current > n {
		return
	}

	// Jika current adalah faktor dari n
	if n%current == 0 {
		if current == n {
			fmt.Printf("%d", current)
		} else {
			fmt.Printf("%d ", current)
		}
	}

	// Panggil rekursif untuk angka selanjutnya
	findFactors(n, current+1)
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
		findFactors(n, 1)
		fmt.Println()
	}
}
