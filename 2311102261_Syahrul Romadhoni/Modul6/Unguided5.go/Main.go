package main

import "fmt"

func printOddNumbers(n, current int) {
    if current > n {
        return
    }
    fmt.Print(current, " ")
    printOddNumbers(n, current + 2)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&n)

    fmt.Print("Keluaran: ")
    printOddNumbers(n, 1)
    fmt.Println()
}
