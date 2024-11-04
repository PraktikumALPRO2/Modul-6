package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga batas tertentu
func tampilkanGanjil(batas, angkaSaatIni int) {
	if angkaSaatIni > batas {
		return
	}
	if angkaSaatIni%2 != 0 {
		fmt.Print(angkaSaatIni, " ")
	}
	tampilkanGanjil(batas, angkaSaatIni+1)
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	tampilkanGanjil(angka, 1)
	fmt.Println()
}
