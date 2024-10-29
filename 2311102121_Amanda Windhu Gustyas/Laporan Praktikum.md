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





