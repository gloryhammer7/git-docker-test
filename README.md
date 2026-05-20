# Lilia Test Server

Простой Go-сервер с одним POST эндпоинтом.

## Запуск

Сервер запустится на `http://localhost:8081`. Порт 8081

### Docker

Сборка образа:

```bash
docker build -t lilia-server .
```

Запуск контейнера:

```bash
docker run -p 8081:8081 lilia-server
```

## API

### GET /

Проверка работоспособности сервера.

### POST /api/hello

Принимает POST-запрос и возвращает приветствие.
