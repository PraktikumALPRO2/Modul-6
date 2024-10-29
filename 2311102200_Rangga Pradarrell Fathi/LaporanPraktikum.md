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
 Rangga Pradarrell Fathi
  <br>
  2311102200
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
Rekursif merupakan konsep pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan suatu masalah. Fungsi rekursif akan terus memanggil dirinya sendiri hingga mencapai suatu kondisi dasar (base case) yang memungkinkan fungsi tersebut berhenti dan mengembalikan nilai.

### <strong> Beberapa karakteristik utama dari rekursif:

- Base Case: Setiap fungsi rekursif harus memiliki kondisi dasar atau base case yang menghentikan rekursi. Jika tidak ada base case, rekursi akan terus berlanjut tanpa batas.
- Pengurangan Masalah: Setiap pemanggilan rekursif harus mengurangi masalah menjadi sub-masalah yang lebih kecil dan mendekati base case.
- Pemanggilan Diri Sendiri: Fungsi rekursif memanggil dirinya sendiri dengan argumen yang berbeda untuk menyelesaikan sub-masalah.
- Kemampuan Memecah Masalah: Rekursif sangat efektif untuk memecah masalah yang kompleks menjadi sub-masalah yang lebih sederhana dan dapat diselesaikan secara terpisah.

### <strong> Contoh penerapan rekursif dalam pemrograman antara lain:
- Menghitung Faktorial: Fungsi rekursif untuk menghitung faktorial suatu bilangan.
   **Faktorial**:
   ```go
   func factorial(n int) int {
       if n == 0 {
           return 1
       }
       return n * factorial(n-1)
   }
   ```
   Fungsi `factorial()` memanggil dirinya sendiri dengan nilai `n-1` sampai `n` mencapai 0, lalu mengembalikan hasil perkalian dari `n` sampai 1.
  
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
//Nama: Rangga Pradarrell Fathi
//NIM : 2311102200 Kelas: IF-11-05
package main

import "fmt"

func main() {
    var a, b int // mendeklarasikan 2 variabel 
    fmt.Scan(&a, &b)  // untuk input dari pengguna
    if a >= b { // jika a >= b, akan ada permutasi
        permutasi(a, b) // memanggil prosedur permutasi
    } else { // jika b >= a, akan ada permutasi
        permutasi(b, a) // memanggil prosedur permutasi
    }
}

// prosedur untuk menghitung faktorial 
func faktorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

// prosedur untuk menghitung dan mencetak permutasi
func permutasi(n, r int) {
    var faktorialN, faktorialNR int

    // menghitung faktorial n
    faktorial(n, &faktorialN)

    // menghitung faktorial (n-r)
    faktorial(n-r, &faktorialNR)

    // menghitung permutasi 
    hasil := faktorialN / faktorialNR
    fmt.Println(hasil)
}
```

### Screenshot Code
![Screenshot 2024-10-19 091209](https://github.com/user-attachments/assets/e9b6c33c-63ba-418d-84d6-a338fa5a7247)


### Deskripsi Program


**Algoritma Program**
Inisialisasi variabel: Program mendeklarasikan dua variabel a dan b untuk menyimpan bilangan bulat yang dimasukkan oleh pengguna.
Input: Program meminta pengguna memasukkan dua angka dengan memanggil prosedur bacaInput.
Perbandingan: Program membandingkan nilai a dan b. Jika a >= b, maka program akan menghitung permutasi dari a sebagai n dan b sebagai r. Sebaliknya, jika b > a, program menghitung permutasi dari b sebagai n dan a sebagai r.
Perhitungan Faktorial: Program menghitung faktorial dari n dan n - r menggunakan prosedur faktorial.
Perhitungan Permutasi: Program menghitung permutasi dengan memanggil prosedur hitungPermutasi dan menampilkan hasilnya.
Output: Program mencetak hasil perhitungan permutasi ke layar.

**Cara Kerja Program:**
Program dimulai dengan memanggil prosedur bacaInput untuk membaca input pengguna (dua bilangan a dan b).
Program kemudian menentukan apakah a >= b atau sebaliknya melalui prosedur permutasi.
Setelah itu, program memanggil hitungPermutasi untuk menghitung permutasi dan mencetak hasilnya ke layar.

## Guided - 2
### Study Case
Membuat program dengan bahasa go untuk mencari sebuah Luas dan Keliling Persegi.

### Source Code
```
//Nama: Rangga Pradarrell Fathi
//NIM : 2311102200 Kelas: IF-11-05
package main

