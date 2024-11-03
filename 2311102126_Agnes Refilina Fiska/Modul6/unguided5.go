package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil
func printOddNumbers(n int, current int) {
	// Basis case
	if current > n {
		return
	}

	// Jika current adalah bilangan ganjil
	if current%2 != 0 {
		if current == 1 {
			fmt.Printf("%d", current)
		} else {
			fmt.Printf(" %d", current)
		}
	}

	// Panggil rekursif untuk angka selanjutnya
	printOddNumbers(n, current+1)
}

func main() {
	var testCases int
	fmt.Println("===================================")
	fmt.Println("    Program Deret Bilangan Ganjil   ")
	fmt.Println("===================================")

	fmt.Print("\nMasukkan jumlah test cases: ")
	fmt.Scan(&testCases)

	for i := 1; i <= testCases; i++ {
		var n int
		fmt.Printf("\n----------------------------------")
		fmt.Printf("\nTest Case %d:", i)
		fmt.Print("\nMasukkan N: ")
		fmt.Scan(&n)

		fmt.Printf("Keluaran: ")
		printOddNumbers(n, 1)
		fmt.Println()
	}
	fmt.Println("----------------------------------")
	fmt.Println("\nProgram Selesai!")
}
