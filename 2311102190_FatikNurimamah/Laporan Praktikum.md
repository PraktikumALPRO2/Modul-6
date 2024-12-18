<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 6</strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
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
4. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori
**1. Definisi Rekursif**

Rekursi adalah teknik di mana sebuah fungsi atau prosedur memanggil dirinya sendiri untuk menyelesaikan tugas tertentu. Dalam pemrograman, rekursi sering dimanfaatkan untuk menyelesaikan masalah yang bisa dipecah menjadi sub-masalah yang lebih kecil. Proses ini berlanjut sampai mencapai kondisi yang menghentikan pemanggilan fungsi tersebut.

**2. Komponen Rekursif**

Fungsi rekursif memiliki dua komponen utama:  
1. Base Case (Kasus Dasar): Kondisi ini adalah saat fungsi berhenti memanggil dirinya sendiri dan menghasilkan hasil akhir. Base case sangat penting untuk menghindari loop tanpa akhir dan memastikan rekursi berhenti pada suatu titik.  
2. Recursive Case (Kasus Rekursif): Pada bagian ini, fungsi memanggil dirinya sendiri dengan argumen yang lebih sederhana atau lebih kecil dari masalah awal, sehingga fungsi perlahan mendekati base case.

```go
package main

import "fmt"

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
    if n <= 1 { // Base case
        return 1
    }
    return n * factorial(n-1) // Recursive case
}

func main() {
    var i int = 5
    fmt.Printf("Faktorial dari %d adalah %d\n", i, factorial(i))
}
```
- Base Case: Jika `n ≤ 1`, fungsi mengembalikan `1`.
  
- Recursive Case: Jika tidak, fungsi memanggil dirinya sendiri dengan `n - 1` dan mengalikan hasilnya dengan `n` .

## Guided 

### 1. Membuat baris bilangan dari n hingga 1
**Base-case: bilangan == 1**

### Source Code :
```go
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	baris (n)
}

func baris(bilangan int){
	if bilangan == 1{
		fmt.Println(1)
	}else{
		fmt.Println(bilangan)
		baris(bilangan-1)
	}
}

```

