# SQLMock Testing Guide - Best Practices

## Overview
This guide covers the common patterns and fixes needed when writing repository tests with `sqlmock` in Go, based on real issues encountered in the golang-template project.

## Key Principles

### 1. Match Database Operations to Mock Expectations

**Rule**: Your mock expectations must match the actual database operations in your repository implementation.

- `db.Exec()` → Use `mock.ExpectExec()`
- `db.Query()` → Use `mock.ExpectQuery()`
- `db.QueryRow()` → Use `mock.ExpectQuery()`

### 2. Common Repository Patterns

#### Pattern A: Simple INSERT (returns LastInsertId)
```go
// Repository Code:
result, err := r.db.Exec(query, args...)
id, err := result.LastInsertId()

// Test Code:
mock.ExpectExec("INSERT INTO table").
    WithArgs(args...).
    WillReturnResult(sqlmock.NewResult(1, 1))
```

#### Pattern B: INSERT + SELECT (Create methods that return the created object)
```go
// Repository Code:
result, err := r.db.Exec(insertQuery, args...)
id, err := result.LastInsertId()
return r.GetByID(int(id)) // This calls QueryRow

// Test Code:
// First expect the INSERT
mock.ExpectExec("INSERT INTO table").
    WithArgs(args...).
    WillReturnResult(sqlmock.NewResult(1, 1))

// Then expect the SELECT
rows := sqlmock.NewRows([]string{"id", "name", "..."}).
    AddRow(1, "value", ...)
mock.ExpectQuery("SELECT (.+) FROM table WHERE id = (.+)").
    WithArgs(1).
    WillReturnRows(rows)
```

#### Pattern C: UPDATE + SELECT (Update methods that return the updated object)
```go
// Repository Code:
result, err := r.db.Exec(updateQuery, args...)
rowsAffected, err := result.RowsAffected()
if rowsAffected == 0 { return nil, sql.ErrNoRows }
return r.GetByID(id) // This calls QueryRow

// Test Code:
// First expect the UPDATE
mock.ExpectExec("UPDATE table").
    WithArgs(args...).
    WillReturnResult(sqlmock.NewResult(1, 1))

// Then expect the SELECT
rows := sqlmock.NewRows([]string{"id", "name", "..."}).
    AddRow(1, "updated_value", ...)
mock.ExpectQuery("SELECT (.+) FROM table WHERE id = (.+)").
    WithArgs(1).
    WillReturnRows(rows)
```

#### Pattern D: Simple SELECT
```go
// Repository Code:
err := r.db.QueryRow(query, args...).Scan(&fields...)

// Test Code:
rows := sqlmock.NewRows([]string{"id", "name", "..."}).
    AddRow(1, "value", ...)
mock.ExpectQuery("SELECT (.+) FROM table WHERE (.+)").
    WithArgs(args...).
    WillReturnRows(rows)
```

#### Pattern E: DELETE
```go
// Repository Code:
result, err := r.db.Exec(deleteQuery, args...)
rowsAffected, err := result.RowsAffected()

// Test Code:
mock.ExpectExec("DELETE FROM table (.+)").
    WithArgs(args...).
    WillReturnResult(sqlmock.NewResult(1, 1)) // or (1, 0) for not found
```

### 3. SQL Pattern Matching Best Practices

#### ✅ Good: Use Simple Regex Patterns
```go
// For simple queries
mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)")
mock.ExpectExec("INSERT INTO users")
mock.ExpectExec("UPDATE users")
mock.ExpectExec("DELETE FROM users (.+)")

// For JOIN queries
mock.ExpectQuery("SELECT (.+) FROM users u INNER JOIN roles r (.+)")
```

#### ❌ Avoid: Exact String Matching for Complex Queries
```go
// This is fragile and error-prone
mock.ExpectQuery("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ? AND status = ?")
```

### 4. Error Simulation Patterns

#### Simulating Database Errors
```go
// For Exec operations
mock.ExpectExec("INSERT INTO table").
    WithArgs(args...).
    WillReturnError(sql.ErrConnDone)

// For Query operations
mock.ExpectQuery("SELECT (.+) FROM table").
    WithArgs(args...).
    WillReturnError(sql.ErrNoRows)
```

#### Simulating No Rows Affected (for DELETE/UPDATE)
```go
mock.ExpectExec("DELETE FROM table (.+)").
    WithArgs(args...).
    WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
```

### 5. Common Mistakes and Fixes

#### Mistake 1: Using ExpectQuery for INSERT/UPDATE/DELETE
```go
// ❌ Wrong
mock.ExpectQuery("INSERT INTO users").WillReturnRows(rows)

// ✅ Correct
mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
```

#### Mistake 2: Forgetting the SELECT part of CREATE methods
```go
// ❌ Wrong - Only expecting INSERT
mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))

// ✅ Correct - Expecting both INSERT and SELECT
mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WillReturnRows(rows)
```

#### Mistake 3: Overly Specific SQL Patterns
```go
// ❌ Fragile - exact match
mock.ExpectQuery("SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at FROM permissions p INNER JOIN role_permissions rp ON p.id = rp.permission_id WHERE rp.role_id = ? ORDER BY p.resource, p.action")

// ✅ Robust - regex pattern
mock.ExpectQuery("SELECT (.+) FROM permissions p INNER JOIN role_permissions rp (.+)")
```

### 6. Debugging Tips

#### Always Check Your Repository Implementation First
Before writing tests, examine what database operations your repository actually performs:

```go
func (r *userRepository) Create(user *models.UserCreate) (*models.User, error) {
    // Step 1: This is an Exec operation
    result, err := r.db.Exec(query, user.Name, user.Email)
    if err != nil {
        return nil, err
    }

    // Step 2: This is a Query operation  
    id, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

    // Step 3: This calls GetByID which is another Query operation
    return r.GetByID(int(id))
}
```

Your test needs to expect: `ExpectExec` → `ExpectQuery`

#### Use Test Table Patterns
```go
testCases := []struct {
    name          string
    input         *models.UserCreate
    mockSetup     func(sqlmock.Sqlmock)
    expectedError error
}{
    {
        name: "successful creation",
        input: &models.UserCreate{Name: "test", Email: "test@example.com"},
        mockSetup: func(mock sqlmock.Sqlmock) {
            // First expect INSERT
            mock.ExpectExec("INSERT INTO users").
                WithArgs("test", "test@example.com").
                WillReturnResult(sqlmock.NewResult(1, 1))
            
            // Then expect SELECT
            rows := sqlmock.NewRows([]string{"id", "name", "email"}).
                AddRow(1, "test", "test@example.com")
            mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").
                WithArgs(1).
                WillReturnRows(rows)
        },
        expectedError: nil,
    },
}
```

### 7. Verification
Always end your tests with:
```go
assert.NoError(t, mock.ExpectationsWereMet())
```

This ensures all expected database operations were actually called.

## Quick Reference Checklist

Before writing repository tests:
- [ ] Check if the repository method uses `db.Exec()` or `db.Query()`
- [ ] For CREATE/UPDATE methods, check if they return the created/updated object
- [ ] Use `ExpectExec()` for INSERT/UPDATE/DELETE operations
- [ ] Use `ExpectQuery()` for SELECT operations
- [ ] Use simple regex patterns for SQL matching
- [ ] Test both success and error scenarios
- [ ] Verify all expectations were met

## Related Files
- Review existing test files: `*_repository_test.go`
- Check repository implementations: `*_repository.go`
- Testing standards: `Testing-standard.md` 