package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	if n > 0 {
		baris(n)
	} else {
		fmt.Println("Masukkan bilangan positif")
	}
}

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else if bilangan > 1 {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
