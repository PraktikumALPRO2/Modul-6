<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL 6 </strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Haifa Zahra Azzimmi /2311102163<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
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
**1. Definisi Rekursif**

Rekursif adalah metode pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk memecahkan masalah yang lebih kecil dari masalah awal hingga mencapai kondisi dasar. Ini digunakan untuk menyelesaikan masalah yang dapat dibagi menjadi sub-masalah yang lebih sederhana dan mirip.

**2. Komponen Rekursif**

Komponen rekursif biasanya mencakup hal-hal berikut:

1. Basis Kasus (Base Case): Kondisi yang akan menghentikan rekursi. Misalnya, dalam perhitungan faktorial, basis kasusnya adalah ketika n sama dengan 0 atau 1.

2. Panggilan Rekursif (Recursive Call): Fungsi memanggil dirinya sendiri dengan argumen yang dimodifikasi untuk mendekati basis kasus. Ini adalah bagian di mana masalah dipecah menjadi sub-masalah yang lebih kecil.

3. Kemajuan Menuju Basis Kasus (Progress Towards Base Case): Setiap panggilan rekursif harus mengubah argumen sedemikian rupa sehingga mendekati basis kasus, mencegah rekursi tak terbatas.


## Guided
### 1. Membuat baris bilangan dari n hingga 1. 
Base-case: bilangan == 1
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
#### Output Program:
![image](https://github.com/user-attachments/assets/7308d33e-d64d-407e-b992-9fe9c7d6fade)

### Deskripsi Program:
Program ini meminta pengguna memasukkan sebuah angka, kemudian menggunakan fungsi rekursif untuk mencetak angka dari angka tersebut hingga angka 1 secara menurun

### Algoritma Program:
1. Mulai
2. Minta pengguna memasukkan sebuah angka n.
3. Baca input n.
4. Panggil prosedur baris dengan parameter n.
5. Dalam prosedur baris:
   - Jika n sama dengan 1, cetak 1.
   - Jika n lebih besar dari 1, cetak nilai n dan panggil kembali prosedur baris dengan parameter n-1.
6. Selesai
   
### Cara Kerja Program:
1. Program dimulai dan meminta pengguna memasukkan sebuah angka n.
2. Input n dibaca dan disimpan dalam variabel n.
3. Program memanggil fungsi baris dengan argumen n.
4. Fungsi baris bekerja secara rekursif dengan langkah-langkah berikut:
   - Jika bilangan adalah 1, maka angka 1 dicetak dan fungsi berhenti.
   - Jika bilangan lebih besar dari 1, nilai bilangan dicetak dan fungsi memanggil dirinya sendiri dengan bilangan - 1.
5. Proses rekursif ini terus berlanjut hingga nilai bilangan mencapai 1.
6. Program selesai setelah mencetak semua angka dari n hingga 1.

### 2. Menghitung hasil penjumlahan 1 hingga n.
Base-case: n == 1
```go
package main
import "fmt"

func penjumlahan(n int) int {
	if n == 1{
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

```

#### Output Program
![image](https://github.com/user-attachments/assets/bfb1119c-6785-4342-8116-ec4c62f9e418)

### Deskripsi Program:
Program ini meminta pengguna memasukkan angka n, kemudian menggunakan fungsi rekursif penjumlahan untuk menghitung dan mencetak penjumlahan semua angka dari 1 hingga n.

### Algoritma Program:
1. Mulai
2. Minta pengguna memasukkan angka n.
3. Baca input n.
4. Panggil fungsi penjumlahan dengan parameter n.
5. Dalam fungsi penjumlahan:
   - Jika n sama dengan 1, kembalikan nilai 1.
   - Jika n lebih besar dari 1, kembalikan nilai n ditambah hasil dari penjumlahan(n-1).
6. Cetak hasil penjumlahan dari 1 hingga n.
7. Selesai

### Cara Kerja Program:
1. Meminta pengguna memasukkan sebuah angka n.
2. Input n dibaca dan disimpan.
3. Program memanggil fungsi penjumlahan dengan n sebagai argumen.
4. Fungsi penjumlahan bekerja secara rekursif:
   - Jika n adalah 1, kembalikan nilai 1.
   - Jika n lebih besar dari 1, kembalikan nilai n ditambah hasil dari penjumlahan(n-1).
5. Hasil penjumlahan dari 1 hingga n dicetak ke layar.

### 3. Mencari dua pangkat n atau 2^n. 
Base-case: n == 0
```go
package main
import "fmt"

//Fungsi rekursif untuk menghitung 2^n
func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
func main() {
	var n int
	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah:", pangkat(n))
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/7e406702-92a2-4884-b1c3-b4ea704a16b1)

### Deskripsi Program:
Program ini meminta pengguna memasukkan nilai n, kemudian menggunakan fungsi rekursif pangkat untuk menghitung hasil dari 2 pangkat n. Hasilnya kemudian dicetak ke layar. Simpel dan efisien.

### Algoritma Program:
1. Mulai
2. Minta pengguna memasukkan nilai n.
3. Baca input n.
4. Panggil fungsi pangkat dengan parameter n.
5. Dalam fungsi pangkat:
   - Jika n sama dengan 0, kembalikan nilai 1.
   - Jika n lebih besar dari 0, kembalikan 2 dikali hasil dari pangkat(n-1).
6. Cetak hasil dari 2 pangkat n.
7. Selesai

### Cara Kerja Program:
1. Meminta pengguna memasukkan nilai n.
2. Input n dibaca dan disimpan.
3. Program memanggil fungsi rekursif pangkat dengan n sebagai argumen.
4. Dalam fungsi pangkat:
   - Jika n sama dengan 0, kembalikan nilai 1 (basis kasus).
   - Jika n lebih besar dari 0, kembalikan nilai 2 dikali hasil dari pemanggilan fungsi pangkat(n-1) (rekursi).
5. Hasil dari perhitungan 2 pangkat n dicetak ke layar.

### 4. Mencari nilai faktorial atau n!. 
Base-case: n == 0 atau n == 1 
```go
package main
import "fmt"

var n int
func faktorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}
func main () {

	fmt. Scan(&n)
	fmt.Println(faktorial(n))
}
```

#### Output Program:
![image](https://github.com/user-attachments/assets/5a318a81-d4bd-40f4-a47f-e544df9a3af6)

### Deskripsi Program:
Program ini meminta pengguna memasukkan angka n, lalu menggunakan fungsi rekursif faktorial untuk menghitung nilai faktorial dari n dan mencetak hasilnya.

### Algoritma Program:
1. Mulai
2. Minta pengguna memasukkan angka n.
3. Baca input n.
4. Panggil fungsi faktorial dengan parameter n.
5. Dalam fungsi faktorial:
   - Jika n sama dengan 0 atau 1, kembalikan nilai 1.
   - Jika n lebih besar dari 1, kembalikan nilai n dikali hasil dari faktorial(n-1).
6. Cetak hasil faktorial dari n.
7. Selesai
   
### Cara Kerja Program:
1. Program dimulai dengan meminta pengguna memasukkan angka n.
2. Input n dibaca dan disimpan.
3. Program memanggil fungsi rekursif faktorial dengan n sebagai argumen.
4. Dalam fungsi faktorial:
   - Jika n adalah 0 atau 1, kembalikan nilai 1 sebagai hasil faktorial.
   - Jika n lebih besar dari 1, fungsi memanggil dirinya sendiri dengan argumen n-1 dan mengalikan hasilnya dengan n.
5. Hasil faktorial dari n dicetak ke layar.


## Unguided 

### 1. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci 
**Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilal suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S₁ =Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10.**

![image](https://github.com/user-attachments/assets/67a77beb-da3c-4ee5-92b0-6ccd4009ce67)

### Source Code :
```go

```

### Output:

### Full code Screenshot:


### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program:


### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user

### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program:


   
### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.**

![image](https://github.com/user-attachments/assets/28a4e99e-3018-4ea4-9f3d-8812a8eeffa8)

### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:


### 4.

![image](https://github.com/user-attachments/assets/873da708-8303-4f19-8db6-b3d70921da7b)

### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 


### Algoritma Program

### Cara Kerja Program:


### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.**

### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

   
### Cara Kerja Program:


### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
**Masukan terdiri dari bilangan bulat x dan y.**

**Keluaran terdiri dari hasil x dipangkatkan y.**

**Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".**

### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program:


