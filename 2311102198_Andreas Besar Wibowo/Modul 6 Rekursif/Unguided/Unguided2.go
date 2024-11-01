// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang
func cetakBintang(n int) {
	if n > 0 {
		fmt.Print("*")
		cetakBintang(n - 1)
	}
}

// Fungsi rekursif untuk mencetak pola bintang
func cetakPola(n, baris int) {
	if baris <= n {
		cetakBintang(baris)
		fmt.Println()
		cetakPola(n, baris+1)
	}
}

func main() {
	// Inputan user untuk baris pola bintang
	var n int
	fmt.Print("Jumlah baris pola bintang : ")
	fmt.Scan(&n)

	// Tampilkan pola bintang
	fmt.Println("\npola bintang :")
	cetakPola(n, 1)
}
