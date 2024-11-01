<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VI</strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Andreas Besar Wibowo / 2311102198<br>
  S1 IF11-05
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

## I. Dasar Teori
### Definisi Rekursif
Fungsi rekursif adalah sebuah fungsi yang dapat memanggil dirinya sendiri, baik secara langsung maupun tidak langsung. Rekursi adalah teknik yang bagus untuk banyak masalah dan penting untuk memproses struktur data rekursif[1].

Recursive function adalah sebuah function yang memanggil/mengeksekusi dirinya sendiri. Recursive function bisa dikatakan salah satu yang bisa kita gunakan untuk melakukan perulangan. Ketika menulis kode aplikasi, terkadang ada kasus dimana akan lebih mudah jika dilakukan dengan recursive function. Contoh penggunaan sederhana recursive function adalah ketika melakukan operasi factorial.

Saat membuat recursive function, kita harus memastikan function tersebut dapat berhenti. Jika tidak, maka akan terkena error stack overflow atau melebihi limit stack karena function terus memanggil dirinya sendiri. Oleh karena itu, biasanya recursive function tidak langsung memanggil dirinya sendiri tetapi bergantung pada kondisi tertentu[2].

## II. GUIDED
**1. Membuat baris bilangan dari n hingga 1, Base-Case bilangan == 1**
#### Source Code
```go
package main

import "fmt"

func main() {

    // Inputan user untuk bilangan 
    var n int
    fmt.Scan(&n)
    // untuk mencetak baris
    baris(n)

}

// fungsi rekursif untuk baris
func baris(bilangan int) {
    if bilangan == 1 {
        fmt.Println(1)
    } else {
        fmt.Println(bilangan)
        baris(bilangan - 1)
    }
}
```
#### Screenshoot Source Code
![Guided1 carbon](https://github.com/user-attachments/assets/bc390140-3856-4a52-a352-9d7d245c33c7)

#### Screenshoot Output
![Guided1 go](https://github.com/user-attachments/assets/e3583961-9db0-44ea-bd9b-6909844c0fca)

#### Deskripsi Program
Program diatas merupakan program untuk mencetak bilangan bulat dan nilai yang diinputkan oleh user hingga 1 secara terbalik. Program ini menggunakan fungsi rekursif untuk mencetak bilangan 1 per 1.

Algoritma dari program tersebut sebagai berikut :
1. Minta inputan dari user untuk bilangan 'n'.
2. Panggil fungsi 'baris'.
3. fungsi 'baris', jika 'bilangan' = 1, maka cetak 1.
4. Jika tidak, maka cetak bilangan - 1.

Cara kerja dari program ini yaitu user menginputkan bilangan. Setelah itu fungsi 'baris' digunakan untuk mencetak bilangan mulai dari n hingga 1. Fungsi ini dicetak kemudian memanggil dirinya sendiri dengan nilai berkurang 1 hingga 1.

**2. Menghitung hasil penjumlahan 1 hingga n, Base-Case n == 1**
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
#### Screenshoot Source Code
![Guided2 carbon](https://github.com/user-attachments/assets/30dc10e9-e0c8-4880-ada8-79f6e9707c19)

#### Screenshoot Output
![Guided2 go](https://github.com/user-attachments/assets/f12fba86-4009-4a6b-a2f0-6f6f31002f37)

#### Deskripsi Program
Program diatas merupakan program  untuk menghitung jumlah bilangan hingga nilai yang diinputkan oleh user. Program menggunakan fungsi rekursif untuk menjumlahkan bilangan.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'penjumlahan' untuk melakukan kondisi.
2. Jika n = 1, kembalikan 1.
3. Jika tidak, kembalikan hasil 'n' dan hasil 'penjumlahan(n-1).
4. Minta input pengguna untuk bilangan.
5. Tampilkan hasil penjumlahan dengan fungsi 'penjumlahan(n).

Cara kerja dari program ini yaitu user menginputkan bilangan. Fungsi 'penjumlahan' dipanggil untuk menghitung 1 hingga nilai dari user. Fungsi ini dijalankan dengan menjumlahkan 'n' dan hasil dirinya sendiri dengan 'n-1' hingga 1.

**3. Mencari dua pangkat n atau 2^n, Base-Case n == 0**
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
#### Screenshoot Source Code
![Guided3 carbon](https://github.com/user-attachments/assets/0e0f8bf1-ee44-46c7-97c5-edcf1983cb2f)

#### Screenshoot Output
![Guided3 go](https://github.com/user-attachments/assets/565007d2-aec4-41ce-adfb-214ca4430b05)

#### Deskripsi Program
Program diatas merupakan program untuk menghitung hasil dari 2 pangkat, yang dimana 'n' merupakan bilangan yang diinputkan oleh user. Program ini menggunakan fungsi rekursif untuk menghitung pangkat.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'pangkat' untuk menghitung ^2.
2. Jika n = 0, kembalikan 1.
3. Jika tidak, kembalikan hasil 'pangkat (n-1).
4. Minta inputan dari pengguna untuk nilai.
5. Tampilkan hasil pangkat dengan fungsi 'pangkat(n)'.

Cara kerja dari program ini yaitu user menginputkan bilangan. Setelah itu, panggil fungsi 'pangkat' untuk menghitung hasil 2^n secara rekursif. Fungsi memanggil dirinya sendiri dengan nilai 'n' yang berkurang hingga 0. Fungsi akan mengalikan 2 dengan hasil pangkat sebelumnya.

**4. Mencari nilai faktorial atau n!, Base-Case n == 0 atau n == 1**
#### Source Code
```go
package main

import "fmt"

var n int
// Fungsi untuk menghitung faktorial
func faktorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main () {

	// Inputan user untuk bilangan
	fmt.Scan(&n)
	// Tampilkan nilai faktorial
	fmt.Println(faktorial(n))
}
```
#### Screenshoot Source Code
![Guided4 carbon](https://github.com/user-attachments/assets/1f6be63e-270d-4451-a556-bd7bbdffbb29)

#### Screenshoot Output
![Guided4 go](https://github.com/user-attachments/assets/ed26be5f-eeec-4ec7-9540-e2aaa2669c63)

#### Deskripsi Program
Program diatas merupakan program untuk menghitung faktorial dari bilangan yang diinputkan oleh user. Faktorial ditulis n! adalah hasil perkalian dari bilangan bulat dari 1 hingga n.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'faktorial' untuk menghitung faktorial.
2. Jika n = 0 atau 1, kembalikan 1.
3. Jika tidak, kembalikan hasil dari n! 'faktorial(n-1).
4. Minta inoutan dari user untuk bilangan.
5. Tampilkan hasil faktorial dengan fungsi 'faktorial(n)'.

Cara kerja dari program ini yaitu user menginputkan bilangan. Fungsi 'faktorial' digunakan untuk menghitung faktorial dari bilangan secara rekursif. Fungsi akan memanggil dengan nilai n yang berkurang hingga 0 atau 1. Setiap panggilan akan mengalikan n dengan hasil faktorial dari n-1.

## III. UNGUIDED
**1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S=Sn-1+Sn-2 Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonассі tersebut.**

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Inputan user untuk jumlah suku fibonacci
	var n int
	fmt.Print("	Jumlah suku Fibonacci : ")
	fmt.Scan(&n)

	// Cetak tabel deret fibonacci
	fmt.Println("\nDeret Fibonacci")
	fmt.Println("n   |   Sn")
	fmt.Println("------------")
	for a := 0; a <= n; a++ {
		fmt.Printf("%-3d | %3d\n", a, fibonacci(a))
	}
}
```
#### Screenshoot Source Code
![Unguided1 carbon](https://github.com/user-attachments/assets/f9dad49c-a291-4598-a4d7-080609d98f21)

#### Screenshoot Output
![Unguided1 go](https://github.com/user-attachments/assets/9f1c4a72-90e2-47dc-91dc-f1ce43024ac3)

#### Deskripsi Program
Program ini merupakan program untuk menghasilkan deret Fibonacci hingga suku ke-n sesuai inputan dari user. Fibonacci adalah urutan angka yang merupakan hadil penjumlahan dari 2 angka sebelumnya, dengan dua angka pertama yaitu 0 dan 1.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'fibonacci' untuk menghitung suku ke-n dengan fungsi rekursif.
2. Jika n<=1, kembalikan n.
3. Jika tidak, hitunglah 'fibonacci(n-1) + fibonacci(n-2)'.
4. Minta inputan dari user untuk jumlah suku.
5. Tampilkan hasil Fibonacci dalam bentuk tabel.

Cara kerja dari program ini yaitu meminta user untuk memasukkan jumlah suku. Lalu, dihitung secara rekurdif dan ditampilkan dalam tabel. Setiap suku dihitung dengan menjumlahkan 2 suku sebelumnya hingga mencapai jumlah suku sesuai inputan user.

**2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.**

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang
func cetakBintang(n int) {
	if n > 0 {
		fmt.Print("*")
		cetakBintang(n - 1)
	}
}

// Fungsi rekursif untuk mencetak pola bintang
func cetakPola(n, baris int) {
	if baris <= n {
		cetakBintang(baris)
		fmt.Println()
		cetakPola(n, baris+1)
	}
}

func main() {
	// Inputan user untuk baris pola bintang
	var n int
	fmt.Print("Jumlah baris pola bintang : ")
	fmt.Scan(&n)

	// Tampilkan pola bintang
	fmt.Println("\npola bintang :")
	cetakPola(n, 1)
}
```
#### Screenshoot Source Code
![Unguided2 carbon](https://github.com/user-attachments/assets/12904ccf-b83b-4ddc-84fd-d7c4a74e5084)

#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/5f5c1da4-e5a8-4788-8bde-978d615ad48d)

#### Deskripsi Program
Program ini merupakan program untuk mencetak pola bintang berdasarkan jumlah baris yang diminta oleh user. Pola yang dihasilkan berbentuk segitiga.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'cetakBintang' untuk mencetak bintang sebanyak 'n' kali.
2. Buat fungsi 'cetakPola' untuk mencetak pola bintang.
3. Jika 'baris' ≤ 'n', panggil 'cetakBintang' dengan 'baris', lalu panggil 'cetakPola' dengan 'baris + 1'.
4. Minta inputan dari user untuk jumlah baris.
5. Tampilkan pola bintang dengan 'cetakPola' mulai dari baris pertama.

Cara kerja dari program ini yaitu meminta user untuk menginputkan jumlah baris. Program menggunakan fungsi 'cetakPola' untuk mencetak bintang sesuai urutan baris. Fungsi mencetak bintang dengan 'cetakBintang' dimulai dari baris 1 hingga jumlah baris sesuai inputan user.

**3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.**                                        
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                        
*Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"


// Fungsi rekursif untuk mencetak faktor bilangan
func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	cetakFaktor(n, i+1)
}

func main() {
	// Inputan pengguna untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan faktor dari inputan user
	fmt.Printf("\nFaktor dari bilangan %d adalah: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}
```
#### Screenshoot Source Code
![Unguided3 carbon](https://github.com/user-attachments/assets/f0277246-9d55-4936-b274-1ac553b64e0c)

#### Screenshoot Output
![Unguided3 go](https://github.com/user-attachments/assets/6c8b649d-53f5-47a0-a3ca-54a470664c45)

#### Deskripsi Program
Program ini merupakan program untuk mencetak faktor dari bilangan bulat yang diinputkan oleh user. Faktor merupakan angka-angka yang bisa dibagi oleh bilangan itu sendiri.

Algoritma dari program tersebut sebagai berikut :
1. Fungsi 'cetakFaktor' untuk mencetak faktor.
2. Jika i > n, hentikan fungsinya.
3. Jika n habis dibagi i, cetak i.
4. Panggil 'cetakFaktor' dengan 'i + 1'.
5. Minta inputan dari user untuk bilangan.
6. Tampilkan faktor bilangan menggunakan fungsi 'cetakFaktor' dimulai dari 1.

Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan. Fungsi 'cetakFaktor' untuk mencetak semua faktor dari bilangan. Fungsi ini mengecek setiap angka dari 1 hingga bilangan yang diinputkan oleh user dan mencetak angka yang bisa membagi bilangan itu.

**4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.**                             
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                                 
*Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

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
#### Screenshoot Source Code
![Unguided4 carbon](https://github.com/user-attachments/assets/8c23c9ce-8a88-4280-9a64-4f9fde2e4493)

#### Screenshoot Output
![Unguided4 go](https://github.com/user-attachments/assets/77116484-4d77-4be1-b568-53e61ae8218c)

#### Deskripsi Program
Program ini merupakan program untuk mencetak barisan bilangan dari suatu bilangan hingga 1, kemudian kembali lagi ke bilangan itu. 

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'cetakBarisan' untuk mencetak barisan.
2. Jika nilai = 1, hentikan fungsi.
3. Panggil 'cetakBarisan' dengan 'nilai - 1'.
4. Minta inputan dari user untuk bilangan.
5. Tampilkan barusan bilangan dari 'n' hingga 1, lalu kembali ke 'n'.

Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan. Fungsi 'cetakBarisan' untuk mencetak nilai dari bilangan yang diinputkan hingga 1, lalu kembali lagi ke bilangan. Fungsi ini digunakan untuk menuruskan nilai mencapai 1, lalu kembali lagi ke bilangan tersebut.

**5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.**                                 
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                        
*Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil
func cetakGanjil(n, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}

	cetakGanjil(n, i+1)
}

func main() {
	// Inputan user untuk bilangan
	var n int
	fmt.Print("Bilangan (Bulat Positif): ")
	fmt.Scan(&n)

	// Tampilkan barisan bilangan ganjil
	fmt.Printf("\nBarisan bilangan ganjil dari 1 hingga %d adalah: ", n)
	cetakGanjil(n, 1)
	fmt.Println()
}
```
#### Screenshoot Source Code
![Unguided5 carbon](https://github.com/user-attachments/assets/bd2cb6b1-f93c-4ad2-b476-55e5643b4320)

#### Screenshoot Output
![Unguided5 go](https://github.com/user-attachments/assets/04eb3deb-6157-4af4-90b7-d9a767b25b2f)

#### Deskripsi Program
Program ini merupakan program untuk mencetak semua bilangan ganjil dari 1 hingga bilangan yang diinputkan oleh user.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'cetakGanjil' untuk mencetak bilangan ganjil.
2. Jika i > n, hentikan fungsi.
3. Jika i  merupakan bilangan ganjil, cetak 'i'.
4. Panggil 'cetakGanjil' dengan 'i + 1'.
5. Minta inputan user untuk bilangan.
6. Tampilkan barisan bilangan ganjil dari 1 hingga 'n'.

Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan. Fungsi 'cetakGanjil' digunakan untuk mencetak semua bilangan ganjil dari 1 hingga bilangan yang sesuai dengan inputan user. Fungsi akan mengecek dengan cara rekursif, mencetak angka jika bilangan ganjil. Lalu, melanjutkan hingga bilangan yang diminta.

**6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.**                                     
*Masukan terdiri dari bilangan bulat x dan y.*                                        
*Keluaran terdiri dari hasil x dipangkatkan y.*                                             
#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF-11-05

package main

import "fmt"

// Fungsi rekursif untuk pangkat
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}

	return x * pangkat(x, y-1)
}

func main() {
	// Inputan user untuk bilangan dan bilangan pangkat
	var x, y int
	fmt.Print("Bilangan : ")
	fmt.Scan(&x)
	fmt.Print("Bilangan pangkat : ")
	fmt.Scan(&y)

	// Tampilkan hasil pangkatnya
	hasil := pangkat(x, y)
	fmt.Printf("\nHasil %d dipangkatkan %d adalah: %d\n", x, y, hasil)
}
```
#### Screenshoot Source Code
![Unguided6 carbon](https://github.com/user-attachments/assets/b78d17bb-ab07-457e-8858-2a1b5212e5bb)

#### Screenshoot Output
![Unguided6 go](https://github.com/user-attachments/assets/36c11a92-4052-4b05-b216-14f4804c5570)

#### Deskripsi Program
Program ini merupakan program untuk menghitung hasil pangkat dari suatu bilangan dari inputan user. Pengguna bisa menginputkan bilangan pangkat yang diinginkan. Program ajan menghitung hasilnya secara rekursif.

Algoritma dari program tersebut sebagai berikut :
1. Buat fungsi 'pangkat' untuk menghitung pangkat.
2. Jika 'y' = 0, kembalikan 1.
3. Jika tidak, kembalikan hasil perkalian 'x' dengan 'pangkat(x, y - 1)'.
4. Minta inputan dari pengguna untuk bilangan dan bilangan pangkat.
5. Hitung hasil pangkat dengan fungsi 'pangkat'.
6. Tampilkan hasilnya.

Cara kerja dari program ini yaitu meminta user untuk menginputkan bilangan dan bilangan pangkat. Fungsi 'pangkat' digunakan untuk menghitung hasil pangkat secara rekursif. Fungsi akan memanggil dirinya sendiri dengan nilai pangkat yang berkurang sampai 0, dimana hasilnya 1.

## Referensi
[1] Ahmad Faisal. (2021). Dasar-Dasar Bahasa Pemrograman Golang. Diakses dari https://pta.pilkommedia.org/progress/upload/AhmadFaisal_A1C615001_Dasar-DasarBahasaPemrogramanGolang.pdf                          
[2]Ruang Developer. (n.d.). Recursive function - Belajar Golang dari dasar. Diakses dari https://blog.ruangdeveloper.com/golang-recursive-function/

