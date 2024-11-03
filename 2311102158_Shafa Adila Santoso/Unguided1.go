package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci pada suku ke-n
func fibonaci_158(n_158 int) int {
	if n_158 <= 1 {
		return n_158
	}
	return fibonaci_158(n_158-1) + fibonaci_158(n_158-2)
}

func main() {
	var n_158 int

	// Meminta input dari pengguna untuk menentukan jumlah suku
	fmt.Print("n: ")
	fmt.Scan(&n_158)

	// Slice untuk menyimpan deret Fibonacci
	simpanBilangan := make([]int, n_158+1)

	for i := 0; i <= n_158; i++ {
		simpanBilangan[i] = fibonaci_158(i)
	}

	// Mencetak deret Fibonacci dalam format yang diinginkan
	fmt.Printf("F : ")
	for i, val := range simpanBilangan {
		if i == n_158 {
			fmt.Printf("%d", val) 
		} else {
			fmt.Printf("%d, ", val) 
		}
	}
	fmt.Println() 
}
