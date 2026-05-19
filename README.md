# Product Requirements Document (PRD)
## Portal Informasi Akademik Kampus

| | |
|---|---|
| **Versi** | 1.0.0 |
| **Tanggal** | 19 Mei 2026 |
| **Status** | Draft |
| **Tim** | 11 Mahasiswa |

---

## 1. Ringkasan Produk

**Portal Informasi Akademik Kampus** adalah aplikasi web fullstack berbasis REST API yang memudahkan pengelolaan data akademik kampus. Aplikasi ini dibangun menggunakan **Go Fiber** sebagai backend dan **MongoDB** sebagai database, di-deploy di **Alwaysdata** dengan CI/CD otomatis via GitHub Actions.

---

## 2. Tujuan

- Menyediakan sistem pengelolaan data akademik yang terpusat
- Menjadi media latihan implementasi REST API + Frontend secara end-to-end
- Setiap mahasiswa berkontribusi membangun modul secara mandiri namun tetap terintegrasi dalam satu aplikasi

---

## 3. Tech Stack

| Layer | Teknologi |
|---|---|
| Backend | Go + Go Fiber v2 |
| Database | MongoDB Atlas |
| Frontend | HTML + CSS + JS (vanilla) |
| Hosting | Alwaysdata (free for life) |
| CI/CD | GitHub Actions |
| Boilerplate | github.com/gocroot/alwaysdata |

---

## 4. Arsitektur Aplikasi

```
GitHub Repo (per mahasiswa, nama = NPM)
        │
        │  git push → GitHub Actions
        │
        ▼
  Alwaysdata Hosting
  ┌─────────────────────────────────┐
  │  Go Fiber App                   │
  │  ┌──────────┐  ┌─────────────┐ │
  │  │  url/    │→ │ controller/ │ │
  │  └──────────┘  └──────┬──────┘ │
  │                        │        │
  │                 ┌──────▼──────┐ │
  │                 │   helper/   │ │
  │                 └──────┬──────┘ │
  │                        │        │
  │                 ┌──────▼──────┐ │
  │                 │   MongoDB   │ │
  │                 └─────────────┘ │
  └─────────────────────────────────┘
```

---

## 5. Fitur Aplikasi

### 5.1 Fitur Global (semua mahasiswa wajib implementasi)

| Fitur | Endpoint | Deskripsi |
|-------|----------|-----------|
| Homepage | `GET /` | Menampilkan nama aplikasi dan status server |
| IP Server | `GET /ip` | Menampilkan IP address server yang sedang berjalan |
| CORS | — | Semua endpoint bisa diakses dari frontend |
| Auto-deploy | — | Setiap `git push` ke `main` otomatis deploy ke Alwaysdata |

---

### 5.2 Fitur Per Modul

Setiap mahasiswa membangun **2 modul Backend + 1 menu Frontend** yang saling berkaitan.

---

#### Modul 1 — Data Mahasiswa & Autentikasi
**Deskripsi:** Mengelola data profil mahasiswa dan sistem autentikasi berbasis nomor telepon.

Backend 1 — **Mahasiswa**
- `GET /mahasiswa` → ambil semua data mahasiswa
- `GET /mahasiswa/:npm` → ambil data mahasiswa berdasarkan NPM
- `POST /mahasiswa` → tambah data mahasiswa baru
- `PUT /mahasiswa/:npm` → update data mahasiswa
- `DELETE /mahasiswa/:npm` → hapus data mahasiswa

Backend 2 — **Auth**
- `POST /auth/login` → login menggunakan nomor telepon, return token
- `GET /auth/profile/:phone` → ambil profil berdasarkan nomor telepon

Frontend — **Halaman Data Mahasiswa**
- Tabel list semua mahasiswa
- Form tambah mahasiswa baru
- Tombol hapus data

---

#### Modul 2 — Data Dosen & Jabatan
**Deskripsi:** Mengelola data dosen beserta jabatan fungsional dan struktural.

Backend 1 — **Dosen**
- `GET /dosen` → ambil semua data dosen
- `GET /dosen/:nidn` → ambil data dosen berdasarkan NIDN
- `POST /dosen` → tambah data dosen baru
- `PUT /dosen/:nidn` → update data dosen
- `DELETE /dosen/:nidn` → hapus data dosen

Backend 2 — **Jabatan**
- `GET /jabatan` → ambil semua jabatan
- `POST /jabatan` → tambah jabatan baru
- `GET /jabatan/:id` → detail jabatan

Frontend — **Halaman Data Dosen**
- Tabel list semua dosen beserta jabatannya
- Form tambah dosen baru
- Filter berdasarkan jabatan

