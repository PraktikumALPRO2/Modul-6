package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bilangan dari N ke 1 dan kembali ke N
func cetakTurunNaik(n int) int {
	if n == 0 {
		return 0 // Base case: jika n = 0, kembalikan 0 dan hentikan rekursi
	} else {
		fmt.Print(n, " ")      // Cetak nilai n saat turun
		cetakTurunNaik(n - 1)  // Rekursif ke bawah hingga n mencapai 0
		fmt.Print(n, " ")      // Cetak nilai n saat naik kembali
	}
	return 0
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)
	cetakTurunNaik(N)
	fmt.Println()
}
