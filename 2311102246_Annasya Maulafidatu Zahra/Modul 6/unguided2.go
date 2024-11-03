package main

import "fmt"

func cetakBintang(jumlah int) {
	if jumlah > 0 {
		fmt.Print("*")
		cetakBintang(jumlah - 1)
	} else {
		fmt.Println()
	}
}

func polaBintang(n, baris int) {
	if baris <= n {
		cetakBintang(baris)
		polaBintang(n, baris+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)
	polaBintang(n, 1)
}
