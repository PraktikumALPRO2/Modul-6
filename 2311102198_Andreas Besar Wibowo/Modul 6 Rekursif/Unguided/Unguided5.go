// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil
func cetakGanjil(n, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}

	cetakGanjil(n, i+1)
}

func main() {
	// Inputan user untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan barisan bilangan ganjil
	fmt.Printf("\nBarisan bilangan ganjil dari 1 hingga %d adalah: ", n)
	cetakGanjil(n, 1)
	fmt.Println()
}
