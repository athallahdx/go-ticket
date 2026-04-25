# Event Ticket Management System - Project Planning & Structure

## 📁 Struktur Folder Project

Proyek ini menggunakan struktur standar *Standard Go Project Layout*. Berikut adalah penjelasan untuk masing-masing folder yang telah dibuat:

- **`cmd/server/`**: Titik masuk (entry point) aplikasi. Berisi `main.go` yang bertugas untuk inisiasi server, membaca konfigurasi, menyambungkan ke database, dan memulai *listener* router HTTP.
- **`internal/`**: Berisi kode sumber aplikasi yang sifatnya *private* dan spesifik untuk project ini. Kode di dalam tidak bisa di-import oleh project lain.
  - **`internal/config/`**: Kode untuk membaca dan memparsing environment variables (`env`).
  - **`internal/domain/`** (atau model): Berisi definisi `struct` untuk entitas utama (seperti `User`, `Ticket`, `Event`) serta interface untuk repository dan service.
  - **`internal/repository/`**: Layer yang bertanggung jawab untuk interaksi langsung dengan database MySQL. Semua *query* SQL berada di sini.
  - **`internal/service/`**: Layer *Business Logic*. Menerima request dari handler, memproses logika aplikasi, dan memanggil fungsi di repository.
  - **`internal/handler/`**: Berisi HTTP handlers (controller). Bertugas menerima request dari user, validasi payload, dan memanggil service, serta mengembalikan response JSON. Menggunakan *Chi Router*.
  - **`internal/middleware/`**: Menyimpan middleware HTTP, contohnya untuk validasi JWT dan proteksi endpoint (Authentication & Authorization).
- **`pkg/`**: Berisi *library* yang sifatnya *public* atau *utility/helper* yang dapat digunakan secara bebas, bahkan bisa dipakai ulang di project lain.
  - **`pkg/hash/`**: Utility untuk *hashing* dan *checking* password menggunakan `bcrypt`.
  - **`pkg/jwt/`**: Utility untuk men-generate dan memvalidasi JSON Web Token (JWT).
  - **`pkg/response/`**: Helper untuk memformat standard response REST API (JSON success/error format).
- **`migrations/`**: Tempat untuk menyimpan file `.sql` untuk *database migration* (pembuatan tabel `users`, `events`, `tickets`, dll).

---

## 📝 Task List Programmer (Project Planning)

Berikut adalah langkah-langkah yang harus dilakukan untuk membangun project ini. Beri tanda centang `[x]` jika task sudah selesai.

### Tahap 1: Inisiasi Project & Konfigurasi Dasar
- [x] Inisialisasi Go Module (`go mod init <nama-module>`).
- [x] Install package *Chi router* (`go get -u github.com/go-chi/chi/v5`).
- [x] Install driver MySQL (`go get -u github.com/go-sql-driver/mysql`).
- [x] Install package untuk JWT (`go get -u github.com/golang-jwt/jwt/v5`).
- [x] Install package untuk bcrypt (`go get -u golang.org/x/crypto/bcrypt`).
- [x] Install package untuk .env, misalnya godotenv (`go get -u github.com/joho/godotenv`).
- [x] Buat file `.env` dan `.env.example` dengan variabel konfigurasi (DB_HOST, DB_USER, DB_PASS, DB_NAME, JWT_SECRET, PORT).
- [ ] Buat fungsi pembaca konfigurasi (`env`) di dalam `internal/config/config.go`.
- [ ] Buat inisialisasi koneksi MySQL di `cmd/server/main.go`.

### Tahap 2: Setup Helper & Middleware
- [ ] Implementasi fungsi Hash Password dan Compare Password dengan `bcrypt` di `pkg/hash/hash.go`.
- [ ] Implementasi fungsi Generate Token dan Validate Token dengan `jwt` di `pkg/jwt/jwt.go`.
- [ ] Buat helper standarisasi format response JSON di `pkg/response/response.go`.
- [ ] Buat middleware Auth menggunakan JWT di `internal/middleware/auth.go` untuk memproteksi endpoint.

### Tahap 3: Database & Entitas (Domain)
- [ ] Buat file migrasi SQL di folder `migrations/` untuk tabel `users`, `events`, dan `tickets`.
- [ ] Tulis definisi struct model `User`, `Event`, dan `Ticket` beserta interface repository/service di `internal/domain/`.

### Tahap 4: Fitur Autentikasi (User)
- [ ] Buat `user_repository` untuk operasi database (Register, GetUserByEmail).
- [ ] Buat `user_service` untuk validasi registrasi dan verifikasi login.
- [ ] Buat `user_handler` untuk HTTP endpoint POST `/register` dan POST `/login`.
- [ ] Daftarkan endpoint user ke Chi Router di `main.go` atau file setup router terpisah.

### Tahap 5: Fitur Manajemen Event (Admin/Organizer)
- [ ] Buat `event_repository` untuk CRUD data event di MySQL.
- [ ] Buat `event_service` untuk logika pembuatan dan pengelolaan event.
- [ ] Buat `event_handler` untuk HTTP endpoint (contoh: POST `/events`, GET `/events`).
- [ ] Proteksi endpoint POST `/events` menggunakan middleware Auth JWT yang sudah dibuat.

### Tahap 6: Fitur Pembelian Tiket (User)
- [ ] Buat `ticket_repository` untuk transaksi pembelian dan pengecekan kuota tiket.
- [ ] Buat `ticket_service` yang berisi logika validasi sisa tiket (concurrency/race condition handling disarankan) dan pembuatan tiket user.
- [ ] Buat `ticket_handler` untuk POST `/tickets/buy` atau POST `/events/{id}/buy`.
- [ ] Pastikan endpoint pembelian tiket dilindungi dengan middleware Auth JWT.

### Tahap 7: Testing & Finalisasi
- [ ] Uji coba semua endpoint (Register, Login, Create Event, Buy Ticket) menggunakan Postman, cURL, atau swagger.
- [ ] (Opsional) Tambahkan *Graceful Shutdown* di `cmd/server/main.go`.
- [ ] Update README.md dengan instruksi cara menjalankan project lokal.
