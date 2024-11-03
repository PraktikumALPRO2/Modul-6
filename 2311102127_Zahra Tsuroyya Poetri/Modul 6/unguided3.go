package main

import "fmt"

// Fungsi rekursif untuk mencari dan mencetak faktor dari N
func cariFaktor(n, habisBagi int) {
    // Base case: jika habisBagi melebihi N, maka proses berhenti
    if habisBagi > n {
        return
    }
    // Cek apakah habisBagi adalah faktor dari N 
    if n % habisBagi == 0 {
        fmt.Print(habisBagi, " ")
    }
    // Memanggil rekursif dengan menambah habisBagi
    cariFaktor(n, habisBagi+1)
}

func main() {
    var N int // Mendeklarasikan variabel nilai N dengan tipe data integer 
    fmt.Print("Masukkan bilangan bulat positif N: ") // Pengguna memasukkan bilangan bulang positif N
    fmt.Scan(&N)

    fmt.Print("Hasil: ") // Cetak hasil
    cariFaktor(N, 1)
    fmt.Println()
}
