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
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung deret Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("n\tSn")
	// Cetak tabel deret Fibonacci hingga suku ke-10
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t%d\n", i, fibonacci(i))
	}
}

```

### Output:
![image](https://github.com/user-attachments/assets/c25403e8-2c4c-4acf-9182-43fc39e8b2b6)

### Deskripsi Program : 
Program ini menghitung dan mencetak deret Fibonacci hingga suku ke-10 menggunakan fungsi rekursif dan loop for. Setiap suku adalah jumlah dari dua suku sebelumnya, dimulai dari 0 dan 1.

### Algoritma Program
1. Mulai
2. Cetak "n" dan "Sn" sebagai header tabel
3. Untuk i dari 0 hingga 10:
   - Panggil fungsi fibonacci(i)
   - Cetak nilai i dan hasil dari fibonacci(i)
4. Fungsi fibonacci:
   - Jika n <= 1 maka kembalikan n
   - Jika tidak, kembalikan hasil dari fibonacci(n-1) + fibonacci(n-2)
5. Selesai
   
### Cara Kerja Program:
1. Program mulai dan mencetak header tabel "n" dan "Sn".
2. Loop for menjalankan perulangan dari i = 0 hingga i = 10.
3. Untuk setiap nilai i, program memanggil fungsi fibonacci(i) untuk menghitung nilai Fibonacci ke-i.
4. Fungsi fibonacci bekerja secara rekursif:
   - Jika n <= 1, fungsi mengembalikan nilai n.
   - Jika n > 1, fungsi mengembalikan jumlah dari fibonacci(n-1) dan fibonacci(n-2).
5. Program mencetak nilai i dan hasil dari fungsi fibonacci(i).
6. Proses ini berlanjut sampai semua nilai dari i = 0 hingga i = 10 dihitung dan dicetak.

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak pola bintang
func cetakBintang(n int) {
	// Kondisi akhir rekursi: jika n <= 0, hentikan fungsi
	if n <= 0 {
		return
	}
	// Panggil cetakBintang dengan nilai n-1 lebih dulu
	cetakBintang(n - 1)
	// Tampilkan bintang sejumlah n
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah baris bintang: ")
	fmt.Scan(&jumlah)

	fmt.Println("Hasil pola bintang:")
	cetakBintang(jumlah)
}
```

