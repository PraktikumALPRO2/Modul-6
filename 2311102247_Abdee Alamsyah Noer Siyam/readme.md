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
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
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


## <strong> Dasar Teori </strong>

<strong><h2>PENGERTIAN REKURSIF</h2></strong>

Rekursi adalah konsep di mana suatu fungsi memanggil dirinya sendiri dengan cara langsung atau tidak langsung. Setiap panggilan ke fungsi rekursif adalah versi yang lebih kecil sehingga ia bertemu di suatu titik. Setiap fungsi rekursif memiliki kasus dasar atau kondisi dasar yang merupakan pernyataan eksekusi terakhir dalam rekursi dan menghentikan panggilan selanjutnya.

Namun rekursi sering kali digunakan untuk menyelesaikan masalah yang lebih kompleks, yaitu :
- yang tidak bisa (sulit) diselesaikan dengan loop biasa
- atau kode implementasinya akan sangat sulit dibaca jika menggunakan loop (iterasi)

<strong><h2>KOMPONEN REKURSIF</h2></strong>

Pada dasarnya, sebuah fungsi rekursif terdiri dari dua komponen utama:

- **Base Case (Kasus Dasar)**
Ini adalah kondisi di mana fungsi tidak akan memanggil dirinya sendiri lagi dan langsung memberikan hasil. Base case diperlukan untuk menghentikan rekursi dan mencegah terjadinya loop tak berujung.

- **Recursive Case (Kasus Rekursif)**
Bagian ini adalah di mana fungsi memanggil dirinya sendiri dengan versi yang lebih sederhana atau lebih kecil dari masalah aslinya.

<strong><h2>JENIS REKURSIF</h2></strong>

Ada berbagai jenis rekursi seperti yang dijelaskan dalam contoh berikut:

### Retur Langsung
Jenis rekursi di mana fungsi memanggil dirinya sendiri secara langsung, tanpa bantuan fungsi lain disebut rekursi langsung. Contoh berikut menjelaskan konsep rekursi langsung:

Contoh :
```go
func faktorial(number int) int { 
 
	if number == 0 || number == 1 { 
		return 1 
	} 
	

	if number < 0 { 
		fmt.Println("Angka Tidak Valid") 
		return -1 
	} 

	return number*faktorial(number - 1) 
} 
```
### Retur Tidak Langsung
Jenis rekursi di mana suatu fungsi memanggil fungsi lain dan fungsi ini, pada gilirannya, memanggil fungsi pemanggilan disebut rekursi tak langsung. Jenis rekursi ini membutuhkan bantuan fungsi lain. Fungsi tersebut memanggil dirinya sendiri, tetapi secara tidak langsung, yaitu melalui fungsi lain.

Contoh :
```go

func print_one(n int) { 
      
    if n >= 0 { 
        fmt.Println("In first function:", n) 
        print_two(n - 1) 
    } 
} 
  
func print_two(n int) { 

    if n >= 0 { 
        fmt.Println("In second function:", n) 
        print_one(n - 1) 
    } 
} 

```


## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

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

