package main

import "fmt"

func printNumbers(n, current int, descending bool) {
    fmt.Print(current, " ")

    if descending {
        if current == 1 {
            descending = false
        } else {
            printNumbers(n, current-1, true)
            return
        }
    }

    if !descending && current < n {
        printNumbers(n, current+1, false)
    }
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&n)

    fmt.Print("Keluaran: ")
    printNumbers(n, n, true)
    fmt.Println()
}
