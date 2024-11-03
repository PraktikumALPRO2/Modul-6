package main 
import "fmt"

func bintang_158(n_158 int) {
	for i := 1; i <= n_158; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main () {
	var n_158 int
	fmt.Print("Mausukkan n: ")
	fmt.Scan(&n_158)

	fmt.Println("Pola bintnag terurut: ")
	bintang_158(n_158)
}