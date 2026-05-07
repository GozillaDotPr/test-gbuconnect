
## Requirements

- PHP 8.0+
- Composer
- MongoDB (local or Atlas)
- PHP MongoDB extension (`ext-mongodb`)

## Install

```bash
# Copy env file
cp .env.example .env

# Edit .env sesuai kebutuhan
nano .env

# Install dependencies
composer install
```

## Run

```bash
php -S localhost:8000 -t public


## API Endpoints

### Login

```bash
curl -X POST http://localhost:8000/login \
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

### Gunakan Token

Semua endpoint `/products` wajib menyertakan header:
```
Authorization: Bearer <token>
```

### Create Product

```bash
curl -X POST http://localhost:8000/products \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","desc":"Laptop gaming","price":15000000}'
```

### Get All Products

```bash
curl http://localhost:8000/products \
  -H "Authorization: Bearer <token>"
```

### Get Product by ID

```bash
curl http://localhost:8000/products/{id} \
  -H "Authorization: Bearer <token>"
```

### Update Product

```bash
curl -X PUT http://localhost:8000/products/{id} \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop Pro","price":18000000}'
```

### Delete Product

```bash
curl -X DELETE http://localhost:8000/products/{id} \
  -H "Authorization: Bearer <token>"
```

## Project Structure

```
app/
├── Handlers/       # Handle HTTP request & response
├── Services/       # Business logic
├── Repositories/   # Database queries
├── Models/         # Data models
├── Middleware/     # JWT middleware
└── Routes/         # Route definitions
public/
└── index.php       # Entry point
```
