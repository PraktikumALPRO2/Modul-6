

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
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scanln(&n)

	// Menampilkan baris pertama untuk nilai n
	fmt.Print("n  ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Menampilkan baris kedua untuk nilai S_n
	fmt.Print("S_n ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

```
#### Output
![image](https://github.com/user-attachments/assets/e97022c0-bbbb-464e-8b7a-6e3cabccd302)
#### Deskripsi program
Program ini dibuat untuk menghitung deret Fibonacci hingga nilai ke-n yang dimasukkan oleh pengguna. Program menampilkan dua baris:

1. Baris pertama menampilkan indeks dari 0 hingga n.

2. Baris kedua menampilkan nilai Fibonacci yang sesuai dengan setiap indeks tersebut.

Fungsi rekursif `fibonacci` digunakan untuk menghitung nilai Fibonacci pada setiap indeks.
#### Algoritma program :
1. Input: Pengguna diminta memasukkan sebuah bilangan bulat `n` yang akan menentukan jumlah suku Fibonacci yang ingin dihitung.

2. Proses:
- Program memiliki fungsi `fibonacci(n)` yang menghitung nilai Fibonacci ke-n secara rekursif.

- Jika nilai `n` adalah 0 atau 1, fungsi akan langsung mengembalikan nilai tersebut sebagai hasil.

- Jika nilai `n` lebih besar dari 1, fungsi akan memanggil dirinya sendiri dengan nilai `n-1` dan `n-2`, kemudian menjumlahkan hasilnya.

3. Output: Program mencetak dua baris:

- Baris pertama berisi indeks dari 0 hingga n.

- Baris kedua berisi nilai Fibonacci yang sesuai dengan indeks masing-masing.
#### Cara kerja program :
1. Program meminta pengguna untuk memasukkan sebuah angka `n` (misalnya 5).

2. Program mencetak baris pertama yang berisi indeks dari 0 sampai 5.

3. Program kemudian memanggil fungsi `fibonacci` untuk setiap indeks dari 0 sampai 5 untuk menghitung nilai Fibonacci.

4. Hasil perhitungan Fibonacci ditampilkan di baris kedua, sesuai dengan indeksnya.

5. Program selesai setelah menampilkan hasil dalam bentuk tabel.
### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user
#### Source code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan sejumlah bintang pada satu baris
func tampilkanBintang(jumlahBintang int) {
	// Jika jumlah bintang mencapai 0, hentikan rekursi
	if jumlahBintang == 0 {
		return
	}
	// Panggil rekursi untuk menurunkan jumlah bintang
	tampilkanBintang(jumlahBintang - 1)
	// Cetak satu bintang
	fmt.Print("*")
}

// Fungsi rekursif untuk mencetak pola bintang dari baris pertama hingga baris terakhir
func cetakPolaBintang(totalBaris int, barisSaatIni int) {
	// Jika baris saat ini melewati total baris, hentikan rekursi
	if barisSaatIni > totalBaris {
		return
	}
	// Tampilkan bintang sesuai dengan nomor baris saat ini
	tampilkanBintang(barisSaatIni)
	// Pindah ke baris baru
	fmt.Println()
	// Panggil fungsi ini lagi untuk baris berikutnya
	cetakPolaBintang(totalBaris, barisSaatIni+1)
}

func main() {
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris (N): ")
	fmt.Scan(&jumlahBaris)

	fmt.Println("Pola bintang:")
	cetakPolaBintang(jumlahBaris, 1)
}

```
#### Output
![image](https://github.com/user-attachments/assets/cdb42ae8-19a8-41f2-bbb3-e0ca0753224d)

#### Deskripsi program
Program ini adalah aplikasi sederhana untuk mencetak pola segitiga bintang menggunakan fungsi rekursif dalam bahasa Go. Program meminta pengguna memasukkan jumlah baris yang diinginkan, dan kemudian menampilkan pola segitiga yang dimulai dari satu bintang pada baris pertama dan bertambah satu bintang pada setiap baris berikutnya hingga mencapai jumlah baris yang diinputkan.
#### Algoritma program :
1. Input: Program meminta jumlah baris dari pengguna.

2. Cetak Pola dengan Rekursi:

- Fungsi `cetakPolaBintang` mencetak bintang di tiap baris dan memanggil dirinya sendiri untuk baris berikutnya.

- Fungsi `tampilkanBintang` mencetak bintang dalam jumlah sesuai nomor baris, menggunakan rekursi hingga selesai.
#### Cara kerja program :
1. Pengguna memasukkan jumlah baris.

2. Program memulai cetak dari baris pertama dan berlanjut hingga mencapai baris terakhir yang diminta.

3. Setiap baris menampilkan jumlah bintang sesuai nomor baris, hingga terbentuk segitiga.
### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.

![image](https://github.com/user-attachments/assets/6f75059e-002d-4b32-9900-589175b9f579)
#### Source code :
```go
package main

import "fmt"

// Fungsi rekursif untuk mencari faktor dari suatu bilangan
func tampilkanFaktor(bilangan, indeks int) {
	if indeks > bilangan {
		return
	}
	if bilangan%indeks == 0 {
		fmt.Print(indeks, " ")
	}
	tampilkanFaktor(bilangan, indeks+1)
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	fmt.Print("Faktor dari ", angka, ": ")
	tampilkanFaktor(angka, 1)
	fmt.Println()
}

```
#### Output
![image](https://github.com/user-attachments/assets/0322aa4a-4a32-4819-b9f1-47ed9cb7d224)

#### Deskripsi program
Program ini dibuat untuk menampilkan faktor-faktor dari sebuah bilangan bulat positif yang dimasukkan oleh pengguna. Faktor-faktor ini adalah bilangan-bilangan yang dapat membagi bilangan input secara tepat (sisa pembagian sama dengan nol).
#### Algoritma program :
1. Input Bilangan: Pengguna memasukkan sebuah bilangan bulat positif.

2. Menampilkan Faktor: Program menggunakan fungsi rekursif `tampilkanFaktor` untuk menemukan dan mencetak setiap faktor dari bilangan tersebut, mulai dari angka 1 hingga bilangan itu sendiri.

3. Pencetakan Hasil: Program mencetak semua faktor bilangan secara terurut dan kemudian menyelesaikan eksekusinya.
#### Cara kerja program :
1. Program dimulai dengan meminta pengguna untuk memasukkan sebuah bilangan, yang disimpan dalam variabel `angka`.

2. Fungsi `tampilkanFaktor` dipanggil dengan `angka` dan dimulai dari `indeks` = 1.

3. Di dalam `tampilkanFaktor`:

- Jika `indeks` lebih besar dari `bilangan`, fungsi berhenti (`return`).

- Jika `indeks` adalah faktor dari `bilangan` (yaitu `bilangan % indeks == 0`), maka nilai `indeks` dicetak sebagai faktor.

- Fungsi ini kemudian memanggil dirinya sendiri dengan menambah nilai `indeks` sebesar 1, sehingga proses rekursif mencari faktor berikutnya.

4. Proses ini berlanjut hingga seluruh faktor dari `bilangan` ditemukan dan dicetak.
### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu
Masukan terdiri dari sebuah bilangan positif N

Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N
![image](https://github.com/user-attachments/assets/a088d67a-43d7-4002-8aa1-3ec9d0bf7763)
#### Source code :
```go
package main

import "fmt"

// Fungsi rekursif untuk menampilkan deret bilangan dari M ke 1 dan kembali ke M
func tampilkanDeret(bilangan, saatIni int) {
	// Kondisi dasar rekursi
	if saatIni == 1 {
		fmt.Print(saatIni, " ")
		return
	}

	// Tampilkan nilai saat ini dan panggil fungsi secara rekursif ke bawah
	fmt.Print(saatIni, " ")
	tampilkanDeret(bilangan, saatIni-1)

	// Tampilkan nilai saat kembali dari rekursi
	fmt.Print(saatIni, " ")
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	fmt.Print("Hasil: ")
	tampilkanDeret(angka, angka)
	fmt.Println()
}

```
#### Output
![image](https://github.com/user-attachments/assets/c64cfe40-3bda-4acb-b250-367c1c7fe2de)

#### Deskripsi program
Program ini bertujuan untuk menampilkan deret bilangan dari `N` hingga `1`, kemudian kembali lagi dari `1` ke `N`. Pengguna memasukkan bilangan `N`, dan program mencetak pola deret tersebut menggunakan fungsi rekursif.
#### Algoritma program :
1. Input Bilangan: Pengguna memasukkan sebuah bilangan` N`.

2. Pencetakan Pola Rekursif:

- Fungsi `tampilkanDeret` dipanggil dengan `N` sebagai parameter awal.

- Fungsi ini mencetak angka dari `N` ke `1` dan kembali lagi ke `N` melalui proses rekursi.

3. Basis Rekursi: Jika nilai `saatIni` sama dengan `1`, fungsi mencetak `1` dan menghentikan pemanggilan rekursif lebih lanjut.

4. Rekursi Balik Naik: Setelah mencapai `1`, fungsi kembali ke setiap pemanggilan sebelumnya dan mencetak angka pada tiap langkah.
#### Cara kerja program :
1. Program meminta pengguna memasukkan bilangan `N`.

2. Fungsi `tampilkanDeret` dipanggil, dimulai dengan `N`.

3. Fungsi mencetak nilai saat ini (`saatIni`) dan memanggil dirinya sendiri dengan `saatIni - 1`, sampai nilai saatIni menjadi `1`.

4. Saat saatIni mencapai `1`, angka `1` dicetak, dan fungsi berhenti memanggil dirinya sendiri lebih jauh.

5. Program kembali mencetak setiap nilai `saatIni` yang tersimpan dari rekursi sebelumnya, membentuk deret yang naik dari `1` kembali ke `N`.
### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.
![image](https://github.com/user-attachments/assets/c47df5f9-c85e-4bed-94cc-8645ac695dcb)
#### Source code :
```go
package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga batas tertentu
func tampilkanGanjil(batas, angkaSaatIni int) {
	if angkaSaatIni > batas {
		return
	}
	if angkaSaatIni%2 != 0 {
		fmt.Print(angkaSaatIni, " ")
	}
	tampilkanGanjil(batas, angkaSaatIni+1)
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)
	tampilkanGanjil(angka, 1)
	fmt.Println()
}

```
#### Output 
![image](https://github.com/user-attachments/assets/c382cd11-9405-4279-b473-c96145c458fa)

#### Deskripsi program
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan berfungsi untuk mencetak semua bilangan ganjil dari 1 hingga angka yang dimasukkan oleh pengguna. Program ini menggunakan pendekatan rekursif untuk menampilkan deret bilangan ganjil.
#### Algoritma program :
1. Deklarasi Fungsi: Buat fungsi rekursif `tampilkanGanjil(batas, angkaSaatIni)` untuk menampilkan bilangan ganjil.

- Input: `batas` (angka maksimum) dan `angkaSaatIni` (angka yang sedang diproses).

- Basis Rekursi: Jika `angkaSaatIni` lebih besar dari `batas`, hentikan rekursi.

- Cek Bilangan Ganjil: Jika `angkaSaatIni` adalah bilangan ganjil (modulus 2 tidak sama dengan 0), cetak `angkaSaatIni`.

- Panggil Rekursi: Panggil `tampilkanGanjil` dengan `angkaSaatIni` yang ditingkatkan 1.

2. Fungsi `main`:

- Minta pengguna untuk memasukkan angka dan simpan dalam variabel `angka`.

- Panggil fungsi `tampilkanGanjil` dengan `angka` sebagai batas dan 1 sebagai angka saat ini.

- Cetak baris baru setelah semua bilangan ganjil ditampilkan.
#### Cara kerja program :
1. Program dimulai dari fungsi `main`.

2. Pengguna diminta untuk memasukkan sebuah bilangan yang akan menjadi batas untuk bilangan ganjil.

3. Fungsi `tampilkanGanjil` dipanggil dengan `angka` sebagai batas dan 1 sebagai angka awal.

4. Fungsi ini memeriksa apakah `angkaSaatIni` lebih besar dari `batas`. Jika ya, proses berhenti.

5. Jika `angkaSaatIni` adalah bilangan ganjil, maka angka tersebut akan dicetak.

6. Fungsi kemudian memanggil dirinya sendiri dengan `angkaSaatIni` yang bertambah 1.

7. Proses berulang hingga semua bilangan ganjil dari 1 hingga batas ditampilkan.
### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
Masukan terdiri dari bilangan bulat x dan y.

Keluaran terdiri dari hasil x dipangkatkan y.

Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".
#### Source code :
```go
package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai a pangkat b
func hitungPangkat(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * hitungPangkat(a, b-1)
}

func main() {
	var dasar, pangkat int
	fmt.Print("Masukkan dasar (a): ")
	fmt.Scan(&dasar)
	fmt.Print("Masukkan pangkat (b): ")
	fmt.Scan(&pangkat)
	hasil := hitungPangkat(dasar, pangkat)
	fmt.Printf("%d pangkat %d = %d\n", dasar, pangkat, hasil)
}

```
#### Output
![image](https://github.com/user-attachments/assets/74931e7b-6fc5-4c21-8ae8-edd0e91dea9c)

#### Deskripsi program
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan bertujuan untuk menghitung nilai pangkat dari suatu bilangan. Pengguna akan diminta untuk memasukkan dua angka: yang pertama adalah bilangan dasar (a) yang ingin dipangkatkan, dan yang kedua adalah pangkat (b) yang menunjukkan seberapa banyak bilangan dasar tersebut akan dikalikan dengan dirinya sendiri. Penghitungan pangkat dilakukan dengan metode rekursif, yaitu teknik di mana fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih kecil hingga mencapai kondisi dasar yang menghentikan rekursi.
#### Algoritma program :
1. Mulai.

2. Deklarasi fungsi `hitungPangkat(a, b int)` int:

- Jika `b` sama dengan 0, kembalikan 1 (kondisi dasar).

- Jika tidak, kembalikan `a` dikali hasil dari `hitungPangkat(a, b-1)`.

3. Dalam fungsi `main()`:

- Deklarasi dua variabel `dasar` dan `pangkat`.

- Minta pengguna memasukkan nilai untuk `dasar` dan `pangkat`.

- Panggil fungsi `hitungPangkat(dasar, pangkat)` dan simpan hasilnya di variabel `hasil`.

- Cetak hasil dalam format yang jelas.

4. Selesai.
#### Cara kerja program :
1. Program dimulai dengan mendeklarasikan fungsi `hitungPangkat` yang menerima dua parameter: `a` (dasar) dan `b` (pangkat).

2. Dalam fungsi tersebut, terdapat kondisi dasar di mana jika `b` sama dengan 0, fungsi akan mengembalikan 1. Ini mengikuti prinsip matematis bahwa setiap bilangan pangkat 0 adalah 1.

3. Jika `b` lebih besar dari 0, fungsi akan mengalikan `a` dengan hasil dari panggilan rekursif `hitungPangkat(a, b-1)`. Ini berarti bahwa program akan terus memanggil dirinya sendiri dengan nilai pangkat yang berkurang satu hingga mencapai kondisi dasar.

4. Dalam fungsi `main`, program meminta pengguna untuk memasukkan nilai dasar dan pangkat melalui input dari konsol.

5. Setelah kedua nilai dimasukkan, fungsi `hitungPangkat` dipanggil dengan nilai yang diberikan, dan hasilnya disimpan dalam variabel `hasil`.

6. Program kemudian mencetak hasil penghitungan dalam format yang menyatakan nilai pangkat yang telah dihitung, sehingga pengguna dapat dengan mudah memahami hasilnya.

7. Program berakhir setelah menampilkan hasil.
