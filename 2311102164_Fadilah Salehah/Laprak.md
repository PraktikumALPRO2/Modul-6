
<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL VI</strong>
  <br>
  <strong>REKURSIF</strong>
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
  Fadilah Salehah
  <br>
  2311102164
  <br>
  S1IF1105
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

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN REKURSIF</strong></span>
<br>
Rekursif adalah konsep pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih besar dengan cara memecahnya menjadi masalah-masalah yang lebih kecil. Dalam Go, rekursif memungkinkan penyelesaian masalah kompleks dengan kode yang lebih elegan dan mudah dipahami.

Setiap fungsi rekursif harus memiliki :
- Base case - kondisi yang menghentikan rekursi
- Recursive case - kondisi di mana fungsi memanggil dirinya sendiri
- Progress towards base case - setiap pemanggilan rekursif harus mendekati base case

**Penggunaan Rekursif**

Rekursif dalam Go umumnya digunakan untuk :

- Perhitungan matematis (faktorial, fibonacci, dll)
- Traversal struktur data (tree, graph)
- Pencarian dan pengurutan (binary search, merge sort)
- Problem solving dengan pendekatan divide and conquer
- File system traversal

<span style="font-size:16px"><strong> ── JENIS REKURSIF</strong></span>
<br>

**Direct Recursion (Rekursif Langsung)**

Fungsi yang memanggil dirinya sendiri secara langsung dalam definisinya.
Karakteristik :
- Pemanggilan fungsi terjadi dalam fungsi yang sama
- Lebih mudah dipahami dan diimplementasikan

Contoh implementasi :

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)  // Memanggil dirinya sendiri secara langsung
}
```

**Indirect Recursion (Rekursif Tidak Langsung)**

Terjadi ketika fungsi A memanggil fungsi B, yang kemudian memanggil kembali fungsi A.
Karakteristik :
- Melibatkan dua atau lebih fungsi yang saling memanggil
- Lebih kompleks untuk di-debug

Contoh implementasi :

```go
func isEven(n int) bool {
    if n == 0 {
        return true
    }
    return isOdd(n - 1)
}

func isOdd(n int) bool {
    if n == 0 {
        return false
    }
    return isEven(n - 1)
}
```

**Tail Recursion (Rekursif Ekor)**

Bentuk rekursif di mana pemanggilan rekursif adalah operasi terakhir yang dilakukan dalam fungsi.
Karakteristik :
- Lebih efisien dalam penggunaan memori
- Dapat dioptimasi oleh compiler menjadi iterasi
- Tidak memerlukan stack frame tambahan

Contoh implementasi :

```go
func factorialTail(n int, accumulator int) int {
    if n == 0 {
        return accumulator
    }
    return factorialTail(n-1, n*accumulator)  // Pemanggilan rekursif di akhir
}
```

**Tree Recursion (Rekursif Pohon)**

Fungsi rekursif yang memanggil dirinya sendiri lebih dari satu kali.
Karakteristik :
- Membentuk struktur seperti pohon dalam pemanggilan
- Biasanya digunakan untuk masalah yang memiliki cabang keputusan

Contoh implementasi :

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)  // Dua pemanggilan rekursif
}
```

<span style="font-size:16px"><strong> ── KELEBIHAN DAN KEKURANGAN REKURSIF</strong></span>
<br>

**Kelebihan**
- Kode lebih bersih dan mudah dibaca
- Ideal untuk masalah yang memiliki struktur rekursif alami
- Mempermudah implementasi algoritma kompleks
- Mengurangi pengulangan kode

**Kekurangan**
- Konsumsi memori stack yang lebih tinggi
- Risiko stack overflow pada rekursi dalam
- Performa bisa lebih lambat dibanding iterasi
- Debugging lebih menantang





## <strong> GUIDED </strong>

### ── Guided 1

#### Source Code

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
	fmt.Scan(&n)
	baris(n)
}
```

#### Output
![image](https://github.com/user-attachments/assets/1b239410-72a7-48fd-bfe0-b0035865104b)

#### Deskripsi Program :
Program ini digunakan untuk menampilkan deretan angka secara menurun mulai dari angka yang diberikan oleh user hingga angka 1. Program akan meminta user untuk memasukkan satu bilangan bulat, lalu akan menampilkan angka tersebut hingga 1 secara berurutan menurun. Program ini menggunakan fungsi rekursif baris yang mencetak angka saat ini dan memanggil dirinya sendiri dengan angka yang dikurangi 1 sampai mencapai angka 1, sehingga menghasilkan deretan angka dari bilangan input hingga 1.

### ── Guided 2

#### Source Code

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
	fmt.Print(penjumlahan(n))
}
```

