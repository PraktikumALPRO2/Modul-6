/* Liya Khoirunnisa - 2311102124 */
package main

import "fmt"

// Fungsi untuk menghitung suku Fibonacci ke-n
func fibonacci(n int) int {
    if n <= 1 {
        return n // jika n adalah 0 atau 1, kembalikan n
    }
    return fibonacci(n-1) + fibonacci(n-2) // algoritma rekursif
}

func main() {
    // deklarasi variabel
    var n int

    // Input jumlah suku fibonacci dari pengguna
    fmt.Print("Masukkan jumlah suku Fibonacci: ")
    fmt.Scan(&n)

    // Memastikan input positif
    if n < 0 {
        fmt.Println("Masukkan angka positif.")
        return
    }

    // Mencetak deret fibonacci
    fmt.Print("Deret Fibonacci: ")
    for i := 0; i < n; i++ {
        fmt.Print(fibonacci(i), " ")
    }
    fmt.Println()
}