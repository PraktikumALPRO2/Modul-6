package main

import "fmt"

func main() {

    // Inputan user untuk bilangan 
    var n int
    fmt.Scan(&n)
    // untuk mencetak baris
    baris(n)

}

// fungsi rekursif untuk baris
func baris(bilangan int) {
    if bilangan == 1 {
        fmt.Println(1)
    } else {
        fmt.Println(bilangan)
        baris(bilangan - 1)
    }
}