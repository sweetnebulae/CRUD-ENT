# 📚 GoEnt REST API

RESTful API sederhana untuk manajemen data buku menggunakan Go, Entgo ORM, dan pendekatan Clean Architecture.

---

## 🔧 Tech Stack

- **Golang** `1.20+`
- **Entgo** (Entity Framework untuk Go)
- **HttpRouter** (`julienschmidt/httprouter`)
- **PostgreSQL** (atau basis data lain sesuai konfigurasi)
- **Clean Architecture** (lapisan model, repository, service, controller)

---

## 📁 Project Structure

```text
go_ent/
├── config/         # Konfigurasi database & HTTP server
├── controller/     # Logika handler HTTP
├── data/
│   ├── request/    # Struct untuk menerima data dari client (DTO)
│   └── response/   # Struct untuk mengirim data ke client
├── helper/         # Fungsi utilitas (error handler, JSON encoder)
├── model/          # Representasi internal data
├── repository/     # Lapisan data access
├── router/         # Routing HTTP
├── schema/         # Skema Entgo (Book.go)
├── service/        # Logika bisnis utama
├── ent/            # Berkas hasil generate Ent
└── main.go         # Entry point aplikasi
```

---

## 🚀 Getting Started

### 1. Clone repository

```bash
git clone https://github.com/username/go_ent.git
cd go_ent
```

### 2. Setup environment

Sesuaikan koneksi database pada file `config/postgres.go`. Contoh:

```go
dsn := "host=localhost port=5432 user=postgres password=yourpassword dbname=go_ent sslmode=disable"
```

### 3. Generate ent schema

```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```

### 4. Run the app

```bash
go run main.go
```

---

## 📡 API Endpoint

| Method | Endpoint           | Description          |
|--------|--------------------|----------------------|
| GET    | `/api/posts`       | Ambil semua buku     |
| GET    | `/api/posts/:id`   | Ambil satu buku      |
| POST   | `/api/post`        | Tambah buku          |
| PUT    | `/api/post/:id`    | Update buku          |
| DELETE | `/api/post/:id`    | Hapus buku           |

---

## 📦 Contoh Request

### POST `/api/post`

```json
{
  "title": "Laskar Pelangi",
  "author": "Andrea Hirata",
  "genres": "Drama",
  "description": "Kisah inspiratif anak-anak Belitong."
}
```

### PUT `/api/post/{id}`

```json
{
  "id": "uuid-di-sini",
  "title": "Laskar Pelangi Revised",
  "author": "Andrea Hirata",
  "genre": "Drama",
  "description": "Versi revisi dengan tambahan konten."
}
```

---

## 🧊 Catatan

- Gunakan tool seperti Postman, Insomnia, atau Curl untuk menguji endpoint.
- Pastikan database aktif dan sesuai dengan konfigurasi.

---

## 🛠️ Credits

Dikembangkan oleh [@sweetnebulae](https://github.com/sweetnebulae) dengan ❤️ untuk latihan dan produksi API yang bersih dan maintainable.