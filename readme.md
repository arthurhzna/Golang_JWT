# Golang JWT Authentication API

A robust REST API built with Go that implements JWT-based authentication with refresh tokens, session management, middleware protection, and automated cleanup using PostgreSQL.

## 📋 Features

- ✅ User Registration & Login
- ✅ JWT Access & Refresh Token Authentication
- ✅ Authentication Middleware Protection
- ✅ Session Management with Database Storage
- ✅ Automated Session Cleanup Scheduler
- ✅ Token Renewal Mechanism
- ✅ Session Revocation & Logout
- ✅ Password Hashing with Bcrypt
- ✅ Input Validation
- ✅ Clean Architecture Pattern
- ✅ PostgreSQL Database Integration
- ✅ Environment Configuration
- ✅ Background Task Processing

## 🏗️ Project Structure

```
Golang_JWT/
├── app/                    # Application configuration
│   ├── database.go         # Database connection setup
│   └── route.go           # HTTP routing configuration
├── middleware/             # Middleware components
│   └── auth_middleware.go # JWT authentication middleware
├── scheduler/             # Background task schedulers
│   └── cleanup_scheduler.go # Session cleanup scheduler
├── controller/            # HTTP handlers
│   ├── user_controller.go
│   └── user_controller_imp.go
├── exception/             # Custom error handling
│   ├── error_handler.go
│   └── not_found_error.go
├── helper/                # Utility functions
│   ├── error.go
│   ├── json.go
│   ├── model.go          # Response mappers
│   ├── password.go       # Password hashing
│   └── tx.go            # Transaction helpers
├── model/               # Data models
│   ├── domain/         # Domain entities
│   │   ├── user.go
│   │   └── sessions.go
│   └── web/           # API request/response models
│       ├── user_*.go
│       ├── renew_access_token_*.go
│       ├── user_claims.go
│       └── web_response.go
├── repository/         # Data access layer
│   ├── user_repository.go
│   └── user_repository_imp.go
├── service/           # Business logic layer
│   ├── user_service.go
│   └── user_service_impl.go
├── token/            # JWT token management
│   ├── user_token.go
│   └── user_token_imp.go
└── main.go          # Application entry point
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

### Public Endpoints (No Authentication Required)

#### Register User
```http
POST /api/register
Content-Type: application/json

{
    "username": "arthur",
    "email": "arthur@example.com",
    "password": "mypassword123"
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "username": "arthur",
        "email": "arthur@example.com"
    }
}
```

#### Login
```http
POST /api/users/login
Content-Type: application/json

{
    "email": "arthur@example.com",
    "password": "mypassword123"
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "session_id": "892d07bf-f4d0-4c79-8a21-306a8976201a",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "access_token_expires_at": "2025-08-09T17:07:25+07:00",
        "refresh_token_expires_at": "2025-08-10T16:52:25+07:00",
        "user": {
            "id": 2,
            "username": "arthur",
            "email": "arthur@example.com"
        }
    }
}
```

#### Renew Access Token
```http
POST /api/users/refresh-token
Content-Type: application/json

{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "access_token_expires_at": "2025-08-09T17:30:25+07:00"
    }
}
```

### Protected Endpoints (Authentication Required)

All protected endpoints require `Authorization: Bearer <access_token>` header.

#### Get All Users
```http
GET /api/users
Authorization: Bearer <access_token>
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "id": 1,
            "username": "user1",
            "email": "user1@example.com"
        },
        {
            "id": 2,
            "username": "arthur",
            "email": "arthur@example.com"
        }
    ]
}
```

#### Get User By ID
```http
GET /api/users/:userId
Authorization: Bearer <access_token>
```

**Response:**
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "username": "arthur",
        "email": "arthur@example.com"
    }
}
```

#### Logout
```http
POST /api/users/logout
Authorization: Bearer <access_token>
```

#### Revoke Session
```http
POST /api/users/revoke-session
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

## 🛡️ Authentication Middleware

### Features
- **JWT Token Validation:** Validates Bearer tokens in Authorization header
- **Context Injection:** Adds user claims to request context for controllers
- **Error Handling:** Returns standardized 401 responses for invalid tokens
- **Panic Recovery:** Safely handles token validation panics
- **Clean Architecture:** Separates authentication logic from business logic

### Usage Example
```go
// Protected routes automatically get user context
func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // User claims available via context if needed
    claims := r.Context().Value(middleware.UserClaimsKey).(*web.UserClaims)
    // ... controller logic
}
```

## 🧹 Automated Session Cleanup

### Background Scheduler
- **Automatic Cleanup:** Runs every 24 hours in background
- **Database Maintenance:** Removes expired refresh tokens and sessions
- **Non-blocking:** Runs as separate goroutine without affecting API performance
- **Error Handling:** Proper transaction management with rollback on errors
- **Startup Cleanup:** Immediate cleanup on application start

### Configuration
```go
// Default: 24 hours interval
cleanupScheduler := scheduler.NewCleanupScheduler(userRepository, db)

// Custom interval (for testing)
cleanupScheduler.SetInterval(1 * time.Hour)
```

## 🔐 Security Features

- **Password Hashing:** Bcrypt with salt
- **JWT Security:** HMAC-SHA256 signing
- **Authentication Middleware:** Route-level protection
- **Session Management:** Database-stored sessions with revocation
- **Token Validation:** Comprehensive token verification with panic recovery
- **Input Validation:** Request payload validation
- **SQL Injection Protection:** Parameterized queries
- **Context Security:** Secure user context injection
- **Automatic Cleanup:** Expired session removal for security hygiene

## 🧪 Testing

Run tests with:
```bash
go test ./...
```

## 🔧 Dependencies

- **HTTP Router:** `github.com/julienschmidt/httprouter`
- **Database Driver:** `github.com/jackc/pgx/v5`
- **JWT Library:** `github.com/golang-jwt/jwt/v5`
- **Validation:** `github.com/go-playground/validator/v10`
- **Environment:** `github.com/joho/godotenv`
- **UUID Generation:** `github.com/google/uuid`
- **Password Hashing:** `golang.org/x/crypto/bcrypt`

## 🆕 New Features

### Authentication Middleware
- Clean, reusable middleware architecture
- Automatic token validation for protected routes
- Context-based user information passing
- Standardized error responses

### Background Session Cleanup
- Automated expired session removal
- Configurable cleanup intervals
- Background processing without API blocking
- Proper database transaction handling

## 📝 Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | Yes |
| `SECRET_KEY` | JWT signing secret key | Yes |

## 🧪 Testing

### API Testing with cURL

**Login and get token:**
```bash
curl -X POST http://localhost:3000/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"arthur@example.com","password":"mypassword123"}'
```

**Access protected endpoint:**
```bash
curl -X GET http://localhost:3000/api/users \
  -H "Authorization: Bearer <your_access_token>"
```

**Run unit tests:**
```bash
go test ./...
```

## 🏛️ Architecture

This project follows **Clean Architecture** principles with additional middleware and scheduler layers:

- **Middleware Layer:** Authentication and request preprocessing
- **Controller Layer:** HTTP request handling
- **Service Layer:** Business logic implementation  
- **Repository Layer:** Data access abstraction
- **Domain Layer:** Core business entities
- **Helper Layer:** Utility functions and cross-cutting concerns
- **Scheduler Layer:** Background task processing

### Request Flow
```
HTTP Request → Middleware (Auth) → Controller → Service → Repository → Database
                     ↓
              Context Injection
```

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
