# Testing Guide

## Overview
This project is covered by various types of tests to ensure code quality and reliability.

## Test Structure

### Unit Tests
- `internal/handlers/handlers_test.go` - tests for HTTP handlers
- `internal/services/ai_test.go` - tests for AI service
- `internal/repository/mongodb_test.go` - tests for MongoDB repository
- `internal/models/models_test.go` - tests for data models
- `pkg/utils/crypto_test.go` - tests for cryptographic utilities
- `pkg/utils/jwt_test.go` - tests for JWT utilities

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
- ✅ Rating endpoint
- ✅ Error handling

### Services (Business logic)
- ✅ AI Service rating functionality
- ✅ Error handling
- ✅ Mock dependencies

### Repository (Data layer)
- ✅ MongoDB rating calculation
- ✅ Consecutive days algorithm
- ✅ Level calculation logic

### Utils (Utilities)
- ✅ Password hashing/verification
- ✅ JWT generation/validation
- ✅ Token expiration handling

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