// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama
func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&n)

	// Memanggil fungsi faktor dan mencetak hasilnya
	fmt.Println("Faktor dari", n, "adalah:", faktor(n, 1))
}

// Fungsi rekursif untuk mencari faktor bilangan
func faktor(n, i int) []int {
	// Basis rekursi: jika i lebih besar dari n, kembalikan nil
	if i > n {
		return nil
	}

	// Jika i adalah faktor dari n, tambahkan i ke dalam slice hasil
	if n%i == 0 {
		return append([]int{i}, faktor(n, i+1)...)
	}

	// Jika i bukan faktor dari n, lanjutkan ke nilai berikutnya
	return faktor(n, i+1)
}
