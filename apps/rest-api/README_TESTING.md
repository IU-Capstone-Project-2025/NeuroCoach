# Testing Guide

## Overview
Этот проект покрыт различными типами тестов для обеспечения качества и надежности кода.

## Структура тестов

### Unit Tests (Модульные тесты)
- `internal/handlers/handlers_test.go` - тесты для HTTP handlers
- `internal/services/ai_test.go` - тесты для AI сервиса
- `internal/repository/mongodb_test.go` - тесты для MongoDB repository
- `internal/models/models_test.go` - тесты для моделей данных
- `pkg/utils/crypto_test.go` - тесты для криптографических утилит
- `pkg/utils/jwt_test.go` - тесты для JWT утилит

### Integration Tests (Интеграционные тесты)
- `test_integration.go` - интеграционные тесты для API endpoints

### Benchmark Tests (Бенчмарки)
- Тесты производительности для критических алгоритмов

## Запуск тестов

### Все тесты
```bash
make test
```

### Только unit тесты
```bash
make test-unit
```

### Интеграционные тесты
```bash
make test-integration
```

### Тесты с покрытием
```bash
make test-coverage
```

### Бенчмарки
```bash
make benchmark
```

## Настройка тестовой среды

### Тестовые базы данных
```bash
make test-setup    # Создать тестовые БД
make test-teardown # Удалить тестовые БД
```

## Покрытие кода

Основные компоненты покрыты тестами:

### Handlers (HTTP обработчики)
- ✅ Register/Login endpoints
- ✅ Rating endpoint
- ✅ Error handling

### Services (Бизнес-логика)
- ✅ AI Service rating functionality
- ✅ Error handling
- ✅ Mock dependencies

### Repository (Работа с данными)
- ✅ MongoDB rating calculation
- ✅ Consecutive days algorithm
- ✅ Level calculation logic

### Utils (Утилиты)
- ✅ Password hashing/verification
- ✅ JWT generation/validation
- ✅ Token expiration handling

### Models (Модели данных)
- ✅ Score calculation
- ✅ Data validation
- ✅ Level progression

## Тестовые данные

### Mock данные для рейтинга
```go
[]models.UserRating{
    {UserID: 1, TotalWorkouts: 25, MaxConsecutive: 7, Score: 32},
    {UserID: 2, TotalWorkouts: 15, MaxConsecutive: 5, Score: 20},
}
```

### Тестовые сценарии
- Базовый расчет рейтинга
- Обработка ошибок
- Валидация данных
- Производительность алгоритмов

## CI/CD

GitHub Actions автоматически запускает:
- Unit тесты
- Интеграционные тесты
- Линтер
- Генерацию отчетов покрытия

## Лучшие практики

1. **Изоляция тестов** - каждый тест независим
2. **Mock зависимости** - используем mock для внешних сервисов
3. **Тестовые данные** - четкие и понятные тестовые случаи
4. **Покрытие** - стремимся к высокому покрытию критического кода
5. **Производительность** - бенчмарки для важных алгоритмов

## Добавление новых тестов

При добавлении нового функционала:

1. Создайте unit тесты для новых функций
2. Добавьте интеграционные тесты для новых endpoints
3. Обновите mock объекты при необходимости
4. Добавьте бенчмарки для критических алгоритмов

## Отладка тестов

```bash
# Запуск конкретного теста
go test -v -run TestSpecificFunction ./internal/handlers

# Запуск с подробным выводом
go test -v ./...

# Запуск с race detection
go test -race ./...
```