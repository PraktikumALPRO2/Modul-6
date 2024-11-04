package main

import "fmt"

// Fungsi rekursif untuk menampilkan deret bilangan dari M ke 1 dan kembali ke M
func tampilkanDeret(bilangan, saatIni int) {
	// Kondisi dasar rekursi
	if saatIni == 1 {
		fmt.Print(saatIni, " ")
		return
	}

	// Tampilkan nilai saat ini dan panggil fungsi secara rekursif ke bawah
	fmt.Print(saatIni, " ")
	tampilkanDeret(bilangan, saatIni-1)

	// Tampilkan nilai saat kembali dari rekursi
	fmt.Print(saatIni, " ")
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	fmt.Print("Hasil: ")
	tampilkanDeret(angka, angka)
	fmt.Println()
}
