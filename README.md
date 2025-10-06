# 🛒 E-Commerce API (Golang + Gin + MySQL)

### 👨‍💻 Developer: Pawit  
Membangun backend sederhana untuk sistem e-commerce dengan fitur autentikasi, produk, kategori, dan pesanan.

---

## 🧱 1. Arsitektur & Teknologi

**Backend:** Go (Gin Framework)  
**Database:** MySQL (hosted via Railway)  
**ORM:** GORM  
**Auth:** JSON Web Token (JWT)  
**Deployment:** Railway App  

### 📂 Struktur Proyek
controllers/ → logic untuk endpoint
models/ → definisi tabel database
routes/ → daftar route API
middlewares/ → JWT & role check
config/ → koneksi database
seeders/ → data awal
main.go → entry point


---

## 🔑 2. Fitur Autentikasi

### 🔹 Register
**Endpoint:** `POST /api/users/register`  
➡️ Pengguna baru dapat mendaftar akun.

### 🔹 Login
**Endpoint:** `POST /api/users/login`  
➡️ Mendapatkan JWT Token untuk autentikasi.

### 🔹 Profile
**Endpoint:** `GET /api/users/me`  
➡️ Menampilkan informasi user yang sedang login.

**Header Authorization:**
Authorization: Bearer <token>


---

## 🛍️ 3. Fitur Produk & Kategori

### Public Routes (tanpa login)
- `GET /api/products` → Lihat semua produk  
- `GET /api/categories` → Lihat semua kategori  

### Admin Routes
- `POST /api/admin/products` → Tambah produk  
- `PUT /api/admin/products/:id` → Update produk  
- `DELETE /api/admin/products/:id` → Hapus produk  
- `POST /api/admin/categories` → Tambah kategori  
- `PUT /api/admin/categories/:id` → Update kategori  
- `DELETE /api/admin/categories/:id` → Hapus kategori  

📌 **Catatan:** Semua route admin hanya bisa diakses oleh user dengan role `admin`.

---

## 🧾 4. Fitur Order (Pesanan)

### Customer
- `POST /api/orders` → Buat pesanan baru  
- `GET /api/orders` → Lihat daftar pesanan sendiri  
- `GET /api/orders/:id` → Detail pesanan  

### Admin
- `GET /api/admin/orders` → Lihat semua pesanan  
- `PUT /api/admin/orders/:id/status` → Ubah status pesanan  

📊 **Status Pesanan:** `pending`, `completed`, `canceled`

---

## 🔐 5. Middleware & Keamanan

- **JWTAuthMiddleware:** Validasi token setiap request.  
- **Role-Based Access:** Pemisahan hak akses admin & customer.  
- **Password Hashing:** Password user disimpan aman di database menggunakan hashing (bcrypt).  

---

## 🚀 6. Deployment

Project ini di-deploy menggunakan **Railway** untuk:
- Hosting backend server  
- Hosting database MySQL  


API siap diuji menggunakan **Postman** melalui URL Railway setelah berhasil deploy.

---

## 🧩 7. Cara Menjalankan di Lokal

### 1️⃣ Clone Repository
```bash
git clone https://github.com/username/ecommerce-api.git
cd ecommerce-api

