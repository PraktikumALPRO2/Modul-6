package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan dari n hingga 1, lalu kembali ke n
func barisanBlgan_158(n, current int) {
	fmt.Printf("%d ", current)

	if current == 1 {
		return
	}

	barisanBlgan_158(n, current-1)

	fmt.Printf("%d ", current)
}

func main() {
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	fmt.Print("Barisan Bilangan: ")
	barisanBlgan_158(n, n)
	fmt.Println()
}
