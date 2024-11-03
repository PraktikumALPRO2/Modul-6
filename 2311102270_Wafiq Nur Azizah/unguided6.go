package main

import (
	"fmt"
)

func angka(x int, y int) int {
	if y == 0 {
		return 1
	}
	if y < 0 {
		return 1 / angka(x, -y)
	}
	return x * angka(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	result := angka(x, y)
	fmt.Printf("Hasil dari %d dipangkatkan %d adalah %d\n", x, y, result)
}
