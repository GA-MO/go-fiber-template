# Testing Standards for Golang Template Application

## Overview
This document establishes the **mandatory** testing standards for all current and future features in this application. Every feature must have comprehensive tests across all architectural layers.

## Testing Architecture

Our application follows Clean Architecture with three main layers that **MUST** be tested:

1. **Handler Layer** (`app/handlers/`) - HTTP request/response handling
2. **Service Layer** (`app/services/`) - Business logic  
3. **Repository Layer** (`app/repositories/`) - Data access

## Mandatory Testing Rule

**🚨 RULE: Every new feature MUST have tests for ALL THREE LAYERS before being considered complete.**

## Testing Patterns & Standards

### 1. Table-Driven Tests
All tests should use the table-driven pattern for comprehensive coverage:

```go
func TestFeatureHandler(t *testing.T) {
    testCaseList := []struct {
        name               string
        url                string
        method             string
        jsonBody           string
        expectedStatusCode int
        mockFunc           func(serviceMock *services.FeatureServiceMock)
    }{
        // Test cases here
    }
    
    for _, testCase := range testCaseList {
        t.Run(testCase.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 2. Mock Naming Convention
All mocks must follow this exact naming pattern:

- **Service Mocks**: `FeatureServiceMock` (not `MockFeatureService`)
- **Repository Mocks**: `FeatureRepositoryMock` (not `MockFeatureRepository`)
- **Constructor Function**: `NewFeatureServiceMock()` and `NewFeatureRepositoryMock()`

### 3. Required Libraries
- `github.com/stretchr/testify/assert` - for assertions
- `github.com/stretchr/testify/mock` - for mocking
- `github.com/DATA-DOG/go-sqlmock` - for database testing (repositories only)

## Layer-Specific Testing Requirements

### Handler Layer Tests (`app/handlers/*_test.go`)

**Required Test Categories:**
- ✅ **Success Cases** - Valid requests return expected responses
- ✅ **Validation Errors** - Invalid input returns 400 status
- ✅ **Service Errors** - Service failures return 500 status  
- ✅ **Parameter Validation** - Invalid URL parameters return 400 status
- ✅ **JSON Parsing** - Malformed JSON returns 400 status

**Naming Pattern:** `TestFeatureHandler`

**Example Test Coverage:**
- Create Feature Success
- Create Feature Validation Error
- Create Feature Service Error
- Update Feature Success
- Update Feature Not Found
- Delete Feature Success
- Get Feature Success
- List Features Success

### Service Layer Tests (`app/services/*_test.go`)

**Required Test Categories:**
- ✅ **Business Logic Validation** - All business rules enforced
- ✅ **Error Scenarios** - Repository errors handled properly
- ✅ **Edge Cases** - Boundary conditions tested
- ✅ **Dependency Validation** - Related entity existence checks

**Naming Pattern:** `TestFeatureService_MethodName`

**Example Test Coverage:**
- Successful creation/update/deletion
- Entity not found scenarios
- Duplicate name/identifier conflicts
- Repository error handling
- Validation logic testing

### Repository Layer Tests (`app/repositories/*_test.go`)

**Required Test Categories:**
- ✅ **CRUD Operations** - All database operations tested
- ✅ **Database Errors** - SQL errors handled properly
- ✅ **Edge Cases** - Empty results, no rows affected
- ✅ **SQL Query Validation** - Correct queries generated

**Naming Pattern:** `TestFeatureRepository_MethodName`

**Database Testing:**
- Use `sqlmock` for database mocking
- Test both success and failure scenarios
- Verify SQL queries match expectations
- Handle empty result sets properly (return `&[]Type{}` not `nil`)

## Test File Organization

### File Naming Convention
```
app/handlers/feature_handler_test.go
app/services/feature_service_test.go  
app/repositories/feature_repository_test.go
```

### Mock File Convention
```
app/services/feature_service_mock.go
app/repositories/feature_repository_mock.go
```

## Established Examples

The following features demonstrate complete test coverage and should be used as references:

### ✅ **Role Management** 
- ✅ `app/handlers/role_handler_test.go` - Complete handler tests
- ✅ `app/services/role_service_test.go` - Complete service tests  
- ✅ `app/repositories/role_repository_test.go` - Complete repository tests
- ✅ All required mocks implemented

### ✅ **Permission Management**
- ✅ `app/handlers/permission_handler_test.go` - Complete handler tests
- ✅ `app/services/permission_service_test.go` - Complete service tests
- ✅ `app/repositories/permission_repository_test.go` - Complete repository tests
- ✅ All required mocks implemented

### ✅ **User-Role Management**
- ✅ `app/handlers/user_role_handler_test.go` - Complete handler tests
- ✅ All user-role and role-permission operations covered
- ✅ Complex relationship testing implemented

### ✅ **Existing User Management**
- ✅ Repository and service tests already implemented
- ✅ Handler tests already implemented

## Testing Commands

### Run All Tests
```bash
go test ./app/... -v
```

### Run Layer-Specific Tests
```bash
go test ./app/handlers -v    # Handler tests
go test ./app/services -v    # Service tests  
go test ./app/repositories -v # Repository tests
```

### Run Feature-Specific Tests
```bash
go test ./app/handlers -v -run TestFeatureHandler
go test ./app/services -v -run TestFeatureService
go test ./app/repositories -v -run TestFeatureRepository
```

## Quality Gates

Before any feature is considered complete, it must pass these quality gates:

1. **✅ All tests passing** - `go test ./app/... -v` returns success
2. **✅ Complete coverage** - Handler, Service, and Repository tests exist
3. **✅ Proper naming** - Follows established naming conventions
4. **✅ Mock implementations** - Required mocks created and functional
5. **✅ Error scenarios** - Both success and failure cases covered

## Implementation Checklist

For every new feature, complete this checklist:

### Handler Layer
- [ ] Create `app/handlers/feature_handler_test.go`
- [ ] Test all HTTP endpoints (Create, Read, Update, Delete, List)
- [ ] Test validation errors (400 status codes)
- [ ] Test service errors (500 status codes)  
- [ ] Test invalid parameters (400 status codes)
- [ ] Test JSON parsing errors (400 status codes)

### Service Layer  
- [ ] Create `app/services/feature_service_test.go`
- [ ] Create `app/services/feature_service_mock.go`
- [ ] Test all business logic methods
- [ ] Test entity existence validation
- [ ] Test duplicate name/identifier scenarios
- [ ] Test repository error handling
- [ ] Implement proper error types (e.g., `FeatureError`)

### Repository Layer
- [ ] Create `app/repositories/feature_repository_test.go` 
- [ ] Create `app/repositories/feature_repository_mock.go`
- [ ] Test all CRUD operations with sqlmock
- [ ] Test database error scenarios
- [ ] Test empty result handling
- [ ] Verify SQL query correctness

### Final Verification
- [ ] Run `go test ./app/... -v` - all tests pass
- [ ] No linter errors
- [ ] Mock constructors follow naming convention
- [ ] Table-driven test structure used
- [ ] Comprehensive error scenario coverage

## Maintenance

This document should be updated whenever:
- New testing patterns are established
- Additional testing requirements are identified  
- New layers are added to the architecture
- Testing tools or libraries are changed

---

**Remember: Complete test coverage across all layers is not optional - it's a requirement for every feature in this application.** 