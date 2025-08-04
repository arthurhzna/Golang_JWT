# Golang JWT Authentication API

A robust REST API built with Go that implements JWT-based authentication with refresh tokens and session management using PostgreSQL.

## 📋 Features

- ✅ User Registration & Login
- ✅ JWT Access & Refresh Token Authentication
- ✅ Session Management with Database Storage
- ✅ Token Renewal Mechanism
- ✅ Session Revocation & Logout
- ✅ Password Hashing with Bcrypt
- ✅ Input Validation
- ✅ Clean Architecture Pattern
- ✅ PostgreSQL Database Integration
- ✅ Environment Configuration

## 🏗️ Project Structure

```
Golang_JWT/
├── app/                    # Application configuration
│   └── database.go         # Database connection setup
├── controller/             # HTTP handlers
│   ├── user_controller.go
│   └── user_controller_imp.go
├── exception/              # Custom error handling
│   ├── error_handler.go
│   └── not_found_error.go
├── helper/                 # Utility functions
│   ├── error.go
│   ├── json.go
│   ├── model.go           # Response mappers
│   ├── password.go        # Password hashing
│   └── tx.go             # Transaction helpers
├── model/                 # Data models
│   ├── domain/           # Domain entities
│   │   ├── user.go
│   │   └── sessions.go
│   └── web/              # API request/response models
│       ├── user_*.go
│       └── renew_access_token_*.go
│       └── user_claims*.go
│       └── web_response*.go
├── repository/            # Data access layer
│   ├── user_repository.go
│   └── user_repository_imp.go
├── service/              # Business logic layer
│   ├── user_service.go
│   └── user_service_impl.go
├── token/                # JWT token management
│   ├── user_token.go
│   └── user_token_imp.go
└── main.go              # Application entry point
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23.4 or higher
- PostgreSQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Golang_JWT
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgres://username:password@localhost/database_name?sslmode=disable
   SECRET_KEY=your-secret-key-for-jwt-signing
   ```

4. **Set up PostgreSQL database**
   ```sql
   -- Create users table
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(100) NOT NULL,
       email VARCHAR(100) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   -- Create sessions table
   CREATE TABLE sessions (
       id VARCHAR(255) PRIMARY KEY,
       user_email VARCHAR(100) NOT NULL,
       refresh_token TEXT NOT NULL,
       is_revoked BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       expires_at TIMESTAMP NOT NULL,
       FOREIGN KEY (user_email) REFERENCES users(email)
   );
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:3000`

## 📡 API Endpoints

### Authentication

#### Register User
```http
POST /api/users/register
Content-Type: application/json

{
    "username": "arthurhzna",
    "email": "arthur@example.com",
    "password": "securepassword123"
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "username": "arthurhzna",
        "email": "arthur@example.com"
    }
}
```

#### Login
```http
POST /api/users/login
Content-Type: application/json

{
    "email": "john@example.com",
    "password": "securepassword123"
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "session_id": "uuid-session-id",
        "access_token": "eyJhbGciOiJIUzI1NiIs...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
        "access_token_expires_at": "2024-01-01T12:15:00Z",
        "refresh_token_expires_at": "2024-01-02T12:00:00Z",
        "user": {
            "id": 1,
            "username": "arthurhzna",
            "email": "arthur@example.com"
        }
    }
}
```

#### Renew Access Token
```http
POST /api/users/renew-token
Content-Type: application/json

{
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIs...",
        "access_token_expires_at": "2024-01-01T12:30:00Z"
    }
}
```

#### Logout
```http
DELETE /api/users/logout/{session_id}
Authorization: Bearer <access_token>
```

#### Revoke Session
```http
PUT /api/users/revoke/{session_id}
Authorization: Bearer <access_token>
```

## 🔧 Configuration

### Token Settings

- **Access Token Expiry:** 15 minutes
- **Refresh Token Expiry:** 24 hours
- **JWT Algorithm:** HMAC-SHA256
- **Token Claims:** User ID, Username, Email, JWT Standard Claims

### Database Configuration

- **Database:** PostgreSQL
- **Connection Pool:** 
  - Max Idle Connections: 5
  - Max Open Connections: 20
  - Connection Max Lifetime: 60 minutes
  - Connection Max Idle Time: 10 minutes

## 🔐 Security Features

- **Password Hashing:** Bcrypt with salt
- **JWT Security:** HMAC-SHA256 signing
- **Session Management:** Database-stored sessions with revocation
- **Token Validation:** Comprehensive token verification
- **Input Validation:** Request payload validation
- **SQL Injection Protection:** Parameterized queries

## 🧪 Testing

Run tests with:
```bash
go test ./...
```

## 🔧 Dependencies

- **Web Framework:** Native Go HTTP server
- **Database Driver:** `github.com/jackc/pgx/v5`
- **JWT Library:** `github.com/golang-jwt/jwt/v5`
- **Validation:** `github.com/go-playground/validator/v10`
- **Environment:** `github.com/joho/godotenv`
- **UUID Generation:** `github.com/google/uuid`
- **Password Hashing:** `golang.org/x/crypto/bcrypt`

## 📝 Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | Yes |
| `SECRET_KEY` | JWT signing secret key | Yes |

## 🏛️ Architecture

This project follows **Clean Architecture** principles:

- **Controller Layer:** HTTP request handling
- **Service Layer:** Business logic implementation
- **Repository Layer:** Data access abstraction
- **Domain Layer:** Core business entities
- **Helper Layer:** Utility functions and cross-cutting concerns

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🐛 Known Issues

Please check the issues section for known bugs and planned features.

## 📞 Support

For support, please open an issue in the GitHub repository.
