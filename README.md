# 🥷 Samurai API

A RESTful API for managing samurai clans and their warriors — built with Go, MongoDB, and tested using Testcontainers.  
Honor flows through the pipeline with GitHub Actions ⚔️

---

## 📦 Features

- 🏯 Clan Management – Add and list samurai clans
- 🥷 Samurai Management – Register warriors under clans
- ✅ Integration Tests – Real MongoDB with Testcontainers
- 🔁 Layered Architecture – Handler → Service → Repository
- 🔄 CI Pipeline – Runs on every push via GitHub Actions

---

## 🚀 Tech Stack

- Go 1.21
- MongoDB
- Testcontainers-Go
- GitHub Actions


---

## 🛠️ Getting Started

```bash
git clone https://github.com/reyyansarikaya/samurai-api.git
cd samurai-api
go mod tidy
``` 

### Start MongoDB (via Docker)
```bash
docker run --name samurai-mongo -d -p 27017:27017 mongo:6
``` 
### Run the API
```bash
go run main.go
```
### Run Tests
```bash
go test ./...
```


## 🔁 CI Pipeline
GitHub Actions automatically:
- Installs dependencies
- Starts MongoDB via Testcontainers
- Runs all tests on every push or PR

## 📬 API Endpoints
Clans
-	POST /clans – Add a new clan
-	GET /clans – List all clans

Samurais
-	POST /samurais – Register a new samurai
-	GET /samurais – List all samurais
    
