# Gin + PostgreSQL CRUD API

Simple REST API with JWT auth using Gin, GORM, and PostgreSQL.

## Requirements

- Go 1.22+
- Docker & Docker Compose

## Setup & Run

```bash
# 1. Copy env
cp .env.example .env

# 2. Start PostgreSQL
docker compose up -d

# 3. Install dependencies
go mod tidy

# 4. Run
go run cmd/main.go
```

Server runs at `http://localhost:8080`. GORM will auto-migrate the `products` table on startup.

## Environment Variables

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=products_db

JWT_SECRET=secret123
JWT_EXPIRES=3600

APP_USERNAME=admin
APP_PASSWORD=admin
```

## API Endpoints

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'
```

Response:
```json
{
  "success": true,
  "data": { "token": "eyJ..." }
}
```

### Use Token

Add header to all `/products` requests:
```
Authorization: Bearer <token>
```

### Create Product

```bash
curl -X POST http://localhost:8080/products \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","desc":"Gaming laptop","price":15000000}'
```

### Get All Products

```bash
curl http://localhost:8080/products \
  -H "Authorization: Bearer <token>"
```

### Get Product by ID

```bash
curl http://localhost:8080/products/{id} \
  -H "Authorization: Bearer <token>"
```

### Update Product

```bash
curl -X PUT http://localhost:8080/products/{id} \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop Pro","price":18000000}'
```

### Delete Product

```bash
curl -X DELETE http://localhost:8080/products/{id} \
  -H "Authorization: Bearer <token>"
```

## Project Structure

```
cmd/
└── main.go                  # Entry point
internal/
├── config/                  # Config loader + DB connection
├── container/               # Dependency injection
├── handler/                 # HTTP handlers
├── middleware/              # JWT middleware
├── models/                  # GORM models
├── repository/              # Database queries
├── routes/                  # Route registration
└── service/                 # Business logic
pkg/
└── response/                # Shared response helpers
```
