# Account Service

Микросервис для управления аккаунтами пользователей с поддержкой gRPC и Kafka.

## Быстрый старт

### Предварительные требования

- Go 1.24.5+
- PostgreSQL
- Apache Kafka

### Установка зависимостей

```bash
make deps
```

### Настройка окружения

Создайте файл `.env` в корне проекта:

```bash
cp .env.example .env
```

### Запуск

#### Локальная разработка

```bash
# Сборка и запуск
make run

# Или только запуск (если уже собран)
make run-only
```

#### Docker

```bash
# Сборка образа
docker build -t account-service .

# Запуск контейнера
docker run -p 50051:50051 --env-file .env account-service
```

### Полезные команды

```bash
# Форматирование кода
make fmt

# Линтинг
make lint

# Тесты
make test

# Тесты с покрытием
make test-coverage

# Очистка артефактов сборки
make clean

# Показать все доступные команды
make help
```
