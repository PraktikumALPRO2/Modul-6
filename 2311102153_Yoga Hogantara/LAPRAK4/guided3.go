package main

import "fmt"

func pangkat(n int)int{
	if n== 0{
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
func main(){
	var n int
	fmt.Print("MASUKAN NILAI n: ")
	fmt.Scan(&n)
	fmt.Println("hasil dari 2 pangkat",n,"adalah :",pangkat(n))
}