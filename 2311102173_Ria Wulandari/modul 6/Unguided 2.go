package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan sejumlah bintang pada satu baris
func tampilkanBintang(jumlahBintang int) {
	// Jika jumlah bintang mencapai 0, hentikan rekursi
	if jumlahBintang == 0 {
		return
	}
	// Panggil rekursi untuk menurunkan jumlah bintang
	tampilkanBintang(jumlahBintang - 1)
	// Cetak satu bintang
	fmt.Print("*")
}

// Fungsi rekursif untuk mencetak pola bintang dari baris pertama hingga baris terakhir
func cetakPolaBintang(totalBaris int, barisSaatIni int) {
	// Jika baris saat ini melewati total baris, hentikan rekursi
	if barisSaatIni > totalBaris {
		return
	}
	// Tampilkan bintang sesuai dengan nomor baris saat ini
	tampilkanBintang(barisSaatIni)
	// Pindah ke baris baru
	fmt.Println()
	// Panggil fungsi ini lagi untuk baris berikutnya
	cetakPolaBintang(totalBaris, barisSaatIni+1)
}

func main() {
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris (N): ")
	fmt.Scan(&jumlahBaris)

	fmt.Println("Pola bintang:")
	cetakPolaBintang(jumlahBaris, 1)
}
