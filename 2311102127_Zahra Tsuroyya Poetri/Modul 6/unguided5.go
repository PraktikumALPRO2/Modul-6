package main

import "fmt"

func bilGanjil(n int) {
	// Base case: jika n lebih dari 1, maka proses berhenti
    if n < 1 { 
        return 
	}
    
	// Jika n adalah genap, kurangi n dengan 1, maka bilangan menjadi ganjil terbesar di bawah n
    if n % 2 == 0 {
		n -= 1
    }

	bilGanjil(n - 2) // Memanggil rekursif untuk bilangan ganjil dari 1 hingga n-2

	fmt.Print(n, " ") // Cetak

}

func main() {
    var N int // Mendeklarasikan variabel nilai N dengan tipe data integer 
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)

    // Memastikan input N adalah bilangan bulat positif
    if N > 0 {
        bilGanjil(N)
    } else {
        fmt.Println("Mohon masukkan bilangan bulat positif.") // Jika input N bukan bilangan bulat positif
    }
}
