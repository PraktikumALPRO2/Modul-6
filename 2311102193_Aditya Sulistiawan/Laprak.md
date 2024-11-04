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
Aditya Sulistiawan
  <br>
  2311102193
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
Dilansir dari Geeks for Geeks fungsi rekursif adalah proses di mana suatu fungsi memanggil dirinya sendiri secara langsung atau tidak langsung. Dengan menggunakan algoritma rekursif, masalah tertentu dapat diselesaikan dengan cukup mudah. Fungsi rekursif juga diartikan sebagai fungsi yang memanggil dirinya sendiri secara berulang dalam proses pengolahan data atau pemanggilan fungsi. Pada beberapa kasus fungsi rekursif bisa lebih mudah dipahami dan sederhana daripada dengan solusi iteratif.

### <strong> Beberapa karakteristik utama dari rekursif:

- harus ada statement kondisional biar aktivitas rekursi-nya berhenti,
- kalau nggak ada kondisional untuk terminate, bakal ada infinite recursion, jadinya system crash,
- makan banyak kekuatan prosesor dan memori karena pemanggilan fungsi bakal ditumpuk sampai manggil fungsi yang terakhir,
- kode lebih simpel (dibanding iteratif).
- 
### <strong> Contoh penerapan rekursif dalam pemrograman antara lain:
- Menghitung Faktorial: Fungsi rekursif untuk menghitung faktorial suatu bilangan.
   **Faktorial**:
   ```go
  package main
import (
 "fmt"
)
func getFactorial(n int) int {
 return n
}
func main() {
 val := getFactorial(5)
}
   ```
  
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
![Gui1](https://github.com/user-attachments/assets/afd70a56-635e-4160-b37f-60484bd7fdf9)

### SS CODE
![SC](https://github.com/user-attachments/assets/039a019f-23cf-4f77-ad59-0dd1544e4c17)


### Deskripsi Program
Fungsi baris melakukan cetak bilangan, kemudian memanggil dirinya sendiri dengan nilai 
bilangan yang dikurangi satu. Proses akan terus berlanjut hingga bilangan mencapai satu, di 
mana ia mencetak satu dan berhenti. Hasilnya yaitu deretan bilangan yang dicetak dari n hingga 
secara menurun 1.

### Algoritma dari program tersebut sebagai berikut :
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
![our](https://github.com/user-attachments/assets/700dc7bc-c750-4e04-99aa-f6ba2be7b6ff)


### Source code
![carbon](https://github.com/user-attachments/assets/41462763-ca2e-49ad-a4cb-170e4e2c5511)


### Deskripsi Program
Program tersebut membaca bilangan n yang diinput pengguna dan akan mencetak hasil 
penjumlahan dari n hingga 1 menggunakan fungsi rekursif penjumlahan. Fungsi itu akan 
menambahkan n dengan hasil pemanggilan rekursif pada n-1 hingga mencapai 1. 

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
![ou](https://github.com/user-attachments/assets/39103292-e460-4ba4-b733-9d91ed4743d2)

### SC
![carbon](https://github.com/user-attachments/assets/6c26ef4b-e3e2-4d6d-b1e0-4ab5912d9aae)


### Deskripsi Program
Program ini ditulis dalam bahasa Go dan berfungsi untuk menghitung nilai dari (2^n) menggunakan pendekatan rekursif. Pengguna diminta untuk memasukkan nilai (n), dan program akan menghitung serta menampilkan hasil dari (2) pangkat (n).

Algoritma dari program tersebut sebagai berikut :
1. Minta input dari pengguna untuk nilai (n).
2. Definisikan fungsi rekursif pangkat yang menerima parameter (n):
	Jika (n) sama dengan 0, kembalikan 1 (karena (2^0 = 1)).
	Jika tidak, kembalikan hasil dari (2) dikali hasil dari pemanggilan fungsi pangkat dengan parameter (n-1).
3. Panggil fungsi pangkat dengan nilai (n) yang dimasukkan oleh pengguna.
4. Tampilkan hasil perhitungan ke layar.

Cara kerja dari program ini yaitu user menginputkan bilangan. Setelah itu, panggil fungsi 'pangkat' untuk menghitung hasil 2^n secara rekursif. Fungsi memanggil dirinya sendiri dengan nilai 'n' yang berkurang hingga 0. Fungsi akan mengalikan 2 dengan hasil pangkat sebelumnya.

## Guided - 4
### Study Case
**Mencari nilai faktorial atau n!, Base-Case n == 0 atau n == 1**
#### Source Code
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

	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}
```
### Screenshot Output
![ou](https://github.com/user-attachments/assets/7867ead3-5289-489a-a5d0-d95604cc46ce)

### SC
![carbon](https://github.com/user-attachments/assets/5e53aef1-0de1-4806-8ca1-2c8a26556e77)



### Deskripsi Program
Program ini ditulis dalam bahasa Go dan berfungsi untuk menghitung nilai faktorial dari sebuah bilangan bulat non-negatif (n). Program ini menggunakan pendekatan rekursif untuk melakukan perhitungan faktorial.

### Algoritma dari program tersebut sebagai berikut :
1. Definisikan fungsi faktorial yang menerima parameter (n):
	- Jika (n) sama dengan 0 atau 1, kembalikan 1 (karena (0! = 1) dan (1! = 1)).
	- Jika tidak, kembalikan hasil dari (n) dikali hasil dari pemanggilan fungsi faktorial dengan parameter (n-1).
2. Dalam fungsi main, minta input dari pengguna untuk nilai (n).
3. Panggil fungsi faktorial dengan nilai (n) yang dimasukkan oleh pengguna.
4. Tampilkan hasil perhitungan faktorial ke layar.

Cara kerja dari program ini yaitu user menginputkan bilangan. Fungsi 'faktorial' digunakan untuk menghitung faktorial dari bilangan secara rekursif. Fungsi akan memanggil dengan nilai n yang berkurang hingga 0 atau 1. Setiap panggilan akan mengalikan n dengan hasil faktorial dari n-1.


# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S=Sn-1+Sn-2 Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonассі tersebut.**

#### Source Code
```go
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print("n: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
	fmt.Print("Sn: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

	fmt.Println()
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/dd1e9df0-1d4b-4622-9f99-84d35e655e50)


### Output
![ou](https://github.com/user-attachments/assets/bd9512f6-a2e8-4828-85c6-5b365fa751cc)


### Deskripsi Program
Program ini ditulis dalam bahasa Go dan berfungsi untuk menghitung dan menampilkan deret Fibonacci hingga suku ke-10. Program menggunakan fungsi rekursif untuk menghitung nilai Fibonacci. Deret Fibonacci adalah deret angka di mana setiap angka adalah jumlah dari dua angka sebelumnya, dimulai dengan 0 dan 1.

### Algoritma Program
1. Definisikan fungsi fibonacci yang menerima parameter (n):
	- Jika (n) kurang dari atau sama dengan 1, kembalikan (n) (karena (F(0) = 0) dan (F(1) = 1)).
	- Jika tidak, kembalikan hasil dari penjumlahan fibonacci(n-1) dan fibonacci(n-2).
2. Dalam fungsi main, tampilkan angka dari 0 hingga 10 untuk menunjukkan indeks deret Fibonacci.
3. Tampilkan hasil dari fungsi fibonacci untuk setiap indeks dari 0 hingga 10.
4. Cetak hasilnya dalam format yang mudah dibaca.

### Cara kerja program
1. Program dimulai dengan mengimpor paket fmt untuk keperluan input dan output.
2. Fungsi fibonacci diimplementasikan untuk menghitung nilai Fibonacci secara rekursif.
3. Dalam fungsi main, program mencetak indeks dari 0 hingga 10, yang merepresentasikan posisi dalam deret Fibonacci.
4. Program kemudian mencetak nilai Fibonacci untuk setiap indeks dari 0 hingga 10 dengan memanggil fungsi fibonacci.
5. Hasil deret Fibonacci ditampilkan dalam format yang terpisah dengan spasi.

## Unguided - 2
### Study Case
**Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.**

#### Source Code
```go
package main
import (
    "fmt"
)

func cetakBintang(n int) {
    if n <= 0 {
        return
    }
    for i := 0; i < n; i++ {
        fmt.Print("*")
    }
    fmt.Println()
    cetakBintang(n - 1) 
}

func tampilkanPola(n int) {
    if n <= 0 {
        fmt.Println("Masukan harus lebih besar dari 0")
        return
    }
    cetakBintang(n)
}

func main() {
    var n int
    
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&n) 
    
    fmt.Printf("\nPola bintang untuk N = %d:\n", n)
    tampilkanPola(n) 
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/43ab1f1b-04fe-4e86-b804-ea42b1ba6306)

### Output
![OU](https://github.com/user-attachments/assets/89892e7f-eb64-49e2-8ace-987d1c0794d2)


### Deskripsi Program
Program ini ditulis dalam bahasa Go dan berfungsi untuk mencetak pola bintang berbentuk segitiga terbalik berdasarkan input angka (n) dari pengguna. Setiap baris akan mencetak bintang sebanyak (n) untuk baris pertama, (n-1) untuk baris kedua, dan seterusnya hingga baris terakhir yang mencetak 1 bintang. Jika pengguna memasukkan angka kurang dari atau sama dengan 0, program akan memberikan pesan kesalahan.

### Algoritma Program
1. Definisikan fungsi cetakBintang(n int):

- Jika (n) kurang dari atau sama dengan 0, keluar dari fungsi.
- Cetak (n) bintang di satu baris.
- Panggil cetakBintang(n - 1) untuk mencetak baris berikutnya dengan jumlah bintang yang lebih sedikit.

2. Definisikan fungsi tampilkanPola(n int):

- Jika (n) kurang dari atau sama dengan 0, tampilkan pesan kesalahan dan keluar dari fungsi.
- Panggil fungsi cetakBintang(n) untuk mencetak pola bintang.

3. Dalam fungsi main:

- Minta input dari pengguna untuk nilai (n).
- Tampilkan pola bintang dengan memanggil tampilkanPola(n).

### Cara kerja program
1. Program dimulai dengan mengimpor paket fmt untuk keperluan input dan output.
2. Fungsi cetakBintang diimplementasikan untuk mencetak bintang secara rekursif. Setiap kali fungsi dipanggil, jumlah bintang yang dicetak berkurang satu hingga mencapai 0.
3. Fungsi tampilkanPola memeriksa validitas input dan memanggil fungsi cetakBintang jika input valid.
4. Dalam fungsi main, program meminta pengguna untuk memasukkan angka, kemudian memanggil tampilkanPola untuk mencetak pola bintang berdasarkan input tersebut.


## Unguided - 3
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.**                                        
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                        
*Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).*

#### Source Code
```go
package main
import (
    "fmt"
)

func findFactors(n int, current int) {
    if current > n {
        return
    }

    if n%current == 0 {
        fmt.Printf("%d ", current)
    }

    findFactors(n, current+1)
}

func main() {
    var n int
    
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Mohon masukkan bilangan positif!")
        return
    }
    
    fmt.Printf("Faktor dari %d adalah: ", n)
    findFactors(n, 1)
    fmt.Println() 
}package main
import (
    "fmt"
)

func findFactors(n int, current int) {
    if current > n {
        return
    }

    if n%current == 0 {
        fmt.Printf("%d ", current)
    }

    findFactors(n, current+1)
}

func main() {
    var n int
    
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Mohon masukkan bilangan positif!")
        return
    }
    
    fmt.Printf("Faktor dari %d adalah: ", n)
    findFactors(n, 1)
    fmt.Println() 
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/4d90e8d8-025f-43bc-a162-ed1ef42ef4cc)


### Output
![ou](https://github.com/user-attachments/assets/a79dce1d-49f0-4ac4-8f2f-f3403312f82c)


### Deskripsi Program
Fungsi printFactors menerima dua parameter: n (angka yang ingin Anda cari faktornya) dan i 
(angka yang ingin Anda periksa untuk melihat apakah faktornya). Jika i adalah faktor dari n (n 
% i == 0), maka i  dicetak. Fungsi ini dijalankan secara rekursif dari nilai i = 1 hingga i > n. 
Pada fungsi utama, pengguna diminta  memasukkan nilai n. Program kemudian  mencetak 
faktor n menggunakan fungsi printFactors. 

### Algoritma Program
1. Definisikan fungsi findFactors(n int, current int):

- Jika current lebih besar dari n, keluar dari fungsi.
- Jika n dapat dibagi habis oleh current (yaitu n % current == 0), cetak current sebagai faktor.
- Panggil findFactors(n, current + 1) untuk melanjutkan pencarian faktor dengan nilai current yang lebih besar.

2. Dalam fungsi main:

- Deklarasikan variabel n untuk menyimpan input pengguna.
- Minta pengguna untuk memasukkan bilangan positif.
- Jika n kurang dari atau sama dengan 0, tampilkan pesan kesalahan dan keluar dari fungsi.
- Tampilkan pesan yang menunjukkan faktor dari n dan panggil findFactors(n, 1) untuk memulai pencarian faktor dari 1.

### Cara kerja program
1. Program dimulai dengan mengimpor paket fmt untuk keperluan input dan output.
2. Fungsi findFactors diimplementasikan untuk mencari dan mencetak faktor dari bilangan (n) secara rekursif.
	- Fungsi ini memeriksa setiap angka mulai dari 1 hingga (n) untuk melihat apakah angka tersebut adalah faktor.
3. Dalam fungsi main, program meminta pengguna untuk memasukkan bilangan positif.
4. Jika pengguna memasukkan bilangan yang tidak valid (kurang dari atau sama dengan 0), program akan menampilkan pesan kesalahan.
	- Jika input valid, program mencetak faktor dari bilangan yang dimasukkan dengan memanggil fungsi findFactors mulai dari 1.
5. Hasil faktor ditampilkan dalam satu baris, dipisahkan oleh spasi.

## Unguided - 4
### Study Case
**4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.**                             
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                                 
*Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.*

#### Source Code
```go
package main

import "fmt"

// Fungsi rekursif untuk mencetak barisan
func cetakBarisan(n, nilai int) {
	fmt.Printf("%d ", nilai)
	if nilai == 1 {
		return
	}

	// Mencetak barisan
	cetakBarisan(n, nilai-1)
	fmt.Printf("%d ", nilai)
}

func main() {
	// Inputan user untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan barisan bilangan
	fmt.Printf("\nBarisan bilangan dari %d hingga 1 dan kembali ke %d adalah: ", n, n)
	cetakBarisan(n, n)
	fmt.Println()
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/bcdd204c-50ac-4345-b456-ea6e24b59663)


### Output
![ou](https://github.com/user-attachments/assets/d60e69ae-f20e-490c-8e47-ca3cca2c94a6)


### Deskripsi Program
n (nilai awal yang dimasukkan oleh pengguna) dan saat ini (nilai yang  diproses). Fungsi ini 
mencetak nilai saat ini dan memanggil dirinya sendiri dengan current-1 hingga mencapai 0. 
Kemudian, jika nilai saat ini tidak sama dengan n, program akan mencetak ulang nilai saat ini 
untuk membuat urutan menaik.

### Algoritma Program
1. Definisikan fungsi cetakBarisan(n int, nilai int):

- Cetak nilai saat ini.
- Jika nilai sama dengan 1, keluar dari fungsi.
- Panggil cetakBarisan(n, nilai - 1) untuk mencetak nilai yang lebih kecil.
- Setelah kembali dari pemanggilan rekursif, cetak nilai lagi.

2. Dalam fungsi main:

- Deklarasikan variabel n untuk menyimpan input pengguna.
- Minta pengguna untuk memasukkan bilangan bulat positif.
- Tampilkan pesan yang menunjukkan barisan bilangan dari n hingga 1 dan kembali ke n.
- Panggil fungsi cetakBarisan(n, n) untuk memulai pencetakan barisan.

### Cara kerja program
1. Program dimulai dengan mengimpor paket fmt untuk keperluan input dan output.
2. Fungsi cetakBarisan diimplementasikan untuk mencetak barisan bilangan secara rekursif.
	- Fungsi ini pertama-tama mencetak nilai saat ini.
	- Jika nilai mencapai 1, fungsi akan berhenti.
	- Fungsi kemudian dipanggil kembali dengan nilai yang berkurang, sehingga mencetak nilai dari (n) hingga 1.
	- Setelah mencapai 1 dan kembali dari pemanggilan rekursif, fungsi mencetak nilai lagi, sehingga membentuk urutan yang naik kembali ke (n).
3. Dalam fungsi main, program meminta pengguna untuk memasukkan bilangan bulat positif.
	- Setelah mendapatkan input, program mencetak barisan bilangan dari (n) hingga 1 dan kembali ke (n) dengan memanggil cetakBarisan(n, n).


## Unguided - 5
### Study Case
**5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.**
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

#### Source Code
```go
package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga N
func bilGanjil(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	bilGanjil(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan: ")
	fmt.Scan(&n)

	bilGanjil(n, 1)
	fmt.Println() // 2311102193
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/b9d8c3da-2ec2-4f58-9e49-0f5ddeeba82a)


### Output
![ou](https://github.com/user-attachments/assets/8d359e1b-4eea-42be-9334-066a036c5a32)


### Deskripsi Program
Fungsi bilangan ganjil menerima dua parameter n (batas atas yang dimasukkan oleh pengguna) 
dan saat ini (angka ganjil yang akan diproses). Fungsi ini mencetak nilai saat ini dan memanggil 
dirinya sendiri, yang menambahkan 2 ke nilai saat ini hingga nilainya melebihi n. Pada fungsi 
utama, program meminta pengguna untuk memasukkan nilai n dan memanggil bilangan ganjil 
untuk mencetak semua bilangan ganjil dari 1 hingga n.

### Algoritma Program
1. Definisikan fungsi bilGanjil(n int, current int):

- Jika current lebih besar dari n, keluar dari fungsi.
- Cetak nilai current.
- Panggil fungsi bilGanjil(n, current + 2) untuk melanjutkan ke bilangan ganjil berikutnya.

2. Dalam fungsi main:

- Deklarasikan variabel n untuk menyimpan input pengguna.
- Minta pengguna untuk memasukkan bilangan bulat.
- Panggil fungsi bilGanjil(n, 1) untuk memulai pencetakan bilangan ganjil dari 1 hingga (n).

### Cara kerja program
Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan. Fungsi 'cetakGanjil' digunakan untuk mencetak semua bilangan ganjil dari 1 hingga bilangan yang sesuai dengan inputan user. Fungsi akan mengecek dengan cara rekursif, mencetak angka jika bilangan ganjil. Lalu, melanjutkan hingga bilangan yang diminta.


## Unguided - 6
### Study Case
** Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.**                                 
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                        
*Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.*

#### Source Code
```go
package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y) // 2311102193

	fmt.Println(pangkat(x, y))
}
```
### Screenshot Code
![carbon](https://github.com/user-attachments/assets/51a11136-4947-4d23-9f56-8a27e300491e)

### Output
![ou](https://github.com/user-attachments/assets/dcf6b292-b86a-4d9b-a2f5-bc8a05fef2d8)


### Deskripsi Program
Jika y adalah 0, fungsi tersebut mengembalikan 1 karena  bilangan yang dipangkatkan 0 adalah 
1. Jika tidak, fungsi tersebut akan memanggil dirinya sendiri dengan y-1 dan mengalikan 
hasilnya dengan x. Pada fungsi utama, pengguna diminta  memasukkan nilai x dan y. 
Kemudian, dengan menggunakan fungsi pangkat, hasil  x pangkat y dihitung dan ditampilkan 
di layar.

### Algoritma Program
1. Definisikan Fungsi pangkat(x int, y int) int:

- Jika ( y ) sama dengan 0, kembalikan 1 (karena setiap bilangan pangkat 0 adalah 1).
- Jika tidak, kembalikan hasil perkalian ( x ) dengan hasil pemanggilan fungsi pangkat(x, y - 1) (rekursif).

2. Fungsi main:

- Deklarasikan dua variabel x dan y untuk menyimpan input pengguna.
- Tampilkan prompt untuk meminta pengguna memasukkan nilai ( x ) dan ( y ).
- Baca nilai ( x ) dan ( y ) dari input pengguna.
- Panggil fungsi pangkat(x, y) dan cetak hasilnya.

### Cara kerja program
Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan dan bilangan pangkat. Fungsi 'pangkat' digunakan untuk menghitung hasil pangkat secara rekursif. Fungsi akan memanggil dirinya sendiri dengan nilai pangkat yang berkurang sampai 0, dimana hasilnya 1.

## Referensi

[1] Fazry, M. R. (2021, December 9). Cara Rekursif di Golang - Muhammad Redyka Fazry - Medium. Medium. https://medium.com/@fazry/cara-rekursif-di-golang-cc202f55e336
