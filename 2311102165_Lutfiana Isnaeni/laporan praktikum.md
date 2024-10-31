<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VI </strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Lutfiana Isnaeni L /2311102165<br>
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






## Guided
### 1. Membuat baris bilangan dari n hingga 1. Base-case: bilangan == 1
```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)

}

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/836e0f29-7159-47bc-92a2-b8c841a1b17b)
### Deskripsi Program
Program ini menerima input berupa angka bulat 𝑛 dan menampilkan urutan angka dari 
𝑛 hingga 1 dalam susunan menurun. Program menggunakan teknik rekursif untuk mencetak angka-angka tersebut.
### Algoritma Program:
1.	Input angka bulat n.
2.	Panggil fungsi baris dengan parameter n.
3.	Di dalam fungsi `baris`:
   
    •	Jika nilai parameter adalah 1, cetak angka 1.
  	
    •	Jika tidak, cetak nilai parameter tersebut, lalu panggil kembali fungsi `baris` dengan mengurangi nilainya sebesar 1.
4.	Fungsi ini terus berjalan secara rekursif hingga parameter mencapai nilai 1, setelah itu proses rekursi berakhir.

### Cara Kerja Program:
1.	Program dimulai dengan menerima angka dari pengguna dan menyimpannya dalam variabel `n`.
2.	Fungsi `baris(n)` dipanggil menggunakan nilai `n` sebagai parameter.
3.	Fungsi `baris` bekerja sebagai berikut:
   
    •	Setiap kali fungsi `baris` dipanggil, ia mencetak nilai parameter jika nilainya lebih besar dari 1.

    •	Fungsi `baris` kemudian dipanggil kembali dengan parameter dikurangi 1, menghasilkan proses rekursif.

    •	Ketika nilai parameter mencapai 1, fungsi mencetak angka 1 dan menghentikan rekursi.

4.	Akhirnya, program menghasilkan deret angka menurun mulai dari `n` hingga 1.

### 2. Menghitung hasil penjumlahan 1 hingga n. Base-case: n == 1
```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}
```

#### Output Program
![image](https://github.com/user-attachments/assets/630cf12f-3af6-4ff8-a4b5-18586210074b)
### Deskripsi Program
Program ini menerima input berupa angka bulat n dan menghitung jumlah dari seluruh angka mulai dari 1 hingga n menggunakan teknik rekursif. Hasil penjumlahan kemudian ditampilkan sebagai output.
### Algoritma Program:
1.	Ambil angka n dari input pengguna.
2.	Panggil fungsi `penjumlahan` dengan parameter n.
3.	Di dalam fungsi `penjumlahan`, lakukan pengecekan:
   
    •	Jika nilai n sama dengan 1, kembalikan nilai 1.
  	
    •	Jika tidak, tambahkan n dengan hasil dari `penjumlahan(n-1)`.
  	
4.	Fungsi ini akan terus berjalan secara rekursif hingga nilai parameter mencapai 1, dan pada titik ini seluruh nilai penjumlahan akan diselesaikan.

### Cara Kerja Program:
1.	Program mulai dengan membaca input pengguna dan menyimpannya di variabel `n`.
2.	Fungsi `penjumlahan(n)` kemudian dipanggil, dan hasilnya dicetak sebagai output.
3.	Fungsi `penjumlahan` bekerja sebagai berikut:
   
    •	Jika dipanggil dengan nilai n>1, fungsi akan mengembalikan nilai n ditambah hasil dari `penjumlahan(n-1)`.
  	
    •	Ketika n bernilai 1, fungsi akan mengembalikan nilai 1 sebagai dasar perhitungan rekursif ini.
  	
    •	Rekursi kemudian menyelesaikan perhitungan jumlah dari 1 hingga n.
  	
4.	Program akhirnya mencetak hasil penjumlahan dari seluruh angka 1 hingga n.

### 3. Mencari dua pangkat n atau 2^n. Base-case: n == 0
```go
package main

import "fmt"

//Fungsi rekrusif untuk menghitung 2^n
func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah: ", pangkat(n))
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/9231022b-a03a-4739-8e5b-8e6de7ea624a)
### Deskripsi Program
Program ini menghitung nilai 2^n(dua pangkat n) dengan menggunakan fungsi rekursif. Pengguna memasukkan nilai n, dan program akan menampilkan hasil pangkat tersebut di layar.
### Algoritma Program:
1.	Terima masukan integer n dari pengguna.
2.	Panggil fungsi `pangkat` dengan parameter nnn.
3.	Di dalam fungsi `pangkat`:
   
    •	Jika n sama dengan 0, kembalikan nilai 1, karena 2^0 = 1
  	
    •	Jika n lebih besar dari 0, kembalikan hasil perkalian 2 dengan hasil pemanggilan fungsi `pangkat(n-1)`.
  	
