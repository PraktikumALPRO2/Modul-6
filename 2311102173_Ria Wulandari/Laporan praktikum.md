

<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL VI</strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Ria Wulandari / 2311102173<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)

## Dasar Teori
#### Pengertian Rekursif
Rekursif adalah metode dalam pemrograman di mana sebuah fungsi memanggil dirinya sendiri secara berulang sampai mencapai kondisi tertentu yang disebut basis rekursi atau base case. Setiap kali fungsi memanggil dirinya sendiri, fungsi tersebut bekerja pada versi yang lebih sederhana atau lebih kecil dari masalah yang sama, hingga mencapai titik di mana pemanggilan lebih lanjut tidak diperlukan dan fungsi dapat mengembalikan hasil.

Rekursi merupakan salah satu cara menyelesaikan masalah yang dapat dipecah menjadi sub-masalah yang lebih kecil dan serupa dengan masalah aslinya. Teknik ini banyak digunakan dalam algoritma karena sering kali menghasilkan solusi yang elegan dan sederhana untuk masalah yang kompleks, terutama yang memiliki sifat berulang seperti perhitungan matematis, manipulasi struktur data, atau algoritma pencarian dan penyortiran.
#### Komponen Rekursif
Untuk memahami cara kerja rekursif, ada dua komponen penting yang harus diperhatikan:

1. Base Case (Kondisi Dasar): Kondisi yang menghentikan rekursi. Fungsi rekursif harus memiliki satu atau lebih base case yang memastikan rekursi tidak berjalan tanpa batas. Base case ini akan mengembalikan nilai tertentu tanpa melakukan pemanggilan rekursif lagi. Tanpa base case, program akan mengalami infinite recursion (rekursi tanpa akhir) yang berpotensi menyebabkan kesalahan atau kehabisan memori (stack overflow).

2. Recursive Case (Kasus Rekursif): Bagian dari fungsi yang memanggil dirinya sendiri dengan versi yang lebih kecil atau lebih sederhana dari masalah yang sama. Pada setiap pemanggilan rekursif, fungsi tersebut semakin mendekati base case, sehingga pada akhirnya rekursi akan berhenti.

#### Contoh sederhana Rekursif
Contoh klasik dari rekursif adalah menghitung faktorial dari suatu bilangan. Secara matematis, faktorial dari sebuah bilangan 𝑛, yang dilambangkan sebagai n!, dapat didefinisikan sebagai:

n!=n x(n-1) x (n-2) x ... x 1

Dalam bentuk rekursif, faktorial dapat dituliskan sebagai :
- Base case 0!=1
- Recursive case n!=n x (n-1)!

#### Berikut contoh implementasi faktorial menggunakan rekursi:
```go
package main

import "fmt"

// Fungsi rekursif untuk menghitung faktorial
func faktorial(n int) int {
    // Base case: 0! = 1
    if n == 0 {
        return 1
    }
    // Recursive case: n! = n * (n-1)!
    return n * faktorial(n-1)
}

func main() {
    // Menghitung faktorial dari 5
    fmt.Println(faktorial(5))  // Output: 120
}
```
Penjelasan :
1. Base case : jika nilai n=0, fungsi akan mengembalikan 1, karena 0!=1.
2. Recursive case : jikan n > 0, fungsi akan memanggil dirinya sendiri dengan parameter n-1.
3. Pada fungsi `main`, faktorial dari 5 dihitung dengan `faktorial(5)`.

## Guided
### 1. Membuat baris bilangan dari n hingga 1
Base-case: bilangan == 1
#### Source code :
```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	baris(n)
}
func baris(bilangan int){
	if bilangan == 1 {
		fmt.Println(1)
	}else{
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
```
#### Deskripsi program
Program ini menggunakan rekursi untuk mencetak deret angka dari nilai input hingga angka 1 secara menurun. Program ini meminta input dari pengguna, yaitu bilangan bulat n, dan kemudian mencetak angka mulai dari n hingga 1. Fungsi `baris()` dipanggil secara rekursif untuk menampilkan setiap angka hingga mencapai nilai 1.
#### Algoritma program :
1. Program meminta input dari pengguna berupa bilangan bulat.

2. Fungsi `baris()` dipanggil dengan parameter input tersebut.

3. Di dalam fungsi `baris()`, program akan mencetak angka yang dikirim sebagai parameter, kemudian memanggil dirinya sendiri dengan nilai parameter yang dikurangi 1.

4. Proses rekursif berlanjut sampai nilai parameter mencapai 1, di mana angka 1 dicetak dan rekursi berhenti.
#### Cara kerja program :
1. Minta Input: Program meminta pengguna memasukkan sebuah angka, misalnya 5.

2. Cetak Angka: Program memanggil fungsi `baris()`, yang pertama-tama mencetak angka 5.

3. Rekursi: Fungsi `baris()` kemudian memanggil dirinya sendiri dengan angka yang lebih kecil, yaitu 4, lalu 3, 2, dan akhirnya 1.

4. Selesai: Ketika angka 1 sudah dicetak, program berhenti.
### 2. Menghitung hasil penjumlahan 1 hingga n
Base-case: n = 1
#### Source code :
```go
package main
import "fmt"

func penjumlahan(n int)int{
	if n == 1{
		return 1
	}else{
		return n + penjumlahan(n-1)
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}
```
#### Deskripsi program
Program ini menghitung penjumlahan dari semua bilangan bulat dari 1 hingga nilai input n menggunakan rekursi. Program meminta input dari pengguna, yaitu bilangan bulat n, dan kemudian memanggil fungsi `penjumlahan()` untuk menghitung hasil penjumlahan rekursif dari 1 hingga n.
#### Algoritma program :
1. Program meminta input dari pengguna berupa bilangan bulat n.

