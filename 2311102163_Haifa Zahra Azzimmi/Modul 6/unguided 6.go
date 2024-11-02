package main

import (
    "fmt"
)

// Fungsi rekursif untuk menghitung x^y
func power(x int, y int) int {
    if y == 0 {
        return 1
    }
    return x * power(x, y-1)
}

func main() {
    var x, y int

    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan (x): ")
    fmt.Scan(&x)
    fmt.Print("Masukkan pangkat (y): ")
    fmt.Scan(&y)

    // Menampilkan hasil
    fmt.Printf("%d ^ %d = %d\n", x, y, power(x, y))
}
