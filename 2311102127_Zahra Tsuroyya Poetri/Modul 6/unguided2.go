package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang
func cetakBintang(n int) {
	// Base case: berhenti ketika n = 0
	if n == 0 { 
		return 
	}
	cetakBintang(n - 1) // Memanggil fungsi rekursif untuk mencetak baris sebelumnya
	for i := 0; i < n; i++ { 
		fmt.Print("*") // Mencetak baris dengan n bintang, setelah baris sebelumnya tercetak
	}
	fmt.Println()
}

func main() {
	var n int // Mendeklarasikan variabel nilai N dengan tipe data integer 
	fmt.Print("N: ") // Pengguna menginputkan nilai N
	fmt.Scan(&n)

	cetakBintang(n)
}