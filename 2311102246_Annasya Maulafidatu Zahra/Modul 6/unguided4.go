package main

import "fmt"

// Fungsi untuk menampilkan barisan bilangan dari N hingga 1 dan kembali ke N
func cetakBarisan(n, current int) {
	if current < 1 {
		return
	}
	fmt.Print(current, " ")
	cetakBarisan(n, current-1)
	if current != n {
		fmt.Print(current, " ")
	}
}

func main() {
	var n int
	fmt.Scanln(&n) // Membaca input tanpa pesan output
	cetakBarisan(n, n)
	fmt.Println() // Menambahkan baris baru setelah output
}