---

#### Modul 3 — Mata Kuliah & KRS
**Deskripsi:** Mengelola data mata kuliah dan pengambilan KRS oleh mahasiswa.

Backend 1 — **Mata Kuliah**
- `GET /matkul` → ambil semua mata kuliah
- `GET /matkul/:kode` → ambil detail mata kuliah berdasarkan kode
- `POST /matkul` → tambah mata kuliah baru
- `PUT /matkul/:kode` → update mata kuliah
- `DELETE /matkul/:kode` → hapus mata kuliah

Backend 2 — **KRS**
- `GET /krs/:npm` → ambil KRS mahasiswa berdasarkan NPM
- `POST /krs` → daftarkan mata kuliah ke KRS
- `DELETE /krs/:id` → batalkan KRS

Frontend — **Halaman Mata Kuliah**
- Tabel list mata kuliah (kode, nama, SKS, semester)
- Form tambah mata kuliah
- Form input KRS mahasiswa

---

#### Modul 4 — Jadwal & Ruangan
**Deskripsi:** Mengelola jadwal perkuliahan dan data ruangan yang tersedia.

Backend 1 — **Jadwal**
- `GET /jadwal` → ambil semua jadwal kuliah
- `GET /jadwal/:id` → detail jadwal
- `POST /jadwal` → tambah jadwal baru
- `PUT /jadwal/:id` → update jadwal
- `DELETE /jadwal/:id` → hapus jadwal

Backend 2 — **Ruangan**
- `GET /ruangan` → ambil semua ruangan
- `POST /ruangan` → tambah ruangan baru
- `GET /ruangan/:kode` → cek ketersediaan ruangan
- `PUT /ruangan/:kode` → update data ruangan

Frontend — **Halaman Jadwal Kuliah**
- Tampilan jadwal per hari/minggu
- Filter jadwal berdasarkan prodi atau dosen
- Info ruangan yang digunakan

---

#### Modul 5 — Nilai & Transkrip
**Deskripsi:** Mengelola input nilai mahasiswa dan rekap transkrip akademik.

Backend 1 — **Nilai**
- `GET /nilai/:npm` → ambil semua nilai mahasiswa berdasarkan NPM
- `POST /nilai` → input nilai mahasiswa
- `PUT /nilai/:id` → update nilai
- `DELETE /nilai/:id` → hapus nilai

Backend 2 — **Transkrip**
- `GET /transkrip/:npm` → ambil rekap seluruh nilai dan total SKS mahasiswa
- `GET /transkrip/:npm/ipk` → hitung dan return nilai IPK

Frontend — **Halaman Input Nilai**
- Form input nilai per mahasiswa per mata kuliah
- Tabel rekap nilai dengan IPK
- Filter berdasarkan semester

---

#### Modul 6 — Absensi & Rekap
**Deskripsi:** Mengelola data kehadiran mahasiswa dan rekap absensi per periode.

Backend 1 — **Absensi**
- `GET /absensi/:npm` → ambil absensi mahasiswa
- `POST /absensi` → input absensi mahasiswa
- `PUT /absensi/:id` → update status absensi
- `GET /absensi/hari-ini` → absensi yang masuk hari ini

Backend 2 — **Rekap Absensi**
- `GET /rekap-absensi/:npm` → rekap persentase kehadiran per matkul
- `GET /rekap-absensi/matkul/:kode` → rekap semua mahasiswa dalam satu matkul

Frontend — **Halaman Form Absensi**
- Form input absensi dengan tanggal dan status (hadir/izin/alpha)
- Tabel rekap persentase kehadiran per mahasiswa

---

#### Modul 7 — Pengumuman & Kategori
**Deskripsi:** Papan pengumuman digital untuk informasi kampus dengan sistem kategori.

Backend 1 — **Pengumuman**
- `GET /pengumuman` → ambil semua pengumuman (terbaru di atas)
- `GET /pengumuman/:id` → detail pengumuman
- `POST /pengumuman` → tambah pengumuman baru
- `PUT /pengumuman/:id` → update pengumuman
- `DELETE /pengumuman/:id` → hapus pengumuman

Backend 2 — **Kategori**
- `GET /kategori` → ambil semua kategori
- `POST /kategori` → tambah kategori baru
- `GET /pengumuman/kategori/:nama` → filter pengumuman berdasarkan kategori

Frontend — **Halaman Board Pengumuman**
- Tampilan card pengumuman terbaru
- Filter berdasarkan kategori
- Form tambah pengumuman

---

#### Modul 8 — Beasiswa & Pendaftaran
**Deskripsi:** Informasi beasiswa yang tersedia dan sistem pendaftaran beasiswa.

