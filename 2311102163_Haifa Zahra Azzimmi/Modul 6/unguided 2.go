package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak pola bintang
func cetakBintang(n int) {
	// Kondisi akhir rekursi: jika n <= 0, hentikan fungsi
	if n <= 0 {
		return
	}
	// Panggil cetakBintang dengan nilai n-1 lebih dulu
	cetakBintang(n - 1)
	// Tampilkan bintang sejumlah n
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah baris bintang: ")
	fmt.Scan(&jumlah)

	fmt.Println("Hasil pola bintang:")
	cetakBintang(jumlah)
}