### Output:
![image](https://github.com/user-attachments/assets/7b6b8b02-3dfa-45c9-910b-871e9ce60a4b)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan jumlah baris bintang, kemudian mencetak pola bintang secara rekursif. Fungsi cetakBintang akan menampilkan satu baris tambahan dengan jumlah bintang yang sesuai untuk setiap panggilan rekursif, hingga mencapai jumlah yang diminta.

### Algoritma Program:
1. Mulai
2. Minta input dari pengguna untuk jumlah baris bintang
3. Panggil fungsi cetakBintang dengan nilai input dari pengguna
4. Fungsi cetakBintang:
   - Jika n <= 0, kembali (berhenti)
   - Panggil ulang fungsi cetakBintang dengan n - 1
5. Cetak bintang sejumlah n
6. Cetak garis baru
7. Selesai

### Cara Kerja Program:
1. Program meminta pengguna memasukkan jumlah baris bintang yang diinginkan.
2. Fungsi cetakBintang dipanggil dengan nilai input yang diberikan pengguna.
3. Fungsi bekerja secara rekursif:
   - Jika n lebih kecil atau sama dengan 0, fungsi berhenti.
   - Jika n lebih besar dari 0, fungsi memanggil dirinya sendiri dengan n - 1.
4. Setelah pemanggilan rekursif, fungsi mencetak sejumlah bintang sesuai nilai n.
5. Proses ini berlanjut hingga semua baris bintang tercetak sesuai jumlah yang diminta.
6. Program selesai dan menampilkan pola bintang di layar.


### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.**

![image](https://github.com/user-attachments/assets/28a4e99e-3018-4ea4-9f3d-8812a8eeffa8)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan faktor-faktor dari N
func cariFaktor(n, i int) {
	// Jika i lebih besar dari n, maka rekursi berhenti
	if i > n {
		return
	}
	// Jika n habis dibagi i, maka i adalah faktor dari n
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	// Panggil fungsi rekursif untuk nilai i berikutnya
	cariFaktor(n, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor-faktor dari %d adalah: ", N)
	cariFaktor(N, 1)
	fmt.Println()
}

```

### Output:
![image](https://github.com/user-attachments/assets/c7925049-d42c-46d2-b7d6-4f6e0fe47b36)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan bilangan bulat positif 𝑁. Kemudian, menggunakan fungsi rekursif cariFaktor, program menampilkan semua faktor dari 𝑁. Fungsi ini memeriksa setiap angka dari 1 hingga 𝑁untuk melihat apakah angka tersebut adalah faktor dari 𝑁. Jika ya, angka tersebut dicetak. Proses ini berlanjut hingga semua faktor ditemukan dan dicetak.

### Algoritma Program:
1. Mulai
2. Minta pengguna untuk memasukkan bilangan bulat positif N
3. Panggil fungsi cariFaktor dengan parameter n = N dan i = 1
4. Fungsi cariFaktor:
   - Jika i > n, berhenti
   - Jika n % i == 0, cetak i sebagai faktor dari n
   - Panggil fungsi cariFaktor dengan i + 1
5. Selesai

### Cara Kerja Program:
1. Program meminta pengguna memasukkan bilangan bulat positif N.
2. Program memanggil fungsi cariFaktor mulai dari i = 1.
3. Fungsi cariFaktor memeriksa setiap angka dari 1 hingga N:
   - Jika i lebih besar dari N, berhenti.
   - Jika N habis dibagi i, maka i adalah faktor dari N dan dicetak.
   - Fungsi dipanggil lagi dengan nilai i ditambah 1.
4. Proses berlanjut hingga semua faktor N ditemukan dan dicetak.
5. Program selesai menampilkan faktor-faktor dari N.

### 4.

![image](https://github.com/user-attachments/assets/873da708-8303-4f19-8db6-b3d70921da7b)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan barisan dari N hingga 1
func turunNaik(n, current int) {
	// Cetak nilai saat ini
	fmt.Print(current, " ")

	// Jika current mencapai 1, mulai mencetak kembali ke N
	if current == 1 {
		return
	}

	// Rekursi untuk turun ke 1
	turunNaik(n, current-1)

	// Cetak kembali saat naik lagi dari 2 hingga N
	fmt.Print(current, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali ke %d: ", N, N)
	turunNaik(N, N)
	fmt.Println()
}

```

### Output:
![image](https://github.com/user-attachments/assets/d6022da3-9d95-4a48-9a93-7cb01ef57298)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan bilangan bulat positif N, lalu mencetak barisan bilangan dari N hingga 1 dan kembali lagi ke N menggunakan fungsi rekursif turunNaik. Program mencetak bilangan menurun hingga 1, kemudian naik kembali ke N.

### Algoritma Program:
1. Mulai
2. Minta input dari pengguna berupa bilangan bulat positif N
3. Panggil fungsi turunNaik dengan parameter n = N dan current = N
4. Fungsi turunNaik:
   - Cetak nilai current
   - Jika current == 1, berhenti
   - Panggil fungsi turunNaik dengan nilai current - 1 (untuk turun sampai ke 1)
   - Cetak kembali nilai current saat naik dari 2 hingga N
5. Selesai

### Cara Kerja Program:
1. Program meminta pengguna memasukkan bilangan bulat positif N.
2. Program memanggil fungsi rekursif turunNaik mulai dari nilai N.
3. Fungsi turunNaik mencetak nilai current:
   - Jika current mencapai 1, fungsi berhenti.
   - Jika current lebih dari 1, fungsi memanggil dirinya sendiri dengan current - 1 untuk mencetak bilangan menurun hingga 1.
   - Setelah mencapai 1, fungsi kembali mencetak nilai current saat naik dari 2 hingga N.
4. Proses berlanjut sampai semua bilangan dari N hingga 1 dan kembali ke N dicetak.
5. Program selesai dan menampilkan rangkaian bilangan sesuai input pengguna.


### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.**

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan ganjil hingga N
func tampilGanjil(n, current int) {
	// Jika current melebihi n, hentikan rekursi
	if current > n {
		return
	}
	// Cetak nilai current jika ganjil
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	// Panggil fungsi untuk bilangan berikutnya
	tampilGanjil(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", N)
	tampilGanjil(N, 1)
	fmt.Println()
}

```

### Output:
![image](https://github.com/user-attachments/assets/87f9e1ab-ceaa-43c2-979c-ecc7b7ea4e9b)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan bilangan bulat positif N. Kemudian, program menampilkan semua bilangan ganjil dari 1 hingga N menggunakan fungsi rekursif tampilGanjil. Fungsi ini mencetak setiap bilangan ganjil dengan memeriksa apakah bilangan tersebut tidak habis dibagi 2 dan kemudian memanggil dirinya sendiri dengan nilai yang ditingkatkan satu hingga mencapai N.

### Algoritma Program:
1. Mulai
2. Minta input dari pengguna berupa bilangan bulat positif N
3. Panggil fungsi tampilGanjil dengan parameter n = N dan current = 1
4. Fungsi tampilGanjil:
   - Jika current > n, berhenti
   - Jika current adalah bilangan ganjil, cetak current
5. Panggil fungsi tampilGanjil dengan current + 1
6. Selesai
   
### Cara Kerja Program:
1. Program meminta pengguna memasukkan bilangan bulat positif N.
2. Program memanggil fungsi rekursif tampilGanjil dimulai dari nilai current = 1.
3. Fungsi tampilGanjil memeriksa setiap bilangan dari 1 hingga N:
   - Jika current lebih besar dari N, fungsi berhenti.
   - Jika current adalah bilangan ganjil, maka current dicetak.
   - Fungsi dipanggil lagi dengan current + 1.
4. Proses ini berlanjut sampai semua bilangan ganjil dari 1 hingga N dicetak.
5. Program selesai dan menampilkan semua bilangan ganjil dari 1 hingga N.

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
**Masukan terdiri dari bilangan bulat x dan y.**

**Keluaran terdiri dari hasil x dipangkatkan y.**

**Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan Import "math".**

### Source Code :
```go
package main

import (
    "fmt"
)

// Fungsi rekursif untuk menghitung x^y
func power(x int, y int) int {
    if y == 0 {
        return 1
    }
    return x * power(x, y-1)
}

func main() {
    var x, y int

    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan (x): ")
    fmt.Scan(&x)
    fmt.Print("Masukkan pangkat (y): ")
    fmt.Scan(&y)

    // Menampilkan hasil
    fmt.Printf("%d ^ %d = %d\n", x, y, power(x, y))
}
```

### Output:
![image](https://github.com/user-attachments/assets/d6785ff4-5e44-4f11-84b7-c324a2a6e1ef)

### Deskripsi Program: 
Program ini meminta pengguna untuk memasukkan bilangan x dan pangkat y. Kemudian, menggunakan fungsi rekursif power, program menghitung hasil dari x dipangkatkan dengan y.

### Algoritma Program:
1. Mulai
2. Minta input pengguna untuk bilangan x dan pangkat y
3. Panggil fungsi power dengan parameter x dan y
4. Fungsi power:
   - Jika y == 0, kembalikan 1
   - Jika tidak, kembalikan x kali hasil dari power(x, y-1)
5. Cetak hasil dari x dipangkatkan dengan y
6. Selesai
   
### Cara Kerja Program:
1. Program meminta pengguna memasukkan bilangan x dan pangkat y.
2. Program memanggil fungsi rekursif power untuk menghitung nilai x dipangkatkan y.
3. Fungsi power bekerja dengan:
   - Jika y adalah 0, fungsi mengembalikan nilai 1 (karena bilangan apapun dipangkatkan 0 adalah 1).
   - Jika y lebih dari 0, fungsi mengembalikan hasil perkalian x dengan hasil dari power(x, y-1), mengurangi y setiap kali sampai menjadi 0.
4. Hasil perhitungan ditampilkan di layar dalam format x ^ y = hasil.
5. Program selesai dan menampilkan hasilnya.