import (
	"fmt"
)

// prosedur untuk menghitung luas persegi
func hitungLuas(sisi float64, luas *float64) {
	*luas = sisi * sisi
}

// prosedur untuk menghitung keliling persegi
func hitungKeliling(sisi float64, keliling *float64) {
	*keliling = 4 * sisi
}

func main() {
	var sisi, luas, keliling float64

	// untuk inputan dari user
	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// menghitung luas persegi menggunakan prosedur hitungLuas
	hitungLuas(sisi, &luas)
	// menghitung keliling persegi menggunakan prosedur hitungKeliling
	hitungKeliling(sisi, &keliling)

	// menampilkan hasil dari perhitungan luas dan keliling persegi
	fmt.Printf("Luas persegi: %.2f\n", luas)
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}
```
### Screenshot Output
![Screenshot 2024-10-19 091312](https://github.com/user-attachments/assets/c07e6b42-4b82-4f94-844b-8f19919b2034)


### Deskripsi Program


**Algoritma Program:**
1. Program pertama-tama meminta input dari pengguna berupa panjang sisi persegi. Input ini disimpan dalam variabel `sisi` yang bertipe `float64` untuk mendukung angka desimal.
2. Setelah input diterima, program menghitung **luas** persegi dengan rumus:
   Luas = sisi x sisi
3. Kemudian, program menghitung **keliling** persegi dengan rumus:
  Keliling = 4 x sisi
4. Setelah perhitungan selesai, hasil **luas** dan **keliling** ditampilkan ke layar dengan format dua angka desimal menggunakan fungsi `fmt.Printf`.

**Cara Kerja Program:**
1. Program dimulai dengan mendeklarasikan variabel `sisi` yang akan menyimpan input dari pengguna. 
2. Program menampilkan pesan "Masukkan panjang sisi persegi:" untuk meminta pengguna memasukkan nilai panjang sisi persegi. Fungsi `fmt.Scan(&sisi)` digunakan untuk membaca input dari pengguna.
3. Setelah nilai sisi dimasukkan, program menghitung luas persegi dengan mengalikan nilai `sisi` dengan dirinya sendiri (sisi * sisi). Hasil perhitungan ini disimpan dalam variabel `luas`.
4. Program juga menghitung keliling persegi dengan mengalikan nilai `sisi` dengan 4 (4 * sisi). Hasilnya disimpan dalam variabel `keliling`.
5. Setelah perhitungan selesai, program menampilkan hasil perhitungan luas dan keliling menggunakan fungsi `fmt.Printf` untuk memformat hasil ke dua angka di belakang koma.
6. Output menampilkan nilai luas dan keliling yang sudah dihitung berdasarkan input panjang sisi yang diberikan oleh pengguna.

------

# UNGUIDED
## Unguided - 1
### Study Case
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap
𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
<br>


### Source Code
```
// Rangga Pradarrell Fathi
// 2311102200 / S1IF11-05
package main

import "fmt"

// Prosedur untuk input bilangan
func inputBilangan(prompt string, bilangan *int) {
    fmt.Print(prompt)
    fmt.Scan(bilangan)
}

// Prosedur untuk mengecek syarat a>=c dan b>=d
func cekSyarat(a, b, c, d int) bool {
    return a >= c && b >= d
}

