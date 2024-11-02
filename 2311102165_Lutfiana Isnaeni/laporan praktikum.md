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

### Source Code: 
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

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
#### Output Program
![image](https://github.com/user-attachments/assets/4296f075-caa2-4d75-89d9-7e19fbd960c6)

### Deskripsi Program
Program ini menampilkan deret Fibonacci dalam bentuk tabel hingga jumlah suku yang ditentukan oleh pengguna. Deret Fibonacci merupakan urutan angka di mana setiap angkanya adalah hasil penjumlahan dari dua angka sebelumnya, dimulai dengan angka 0 dan 1. 
### Algoritma Program :
1.	Definisikan fungsi `fibonacci` untuk menghitung nilai Fibonacci dari suatu suku `n` secara rekursif:
   
	•	Jika `n` adalah 0 atau 1, fungsi mengembalikan nilai `n` (karena dua suku pertama Fibonacci adalah 0 dan 1).

	•	Jika n lebih dari 1, fungsi memanggil dirinya sendiri untuk menghitung `fibonacci(n-1) + fibonacci(n-2)`.

2.	Dalam fungsi `main()`:

	•	Program meminta pengguna untuk memasukkan jumlah suku Fibonacci `(n)` yang akan dihitung.

	•	Cetak baris pertama tabel dengan label `n`, diikuti oleh angka-angka urutan dari 0 hingga `n`.

	•	Cetak baris kedua tabel dengan label `S_n`, diikuti oleh nilai-nilai Fibonacci untuk setiap suku dari 0 hingga `n`.

3.	Hasilnya ditampilkan dalam bentuk tabel.


### Cara Kerja Program : 
1.	Program dimulai dengan fungsi `main()`, di mana pengguna memasukkan jumlah suku Fibonacci yang akan ditampilkan.
   
2.	Setelah menerima input `n`, program membuat dua baris untuk tabel:
   
	•	Baris `n`: berisi nomor urutan setiap suku dari 0 hingga `n`.

	•	Baris `S_n`: berisi hasil perhitungan deret `Fibonacci` dari suku ke-0 hingga ke-`n`, menggunakan fungsi fibonacci.

3.	Fungsi `fibonacci(n)` dipanggil secara rekursif untuk setiap nilai `i` dari 0 hingga `n`:
   
	•	Jika `i` bernilai 0 atau 1, fungsi mengembalikan nilai `i`.

	•	Jika `i` lebih besar dari 1, fungsi memanggil `fibonacci(i-1)` dan `fibonacci(i-2)`, lalu menjumlahkan hasilnya.

4.	Program kemudian menampilkan deret Fibonacci dalam format tabel, dengan `n` untuk indeks dan `S_n` untuk nilai deret Fibonacci.


### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user

![image](https://github.com/user-attachments/assets/1af3fed6-0e4b-437f-be94-e7bc4557930c)

### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menampilkan pola bintang
func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var input int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&input)
	printStars(input)
}
```


#### Output Program
![image](https://github.com/user-attachments/assets/87c25bca-d297-4436-a67e-66769ab2176a)
### Deskripsi Program
Program di atas berfungsi untuk mencetak pola bintang berbentuk segitiga menurun menggunakan rekursi. Program meminta pengguna untuk memasukkan jumlah baris yang diinginkan, kemudian menampilkan pola bintang dengan jumlah bintang yang semakin berkurang pada tiap baris berikutnya.

### Algoritma Program :

1.	Fungsi `main()` meminta input dari pengguna, yaitu jumlah baris `(n)`.
   
2.	Nilai `n` ini akan diproses oleh fungsi rekursif `printStars()`, yang berperan untuk mencetak pola bintang.
   
3.	Di dalam fungsi `printStars()`:
    
	•	Fungsi akan memanggil dirinya sendiri dengan mengurangi nilai `n` (yaitu `n-1`) hingga mencapai nilai 0 (sebagai kondisi akhir rekursi).

	•	Setelah mencapai nilai 0, fungsi berhenti memanggil dirinya sendiri dan mencetak n bintang `(*)` pada setiap tingkat rekursi dalam proses kembali 		ke tingkat awal.

4.	Ketika fungsi kembali ke tingkat awal, setiap kali nilai `n` tercetak dalam bentuk baris berisi `n` bintang.

5.	Proses ini terus berlanjut hingga seluruh pola bintang tercetak.

### Cara Kerja Program : 
1.	Program dimulai dari fungsi `main()`, di mana pengguna diminta memasukkan jumlah baris yang diinginkan.
   
2.	`Input` tersebut disimpan dalam variabel input.
3.	Fungsi `printStars(input)` dipanggil untuk mencetak pola bintang.
   
4.	Dalam fungsi `printStars()`:
    
	•	Jika nilai `n` lebih besar dari 0, fungsi akan memanggil dirinya sendiri dengan `n - 1`.

	•	Setelah pemanggilan rekursif selesai, program akan mencetak `n` bintang dalam satu baris.

5.	Hasil akhirnya adalah pola bintang menurun dari `n` bintang pada baris pertama hingga satu bintang di baris terakhir.


### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.

![image](https://github.com/user-attachments/assets/c84c8623-23a2-4bd9-8c6e-9fe4e050bd2f)

### Source Code
``` go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menemukan faktor bilangan
func printFactors(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	printFactors(n, i+1)
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	fmt.Print("Faktor dari ", input, ": ")
	printFactors(input, 1)
	fmt.Println()
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/60e45ec9-e03f-4967-99d8-273350951d29)
### Deskripsi Program
Program di atas berfungsi untuk menampilkan faktor-faktor dari sebuah bilangan yang diinput oleh pengguna. Program ini menggunakan fungsi rekursif untuk menemukan bilangan-bilangan yang merupakan faktor dari bilangan tersebut. Setelah pengguna memasukkan bilangan, program akan menampilkan bilangan-bilangan yang dapat membagi bilangan input secara sempurna.

