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
  Yesika Widiyani / 2311102195 <br>
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

# DASAR TEORI
## REKRUSIF
Rekursif adalah teknik pemrograman di mana suatu fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih kompleks dengan memecahnya menjadi masalah-masalah kecil. Fungsi rekursif akan terus memanggil dirinya hingga mencapai kondisi tertentu yang menghentikan pemanggilan berulang ini, yang disebut dengan base case.</br>

### **1. Elemen Utama Rekusif**
Dalam setiap fungsi rekursif, ada dua elemen utama:

a) Base Case (Kondisi Dasar): Kondisi yang menghentikan rekursi. Base case menentukan kapan fungsi harus berhenti memanggil dirinya sendiri. Tanpa base case yang benar, fungsi rekursif dapat menjadi infinite loop dan menghabiskan memori (stack overflow).</br>
b) Recursive Case (Kasus Rekursif): Bagian dari fungsi yang memanggil dirinya sendiri untuk mendekati base case. Ini adalah inti dari rekursi yang secara bertahap memperkecil masalah.</br>

### **2. Keunggulan dan Kelemahan Rekusif**

**Keunggulan:**
Sederhana dan Elegan: Rekursi dapat membuat kode lebih mudah dibaca dan dipahami, terutama untuk masalah yang dapat dibagi menjadi sub-masalah berulang (seperti faktorial, Fibonacci, pencarian dalam struktur data pohon).
Mendekati Penyelesaian Top-Down: Pendekatan top-down memungkinkan kita fokus pada pemecahan masalah secara umum tanpa langsung ke detail implementasi.</br>

**Kelemahan:**
Mengonsumsi Memori Lebih Banyak: Setiap pemanggilan fungsi rekursif menyimpan status sebelumnya di stack, sehingga memori bisa cepat habis (stack overflow) jika pemanggilan terlalu dalam.
Tidak Efisien untuk Kasus Tertentu: Contoh kasus Fibonacci yang memiliki tumpang tindih sub-masalah, membuat fungsi menghitung ulang nilai yang sama berulang kali. Ini bisa diatasi dengan optimasi seperti memoization.</br>

### **3. Cara Kerja Rekursif**
Cara kerja rekursif dalam pemrograman adalah ketika sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan masalah dengan cara membaginya menjadi sub-masalah yang lebih kecil hingga mencapai kondisi yang sudah ditentukan untuk berhenti (base case). Pada dasarnya, cara kerja rekursif adalah *divide and conquer*, atau "pecah dan taklukkan."</br>

Berikut adalah penjelasan tahapan cara kerja rekursif:</br>

**1. Memanggil Fungsi Utama**
Pemanggilan fungsi rekursif dimulai dari fungsi utama yang memanggil dirinya sendiri dengan parameter awal. Dalam proses ini:</br>
   - Fungsi pertama kali dipanggil dengan nilai atau data input.</br>
   - Fungsi tersebut akan menguji apakah sudah mencapai kondisi berhenti (base case) atau belum.</br>

**2. Mengecek Kondisi Dasar (Base Case)**
**Base case** adalah kondisi yang menentukan kapan fungsi tidak lagi memanggil dirinya sendiri. Ini merupakan syarat utama dalam rekursif untuk mencegah terjadinya infinite loop atau *stack overflow* (jika fungsi terus memanggil dirinya sendiri tanpa batas).</br>
   - Jika base case terpenuhi, fungsi akan berhenti melakukan pemanggilan berulang dan mengembalikan hasil.</br>
   - Jika base case belum terpenuhi, fungsi akan memanggil dirinya sendiri dengan input atau parameter yang telah diubah untuk mendekati base case.</br>

**3. Proses Pemanggilan Ulang (Recursive Case)**
Setiap kali fungsi belum mencapai base case, fungsi tersebut akan memanggil dirinya lagi dengan nilai parameter yang lebih mendekati kondisi berhenti. Dalam recursive case:</br>
   - Nilai parameter pada fungsi dikurangi atau diubah setiap kali pemanggilan baru dilakukan.</br>
   - Pemanggilan fungsi ini akan terus berjalan hingga base case tercapai.</br>

