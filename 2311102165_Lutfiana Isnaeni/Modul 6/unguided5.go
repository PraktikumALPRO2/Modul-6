// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil dari 1 hingga N
func printOddSequence(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printOddSequence(n, current+1)
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	printOddSequence(input, 1)
	fmt.Println()
}
