// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama
func main() {
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║         Bilangan Ganjil        ║")
	fmt.Println("╚════════════════════════════════╝")

	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	// Menampilkan header tabel
	fmt.Printf("\nBilangan ganjil dari 1 hingga %d:\n", n)
	fmt.Println("╔════════════════════════════════════════╗")
	printOddNumbers(1, n, 5) // Memanggil fungsi untuk menampilkan bilangan ganjil dengan 5 kolom per baris
	fmt.Println("╚════════════════════════════════════════╝")
}

// Fungsi untuk mencetak bilangan ganjil dari 1 hingga n dalam format tabel
func printOddNumbers(i, n, columns int) {
	count := 0 // Variabel untuk menghitung jumlah angka pada baris saat ini
	for ; i <= n; i += 2 {
		fmt.Printf("║ %-6d", i) // Menampilkan bilangan ganjil dengan lebar 6 karakter
		count++

		// Jika sudah mencapai jumlah kolom, pindah ke baris berikutnya
		if count%columns == 0 {
			fmt.Print("║\n") // Menutup baris dan mulai yang baru
		}
	}

	// Menutup baris jika kolom belum penuh di akhir loop
	if count%columns != 0 {
		for count%columns != 0 {
			fmt.Print("║       ") // Tambahkan ruang kosong jika kolom belum penuh
			count++
		}
		fmt.Println("║") // Menutup baris terakhir setelah kolom penuh
	}
}