**4. Penyimpanan di Stack Memori**
Setiap kali fungsi dipanggil, komputer menyimpan setiap status fungsi (nilai variabel dan kondisi saat ini) dalam *stack memori*. Stack ini bekerja dengan prinsip *Last In, First Out (LIFO)*, di mana pemanggilan terakhir yang dilakukan akan diselesaikan lebih dahulu.</br>
   - Ketika fungsi mencapai base case, panggilan berhenti, dan fungsi mulai mengembalikan nilai dari panggilan terakhir ke panggilan sebelumnya secara berurutan.</br>
   - Hasil perhitungan dari setiap pemanggilan akan dikembalikan ke tingkat panggilan sebelumnya, dan nilai akhir diperoleh setelah semua panggilan selesai.</br>

**5. Unwinding: Mengembalikan Nilai dari Fungsi**
Setelah mencapai base case, stack mulai “di-unwind”, yaitu melepas satu per satu panggilan dari yang terakhir dipanggil hingga kembali ke pemanggilan fungsi utama.</br>
   - Setiap level pemanggilan akan menyelesaikan kalkulasinya berdasarkan hasil dari panggilan berikutnya, hingga akhirnya fungsi utama mendapatkan hasil akhir.</br>

**Menghindari Infinite Loop dengan Base Case**
Untuk memastikan rekursif tidak berjalan terus-menerus, base case harus didefinisikan dengan benar. Jika tidak, fungsi akan terus memanggil dirinya tanpa henti dan menyebabkan program mengalami *stack overflow*.

Itulah cara kerja rekursif dalam bahasa pemrograman, termasuk Go. Konsep dasar ini penting untuk memahami bagaimana sebuah masalah besar dapat dipecah menjadi sub-masalah kecil dalam proses rekursif.

-------
# GUIDED
## Guided - 1
### Study Case
a.	Membuat baris bilangan dari n hingga 1</br>
Base-case: bilangan == 1</br>

