# Go REST API Project dengan React

Proyek ini adalah implementasi **REST API** menggunakan **Go** untuk backend dan **React** untuk frontend. API ini menangani data tentang **judul buku** dan **penulis**.

## Deskripsi

Proyek ini bertujuan untuk menyediakan REST API sederhana yang dapat menangani informasi tentang buku, termasuk:
- **Judul Buku**: Nama buku yang dimiliki.
- **Penulis**: Nama penulis dari buku tersebut.

Backend menggunakan **Go** dengan **Gin Framework** untuk menangani HTTP requests, sementara frontend menggunakan **React** untuk menampilkan data buku dan penulis di browser.

### Fitur:
- **CRUD** (Create, Read, Update, Delete) untuk buku dan penulis.
- Pengambilan data buku dan penulis dari backend menggunakan REST API.
- Frontend interaktif untuk menampilkan data buku dan penulis.
  
## Teknologi yang Digunakan

- **Backend**:
  - **Go**: Bahasa pemrograman untuk backend.
  - **Gin Framework**: Framework web untuk Go yang ringan dan cepat.
  - **GORM**: ORM untuk berinteraksi dengan database.
  - **PostgreSQL**: Database relasional untuk menyimpan data buku dan penulis.
  
- **Frontend**:
  - **React**: Library JavaScript untuk membangun antarmuka pengguna.
  - **Axios**: Untuk melakukan HTTP requests ke backend.
