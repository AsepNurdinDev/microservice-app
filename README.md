# Microservices APP 🌿

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-14-black?style=flat&logo=next.js)](https://nextjs.org/)
[![Docker](https://img.shields.io/badge/Docker-Container-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Managed-4169E1?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-NoSQL-47A248?style=flat&logo=mongodb)](https://www.mongodb.com/)

Platform Wellness (Tenangin) berbasis arsitektur **Microservices**. Proyek ini mendemonstrasikan implementasi sistem terdistribusi dengan pemisahan tanggung jawab (*Separation of Concerns*) yang ketat, keamanan berbasis JWT, dan manajemen infrastruktur menggunakan Docker.

## 🏗️ Arsitektur Sistem

Proyek ini menggunakan beberapa layanan independen yang berkomunikasi dalam jaringan internal Docker:

1.  **API Gateway**: Dibangun dengan Golang, berfungsi sebagai entry point tunggal yang meneruskan request ke layanan terkait.
2.  **Auth Service**: Menangani manajemen pengguna, pendaftaran, dan autentikasi. Menggunakan **PostgreSQL** untuk integritas data transaksional.
3.  **Article Service**: Mengelola konten edukasi wellness. Menggunakan **MongoDB** untuk fleksibilitas skema dokumen artikel.


## 🛠️ Tech Stack

- **Backend**: Golang (Gin Framework)
- **Frontend**: Next.js & Tailwind CSS
- **Database**: PostgreSQL (Relational) & MongoDB (NoSQL)
- **Security**: JWT (JSON Web Token) dengan Middleware kustom
- **Infrastruktur**: Docker & Docker Compose
- **Reverse Proxy**: Nginx

## 🚀 Panduan Instalasi

### 1. Persiapan
Pastikan Anda sudah menginstal **Docker** dan **Docker Compose** di mesin Anda (Ubuntu/Windows/Mac).

### 2. Konfigurasi Environment
Buat file `.env` di direktori root proyek (`microservice-app/.env`) dan isi dengan variabel berikut:

```env
# Database Auth (PostgreSQL)
AUTH_DB_USER=user
AUTH_DB_PASSWORD=281205
AUTH_DB_NAME=auth_db
AUTH_DB_HOST=postgres
AUTH_DB_PORT=5432

# Database Article (MongoDB)
MONGO_URI=mongodb://mongodb:27017
ARTICLE_DB_NAME=article_db

# Security
JWT_SECRET=GunakanStringRahasiaYangSangatPanjangDisini123!


# Sturucture Folder:

microservices-app/
│
├── docker-compose.yml
├── .env
│
├── services/
│   ├── auth-service/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   │
│   │   ├── internal/
│   │   │   ├── domain/
│   │   │   │   └── user.go
│   │   │   │
│   │   │   ├── repository/
│   │   │   │   └── user_repository.go
│   │   │   │
│   │   │   ├── usecase/
│   │   │   │   └── auth_usecase.go
│   │   │   │
│   │   │   ├── delivery/
│   │   │   │   └── http/
│   │   │   │       └── handler.go
│   │   │   │
│   │   │   └── infrastructure/
│   │   │       ├── database/
│   │   │       │   └── postgres.go
│   │   │       └── jwt/
│   │   │           └── jwt.go
│   │   │
│   │   ├── pkg/
│   │   │   └── utils/
│   │   │       └── response.go
│   │   │
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── article-service/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   │
│   │   ├── internal/
│   │   │   ├── domain/
│   │   │   │   └── article.go
│   │   │   │
│   │   │   ├── repository/
│   │   │   │   └── article_repository.go
│   │   │   │
│   │   │   ├── usecase/
│   │   │   │   └── article_usecase.go
│   │   │   │
│   │   │   ├── delivery/
│   │   │   │   └── http/
│   │   │   │       └── handler.go
│   │   │   │
│   │   │   └── infrastructure/
│   │   │       ├── database/
│   │   │       │   └── mongodb.go
│   │   │       └── middleware/
│   │   │           └── auth_middleware.go
│   │   │
│   │   ├── pkg/
│   │   │   └── utils/
│   │   │       └── response.go
│   │   │
│   │   ├── Dockerfile
│   │   └── go.mod
│
└── gateway/
    ├── nginx.conf
    └── Dockerfile