### Source Code
![guid1](https://github.com/user-attachments/assets/bf9c11ae-d1c9-4973-b3d3-75c01ad4d955)

### Screenshot Output
![ssg1](https://github.com/user-attachments/assets/75ac2072-f840-4980-bbed-201bec49501c)

### Deskripsi Program
Program ini meminta input berupa bilangan bulat 𝑛 dari pengguna dan kemudian mencetak bilangan tersebut dalam urutan menurun hingga mencapai angka 1 menggunakan rekursif.</br>

**Alur Algoritma**
1. Pengguna memasukkan nilai 𝑛.</br>
2. Fungsi baris dipanggil dengan parameter n.</br>
3. Jika bilangan sama dengan 1, program mencetak angka 1 dan berhenti.</br>
4. Jika tidak, program mencetak bilangan dan memanggil baris dengan bilangan - 1.</br>
5. Proses ini berulang hingga bilangan mencapai 1.</br>

**Cara Kerja Program**
a. Program pertama-tama membaca input nilai 𝑛. </br>
b. Fungsi baris bekerja secara rekursif. Pada setiap pemanggilan, fungsi mencetak angka saat ini dan memanggil dirinya sendiri dengan angka yang dikurangi 1.</br>
c. Ketika nilai mencapai 1, program berhenti karena base case terpenuhi.</br>

## Guided - 2
### Study Case
b.	Menghitung hasil penjumlahan 1 hingga n </br>
Base-case: n == 1</br>

### Source Code
![guid2](https://github.com/user-attachments/assets/f9cb6795-fb10-45fe-a15c-aa9dd3792317)

### Screenshot Output
![ssg2](https://github.com/user-attachments/assets/42b2e7e7-53ee-4052-be06-59a894b983b8)

### Deskripsi Program
Program ini meminta input berupa bilangan 𝑛 dari pengguna dan menghitung jumlah total dari 1 hingga 𝑛 secara rekursif.</br>

**Alur Algoritma**
1. Pengguna memasukkan nilai 𝑛.</br>
2. Fungsi penjumlahan dipanggil dengan parameter n.</br>
3. Jika n sama dengan 1, fungsi mengembalikan 1 (base case).</br>
4. Jika tidak, fungsi mengembalikan n + penjumlahan(n-1), di mana penjumlahan(n-1) adalah hasil penjumlahan rekursif berikutnya.</br>
5. Nilai yang dikembalikan dari setiap panggilan rekursif dijumlahkan hingga mencapai nilai total.</br>

**Cara Kerja Program**
a. Program membaca input 𝑛 dari pengguna.</br>
b. Fungsi penjumlahan bekerja secara rekursif, mengurangi nilai 𝑛 setiap kali dipanggil hingga mencapai base case, yaitu ketika 𝑛 = 1</br>
c. Setelah mencapai base case, fungsi menghitung mundur dan mengembalikan jumlah total dari 1 hingga 𝑛 sebagai hasil.</br>

## Guided - 3
### Study Case
c.	Mencari dua pangkat n atau 2𝑛</br>
Base-case: n == 0</br>

### Source Code
![guid3](https://github.com/user-attachments/assets/5ba98e5c-37b7-4085-b718-374064c8d509)

### Screenshot Output
![ssg3](https://github.com/user-attachments/assets/33e42dd0-9a8f-4bac-83e2-f9e1cad5a76c)

### Deskripsi Program
Program ini meminta input berupa bilangan bulat 𝑛 dan menghitung nilai 2^𝑛 secara rekursif.

**Alur Algoritma**
1. Pengguna memasukkan nilai 𝑛.</br>
2. Fungsi pangkat dipanggil dengan parameter n.</br>
3. Jika n sama dengan 0, fungsi mengembalikan 1 (karena 2^0 = 1 adalah base case).</br>
4. Jika tidak, fungsi mengembalikan hasil dari 2 * pangkat(n-1).</br>
5. Setiap pemanggilan rekursif mengalikan 2 dengan hasil panggilan sebelumnya hingga mencapai base case.</br>

**Cara Kerja Program**
a. Program meminta pengguna untuk memasukkan nilai 𝑛 sebagai eksponen.</br>
b. Fungsi pangkat bekerja secara rekursif dengan mengalikan 2 pada setiap pemanggilan hingga mencapai base case.</br>
c. Setelah mencapai base case, hasil pangkat dihitung secara mundur hingga hasil akhirnya diperoleh.</br>

## Guided - 4
### Study Case
d.	Mencari nilai faktorial atau n! </br>
Base-case: n == 0 atau n == 1</br>

### Source Code
![guid4](https://github.com/user-attachments/assets/c76b5148-c4bf-4402-9e18-6714cc9d613b)

### Screenshot Output
![ssg4](https://github.com/user-attachments/assets/c0a3fe9c-6fdf-4687-8043-cdbdec110b84)

### Deskripsi Program
Program ini meminta input bilangan bulat 𝑛 dan menghitung faktorialnya (𝑛!) secara rekursif.</br>

**Alur Algoritma**
1. Pengguna memasukkan nilai 𝑛.</br>
2. Fungsi faktorial dipanggil dengan parameter n.</br>
3. Jika n sama dengan 0 atau 1, fungsi mengembalikan 1 (base case: 0!=1 dan 1!=1).</br>
4. Jika tidak, fungsi mengembalikan n * faktorial(n-1), yang mengalikan 𝑛 dengan hasil faktorial dari nilai sebelumnya.</br>
5. Fungsi akan terus memanggil dirinya dengan nilai yang berkurang hingga mencapai base case.</br>

**Cara Kerja Program**
a. Program membaca nilai 𝑛 dari pengguna.</br>
b. Fungsi faktorial menggunakan pemanggilan rekursif, mengurangi 𝑛 pada setiap pemanggilan hingga base case tercapai.</br>
c. Ketika base case tercapai, setiap nilai dari pemanggilan sebelumnya dikalikan secara bertahap hingga menghasilkan faktorial dari 𝑛.</br>

## UNGUIDED
## Unguided - 1
### Study Case
1)	Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan 𝑆𝑛 = 𝑆𝑛+1 + 𝑆𝑛+2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.</br>
![image](https://github.com/user-attachments/assets/bd2cd29a-41f0-4cc7-88d9-242c829510a4)

### Source Code
![ung1](https://github.com/user-attachments/assets/bc76cf09-0c76-4f67-80c5-95b7ea2cc6de)

### Screenshot Output
![image](https://github.com/user-attachments/assets/b5b58e47-ff4c-4c57-9dca-c78303ed6fa5)

### Deskripsi Program
Program ini menghitung dan mencetak deret Fibonacci hingga suku ke-10 menggunakan fungsi rekursif `fibonacci`.

**Alur Program**
1. Nilai `n = 10` ditetapkan, mewakili batas deret Fibonacci.
2. Dalam perulangan `for`, fungsi `fibonacci(i)` dipanggil dari suku ke-0 hingga ke-10.
3. Fungsi `fibonacci` bekerja rekursif:
   - Jika `n` adalah 0 atau 1, fungsi mengembalikan nilai `n`.
   - Jika `n > 1`, fungsi mengembalikan `fibonacci(n-1) + fibonacci(n-2)`.
4. Setiap hasil dikembalikan dan ditampilkan sebagai "Suku ke-i: [nilai]".

**Cara Kerja Program**

1. **Inisialisasi dan Output Pertama**:
   - Program menetapkan `n = 10` dan mencetak "Deret Fibonacci hingga suku ke-10".

2. **Perulangan untuk Mencetak Setiap Suku Fibonacci**:
   - Perulangan `for i := 0; i <= n; i++` berjalan dari `i = 0` hingga `i = 10`.
   - Pada setiap iterasi, `fibonacci(i)` dipanggil untuk menghitung nilai Fibonacci pada suku ke-\( i \).

3. **Pemanggilan Rekursif pada Fungsi `fibonacci`**:
   - Fungsi `fibonacci` menghitung nilai Fibonacci dari \( i \) secara rekursif.
   - Jika \( i = 0 \) atau \( i = 1 \), fungsi mengembalikan nilai \( i \) langsung (base case).
   - Untuk \( i > 1 \), fungsi memanggil dirinya sendiri dua kali: `fibonacci(i-1)` dan `fibonacci(i-2)`, kemudian menjumlahkan hasilnya.

4. **Pengembalian Hasil dan Penampilan Suku Fibonacci**:
   - Setiap kali hasil suku Fibonacci dikembalikan dari fungsi, program mencetak hasil tersebut dalam format "Suku ke-\( i \): [nilai]".
   - Proses ini berulang hingga program menampilkan suku ke-10 dari deret Fibonacci.

## Unguided - 2
### Study Case
2)	Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.</br>
Contoh masukan dan keluaran:</br>
![image](https://github.com/user-attachments/assets/266dcbeb-47a0-46fb-b57a-674abf4e0aa7)

### Source Code
![ung2](https://github.com/user-attachments/assets/ec8761f7-e43e-4e4d-b741-de5340076e24)

### Screenshot Output
![image](https://github.com/user-attachments/assets/f2b2f9f4-5aaa-42a3-a974-c4b49e94c9f4)


### Deskripsi Program
Program ini menggunakan rekursi untuk mencetak pola bintang dalam bentuk segitiga. Program meminta input bilangan \( n \), yang menentukan jumlah baris pada pola. Setiap baris berisi jumlah bintang yang meningkat sesuai dengan nomor barisnya (baris pertama memiliki 1 bintang, baris kedua memiliki 2 bintang, dan seterusnya).</br>

**Alur Algoritma Program**
1. **Input Pengguna**:
   - Program meminta pengguna memasukkan nilai \( n \), yang menunjukkan jumlah baris dalam pola bintang.

2. **Fungsi `printPattern`**:
   - Fungsi `printPattern(n, current)` bertanggung jawab mencetak baris-baris bintang.
   - Base case: Jika `current` melebihi `n`, fungsi berhenti (kembali dari rekursi).
   - Jika tidak, fungsi memanggil `printStars(current)` untuk mencetak `current` jumlah bintang.
   - Setelah mencetak bintang, fungsi berpindah ke baris berikutnya dengan memanggil `printPattern(n, current+1)`.

3. **Fungsi `printStars`**:
   - Fungsi `printStars(n)` bertugas mencetak bintang dalam satu baris.
   - Base case: Jika `n` sama dengan 0, fungsi berhenti.
   - Jika tidak, fungsi memanggil dirinya sendiri dengan `n-1`, kemudian mencetak satu bintang (`*`).
   - Proses ini berulang hingga jumlah bintang yang diperlukan untuk baris tersebut tercapai.

**Cara Kerja Program**
1. **Inisialisasi dan Input**:
   - Program meminta pengguna memasukkan nilai \( n \), lalu mencetak "Pola Bintang:" sebagai header pola bintang.

2. **Pemanggilan Rekursif pada `printPattern`**:
   - Program memanggil `printPattern(n, 1)` untuk mulai mencetak dari baris pertama.
   - Setiap pemanggilan `printPattern` mencetak baris `current` dengan memanggil `printStars(current)`, yang mencetak jumlah bintang sesuai nomor baris.

3. **Mencetak Pola Bintang**:
   - `printStars` menggunakan rekursi untuk mencetak `current` jumlah bintang dalam satu baris, lalu kembali ke `printPattern`.
   - Setelah satu baris selesai, `printPattern` menambah `current` dan memanggil dirinya lagi untuk mencetak baris berikutnya hingga mencapai baris ke-\( n \).

## Unguided - 3
### Study Case
3)	Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. </br>
Masukan terdiri dari sebuah bilangan bulat positif N. </br>

Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).</br>

Contoh masukan dan keluaran:</br>
![image](https://github.com/user-attachments/assets/978f6ddb-51df-43f7-ab08-330b3b81801f)

### Source Code
![ugn3](https://github.com/user-attachments/assets/b50c84be-e727-4c52-a5f2-12fb93be791f)

### Screenshot Output
![image](https://github.com/user-attachments/assets/3c255230-0a77-4ea5-baba-b60311e31c1f)

### **Deskripsi Program**
Program ini menggunakan fungsi rekursif untuk mencari dan menampilkan faktor dari suatu bilangan \( n \). Faktor adalah bilangan yang dapat membagi \( n \) secara sempurna tanpa sisa. Program meminta input bilangan \( n \) dari pengguna, lalu mencetak semua faktor dari 1 hingga \( n \).</br>

**Alur Algoritma Program**
1. **Input Pengguna**:
   - Program meminta pengguna memasukkan nilai \( n \), yaitu bilangan yang akan dicari faktornya.

2. **Fungsi `printFactors`**:
   - Fungsi `printFactors(n, current)` mencari faktor-faktor dari \( n \).
   - Base case: Jika `current` lebih besar dari \( n \), fungsi berhenti.
   - Jika `n % current == 0`, artinya `current` adalah faktor dari \( n \), dan program mencetaknya.
   - Fungsi kemudian memanggil dirinya sendiri dengan `current+1` untuk memeriksa angka berikutnya sebagai faktor potensial.

3. **Proses Cetak Faktor**:
   - Pada setiap pemanggilan `printFactors`, fungsi mengecek apakah `current` adalah faktor dari \( n \).
   - Setiap faktor yang ditemukan dicetak dengan spasi sebagai pemisah, lalu proses berlanjut hingga semua faktor tercetak.

**Cara Kerja Program**
1. **Inisialisasi dan Input**:
   - Program meminta pengguna memasukkan nilai \( n \) dan mencetak "Faktor dari \( n \):" sebagai judul output.

2. **Pemanggilan Rekursif pada `printFactors`**:
   - Fungsi `printFactors(n, 1)` dipanggil untuk memulai pencarian faktor dari angka 1.
   - Setiap kali `printFactors` dipanggil, program memeriksa apakah `current` adalah faktor dari \( n \).
   - Jika iya, faktor tersebut dicetak, dan fungsi melanjutkan ke angka berikutnya dengan `current + 1`.

3. **Pengulangan hingga `current > n`**:
   - Fungsi terus berlanjut hingga nilai `current` melewati \( n \).
   - Setelah semua faktor ditemukan, program mencetak hasil akhir di satu baris.

## Unguided - 4
### Study Case
4)	Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.
Masukan terdiri dari sebuah bilangan bulat positif N.</br>
Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.</br>

Contoh masukan dan keluaran:</br>
![image](https://github.com/user-attachments/assets/0074a381-8f94-4fbd-ba7b-21af15790069)

### Source Code
![ung4](https://github.com/user-attachments/assets/8b5cb64a-ac1a-4b30-a8df-1fdc7b52ddbd)

### Screenshot Output
![image](https://github.com/user-attachments/assets/197dfb31-314e-4f75-97dd-b50b02977106)

### Deskripsi Program
Program ini menggunakan fungsi rekursif untuk mencetak urutan angka dari \( n \) ke 1, lalu kembali dari 1 ke \( n \). Program meminta input bilangan \( n \), kemudian mencetak pola angka secara simetris dengan urutan menurun dari \( n \) hingga 1 dan kembali naik dari 1 hingga \( n \).</br>

**Alur Algoritma Program**
1. **Input Pengguna**:
   - Program meminta pengguna memasukkan nilai \( n \), yaitu batas atas dari urutan angka yang akan dicetak.

2. **Fungsi `printSequence`**:
   - Fungsi `printSequence(n)` mencetak angka dari \( n \) hingga 1 dan kemudian kembali mencetak dari 1 hingga \( n \).
   - Base case: Jika `n == 1`, fungsi mencetak angka `1` dan kembali ke pemanggilan sebelumnya (rekursi berhenti).
   - Jika `n > 1`, fungsi mencetak nilai `n`, memanggil dirinya sendiri dengan `n - 1`, lalu mencetak kembali `n` setelah pemanggilan rekursif selesai.

3. **Polanya Simetris**:
   - Karena nilai \( n \) dicetak sebelum dan sesudah pemanggilan rekursif, hasilnya adalah pola simetris: menurun dari \( n \) hingga 1, lalu naik kembali dari 1 hingga \( n \).

**Cara Kerja Program**
1. **Inisialisasi dan Input**:
   - Program meminta pengguna memasukkan nilai \( n \) dan langsung memanggil fungsi `printSequence(n)` untuk mulai mencetak urutan.

2. **Pemanggilan Rekursif pada `printSequence`**:
   - Setiap kali `printSequence(n)` dipanggil, program mencetak nilai `n`, lalu memanggil dirinya sendiri dengan `n-1`.
   - Setelah mencapai base case (saat `n == 1`), program mulai kembali mencetak nilai `n` pada setiap langkah pemanggilan sebelumnya.

3. **Polanya Terbentuk**:
   - Angka-angka dicetak menurun hingga 1, dan saat kembali dari pemanggilan rekursif, angka-angka tersebut dicetak naik kembali, menghasilkan pola simetris.

Urutan ini menunjukkan pola menurun dari 4 ke 1 dan naik kembali dari 1 ke 4.

## Unguided - 5
### Study Case
5)	Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
Masukan terdiri dari sebuah bilangan bulat positif N.</br>
Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.</br>

Contoh masukan dan keluaran:</br>
![image](https://github.com/user-attachments/assets/34aec91d-8c8f-48c6-8344-3234a5630c9b)

### Source Code
![ung5](https://github.com/user-attachments/assets/9c2cd03e-3c31-4d8b-beb6-e69cf810dede)

### Screenshot Program
![image](https://github.com/user-attachments/assets/deea591d-58fe-40e4-a28f-f6482a992cd0)

### Deskripsi Program
Program ini mencetak deret bilangan ganjil secara rekursif, dimulai dari angka 1 hingga batas angka \( n \) yang ditentukan oleh pengguna. Program menggunakan fungsi `printOddSequence` untuk mencetak setiap bilangan ganjil secara bertahap dengan menambahkan 2 pada setiap pemanggilan rekursif.

**Alur Algoritma Program**
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan nilai \( n \), yaitu batas atas dari bilangan ganjil yang ingin dicetak.

2. **Fungsi `printOddSequence`**:
   - Fungsi `printOddSequence(n, current)` bertugas mencetak bilangan ganjil secara berurutan dari angka 1 hingga \( n \).
   - Base case: Jika `current` lebih besar dari \( n \), fungsi berhenti dan kembali ke pemanggilan sebelumnya.
   - Jika `current` masih dalam batas \( n \), program mencetak nilai `current` dan memanggil `printOddSequence` lagi dengan nilai `current + 2` untuk beralih ke bilangan ganjil berikutnya.

3. **Pencetakan Rekursif**:
   - Fungsi mencetak nilai `current`, lalu memanggil dirinya sendiri dengan menambahkan 2 pada `current`, menghasilkan deret bilangan ganjil hingga mencapai atau melewati batas \( n \).

### **Cara Kerja Program**
1. **Inisialisasi dan Input**:
   - Program meminta pengguna memasukkan nilai \( n \) dan langsung memanggil fungsi `printOddSequence(n, 1)` untuk mulai mencetak deret bilangan ganjil dari 1.

2. **Pemanggilan Rekursif pada `printOddSequence`**:
   - Fungsi mencetak nilai `current`, dimulai dari 1, kemudian memanggil `printOddSequence` dengan nilai `current + 2` pada setiap langkah.
   - Proses ini berulang hingga `current` melebihi \( n \).

3. **Hasil Akhir**:
   - Semua bilangan ganjil dari 1 hingga batas \( n \) tercetak secara berurutan, dan fungsi berakhir saat `current` melebihi \( n \).

Program mencetak semua bilangan ganjil dari 1 hingga 9, lalu berhenti karena nilai berikutnya (11) melebihi batas \( n = 10 \).</br>

## Unguided - 6
### Study Case
6)	Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.</br>
Masukan terdiri dari bilangan bulat x dan y.</br>
Keluaran terdiri dari hasil x dipangkatkan y.</br>

Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan import "math".</br>

Contoh masukan dan keluaran:</br>
![image](https://github.com/user-attachments/assets/fd501a43-2aa1-48b5-a16a-d213762a3811)

### Source Code
![ung6](https://github.com/user-attachments/assets/0be0c948-e9bd-4e2d-86b9-e74b7ffb02c5)

### Screenshot Code
![image](https://github.com/user-attachments/assets/34e2a185-6458-47fb-995a-010a7fbb647a)

### Deskripsi Program
Program ini menghitung nilai perpangkatan \( x \) pangkat \( y \) secara rekursif menggunakan fungsi `power`. Program menerima dua input, yaitu basis \( x \) dan eksponen \( y \), lalu mengembalikan hasil \( x^y \).

**Alur Algoritma Program**
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan nilai \( x \) (basis) dan \( y \) (eksponen).

2. **Fungsi `power`**:
   - Fungsi `power(x, y)` menghitung hasil perpangkatan \( x \) pangkat \( y \) secara rekursif.
   - Base case: Jika `y` adalah 0, fungsi mengembalikan 1 karena \( x^0 = 1 \).
   - Jika \( y > 0 \), fungsi mengembalikan `x * power(x, y-1)`, yang berarti mengalikan \( x \) dengan hasil perpangkatan \( x^{y-1} \).

3. **Proses Rekursif**:
   - Fungsi memanggil dirinya sendiri dengan mengurangi nilai eksponen \( y \) sebanyak 1 pada setiap pemanggilan hingga mencapai base case (`y == 0`).
   - Setiap hasil dikalikan dengan \( x \), sehingga proses perkalian dilakukan sebanyak \( y \) kali untuk menghitung \( x^y \).

**Cara Kerja Program**
1. **Inisialisasi dan Input**:
   - Program meminta pengguna memasukkan nilai \( x \) (basis) dan \( y \) (eksponen) dan menyimpannya dalam variabel `x` dan `y`.

2. **Pemanggilan Fungsi `power`**:
   - Program memanggil `power(x, y)` dan menyimpan hasilnya dalam variabel `result`.
   - Fungsi `power` kemudian bekerja secara rekursif untuk menghitung hasil dari \( x^y \), menggunakan base case `y == 0` sebagai titik berhenti.

3. **Hasil Akhir**:
   - Setelah rekursi selesai, hasilnya dikembalikan ke `main` dan dicetak dalam format "Hasil: [result]".
