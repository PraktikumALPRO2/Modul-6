package main

import "fmt"

func hitungPangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * hitungPangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x dan y: ")
	fmt.Scan(&x, &y)

	fmt.Printf("Hasil %d pangkat %d: %d\n", x, y, hitungPangkat(x, y))
}
