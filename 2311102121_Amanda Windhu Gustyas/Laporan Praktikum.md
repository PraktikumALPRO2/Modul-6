# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 5</h2>
# <h2 align="center">REKURSIF</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI

### Pengantar Rekursif<br/>
Rekursif secara sederhana dapat diartikan sebagai metode penyelesaian masalah dengan cara menyelesaikan sub-masalah yang serupa dengan masalah utamanya. Sebagai
contoh, perhatikan prosedur cetak berikut!<br/>
![image](https://github.com/user-attachments/assets/ad6522dc-1fd2-451b-b74d-55f1945c6e12)<br/>
Jika diperhatikan, pada subprogram `cetak()` di atas, terlihat pada baris ke-4 terdapat pemanggilan kembali subprogram `cetak()`. Misalnya, jika kita menjalankan perintah `cetak(5)`, maka akan menghasilkan tampilan angka 5, 6, 7, 8, 9, dan seterusnya tanpa henti. Ini berarti setiap kali subprogram `cetak()` dipanggil, nilai x akan terus bertambah 1 (increment by one) tanpa henti.<br/>
![image](https://github.com/user-attachments/assets/ee5b2507-f3ba-464a-aac1-2bfba2df651c)<br/>
Oleh karena itu bisanya ditambahkan struktur kontrol percabangan (if-then) untuk menghentikan proses rekursif ini. Kondisi ini disebut juga dengan base-case, artinya apabila kondisi base-case bernilai true maka proses rekursif akan berhenti. Sebagai contoh misalnya base case adalah ketika x bernilai 10 atau x == 10, maka tidak perlu dilakukan rekursif.<br/>
![image](https://github.com/user-attachments/assets/a630d60e-9254-4f10-8359-3e048b1b8209)<br/>
Apabila diperhatikan pada baris ke-3 di Program di atas, kita telah menambahkan base-case seperti penjelasan sebelumnya. Selanjutnya pada bagian aksi else dari baris ke-6 dan ke-7 kita namakan recursive-case atau kasus pemanggilan dirinya sendiri tersebut terjadi. Kondisi dari recursive-case ini adalah negasi dari kondisi base-case atau ketika nilai x != 10.<br/>
![image](https://github.com/user-attachments/assets/1d56cf6f-d51e-4bc4-bad3-be86734048a0)<br/>
Apabila program di atas ini dijalankan maka akan tampil angka 5 6 7 8 9 10. Terlihat bahwa proses rekursif berhasil dihentikan ketika x == 10.<br/>
![image](https://github.com/user-attachments/assets/0d3fd753-4a9c-4ede-84a5-82b53bea7c06)<br/>
Pada Gambar 1 memperlihatkan saat subprogram dipanggil secara rekursif, maka subprogram akan terus melakukan pemanggilan (forward) hingga berhenti pada saat kondisi base-case terpenuhi atau true. Setelah itu akan terjadi proses backward atau kembali ke subprogram yang sebelumnya. Artinya setelah semua instruksi cetak(10) selesai di eksekusi, maka program akan kembali ke cetak(9) yang memanggil cetak(10) tersebut. Begitu seterusnya hingga kembali ke cetak(5).
Perhatikan modifikasi program di atas dengan menukar posisi baris 10 dan 11, mengakibatkan ketika program dijalankan maka akan menampilkan 10 9 8 7 6 5. Kenapa bisa demikian? Pahami proses backward pada Gambar 1<br/>
![image](https://github.com/user-attachments/assets/56caeef3-fd55-4ca1-af98-759f49a5b495)<br/>
Catatan:<br/>
1. Teknik rekursif ini merupakan salah satu alternatif untuk mengganti struktur kontrol perulangan dengan memanfaatkan subprogram (bisa fungsi ataupun prosedur).
2. Untuk menghentikan proses rekursif digunakan percabangan (if-then).
3. Base-case adalah kondisi proses rekursif berhenti. Base-case merupakan hal terpenting dan pertama yang harus diketahui ketika akan membuat program rekursif. Mustahil membuat program rekursif tanpa mengetahui base-case terlebih dahulu.
4. Recursive-case adalah kondisi dimana proses pemanggilan dirinya sendiri dilakukan. Kondisi recursive-case adalah komplemen atau negasi dari base-case.
5. Setiap algoritma rekursif selalu memiliki padanan dalam bentuk algoritma iteratif.

### Komponen Rekursif
Algoritma rekursif terdiri dari dua komponen utama:
1. Bae-case (Basis), yaitu bagian untuk menghentikan proses rekursif dan menjadi komponen terpenting di dalam sebuah rekursif.
2. Recursive-case, yaitu bagian pemanggilan subprogramnya

### Contoh Program dengan menggunakan Rekursif
1. Membuat baris bilangan dari n hingga 1
   Base-case: bilangan == 1<br/>
   ![image](https://github.com/user-attachments/assets/e0b66fdd-e2d6-4973-ab56-7010e254c8f5)
2. Menghitung hasil penjumlahan 1 hingga n
   Base-case: n == 1
   ![image](https://github.com/user-attachments/assets/5f1520b0-8a51-41d2-ab85-67e5d54735e7)
3. Mencari dua pangkat n atau 2^n
   Base-case: n == 0
   ![image](https://github.com/user-attachments/assets/8dd3cd40-a81c-4525-ab45-a94c359cad46)
4. Mencari nilai faktorial atau n!
   Base-case: n == 0 atau n == 1
   ![image](https://github.com/user-attachments/assets/ced93f94-7bff-46f3-837b-47a501be16de)

## II. GUIDED

### 1. Contoh Program Rekursif Menampilkan Deret Angka Menurun

```go
package main
import "fmt"
func main (){
	var n int
	fmt.Scan(&n)
	baris(n)
}

func baris(bilangan int){
	if bilangan == 1{
		fmt.Println(1)
	}else{
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
```
## Output: ![image](https://github.com/user-attachments/assets/dd416d2e-083c-4235-9280-a37a73844e2b)

Kode di atas untuk mencetak bilangan dari input yang diberikan hingga angka 1. Pertama, program mengimpor paket fmt untuk menangani fungsi input dan output. Pada fungsi utama (main), program meminta pengguna untuk memasukkan sebuah bilangan bulat yang kemudian disimpan dalam variabel n. Setelah itu, fungsi baris dipanggil dengan n sebagai argumennya. Fungsi baris bertanggung jawab untuk mencetak bilangan mulai dari n hingga 1 secara berurutan. Di dalam fungsi ini, terdapat pengecekan kondisi: jika bilangan sama dengan 1, maka program hanya mencetak angka 1, yang menjadi kondisi dasar rekursi agar berhenti. Namun, jika bilangan lebih besar dari 1, program akan mencetak angka tersebut, lalu memanggil fungsi baris secara rekursif dengan nilai bilangan - 1. Dengan cara ini, program akan terus mencetak angka-angka secara menurun hingga mencapai kondisi dasar.

### 2. Program Rekursif untuk Menghitung Penjumlahan Bilangan dari 1 hingga n

```go
package main

import "fmt"

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}
```
## Output: ![image](https://github.com/user-attachments/assets/f0afb5d2-eb61-437b-b9b3-959f57ec886c)

Kode di atas untuk menghitung jumlah bilangan dari 1 hingga bilangan n yang diberikan oleh pengguna, menggunakan pendekatan rekursif. Program diawali dengan mengimpor paket fmt untuk fungsi input dan output. Di dalam fungsi utama (main), pengguna diminta untuk memasukkan sebuah bilangan bulat yang kemudian disimpan dalam variabel n. Setelah itu, fungsi penjumlahan dipanggil dengan argumen n. Fungsi penjumlahan memiliki mekanisme rekursif untuk menghitung total penjumlahan dari 1 hingga n. Pada fungsi ini, terdapat pengecekan kondisi dasar, yaitu jika n bernilai 1, maka fungsi mengembalikan nilai 1, yang menjadi titik penghentian rekursi. Jika n lebih besar dari 1, maka fungsi akan mengembalikan hasil penjumlahan antara n dan hasil pemanggilan penjumlahan(n-1). Proses ini berlanjut hingga seluruh bilangan dari n hingga 1 dijumlahkan, dan hasil akhir penjumlahan ditampilkan melalui fungsi fmt.Println.

### 3. Program Rekursif untuk Menghitung 2 Pangkat n

```go
package main
import "fmt"

// Fungsi rekursif untuk menghitung 2^n
func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
func main(){
	var n int 
	fmt.Print("Masukkan nilai n:")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah:", pangkat(n))
}
```
## Output: ![image](https://github.com/user-attachments/assets/ff7edf13-8f21-40b6-acf3-146c1d57500d)

Kode di atas untuk menghitung hasil dari 2^n menggunakan metode rekursif. rogram dimulai dengan mengimpor paket fmt untuk mendukung fungsi input dan output. Dalam fungsi utama (main), pengguna diminta untuk memasukkan nilai n, yang merupakan eksponen untuk menghitung pangkat dua. Nilai n ini disimpan dalam variabel yang sama, dan kemudian fungsi pangkat dipanggil untuk menghitung hasilnya. Fungsi pangkat berisi logika rekursif untuk menghitung 2^n, di mana terdapat kondisi dasar yang mengembalikan nilai 1 jika n sama dengan 0 (karena 2^0 = 1). Jika n lebih besar dari 0, maka fungsi mengalikan 2 dengan hasil pemanggilan pangkat(n-1), yang secara rekursif terus mengurangi nilai n hingga mencapai kondisi dasar. Hasil akhir perhitungan pangkat ini kemudian ditampilkan kepada pengguna.

### 4. Program Rekursif untuk Menghitung Faktorial dari Bilangan n

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
func main (){

	fmt. Scan(&n)
	fmt.Println(faktorial(n))
}
```
## Output: ![image](https://github.com/user-attachments/assets/6f74ab8b-0766-4008-be74-16baaa99842a)

Kode di atas untuk menghitung nilai faktorial dari bilangan n yang diberikan oleh pengguna, menggunakan metode rekursif. Program dimulai dengan mengimpor paket fmt untuk fungsi input dan output. Pada bagian awal, terdapat variabel n yang dideklarasikan secara global. Fungsi utama (main) bertugas untuk membaca input pengguna yang dimasukkan melalui fmt.Scan(&n). Selanjutnya, fungsi faktorial dipanggil untuk menghitung faktorial dari n. Fungsi faktorial memiliki kondisi dasar yang memeriksa apakah n bernilai 0 atau 1, karena faktorial dari 0 dan 1 adalah 1. Jika n lebih besar dari 1, fungsi ini akan mengalikan n dengan hasil pemanggilan faktorial(n-1), sehingga proses rekursif akan berlanjut hingga mencapai kondisi dasar. Setelah selesai, hasil faktorial ditampilkan sebagai output.

## III. UNGUIDED

### 1. Deret Fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn-1 + Sn-2. Berikut ini adalah contoh nilai deret Fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret Fibonacci tersebut.<br/>
![image](https://github.com/user-attachments/assets/ae90637e-9f70-43a6-bdca-2440363043a7)

```go
package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var inp int
	var bilke, fibo string
	fmt.Scan(&inp)
	for i := 1; i <= inp; i++ {
		bilke += "\t" + strconv.Itoa(i)
		fibo += "\t" + strconv.Itoa(fibonacci(i))
	}
	fmt.Printf("n: %s \nSn: %s \n", bilke, fibo)
}
```
## Output: ![image](https://github.com/user-attachments/assets/f824e2c3-ea91-41c4-b275-1b03bf5ad04c)

Kode di atas untuk menghitung dan menampilkan deret bilangan Fibonacci hingga suku ke-n, menggunakan metode rekursif. Program dimulai dengan mengimpor paket fmt untuk input/output dan strconv untuk mengonversi bilangan bulat menjadi string. Fungsi fibonacci digunakan untuk menghitung nilai Fibonacci dari suku n tertentu, dengan kondisi dasar bahwa jika n bernilai 0 atau 1, fungsi mengembalikan nilai n itu sendiri. Jika n lebih besar dari 1, fungsi mengembalikan hasil penjumlahan dari fibonacci(n-1) + fibonacci(n-2), yang secara rekursif menghasilkan nilai Fibonacci yang diinginkan.

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

```go
package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

func cetakPola(n, current int) {
	if current > n {
		return
	}
	cetakBintang(current)
	fmt.Println()
	cetakPola(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	cetakPola(n, 1)
}
```
## Output: ![image](https://github.com/user-attachments/assets/02a9bd38-1f80-4977-a86d-e1d62c9a4f19)

Program ini mencetak pola segitiga bintang menggunakan fungsi rekursif. Terdapat dua fungsi utama: cetakBintang dan cetakPola. Fungsi cetakBintang(n) mencetak n bintang dalam satu baris menggunakan rekursi. Fungsi cetakPola(n, current) mencetak pola segitiga dengan memanggil cetakBintang(current) untuk setiap baris, kemudian melanjutkan ke baris berikutnya hingga mencapai n baris. Di fungsi main, program meminta pengguna memasukkan nilai n (jumlah baris), lalu mencetak pola bintang yang bertambah setiap baris hingga mencapai n bintang.

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.<br/> Masukan terdiri dari sebuah bilangan bulat positif N. <br/> Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

```go
package main

import "fmt"

func rekursif(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	rekursif(a)
}
```
## Output: ![image](https://github.com/user-attachments/assets/4f823272-bdfe-4ab5-ae18-771496b3b62f)

Program ini mencetak faktor-faktor dari suatu bilangan bulat positif yang dimasukkan oleh pengguna. Fungsi rekursif(n) menggunakan perulangan for untuk memeriksa setiap bilangan dari 1 hingga n. Jika n habis dibagi oleh i (artinya n % i == 0), maka i adalah faktor dari n dan dicetak. Di fungsi main, program meminta pengguna memasukkan nilai a, lalu memanggil rekursif(a) untuk mencetak semua faktor dari nilai tersebut.

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.<br/> Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

```go
package main

import "fmt"

func cetakUrutan(jumlah int) {
	if jumlah == 1 {
		fmt.Print(jumlah, " ")
		return
	}
	fmt.Print(jumlah, " ")
	cetakUrutan(jumlah - 1)
	fmt.Print(jumlah, " ")
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scanln(&jumlah)

	fmt.Print("Keluaran : ")
	cetakUrutan(jumlah)
	fmt.Println()
}
```
## Output: ![image](https://github.com/user-attachments/assets/d0e6f9b2-54ef-4f41-a7cd-e1a92f91ee31)

Program ini mencetak urutan angka dengan pola menurun dan menaik menggunakan rekursi. Fungsi cetakUrutan bertugas mencetak angka dari nilai input jumlah hingga 1, kemudian kembali naik ke nilai jumlah, membentuk pola simetris. Jika jumlah bernilai 1, fungsi langsung mencetak 1 dan berhenti sebagai kasus dasar (base case). Jika jumlah lebih dari 1, fungsi mencetak angka tersebut, lalu memanggil dirinya sendiri dengan nilai jumlah - 1, dan setelah selesai, mencetak kembali angka tersebut untuk pola cermin. Di fungsi main, program meminta input dari pengguna berupa angka jumlah yang menentukan batas urutan. Setelah itu, fungsi cetakUrutan dipanggil untuk mencetak urutan angka menurun dan menaik sesuai input, sehingga terbentuk pola simetris yang dimulai dari angka input, turun ke 1, dan kembali naik ke angka input.

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.<br/> Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

```go
package main

import "fmt"

func cetakBilanganGanjil(n int) {
	if n < 1 {
		return
	}
	cetakBilanganGanjil(n - 2)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	fmt.Print("Keluaran : ")
	cetakBilanganGanjil(N)
	fmt.Println()
}
```
## Output: ![image](https://github.com/user-attachments/assets/70684f2f-40f1-4e43-b6b1-5ef1ec895373)

Program ini mencetak bilangan ganjil dari 1 hingga nilai maksimum N yang diberikan oleh pengguna, menggunakan fungsi rekursif. Fungsi cetakBilanganGanjil menerima parameter n dan akan berhenti jika n kurang dari 1 (basis rekursi), karena tidak ada bilangan ganjil yang dapat dicetak. Jika n lebih besar dari atau sama dengan 1, fungsi ini pertama-tama memanggil dirinya sendiri dengan nilai n - 2, sehingga rekursi berjalan dari nilai terkecil hingga mencapai nilai n. Setelah kembali dari setiap pemanggilan rekursif, fungsi memeriksa apakah n ganjil (n % 2 != 0). Jika ya, maka nilai n dicetak, sehingga hasil akhirnya adalah urutan bilangan ganjil dari 1 hingga N secara berurutan. Di bagian main, program meminta input N dari pengguna, lalu memanggil cetakBilanganGanjil(N) untuk mencetak semua bilangan ganjil hingga batas N.
















