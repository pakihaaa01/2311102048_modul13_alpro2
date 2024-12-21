# 2311102048_modul13_alpro2
Deskripsi Program latihan soal nomor 1 : Program mengimplementasikan algoritma pengurutan dan pemeriksaan jarak yang sama pada sekumpulan bilangan bulat. Berikut adalah deskripsi singkat dari program tersebut:
1. **Pengumpulan Input**: Program meminta pengguna untuk memasukkan sekumpulan bilangan bulat. Pengguna dapat memasukkan bilangan satu per satu, dan proses input akan berhenti ketika pengguna memasukkan bilangan negatif.
2. **Pengurutan**: Setelah semua bilangan dimasukkan, program menggunakan algoritma *insertion sort* untuk mengurutkan bilangan-bilangan tersebut dalam urutan menaik.
3. **Pemeriksaan Jarak Sama**: Setelah pengurutan, program memeriksa apakah bilangan-bilangan yang telah diurutkan memiliki jarak yang sama antara satu bilangan dengan bilangan berikutnya. Jika jarak antara setiap bilangan adalah sama, program akan mengembalikan nilai `true` dan jarak tersebut. Jika tidak, program akan mengembalikan `false`.
4. **Output**: Program kemudian mencetak bilangan-bilangan yang telah diurutkan. Jika bilangan-bilangan tersebut memiliki jarak yang sama, program akan mencetak jarak tersebut. Jika tidak, program akan mencetak pesan bahwa jarak tidak tetap.
Program ini berguna untuk memeriksa apakah sekumpulan bilangan memiliki pola aritmetika setelah diurutkan.

Deskripsi Program latihan soal nomor 2 : program aplikasi manajemen buku sederhana yang ditulis dalam bahasa Go. Berikut adalah deskripsi singkat dari fungsionalitas program:
1. **Struktur Data**: 
   - `Buku`: Struktur yang menyimpan informasi tentang buku, termasuk `id`, `judul`, `penulis`, `penerbit`, `eksemplar`, `tahun`, dan `rating`.
   - `DaftarBuku`: Array yang dapat menampung hingga `nMax` buku.
2. **Fungsi Utama**:
   - `DaftarkanBuku`: Mengambil input dari pengguna untuk mengisi data buku ke dalam `DaftarBuku`.
   - `CetakTerfavorit`: Mencetak buku dengan rating tertinggi dari daftar.
   - `UrutBuku`: Mengurutkan buku berdasarkan rating dari yang tertinggi ke terendah menggunakan algoritma insertion sort.
   - `Cetak5Terbaru`: Mencetak judul dari lima buku pertama dalam daftar setelah diurutkan.
   - `CariBuku`: Mencari buku berdasarkan rating menggunakan pencarian biner dan mencetak detail buku jika ditemukan.
3. **Fungsi `main`**:
   - Mengambil jumlah buku (`n`) dari input pengguna.
   - Memanggil fungsi untuk mendaftarkan buku, mencetak buku terfavorit, mengurutkan buku, mencetak lima buku terbaru, dan mencari buku berdasarkan rating yang diberikan pengguna.
Program ini mengandalkan input dari pengguna untuk mengisi data buku dan melakukan operasi pencarian serta pengurutan berdasarkan rating buku.

deskripsi program latihan soal nomor 2 : 
