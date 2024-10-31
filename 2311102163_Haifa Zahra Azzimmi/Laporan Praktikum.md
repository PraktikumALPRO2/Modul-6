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
### 1. Membuat baris bilangan dari n hingga 1. Base-case: bilangan == 1
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


### Algoritma Program:

### Cara Kerja Program:


### 2. Menghitung hasil penjumlahan 1 hingga n. Base-case: n == 1
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

### Deskripsi Program

### Algoritma Program:

### Cara Kerja Program:


### 3. Mencari dua pangkat n atau 2^n. Base-case: n == 0
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

### Deskripsi Program

### Algoritma Program:


### Cara Kerja Program:


### 4. Mencari nilai faktorial atau n!. Base-case: n == 0 atau n == 1 
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

### Algoritma Program:

### Cara Kerja Program:
