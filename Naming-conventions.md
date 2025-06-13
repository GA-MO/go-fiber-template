# Golang Naming Conventions

This document outlines recommended naming conventions for Go (Golang) code, with good and bad examples for each case. These conventions are based on production-grade best practices and the Go community's standards.

## 📝 General Principles
- Use **CamelCase** for exported names (types, functions, variables).
- Use **camelCase** for unexported names.
- Use **ALL_CAPS** with underscores for acronyms only if they are at the start of a name (e.g., `ID`, `HTTP`).
- Keep names short, clear, and descriptive.
- Avoid stuttering (repeating the package name in the type or function name).
- Only export names that need to be accessed from other packages.
- Avoid abbreviations unless they are well-known (e.g., `URL`, `ID`).
- Be consistent with naming across the codebase.

---

## 📁 Folders and Filenames
- Use short, meaningful, all-lowercase names for folders and files.
- Avoid underscores, plurals, and uppercase letters.
- Filenames should match the main type or purpose (e.g., `user.go` for `User` type).
- Test files should end with `_test.go`.
- Avoid generic names like `utils.go` or `common.go`.

**Good:**
```
validator/
├── validator.go
├── validator_test.go
└── httpclient/
    ├── httpclient.go
    └── httpclient_test.go
└── user/
    ├── user.go
    └── service.go
└── database/
    ├── connection.go
    └── connection_test.go
└── logger/
    ├── logger.go
    └── logger_test.go
└── middleware/
    ├── auth.go
    └── auth_test.go
```
**Bad:**
```
/Validators/           // Don't use uppercase or plurals
/http_client/          // Don't use underscores
/User.go               // Don't use uppercase
/user/Utils.go         // Don't use uppercase or generic names
/user/common.go        // Avoid generic names
/user/user_test.go     // Use _test.go for test files, but keep lowercase
```
----

## 🧮 Variables
- Boolean variables should imply true/false (e.g., `isReady`, `hasPermission`, `shouldRetry`).
- Avoid single-letter names except for receivers and short-lived variables in small scopes.

**Good:**
```go
userName := "Alice"
count := 10
isReady := true
hasPermission := false
shouldRetry := true
```
**Bad:**
```go
User_name := "Alice"  // Don't use underscores
cnt := 10             // Too short, not descriptive
is_ready := true      // Don't use underscores
x := 5                // Avoid single-letter names for important variables
```

## 🏃‍♂️ Functions
- Function names should be verbs or verb phrases (e.g., `sendEmail`, `calculateTotal`).
- If a function returns an error, name it to indicate what can go wrong (e.g., `ParseConfig`).

**Good:**
```go
func getUserName() string {}
func calculateTotal() int {}
func parseConfig() error {}
```
**Bad:**
```go
func Get_user_name() string {} // Don't use underscores
func calc() int {}            // Too short, not descriptive
func do() error {}            // Not descriptive
```

## 🏷️ Constants
- Use CamelCase for exported constants.
- Use UPPER_SNAKE_CASE for environment variable names or when required by external systems.
- Group related constants using iota for enums.

**Good:**
```go
const MaxUsers = 100
const DefaultTimeout = 30

const (
    HTTPStatusOK       = 200
    HTTPStatusNotFound = 404
)

// For environment variables or external systems
const (
    ENV_DB_HOST = "DB_HOST"
    ENV_API_KEY = "API_KEY"
)

type Status int
const (
    StatusUnknown Status = iota
    StatusActive
    StatusInactive
)
```
**Bad:**
```go
const max_users = 100      // Don't use underscores for Go constants
const DEFAULT_TIMEOUT = 30 // Don't use all caps for Go constants
const http_status_ok = 200 // Don't use snake_case
```

## 🏗️ Structs
- Avoid embedding context.Context or sync.Mutex directly; use named fields (e.g., `ctx context.Context`, `mu sync.Mutex`).
- Group related fields and keep exported fields at the top.

**Good:**
```go
type User struct {
    Name string
    Email string
    IsActive bool
    mu sync.Mutex
}
```
**Bad:**
```go
type user_struct {
    name string
    age  int
}
type User struct {
    sync.Mutex // Don't embed directly
    Name string
}
```

## 🔌 Interfaces
- Prefer small, focused interfaces (often a single method).
- If an interface has one method, name it as the method plus -er (e.g., `Reader`, `Writer`).

**Good:**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Stringer interface {
    String() string
}
```
**Bad:**
```go
type IReader interface { // Don't prefix with 'I'
    Read(p []byte) (n int, err error)
}
type ReadAndWrite interface { // Too broad
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```

## 📦 Packages
- Use short, meaningful, lower-case names with no underscores or plurals.
- Avoid stuttering (e.g., `user.User`).
- Avoid names like `util`, `common`, or `helpers`.

**Good:**
```
package validator
package httpclient
```
**Bad:**
```
package Validator // Don't use uppercase
package my_http_client // Don't use underscores
package users // Don't use plurals
package util // Avoid generic names
```

## 📤 API JSON Responses
- Always use camelCase for JSON tags.
- Use `omitempty` for optional fields.
- Keep JSON field names consistent and descriptive.

**Good:**
```go
type User struct {
    UserName string `json:"userName"`
    Email    string `json:"email"`
    IsActive bool   `json:"isActive,omitempty"`
}
```
**Bad:**
```go
type User struct {
    UserName string `json:"UserName"` // Don't use PascalCase in JSON tags
    Email    string `json:"Email"`
    IsActive bool   `json:"is_active"` // Don't use snake_case in JSON tags
}
```

## 👤 Receivers
- Use short, meaningful receiver names (e.g., `u` for `User`, `srv` for `Server`).
- Avoid generic names like `this` or `self`.

**Good:**
```go
func (u *User) IsActive() bool {
    return u.IsActive
}
func (srv *Server) Start() {}
```
**Bad:**
```go
func (user *User) IsActive() bool { // Too long
    return user.IsActive
}
func (this *Server) Start() {} // Avoid 'this'
```

## 📚 Additional Must-Follow Conventions
- **Acronyms:** Capitalize acronyms (e.g., `HTTPServer`, `APIClient`, not `HttpServer` or `ApiClient`).
- **Avoid Hungarian Notation:** Don't prefix names with type information (e.g., `strName`, `iCount`).
- **No Get Prefix:** Omit `Get` from getter method names (e.g., `Name()` not `GetName()`).
- **No underscores in Go names (except for test functions like `TestX_Y`).**
- **No package import aliases unless necessary.**
- **Avoid stuttering in function/type names.**

---

For more details, see the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#package-names) and [Effective Go](https://golang.org/doc/effective_go.html#names). 