### Output:
![Screenshot 2024-10-28 210740](https://github.com/user-attachments/assets/44b30c0e-9830-4f3b-8988-36a5821430d7)

### Full code Screenshot:
![Screenshot 2024-10-28 210857](https://github.com/user-attachments/assets/8ae6f0d5-37ad-4763-b76c-7a26526435d7)

### Deskripsi Program : 
Program tersebut adalah program sederhana yang meminta pengguna untuk memasukkan sebuah bilangan bulat positif n. Program kemudian mencetak angka dari n hingga 1 secara menurun menggunakan teknik rekursif.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan bilangan n.
2. Fungsi `baris(bilangan int)` dipanggil dengan `n` sebagai parameter.
3. Di dalam fungsi `baris`, jika nilai `bilangan` sama dengan `1`, maka fungsi mencetak `1` dan berhenti.
4. Jika `bilangan` lebih besar dari `1`, fungsi akan mencetak nilai `bilangan`, kemudian memanggil dirinya sendiri dengan nilai `bilangan - 1`.
5. Proses rekursif terus berlanjut hingga bilangan bernilai `1`.

### Cara Kerja Program:
1. Meminta Input: Program menerima masukan nilai n dari pengguna, yang merupakan bilangan bulat positif.
2. Pemanggilan Fungsi Rekursif: Program memanggil fungsi `baris` dengan parameter `n`, dan fungsi ini memeriksa nilai `bilangan`:
   - Jika `bilangan` bernilai `1`, program akan mencetak `1` dan menghentikan rekursi (ini adalah kondisi dasar rekursi).
   - Jika `bilangan` lebih besar dari `1`, program akan mencetak nilai `bilangan` dan kemudian memanggil fungsi `baris` dengan parameter `bilangan - 1`.
3. Output: Program akan mencetak deret angka menurun dari `n` hingga 1.

### 2. Menghitung hasil penjumlahan 1 hingga n 
**Base-case: n = 1**

### Source Code :
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

### Output:
![Screenshot 2024-10-28 211927](https://github.com/user-attachments/assets/014d717b-99f1-4be9-9928-d10f0d1fd3d8)

### Full code Screenshot:
![Screenshot 2024-10-28 212007](https://github.com/user-attachments/assets/eb7eaa10-e257-4a05-af32-ba9667fb4ef7)

### Deskripsi Program : 
Program tersebut adalah program sederhana yang dibuat untuk menerima input berupa bilangan bulat positif n dari pengguna. Selanjutnya, program ini menghitung jumlah semua bilangan dari 1 hingga n dengan menggunakan fungsi rekursif, kemudian menampilkan hasil penjumlahannya.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan nilai n.
2. Fungsi `penjumlahan(n int)` dipanggil dengan parameter n untuk menghitung jumlah dari 1 sampai n.
3. Di dalam fungsi `penjumlahan`:
   - Jika `n` sama dengan `1`, fungsi akan mengembalikan nilai `1`, yang berfungsi sebagai kondisi dasar (basis rekursi).
   - Jika `n` lebih dari `1`, fungsi akan mengembalikan hasil dari `n` ditambah dengan pemanggilan fungsi `penjumlahan` dengan parameter `n-1`.
4. Hasil yang diperoleh dari fungsi `penjumlahan` akan dicetak oleh fungsi `main`.
   
### Cara Kerja Program:
1. Meminta Input: Program meminta pengguna untuk memasukkan nilai n.
2. Pemanggilan Fungsi Rekursif: Program memanggil fungsi `penjumlahan` dengan `n` sebagai parameter:
   - Jika `n` sama dengan `1`, fungsi mengembalikan nilai `1` (kondisi dasar rekursi).
   - Jika `n` lebih dari `1`, fungsi akan mengembalikan hasil dari `n` ditambah hasil dari pemanggilan fungsi `penjumlahan(n-1)`.
3. Output: Setelah proses rekursi selesai dan semua penjumlahan dari `1` hingga `n` telah dilakukan, hasil akhirnya akan dicetak di layar.
   
### 3. Mencari dua pangkat n atau 2^n
**Base-case: n == 0**

### Source Code :
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

### Output:
![Screenshot 2024-10-28 213222](https://github.com/user-attachments/assets/44a18588-7201-40d5-938d-b693b59d6b33)

### Full code Screenshot:
![Screenshot 2024-10-28 213301](https://github.com/user-attachments/assets/aceed035-94e2-462b-9590-65cd1db5c89d)

### Deskripsi Program : 
Program tersebut dibuat untuk menghitung nilai `2^n`, di mana `n` merupakan bilangan bulat non-negatif yang dimasukkan oleh pengguna. Program ini memanfaatkan fungsi rekursif untuk melakukan perhitungan dan menampilkan hasilnya.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan nilai `n`.
2. Fungsi `pangkat(n int)` dipanggil dengan parameter `n` untuk menghitung `2^n`.
3. Di dalam fungsi `pangkat`:
   - Jika `n` sama dengan `0`, fungsi akan mengembalikan nilai `1` (karena `2^0 = 1`).
   - Jika `n` lebih besar dari `0`, fungsi akan mengembalikan hasil dari `2` dikalikan dengan hasil pemanggilan fungsi `pangkat` dengan parameter `n-1`.
4. Hasil akhir dari fungsi `pangkat` akan dicetak oleh fungsi `main`.

### Cara Kerja Program:
1. Meminta Input: Program meminta pengguna untuk memasukkan nilai `n`.
2. Pemanggilan Fungsi Rekursif: Program memanggil fungsi `pangkat` dengan parameter `n`:
   - Jika `n` adalah `0`, fungsi mengembalikan nilai `1`.
   - Jika `n` lebih dari `0`, fungsi akan mengembalikan hasil dari `2` dikalikan dengan hasil dari pemanggilan `pangkat(n-1)`.
3. Output: Setelah proses perhitungan rekursif selesai, hasil dari `2^n` akan ditampilkan di layar.

### 4. Mencari nilai faktorial atau n!
**Base-case: n == 0 atau n == 1**

### Source Code :
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

### Output:
![Screenshot 2024-10-28 214445](https://github.com/user-attachments/assets/e6622cf8-1eec-49a3-b1dd-376a3e9bf957)

### Full code Screenshot:
![Screenshot 2024-10-28 214519](https://github.com/user-attachments/assets/ed5b729e-90c8-4989-abb6-223cd94693e4)

### Deskripsi Program : 
Program tersebut adalah program sederhana untuk menghitung faktorial dari bilangan bulat non-negatif n yang dimasukkan oleh pengguna. Dengan menggunakan fungsi rekursif, program ini melakukan perhitungan dan menampilkan hasilnya.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan nilai `n`.
2. Fungsi `faktorial(n int)` dipanggil dengan parameter `n` untuk menghitung faktorial dari nilai tersebut.
3. Di dalam fungsi `factorial`:
   - Jika `n` adalah `0` atau `1`, fungsi akan mengembalikan `1` (karena `0! = 1` dan `1! = 1`).
   - Jika `n` lebih dari `1`, fungsi akan mengembalikan hasil dari `n` dikalikan dengan pemanggilan fungsi `factorial` dengan parameter `n-1`.
4. Hasil dari fungsi `factorial` akan dicetak oleh fungsi `main`.

### Cara Kerja Program:
1. Meminta Input: Program meminta pengguna untuk memasukkan nilai `n`.
2. Pemanggilan Fungsi Rekursif: Program memanggil fungsi `factorial` dengan parameter `n`:
   - Jika `n` adalah `0` atau `1`, fungsi mengembalikan nilai `1`.
   - Jika `n` lebih dari `1`, fungsi akan mengembalikan hasil dari `n` dikalikan dengan hasil dari pemanggilan `faktorial(n-1)`.
3. Output: Hasil faktorial dari `n` akan ditampilkan di layar.



## Unguided 

### 1. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci 
**Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilal suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S₁ =Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10.**

![Screenshot 2024-10-31 135729](https://github.com/user-attachments/assets/6f1bf39c-81b5-4142-b20f-67597a737f63)


### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung suku ke-x dalam deret Fibonacci
func MenghitungFibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return MenghitungFibonacci(x-1) + MenghitungFibonacci(x-2)
}

func main() {
	var jumlahSuku int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scan(&jumlahSuku)
	fmt.Printf("n   |")
	for i := 0; i <= jumlahSuku; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()

	fmt.Printf("Sn  |")
	for i := 0; i <= jumlahSuku; i++ {
		fmt.Printf(" %d ", MenghitungFibonacci(i))
	}
	fmt.Println()
}
	
```

### Output:
![Screenshot 2024-10-28 215633](https://github.com/user-attachments/assets/f2119486-2a1b-4823-ba7f-c283ec07a51e)

### Full code Screenshot:
![Screenshot 2024-10-28 215713](https://github.com/user-attachments/assets/c14c0cd3-03dd-4908-86e8-4b45de41707a)

### Deskripsi Program : 
Program tersebut adalah program sederhana yang digunakan untuk menghitung dan menampilkan suku-suku dalam deret Fibonacci. Pengguna diminta untuk memasukkan jumlah suku yang diinginkan, dan program ini menggunakan fungsi rekursif untuk menghitung setiap suku dalam deret tersebut.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan jumlah suku Fibonacci yang ingin dihitung.
2. Fungsi `MenghitungFibonacci(x int)` dipanggil untuk mendapatkan suku ke-x dalam deret Fibonacci.
3. Dalam fungsi `MenghitungFibonacci`:
   - Jika `x` kurang dari atau sama dengan `1`, fungsi mengembalikan nilai `x` 
   - Jika `x` lebih dari `1`, fungsi mengembalikan hasil penjumlahan dari suku ke-`x-1` dan suku ke-`x-2` melalui pemanggilan fungsi `MenghitungFibonacci` secara rekursif.
4. Hasil dari fungsi `MenghitungFibonacci` dicetak, di mana baris pertama menunjukkan indeks suku dan baris kedua menampilkan nilai-nilai suku Fibonacci.

### Cara Kerja Program:
1. Meminta Input: Program meminta pengguna untuk memasukkan jumlah suku Fibonacci yang diinginkan.
2. Pemanggilan Fungsi Rekursif: Program menggunakan loop untuk memanggil fungsi `MenghitungFibonacci` untuk setiap nilai dari `0` hingga `jumlahSuku`:
   - Fungsi akan mengembalikan nilai Fibonacci sesuai dengan rumus yang ada.
3. Output: Program mencetak hasilnya  yang menunjukkan indeks suku `(n)` dan nilai suku Fibonacci `(Sn)`.


### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang dengan menggunakan fungsi rekursif. N adalah masukan dari user

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencetak bintang pada setiap baris
func cetakJumlahBintang(jumlahBintang int) {
	if jumlahBintang == 0 {
		return
	}
	cetakJumlahBintang(jumlahBintang - 1)
	fmt.Print("*")
}

// Fungsi rekursif untuk mencetak pola bintang 
func polaBintang(totalBaris int, barisSaatIni int) {
	if barisSaatIni > totalBaris {
		return
	}
	cetakJumlahBintang(barisSaatIni)
	fmt.Println()
	polaBintang(totalBaris, barisSaatIni+1)
}

func main() {
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris (N): ")
	fmt.Scan(&jumlahBaris)

	fmt.Println("Pola bintang:")
	polaBintang(jumlahBaris, 1)
}

```

### Output:
![Screenshot 2024-10-28 220824](https://github.com/user-attachments/assets/ecdd4c3d-eced-497f-8700-600c94096234)

### Full code Screenshot:
![Screenshot 2024-10-28 220859](https://github.com/user-attachments/assets/6a562f5a-2460-4569-b349-b99bbf797cef)

### Deskripsi Program : 
Program tersebut adalah program sederhana untuk mencetak pola bintang berdasarkan jumlah baris yang diinputkan oleh pengguna. Menggunakan metode rekursif, program mencetak bintang di setiap baris dengan jumlah yang meningkat sesuai dengan nomor baris.

### Algoritma Program
1. Program meminta pengguna untuk menentukan jumlah baris yang diinginkan.
2. Fungsi `polaBintang(totalBaris int, barisSaatIni int)` dipanggil untuk mencetak pola bintang:
   - Jika `barisSaatIni` melebihi `totalBaris`, fungsi akan berhenti.
   - Jika tidak, fungsi akan memanggil `cetakJumlahBintang` untuk mencetak bintang sesuai dengan nomor baris saat ini.
   - Setelah bintang dicetak, fungsi akan memanggil dirinya sendiri untuk melanjutkan ke baris berikutnya `(barisSaatIni + 1)`.
3. Fungsi `cetakJumlahBintang(jumlahBintang int)` digunakan untuk mencetak bintang di satu baris:
   - Jika `jumlahBintang` sama dengan `0`, fungsi akan kembali tanpa mencetak apapun.
   - Jika tidak, fungsi akan memanggil dirinya sendiri dengan parameter `jumlahBintang - 1`, lalu mencetak satu bintang setelahnya.
### Cara Kerja Program:
1. Pengambilan Input: Program meminta pengguna untuk memasukkan jumlah baris yang diinginkan.
2. Pemanggilan Fungsi Rekursif:
   - Fungsi `polaBintang` dipanggil dengan total baris dan memulai `barisSaatIni` dari 1.
   - Fungsi `cetakJumlahBintang` dipanggil untuk mencetak bintang pada setiap baris, sesuai dengan jumlah bintang yang ditentukan oleh nomor baris.
3. Hasil Output: Program menampilkan pola bintang berdasarkan jumlah baris yang dimasukkan oleh pengguna.

   
### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N.**

![Screenshot 2024-10-31 140054](https://github.com/user-attachments/assets/ad20f53d-ffbb-4c04-af06-1e486b613445)


### Source Code :
```go
package main

import (
	"fmt"
)

func cariFaktor(n, x int, daftarfaktor []int) []int {
	// Kondisi berhenti: jika x melebihi n, kembalikan daftar faktor
	if x > n {
		return daftarfaktor
	}
	// Cek apakah x adalah faktor dari n
	if n%x == 0 {
		daftarfaktor = append(daftarfaktor, x)
	}
	// Panggilan rekursif untuk memeriksa angka berikutnya
	return cariFaktor(n, x+1, daftarfaktor)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&n)

	daftarfaktor := cariFaktor(n, 1, []int{})
	fmt.Println("Faktor-faktor dari", n, "adalah:", daftarfaktor)
}

```

### Output:
![Screenshot 2024-10-29 184147](https://github.com/user-attachments/assets/1848cd95-9254-46f4-92b6-f45e269283f3)

### Full code Screenshot:
![Screenshot 2024-10-29 184247](https://github.com/user-attachments/assets/4c028061-e1e2-4198-a9e3-bf693a64c9f2)

### Deskripsi Program : 
Program ini menggunakan fungsi rekursif untuk mencari semua faktor dari sebuah bilangan bulat positif yang dimasukkan oleh pengguna. Faktor-faktor tersebut kemudian dikumpulkan dalam sebuah daftar dan ditampilkan di akhir program.

### Algoritma Program
1. Program meminta pengguna untuk menginputkan suatu bilangan bulat positif.
2. Fungsi `cariFaktor(n, x int, daftarfaktor []int) []int` digunakan untuk mencari faktor-faktor dari `n`:
   - Jika `x`  melebihi `n `, proses berakhir, dan daftar faktor dikembalikan.
   - Jika `x ` merupakan faktor dari `n` (artinya hasil `n%x` adalah `0`), maka `x` dimasukkan ke dalam `daftarfaktor`.
   - Fungsi dipanggil kembali dengan nilai `x+1` untuk melanjutkan pengecekan pada bilangan berikutnya.
3. Daftar faktor yang berhasil dikumpulkan kemudian dikembalikan ke fungsi` main` dan ditampilkan.

### Cara Kerja Program:
1. Pengambilan Input: Program meminta pengguna menginputkan bilangan bulat positif.
2. Pemanggilan Fungsi Rekursif:
   - Fungsi `cariFaktor` dimulai dari `x = 1` dan terus memeriksa hingga `x` melebihi `n`.
   - Setiap kali `x` terbukti sebagai faktor, nilainya ditambahkan ke `daftarfaktor`.
3. Hasil Output: Setelah semua faktor ditemukan, program menampilkan daftar faktor yang sudah terkumpul ke layar.


### 4.

![Screenshot 2024-10-31 140209](https://github.com/user-attachments/assets/7d2dc88e-5855-41e0-95ab-c9bd6a78337f)

### Source Code :
```go
package main

import (
	"fmt"
)

func cetakBarisan(bilangan int) {
	// Tampilkan bilangan saat ini
	fmt.Print(bilangan, " ")

	// Kondisi berhenti saat mencapai 1
	if bilangan == 1 {
		return
	}
	// Panggilan rekursif untuk menurunkan bilangan
	cetakBarisan(bilangan - 1)

	// Tampilkan bilangan lagi untuk membalik urutan dari 2 ke N
	fmt.Print(bilangan, " ")
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&bilangan)

	// Panggil fungsi rekursif
	cetakBarisan(bilangan)
	fmt.Println()
}
```

### Output:
![Screenshot 2024-10-29 185946](https://github.com/user-attachments/assets/81fda3ff-f09a-4a89-a9c2-58d85859d6ad)

### Full code Screenshot:
![Screenshot 2024-10-29 190021](https://github.com/user-attachments/assets/586c062d-8733-409b-8b99-3d9ddbb4b12c)

### Deskripsi Program : 
Program tersebut dibuat untuk mencetak deret angka menurun mulai dari bilangan positif yang diinputkan pengguna hingga 1, lalu mencetak angka kembali naik hingga mencapai bilangan awal. Pola ini dibentuk dengan menggunakan fungsi rekursif.

### Algoritma Program
1. Program meminta pengguna memasukkan bilangan bulat positif.
2. Fungsi `cetakBarisan(bilangan int)` dipanggil untuk membentuk deret angka:
   - `Bilangan` saat ini langsung dicetak.
   - Jika `bilangan` sama dengan 1, fungsi berhenti (kondisi dasar).
   - Jika tidak, fungsi akan memanggil dirinya sendiri dengan bilangan dikurangi 1.
   - Setelah rekursi selesai, bilangan yang dipanggil kembali dicetak sehingga membentuk pola urutan naik.
3. Fungsi ini berlanjut hingga mencapai angka 1 dan kembali ke angka awal, membentuk pola yang simetris.
### Cara Kerja Program:
1. Pengambilan Input: Program meminta pengguna untuk menginputkan sebuah bilangan bulat positif.
2. Pemanggilan Fungsi Rekursif:
   - `Bilangan` saat ini langsung dicetak, kemudian fungsi dipanggil lagi dengan nilai `bilangan` yang dikurangi 1.
   - Ketika nilai mencapai 1, rekursi selesai dan `bilangan` kembali dicetak dalam urutan naik.
3. Hasil Output: Program mencetak pola bilangan yang menurun dari angka awal ke 1, lalu berbalik arah naik hingga mencapai bilangan awal.

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
**Masukan terdiri dari sebuah bilangan bulat positif N.**

**Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.**

### Source Code :
```go
package main

import (
	"fmt"
)

func cetakBilanganGanjil(bilangan, batasAkhir int) {
	// Kondisi berhenti saat bilangan melebihi batas
	if bilangan > batasAkhir {
		return
	}

	// Tampilkan bilangan jika ganjil
	if bilangan%2 != 0 {
		fmt.Print(bilangan, " ")
	}

	// Panggilan rekursif untuk bilangan berikutnya
	cetakBilanganGanjil(bilangan+1, batasAkhir)
}

func main() {
	var batasAkhir int
	fmt.Print("Masukkan bilangan bulat positif (N): ")
	fmt.Scan(&batasAkhir)

	// Panggil fungsi rekursif dimulai dari angka 1
	cetakBilanganGanjil(1, batasAkhir)
	fmt.Println()
}

```

### Output:
![Screenshot 2024-10-29 195329](https://github.com/user-attachments/assets/711fef96-a9c2-4f94-ab67-e30fdaff3c50)

### Full code Screenshot:
![Screenshot 2024-10-31 144714](https://github.com/user-attachments/assets/5abcc8bc-d04a-42a5-ab2b-69ae1caeb80a)

### Deskripsi Program : 
Program ini menggunakan rekursi untuk mencetak semua bilangan ganjil dari angka 1 hingga batas akhir yang ditentukan oleh pengguna.

### Algoritma Program
1. Program meminta pengguna memasukkan batas akhir berupa bilangan bulat positif.
2. Fungsi `cetakBilanganGanjil(bilangan, batasAkhir int)` digunakan untuk mencetak bilangan ganjil dari 1 sampai N:
   - Fungsi berhenti jika bilangan saat ini sudah melewati batas akhir (kondisi dasar).
   - Jika bilangan saat ini ganjil (memenuhi kondisi `bilangan % 2 != 0`), maka bilangan tersebut dicetak.
   - Fungsi dipanggil kembali dengan bilangan yang ditambah 1 untuk melanjutkan ke angka berikutnya.
3. Proses ini terus berlangsung sampai seluruh bilangan ganjil dalam rentang 1 hingga batas akhir dicetak.
   
### Cara Kerja Program:
1. Pengambilan Input: Program meminta pengguna untuk memasukkan batas akhir dalam bentuk bilangan bulat positif.
2. Pemanggilan Fungsi Rekursif:
   - Fungsi `cetakBilanganGanjil` dimulai dari bilangan 1, lalu memeriksa setiap bilangan hingga mencapai batas akhir.
   - Setiap bilangan ganjil dicetak, kemudian fungsi dipanggil kembali dengan bilangan yang bertambah 1.
3. Hasil Output: Program mencetak seluruh bilangan ganjil dari 1 hingga batas akhir yang diinputkan oleh pengguna.

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

func HitungPangkat(x, y int) int {
	// Kondisi dasar: jika pangkat (y) adalah 0, hasilnya 1
	if y == 0 {
		return 1
	}
	return x * HitungPangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	// Panggil fungsi rekursif untuk menghitung pangkat
	hasilAkhir := HitungPangkat(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah: %d\n", x, y, hasilAkhir)
}

```

### Output:
![Screenshot 2024-10-29 201446](https://github.com/user-attachments/assets/fc9002fc-c186-4bbc-adf3-d8795db8733d)

### Full code Screenshot:
![Screenshot 2024-10-29 201603](https://github.com/user-attachments/assets/21d12f3c-9f89-455e-aca7-064ed38bc966)

### Deskripsi Program : 
Program tersebut menerima dua masukan dari pengguna, yaitu bilangan dasar dan pangkatnya, lalu menghitung hasil perpangkatan bilangan tersebut secara rekursif.

### Algoritma Program
1. Program meminta pengguna untuk memasukkan bilangan dasar `(x)` dan pangkat `(y)`.
2. Fungsi `HitungPangkat(x, y int)` digunakan untuk menghitung perpangkatan dengan langkah-langkah berikut:
   - Jika eksponen `(y)` bernilai 0, fungsi akan mengembalikan 1, karena bilangan apa pun yang dipangkatkan 0 menghasilkan 1 (kondisi dasar).
   - Jika tidak, fungsi akan mengembalikan hasil kali dari `x` dengan `HitungPangkat(x, y-1)`, yaitu memanggil dirinya sendiri dengan `y` yang dikurangi 1 hingga mencapai 0.
3. Fungsi `HitungPangkat` akan dipanggil terus-menerus secara rekursif sampai mencapai kondisi dasar.
   
### Cara Kerja Program:
1. Pengambilan Input: Program meminta pengguna untuk memasukkan nilai bilangan dasar `(x)` dan pangkat (y).
2. Pemanggilan Fungsi Rekursif:
   - Fungsi `HitungPangkat` bekerja dengan mengalikan x dengan hasil dari `HitungPangkat(x, y-1)`, mengurangi `y` secara bertahap hingga mencapai 0.
   - Fungsi berhenti ketika nilai `(y)` mencapai 0, kemudian mengembalikan hasil akhir.
3. Output: Program menampilkan hasil dari `x` yang dipangkatkan `y`.

## Daftar Pustaka
[1] "Pengertian Algoritma Rekursi". (n.d.). Retrieved from Universitas Janabadra: https://janabadra.ac.id/2023/algoritma-rekursi/
