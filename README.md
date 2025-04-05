# 🥷 Samurai API

A RESTful API for managing samurai clans and their warriors — built with Go, MongoDB, and tested using Testcontainers.  
Honor flows through the pipeline with GitHub Actions ⚔️

---

## 📦 Features

- 🏯 Clan Management – Add and list samurai clans
- 🥷 Samurai Management – Register warriors under clans
- ✅ Integration Tests – Real MongoDB with Testcontainers
- 🐳 Dockerized – Build & run with Docker or Docker Compose
- 🔁 Layered Architecture – Handler → Service → Repository
- 🔄 CI Pipeline – GitHub Actions: test & build automation

---

## 🚀 Tech Stack

- Go 1.21+
- MongoDB 6
- Testcontainers-Go
- GitHub Actions
- Docker & Docker Compose

---

## 🛠️ Getting Started

### 🔃 Clone & Prepare
```bash
git clone https://github.com/reyyansarikaya/samurai-api.git
cd samurai-api
go mod tidy
```

---

### 🐳 Option 1: Run with Docker Compose
Spin up both MongoDB and Samurai API together:

```bash
docker compose up --build
```

> API will be available at: `http://localhost:1600`

---

### ⚙️ Option 2: Manual Run (Mongo via Docker)

Start MongoDB:
```bash
docker run --name samurai-mongo -d -p 27017:27017 mongo:6
```

Run the API:
```bash
go run main.go
```

---

### 🧪 Run Tests

```bash
go test ./...
```

---

## 📬 API Endpoints

### 📁 Clans

- `POST /clans` – Add a new clan
- `GET /clans` – List all clans

### 🥷 Samurais

- `POST /samurais` – Register a new samurai
- `GET /samurais` – List all samurais

---

## 🔄 CI Pipeline

GitHub Actions will automatically:
- Install dependencies
- Build the project
- Run integration tests with real MongoDB (via Testcontainers)
- Fail on errors or failed assertions

---

## 📁 Project Structure (Simplified)

```
.
├── main.go
├── Dockerfile
├── docker-compose.yml
├── internal/
│   └── banner/
├── db/
├── handlers/
├── models/
├── repository/
├── service/
├── tests/
└── vendor/
```

---

## 👤 Author

Made with discipline by [@reyyansarikaya](https://github.com/reyyansarikaya) 🥷
