# 📚 LibraryApp - Full Stack Library Management System

A modern library management system with book CRUD operations and URL processing capabilities.

**Tech Stack:** Go + Gin + GORM + PostgreSQL + React + TypeScript + TailwindCSS + Docker

## 🚀 Quick Setup

### Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 13+
- Docker & Docker Compose

### Installation

```bash
# Clone repository
git clone https://github.com/tayfunyaltur/libraryapp.git
cd libraryapp

# Setup with Docker (Recommended)
docker-compose up --build

# Access the application
# Frontend: http://localhost:5173
# Backend API: http://localhost:8080
# API Documentation: http://localhost:8080/swagger/index.html
```

### Manual Setup

```bash
# 1. Database
createdb library

# 2. Backend
cd library-backend
cp .env.example .env  # Edit database credentials
go mod download
go run cmd/server/main.go

# 3. Frontend (new terminal)
cd library-frontend
cp .env.example .env  # Edit API URL if needed
npm install
npm run dev
```

## 📁 Project Structure

```
libraryapp/
├── library-backend/           # Go API Server
│   ├── cmd/server/main.go    # Entry point
│   ├── internal/
│   │   ├── api/handlers/     # HTTP handlers
│   │   ├── models/           # Data models
│   │   ├── service/          # Business logic
│   │   └── config/           # Configuration
│   ├── pkg/database/         # Database connection
│   └── docs/                 # Swagger documentation
├── library-frontend/         # React Application
│   ├── src/
│   │   ├── components/       # UI components
│   │   ├── pages/           # Route pages
│   │   ├── services/        # API integration
│   │   └── context/         # State management
│   └── package.json
├── docker-compose.yml        # Docker orchestration
└── README.md
```

## 🔌 API Endpoints

### Books API

| Method   | Endpoint                       | Description     |
| -------- | ------------------------------ | --------------- |
| `GET`    | `/api/v1/books`                | List all books  |
| `POST`   | `/api/v1/books`                | Create new book |
| `GET`    | `/api/v1/books/{id}`           | Get book by ID  |
| `PUT`    | `/api/v1/books/{id}`           | Update book     |
| `DELETE` | `/api/v1/books/{id}`           | Delete book     |
| `GET`    | `/api/v1/books/search?q=query` | Search books    |

### URL Processing API

| Method | Endpoint              | Description                |
| ------ | --------------------- | -------------------------- |
| `POST` | `/api/v1/process-url` | Process URL with operation |
| `GET`  | `/api/v1/url-stats`   | Get processing statistics  |

## 📋 API Usage Examples

### Create Book

```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "year": 2015,
    "isbn": "9780134190440",
    "description": "Comprehensive guide to Go"
  }'
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "year": 2015,
    "isbn": "9780134190440",
    "description": "Comprehensive guide to Go",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "message": "Book created successfully"
}
```

### Process URL

```bash
curl -X POST http://localhost:8080/api/v1/process-url \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
    "operation": "all"
  }'
```

**Response:**

```json
{
  "success": true,
  "processed_url": "https://www.byfood.com/food-experiences",
  "original_url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
  "operation": "all"
}
```

### Get Books with Filters

```bash
curl "http://localhost:8080/api/v1/books?limit=10&offset=0&author=Donovan"
```

## 🧪 Running Tests

### Backend Tests

```bash
cd library-backend

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/service/...

# Verbose output
go test -v ./...
```

### Frontend Tests

```bash
cd library-frontend

# Run unit tests
npm test

# Run tests with coverage
npm run test:coverage

# Run tests in watch mode
npm run test:watch
```

### API Test Examples

**Test Book Creation:**

```bash
# Valid request
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Book","author":"Test Author","year":2023}'

# Invalid request (missing required fields)
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Book"}'
# Expected: 400 Bad Request with validation errors
```

**Test URL Processing Edge Cases:**

```bash
# Test canonical operation
curl -X POST http://localhost:8080/api/v1/process-url \
  -d '{"url":"https://BYFOOD.com/food-EXPeriences?query=abc/","operation":"canonical"}'
# Expected: "https://BYFOOD.com/food-EXPeriences"

# Test redirection operation
curl -X POST http://localhost:8080/api/v1/process-url \
  -d '{"url":"https://BYFOOD.com/food-EXPeriences","operation":"redirection"}'
# Expected: "https://www.byfood.com/food-experiences"

# Test invalid operation
curl -X POST http://localhost:8080/api/v1/process-url \
  -d '{"url":"https://example.com","operation":"invalid"}'
# Expected: 400 Bad Request
```

## 📱 Application Screenshots

### 1. Dashboard - Book List

![Dashboard](screenshots/dashboard.png)
_Main dashboard showing book grid with search and filter options_

### 2. Book Creation Modal

![Create Book](screenshots/create-book.png)
_Modal form for creating new books with validation_

### 3. Book Detail View

![Book Detail](screenshots/book-detail.png)
_Detailed view of individual book with edit/delete options_

### 4. API Documentation

![Swagger API](screenshots/swagger-api.png)
_Interactive Swagger API documentation_

### 5. Mobile Responsive

![Mobile View](screenshots/mobile-view.png)
_Mobile-responsive design on various screen sizes_

## 🐳 Docker Deployment

### Build and Run

```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Services

- **PostgreSQL**: Port 5432 (database)
- **Backend**: Port 8080 (Go API)
- **Frontend**: Port 5173 (React app)

## 🔧 Development

### Backend Development

```bash
cd library-backend

# Install dependencies
go mod download

# Run with hot reload (install air first)
go install github.com/cosmtrek/air@latest
air

# Generate Swagger docs
swag init -g cmd/server/main.go -o docs

# Build binary
go build -o bin/server cmd/server/main.go
```

### Frontend Development

```bash
cd library-frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Type checking
npm run type-check
```

## 🌟 Features

### Backend Features

- ✅ RESTful API with Gin framework
- ✅ GORM auto-migration and models
- ✅ PostgreSQL database integration
- ✅ Input validation and error handling
- ✅ Swagger API documentation
- ✅ Structured logging
- ✅ URL processing service with 3 operations

### Frontend Features

- ✅ React 18 with TypeScript
- ✅ TailwindCSS for styling
- ✅ Context API for state management
- ✅ React Hook Form with validation
- ✅ Modal forms for book operations
- ✅ Responsive design
- ✅ Real-time search and filtering
- ✅ Pagination support

## 🐛 Troubleshooting

### Common Issues

**Database Connection Error:**

```bash
# Check if PostgreSQL is running
pg_isready -h localhost -p 5432

# Check database exists
psql -U postgres -l | grep library
```

**Port Already in Use:**

```bash
# Kill process using port 8080
lsof -ti:8080 | xargs kill -9

# Kill process using port 5173
lsof -ti:5173 | xargs kill -9
```

**Docker Issues:**

```bash
# Clean Docker system
docker system prune -f

# Rebuild without cache
docker-compose build --no-cache
```

## 📞 Contact

- **GitHub**: [@tayfunyaltur](https://github.com/tayfunyaltur)
- **Repository**: [https://github.com/tayfunyaltur/libraryapp](https://github.com/tayfunyaltur/libraryapp)

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.