// Prosedur untuk menampilkan hasil perhitungan
func tampilkanHasil(label string, permutasi, kombinasi int) {
    fmt.Printf("\nPermutasi %s: %d\n", label, permutasi)
    fmt.Printf("Kombinasi %s: %d\n", label, kombinasi)
}

// Prosedur untuk menghitung faktorial
func faktorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

// Prosedur untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
    var faktorialN, faktorialNR int
    faktorial(n, &faktorialN)
    faktorial(n-r, &faktorialNR)
    *hasil = faktorialN / faktorialNR
}

// Prosedur untuk menghitung kombinasi
func kombinasi(n, r int, hasil *int) {
    var faktorialN, faktR, faktorialNR int
    faktorial(n, &faktorialN)
    faktorial(r, &faktR)
    faktorial(n-r, &faktorialNR)
    *hasil = faktorialN / (faktR * faktorialNR)
}

// Prosedur untuk melakukan perhitungan permutasi dan kombinasi
func hitungPermutasiKombinasi(n, r int) (int, int) {
    var hasilPermutasi, hasilKombinasi int
    permutasi(n, r, &hasilPermutasi)
    kombinasi(n, r, &hasilKombinasi)
    return hasilPermutasi, hasilKombinasi
}

func main() {
    var a, b, c, d int

    // Input bilangan menggunakan prosedur
    inputBilangan("Bilangan a: ", &a)
    inputBilangan("Bilangan b: ", &b)
    inputBilangan("Bilangan c: ", &c)
    inputBilangan("Bilangan d: ", &d)

    // Cek syarat menggunakan prosedur
    if cekSyarat(a, b, c, d) {
        // Hitung permutasi dan kombinasi untuk (a,c)
        permutasiAC, kombinasiAC := hitungPermutasiKombinasi(a, c)
        tampilkanHasil("(a, c)", permutasiAC, kombinasiAC)

        // Hitung permutasi dan kombinasi untuk (b,d)
        permutasiBD, kombinasiBD := hitungPermutasiKombinasi(b, d)
        tampilkanHasil("(b, d)", permutasiBD, kombinasiBD)
    } else {
        fmt.Println("Input tidak sesuai dengan syarat yang ada")
    }
}
```
### Screenshot Output
![Screenshot 2024-10-19 091530](https://github.com/user-attachments/assets/5ab5f8b2-750c-45bc-ad0b-a94cd7f080cf)


### Deskripsi Program
Program di atas dibuat untuk menghitung **permutasi** dari dua bilangan `a` dan `b`. Permutasi adalah operasi yang menghitung banyaknya cara menyusun `r` objek dari `n` objek yang tersedia, di mana urutan sangat diperhatikan. Program ini meminta input dua bilangan (`a` dan `b`) dari pengguna dan kemudian menghitung permutasi dengan formula:

P(n,r) = n!/(n-r)!

di mana `n!` adalah faktorial dari `n`, yaitu hasil perkalian dari semua bilangan bulat positif dari 1 hingga `n`.

#### Algoritma Program

1. **Input dari Pengguna:**
   - Program pertama kali meminta dua bilangan `a` dan `b` dari pengguna.
   - Jika `a >= b`, program akan menghitung permutasi dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` akan ditukar sehingga `b` dianggap sebagai `n` dan `a` sebagai `r`.

2. **Perhitungan Faktorial:**
   - Fungsi `faktorial(n int) int` menghitung faktorial dari `n`. 
   - Fungsi ini menggunakan sebuah loop dari 1 hingga `n`, di mana setiap iterasi mengalikan nilai sebelumnya dengan `i` untuk menghitung faktorial.

3. **Perhitungan Permutasi:**
   - Fungsi `permutasi(n, r int) int` menghitung permutasi menggunakan rumus : Permutasi = faktorial(n)/faktorial(n-r).
   - Fungsi ini menggunakan nilai faktorial dari `n` dan `n - r`, lalu membaginya untuk mendapatkan hasil permutasi.

4. **Output:**
   - Hasil dari perhitungan permutasi ditampilkan ke layar menggunakan `fmt.Println`.

