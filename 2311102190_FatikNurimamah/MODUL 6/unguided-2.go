package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bintang pada setiap baris
func cetakJumlahBintang(jumlahBintang int) {
	if jumlahBintang == 0 {
		return
	}
	cetakJumlahBintang(jumlahBintang - 1)
	fmt.Print("*")
}

// Fungsi rekursif untuk mencetak pola bintang 
func polaBintang(totalBaris int, barisSaatIni int) {
	if barisSaatIni > totalBaris {
		return
	}
	cetakJumlahBintang(barisSaatIni)
	fmt.Println()
	polaBintang(totalBaris, barisSaatIni+1)
}

func main() {
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris (N): ")
	fmt.Scan(&jumlahBaris)

	fmt.Println("Pola bintang:")
	polaBintang(jumlahBaris, 1)
}
