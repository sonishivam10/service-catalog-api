# 💠 Service Catalog API

A production-ready, containerized Go microservice for managing services and their versions. Built with PostgreSQL, JWT auth, and Swagger.

---

## 🚀 Features

* RESTful API (CRUD for services & versions)
* Pagination, sorting, filtering
* JWT-based authentication
* OpenAPI docs with Swagger UI
* SQL migrations and seed data
* Dockerized setup (Postgres + Go app)
* Graceful shutdown, health check
* Makefile automation

---

## ⚙️ Tech Stack

* **Golang 1.24+**
* **PostgreSQL 15**
* **Docker + Compose**
* **JWT (HS256)**
* **Swaggo** (OpenAPI generation)
* **sqlx + pq** (DB access)

---

## 📦 Project Structure

```
.
├── cmd/server/         # Main Go entrypoint
├── internal/           # Handlers, services, repos, middleware
├── migrations/         # SQL migration files
├── scripts/            # Seeder and token generator
├── docs/               # Auto-generated Swagger files
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 🥪 Local Setup (1-Step)

Make sure you have:

* Docker & Docker Compose
* Go 1.24+

```bash
make setup
```

This will:

* Build and start Postgres
* Run DB migrations
* Seed with 10 services and versions
* Start the Go API

---

## 🔐 Auth & Tokens

This API uses **JWT** via `Authorization: Bearer <token>`.

### ✅ Generate a Token:

```bash
make token
```

Copy the output and use it in:

* Swagger Authorize
* `curl` requests

---

## 📘 Swagger Docs

```http
http://localhost:8080/swagger/index.html
```

---

## 📈 Health Check

```http
GET /healthz → 200 OK
```

---

## 🧰 Dev Commands

```bash
make setup       # Full startup: db + migrations + seed + app
make stop        # Stop all containers
make token       # Generate JWT token
make swag        # Regenerate Swagger docs
make run         # Run Go app locally (not in Docker)
make test        # Run Go tests
```

---

## 🧼 Clean Up

```bash
make stop
docker volume rm service_catalog_pgdata
```

---

## 💡 Design Considerations & Trade-offs

* **Modularity**: The app is broken into `handler`, `service`, and `repository` layers to follow clean architecture principles. This separation simplifies testing and future extension.
* **Postgres in Docker**: Chosen for portability and consistency across environments.
* **Seed and Migration Ordering**: Managed via `depends_on` and health checks in Docker Compose to ensure the database is ready before seeding(idempotent).
* **Swagger over Postman**: Swagger UI provides both documentation and an interactive way to test APIs.
* **JWT over sessions**: Simpler for stateless microservices and scales better in distributed systems.
* **Not all CRUD covered yet**: For simplicity, only read endpoints were initially implemented : creation and deletion can be added in a real-world build.
* **No ORM**: Uses `sqlx` for lighter version and explicit SQL over heavier ORMs like GORM.
* **Minimal dependencies**: To keep the image lean, everything is built into Alpine containers.

---