#### Cara Kerja Program

1. Program dimulai dengan mendeklarasikan dua variabel `a` dan `b` sebagai bilangan bulat (`int`).
2. Program meminta input dua bilangan dari pengguna dengan menggunakan `fmt.Scan(&a, &b)`.
3. Program mengecek apakah nilai `a` lebih besar atau sama dengan `b`. Jika ya, permutasi dihitung dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` ditukar.
4. Fungsi `permutasi` dipanggil dengan parameter `n` dan `r`, dan hasilnya dihitung dengan membagi faktorial `n` dengan faktorial `(n - r)`.
5. Setelah hasil permutasi diperoleh, hasil tersebut dicetak ke layar.

## Unguided - 2
### Study Case
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
![image](https://github.com/user-attachments/assets/c92445af-a30b-46d8-b6f8-4003870afefc)

Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

Keterangan:

Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.

### Source Code
```
// Rangga Pradarrell Fathi
// 2311102200 / S1IF11-05
package main

import (
    "fmt"
)

// Struct untuk menyimpan data peserta
type Peserta struct {
    nama      string
    waktu     [8]int
    soalBenar int
    totalWaktu int
}

// Prosedur untuk input jumlah peserta
func inputJumlahPeserta() int {
    var n int
    fmt.Print("Jumlah Peserta: ")
    fmt.Scan(&n)
    return n
}

// Prosedur untuk input data peserta
func inputDataPeserta(nomorPeserta int) Peserta {
    var peserta Peserta
    
    // Input nama peserta
    fmt.Printf("\nNama peserta %d: ", nomorPeserta+1)
    fmt.Scan(&peserta.nama)
    
    // Input waktu pengerjaan
    fmt.Print("Waktu Pengerjaan Soal (8 soal): ")
    for j := 0; j < 8; j++ {
        fmt.Scan(&peserta.waktu[j])
    }
    
    return peserta
}

// Prosedur untuk menghitung total soal yang dikerjakan dan total waktu
func hitungSkor(waktu [8]int, soal *int, totalWaktu *int) {
    *soal = 0
    *totalWaktu = 0
    for i := 0; i < 8; i++ {
        if waktu[i] <= 300 { // jika waktu pengerjaan kurang dari 300 menit, soal selesai
            *soal++
            *totalWaktu += waktu[i] // hanya tambahkan waktu soal yang selesai
        }
    }
}

// Prosedur untuk menentukan pemenang
func tentukanPemenang(peserta Peserta, pemenangSekarang *Peserta) {
    if peserta.soalBenar > pemenangSekarang.soalBenar || 
       (peserta.soalBenar == pemenangSekarang.soalBenar && peserta.totalWaktu < pemenangSekarang.totalWaktu) {
        *pemenangSekarang = peserta
    }
}

// Prosedur untuk menampilkan hasil akhir
func tampilkanHasil(pemenang Peserta) {
    fmt.Printf("\nNama pemenang: %s\n", pemenang.nama)
    fmt.Printf("Jumlah soal yang selesai: %d\n", pemenang.soalBenar)
    fmt.Printf("Total waktu yang dihabiskan: %d menit\n", pemenang.totalWaktu)
}