#### Output
![image](https://github.com/user-attachments/assets/8d6c6275-3557-4d8a-aea3-c1954d0e6ebd)

#### Deskripsi Program :
Program ini digunakan untuk menghitung jumlah total dari angka 1 hingga angka yang diberikan oleh user. Program akan meminta user memasukkan satu bilangan bulat, lalu menghitung jumlah dari semua angka mulai dari 1 sampai bilangan tersebut. Program ini menggunakan fungsi rekursif penjumlahan yang menjumlahkan angka saat ini dengan hasil dari pemanggilan fungsi berikutnya yang dikurangi 1, hingga mencapai angka 1. Hasil akhirnya adalah jumlah total dari semua angka mulai dari 1 hingga angka input user.

### ── Guided 3

#### Source Code

```go
package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan Nilai n : ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah", pangkat(n))
}
```

#### Output
![image](https://github.com/user-attachments/assets/f3b19d92-0a8f-4e20-b95a-f68bb92db2de)

#### Deskripsi Program :
Program ini digunakan untuk menghitung hasil dari 2 pangkat suatu bilangan yang diberikan oleh user. Program akan meminta user memasukkan sebuah bilangan bulat, kemudian menghitung 2 pangkat bilangan tersebut. Program ini menggunakan fungsi rekursif pangkat yang mengalikan 2 dengan hasil pemanggilan fungsi berikutnya yang dikurangi 1, hingga mencapai 2 pangkat 0, yang nilainya adalah 1. Hasil akhirnya adalah nilai dari 2 pangkat bilangan yang diinputkan oleh user.

### ── Guided 4

#### Source Code

```go
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}
```

#### Output
![image](https://github.com/user-attachments/assets/960254b5-34f8-46b5-b8ef-abc27281bfe0)

#### Deskripsi Program :
Program ini digunakan untuk menghitung faktorial dari sebuah bilangan yang diberikan oleh user. Program akan meminta user memasukkan satu bilangan bulat, lalu menghitung faktorial dari bilangan tersebut. Program ini menggunakan fungsi rekursif faktorial yang mengalikan bilangan saat ini dengan hasil pemanggilan fungsi berikutnya yang dikurangi 1, hingga mencapai 0 atau 1, di mana nilai faktorialnya adalah 1. Hasil akhirnya adalah nilai faktorial dari bilangan yang diinputkan oleh user.

## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case

Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan $S_n = S_{n-1} + S_{n-2}$. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

| n  | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8  | 9  | 10 |
|----|---|---|---|---|---|---|---|---|----|----|----|
| $S_n$ | 0 | 1 | 1 | 2 | 3 | 5 | 8 | 13 | 21 | 34 | 55 |

#### Source Code

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Fibonacci(", i, ") =", fibonacci(i))
	}
}

```

#### Output
![image](https://github.com/user-attachments/assets/28cc8b78-e1db-4223-a4ba-b0b49cf8cb30)


#### Deskripsi Program :
Program ini digunakan untuk menghitung dan menampilkan deret Fibonacci hingga angka ke-10. Fungsi fibonacci bekerja secara rekursif untuk menghitung nilai Fibonacci dari suatu bilangan n. Jika n adalah 0 atau 1, fungsi akan mengembalikan nilai n. Jika lebih besar dari 1, fungsi menambahkan hasil dari dua pemanggilan sebelumnya (fibonacci(n-1) + fibonacci(n-2)). Program utama menampilkan deret Fibonacci dari 0 hingga 10.

### ── Unguided 2

#### Study Case

Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran        |
   |----|---------|-----------------|
   | 1  | 5       | *               |
   |    |         | **              |
   |    |         | ***             |
   |    |         | ****            |
   |    |         | *****           |
   | 2  | 1       | *               |
   | 3  | 3       | *               |
   |    |         | **              |
   |    |         | ***             |

#### Source Code

```go
package main

import (
	"fmt"
)

func cetakSegitiga(n int) {
	if n == 0 {
		return
	}
	cetakSegitiga(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&n)
	cetakSegitiga(n)
}
```

#### Output
![image](https://github.com/user-attachments/assets/0ae43afa-5490-497a-aec8-1d3eef0e20eb)


#### Deskripsi Program :
Program ini digunakan untuk mencetak pola segitiga bintang dengan tinggi yang ditentukan user. Program meminta user memasukkan angka, yang digunakan untuk menentukan tinggi segitiga. Fungsi cetakSegitiga bekerja secara rekursif untuk mencetak setiap baris segitiga dengan jumlah bintang yang berkurang setiap kali pemanggilan. Setiap kali fungsi dipanggil, pola bintang bertambah hingga mencapai tinggi yang dimasukkan.

### ── Unguided 3

#### Study Case

Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.

   **Masukan** terdiri dari sebuah bilangan bulat positif N.

   **Keluaran** terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran         |
   |----|---------|------------------|
   | 1  | 5       | 1 5             |
   | 2  | 12      | 1 2 3 4 6 12    |

#### Source Code

```go
package main

import (
	"fmt"
)

func faktorisasi(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorisasi(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah : ")
	faktorisasi(n, 1)
	fmt.Println()
}
```

#### Output
![image](https://github.com/user-attachments/assets/c9f41f91-cd6b-4f77-973a-0c759cdd0bd9)


#### Deskripsi Program :
Program ini digunakan untuk mencari dan menampilkan semua faktor dari bilangan yang diberikan user. Program meminta user memasukkan bilangan bulat positif, lalu fungsi faktorisasi memeriksa secara rekursif apakah setiap angka dari 1 hingga n adalah faktor dari bilangan tersebut. Jika n habis dibagi oleh i, maka i akan dicetak sebagai faktor dari n.

### ── Unguided 4

#### Study Case

Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.

   **Masukan** terdiri dari sebuah bilangan bulat positif N.

   **Keluaran** terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran             |
   |----|---------|----------------------|
   | 1  | 5       | 5 4 3 2 1 2 3 4 5    |
   | 2  | 9       | 9 8 7 6 5 4 3 2 1 2 3 4 5 6 7 8 9 |

#### Source Code

```go
package main

import "fmt"

func cetakFaktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		cetakFaktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah: ")
	cetakFaktor(n, 1)
	fmt.Println()
}

```

#### Output
![image](https://github.com/user-attachments/assets/9373b84c-b5f5-4266-83c5-3052a10198c4)


#### Deskripsi Program :
Program ini digunakan untuk mencetak bilangan dalam urutan menurun dari angka yang diberikan hingga 1, lalu dilanjutkan dengan mencetak urutan menaik dari 1 hingga angka tersebut. Program meminta user memasukkan bilangan bulat positif, kemudian menggunakan fungsi descending untuk mencetak bilangan secara menurun, dan ascending untuk mencetak bilangan secara menaik, masing-masing secara rekursif.


### ── Unguided 5

#### Study Case

Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.

   **Masukan** terdiri dari sebuah bilangan bulat positif N.

   **Keluaran** terdiri dari barisan bilangan ganjil dari 1 hingga N.

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran                   |
   |----|---------|----------------------------|
   | 1  | 5       | 1 3 5                      |
   | 2  | 20      | 1 3 5 7 9 11 13 15 17 19   |

#### Source Code

```go
package main

import "fmt"

func cetakGanjil(n, current int) {
	if current <= n {
		if current%2 != 0 {
			fmt.Print(current, " ")
		}
		cetakGanjil(n, current+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	fmt.Print("Bilangan ganjil dari 1 hingga ", n, " adalah: ")
	cetakGanjil(n, 1)
	fmt.Println()
}

```

#### Output
![image](https://github.com/user-attachments/assets/c0af873c-fb4b-4452-bf23-ced04d738d1d)


#### Deskripsi Program :
Program ini digunakan untuk menampilkan bilangan ganjil dari 1 hingga bilangan yang diberikan oleh user. Program meminta user memasukkan bilangan bulat positif N, kemudian menggunakan fungsi rekursif printGanjil untuk memeriksa setiap angka dari 1 hingga N. Jika suatu angka ganjil, angka tersebut akan dicetak.

### ── Unguided 6

#### Study Case

Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.

   **Masukan** terdiri dari bilangan bulat x dan y.

   **Keluaran** terdiri dari hasil x dipangkatkan y.

   Catatan: diperbolehkan menggunakan asterisk "*", tapi dilarang menggunakan import "math".

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran |
   |----|---------|----------|
   | 1  | 2 2     | 4        |
   | 2  | 5 3     | 125      |

#### Source Code

```go
package main

import "fmt"

func hitungPangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * hitungPangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x dan y: ")
	fmt.Scan(&x, &y)

	fmt.Printf("Hasil %d pangkat %d: %d\n", x, y, hitungPangkat(x, y))
}

```

#### Output
![image](https://github.com/user-attachments/assets/fda49526-43ac-4312-90a5-5af11a2c3322)

#### Deskripsi Program :
Program ini digunakan untuk menghitung hasil pemangkatan dari dua bilangan bulat yang diberikan user. Program meminta user memasukkan dua bilangan bulat x dan y, kemudian menggunakan fungsi rekursif pangkat untuk menghitung hasil x pangkat y. Jika y sama dengan 0, fungsi mengembalikan nilai 1. jika tidak, fungsi mengalikan x dengan hasil pemanggilan fungsi dengan eksponen y yang berkurang 1.


## <strong> REFERENSI </strong>

#### [1] geeksforgeeks. “Different Types of Recursion in Golang”, 2020. https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/
