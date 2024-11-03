package main

import "fmt"

// Fungsi rekursif untuk mencari faktor bilangan
func printFactors(n, current int) {
	if current > n {
		return
	}
	if n%current == 0 {
		fmt.Printf("%d ", current)
	}
	printFactors(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d: ", n)
	printFactors(n, 1)
	fmt.Println()
}