4.	Proses rekursif akan terus berjalan hingga n mencapai 0, kemudian semua hasil perkalian akan dihitung.
5.	Tampilkan hasil perhitungan pangkat yang diperoleh.

### Cara Kerja Program:
1.	Program mulai dengan meminta pengguna memasukkan nilai untuk variabel n.
2.	Fungsi pangkat(n) dipanggil dengan nilai n sebagai argumen, dan hasilnya ditampilkan sebagai output.
3.	Fungsi pangkat bekerja sebagai berikut:
   
    •	Jika n = 0, fungsi mengembalikan 1 sebagai dasar perhitungan rekursif (karena 2^0 = 1)
  	
    •	Jika n > 0, fungsi mengembalikan hasil dari 2 dikalikan dengan hasil pangkat(n-1), sehingga melanjutkan proses rekursif.
  	
    •	Rekursi berakhir saat nnn bernilai 0, dan seluruh hasil perkalian dikembalikan ke fungsi utama.
4.	Akhirnya, program menampilkan hasil perhitungan 2^n.


### 4. Mencari nilai faktorial atau n!. Base-case: n == 0 atau n == 1 
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

#### Output Program
![image](https://github.com/user-attachments/assets/3e108bbb-99ed-4386-8d78-e7db7140ccbd)
### Deskripsi Program
Program ini menghitung nilai faktorial dari angka bulat n yang diinput oleh pengguna. Faktorial adalah hasil kali dari semua bilangan bulat positif hingga n. Program ini menggunakan pendekatan rekursif untuk melakukan perhitungan faktorial.
### Algoritma Program:
1.	Ambil input bilangan bulat nnn dari pengguna.
2.	Panggil fungsi `faktorial` dengan parameter n.
3.	Di dalam fungsi `faktorial`:
   
    •	Jika n sama dengan 0 atau 1, kembalikan 1 (karena 0! = 1  dan 1! = 1)
  	
    •	Jika tidak, kembalikan hasil perkalian n dengan hasil pemanggilan fungsi `factorial (n-1)`.
  	
4.	Proses rekursi akan terus berlanjut hingga n mencapai 0 atau 1, dan hasil akhir dari faktorial akan dikembalikan ke pemanggil fungsi.

### Cara Kerja Program:
1.	Program mulai dengan membaca input dari pengguna dan menyimpan nilai tersebut dalam variabel n.
2.	Fungsi `faktorial(n)` kemudian dipanggil dengan n sebagai argumen, dan hasilnya dicetak sebagai output.
3.	Fungsi `faktorial` bekerja dengan langkah -langkah berikut:
    
    •	Jika n adalah 0 atau 1, fungsi akan mengembalikan 1.
  	
    •	Jika n lebih besar dari 1, fungsi mengembalikan hasil n dikalikan dengan pemanggilan `faktorial(n-1)`, sehingga melanjutkan proses rekursif.
  	
    •	Proses ini terus berlanjut sampai mencapai kondisi dasar, yaitu saat n menjadi 0 atau 1.
  	
4.	Kemudian, program menampilkan hasil perhitungan faktorial dari nilai n.


## Unguided
### 1.Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci 
Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilal suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S₁ =Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. 
![image](https://github.com/user-attachments/assets/096ae557-95b8-46d4-8dd4-a0f8e95dcde2)

```go
#### Output Program

### Deskripsi Program


### Algoritma Program :

### Cara Kerja Program : 

### 2. 

```go

#### Output Program

### Deskripsi Program

### Algoritma Program :

### Cara Kerja Program : 

### 3. 

```go

#### Output Program

### Deskripsi Program

### Algoritma Program :

### Cara Kerja Program : 

### 4. 

```go

#### Output Program

### Deskripsi Program

### Algoritma Program :

### Cara Kerja Program : 

### 5. 

go```

#### Output Program

### Deskripsi Program

### Algoritma Program :

### Cara Kerja Program : 

### 6. 

go```

#### Output Program

### Deskripsi Program

### Algoritma Program :

### Cara Kerja Program : 