Backend 1 — **Beasiswa**
- `GET /beasiswa` → ambil semua beasiswa
- `GET /beasiswa/:id` → detail beasiswa
- `POST /beasiswa` → tambah data beasiswa
- `PUT /beasiswa/:id` → update beasiswa
- `DELETE /beasiswa/:id` → hapus beasiswa

Backend 2 — **Pendaftaran Beasiswa**
- `POST /beasiswa/daftar` → daftarkan mahasiswa ke beasiswa
- `GET /beasiswa/pendaftar/:id` → lihat daftar pendaftar beasiswa
- `GET /beasiswa/status/:npm` → cek status pendaftaran beasiswa mahasiswa

Frontend — **Halaman List Beasiswa**
- Tabel list beasiswa (nama, syarat, deadline)
- Form pendaftaran beasiswa
- Status pendaftaran per mahasiswa

---

#### Modul 9 — Perpustakaan & Peminjaman
**Deskripsi:** Katalog buku perpustakaan dan sistem peminjaman buku oleh mahasiswa.

Backend 1 — **Buku**
- `GET /buku` → ambil semua buku
- `GET /buku/:id` → detail buku
- `GET /buku/cari?judul=` → cari buku berdasarkan judul
- `POST /buku` → tambah buku baru
- `PUT /buku/:id` → update data buku

Backend 2 — **Peminjaman**
- `POST /peminjaman` → pinjam buku
- `PUT /peminjaman/:id/kembali` → kembalikan buku
- `GET /peminjaman/:npm` → riwayat peminjaman mahasiswa
- `GET /peminjaman/aktif` → daftar buku yang sedang dipinjam

Frontend — **Halaman Cari Buku**
- Search bar pencarian buku
- Tabel hasil pencarian dengan status ketersediaan
- Form peminjaman buku

---

#### Modul 10 — Prestasi & Kategori
**Deskripsi:** Pencatatan prestasi mahasiswa beserta kategori jenis prestasi.

Backend 1 — **Prestasi**
- `GET /prestasi` → ambil semua prestasi
- `GET /prestasi/:npm` → prestasi mahasiswa tertentu
- `POST /prestasi` → input prestasi baru
- `PUT /prestasi/:id` → update prestasi
- `DELETE /prestasi/:id` → hapus prestasi

Backend 2 — **Kategori Prestasi**
- `GET /kategori-prestasi` → ambil semua kategori (akademik, non-akademik, dll)
- `POST /kategori-prestasi` → tambah kategori baru
- `GET /prestasi/kategori/:nama` → filter prestasi berdasarkan kategori

Frontend — **Halaman Input Prestasi**
- Form input prestasi (nama event, tingkat, juara, tanggal)
- Tabel list prestasi mahasiswa dengan filter kategori

---

#### 🎓 Modul 11 — Alumni & Lowongan Kerja
**Deskripsi:** Data alumni kampus dan informasi lowongan kerja yang relevan.

Backend 1 — **Alumni**
- `GET /alumni` → ambil semua data alumni
- `GET /alumni/:npm` → detail alumni
- `POST /alumni` → tambah data alumni baru
- `PUT /alumni/:npm` → update data alumni
- `GET /alumni/angkatan/:tahun` → filter alumni berdasarkan angkatan

Backend 2 — **Lowongan Kerja**
- `GET /lowongan` → ambil semua lowongan
- `GET /lowongan/:id` → detail lowongan
- `POST /lowongan` → tambah lowongan baru
- `PUT /lowongan/:id` → update lowongan
- `DELETE /lowongan/:id` → hapus lowongan

Frontend — **Halaman Data Alumni & Lowongan**
- Tabel data alumni dengan info pekerjaan
- List lowongan kerja terbaru
- Filter lowongan berdasarkan bidang

---

## 6. Pembagian Tugas (11 Mahasiswa)

| No | Modul BE 1 | Modul BE 2 | Menu FE |
|----|-----------|-----------|---------|
| 1 | Mahasiswa (CRUD) | Auth (login phone) | Halaman Data Mahasiswa |
| 2 | Dosen (CRUD) | Jabatan (CRUD) | Halaman Data Dosen |
| 3 | Mata Kuliah (CRUD) | KRS (daftar matkul) | Halaman Mata Kuliah |
| 4 | Jadwal (CRUD) | Ruangan (CRUD) | Halaman Jadwal Kuliah |
| 5 | Nilai (CRUD) | Transkrip (rekap + IPK) | Halaman Input Nilai |
| 6 | Absensi (CRUD) | Rekap Absensi | Halaman Form Absensi |
| 7 | Pengumuman (CRUD) | Kategori (CRUD) | Halaman Board Pengumuman |
| 8 | Beasiswa (CRUD) | Pendaftaran Beasiswa | Halaman List Beasiswa |
| 9 | Buku/Perpustakaan (CRUD) | Peminjaman Buku | Halaman Cari Buku |
| 10 | Prestasi (CRUD) | Kategori Prestasi | Halaman Input Prestasi |
| 11 | Alumni (CRUD) | Lowongan Kerja (CRUD) | Halaman Alumni & Lowongan |

