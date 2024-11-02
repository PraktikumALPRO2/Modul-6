// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menampilkan pola bintang
func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var input int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&input)
	printStars(input)
}