![image](https://github.com/user-attachments/assets/bf4ccc46-918f-4919-8053-3bdc52bd99cd)


Deskripsi Program : 
Program di atas adalah aplikasi sederhana yang mencetak urutan angka menurun dari bilangan yang dimasukkan pengguna hingga 1. Program ini terdiri dari dua fungsi yaitu fungsi baris dan fungsi main. Fungsi baris adalah fungsi rekursif yang menerima satu bilangan. Jika bilangan tersebut bernilai 1, fungsi mencetak angka 1 dan berhenti. Jika bilangan lebih dari 1, fungsi mencetak nilai saat ini, lalu memanggil dirinya sendiri dengan bilangan - 1, hingga mencapai angka 1. Urutan angka yang dihasilkan ditampilkan sesuai dengan input pengguna.

### <h2>GUIDED 2</h2>

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

![image](https://github.com/user-attachments/assets/893ca55e-8b64-4ab1-816c-41f810094770)

Deskripsi Program : 
Program di atas adalah aplikasi sederhana yang menghitung jumlah total dari bilangan yang dimasukkan pengguna hingga 1. Program ini memiliki dua fungsi, yaitu fungsi penjumlahan dan fungsi main. Fungsi penjumlahan adalah fungsi rekursif yang menerima satu parameter n. Jika n bernilai 1, fungsi akan mengembalikan 1 dan menghentikan rekursi. Jika n lebih besar dari 1, fungsi akan mengembalikan nilai n ditambah hasil dari penjumlahan(n-1), sehingga menghitung total penjumlahan dari n hingga 1. Hasil penjumlahan ini kemudian ditampilkan di layar sesuai dengan angka input dari pengguna.

### <h2>GUIDED 3</h2>

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

![image](https://github.com/user-attachments/assets/c7775d40-ecaa-4bcb-b874-3a2a4019a6bd)

Deskripsi Program : 
Program di atas adalah aplikasi sederhana yang menghitung nilai  $ 2^n $ (2 pangkat n) menggunakan fungsi rekursif. Program terdiri dari dua fungsi: fungsi pangkat dan fungsi main. Fungsi pangkat menerima satu parameter n. Jika n bernilai 0, fungsi mengembalikan 1 karena $ 2^0 = 1 $. Jika n lebih besar dari 0, fungsi mengembalikan hasil perkalian 2 dengan pangkat(n-1), yang secara rekursif menghitung nilai $ 2^n $. Hasil perhitungan $ 2^n $ ini ditampilkan di layar sesuai dengan nilai n yang dimasukkan oleh pengguna.

### <h2>GUIDED 4</h2>

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

![image](https://github.com/user-attachments/assets/26875c7e-d0fa-4a5a-98e8-eb7a9a7e844d)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang menghitung nilai faktorial dari bilangan yang dimasukkan oleh pengguna. Program ini memiliki dua fungsi : fungsi faktorial dan fungsi main. Fungsi faktorial adalah fungsi rekursif yang menghitung hasil faktorial dari n. Jika n bernilai 0 atau 1, fungsi mengembalikan 1 (karena $0! = 1! = 1$). Jika n lebih besar dari 1, fungsi mengembalikan hasil perkalian n dengan faktorial(n-1), yang secara rekursif menghitung faktorial hingga mencapai 1. Hasil faktorial ditampilkan sesuai dengan nilai input pengguna.

## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question

Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan $S_n = S_{n-1} + S_{n-2}$. Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

| n  | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8  | 9  | 10 |
|----|---|---|---|---|---|---|---|---|----|----|----|
| $S_n$ | 0 | 1 | 1 | 2 | 3 | 5 | 8 | 13 | 21 | 34 | 55 |

#### Source Code

```go
package main

import (
	"fmt"
)

func deretFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return deretFibonacci(n-1) + deretFibonacci(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("n(%d) = Sn(%d)\n", i, deretFibonacci(i))
	}
}

```

#### Output

![image](https://github.com/user-attachments/assets/7ea5469b-a327-4cad-85d2-e44bd2b3e377)

Deskripsi Program : 
Program di atas adalah aplikasi sederhana yang menghitung deret Fibonacci hingga elemen ke-10. Deret Fibonacci adalah urutan angka di mana setiap angka adalah hasil penjumlahan dari dua angka sebelumnya, dimulai dari 0 dan 1.Program ini terdiri dari dua fungsi : fungsi deretFibonacci dan fungsi main. Fungsi deretFibonacci adalah fungsi rekursif yang menerima satu parameter n. Jika n bernilai 0 atau 1, fungsi mengembalikan n (karena $F(0)=0$ dan $F(1)=1$). Jika n lebih besar dari 1, fungsi mengembalikan hasil penjumlahan dari $deretFibonacci(n-1)$ dan $deretFibonacci(n-2)$, yang secara rekursif menghitung nilai Fibonacci ke-n.Fungsi main mencetak deret Fibonacci dari elemen ke-0 hingga ke-10, menampilkan setiap nilai dalam format $n(i) = Sn(i)$.

### <h2>UNGUIDED 2</h2>

#### Question

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

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	cetakBintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("N : ")
	fmt.Scan(&n)
	cetakBintang(n)
}


```

#### Output

![image](https://github.com/user-attachments/assets/bb4422b2-44d8-4a5e-bef9-a1dc69fb93f0)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang mencetak pola bintang bertingkat berdasarkan bilangan yang dimasukkan pengguna. Program membaca input integer n dan mencetak pola bintang sebanyak n baris, di mana setiap baris memiliki jumlah bintang yang menurun dari atas ke bawah. Program ini memiliki dua fungsi : fungsi cetakBintang dan fungsi main. Fungsi cetakBintang adalah fungsi rekursif yang mencetak bintang secara bertingkat. Jika n bernilai 0, fungsi berhenti. Jika n lebih besar dari 0, fungsi memanggil dirinya sendiri dengan n-1 untuk mencetak baris sebelumnya terlebih dahulu, lalu mencetak n bintang pada baris saat ini. Hasilnya adalah pola bintang yang menurun, mulai dari n bintang di baris pertama hingga 1 bintang di baris terakhir.

### <h2>UNGUIDED 3</h2>

#### Question

Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.

   Masukan terdiri dari sebuah bilangan bulat positif N.

   Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

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

func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d", i)
	}
	cetakFaktor(n, i+1)
}

func main() {
	var angka int
	fmt.Print("N : ")
	fmt.Scan(&angka)

	fmt.Printf("Faktor %d : ", angka)
	cetakFaktor(angka, 1)
	fmt.Println()
}

```

#### Output

![image](https://github.com/user-attachments/assets/e0de6765-8caa-4fad-9a2c-3f0390ac20ec)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang mencetak faktor dari bilangan yang dimasukkan pengguna. Program terdiri dari dua bagian fungsi : fungsi cetakFaktor dan fungsi main. Fungsi cetakFaktor adalah fungsi rekursif yang menerima dua parameter: n (bilangan yang ingin dicari faktornya) dan i (angka yang digunakan untuk membagi n). Jika i lebih besar dari n, fungsi berhenti. Jika n dapat dibagi oleh i (yaitu n % i == 0), maka i dicetak sebagai faktor. Fungsi main mencetak faktor dari angka, dimulai dari 1 hingga bilangan itu sendiri, dan hasilnya ditampilkan di layar.

### <h2>UNGUIDED 4</h2>

#### Question

Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.

   Masukan terdiri dari sebuah bilangan bulat positif N.

   Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran             |
   |----|---------|----------------------|
   | 1  | 5       | 5 4 3 2 1 2 3 4 5    |
   | 2  | 9       | 9 8 7 6 5 4 3 2 1 2 3 4 5 6 7 8 9 |

#### Source Code

```go
package main

import (
	"fmt"
)

func desc(n int) {
	if n < 1 {
		return
	}
	fmt.Printf("%d ", n)
	desc(n - 1)
}

func asce(n, mulai int) {
	if mulai > n {
		return
	}
	fmt.Printf("%d ", mulai)
	asce(n, mulai+1)
}

func main() {
	var angka int
	fmt.Print("N : ")
	fmt.Scan(&angka)
	fmt.Print("Hasil : ")
	desc(angka)
	asce(angka, 1)
	fmt.Println()
}

```

#### Output

![image](https://github.com/user-attachments/assets/07f17911-1fef-46f2-a513-f30c9459bd8b)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang mencetak urutan angka secara menurun dan kemudian secara menaik berdasarkan bilangan yang dimasukkan pengguna. Program terdiri dari tiga fungsi : fungsi desc, fungsi asce, dan fungsi main. Fungsi desc adalah fungsi rekursif yang mencetak angka dari n hingga 1. Jika n kurang dari 1, fungsi berhenti. Jika tidak, fungsi mencetak nilai n dan memanggil dirinya sendiri dengan n - 1. Fungsi asce adalah fungsi rekursif yang mencetak angka dari 1 hingga n. Jika mulai lebih besar dari n, fungsi berhenti. Jika tidak, fungsi mencetak nilai mulai dan memanggil dirinya sendiri dengan mulai + 1.

Fungsi main menggabungkan kedua fungsi tersebut. Setelah membaca nilai angka, program mencetak hasil dari fungsi desc diikuti dengan hasil dari fungsi asce, sehingga menghasilkan urutan menurun diikuti oleh urutan menaik.

### <h2>UNGUIDED 5</h2>

#### Question

Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.

   Masukan terdiri dari sebuah bilangan bulat positif N.

   Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

   Contoh masukan dan keluaran:

   | No | Masukan | Keluaran                   |
   |----|---------|----------------------------|
   | 1  | 5       | 1 3 5                      |
   | 2  | 20      | 1 3 5 7 9 11 13 15 17 19   |

#### Source Code

```go
package main

import "fmt"

func cetakGanjil(batas int, nilai int) {
	if nilai > batas {
		return
	}
	if nilai%2 != 0 {
		fmt.Printf("%d ", nilai)
	}
	cetakGanjil(batas, nilai+1)
}

func main() {
	var batas int
	fmt.Print("N : ")
	fmt.Scan(&batas)
	fmt.Printf("Bilangan ganjil sampai %d adalah : ", batas)
	cetakGanjil(batas, 1)
	fmt.Println()
}

```

#### Output

![image](https://github.com/user-attachments/assets/8d20167c-be0b-464b-bb4f-a8e3a977ef2b)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang mencetak bilangan ganjil dari 1 hingga batas yang ditentukan pengguna. Program terdiri dari dua fungsi : fungsi cetakGanjil dan fungsi main. Fungsi cetakGanjil adalah fungsi rekursif yang menerima parameter batas (nilai maksimum yang akan dicetak) dan nilai (angka saat ini yang akan diperiksa). Jika nilai lebih besar dari batas, fungsi akan berhenti. Jika nilai adalah bilangan ganjil (yaitu nilai % 2 != 0), fungsi mencetak nilai tersebut. Fungsi kemudian memanggil dirinya sendiri dengan nilai + 1, untuk memeriksa bilangan berikutnya. Fungsi main mencetak bilangan ganjil dari 1 hingga batas yang dimasukkan oleh pengguna, memberikan hasilnya di layar.

### <h2>UNGUIDED 6</h2>

#### Question

Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.

   Masukan terdiri dari bilangan bulat x dan y.

   Keluaran terdiri dari hasil x dipangkatkan y.

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

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan 2 bilangan bulat : ")
	fmt.Scan(&x, &y)

	result := pangkat(x, y)
	fmt.Printf("Hasil pangkat : %d\n", result)
}

```

#### Output

![image](https://github.com/user-attachments/assets/7fdeddc5-589c-4527-bba6-db75aebd5f46)

Deskripsi Program :
Program di atas adalah aplikasi sederhana yang menghitung nilai pangkat dari dua bilangan bulat yang dimasukkan pengguna. Program terdiri dari dua fungsi : fungsi pangkat dan fungsi main. Fungsi pangkat adalah fungsi rekursif yang menerima dua parameter: x (basis) dan y (pangkat). Jika y bernilai 0, fungsi mengembalikan 1 (karena $x^0 = 1$). Jika tidak, fungsi mengembalikan hasil perkalian x dengan pemanggilan rekursif pangkat(x, y-1), yang menghitung pangkat y dengan mengurangi satu pada setiap langkah. Fungsi main mencetak hasil perhitungan pangkat berdasarkan bilangan yang dimasukkan pengguna. Setelah membaca nilai x dan y, hasil perhitungan ditampilkan di layar.

## <strong> Kesimpulan </strong>

Rekursi adalah konsep pemrograman di mana suatu fungsi memanggil dirinya sendiri, baik secara langsung maupun tidak langsung, dengan tujuan menyelesaikan masalah yang kompleks yang sulit diselesaikan dengan metode iterasi biasa atau yang akan sulit dibaca jika menggunakan loop. Setiap fungsi rekursif terdiri dari dua komponen utama: kasus dasar (base case), yang menghentikan pemanggilan fungsi untuk mencegah loop tak berujung dan memberikan hasil akhir, serta kasus rekursif (recursive case), di mana fungsi memanggil dirinya sendiri dengan versi yang lebih sederhana dari masalah tersebut. Rekursi dapat dibedakan menjadi dua jenis: rekursi langsung, di mana fungsi memanggil dirinya sendiri secara langsung, seperti dalam contoh perhitungan faktorial, dan rekursi tidak langsung, di mana suatu fungsi memanggil fungsi lain yang pada gilirannya memanggil fungsi pemanggil, seperti dalam contoh fungsi yang saling memanggil untuk mencetak nilai. Penggunaan rekursi memungkinkan penyelesaian masalah yang lebih elegan dan mudah dipahami dalam banyak kasus.

## <strong> Referensi </strong>

#### [1] Berbagai Jenis Rekursi di Golang. Diakses melalui website https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/

#### [2] APA ITU Rekursif : Pengertian dan Informasi. Diakses melalui website https://sko.dev/wiki/rekursif