---

## 7. Struktur Folder Per Mahasiswa

Mengikuti struktur boilerplate `gocroot/alwaysdata`:

```
[NPM]/
├── .github/
│   └── workflows/
│       └── alwaysdata.yml       ← CI/CD config
├── config/
│   ├── api.go
│   ├── config.go
│   ├── cors.go
│   ├── db.go
│   └── token.go
├── controller/
│   ├── controller.go            ← homepage, ip
│   └── [modul1].go              ← modul BE 1 (buat sendiri)
│   └── [modul2].go              ← modul BE 2 (buat sendiri)
├── model/
│   ├── model.go
│   └── [modul].go               ← struct model (buat sendiri)
├── helper/
│   ├── helper.go
│   └── mongodb.go
├── url/
│   └── url.go                   ← daftarkan semua route di sini
├── frontend/
│   └── index.html               ← 1 menu frontend
├── main.go
├── go.mod
└── .env
```

---

## 8. Standar REST API

Setiap modul wajib mengikuti standar berikut:

### Request & Response

```
GET    /[resource]          → ambil semua data
GET    /[resource]/:id      → ambil satu data
POST   /[resource]          → tambah data baru
PUT    /[resource]/:id      → update data
DELETE /[resource]/:id      → hapus data
```

### Format Response Sukses
```json
{
  "status": "success",
  "data": { ... }
}
```

### Format Response Error
```json
{
  "status": "error",
  "message": "deskripsi error"
}
```

---

## 9. Flow Pengerjaan

```
1. Setup Awal
   ├── Fork/clone boilerplate gocroot/alwaysdata
   ├── Push ke repo NPM di organisasi GitHub
   ├── Daftar Alwaysdata → dapat subdomain [nama].alwaysdata.net
   ├── Setup GitHub Secrets (sshhost, sshpassword, dll)
   └── Test CI/CD: push → cek apakah auto-deploy berjalan

2. Bangun Backend
   ├── Buat model struct di model/[modul].go
   ├── Buat controller di controller/[modul].go
   ├── Daftarkan route di url/url.go
   └── Test endpoint dengan Postman / curl

3. Bangun Frontend
   ├── Buat file frontend/index.html
   ├── Koneksikan ke endpoint BE via fetch/axios
   └── Tampilkan data dalam tabel atau form

4. Finishing
   ├── Dokumentasi API (endpoint, method, contoh request & response)
   ├── Debugging & fix bug
   └── Push final → pastikan live di Alwaysdata
```

---

## 10. GitHub Secrets yang Dibutuhkan

| Secret | Keterangan |
|--------|-----------|
| `sshhost` | SSH host Alwaysdata |
| `sshusername` | Username SSH |
| `sshpassword` | Password SSH (strong password) |
| `sshport` | Port SSH (default 22) |
| `apikey` | API key dari Alwaysdata dashboard |
| `appid` | App ID dari menu Web > Sites |
| `folder` | Path folder home di Alwaysdata |

---

## 11. Environment Variables

```sh
MONGOSTRING=mongodb+srv://user:pass@cluster.mongodb.net/
PORT=porthost
IP=iphost
```

---

## 12. Kriteria Selesai (Definition of Done)

- [ ] App berhasil di-deploy di Alwaysdata dan bisa diakses via URL
- [ ] CI/CD berjalan: setiap push ke `main` otomatis deploy
- [ ] Minimal **2 endpoint BE** per mahasiswa berjalan (GET + POST)
- [ ] **1 menu FE** bisa menampilkan dan/atau mengirim data ke BE
- [ ] Dokumentasi API tersedia (minimal dalam README.md)
- [ ] Link live app + repo dikumpulkan di folder NPM dalam organisasi GitHub

---

## 13. Referensi

- Boilerplate: https://github.com/gocroot/alwaysdata
- Go Fiber Docs: https://gofiber.io/
- MongoDB Atlas: https://www.mongodb.com/atlas
- Alwaysdata: https://www.alwaysdata.com
- GitHub Actions Docs: https://docs.github.com/actions
