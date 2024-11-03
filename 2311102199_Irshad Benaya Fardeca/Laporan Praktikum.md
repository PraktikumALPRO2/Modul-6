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

Fungsi dapat bersifat rekursif, artinya mereka dapat memanggil diri mereka sendiri, baik secara langsung maupun tidak langsung. Rekursif adalah teknik yang kuat untuk banyak masalah dan tentu saja sangat penting dalam pemrosesan struktur data rekursif.

rekursi adalah proses pengulangan item dengan cara yang serupa. konsep yang sama juga berlaku dalam bahasa pemograman. jika sebuah program memungkinkan untuk memanggil fungsi di dalam fungsi yang sama, maka itu disebut pemanggilan fungsi rekursif. contoh

func rekursif(){
  /* fungsi memanggil dirinya sendiri */
  rekursif()
}

func main(){
  rekursif()
}


bahasa pemograman Go mendukung rekursi. artinya, memungkinkan suatu fungsi untuk memanggil dirinya sendiri. tetapi saat menggunakan rekursi, pemogram harus berhati-hati untuk mementukan kondisi keluar dari fungsi, jika tidak maka akan menjadi loop tidak terbatas.



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
Program ini dirancang untuk menghitung jumlah dari semua bilangan bulat positif mulai dari 1 hingga suatu bilangan tertentu yang diinputkan pengguna. Fungsi penjumlahan menggunakan rekursi untuk memecahkan masalah ini. Jika bilangan yang diberikan adalah 1, fungsi langsung mengembalikan 1. Jika tidak, fungsi akan menjumlahkan bilangan tersebut dengan hasil rekursif dari bilangan yang satu lebih kecil. Fungsi main berfungsi sebagai titik masuk program, meminta pengguna untuk memasukkan sebuah bilangan, kemudian memanggil fungsi penjumlahan dan mencetak hasilnya ke layar.

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
Program ini dirancang untuk menghitung nilai pangkat dua dari suatu bilangan bulat non-negatif yang diinputkan pengguna. Fungsi pangkat secara rekursif mengalikan bilangan 2 sebanyak n kali, di mana n adalah nilai yang diberikan pengguna. Jika n sama dengan 0, fungsi akan mengembalikan nilai 1 sesuai dengan definisi pangkat nol. Hasil perhitungan kemudian ditampilkan ke layar bersama dengan nilai n yang diinputkan pengguna.

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
Program ini dirancang untuk menghitung faktorial dari sebuah bilangan bulat non-negatif yang diinputkan oleh pengguna. Program dimulai dengan meminta pengguna memasukkan sebuah bilangan. Bilangan ini kemudian akan diproses oleh fungsi faktorial yang menggunakan konsep rekursif. Fungsi ini akan terus memanggil dirinya sendiri dengan nilai n yang dikurangi 1 hingga mencapai kondisi dasar, yaitu ketika n bernilai 0 atau 1. Hasil perkalian dari setiap pemanggilan fungsi inilah yang kemudian akan menjadi nilai faktorial dari bilangan yang diinputkan. Hasil akhir dari perhitungan faktorial tersebut akan ditampilkan di layar.

<br></br>


#### III. UNGUIDED

##### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program


##### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program

##### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program


##### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program


##### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program


##### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y. Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".

#### Source Code
```go

```
##### Screenshoot Output

##### Deskripsi Program




### Referensi
[1] Donovan, A., Kernighan, B. (2015). The Go Programming Language. United Kingdom: Pearson Education.

