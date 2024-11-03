package main

import "fmt"

func baris(n int) {
	// Base case: jika nilai n adalah 1, maka proses berhenti
    if n == 1 { 
        fmt.Print(n, " ")
        return // Mengembalikan fungsi ke pemanggilan sebelumnya
    }
    
    // Menampilkan bilangan menurun
    fmt.Print(n, " ")
    baris(n - 1) // Melakukan iterasi menurun (nilai n berkurang 1 setiap kali)
    
    // Menampilkan bilangan meningkat
    fmt.Print(n, " ") // Mencetak nilai n setiap tahap meningkat
}

func main() {
    var N int // Mendeklarasikan variabel nilai N dengan tipe data integer 
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)

    // Memastikan input N adalah bilangan bulat positif
    if N > 0 { 
        baris(N)
    } else {
        fmt.Println("Mohon masukkan bilangan bulat positif.") // Jika input N bukan bilangan bulat positif
    }
}
