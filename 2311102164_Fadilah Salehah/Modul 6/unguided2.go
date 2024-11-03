package main

import (
	"fmt"
)

func cetakSegitiga(n int) {
	if n == 0 {
		return
	}
	cetakSegitiga(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&n)
	cetakSegitiga(n)
}