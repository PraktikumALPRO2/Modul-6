package main

import "fmt"

// Fungsi rekursif untuk menghitung Fibonacci ke-n
func fibonacci(n int) int {
     if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}

func main() {
    // Menampilkan deret Fibonacci hingga suku ke-10
	var n int 
	print("masukan bilangan: ")
	fmt.Scanln(&n)
    for i := 0; i <= n; i++ {
        fmt.Printf("Fibonacci ke-%d: %d\n", i, fibonacci(i))
    }
    
}
