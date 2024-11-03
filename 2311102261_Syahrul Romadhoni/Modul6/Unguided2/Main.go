package main

import "fmt"

func printStars(n int) {
    if n > 0 {
        fmt.Print("*")
        printStars(n - 1)
    }
}

func printPattern(n int, current int) {
    if current > n {
        return
    }
    printStars(current)
    fmt.Println()
    printPattern(n, current + 1)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N untuk pola bintang: ")
    fmt.Scan(&n)

    printPattern(n, 1)
}
