// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama
func main() {
	// Deklarasi variabel n untuk menyimpan input pengguna
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	// Memanggil fungsi rekursif baris dengan input n
	baris(n)
}

// Fungsi rekursif untuk mencetak angka menurun dari bilangan ke 1
func baris(bilangan int) {
	// Kondisi dasar: jika bilangan == 1, cetak 1 dan hentikan rekursi
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		// Jika bilangan > 1, cetak bilangan, lalu panggil fungsi baris dengan bilangan - 1
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
