package main

import "fmt"

func printFactors(n, divisor int) {
    if divisor > n {
        return
    }
    if n % divisor == 0 {
        fmt.Print(divisor, " ")
    }
    printFactors(n, divisor + 1)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&n)

    fmt.Print("Faktor dari ", n, " adalah: ")
    printFactors(n, 1)
    fmt.Println()
}
