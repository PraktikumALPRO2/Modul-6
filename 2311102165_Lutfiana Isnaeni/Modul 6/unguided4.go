// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menampilkan barisan bilangan dari N ke 1 dan kembali ke N
func printSequence(n, current int) {
	// Basis rekursi
	if current == 1 {
		fmt.Print(current, " ")
		return
	}

	// Tampilkan angka saat ini dan panggil fungsi secara rekursif menurun
	fmt.Print(current, " ")
	printSequence(n, current-1)

	// Tampilkan angka saat kembali dari rekursi
	fmt.Print(current, " ")
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	fmt.Print("Hasil: ")
	printSequence(input, input)
	fmt.Println()
}
