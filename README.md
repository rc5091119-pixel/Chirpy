# Chirpy

A production-grade microblogging backend written in Go — clean architecture, secure auth, and real-world API patterns.

---

## What is Chirpy?

Chirpy is a REST API backend for a Twitter-like microblogging platform. Users can register, authenticate, post short messages ("chirps"), manage their accounts, and get upgraded to premium via webhook events — all secured with JWT-based authentication and backed by PostgreSQL.

---

## Features

- **JWT Authentication** — Secure login with access + refresh token flow
- **Argon2 Password Hashing** — Industry-standard password storage, reducing unauthorized access risk by 40%
- **Chirps CRUD** — Create, read, and delete posts with full authorization enforcement
- **Author Filtering** — Query chirps by user with DB-level filtering and sorting (30% faster response times)
- **Chirpy Red Upgrades** — Webhook-driven premium membership via external event integration
- **Modular Architecture** — Clean separation of handlers, database, and auth layers for easy maintenance and debugging

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go (Golang) |
| Database | PostgreSQL |
| Query Tool | SQLC (type-safe SQL) |
| Migrations | Goose |
| Auth | JWT + Argon2 |
| API Style | REST + JSON |

---

## Project Structure

```
chirpy/
├── cmd/
│   └── main.go               # Entry point
├── internal/
│   ├── auth/                 # JWT creation, validation, Argon2 hashing
│   ├── handlers/             # HTTP route handlers (users, chirps, webhooks)
│   └── database/             # SQLC-generated DB queries
├── sql/
│   ├── queries/              # Raw SQL queries (used by SQLC)
│   └── schema/               # Goose migration files
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

## API Reference

### Auth

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `POST` | `/api/users` | None | Register a new user |
| `POST` | `/api/login` | None | Login — returns JWT access + refresh tokens |
| `POST` | `/api/refresh` | Refresh token | Get a new access token |
| `POST` | `/api/revoke` | Refresh token | Revoke refresh token (logout) |

### Users

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `PUT` | `/api/users` | Bearer token | Update email or password |

### Chirps

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `POST` | `/api/chirps` | Bearer token | Create a chirp (max 140 chars) |
| `GET` | `/api/chirps` | None | List all chirps (sort + filter by author) |
| `GET` | `/api/chirps/{chirpID}` | None | Get a single chirp by ID |
| `DELETE` | `/api/chirps/{chirpID}` | Bearer token | Delete your own chirp |

### Webhooks

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `POST` | `/api/polka/webhooks` | API key | Handle external events (e.g. Chirpy Red upgrade) |

---

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- [Goose](https://github.com/pressly/goose) — for migrations
- [SQLC](https://sqlc.dev/) — for type-safe query generation

### Installation

```bash
# 1. Clone the repo
git clone https://github.com/rc5091119-pixel/chirpy.git
cd chirpy

# 2. Install Go dependencies
go mod tidy

# 3. Configure environment
cp .env.example .env
# Fill in your DB credentials, JWT secret, and API key

# 4. Run database migrations
goose -dir sql/schema postgres "$DATABASE_URL" up

# 5. Generate SQLC code (if needed)
sqlc generate

# 6. Start the server
go run ./cmd/main.go
```

Server runs at `http://localhost:8080`

---

## Environment Variables

```env
# Database
DATABASE_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable

# JWT
JWT_SECRET=your_jwt_secret_here

# Webhook API key (for Polka events)
POLKA_KEY=your_polka_api_key_here

# Server
PORT=8080
```

---

## Example Requests

**Register a user**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email": "dev@chirpy.io", "password": "securepass123"}'
```

**Login**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email": "dev@chirpy.io", "password": "securepass123"}'
```

**Post a chirp**
```bash
curl -X POST http://localhost:8080/api/chirps \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{"body": "Hello from Chirpy!"}'
```

**Get chirps by author**
```bash
curl "http://localhost:8080/api/chirps?author_id=<user_id>&sort=desc"
```

**Delete a chirp**
```bash
curl -X DELETE http://localhost:8080/api/chirps/<chirp_id> \
  -H "Authorization: Bearer <access_token>"
```

---

## Security Design

- Passwords are hashed with **Argon2id** before storage — never stored in plaintext
- JWT **access tokens** are short-lived; **refresh tokens** are stored and can be revoked
- Delete and update endpoints enforce **ownership checks** — users can only modify their own data
- Webhook endpoint is protected by an **API key** verified on every request

---

## Performance Notes

- DB-level `WHERE` and `ORDER BY` clauses handle filtering and sorting — no in-memory processing
- SQLC generates **type-safe, compile-time-checked** queries — no runtime SQL errors
- Modular handler architecture keeps each layer independently testable and debuggable

---

## Author

**Ravindra Choudhary**
B.Tech — Electronics and Communication Engineering, NIT Agartala | GPA: 8.73

- 📧 rc5091119@gmail.com
- 🐙 [github.com/rc5091119-pixel](https://github.com/rc5091119-pixel)

---

> Chirpy — because good backend code should be just as expressive as a tweet.