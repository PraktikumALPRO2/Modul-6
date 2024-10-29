[<p align="center">
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
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
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

### Screenshot Code


### Deskripsi Program


**Algoritma Program**

**Cara Kerja Program:**

     - Setelah mencapai 1, program juga mencetak angka 1 sebagai suku terakhir.

4. **Output**:
   - Hasil deret bilangan ditampilkan di layar sebagai output akhir program, di mana semua suku dipisahkan oleh spasi.
