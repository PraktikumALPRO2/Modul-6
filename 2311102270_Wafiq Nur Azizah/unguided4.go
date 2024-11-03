package main

import (
    "fmt"
)

func printSequence(n int) {
    if n < 1 {
        return
    }
    fmt.Print(n, " ")
    if n > 1 {
        printSequence(n - 1)
        fmt.Print(n, " ")
    }
}

func main() {
    var N int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&N)

    printSequence(N)
    fmt.Println()
}
