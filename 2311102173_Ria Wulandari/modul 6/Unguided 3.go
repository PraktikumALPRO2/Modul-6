package main

import "fmt"

// Fungsi rekursif untuk mencari faktor dari suatu bilangan
func tampilkanFaktor(bilangan, indeks int) {
	if indeks > bilangan {
		return
	}
	if bilangan%indeks == 0 {
		fmt.Print(indeks, " ")
	}
	tampilkanFaktor(bilangan, indeks+1)
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	fmt.Print("Faktor dari ", angka, ": ")
	tampilkanFaktor(angka, 1)
	fmt.Println()
}
