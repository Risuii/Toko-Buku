# Toko Buku

## Deskripsi
Simple CRUD API untuk toko buku yang  memiliki fitur untuk CRUD toko serta CRUD buku yang ada di dalam toko tersebut

## Stack

- Golang (go1.17.8)
- Framework: Gin-Gonic
- Database: MongoDB

## Cara Menjalankan

1. Buka terminal di dalam visual studio code 
2. Pastikan alamat folder di terminal sudah sesuai dengan folder dari code
3. ketik `go run main.go` untuk menjalankan program

## Fungsional

* Import Postman Collection di dalam folder postman dan rubah variabel {{base3000}} menjadi localhost:3000

* Akses `localhost:3000/v1/stotre/create` dengan method POST untuk membuat toko

* Akses `localhost:3000/v1/store/toko/:kodeToko` dengan method GET untuk melihat toko yang di inginkan sesuai dengan ID yang didapatkan

* Akses `localhost:3000/v1/store/getall` dengan method GET untuk melihat semua toko yang ada

* Akses `localhost:3000/v1/store/edit/:kodeToko` dengan method PUT untuk mengubah atau mengedit toko yang di inginkan sesuai dengan kode tokonya

* Akses `localhost:3000/v1/delete/:kodeToko` dengan method DELETE untuk menghapus toko yang di inginkan sesuai dengan kode tokonya

* Akkses `localhost:3000/v1/store/book/:kodeToko` dengan method POST untuk membuat buku didalam toko yang di inginkan

* Akses `localhost:3000/v1/store/book/:kodeToko/:kodeBuku` dengan method GET untuk melihat buku yang di inginkan di dalam toko yang di inginkan

* Akses `localhost:3000/v1/store/book/all/:kodeToko` dengan method GET untuk melihat semua buku yang ada di dalam toko yang di inginkan

* Akses `localhost:3000/v1/store/book/:kodeToko/:kodeBuku` dengan method PUT untuk mengubah atau mengedit buku yang di inginkan di dalam toko yang di inginkan

* Akses `localhost:3000/v1/store/book/:kodeToko/:kodeBuku` dengan method DELETE untuk menghapus buku yang di inginkan di dalam toko yang di inginkan
