# 🚀 Chirpy - Backend API (Twitter-like Service)

## 📌 What This Project Does

Chirpy is a backend REST API for a microblogging platform (similar to Twitter/X).  
It allows users to:

- 🧑‍💻 Create accounts and authenticate using JWT
- ✍️ Post short messages called "chirps"
- 📥 Fetch all chirps or filter by author
- 🔐 Update user credentials securely
- 💰 Upgrade users to "Chirpy Red" via webhook events
- 🗑️ Delete their own chirps (authorization enforced)

---

## 💡 Why Someone Should Care

This project demonstrates real-world backend engineering concepts:

- 🔐 Authentication & Authorization (JWT-based)
- 🧠 Secure password hashing
- ⚡ Efficient database queries (PostgreSQL + SQLC)
- 🔄 Database migrations using Goose
- 🌐 REST API design with proper HTTP status codes
- 🧩 Webhook handling (external event integration)
- 🏗️ Clean architecture & modular code structure

👉 This makes it a strong **resume-ready backend project**.

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **ORM/Query Tool:** SQLC
- **Migrations:** Goose
- **Auth:** JWT (JSON Web Tokens)

---

## ⚙️ How to Install and Run

### 1️⃣ Clone the repository

```bash
git clone https://github.com/your-username/chirpy.git
cd chirpy