package main 

import "fmt"

// Fungsi rekursif untuk menghitung pangkat
func pangkat(n int) int {
	if n == 0{
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}

func main () {

	// Inputan user untuk bilangan
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	// Tampilkan hasil pangkat
	fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))

}