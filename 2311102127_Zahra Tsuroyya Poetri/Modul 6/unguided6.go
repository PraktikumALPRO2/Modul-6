package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func pangkat(x int, y int) int {
	// Base case: jika y adalah 0, maka hasilnya 1 (x^0 = 1)
	if y == 0 { 
		return 1
	}

	// Base case: jika y adalah 1, maka hasilnya x (x^1 = x)
	if y == 1{ 
		return x
	}

	//  Jika y kurang dari 0 atau negatif, maka hitug pangkat positif dan menjadi kebalikannya
	if y < 0 { 
		return 1 / pangkat(x, -y) // Menggunakan pembagian untuk menangani pangkat negatif
	}

	return x * pangkat(x, y-1) // Memanggil rekursif untuk menghitung x pangkat (y - 1) dan kalikan dengan x
}

func main() {
	var x, y int // Mendeklarasikan nilai x dan y dengan tipe data integer
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	// Menghitung hasil pangkat
	hasil := pangkat(x, y) 
	fmt.Printf("%d pangkat %d adalah %d\n", x, y, hasil)
}