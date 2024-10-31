package main

import (
	"fmt"
)

func cetakBilanganGanjil(bilangan, batasAkhir int) {
	// Kondisi berhenti saat bilangan melebihi batas
	if bilangan > batasAkhir {
		return
	}

	// Tampilkan bilangan jika ganjil
	if bilangan%2 != 0 {
		fmt.Print(bilangan, " ")
	}

	// Panggilan rekursif untuk bilangan berikutnya
	cetakBilanganGanjil(bilangan+1, batasAkhir)
}

func main() {
	var batasAkhir int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&batasAkhir)

	// Panggil fungsi rekursif dimulai dari angka 1
	cetakBilanganGanjil(1, batasAkhir)
	fmt.Println()
}
