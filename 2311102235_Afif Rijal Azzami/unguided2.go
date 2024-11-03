package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak sejumlah bintang dalam satu baris
func cetakBintang(n int) {
	if n > 0 {
		fmt.Print("*")
		cetakBintang(n - 1) // Memanggil fungsi rekursif untuk mencetak sisa bintang
	}
}

// Fungsi rekursif untuk mencetak pola bintang per baris
func cetakPola(rows int) {
	if rows > 0 {
		cetakPola(rows - 1) // Panggilan rekursif ke baris sebelumnya untuk urutan dari atas ke bawah
		cetakBintang(rows)       // Cetak bintang sebanyak nomor baris
		fmt.Println()          // Pindah ke baris baru setelah mencetak bintang
	}
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&N)
	cetakPola(N)
}