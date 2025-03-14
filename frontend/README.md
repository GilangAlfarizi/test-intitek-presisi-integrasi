# 🛠️ Full Stack Project (Go + React)

Project ini merupakan aplikasi full stack yang terdiri dari **backend** menggunakan **Go (Gin Framework)** dan **frontend** menggunakan **React (Vite + Flowbite-React)**.

---

## 🚀 **Backend Setup**

Backend menggunakan **Go (Gin Framework)** untuk API server.

### 📌 **Langkah-langkah menjalankan Backend**

1. **Pastikan Go sudah terinstall** di komputer kamu. Jika belum, silakan install dari [golang.org](https://golang.org/dl/).
2. **Buat database kosong** sesuai dengan **nama database** yang terdapat dalam **DSN** di file `main.go`.
3. **Jalankan server backend** dengan perintah:
   cd backend
   go run main.go

> Sebelum menjalankan server backend, Pastikan database kosong sudah dibuat dengan nama yang sesuai pada **dsn** di main.go.
> Server akan berjalan di http://localhost:8080 secara default.

---

## 🚀 **Frontend Setup**

Backend menggunakan **Go (Gin Framework)** untuk API server.

### 📌 **Langkah-langkah menjalankan Frontend**\

1. **Pastikan Node.js sudah terinstall** di komputer kamu. Jika belum, silakan install dari [nodejs.org](https://nodejs.org/).
2. **Jalankan server frontend** dengan perintah:
   cd frontend (masuk ke dalam direktori frontend)
   npm install
   npm run dev

> Pastikan backend sudah berjalan di terminal lain sebelum menjalankan frontend.
> Akses aplikasi di http://localhost:5173.
