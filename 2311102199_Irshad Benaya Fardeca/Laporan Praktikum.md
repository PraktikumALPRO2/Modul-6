<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 6</h2>
<h2 align="center">REKURSIF</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI

Fungsi dapat bersifat rekursif, artinya fungsi dapat memanggil dirinya sendiri, baik secara langsung maupun tidak langsung. Rekursif adalah teknik yang kuat untuk banyak masalah dan tentu saja sangat penting dalam pemrosesan struktur data rekursif. Contoh : 

```
func rekursif(){
  /* fungsi memanggil dirinya sendiri */
  rekursif()
}

func main(){
  rekursif()
}
```

Bahasa pemograman Go mendukung rekursif. artinya, memungkinkan suatu fungsi untuk memanggil dirinya sendiri. tetapi saat menggunakan rekursif, pemogram harus berhati-hati untuk mementukan kondisi keluar dari fungsi, jika tidak maka akan menjadi loop tidak terbatas.



<br></br>


#### II. GUIDED
##### 1. Menampilkan baris suatu bilangan

##### Source code
```go
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

	fmt.Scan(&n)
	baris(n)
}

```
##### Screenshoot Output
![Screenshot 2024-11-01 222002](https://github.com/user-attachments/assets/f505f67a-a120-4b85-914e-6827830a7062)

##### Deskripsi Program
Program ini dirancang untuk mencetak angka dari suatu bilangan input secara menurun hingga mencapai angka 1. Proses pencetakan ini dilakukan secara rekursif, di mana fungsi baris memanggil dirinya sendiri dengan parameter yang lebih kecil hingga mencapai kondisi dasar ketika bilangan sama dengan 1.

##### 2. Penjumlahan baris bilangan
##### Source code
```go
package main

import "fmt"

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

```
##### Screenshoot Output
![Screenshot 2024-11-01 222136](https://github.com/user-attachments/assets/1aa6dcb6-9791-4511-8816-f807a4bc868e)

##### Deskripsi Program
Program ini dirancang untuk menghitung jumlah dari semua bilangan bulat positif mulai dari 1 hingga suatu bilangan tertentu yang diinputkan user. Fungsi penjumlahan menggunakan rekursif untuk memecahkan masalah ini. Jika bilangan yang diberikan adalah 1, fungsi langsung mengembalikan 1. Jika tidak, fungsi akan menjumlahkan bilangan tersebut dengan hasil rekursif dari bilangan yang satu lebih kecil. Fungsi main berfungsi sebagai titik masuk program, meminta user untuk memasukkan sebuah bilangan, kemudian memanggil fungsi penjumlahan dan mencetak hasilnya ke layar.

##### 3. Hasil dari 2^n

##### Source code
```go
package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
func main() {
	var n int
	fmt.Print("Masukkan Nilai n : ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))
}

```
##### Screenshoot Output
![Screenshot 2024-11-01 222240](https://github.com/user-attachments/assets/d41255a2-648c-4235-b9b0-ce5a1da182ff)

##### Deskripsi Program
Program ini dirancang untuk menghitung nilai pangkat dua dari suatu bilangan bulat non-negatif yang diinputkan user. Fungsi pangkat secara rekursif mengalikan bilangan 2 sebanyak n kali, di mana n adalah nilai yang diberikan user. Jika n sama dengan 0, fungsi akan mengembalikan nilai 1 sesuai dengan definisi pangkat nol. Hasil perhitungan kemudian ditampilkan ke layar bersama dengan nilai n yang diinputkan user.

##### 4. Menghitung faktorial suatu bilangan

##### Source code
```go
package main

import "fmt"

var n int

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {

	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}

```
##### Screenshoot Output
![Screenshot 2024-11-01 222328](https://github.com/user-attachments/assets/b3213121-69e0-48c1-9866-178f591669ae)

##### Deskripsi Program
Program ini dirancang untuk menghitung faktorial dari sebuah bilangan bulat non-negatif yang diinputkan oleh user. Program dimulai dengan meminta user memasukkan sebuah bilangan. Bilangan ini kemudian akan diproses oleh fungsi faktorial yang menggunakan konsep rekursif. Fungsi ini akan terus memanggil dirinya sendiri dengan nilai n yang dikurangi 1 hingga mencapai kondisi dasar, yaitu ketika n bernilai 0 atau 1. Hasil perkalian dari setiap pemanggilan fungsi inilah yang kemudian akan menjadi nilai faktorial dari bilangan yang diinputkan. Hasil akhir dari perhitungan faktorial tersebut akan ditampilkan di layar.

<br></br>


#### III. UNGUIDED

##### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### Source Code
```go
package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	fmt.Print("Deret Fibonaccinya = ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibo(i))
	}
	fmt.Println()
}

```
##### Screenshoot Output
![Screenshot 2024-11-03 212901](https://github.com/user-attachments/assets/1ef2a8da-ab2c-4935-a534-1c12715a4531)

##### Deskripsi Program
Program ini dirancang untuk menghasilkan deret Fibonacci hingga suatu bilangan tertentu. Program dimulai dengan meminta user untuk memasukkan sebuah bilangan bulat. Bilangan ini kemudian akan menjadi batas atas dari deret Fibonacci yang akan dihasilkan. Fungsi fibo digunakan untuk menghitung nilai setiap suku dalam deret Fibonacci secara rekursif. Fungsi main akan mencetak deret Fibonacci mulai dari suku ke-0 hingga suku ke-n sesuai dengan input user. Dengan demikian, program ini memberikan cara yang efisien untuk mendapatkan deret Fibonacci dari sebuah bilangan tertentu.

##### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
![Screenshot 2024-11-03 210917](https://github.com/user-attachments/assets/6e1cf318-7edd-40e5-8a07-0b845bf6d0f4)

#### Source Code
```go
package main

import "fmt"

func bintang(n int) {
	if n <= 0 {
		return
	}

	bintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	bintang(n)
}

```
##### Screenshoot Output
![Screenshot 2024-11-03 212659](https://github.com/user-attachments/assets/5b41892f-5148-42c2-8add-f973140baad2)

##### Deskripsi Program
Program ini dirancang untuk mencetak pola menggunakan tanda bintang. User diminta untuk memasukkan sebuah bilangan bulat yang akan menentukan tinggi dari segitiga tersebut. Fungsi bintang bekerja secara rekursif untuk mencetak setiap baris dari bintang. Pada setiap pemanggilan fungsi, program akan terlebih dahulu memanggil dirinya sendiri dengan nilai n yang dikurangi 1, sehingga membentuk pola bintang yang semakin besar ke bawah. Setelah pemanggilan rekursif, program akan mencetak sejumlah tanda bintang sesuai dengan nilai n pada baris tersebut. Dengan demikian, program akan mencetak pola bintang.

##### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

#### Source Code
```go
package main

import "fmt"

func faktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
		faktor(n, i+1)
	}
}

func main() {
	var bil int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bil)

	fmt.Printf("Faktor dari %d adalah : ", bil)
	faktor(bil, 1)
	fmt.Println()
}
```
##### Screenshoot Output
![Screenshot 2024-11-03 212528](https://github.com/user-attachments/assets/5a6078dc-7c11-463f-8b5b-e44cc7b4fb2e)

##### Deskripsi Program
Program ini dirancang untuk mencari dan mencetak semua faktor dari sebuah bilangan bulat positif yang diinputkan oleh user. Program ini bekerja dengan cara rekursif, di mana fungsi faktor akan terus memanggil dirinya sendiri untuk memeriksa setiap bilangan mulai dari 1 hingga bilangan yang diinputkan. Jika suatu bilangan merupakan faktor dari bilangan yang diinputkan (artinya, hasil bagi keduanya adalah bilangan bulat), maka bilangan tersebut akan dicetak. Proses ini berulang hingga semua kemungkinan faktor telah diperiksa. Dengan demikian, program ini akan memberikan daftar lengkap semua faktor dari bilangan yang diberikan.

##### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

#### Source Code
```go
package main

import "fmt"

func baris(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	baris(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)

	fmt.Printf("Barisan bilangan dari %d : ", bilangan)
	baris(bilangan)
	fmt.Println()
}

```
##### Screenshoot Output
![Screenshot 2024-11-03 212318](https://github.com/user-attachments/assets/05323282-04d8-4f5f-9d63-84dfbc3533cd)

##### Deskripsi Program
Program ini dirancang untuk mencetak sebuah barisan bilangan dari suatu bilangan bulat yang diinputkan oleh user, mulai dari bilangan tersebut hingga 1, lalu kembali lagi ke bilangan tersebut. Program akan meminta user untuk memasukkan sebuah bilangan bulat. Kemudian, fungsi baris akan dipanggil secara rekursif untuk mencetak barisan bilangan tersebut. Fungsi baris bekerja dengan cara mencetak bilangan saat ini, kemudian memanggil dirinya sendiri dengan nilai n yang dikurangi 1. Setelah mencapai bilangan 1, fungsi akan kembali mencetak bilangan-bilangan sebelumnya dalam urutan yang berlawanan.

##### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

#### Source Code
```go
package main

import "fmt"

func ganjil(n, cur int) {
	if cur > n {
		return
	}

	if cur%2 != 0 {
		fmt.Print(cur, " ")
	}

	ganjil(n, cur+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d : ", n)
	ganjil(n, 1)
	fmt.Println()
}

```
##### Screenshoot Output
![Screenshot 2024-11-03 212153](https://github.com/user-attachments/assets/27213a1a-df5f-49cb-91d8-f664c88c483d)

##### Deskripsi Program
Program ini dirancang untuk mencetak semua bilangan ganjil dari 1 hingga suatu bilangan bulat tertentu yang diinputkan oleh user. Program bekerja dengan cara rekursif, di mana fungsi ganjil terus memanggil dirinya sendiri untuk memeriksa setiap bilangan mulai dari 1 hingga bilangan yang diinputkan. Jika bilangan tersebut ganjil, maka akan langsung dicetak. Proses ini berulang hingga semua bilangan telah diperiksa. 

##### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y. Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".

#### Source Code
```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else if y > 0 {
		return x * pangkat(x, y-1)
	} else {
		return 1 / pangkat(x, -y)
	}
}

func main() {
	var bil1, bil2 int

	fmt.Print("Masukkan bilangan (x) : ")
	fmt.Scan(&bil1)
	fmt.Print("Masukkan bilangan (y) : ")
	fmt.Scan(&bil2)

	fmt.Printf("%d pangkat %d = %d\n", bil1, bil2, pangkat(bil1, bil2))
}

```
##### Screenshoot Output
![Screenshot 2024-11-03 211459](https://github.com/user-attachments/assets/a7e44cef-e13b-44c4-9dba-72a0b0554a2a)

##### Deskripsi Program
Program ini dirancang untuk menghitung hasil pangkat dari dua bilangan bulat yang diinputkan oleh user. Program akan meminta user memasukkan bilangan pokok (x) dan pangkat (y). Fungsi pangkat kemudian akan menghitung hasil pangkat dengan menggunakan rekursif. Fungsi ini memiliki tiga kondisi: jika pangkat (y) adalah 0, maka hasilnya adalah 1; jika pangkat (y) positif, maka hasil dihitung dengan mengalikan bilangan pokok dengan hasil pangkat dari bilangan pokok yang dipangkatkan dengan y-1; dan jika pangkat (y) negatif, maka hasil dihitung dengan membagi 1 dengan hasil pangkat dari bilangan pokok yang dipangkatkan dengan nilai absolut dari y. Hasil akhir dari perhitungan pangkat tersebut kemudian akan ditampilkan di layar.


### Referensi
[1] Donovan, A., Kernighan, B. (2015). The Go Programming Language. United Kingdom: Pearson Education.
[2] Bellshade, Rekursif (2021), https://www.australianhumanitiesreview.org/archive/Issue-June-1997/dessaix.html
