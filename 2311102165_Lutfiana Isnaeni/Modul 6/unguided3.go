// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menemukan faktor bilangan
func printFactors(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	printFactors(n, i+1)
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	fmt.Print("Faktor dari ", input, ": ")
	printFactors(input, 1)
	fmt.Println()
}
