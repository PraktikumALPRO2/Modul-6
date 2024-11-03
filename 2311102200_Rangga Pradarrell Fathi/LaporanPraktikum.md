<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL III</strong>
  <br>
  <strong>Rekursif</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
 Rangga Pradarrell Fathi
  <br>
  2311102200
  <br>
  IF-11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

<strong><h2>Definisi Rekursif</h2></strong>
Rekursif merupakan konsep pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan suatu masalah. Fungsi rekursif akan terus memanggil dirinya sendiri hingga mencapai suatu kondisi dasar (base case) yang memungkinkan fungsi tersebut berhenti dan mengembalikan nilai.

### <strong> Beberapa karakteristik utama dari rekursif:

- Base Case: Setiap fungsi rekursif harus memiliki kondisi dasar atau base case yang menghentikan rekursi. Jika tidak ada base case, rekursi akan terus berlanjut tanpa batas.
- Pengurangan Masalah: Setiap pemanggilan rekursif harus mengurangi masalah menjadi sub-masalah yang lebih kecil dan mendekati base case.
- Pemanggilan Diri Sendiri: Fungsi rekursif memanggil dirinya sendiri dengan argumen yang berbeda untuk menyelesaikan sub-masalah.
- Kemampuan Memecah Masalah: Rekursif sangat efektif untuk memecah masalah yang kompleks menjadi sub-masalah yang lebih sederhana dan dapat diselesaikan secara terpisah.

### <strong> Contoh penerapan rekursif dalam pemrograman antara lain:
- Menghitung Faktorial: Fungsi rekursif untuk menghitung faktorial suatu bilangan.
   **Faktorial**:
   ```go
   func factorial(n int) int {
       if n == 0 {
           return 1
       }
       return n * factorial(n-1)
   }
   ```
   Fungsi `factorial()` memanggil dirinya sendiri dengan nilai `n-1` sampai `n` mencapai 0, lalu mengembalikan hasil perkalian dari `n` sampai 1.
  
- Deret Fibonacci: Fungsi rekursif untuk menghitung suku ke-n dari deret Fibonacci.
   **Deret Fibonacci**:
   ```go
   func fibonacci(n int) int {
       if n <= 1 {
           return n
       }
       return fibonacci(n-1) + fibonacci(n-2)
   }
   ```
   Fungsi `fibonacci()` memanggil dirinya sendiri dengan `n-1` dan `n-2` untuk menghitung suku ke-n dalam deret Fibonacci.
  
- Pencarian Biner (Binary Search): Fungsi rekursif untuk melakukan pencarian elemen dalam array yang terurut.
 **Pencarian Biner (Binary Search)**:
   ```go
   func binarySearch(arr []int, target, left, right int) int {
       if left > right {
           return -1 // Target tidak ditemukan
       }
       mid := (left + right) / 2
       if arr[mid] == target {
           return mid
       } else if arr[mid] < target {
           return binarySearch(arr, target, mid+1, right)
       } else {
           return binarySearch(arr, target, left, mid-1)
       }
   }
   ```
   Fungsi `binarySearch()` memanggil dirinya sendiri dengan rentang indeks yang semakin sempit untuk menemukan target dalam array yang terurut.

# <strong> Guided </strong>
## Guided - 1
### Study Case
**Membuat baris bilangan dari n hingga 1.** 
**Base-case: bilangan == 1**

### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import "fmt"

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")	
	fmt.Scan(&n)
	baris(n)
}
```

### Screenshot Output
![Guided 1](https://github.com/user-attachments/assets/acfc97ef-0450-4912-bfa9-3f5c7ab7defb)


### Deskripsi Program

Program diatas merupakan program untuk mencetak bilangan bulat dan nilai yang diinputkan oleh user hingga 1 secara terbalik. Program ini menggunakan fungsi rekursif untuk mencetak bilangan 1 per 1.

Algoritma dari program tersebut sebagai berikut :
1. Minta inputan dari user untuk bilangan 'n'.
2. Panggil fungsi 'baris'.
3. fungsi 'baris', jika 'bilangan' = 1, maka cetak 1.
4. Jika tidak, maka cetak bilangan - 1.

Cara kerja dari program ini yaitu user menginputkan bilangan. Setelah itu fungsi 'baris' digunakan untuk mencetak bilangan mulai dari n hingga 1. Fungsi ini dicetak kemudian memanggil dirinya sendiri dengan nilai berkurang 1 hingga 1.

## Guided - 2
### Study Case
**Menghitung hasil penjumlahan 1 hingga n, Base-Case n == 1**
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import "fmt"

// fungsi rekursif untuk penjumlahan
func penjumlahan (n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}

func main() {

	// Inputan user untuk bilangan
	var n int
	fmt.Scan(&n)
	// Tampilkan hasil penjumlahan
	fmt.Println(penjumlahan(n))
	
}
```
### Screenshot Output
![Guided 2](https://github.com/user-attachments/assets/321c6b05-275a-494a-8298-2936347cd4d1)


### Deskripsi Program
Program diatas merupakan program  untuk menghitung jumlah bilangan hingga nilai yang diinputkan oleh user. Program menggunakan fungsi rekursif untuk menjumlahkan bilangan.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'penjumlahan' untuk melakukan kondisi.
2. Jika n = 1, kembalikan 1.
3. Jika tidak, kembalikan hasil 'n' dan hasil 'penjumlahan(n-1).
4. Minta input pengguna untuk bilangan.
5. Tampilkan hasil penjumlahan dengan fungsi 'penjumlahan(n).

Cara kerja dari program ini yaitu user menginputkan bilangan. Fungsi 'penjumlahan' dipanggil untuk menghitung 1 hingga nilai dari user. Fungsi ini dijalankan dengan menjumlahkan 'n' dan hasil dirinya sendiri dengan 'n-1' hingga 1.

## Guided - 3
### Study Case 
**Mencari dua pangkat n atau 2^n, Base-Case n == 0**
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main 

import "fmt"

// Fungsi rekursif untuk menghitung pangkat
func pangkat(n int) int {
	if n == 0{
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}

func main () {

	// Inputan user untuk bilangan
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	// Tampilkan hasil pangkat
	fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))

}
```
### Screenshot Output
![Guided 3](https://github.com/user-attachments/assets/34ce676b-b011-44e7-9538-35b3bd624113)


### Deskripsi Program
Program diatas merupakan program untuk menghitung hasil dari 2 pangkat, yang dimana 'n' merupakan bilangan yang diinputkan oleh user. Program ini menggunakan fungsi rekursif untuk menghitung pangkat.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'pangkat' untuk menghitung ^2.
2. Jika n = 0, kembalikan 1.
3. Jika tidak, kembalikan hasil 'pangkat (n-1).
4. Minta inputan dari pengguna untuk nilai.
5. Tampilkan hasil pangkat dengan fungsi 'pangkat(n)'.

Cara kerja dari program ini yaitu user menginputkan bilangan. Setelah itu, panggil fungsi 'pangkat' untuk menghitung hasil 2^n secara rekursif. Fungsi memanggil dirinya sendiri dengan nilai 'n' yang berkurang hingga 0. Fungsi akan mengalikan 2 dengan hasil pangkat sebelumnya.

## Guided - 4
### Study Case
**Mencari nilai faktorial atau n!, Base-Case n == 0 atau n == 1**
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import "fmt"

var n int
// Fungsi untuk menghitung faktorial
func faktorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main () {

	// Inputan user untuk bilangan
	fmt.Scan(&n)
	// Tampilkan nilai faktorial
	fmt.Println(faktorial(n))
}
```
### Screenshot Output
![Guided 4](https://github.com/user-attachments/assets/98761df1-a1f8-42c7-be9b-cff2eed13b84)


### Deskripsi Program
Program diatas merupakan program untuk menghitung faktorial dari bilangan yang diinputkan oleh user. Faktorial ditulis n! adalah hasil perkalian dari bilangan bulat dari 1 hingga n.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'faktorial' untuk menghitung faktorial.
2. Jika n = 0 atau 1, kembalikan 1.
3. Jika tidak, kembalikan hasil dari n! 'faktorial(n-1).
4. Minta inoutan dari user untuk bilangan.
5. Tampilkan hasil faktorial dengan fungsi 'faktorial(n)'.

Cara kerja dari program ini yaitu user menginputkan bilangan. Fungsi 'faktorial' digunakan untuk menghitung faktorial dari bilangan secara rekursif. Fungsi akan memanggil dengan nilai n yang berkurang hingga 0 atau 1. Setiap panggilan akan mengalikan n dengan hasil faktorial dari n-1.


# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S=Sn-1+Sn-2 Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonассі tersebut.**

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func FibonacciAngka(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciAngka(n-1) + FibonacciAngka(n-2)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci yang ingin dicetak: ")
	fmt.Scanln(&n)

	fmt.Println("Deret Fibonacci hingga suku ke-", n)
	for i := 0; i < n; i++ {
		fmt.Print(FibonacciAngka(i), " ")
	}
	fmt.Println()
}
```
### Screenshot Code


### Deskripsi Program


## Unguided - 2
### Study Case
**Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.**

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func printBintang(n int) {
	if n == 0 {
		return
	}
	printBintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scanln(&n)

	printBintang(n)
}
```
### Screenshot Code


### Deskripsi Program


## Unguided - 3
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.**                                        
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                        
*Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func findFactors(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	findFactors(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	findFactors(n, 1)
	fmt.Println()
}
```
### Screenshot Code


### Deskripsi Program


## Unguided - 4
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.**                             
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                                 
*Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func printSekuensial(n int) {
	if n == 1 {
		fmt.Print(n, " ")
	} else {
		fmt.Print(n, " ")
		printSekuensial(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Print("Keluaran: ")
	printSekuensial(n)
	fmt.Println()
}
```
### Screenshot Code


### Deskripsi Program


## Unguided - 5
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.**                                 
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                        
*Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func cetakBilanganGanjil(n int) {
	if n >= 1 {
		cetakBilanganGanjil(n - 2)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan ganjil dari 1 hingga", n, "adalah:")
	cetakBilanganGanjil(n)
}
```
### Screenshot Code


### Deskripsi Program

## Unguided - 6
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.**                                     
*Masukan terdiri dari bilangan bulat x dan y.*                                        
*Keluaran terdiri dari hasil x dipangkatkan y.*    

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	hasil := 0
	for i := 0; i < y; i++ {
		if i == 0 {
			hasil = x
		} else {
			hasil += x 
		}
	}
	return hasil
}

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scanln(&y)

	hasil := pangkat(x, y)
	fmt.Printf("%d pangkat %d = %d\n", x, y, hasil)
}
```
### Screenshot Code


### Deskripsi Program