func main() {
    // Input jumlah peserta
    n := inputJumlahPeserta()
    
    // Inisialisasi data pemenang sementara
    pemenang := Peserta{
        soalBenar: -1,
        totalWaktu: 1000,
    }
    
    // Proses setiap peserta
    for i := 0; i < n; i++ {
        // Input data peserta
        peserta := inputDataPeserta(i)
        
        // Hitung skor peserta
        hitungSkor(peserta.waktu, &peserta.soalBenar, &peserta.totalWaktu)
        
        // Tentukan pemenang sementara
        tentukanPemenang(peserta, &pemenang)
    }
    
    // Tampilkan hasil akhir
    tampilkanHasil(pemenang)
}
```
### Screenshot Output
![Screenshot 2024-10-19 092255](https://github.com/user-attachments/assets/b79b6736-1151-49c5-9fcf-aadad1789882)


### Deskripsi Program
Program ini memiliki aplikasi untuk lomba atau kompetisi yang melibatkan peserta mengerjakan beberapa soal dalam waktu terbatas, dengan kriteria penentuan pemenang berdasarkan kecepatan dan jumlah soal yang diselesaikan.

### **Algoritma Program**

1. **Inisialisasi Variabel**:
   - Deklarasikan variabel untuk menyimpan nama pemenang, jumlah soal yang diselesaikan oleh pemenang, dan total waktu pemenang.
   
2. **Menerima Input**:
   - Dalam loop, terus terima input nama peserta dan waktu pengerjaan hingga pengguna memasukkan kata "Selesai".
   - Jika input nama peserta adalah "Selesai", keluar dari loop.

3. **Membaca Waktu Pengerjaan**:
   - Untuk setiap peserta, baca 8 angka yang merepresentasikan waktu pengerjaan masing-masing soal.
   
4. **Menghitung Skor**:
   - Panggil prosedur `hitungSkor`, yang menerima waktu pengerjaan dan mengembalikan jumlah soal yang diselesaikan dan total waktu yang dibutuhkan.
   
5. **Menentukan Pemenang**:
   - Bandingkan jumlah soal yang diselesaikan dengan pemenang sebelumnya. Jika peserta saat ini menyelesaikan lebih banyak soal, atau jika sama banyak tetapi dengan waktu lebih singkat, maka peserta ini menjadi pemenang sementara.
   
6. **Output**:
   - Setelah semua peserta dimasukkan, tampilkan nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang digunakan.

### **Cara Kerja Program**

1. **Input Peserta**:
   - Program meminta pengguna untuk memasukkan nama peserta diikuti oleh waktu pengerjaan 8 soal. Peserta dapat menginput waktu dalam menit. Jika seorang peserta tidak menyelesaikan suatu soal, waktu yang digunakan dianggap 301 menit.

2. **Memanggil Prosedur `hitungSkor`**:
   - Prosedur ini menganalisis waktu yang diinput. Setiap waktu yang lebih dari 300 menit diabaikan (dianggap tidak menyelesaikan soal), dan hanya soal yang diselesaikan dalam 300 menit atau kurang yang dihitung. Total waktu dari soal-soal yang diselesaikan dijumlahkan.

3. **Menentukan Pemenang**:
   - Program memeriksa apakah jumlah soal yang diselesaikan oleh peserta saat ini lebih banyak dari pemenang sebelumnya. Jika sama, program membandingkan total waktu untuk menentukan pemenang. Ini dilakukan hingga semua peserta diinput.

4. **Menampilkan Hasil**:
   - Setelah pengguna memasukkan "Selesai", program mencetak nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang digunakan.

# Unguided - 3
### Study Case
Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:
	22 11  34  17  52  26  13  40  20  10  5  16  8  4  2  1
 
Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.
Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

![image](https://github.com/user-attachments/assets/3ec0d166-c7a7-4d35-8da1-35927869e010)

Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.

Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.

### Source Code
```
// Rangga Pradarrell Fathi
// 2311102200 / S1IF11-05
package main

import (
    "fmt"
)

// Prosedur untuk input bilangan
func inputBilangan() int {
    var n int
    fmt.Print("Bilangan: ")
    fmt.Scan(&n)
    return n
}

// Prosedur untuk validasi input
func validasiInput(n int) bool {
    if n <= 0 || n >= 1000000 {
        fmt.Println("Nilai harus positif dan kurang dari 1000000.")
        return false
    }
    return true
}

// Prosedur untuk menghitung bilangan berikutnya dalam deret
func hitungBilanganBerikutnya(n int) int {
    if n%2 == 0 {
        return n / 2
    }
    return 3*n + 1
}

