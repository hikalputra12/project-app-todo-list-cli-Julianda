# 📋 project-app-todo-list-cli-Julianda

Aplikasi Command Line Interface (CLI) sederhana yang dibangun dengan Go untuk mengelola daftar tugas harian (To-Do List). Aplikasi ini menggunakan file JSON untuk penyimpanan data persisten.

## ✨ Fitur Utama

Aplikasi ini menyediakan fungsionalitas CRUD (Create, Read, Update, Delete) inti untuk mengelola tugas:

* **List Tasks:** Menampilkan semua tugas dalam format tabel yang terstruktur.
* **Create Task:** Menambahkan tugas baru secara interaktif dengan ID otomatis.
* **Update Task:** Memperbarui status atau prioritas tugas yang ada berdasarkan nomor urut.
* **Delete Task:** Menghapus tugas secara permanen berdasarkan nomor urut.
* **Find Task:** Mencari dan menampilkan tugas berdasarkan judulnya.

## ⚙️ Struktur Proyek

Aplikasi ini mengikuti arsitektur berlapis standar (Handler-Service-Data Access):

| Direktori | Tujuan |
| :--- | :--- |
| `cmd/` | Berisi logika **Handler** (I/O dan *output* presentasi). |
| `service/` | Berisi **Logika Bisnis** (`TaskService`) dan validasi. |
| `model/` | Berisi definisi struktur data utama (`Task` struct). |
| `utils/` | Berisi fungsi utilitas (I/O file JSON, konstanta Status/Priority, dan `ReadLine` untuk input aman). |
| `data/` | **Penyimpanan Data** (berisi `todos.json`). |

## 🚀 Memulai Aplikasi

### Prasyarat

Anda memerlukan Go terinstal pada sistem Anda (Go versi 1.25.4 atau yang lebih baru).

### Instalasi

1.  *Clone* repositori ini:
    ```bash
    git clone [Your Repository URL]
    cd project-app-todo-list-cli-Julianda
    ```
2.  Instal dependensi (seperti `go-pretty`):
    ```bash
    go mod tidy
    ```
### Cara Penggunaan (Menu Interaktif)

Aplikasi ini menggunakan *loop* menu interaktif. Cukup jalankan `main.go` untuk menampilkan menu:

```bash
go run main.go
```
### 📺 Tutorial Video 

[Video Tutorial](URL https://drive.google.com/file/d/1q34SSCwPqws3MlAyQy6qKz70rDyYeL6Q/view?usp=drive_link)

