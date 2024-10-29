package main
import "fmt"

var n int
func faktorial (n int) int {
	if n == 0 || n ==1 {  // n =3
		return 1
	} else {
		return n * faktorial(n-1) // 3 * 2 * 1 = 6 
	}
}

func main () {
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}