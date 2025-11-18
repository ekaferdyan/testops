
```markdown
# 🧠 TestOps Dashboard (Ultimate Edition)

> **TestOps Dashboard** adalah platform analitik untuk mengumpulkan, memvisualisasikan, dan menganalisis hasil pengujian otomatis lintas framework seperti Selenium, Playwright, WebdriverIO, dan Katalon.

User dapat:
1. Memilih framework pengujian yang digunakan.
2. Mengunggah file report hasil testing (format `.json`).
3. Melihat hasil analisis dalam bentuk **grafik & insight visual** yang mudah dipahami oleh QA Lead dan QA Manager.

---

## 🚀 Tech Stack

### 🧩 Backend
- **Language:** Golang  
- **Framework:** Fiber v2  
- **Database:** PostgreSQL  
- **ORM:** GORM  
- **Authentication:** JWT (JSON Web Token)  
- **File Upload:** Local Storage (`/uploads`)  

### 🌐 Frontend
- **Framework:** React (TypeScript)  
- **Styling:** TailwindCSS  
- **HTTP Client:** Axios  
- **Chart:** Recharts  
- **Routing:** React Router DOM  

---

## 📂 Project Structure

```

sambel-ulek/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── connection.go
│   ├── models/
│   │   ├── user.go
│   │   ├── project.go
│   │   └── report.go
│   ├── platform/
│   │   ├── validator
│   │   │  ├── validator.go
│   ├── controllers/
│   │   ├── auth_controller.go
│   │   ├── project_controller.go
│   │   └── report_controller.go
│   ├── routes/
│   │   ├── auth_routes.go
│   │   ├── project_routes.go
│   │   └── report_routes.go
│   ├── middlewares/
│   │   └── jwt_middleware.go
│   ├── services/
│   │   └── jwt_service.go
│   └── uploads/
│       └── (file report json disimpan di sini)
│
├── frontend/
│   ├── package.json
│   ├── tailwind.config.js
│   ├── tsconfig.json
│   ├── public/
│   └── src/
│       ├── App.tsx
│       ├── index.tsx
│       ├── pages/
│       │   ├── Login.tsx
│       │   ├── Dashboard.tsx
│       │   └── UploadReport.tsx
│       ├── components/
│       │   ├── Navbar.tsx
│       │   └── ChartCard.tsx
│       └── services/
│           └── api.ts
│
└── .gitignore

````

---

## 🧱 Database Schema (PostgreSQL)

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  framework VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE test_reports (
  id SERIAL PRIMARY KEY,
  project_id INT REFERENCES projects(id) ON DELETE CASCADE,
  file_path TEXT NOT NULL,
  summary JSONB,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
````

---

## 🧰 Backend Setup (Golang + PostgreSQL)

### 1. Install dependencies

```bash
cd backend
go mod init sambel-ulek/backend
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/fiber/v2/middleware/cors
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

### 2. Setup `.env` (opsional)

Buat file `.env` di `backend/`:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=testops_dashboard
JWT_SECRET=secret123
```

### 3. Jalankan PostgreSQL (contoh Docker)

```bash
docker run --name testops-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testops_dashboard -p 5432:5432 -d postgres
```

### 4. Jalankan server backend

```bash
go run main.go
```

Server akan berjalan di:
👉 [http://localhost:8080](http://localhost:8080)

---

## 🌐 Frontend Setup (React + Tailwind)

### 1. Setup project React

```bash
cd frontend
npm create vite@latest .
# pilih React + TypeScript
npm install
npm install axios react-router-dom recharts tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

### 2. Konfigurasi Tailwind

Edit `tailwind.config.js`:

```js
content: [
  "./index.html",
  "./src/**/*.{js,ts,jsx,tsx}",
],
theme: { extend: {} },
plugins: [],
```

### 3. Jalankan frontend

```bash
npm run dev
```

Frontend berjalan di:
👉 [http://localhost:5173](http://localhost:5173)

---

## 🔑 Authentication Flow

1. User **register** → `/register`

   * Email + password disimpan di database (password di-hash dengan bcrypt).

2. User **login** → `/login`

   * Server memvalidasi password, lalu mengembalikan **JWT Token**.

3. Frontend menyimpan token di `localStorage`.

4. Semua request API berikutnya menyertakan header:

   ```
   Authorization: Bearer <token>
   ```

5. Middleware `JWTProtected` di backend memvalidasi token sebelum mengizinkan akses.

---

## 📤 Upload Report Flow

1. User login, lalu masuk ke halaman **Upload Report**.
2. User memilih framework (misal: `Playwright`) dan mengunggah file `.json`.
3. Backend menyimpan file di folder `/uploads/` dan merekam metadata ke database.
4. Backend membaca isi file JSON, mengekstrak summary hasil testing (passed/failed/etc).
5. Data disimpan dalam kolom `summary` (tipe JSONB).
6. Dashboard menampilkan visualisasi summary per project.

---

## 🧭 API Endpoints

| Method | Endpoint          | Description                     | Auth |
| ------ | ----------------- | ------------------------------- | ---- |
| `POST` | `/register`       | Register user baru              | ❌    |
| `POST` | `/login`          | Login & get JWT token           | ❌    |
| `GET`  | `/projects`       | Ambil daftar project milik user | ✅    |
| `POST` | `/projects`       | Tambah project baru             | ✅    |
| `POST` | `/reports/upload` | Upload file report JSON         | ✅    |
| `GET`  | `/reports`        | Ambil semua report user         | ✅    |

---

## 🧮 Dashboard Preview (React)

Setelah user login, dashboard menampilkan:

* Daftar framework yang digunakan
* Jumlah test yang **passed**, **failed**, dan **skipped**
* Grafik interaktif (menggunakan Recharts)

Contoh tampilan:

```
+------------------------------------+
| Framework: Playwright              |
| Total Test: 52                     |
| Passed: 47 | Failed: 3 | Skipped: 2 |
|                                    |
|   ████░░░░░░  < Bar Chart >        |
+------------------------------------+
```

---

## 🧾 .gitignore

```
# Backend
/backend/uploads
/backend/.env
/backend/go.sum
/backend/go.mod

# Frontend
/frontend/node_modules
/frontend/dist
/frontend/.env

# OS
.DS_Store
```

---

## 🧱 Future Plans (Next Iteration)

* [ ] Upload report langsung dari CI/CD pipeline (GitHub Actions / Jenkins)
* [ ] Integrasi ke Slack / Telegram notification
* [ ] Role-based access control (QA Lead / Manager)
* [ ] Visual Regression tracking
* [ ] Comparison antar-run untuk setiap project

---

## 👨‍💻 Author

**Ferdyan Eka Saputra**
QA/QC Engineer & Test Engineer
Tools: Selenium | WebDriverIO | Mocha | TestNG | Rust | Go | React
🚀 PT Omni Digitama Internusa | Indonesia

---

## ⚖️ License

MIT License © 2025 — TestOps Dashboard Project

```

---

Kalau kamu mau der, saya bisa lanjutkan:
1. Buat **endpoint upload report JSON lengkap (Fiber)**  
2. Tambah **parsing otomatis JSON → summary (passed/failed)**  
3. Dan tampilkan **grafik di Dashboard (React Recharts)**  

Mau saya lanjutkan ke situ ya der (upload + chart)?
```
