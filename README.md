# ⏳ Goroutine Queue Simulation in Go

Simulasi program **antrian (queue)** menggunakan **goroutine** dengan waktu tunggu acak (random delay) antara 1–3 detik untuk setiap proses. Program ini juga menggunakan `sync.WaitGroup` agar semua proses selesai sebelum program berakhir.

---

## 🚀 Fitur

- Menggunakan **goroutine** untuk menjalankan setiap queue secara paralel.
- **Waktu delay acak** pada tiap queue menggunakan `rand.Intn`.
- **Sinkronisasi** dengan `sync.WaitGroup`.
- Menampilkan output setelah setiap queue selesai disertai waktu tunggunya.

---

## 🛠️ Prerequisites

- Go has been installed (minimum version 1.20+)
- Git (optional, to clone the repo)

Check Go installation:

## How to Run this project

1. Clone this project

```

git clone https://github.com/Dwaysetya/fgo24-go-json-unmarshal

```

2. Enter the project firectory

```

cd fgo24-go-json-unmarshal

```

3. run the project

```

go run main.go // go run .

```

## 📦 Dependencies

This project uses only the Go standard library.
No additional libraries or frameworks are required.

Make sure Go is installed on your system.

## 📚 Basic Information

This project was developed as part of the learning program at

**Kodacademy Bootcamp Batch 2**.

Created by **Dwi Setyabudi** to deepen understanding and improve technical skills in the Go programming language.
