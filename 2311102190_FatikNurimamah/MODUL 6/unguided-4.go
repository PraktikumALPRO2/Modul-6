package main

import (
	"fmt"
)

func cetakBarisan(bilangan int) {
	// Tampilkan bilangan saat ini
	fmt.Print(bilangan, " ")

	// Kondisi berhenti saat mencapai 1
	if bilangan == 1 {
		return
	}
	// Panggilan rekursif untuk menurunkan bilangan
	cetakBarisan(bilangan - 1)

	// Tampilkan bilangan lagi untuk membalik urutan dari 2 ke N
	fmt.Print(bilangan, " ")
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&bilangan)

	// Panggil fungsi rekursif
	cetakBarisan(bilangan)
	fmt.Println()
}