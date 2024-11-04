package main
import"fmt"

//fungsi rekursif untuk menghitung 2^n
func pangkat(n int) int{
	if n == 0{
		return 1
	} else {
		return 2 * pangkat (n-1)
	}
}
func main(){
	var n int
	fmt.Print("masukkan nilai n :")
	fmt.Scan(&n)
	fmt.Println("hasil dari 2 pangkat", n, "adalah : ", pangkat(n))
}