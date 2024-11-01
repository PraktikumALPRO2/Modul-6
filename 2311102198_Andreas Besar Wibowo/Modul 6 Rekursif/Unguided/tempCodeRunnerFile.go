// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk pangkat
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}

	return x * pangkat(x, y-1)
}

func main() {
	// Inputan user untuk bilangan dan bilangan pangkat
	var x, y int
	fmt.Print("Bilangan : ")
	fmt.Scan(&x)
	fmt.Print("Bilangan pangkat : ")
	fmt.Scan(&y)

	// Tampilkan hasil pangkatnya
	hasil := pangkat(x, y)
	fmt.Printf("\nHasil %d dipangkatkan %d adalah: %d\n", x, y, hasil)
}
