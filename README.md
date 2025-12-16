# SSO Go Gin

![Go Version](https://img.shields.io/github/go-mod/go-version/Yurills/sso-go-gin)

A robust Single Sign-On (SSO) authentication service built with **Go** and the **Gin Web Framework**. This project includes a backend API for managing user identities and sessions, coupled with a **TypeScript** frontend for user interaction.

## 🚀 Features

- **User Authentication**: Secure Login and Registration endpoints.
- **SSO Capabilities**: JSON Web Token (JWT) based authentication for stateless session management.
- **RESTful API**: Clean and structured API built with Gin.
- **Frontend Client**: TypeScript-based UI for interacting with the auth system.
- **Dockerized**: Includes `docker-compose` for easy deployment of the application and database.
- **Clean Architecture**: organized into `internal`, `pkg`, and `cmd` for maintainability.

## 🛠️ Tech Stack

**Backend**
- **Language**: [Go (Golang)](https://go.dev/)
- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Configuration**: Viper (implied for config management)
- **Database**: PostgreSQL (Default recommended via Docker)

**Frontend**
- **Language**: TypeScript
- **Runtime**: Node.js

**DevOps**
- Docker & Docker Compose

## 📂 Project Structure

```bash
sso-go-gin/
├── cmd/
│   └── api/            # Entry point for the application
├── config/             # Configuration files (env, yaml)
├── frontend/           # TypeScript frontend source code
├── internal/           # Private application and business logic
├── pkg/                # Library code that can be used by external applications
├── docker-compose.yml  # Docker services definition
├── go.mod              # Go dependencies
└── README.md
