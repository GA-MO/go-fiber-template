# Golang Template Architecture Guidelines

This document outlines the architecture and folder structure guidelines for the Golang template project. The template follows **Clean Architecture** principles with clear separation of concerns and dependency injection patterns.

## 🏗️ Architecture Overview

The project follows a **layered architecture** pattern with the following principles:
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Single Responsibility**: Each component has one reason to change
- **Interface Segregation**: Small, focused interfaces
- **Dependency Injection**: Dependencies are injected from outside

## 📁 Project Structure

```
golang-template/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── go.sum                     # Go module checksums
├── Makefile                   # Build and development commands
├── .gitignore                 # Git ignore patterns
├── app.db                     # SQLite database file
├── Readme.md                  # Project documentation
├── Naming-conventions.md      # Code naming standards
├── Architecture-guidelines.md # This file
│
├── app/                       # Core application layer
│   ├── models/               # Domain models/entities
│   │   └── user.go
│   ├── repositories/         # Data access layer
│   │   ├── user_repository.go
│   │   ├── user_repository_test.go
│   │   └── user_repository_mock.go
│   ├── services/            # Business logic layer
│   │   ├── user_service.go
│   │   ├── user_service_test.go
│   │   └── user_service_mock.go
│   ├── handlers/            # HTTP presentation layer
│   │   ├── user_handler.go
│   │   └── user_handler_test.go
│   └── response.go          # Common response structures
│
├── database/                # Database connection and configuration
│   └── sqlite.go
│
├── middleware/              # HTTP middleware components
│   ├── logger.go
│   ├── recover.go
│   └── recover_test.go
│
├── logger/                  # Logging utilities
│   └── logger.go
│
├── httpclient/             # External HTTP client utilities
│   └── httpclient.go
│
└── validator/              # Input validation utilities
    ├── validator.go
    └── validator_test.go
```

---

## 📂 Folder Descriptions

### 🎯 `/` (Root Level)
**Purpose**: Application entry point and configuration files.

**Contains**:
- `main.go` - Application bootstrap, dependency injection, and server startup
- `go.mod/go.sum` - Go module dependencies
- `Makefile` - Build scripts and development commands
- Documentation files (`.md`)

**Guidelines**:
- Keep `main.go` focused on wiring dependencies and starting the server
- Avoid business logic in the main package
- Use dependency injection to connect components

### 🏢 `/app`
**Purpose**: Core application business layer containing domain logic.

This is the heart of your application following the **Repository-Service-Handler** pattern:

#### `/app/models`
**Purpose**: Domain entities and data structures.

**Guidelines**:
- Define core business entities (structs)
- Include JSON tags for API responses
- Keep models pure - no business logic here
- Use validation tags for input validation

**Example Structure**:
```go
type User struct {
    ID       int    `json:"id" db:"id"`
    Name     string `json:"name" db:"name" validate:"required"`
    Email    string `json:"email" db:"email" validate:"required,email"`
    IsActive bool   `json:"isActive" db:"is_active"`
}
```

#### `/app/repositories`
**Purpose**: Data access layer - abstracts database operations.

**Guidelines**:
- Define interfaces for each repository
- Implement database operations (CRUD)
- Handle database-specific errors
- Include both interface and implementation
- Provide mock implementations for testing

**Pattern**:
```go
// Interface definition
type UserRepository interface {
    GetByID(id int) (*models.User, error)
    Create(user *models.User) error
    // ... other methods
}

// Implementation
type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}
```

#### `/app/services`
**Purpose**: Business logic layer - contains application use cases.

**Guidelines**:
- Define interfaces for each service
- Implement business rules and workflows
- Orchestrate repository calls
- Handle business-specific errors
- Include validation logic
- Provide mock implementations for testing

**Pattern**:
```go
type UserService interface {
    GetUser(id int) (*models.User, error)
    CreateUser(user *models.User) error
    // ... other methods
}

type userService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}
```

#### `/app/handlers`
**Purpose**: HTTP presentation layer - handles web requests and responses.

**Guidelines**:
- Handle HTTP-specific concerns (parsing, responses)
- Call service layer for business logic
- Handle HTTP status codes and error responses
- Implement route registration functions
- Keep handlers thin - delegate to services

**Pattern**:
```go
type UserHandler struct {
    userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func RegisterUserRoutes(router fiber.Router, handler *UserHandler) {
    router.Get("/:id", handler.GetUser)
    router.Post("/", handler.CreateUser)
}
```

### 🗄️ `/database`
**Purpose**: Database connection management and configuration.

**Guidelines**:
- Handle database connections
- Provide database factory functions
- Include migration logic if needed
- Keep database-specific configurations here

### 🔧 `/middleware`
**Purpose**: HTTP middleware components for cross-cutting concerns.

**Guidelines**:
- Implement reusable middleware functions
- Handle authentication, logging, recovery, etc.
- Follow Fiber middleware patterns
- Include comprehensive tests

