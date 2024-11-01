// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"


// Fungsi rekursif untuk mencetak faktor bilangan
func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	cetakFaktor(n, i+1)
}

func main() {
	// Inputan pengguna untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan faktor dari inputan user
	fmt.Printf("\nFaktor dari bilangan %d adalah: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}
