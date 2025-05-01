# WorkoutPostingService

A RESTful API service for managing workout posts, built with Go and PostgreSQL.

## Overview

WorkoutPostingService is a backend service that allows users to create, read, update, and delete workout posts. It features user authentication, secure token management, and a robust API structure.

### Architecture

The service follows a clean architecture pattern with clear separation of concerns:

1. **API Layer** (`internal/api/`)
   - Handles HTTP request/response processing
   - Input validation and data transformation
   - Response formatting and error handling

2. **Application Layer** (`internal/app/`)
   - Contains core business logic
   - Manages application state and configuration
   - Coordinates between different layers
   - Implements dependency injection pattern

3. **Middleware Layer** (`internal/middleware/`)
   - Authentication and authorization
   - Request logging and tracing
   - Error handling middleware
   - Request context enrichment

4. **Data Layer** (`internal/store/`)
   - Database interactions and queries
   - Transaction management
   - Data model definitions
   - Repository pattern implementation

5. **Security Layer** (`internal/tokens/`)
   - JWT token generation and validation
   - User authentication logic
   - Security utilities

6. **Infrastructure**
   - PostgreSQL for persistent storage
   - Docker containers for isolation and deployment
   - Separate environments for testing and development
   - Database migration management with Goose

### Key Design Principles

- **Dependency Injection**: Components are loosely coupled through dependency injection
- **Interface-Driven Design**: Core business logic is defined through interfaces
- **Middleware Chain**: Request processing through composable middleware
- **Repository Pattern**: Database operations are abstracted through repositories
- **Clean Architecture**: Dependencies point inward, with domain at the center

## Features

- User registration and authentication
- JWT-based token authentication
- CRUD operations for workout posts
- PostgreSQL database integration
- Docker containerization
- Separate test database environment

## Tech Stack

- **Backend**: Go 1.23.3
- **Database**: PostgreSQL 12.4
- **Router**: Chi v5
- **Database Driver**: pgx v4
- **Migration Tool**: Goose v3
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.23.3 or higher
- Docker and Docker Compose
- Make (optional, for convenience)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/mjossany/workoutPostingService.git
cd workoutPostingService
```

### 2. Start the Database

```bash
docker-compose up -d
```

This will start two PostgreSQL instances:
- Main database on port 5432
- Test database on port 5433

### 3. Run the Application

```bash
go run main.go
```

By default, the server runs on port 8080. You can specify a different port using the `-port` flag:

```bash
go run main.go -port 3000
```

## API Endpoints

### Authentication

- `POST /users` - Register a new user
- `POST /tokens/authentication` - Create authentication token

### Workouts (Protected Routes)

- `GET /workouts/{id}` - Get a specific workout
- `POST /workouts` - Create a new workout
- `PUT /workouts/{id}` - Update an existing workout
- `DELETE /workouts/{id}` - Delete a workout

### Health Check

- `GET /health` - Check service health

## Project Structure

```
WorkoutPostingService/
├── database/           # Database data directories
├── internal/          
│   ├── api/           # API handlers and types
│   ├── app/           # Application setup and configuration
│   ├── middleware/    # HTTP middleware
│   ├── routes/        # Route definitions
│   ├── store/         # Database operations
│   ├── tokens/        # Token management
│   └── utils/         # Utility functions
├── migrations/        # Database migrations
└── docker-compose.yml # Docker services configuration
```

## Development

### Running Tests

```bash
# Make sure the test database is running
docker-compose up -d test-db

# Run tests
go test ./...
```

### Database Migrations

The project uses Goose for database migrations. Migration files are located in the `migrations/` directory.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
