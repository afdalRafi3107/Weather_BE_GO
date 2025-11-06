# Weather API Backend (Go)

Backend project untuk aplikasi **Weather App** yang dibangun menggunakan **Golang**, framework **Gin Gonic**, dan konfigurasi environment dengan **Viper**.

---

## 🚀 Tech Stack

| Category | Tech / Library |
|----------|----------------|
| Language | Go 1.24.5 |
| Web Framework | Gin Gonic |
| Environment Config | Viper, Gotenv |
| JSON Handling | go-json, sonic, json-iterator |
| Validation | go-playground/validator.v10 |
| CORS | gin-contrib/cors |
| Logging / Utils | go.uber.org/atomic, multierr |
| Build & Run | `go run`, `go build` |

---

## 📂 Project Structure

```
backend/
├── controllers/            # Handler logic
│   └── weather.go
├── models/                 # Data structures / DTO
│   └── weather.go
├── tmp/                    # Build binary & logs
│   ├── build-errors.log
│   └── main.exe
├── utils/                  # Helpers & config loader
│   └── config.go
├── .env                    # Environment variables
├── go.mod                  # Module dependencies
├── go.sum
├── main.go                 # App entry point (server)
└── README.md
```

---

## 📥 Cara Clone & Menjalankan Project

### 1️⃣ Clone Repository

```sh
git clone https://github.com/afdalRafi3107/Weather_BE_GO.git
cd Weather_BE_GO/backend
```

### 2️⃣ Install Dependencies

```sh
go mod tidy
```

### 3️⃣ Setup Environment

Buat file `.env` di root folder:

```
API_KEY=your_openweather_api_key
PORT=8080
```

### 4️⃣ Jalankan Server

```sh
go run main.go
```

Atau **build executable**:

```sh
go build -o weather-api
./weather-api
```

Server akan berjalan di:
```
http://localhost:8080
```

---

## ✅ Example Endpoint

| Method | Route | Description |
|--------|-------|-------------|
| `GET` | `/weather/:city` | Get weather data by city |

Contoh request:
```
GET http://localhost:8080/weather/Jakarta
```

---

## 📌 Requirements

- Go **v1.20+**
- Git
- API Key (ex: OpenWeatherMap)

---

## 📄 License

MIT License – open source & free to modify.

---

## 🤝 Contributing

Silakan fork, buat branch baru, dan kirim pull request.

---

Jika ingin ditambahkan contoh response JSON, dokumentasi Swagger, atau screenshot Postman, cukup beri tahu saya 👍