### Algoritma Program :
1.	Fungsi `main()` meminta pengguna untuk memasukkan sebuah bilangan bulat `(n)`.
   
3.	Program lalu memanggil fungsi `printFactors(n, 1)` untuk menemukan dan mencetak faktor-faktor dari bilangan `n`.
   
3.	Di dalam fungsi printFactors`(n, i)`:
   
	•	Jika `i` lebih besar dari `n`, maka fungsi berhenti (kondisi akhir dari rekursi).

	•	Jika `i` adalah faktor dari `n` (yaitu `n % i == 0)`, maka `i` dicetak sebagai faktor.

	•	Fungsi kemudian memanggil dirinya sendiri dengan `i + 1`, untuk melanjutkan pencarian faktor berikutnya hingga mencapai nilai `n`.

4.	Semua faktor dicetak dalam satu baris.

### Cara Kerja Program : 
1.	Program dimulai dengan fungsi `main()`, di mana pengguna diminta memasukkan bilangan yang ingin dicari faktornya.
    
2.	Input dari pengguna disimpan dalam variabel input.
   
3.	Program menampilkan teks "Faktor dari (input):" lalu memanggil `printFactors(input, 1)` untuk mencari faktor-faktornya.
   
4.	Dalam fungsi `printFactors()`:
   
•	Jika `i` melebihi `n`, fungsi berhenti.

•	Jika `i` memenuhi syarat sebagai faktor dari `n` (dengan `n % i == 0`), maka `i` dicetak.

•	Fungsi memanggil dirinya lagi dengan `i + 1`, dan proses ini berlanjut hingga semua faktor ditemukan.

5.	Setelah semua faktor selesai dicetak, program menutup dengan baris kosong.

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu

Masukan terdiri dari bilangan bulat positif N

Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.
![image](https://github.com/user-attachments/assets/8c7495fa-b451-4a6b-926c-15b074a3786c)

### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menampilkan barisan bilangan dari N ke 1 dan kembali ke N
func printSequence(n, current int) {
	// Basis rekursi
	if current == 1 {
		fmt.Print(current, " ")
		return
	}

	// Tampilkan angka saat ini dan panggil fungsi secara rekursif menurun
	fmt.Print(current, " ")
	printSequence(n, current-1)

	// Tampilkan angka saat kembali dari rekursi
	fmt.Print(current, " ")
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	fmt.Print("Hasil: ")
	printSequence(input, input)
	fmt.Println()
}
```
#### Output Program

![image](https://github.com/user-attachments/assets/9d94f93b-c4a0-44c4-828d-5c1114d5b06f)

### Deskripsi Program
untuk mencetak urutan angka yang dimulai dari bilangan `N` menurun hingga `1`, lalu kembali lagi dari `1` naik hingga `N`. Program ini menggunakan fungsi rekursif untuk mengatur pencetakan angka secara menurun dan naik. Pengguna cukup memasukkan bilangan bulat `N`, kemudian program akan menampilkan urutan angka yang diinginkan sesuai dengan format tersebut.
### Algoritma Program :
1.	Program meminta pengguna memasukkan bilangan bulat positif `N`.
   
2.	Program kemudian memanggil fungsi `printSequence(N, N)`, di mana `N` adalah batas maksimum dari urutan angka yang akan dicetak.
   
3.	Di dalam fungsi `printSequence(n, current)`:
   
	•	Jika current bernilai 1, program akan mencetak 1 dan menghentikan panggilan rekursif (sebagai kondisi dasar rekursi).

	•	Jika current lebih dari 1, program akan mencetak current, kemudian memanggil` printSequence(n, current - 1)` untuk mencetak angka berikutnya yang 		lebih kecil, hingga mencapai `1`.

	•	Setelah seluruh urutan menurun tercetak, fungsi akan kembali ke setiap pemanggilan sebelumnya dan mencetak angka current lagi, membentuk urutan 		angka naik.

4.	Proses rekursif ini berakhir setelah seluruh urutan angka dari `N` ke `1` dan kembali ke `N` selesai dicetak.
   
5.	Program kemudian menampilkan hasilnya dalam satu baris.

### Cara Kerja Program : 
1.	Program dimulai dengan fungsi main(), di mana pengguna diminta memasukkan bilangan `N`.
   
2.	Input tersebut disimpan dalam variabel `input`.
   
3.	Program kemudian mencetak teks "Hasil: " dan memanggil fungsi `printSequence(input, input)` untuk mencetak urutan angka yang diinginkan.
   
4.	Di dalam fungsi `printSequence(n, current)`:
   
	•	Jika `current` bernilai 1, fungsi mencetak 1 dan berhenti, yang berfungsi sebagai batas rekursi agar tidak terjadi pemanggilan tak berujung.

	•	Jika `current` lebih dari 1:

	-	Program mencetak nilai current.
  
	-	Fungsi `printSequence()` dipanggil kembali dengan `current - 1`, sehingga nilai `current` berkurang setiap kali hingga mencapai 1.
  
	-	Setelah mencapai dan mencetak `1`, fungsi kembali ke pemanggilan sebelumnya dan mencetak `current` lagi, membentuk urutan angka yang naik 			kembali ke `N`.
  
5.	Hasil akhirnya adalah urutan angka dari `N` ke `1`, lalu kembali naik hingga `N` dalam satu baris.

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
Masukan terdiri dari sebuah bilangan bulat positif N.

Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

![image](https://github.com/user-attachments/assets/e8d2f31d-e831-4a00-a1ed-2190c9a578e8)

### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil dari 1 hingga N
func printOddSequence(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printOddSequence(n, current+1)
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)
	printOddSequence(input, 1)
	fmt.Println()
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/a33fa7a3-4d31-4085-afd5-d33bc7210343)
### Deskripsi Program
Program di atas digunakan untuk menampilkan bilangan ganjil dari 1 hingga bilangan `N` yang dimasukkan oleh pengguna. Program menggunakan fungsi rekursif untuk memeriksa setiap angka dari 1 hingga `N` dan mencetak hanya angka-angka ganjil. Pengguna memasukkan sebuah bilangan bulat positif, dan program akan menampilkan bilangan-bilangan ganjil dalam rentang tersebut.

### Algoritma Program :
1.	Program meminta pengguna memasukkan sebuah bilangan bulat positif `N`.
   
2.	Program memanggil fungsi `printOddSequence(N, 1)`, di mana `N` adalah batas atas, dan `1` adalah angka awal yang akan diperiksa.
   
3.	Di dalam fungsi `printOddSequence(n, current)`:
   
	•	Jika `current` melebihi `n`, fungsi berhenti (sebagai kondisi dasar rekursi).

	•	Jika `current` adalah bilangan ganjil (yaitu `current % 2 != 0)`, maka current dicetak.

	•	Fungsi kemudian memanggil dirinya sendiri dengan `current + 1`, sehingga berlanjut hingga mencapai `N`.

4.	Proses ini berlanjut hingga semua bilangan ganjil dalam rentang dari 1 hingga `N` dicetak.
   
5.	Setelah rekursi selesai, program menampilkan hasilnya dalam satu baris.

### Cara Kerja Program :
1.	Program dimulai di fungsi `main()`, di mana pengguna diminta untuk memasukkan nilai `N`.
   
2.	Input disimpan dalam variabel input.
   
3.	Program memanggil fungsi `printOddSequence(input, 1)` untuk mencetak bilangan ganjil dari 1 hingga `N`.
   
4.	Di dalam `printOddSequence(n, current)`:
   
	•	Jika `current` lebih besar dari `n`, fungsi berhenti, menghindari pemanggilan rekursif tanpa henti.

	•	Jika current adalah ganjil `(current % 2 != 0)`, maka current dicetak.

	•	Fungsi memanggil dirinya kembali dengan `current + 1`, sehingga current bertambah hingga mencapai `n`.

5.	Setelah semua bilangan ganjil dari 1 hingga `N` tercetak, program mengakhiri output dengan baris kosong.

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
Masukan terdiri dari bilangan bulat x dan y.

Keluaran terdiri dari hasil x dipangkatkan y.

Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".
![image](https://github.com/user-attachments/assets/511b08ca-55e7-4be0-a3d6-2eb7ff846f99)

### Source Code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var base, exponent int
	fmt.Print("Masukkan basis (x): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scan(&exponent)
	result := power(base, exponent)
	fmt.Printf("%d pangkat %d = %d\n", base, exponent, result)
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/44a55ad0-41b8-4c58-9ccc-1afd06a9a9c3)

### Deskripsi Program
Program di atas digunakan untuk menghitung nilai x pangkat y menggunakan metode rekursif. Pengguna diminta untuk memasukkan dua angka, yaitu x sebagai basis dan y sebagai eksponen. Program akan menghitung hasil dari x pangkat y melalui fungsi rekursi dan menampilkan hasilnya. Contoh, jika x = 2x dan y = 3y, maka hasilnya adalah 2^3 = 8.
### Algoritma Program :
1.	Program meminta pengguna untuk memasukkan dua bilangan bulat: x (basis) dan y (eksponen).
   
2.	Program memanggil fungsi `power(x, y)` untuk menghitung nilai x pangkat y secara rekursif.
   
3.	Dalam fungsi `power(x, y)`:
   
	•	Jika y sama dengan 0, fungsi mengembalikan 1, karena setiap bilangan yang dipangkatkan 0 hasilnya adalah 1.

	•	Jika y lebih besar dari 0, fungsi mengembalikan hasil perkalian x . power(x,y−1) di mana y dikurangi satu pada setiap pemanggilan rekursif hingga
		mencapai kondisi dasar y == 0

4.	Setelah perhitungan selesai, program menampilkan hasilnya dalam format x pangkat y = hasil y = hasil.

### Cara Kerja Program : 
1.	Program dimulai pada fungsi `main()`, di mana pengguna diminta untuk memasukkan nilai basis x dan eksponen y.
   
2.	Input dari pengguna disimpan dalam variabel `base` (sebagai x) dan `exponent` (sebagai y).
   
2.	Program memanggil fungsi `power(base, exponent)` untuk menghitung x pangkat y.
   
4.	Dalam fungsi `power(x, y)`:
   
	•	Jika y sama dengan 0, fungsi mengembalikan 1, yang merupakan kondisi dasar dari rekursi.

	•	Jika y lebih besar dari 0, fungsi mengalikan x dengan hasil pemanggilan `power(x, y-1)`, sehingga nilai y berkurang 1 setiap kali hingga mencapai 		0.

5.	Hasil dari fungsi power() disimpan dalam variabel `result`.
    
6.	Program mencetak hasil dalam format yang sesuai.

