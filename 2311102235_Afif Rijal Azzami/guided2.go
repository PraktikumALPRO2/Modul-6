package main
import (
	"fmt"
)

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n - 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukan bilingan: ")
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}