2. Fungsi `penjumlahan()` dipanggil dengan parameter n, yang menghitung penjumlahan semua bilangan dari 1 hingga n secara rekursif.

3. Fungsi `penjumlahan()` memiliki dua kasus:
- Base case: Jika n=1, fungsi akan mengembalikan 1.
- Recursive case: Jika n>1, ungsi akan memanggil dirinya sendiri dengan parameter n-1, dan menambahkan nilai n pada hasil penjumlahan rekursif dari n-1.

4. Hasil penjumlahan ditampilkan menggunakan `fmt.Println()`.
#### Cara kerja program :
1. Minta Input: Program meminta pengguna memasukkan sebuah angka, misalnya 5.

2. Rekursi: Program menghitung penjumlahan dari 1 hingga angka tersebut dengan memanggil fungsi `penjumlahan()`.

3. Proses Rekursi: Fungsi menambahkan angka yang dimasukkan dengan hasil dari fungsi yang dipanggil lagi dengan angka lebih kecil, hingga mencapai 1.

4. Cetak Hasil: Ketika sampai di angka 1, program menjumlahkan semua angka dan mencetak hasilnya.
### 3. Mencari dua pangkat n atau 2^n
Base-case: n == 0
#### Source code :
```go
package main

import "fmt"

//Fungsi rekrusif untuk menghitung 2^n
func pangkat(n int) int{
	if n == 0{
		return 1
	}else{
		return 2*pangkat(n-1)
	}
}

func main(){
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah: ", pangkat(n))
}
```
#### Deskripsi program
Program ini menghitung nilai 2^n (dua pangkat n) menggunakan fungsi rekursif. Program akan meminta input bilangan bulat nari pengguna, kemudian memanggil fungsi rekursif `pangkat()` untuk menghitung hasil dari 2^n.
#### Algoritma program :
1. Program meminta pengguna untuk memasukkan bilangan bulat n.

2. Fungsi `pangkat()` dipanggil dengan parameter n untuk menghitung hasil dari 2^n.

3. Di dalam fungsi `pangkat()`, ada dua kasus:
- Base case: Jika n=0, program mengembalikan nilai 1, karena 2^0=1.
- Recursive case: jika n>0, program memanggil fungsi `pangkat()` dengan n−1, dan mengalikan hasilnya dengan 2.

4. Hasil dari 2^n dicetak sebagai output.
#### Cara kerja program :
1. Input Pengguna: Program meminta pengguna memasukkan angka, seperti 3.

2. Hitung 2^n: Program menggunakan fungsi `pangkat(n)` untuk menghitung 2^n

3. Logika Fungsi:
- Jika n adalah 0, fungsi mengembalikan 1.
- Jika n lebih besar dari 0, fungsi memanggil dirinya sendiri dengan n−1 dan mengalikan hasilnya dengan 2.

4. Hasil: Setelah menghitung, program menampilkan hasilnya.
### 4. Mencari nilai faktorial atau n!
Base-case: n == 0 atau n == 1
```go
package main

import "fmt"

var n int
func faktorial(n int) int{
	if n == 0 || n == 1{
		return 1
	}else{
		return n * faktorial(n-1)
	}
}

func main(){
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}
```
#### Deskripsi program
Program ini menghitung faktorial dari sebuah bilangan bulat non-negatif n menggunakan fungsi rekursif. Faktorial dari n, dilambangkan dengan n!, adalah hasil kali semua bilangan bulat positif dari 1 hingga n. Program meminta input dari pengguna dan menampilkan hasil faktorial dari angka tersebut.
#### Algoritma program :
1. Program meminta pengguna untuk memasukkan bilangan bulat non-negatif n.

2. Fungsi `faktorial()` dipanggil dengan parameter n untuk menghitung hasil faktorial.

3. Di dalam fungsi `faktorial()`, terdapat dua kasus:
- Base case: Jika n=0 atau n=1,fungsi akan mengembalikan nilai 1.
- Recursive case: Jika n>1, fungsi akan memanggil dirinya sendiri dengan parameter n−1 dan mengalikan hasilnya dengan n.

4. Hasil dari faktorial ditampilkan sebagai output.
#### Cara kerja program :
1. Input Pengguna: Program meminta pengguna untuk memasukkan angka, misalnya 5.

2. Hitung Faktorial: Program memanggil fungsi faktorial(n) untuk menghitung faktorial dari angka yang dimasukkan.

3. Logika Fungsi:
- Jika n adalah 0 atau 1, fungsi mengembalikan 1.
- Jika n lebih besar dari 1, fungsi memanggil dirinya sendiri dengan n−1 dan mengalikan hasilnya dengan n.

4. Hasil: Setelah semua pemanggilan selesai, program menampilkan hasil faktorial.

## Unguided
### 1. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci 
Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilal suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S₁ =Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10.
![image](https://github.com/user-attachments/assets/4263acff-dfcb-46bd-b7dc-a80c25e7bbb2)
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.

![image](https://github.com/user-attachments/assets/6f75059e-002d-4b32-9900-589175b9f579)
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu
Masukan terdiri dari sebuah bilangan positif N

Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N
![image](https://github.com/user-attachments/assets/a088d67a-43d7-4002-8aa1-3ec9d0bf7763)
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.
![image](https://github.com/user-attachments/assets/c47df5f9-c85e-4bed-94cc-8645ac695dcb)
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
Masukan terdiri dari bilangan bulat x dan y.

Keluaran terdiri dari hasil x dipangkatkan y.

Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".
#### Source code :
```go

```
#### Deskripsi program
#### Algoritma program :
#### Cara kerja program :
