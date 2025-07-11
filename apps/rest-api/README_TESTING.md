# Testing Guide

## Overview
This project is covered by various types of tests to ensure code quality and reliability.

## Test Structure

### Unit Tests (22 test files)
**Handlers (6 files)**
- `internal/handlers/handlers_test.go` - HTTP handler base functionality
- `internal/handlers/health_test.go` - health check endpoint
- `internal/handlers/simple_test.go` - JSON responses, error handling
- `internal/handlers/chat_test.go` - chat endpoints, invalid JSON
- `internal/handlers/workout_test.go` - workout plan endpoints
- `internal/handlers/profile_test.go` - profile CRUD endpoints

**Middleware (3 files)**
- `internal/middleware/auth_test.go` - JWT authentication, token validation
- `internal/middleware/logging_test.go` - request/response logging
- `internal/middleware/validation_test.go` - request validation

**Services (6 files)**
- `internal/services/ai_test.go` - AI service rating functionality
- `internal/services/auth_test.go` - user registration/login
- `internal/services/health_test.go` - database connectivity
- `internal/services/profile_test.go` - profile management
- `internal/services/openrouter_test.go` - OpenRouter client
- `internal/services/service_test.go` - base service functionality

**Repository (2 files)**
- `internal/repository/mongodb_test.go` - MongoDB operations
- `internal/repository/postgres_test.go` - PostgreSQL operations

**Core (5 files)**
- `cmd/server/main_test.go` - main application and migrations
- `internal/models/models_test.go` - data models validation
- `internal/config/config_test.go` - configuration loading
- `pkg/utils/crypto_test.go` - password hashing
- `pkg/utils/jwt_test.go` - JWT token operations
- `pkg/utils/validation_test.go` - data validation helpers

### Integration Tests
- `test_integration.go` - integration tests for API endpoints

### Benchmark Tests
- Performance tests for critical algorithms

## Running Tests

### All tests
```bash
make test
```

### Unit tests only
```bash
make test-unit
```

### Integration tests
```bash
make test-integration
```

### Tests with coverage
```bash
make test-coverage
```

### Benchmarks
```bash
make benchmark
```

## Test Environment Setup

### Test databases
```bash
make test-setup    # Create test databases
make test-teardown # Remove test databases
```

## Code Coverage

Main components covered by tests:

### Handlers (HTTP handlers)
- ✅ Register/Login endpoints
- ✅ Health check endpoint
- ✅ Rating endpoint
- ✅ Error handling

### Middleware (Request processing)
- ✅ Authentication middleware
- ✅ JWT token validation
- ✅ Logging middleware
- ✅ Request validation middleware
- ✅ Request/response tracking

### Services (Business logic)
- ✅ AI Service rating functionality
- ✅ Authentication service
- ✅ Health service
- ✅ Profile service
- ✅ OpenRouter client
- ✅ Model switching logic
- ✅ Error handling
- ✅ Mock dependencies

### Repository (Data layer)
- ✅ MongoDB rating calculation
- ✅ PostgreSQL user management
- ✅ Consecutive days algorithm
- ✅ Level calculation logic
- ✅ Mock repositories

### Configuration
- ✅ Environment variable loading
- ✅ Default value handling
- ✅ Duration parsing

### Utils (Utilities)
- ✅ Password hashing/verification
- ✅ JWT generation/validation
- ✅ Token expiration handling
- ✅ Data validation helpers

### Models (Data models)
- ✅ Score calculation
- ✅ Data validation
- ✅ Level progression

## Test Data

### Mock data for rating
```go
[]models.UserRating{
    {UserID: 1, TotalWorkouts: 25, MaxConsecutive: 7, Score: 32},
    {UserID: 2, TotalWorkouts: 15, MaxConsecutive: 5, Score: 20},
}
```

### Test scenarios
- Basic rating calculation
- Error handling
- Data validation
- Algorithm performance

## CI/CD

GitHub Actions automatically runs:
- Unit tests
- Integration tests
- Linter
- Coverage report generation

## Best Practices

1. **Test isolation** - each test is independent
2. **Mock dependencies** - use mocks for external services
3. **Test data** - clear and understandable test cases
4. **Coverage** - aim for high coverage of critical code
5. **Performance** - benchmarks for important algorithms

## Adding New Tests

When adding new functionality:

1. Create unit tests for new functions
2. Add integration tests for new endpoints
3. Update mock objects when necessary
4. Add benchmarks for critical algorithms

## Test Debugging

```bash
# Run specific test
go test -v -run TestSpecificFunction ./internal/handlers

# Run with verbose output
go test -v ./...

# Run with race detection
go test -race ./...
```