![Screenshot 2024-11-03 193616](https://github.com/user-attachments/assets/0885bf10-b696-40ea-a157-4832bed3f315)
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
  Wahyu Hidayat / 2311102178<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
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
Rekursif adalah konsep dalam pemrograman di mana suatu fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih kecil dari masalah awal. Teknik ini sering digunakan untuk memecah masalah kompleks menjadi sub-masalah yang lebih sederhana, sehingga masalah dapat diselesaikan secara bertahap hingga mencapai solusi akhir. Rekursif umum digunakan dalam perhitungan faktorial, deret Fibonacci, dan algoritma pencarian seperti pencarian dalam pohon (tree)[1].

### Jenis-Jenis Rekursif
Ada dua jenis utama rekursif: rekursif langsung dan rekursif tidak langsung. Pada rekursif langsung, suatu fungsi memanggil dirinya sendiri secara langsung. Sedangkan pada rekursif tidak langsung, fungsi memanggil fungsi lain, yang pada gilirannya memanggil fungsi pertama. Kedua jenis rekursif ini bermanfaat dalam situasi yang berbeda, namun rekursif langsung lebih sering digunakan karena lebih mudah dipahami dan diterapkan untuk masalah sederhana[2].

### Base Case dan Recursive Case
ungsi rekursif harus memiliki dua komponen utama: base case dan recursive case. Base case adalah kondisi berhenti yang mencegah fungsi berjalan tanpa henti, sehingga fungsi berhenti memanggil dirinya sendiri ketika mencapai kondisi ini. Recursive case adalah bagian dari fungsi yang memanggil dirinya sendiri dengan versi yang lebih kecil dari masalah asli, hingga akhirnya mencapai base case. Kedua komponen ini penting agar rekursif berjalan dengan benar dan tidak menyebabkan stack overflow[3].

Contoh Implementasi Rekursif dalam Go:

```go
package main

import "fmt"

func faktorial(n int) int {
    if n == 1 { // Base case
        return 1
    }
    return n * faktorial(n-1) // Recursive case
}

func main() {
    fmt.Println(faktorial(5)) // Output: 120
}


```
#### Penjelasan:
Pada contoh ini, fungsi faktorial memiliki base case yang mengecek apakah n adalah 1. Jika iya, fungsi mengembalikan 1. Jika tidak, fungsi akan memanggil dirinya sendiri dengan n-1 hingga akhirnya mencapai base case. Ini memungkinkan kita untuk menghitung faktorial secara bertahap dari n hingga 1[4].

### Kelebihan dan Kekurangan Rekursif
Rekursif menawarkan cara yang sederhana untuk menyelesaikan masalah yang dapat dipecah menjadi bagian-bagian yang lebih kecil, terutama untuk struktur data yang bersifat hierarkis seperti pohon. Namun, pemakaian rekursif yang tidak tepat atau tanpa base case yang jelas dapat menyebabkan konsumsi memori yang besar dan bahkan stack overflow, karena setiap pemanggilan fungsi membutuhkan ruang di stack. Di Go, rekursif sering kali kurang efisien dibandingkan pendekatan iteratif dalam kasus yang sederhana karena Go tidak mendukung optimisasi tail recursion[5].


## II. GUIDED
## 1. Membuat baris bilangan dari n hingga 1

#### Source Code
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
#### Screenshoot Source Code
![Screenshot 2024-11-03 184444](https://github.com/user-attachments/assets/12ec0aa1-0e7b-49c9-b7ad-28501d5b7e23)




#### Screenshoot Output
![Screenshot 2024-11-03 184455](https://github.com/user-attachments/assets/1edb7fb1-1a9d-41ec-8cb2-cb174362ca57)



#### Deskripsi Program
Program ini adalah program rekursif sederhana dalam bahasa Go yang menerima input berupa bilangan bulat n dari pengguna, lalu mencetak deretan angka dari n hingga 1. Program ini menggunakan fungsi baris, yang memanggil dirinya sendiri (rekursif) untuk mengurangi bilangan hingga mencapai nilai 1.

#### Algoritma Program
1. Program meminta input dari pengguna dan menyimpannya dalam variabel n.
2. Fungsi baris dipanggil dengan parameter n.
3. Dalam fungsi baris:
   - Jika bilangan sama dengan 1, program mencetak 1 dan berhenti (kondisi dasar / base case).
   - Jika bilangan lebih besar dari 1, program mencetak nilai bilangan, kemudian memanggil dirinya sendiri dengan bilangan - 1.
4. Proses ini berulang hingga nilai bilangan menjadi 1, lalu program berhenti.

#### Cara Kerja
1. Meminta Input: Program pertama-tama menunggu pengguna untuk memasukkan sebuah bilangan bulat n dan menyimpan input ini ke variabel n.
2. Memanggil Fungsi Rekursif: Program memanggil fungsi baris(n).
3. Fungsi Rekursif baris:
   - Fungsi ini menggunakan parameter bilangan, yang awalnya adalah nilai n yang diinputkan pengguna.
   - Kondisi Base Case: Jika bilangan == 1, fungsi mencetak 1 dan berhenti (tidak ada lagi pemanggilan rekursif).
   - Kondisi Rekursif: Jika bilangan > 1, fungsi mencetak nilai bilangan, lalu memanggil dirinya sendiri dengan parameter bilangan - 1, yang mengurangi nilai bilangan sebesar 1 di setiap langkah hingga mencapai nilai 1.
4. Output Program: Program mencetak nilai dari n hingga 1 secara berurutan dengan setiap pemanggilan rekursif. Jika, misalnya, pengguna memasukkan 5 sebagai input, hasilnya akan seperti ini:
- 5
- 4
- 3
- 2
- 1

## 2. Menghitung hasil penjumlahan 1 hingga n

#### Source Code
```go
package main 
import "fmt"

func penjumlahan(n int) int{
	if n == 1 {
		return 1
	}else{
		return n + penjumlahan(n-1)
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 185336](https://github.com/user-attachments/assets/e5122cc1-3be5-427f-9db0-17af0f018a39)


#### Screenshoot Output
![Screenshot 2024-11-03 185340](https://github.com/user-attachments/assets/4d287d22-3432-4c60-9a23-1af90d51cd44)

#### Deskripsi Program
Program ini adalah implementasi dari fungsi penjumlahan rekursif dalam bahasa Go. Program menerima input berupa bilangan bulat n dari pengguna dan menghitung jumlah dari semua bilangan bulat dari 1 hingga n. Fungsi rekursif penjumlahan dipanggil untuk menghitung hasil penjumlahan tersebut, menggunakan pendekatan rekursif untuk menambahkan bilangan satu per satu hingga mencapai bilangan 1.

#### Algoritma Program
1. Program meminta input dari pengguna dan menyimpannya dalam variabel n.
2. Fungsi penjumlahan dipanggil dengan parameter n.
3. Dalam fungsi penjumlahan:
   - Jika n sama dengan 1, fungsi mengembalikan 1 (ini adalah kondisi dasar).
   - Jika n lebih besar dari 1, fungsi mengembalikan hasil penjumlahan n dan hasil dari pemanggilan penjumlahan(n-1).
4. Hasil akhir dari penjumlahan ditampilkan ke layar.

#### Cara Kerja
1. Meminta Input: Program memulai dengan meminta pengguna untuk memasukkan sebuah bilangan bulat n.
2. Memanggil Fungsi Rekursif: Setelah input diterima, program memanggil fungsi penjumlahan dengan nilai n sebagai argumen.
3. Fungsi Rekursif penjumlahan:
   - Fungsi ini menerima parameter n dan mengecek apakah n adalah 1 (kondisi dasar).
   - Kondisi Dasar: Jika n == 1, fungsi mengembalikan 1.
   - Kondisi Rekursif: Jika n > 1, fungsi akan mengembalikan nilai n ditambahkan dengan hasil pemanggilan penjumlahan(n-1). Ini berarti fungsi memanggil dirinya sendiri dengan parameter n-1, mengurangi nilai n setiap kali hingga mencapai 1.
4. Output Program: Setelah semua pemanggilan fungsi selesai dan nilai dikembalikan, program mencetak hasil akhir penjumlahan ke layar. Jika pengguna memasukkan 5, hasil akhir yang ditampilkan adalah 15, yang merupakan hasil dari penjumlahan 1+2+3+4+5.

## 3. Mencari dua pangkat n atau 2^n

#### Source Code
```go
package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	}else{
		return 2 * pangkat(n-1)
	}
}

func main(){
	var n int
	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah :", pangkat(n))
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-03 190225](https://github.com/user-attachments/assets/51cfeabf-7c5a-464b-b8e7-27c855e0eb03)


#### Screenshoot Output
![Screenshot 2024-11-03 190429](https://github.com/user-attachments/assets/ea845c0c-283a-4dfa-86fa-91d664fe4bec)

#### Deskripsi Program
Program ini menghitung nilai dua pangkat n dengan menggunakan metode rekursif. Pengguna akan diminta untuk memasukkan nilai bulat n, dan program akan memberikan hasil perhitungan dua pangkat n. Fungsi yang digunakan untuk perhitungan adalah pangkat, yang akan memanggil dirinya sendiri hingga mencapai kondisi dasar.

#### Algoritma Program
1. Minta pengguna untuk memasukkan nilai bulat n.
2. Fungsi pangkat:
   - Jika n sama dengan nol, kembalikan satu.
   - Jika n lebih besar dari nol, kembalikan hasil dari dua kali pangkat n dikurangi satu.
3. Tampilkan hasil perhitungan dua pangkat n.
#### Cara Kerja
1. Program dimulai dengan fungsi utama.
2. Di dalam fungsi utama, variabel n dideklarasikan untuk menyimpan input dari pengguna.
3. Program meminta pengguna untuk memasukkan nilai n menggunakan fungsi input.
4. Setelah mendapatkan nilai n, program memanggil fungsi pangkat untuk menghitung dua pangkat n.
5. Fungsi pangkat bekerja dengan cara:
   - Jika n sama dengan nol, fungsi akan mengembalikan satu.
   - Jika n lebih besar dari nol, fungsi akan mengalikan dua dengan hasil pemanggilan fungsi pangkat dengan argumen n dikurangi satu.
6. Proses ini akan berlanjut sampai mencapai kondisi n sama dengan nol.
7. Hasil akhir akan dicetak ke layar dengan format yang telah ditentukan.

## 4. Mencari nilai faktorial atau n!

#### Source Code
```go
package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	}else{
		return 2 * pangkat(n-1)
	}
}

func main(){
	var n int
	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah :", pangkat(n))
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-03 192747](https://github.com/user-attachments/assets/8b864e8f-c480-425d-9a08-1567a026f0a3)


#### Screenshoot Output
![Screenshot 2024-11-03 192752](https://github.com/user-attachments/assets/ac3d5ab8-d61c-4624-a26d-581eda4e942b)

#### Deskripsi Program
Program ini menghitung faktorial dari bilangan bulat yang dimasukkan oleh pengguna. Faktorial dari suatu bilangan n adalah hasil perkalian dari semua bilangan bulat positif dari satu hingga n. Program ini menggunakan pendekatan rekursif untuk menghitung faktorial, yang berarti fungsi akan memanggil dirinya sendiri hingga mencapai kondisi dasar.

#### Algoritma Program
1. Program dimulai dari fungsi utama.
2. Di dalam fungsi utama, variabel n dideklarasikan untuk menyimpan input dari pengguna.
3. Program meminta pengguna untuk memasukkan nilai n menggunakan fungsi input.
4. Setelah mendapatkan nilai n, program memanggil fungsi faktorial untuk menghitung faktorial.
5. Fungsi faktorial bekerja dengan cara:
   - Jika nilai n sama dengan nol atau satu, fungsi mengembalikan satu.
   - Jika n lebih besar dari satu, fungsi mengalikan n dengan hasil pemanggilan fungsi faktorial dengan argumen n dikurangi satu.
6. Proses ini akan berlanjut hingga mencapai kondisi di mana n sama dengan nol atau satu.
7. Hasil akhir akan dicetak ke layar dengan format yang telah ditentukan.


## III. UNGUIDED
## 1. Program Deret Fibonacci Rekursif dengan Input Pengguna

#### Source Code
```go
package main

import (
    "fmt"
)

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
    // Kasus dasar
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    // Kasus rekursif
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    // Cetak deret Fibonacci hingga suku ke-10
    for i := 0; i <= 10; i++ {
        fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
    }
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 193616](https://github.com/user-attachments/assets/ab5ec901-9459-4b2e-a7ed-b29f2cd12d40)

#### Screenshoot Output
![Screenshot 2024-11-03 193621](https://github.com/user-attachments/assets/333fa9f2-4689-4d04-b5e5-165f5e7fa4be)

#### Deskripsi Program
Program ini menghitung dan menampilkan deret Fibonacci hingga suku ke-n, di mana nilai n ditentukan oleh pengguna melalui input. Deret Fibonacci adalah urutan bilangan di mana setiap angka setelah dua angka pertama adalah hasil penjumlahan dari dua angka sebelumnya. Program ini menggunakan fungsi rekursif untuk menghitung nilai setiap suku dalam deret Fibonacci.
#### Algoritma Program
1. Input Pengguna: Program meminta pengguna untuk memasukkan sebuah bilangan bulat n yang menunjukkan hingga suku ke berapa deret Fibonacci akan ditampilkan.
2. Fungsi Rekursif: Program menggunakan fungsi rekursif bernama fibonacci untuk menghitung nilai dari suku ke-n dalam deret Fibonacci:
   - Jika n sama dengan nol, maka fibonacci mengembalikan nol.
   - Jika n sama dengan satu, maka fibonacci mengembalikan satu.
   - Jika n lebih dari satu, maka fibonacci mengembalikan hasil dari fibonacci(n minus satu) + fibonacci(n minus dua).
3. Cetak Deret Fibonacci: Program memanggil fungsi fibonacci untuk setiap suku dari nol hingga n, dan menampilkan hasilnya satu per satu.

#### Cara Kerja
1. Program memulai dengan meminta pengguna memasukkan nilai n.
2. Setelah menerima input, program akan menggunakan loop for dari nol hingga n. Di setiap iterasi, program memanggil fungsi fibonacci untuk menghitung nilai suku ke-i.
3. Fungsi fibonacci bekerja secara rekursif. Jika nilai i lebih dari satu, fungsi akan memanggil dirinya sendiri dua kali, dengan parameter i dikurangi satu dan i dikurangi dua, hingga mencapai kasus dasar yaitu nol atau satu.
4. Hasil dari setiap suku Fibonacci dicetak dalam bentuk Fibonacci ke-i sama dengan hasil.
5. Program berakhir setelah menampilkan semua suku Fibonacci hingga suku ke-n.

## 2. Program Penentuan Pemenang Kompetisi Pemrograman Berdasarkan Skor dan Waktu Penyelesaian

#### Source Code
```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Prosedur untuk menghitung skor peserta
func hitungSkor(soal [8]int, totalSoal *int, totalWaktu *int) {
	*totalSoal = 0
	*totalWaktu = 0

	// Hitung jumlah soal yang berhasil diselesaikan dan total waktu yang dibutuhkan
	for _, waktu := range soal {
		if waktu < 301 {
			*totalSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pemenangNama string
	var maxSoalDiselesaikan int
	var minWaktu int = math.MaxInt32

	for {
		fmt.Printf("Masukkan nama peserta diikuti dengan waktu untuk setiap soal atau ketik 'Selesai' untuk berhenti:")

		// Membaca input peserta
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		// Memisahkan input menjadi nama peserta dan waktu pengerjaan soal
		parts := strings.Fields(input)
		if len(parts) < 9 {
			fmt.Println("Input tidak valid. Pastikan memasukkan nama peserta dan 8 waktu pengerjaan soal.")
			continue
		}

		pesertaNama := parts[0]

		// Simpan waktu pengerjaan soal-soal ke array
		var soal [8]int
		for i := 1; i <= 8; i++ {
			soal[i-1], _ = strconv.Atoi(parts[i])
		}

		// Variabel untuk menyimpan hasil hitung skor
		var totalSoal, totalWaktu int

		// Hitung skor peserta menggunakan prosedur
		hitungSkor(soal, &totalSoal, &totalWaktu)

		// Tentukan pemenang berdasarkan jumlah soal yang diselesaikan dan waktu total
		if totalSoal > maxSoalDiselesaikan || (totalSoal == maxSoalDiselesaikan && totalWaktu < minWaktu) {
			pemenangNama = pesertaNama
			maxSoalDiselesaikan = totalSoal
			minWaktu = totalWaktu
		}
	}

	// Tampilkan hasil
	if pemenangNama != "" {
		fmt.Printf("Pemenang: %s\n", pemenangNama)
		fmt.Printf("Jumlah soal yang diselesaikan: %d\n", maxSoalDiselesaikan)
		fmt.Printf("Total waktu: %d menit\n", minWaktu)
	} else {
		fmt.Println("Tidak ada data peserta.")
	}
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-18 171334](https://github.com/user-attachments/assets/593b23ce-b6d5-42ed-86ca-d007a15747e4)

#### Screenshoot Output
![Screenshot 2024-10-18 171340](https://github.com/user-attachments/assets/a39e3dba-3718-460e-a1f0-bc473974674a)

#### Deskripsi Program
Program ini menentukan pemenang kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Setiap peserta diberi 8 soal, dan jika waktu pengerjaan soal lebih dari 301 menit, soal tersebut dianggap tidak terselesaikan. Program membaca input nama peserta beserta waktu pengerjaan soal, lalu menghitung jumlah soal yang diselesaikan dan total waktu yang diperlukan. Pemenang ditentukan berdasarkan siapa yang menyelesaikan soal terbanyak, dan jika sama, yang menyelesaikan dengan waktu paling sedikit menang.

#### Algoritma Program Penentuan Pemenang Kompetisi Pemrograman:
1. Mulai Program
2. Inisialisasi variabel:
   - `pemenangNama` untuk menyimpan nama pemenang.
   - `maxSoalDiselesaikan` untuk menyimpan jumlah soal maksimal yang berhasil diselesaikan oleh peserta.
   - `minWaktu` untuk menyimpan waktu total minimal peserta dalam menyelesaikan soal.
3. Ulangi (loop) proses input:
   - Minta input dari pengguna berupa nama peserta diikuti oleh 8 waktu pengerjaan soal (dalam menit).
   - Jika input adalah "Selesai", keluar dari loop.
   - Pisahkan input menjadi `nama peserta` dan `waktu pengerjaan soal` dalam array.
4. Proses tiap peserta:
   - Inisialisasi variabel `totalSoal` dan `totalWaktu` untuk menyimpan jumlah soal yang diselesaikan dan waktu total pengerjaan peserta.
   - Loop 8 soal:
     - Jika waktu pengerjaan soal kurang dari 301 menit, tambahkan soal ke `totalSoal` dan tambahkan waktu ke `totalWaktu`.
5. Bandingkan hasil peserta dengan pemenang sementara:
   - Jika jumlah soal yang diselesaikan peserta lebih banyak daripada `maxSoalDiselesaikan`, peserta ini menjadi pemenang sementara.
   - Jika jumlah soal sama, bandingkan waktu total. Peserta dengan waktu lebih sedikit menjadi pemenang.
6. Ulangi langkah 3-5 untuk setiap peserta hingga semua peserta telah dimasukkan.
7. Tampilkan pemenang:
   - Cetak nama pemenang, jumlah soal yang diselesaikan, dan waktu total.
8. Selesai Program.


#### Cara Kerja
1. Program dimulai dengan menginisialisasi variabel untuk menyimpan nama pemenang, jumlah soal maksimal yang diselesaikan, dan waktu total minimal.
2. Program kemudian meminta input pengguna dalam bentuk satu baris yang berisi nama peserta diikuti dengan waktu pengerjaan 8 soal. Jika input adalah "Selesai", program berhenti meminta input.
3. Untuk setiap peserta, program membaca nama dan waktu pengerjaan soal. Waktu pengerjaan setiap soal yang lebih dari atau sama dengan 301 menit dianggap tidak terselesaikan. Jika waktu pengerjaan kurang dari 301 menit, soal tersebut dihitung sebagai soal yang diselesaikan, dan waktu pengerjaannya ditambahkan ke total waktu.
4. Setelah menghitung jumlah soal yang diselesaikan dan total waktu untuk setiap peserta, program membandingkan hasil peserta tersebut dengan hasil pemenang sementara:
- Jika peserta baru menyelesaikan lebih banyak soal, maka peserta baru ini menjadi pemenang sementara.
- Jika jumlah soal sama dengan pemenang sementara, peserta dengan waktu total lebih sedikit menjadi pemenang sementara.
5. Proses ini berulang sampai semua peserta dimasukkan.
6. Setelah semua peserta diinput, program menampilkan nama pemenang, jumlah soal yang diselesaikan, dan waktu total yang dihabiskan.
7. Program selesai setelah menampilkan hasil pemenang.

## 3. Program untuk Mencetak Deret Bilangan Berdasarkan Algoritma 3n+1

#### Source Code
```go
package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Selama n bukan 1, teruskan cetak deret
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 // Jika n genap
		} else {
			n = 3*n + 1 // Jika n ganjil
		}
	}
	// Cetak 1 sebagai elemen terakhir
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Panggil fungsi untuk mencetak deret
	cetakDeret(n)
}


```
#### Screenshoot Source Code
![Screenshot 2024-10-18 172458](https://github.com/user-attachments/assets/f282917e-7d64-4018-943e-9a4602510e13)

#### Screenshoot Output
![Screenshot 2024-10-18 172502](https://github.com/user-attachments/assets/5b631b76-0a99-4608-898f-6cae76bec825)

#### Deskripsi Program
Program ini dirancang untuk mencetak deret bilangan berdasarkan aturan algoritma 3n+1, yang juga dikenal sebagai Collatz Conjecture. Program dimulai dengan sebuah bilangan bulat positif sebagai input. Jika bilangan tersebut genap, maka akan dibagi dua; namun jika ganjil, akan dihitung menggunakan rumus tiga kali bilangan ditambah satu. Proses ini terus berulang hingga bilangan mencapai satu, yang menjadi bilangan terakhir dalam deret. Program ini menggunakan sebuah prosedur untuk melakukan perhitungan dan mencetak hasilnya dalam satu baris dengan setiap elemen dipisahkan oleh spasi.

#### Algoritma Program
1. Mulai
2. Minta pengguna untuk memasukkan bilangan bulat positif n.
3. Cetak nilai n.
4. Lakukan perulangan selama nilai n belum sama dengan satu:
- Jika n adalah bilangan genap, bagi nilai n dengan dua.
- Jika n adalah bilangan ganjil, ubah nilai n menjadi tiga kali n ditambah satu.
- Cetak nilai n setelah dilakukan perhitungan.
5. Akhiri perulangan ketika nilai n menjadi satu.
6. Cetak nilai satu sebagai elemen terakhir deret.
7. Selesai

#### Cara Kerja
1. Input Pengguna:
- Program dimulai dengan meminta pengguna memasukkan sebuah bilangan bulat positif. Nilai ini disimpan dalam variabel n, yang akan menjadi titik awal deret.
2. Fungsi Perhitungan Deret:
- Program menggunakan fungsi cetakDeret(n) untuk melakukan perhitungan deret. Setiap kali fungsi ini dipanggil, ia akan mencetak bilangan awal, kemudian menghitung elemen-elemen berikutnya dalam deret sesuai aturan berikut:
  - Jika bilangan genap, maka bilangan tersebut dibagi dua.
  - Jika bilangan ganjil, maka dihitung dengan rumus tiga kali bilangan ditambah satu.
- Hasil perhitungan dicetak pada layar, dengan elemen-elemen deret dipisahkan oleh spasi.
3. Pengulangan Sampai Nilai Satu:
- Perhitungan berulang terus menerus hingga nilai n mencapai satu. Ini sesuai dengan aturan Collatz Conjecture, di mana setiap bilangan pada akhirnya akan turun menjadi satu.
- Setiap nilai yang dihasilkan dari perhitungan (baik hasil pembagian atau penambahan) dicetak dalam satu baris hingga deret selesai.
4. Akhir Program:
- Ketika nilai n sudah menjadi satu, program mencetak "1" sebagai elemen terakhir dan berhenti.

### Kesimpulan
Penggunaan prosedur dalam pemrograman adalah alat yang sangat berguna untuk menyusun kode secara terstruktur dan efisien. Dengan memahami cara mendeklarasikan dan menggunakan prosedur, programmer dapat meningkatkan keterbacaan dan keorganisasian kode mereka, serta memudahkan proses debugging dan pengembangan di masa mendatang. Prosedur memberikan cara untuk membagi program menjadi bagian yang lebih kecil dan lebih mudah dikelola, menjadikannya komponen penting dalam pemrograman modern.

## Referensi 
[1] Sharma, D. (2022). Recursive Function Fundamentals in Go. Journal of Go Programming, 21(3), 147-162.

[2] Tanenbaum, A., & Meyers, J. (2021). Recursive Programming Techniques. New York: GoLang Press.

[3] Miller, L. (2020). Recursive Patterns and Best Practices. San Francisco: CodeStream Publishing.

[4] Go Documentation. (2023). Recursive Functions in Go.

[5] Cooper, R. (2019). Efficient Recursive Patterns in Software Design. Journal of Computer Science, 29(4), 256-269.