### 📝 `/logger`
**Purpose**: Centralized logging utilities.

**Guidelines**:
- Provide structured logging interface
- Support different log levels
- Include contextual logging capabilities
- Make it easy to swap logging implementations

### 🌐 `/httpclient`
**Purpose**: External HTTP client utilities for API integrations.

**Guidelines**:
- Provide reusable HTTP client configurations
- Handle retries, timeouts, and error handling
- Include common request/response patterns

### ✅ `/validator`
**Purpose**: Input validation utilities and custom validators.

**Guidelines**:
- Provide validation helpers
- Include custom validation rules
- Handle validation error formatting

---

## 🔄 Data Flow Architecture

```
HTTP Request
     ↓
[Middleware] → Authentication, Logging, Recovery
     ↓
[Handler] → Parse request, validate input
     ↓
[Service] → Business logic, validation, orchestration
     ↓
[Repository] → Data access, database operations
     ↓
[Database] → Persistence layer
```

**Dependency Direction**:
```
Handler → Service → Repository → Database
```

**Key Principles**:
- Each layer only depends on the layer directly below it
- Higher layers define interfaces that lower layers implement
- Dependencies flow inward (toward business logic)

---

## 🧪 Testing Strategy

### File Naming Conventions
- Test files: `*_test.go`
- Mock files: `*_mock.go`
- Place tests alongside the code they test

### Testing Layers
1. **Unit Tests**: Test individual functions and methods
2. **Integration Tests**: Test component interactions
3. **Mock Testing**: Use generated mocks for dependencies

### Mock Generation
Provide mock implementations for interfaces:
```go
// user_service_mock.go
type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUser(id int) (*models.User, error) {
    args := m.Called(id)
    return args.Get(0).(*models.User), args.Error(1)
}
```

---

## 🎯 Best Practices

### 1. **Dependency Injection**
```go
// Good: Inject dependencies through constructors
func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

// Bad: Create dependencies inside the service
func NewUserService() UserService {
    db := database.GetSQLite() // Don't do this
    userRepo := repositories.NewUserRepository(db)
    return &userService{userRepo: userRepo}
}
```

### 2. **Interface Segregation**
```go
// Good: Small, focused interfaces
type UserReader interface {
    GetByID(id int) (*models.User, error)
}

type UserWriter interface {
    Create(user *models.User) error
    Update(user *models.User) error
}

// Bad: Large, monolithic interfaces
type UserEverything interface {
    GetByID(id int) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    Delete(id int) error
    SendEmail(user *models.User) error  // Not related to user data
    ValidateUser(user *models.User) error // Should be in service layer
}
```

### 3. **Error Handling**
```go
// Good: Wrap errors with context
func (s *userService) GetUser(id int) (*models.User, error) {
    user, err := s.userRepo.GetByID(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user %d: %w", id, err)
    }
    return user, nil
}
```

### 4. **Separation of Concerns**
- **Handlers**: Only HTTP concerns (parsing, status codes)
- **Services**: Only business logic
- **Repositories**: Only data access
- **Models**: Only data structures

---

## 🚀 Getting Started with New Features

### Adding a New Entity (e.g., Product)

1. **Create Model**: `app/models/product.go`
2. **Create Repository**: 
   - `app/repositories/product_repository.go`
   - `app/repositories/product_repository_test.go`
   - `app/repositories/product_repository_mock.go`
3. **Create Service**:
   - `app/services/product_service.go`
   - `app/services/product_service_test.go`
   - `app/services/product_service_mock.go`
4. **Create Handler**:
   - `app/handlers/product_handler.go`
   - `app/handlers/product_handler_test.go`
5. **Wire in main.go**:
   ```go
   productRepo := repositories.NewProductRepository(db)
   productService := services.NewProductService(productRepo)
   productHandler := handlers.NewProductHandler(productService)
   handlers.RegisterProductRoutes(api.Group("/v1/product"), productHandler)
   ```

### Adding Middleware
1. Create file in `middleware/` directory
2. Follow Fiber middleware signature
3. Add to `main.go` middleware chain
4. Include comprehensive tests

### Adding External Dependencies
1. Add HTTP clients to `httpclient/`
2. Add validation rules to `validator/`
3. Add logging utilities to `logger/`

---

## 📋 Code Review Checklist

- [ ] Does the code follow the layered architecture?
- [ ] Are dependencies injected properly?
- [ ] Does each layer have a single responsibility?
- [ ] Are interfaces used instead of concrete types?
- [ ] Are tests provided for each component?
- [ ] Are mocks provided for external dependencies?
- [ ] Does error handling follow conventions?
- [ ] Are naming conventions followed?

---

This architecture provides a solid foundation for scalable, maintainable, and testable Go applications. Follow these guidelines to ensure consistency and quality across your codebase. 