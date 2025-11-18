---

# 📘 **Dependency Injection (DI)**

## **🔹 Alur utama arsitektur:**

```
Routes → Controller → Service → Repository → Database
```

---

# 🧩 **Penjelasan Setiap Layer Dengan Analogi**

### **1. Routes → “Polisi Lalu Lintas”**

* Tugas: hanya mengarahkan request ke controller yang benar.
* Tidak punya logika bisnis.
* Tidak perlu tahu database, service, dll.

➡ Kalau ada request **POST /register**, routes cuma bilang:
“Supir, kamu yang handle!”

---

### **2. Controller → “Supir”**

* Tugas: menerima request dari routes.
* Mengemudikan alur proses.
* Memanggil service untuk menjalankan logika.
* Menyiapkan response HTTP untuk dikirim kembali ke client.

➡ Controller = mengarahkan alur, bukan menjalankan logika berat.

---

### **3. Service → “Mesin”**

* Tugas: tempat semua **logika bisnis** berlangsung.
* Validasi bisnis tambahan.
* Proses data.
* Memanggil repository untuk akses database.

➡ Mesin = tempat kerja berat.

---

### **4. Repository → “Bensin”**

* Tugas: berkomunikasi langsung dengan database.
* Menjalankan query GORM.
* CRUD data.

➡ Repository = sumber bahan bakar (data).

---

### **5. DB (Database) → “Stasiun Bensin”**

* Tempat data disimpan.
* Tidak tahu apa-apa soal service atau controller.

➡ DB hanya menyediakan “bahan bakar”.

---

# 💡 **Kenapa `UserRoutes` menyimpan controller?**

Karena Routes hanya butuh 1 hal:

> “Ketika user memanggil endpoint X, jalankan fungsi controller ini.”

Routes **tidak butuh DB, service, repository.**
Controller-lah yang mengarahkan semuanya.

---

# 💡 **Kenapa DI (Dependency Injection) seperti ini?**

Di `NewUserRoutes`:

```
DB → Repository → Service → Controller → Routes
```

Injeksi dependency:

* DB masuk ke repository
* Repository ke service
* Service ke controller
* Controller ke routes

Sehingga semua bagian terhubung **rapi dan modular**.

---

# 💡 **Kenapa Routes cuma punya controller?**

Karena Routes:

* Tidak menjalankan bisnis logic
* Tidak menjalankan database
* Tidak memproses data

Routes hanya mapping:

```
POST /register → controller.Register
```

Makanya struct-nya:

```go
type UserRoutes struct {
    controller *UserController
}
```

---

