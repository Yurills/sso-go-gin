# SSO Go Gin

![Go Version](https://img.shields.io/github/go-mod/go-version/Yurills/sso-go-gin)
![License](https://img.shields.io/github/license/Yurills/sso-go-gin)
![Build Status](https://img.shields.io/github/actions/workflow/status/Yurills/sso-go-gin/go.yml?branch=main)

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
````

## ⚡ Getting Started

### Prerequisites

Ensure you have the following installed:

  - [Go](https://go.dev/dl/) (1.20+)
  - [Docker](https://www.docker.com/) & Docker Compose
  - [Node.js](https://nodejs.org/) & npm (for frontend development)

### 🔧 Installation

1.  **Clone the repository**

    ```bash
    git clone [https://github.com/Yurills/sso-go-gin.git](https://github.com/Yurills/sso-go-gin.git)
    cd sso-go-gin
    ```

2.  **Setup Environment Variables**
    Create a `.env` file in the root directory (or inside `config/` if required) based on your configuration needs.

    ```env
    PORT=8080
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=sso_db
    JWT_SECRET=your_super_secret_key
    ```

### 🐳 Running with Docker (Recommended)

The easiest way to run the application is using Docker Compose.

```bash
docker-compose up --build
```

This will start the backend service, the database, and the frontend server.

### 🏃 Running Manually

**Backend:**

1.  Install Go dependencies:
    ```bash
    go mod download
    ```
2.  Run the server:
    ```bash
    go run cmd/api/main.go
    ```

**Frontend:**

1.  Navigate to the frontend directory:
    ```bash
    cd frontend
    ```
2.  Install dependencies:
    ```bash
    npm install
    ```
3.  Start the development server:
    ```bash
    npm start
    ```

## 📡 API Endpoints

Below are the core endpoints provided by the service (example):

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/auth/register` | Register a new user |
| `POST` | `/auth/login` | Login and receive JWT |
| `POST` | `/auth/logout` | Invalidate session (if applicable) |
| `GET` | `/auth/me` | Get current user profile (Protected) |
| `POST` | `/auth/refresh` | Refresh access token |

*Note: Check `cmd/api/routes.go` or the swagger documentation (if enabled) for the full list of routes.*
## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.
