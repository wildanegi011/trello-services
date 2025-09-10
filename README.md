# Trello Services

A backend service for managing boards, inspired by Trello.  
Supports board CRUD operations, pagination, and is built with Go, Gin, and GORM.

---

## 🚀 Features

- **Boards:** Retrieve all boards with pagination.
- **Boards:** Fetch a board by its ID.
- **Unit Tests:** Includes tests for board retrieval endpoints.
- **Developer Experience:** Hot reload in development using [Air](https://github.com/air-verse/air).
- **Docker:** Ready for containerized deployment.

---

## 🏁 Getting Started

### Prerequisites

- Go 1.25+
- Docker (for containerized runs)
- PostgreSQL database

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/trello-services.git
   cd trello-services
   ```

2. **Configure environment variables:**
   - Copy `.env.example` to `.env` and update values, or edit `.env` directly.

3. **Run locally (development hot reload):**
   ```sh
   make dev
   ```
   > Requires [Air](https://github.com/air-verse/air):  
   > `go install github.com/air-verse/air@latest`

4. **Run tests:**
   ```sh
   make test
   # or
   go test -v ./...
   ```

5. **Build and run with Docker:**
   ```sh
   make docker-build
   make docker-run
   ```

---

## 🛠 Environment Variables

See `.env` for configuration options:

- `DATABASE_URL` – PostgreSQL connection string
- `SERVER_PORT` – API port (default: 8080)
- `MODE` – Gin mode (`development` or `production`)

---

## 📦 API Endpoints

| Method | Endpoint         | Description                  |
|--------|------------------|-----------------------------|
| GET    | `/api/v1/boards` | List all boards (paginated) |
| GET    | `/api/v1/boards/:id` | Get board by ID         |
| POST   | `/api/v1/boards` | Create a new board          |

---

## 🧪 Testing

Unit tests are provided for handlers and usecases.  
Run all tests with:

```sh
go test -v ./...
```

---

## 📝 Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

---

## 📄 License

MIT

---

## 🤝 Contributing

Pull requests and issues are welcome!