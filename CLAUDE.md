# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web server built with the Beego framework, implementing user authentication, RBAC (Role-Based Access Control), and API endpoints. The server provides user management, permission control, and authentication features.

## Commands

### Running the Server
```bash
go run main.go
```
The server will start on port 8080 (configured in conf/app.conf).

### Testing
```bash
go test ./tests
```
Uses GoConvey for testing framework.

### Building
```bash
go build -o cmc-server
```

### Dependencies
```bash
go mod tidy
go mod download
```

## Architecture

### Framework & Core Components
- **Beego v2**: Web framework handling routing, middleware, and HTTP server
- **XORM**: ORM for database operations with MySQL
- **Redis**: Caching and session management
- **Custom JWT**: Token-based authentication with encryption

### Project Structure
```
cmc-server/
├── main.go                 # Application entry point
├── conf/app.conf           # Beego configuration
├── routers/                # Route definitions
│   ├── router.go          # Main router setup
│   └── v1.go              # API v1 routes (/api/v1)
├── controllers/           # HTTP request handlers
├── service/               # Business logic layer
├── models/                # Database models (User, Role, Promission)
├── dto/                   # Data transfer objects
├── components/            # Reusable components
│   ├── jwt/              # JWT authentication
│   ├── rbac/             # Role-based access control
│   ├── orm/              # Database connection
│   ├── redis/            # Redis client
│   └── logger/           # Logging
├── filter/               # Middleware and filters
├── static/               # Static data files (role.json, promission.json)
└── tests/                # Test files
```

### Key Architecture Patterns

#### Routing Pattern
- Routes are defined in `routers/v1.go` using Beego's namespace system
- All API endpoints are prefixed with `/api/v1`
- Authentication bypass paths use `/noAuth` prefix (defined in `jwt.NoAuthPathPrefix`)

#### Authentication Flow
- JWT tokens are encrypted using custom encryption (not standard JWT)
- Tokens have 30-day expiration
- `JwtFilter` middleware intercepts all requests except `/noAuth` paths
- User ID is extracted from token and available in controllers via `JwtDataPayload`

#### RBAC System
- Permissions are stored in database and initialized from `static/promission.json`
- Roles are defined in `static/role.json` with associated permissions
- Admin user is auto-created with ID "admin" and password "1.123."

#### Database Layer
- XORM handles database operations with auto-sync for model changes
- Models extend `common.BaseEntry` for consistent fields
- Output transformation via `Output()` method for DTOs

### Database Configuration
Database connection is hardcoded in `components/orm/mod.go`:
- Host: mysql2.sqlpub.com:3307
- Database: cursor_pool
- User: admin_dys

### Important Notes
- The server automatically initializes RBAC data on startup
- Admin user credentials are hardcoded (ID: "admin", password: "1.123.")
- JWT implementation uses custom encryption, not standard JWT libraries
- Static data files (role.json, promission.json) are used to seed the database