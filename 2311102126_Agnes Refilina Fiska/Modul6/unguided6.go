package main

import "fmt"

// Fungsi rekursif untuk menghitung pangkat
func power(base int, exp int) int {
	// Basis case
	if exp == 0 {
		return 1
	}

	// Recursive case
	return base * power(base, exp-1)
}

func main() {
	fmt.Println("===================================")
	fmt.Println("    Program Menghitung Pangkat      ")
	fmt.Println("===================================")

	var testCases int
	fmt.Print("\nMasukkan jumlah test cases: ")
	fmt.Scan(&testCases)

	for i := 1; i <= testCases; i++ {
		var x, y int
		fmt.Printf("\n----------------------------------")
		fmt.Printf("\nTest Case %d:", i)
		fmt.Print("\nMasukkan nilai x dan y (dipisah spasi): ")
		fmt.Scan(&x, &y)

		result := power(x, y)
		fmt.Printf("Hasil %d pangkat %d: %d\n", x, y, result)
	}

	fmt.Println("----------------------------------")
	fmt.Println("\nProgram Selesai!")
}
