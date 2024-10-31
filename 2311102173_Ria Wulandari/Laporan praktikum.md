

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
#### Deskripsi