// Prosedur untuk menyimpan deret dalam slice
func buatDeret(n int) []int {
    deret := make([]int, 0)
    deret = append(deret, n)
    
    for n != 1 {
        n = hitungBilanganBerikutnya(n)
        deret = append(deret, n)
    }
    
    return deret
}

// Prosedur untuk mencetak deret
func cetakDeret(deret []int) {
    fmt.Print("Suku dan deret: ")
    for i, nilai := range deret {
        if i == len(deret)-1 {
            fmt.Printf("%d", nilai)
        } else {
            fmt.Printf("%d ", nilai)
        }
    }
    fmt.Println()
}

// Prosedur untuk analisis deret
func analisisDeret(deret []int) {
    panjangDeret := len(deret)
    nilaiTerbesar := deret[0]
    
    for _, nilai := range deret {
        if nilai > nilaiTerbesar {
            nilaiTerbesar = nilai
        }
    }
    
    fmt.Printf("\nAnalisis Deret:\n")
    fmt.Printf("Panjang deret: %d\n", panjangDeret)
    fmt.Printf("Nilai terbesar dalam deret: %d\n", nilaiTerbesar)
}

func main() {
    // Input bilangan
    n := inputBilangan()
    
    // Validasi input
    if validasiInput(n) {
        // Buat deret
        deret := buatDeret(n)
        
        // Cetak deret
        cetakDeret(deret)
        
        // Tampilkan analisis deret
        analisisDeret(deret)
    }
}
```

### Screenshot Output
![Screenshot 2024-10-19 092409](https://github.com/user-attachments/assets/c1cd761b-8dc5-44de-9dfd-3d2123fdc865)


### Deskripsi Program
Program ini merupakan implementasi yang menghasilkan deret angka berdasarkan aturan dari Collatz conjecture atau 3n + 1 problem. Program ini meminta input bilangan dari pengguna, membentuk deret berdasarkan aturan tertentu, menampilkan deret, dan memberikan analisis terkait deret tersebut.

#### **Algoritma Program**
1. **Inisialisasi**:
   - Mulai dengan meminta input dari pengguna untuk bilangan bulat positif \( n \) (kurang dari 1.000.000).

2. **Validasi Input**:
   - Periksa apakah \( n \) valid (harus lebih besar dari 0 dan kurang dari 1.000.000). Jika tidak valid, tampilkan pesan kesalahan dan akhiri program.

3. **Prosedur `cetakDeret`**:
   - Menerima parameter \( n \).
   - Dalam loop, lakukan hal berikut:
     - Cetak nilai \( n \).
     - Jika \( n \) genap, ubah \( n \) menjadi \( \frac{n}{2} \).
     - Jika \( n \) ganjil, ubah \( n \) menjadi \( 3n + 1 \).
   - Akhiri loop ketika \( n \) menjadi 1, dan cetak 1 sebagai suku terakhir.

4. **Output**:
   - Cetak deret bilangan yang dihasilkan di baris yang sama, dipisahkan oleh spasi.

#### **Cara Kerja Program**
1. **Input**:
   - Program meminta pengguna untuk memasukkan bilangan bulat positif.
   - Pengguna memasukkan nilai, yang disimpan dalam variabel \( n \).

2. **Validasi**:
   - Program memeriksa nilai \( n \):
     - Jika \( n \) lebih kecil atau sama dengan 0 atau lebih besar atau sama dengan 1.000.000, program akan menampilkan pesan kesalahan dan menghentikan eksekusi.

3. **Menghasilkan Deret**:
   - Program memanggil prosedur `cetakDeret` dengan nilai \( n \):
     - Di dalam prosedur, selama nilai \( n \) tidak sama dengan 1, program mencetak nilai \( n \) dan memperbarui nilainya sesuai dengan aturan (genap atau ganjil).
     - Setelah mencapai 1, program juga mencetak angka 1 sebagai suku terakhir.

4. **Output**:
   - Hasil deret bilangan ditampilkan di layar sebagai output akhir program, di mana semua suku dipisahkan oleh spasi.
