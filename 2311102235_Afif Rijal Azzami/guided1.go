package main
import (
	"fmt"
)

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Print(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}

func main() {
	var n int
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&n)
	baris(n)
}