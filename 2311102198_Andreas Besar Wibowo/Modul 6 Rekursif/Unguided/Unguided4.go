// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk mencetak barisan
func cetakBarisan(n, nilai int) {
	fmt.Printf("%d ", nilai)
	if nilai == 1 {
		return
	}

	// Mencetak barisan
	cetakBarisan(n, nilai-1)
	fmt.Printf("%d ", nilai)
}

func main() {
	// Inputan user untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan barisan bilangan
	fmt.Printf("\nBarisan bilangan dari %d hingga 1 dan kembali ke %d adalah: ", n, n)
	cetakBarisan(n, n)
	fmt.Println()
}
