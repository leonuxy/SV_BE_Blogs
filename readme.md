Postman Link : https://gm8888-4541.postman.co/workspace/GM-Workspace~06ef7a90-b7c2-41d3-8a6d-55b1a295a3d0/collection/23706900-0dd114db-9547-4f47-b9d2-f4c997c6fda3?action=share&creator=23706900


# Backend Blogs

A RESTful API backend for a blog application built with Go, Gin, and GORM.

---

## Features

- RESTful API structure using Gin
- ORM with GORM
- Environment configuration via `.env`
- MySQL database

---

## Project Structure
. ├── config/ # Database and config setup ├── controllers/ # Handle requests from user ├── models/ # GORM models ├── routes/ # API route grouping ├── migrations/ # Migration files ├── .env # Environment variables ├── go.mod # Go module file ├── go.sum # Go module checksums ├── main.go # Entry point └── README.md

---

---

##  Starting Up

### 1. Clone the repo

```bash
git clone https://github.com/leonuxy/SV_BE_Blogs.git
cd SV_BE_Blogs


setup env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=dbname

run migration
go run migrations/20250411-Post.go

start server
go run main.go

The server runs on: http://localhost:8080
