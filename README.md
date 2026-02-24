# Task Manager API (Go + Gin + PostgreSQL)

##  О проекте

REST API для управления задачами с JWT-аутентификацией.

Проект реализован как pet-project для демонстрации навыков backend-разработки на Go.

## Стек технологий

- Go
- Gin
- PostgreSQL
- JWT
- Docker
- bcrypt

## Архитектура

Проект использует layered architecture:

- **handler** — HTTP слой
- **service** — бизнес-логика
- **repository** — работа с БД
- **middleware** — JWT авторизация

## База данных

### users

| Поле | Тип | Описание |
|------|-----|----------|
| id | SERIAL | Primary key |
| email | TEXT | Unique |
| password_hash | TEXT | Hashed пароль |
| created_at | TIMESTAMP | Дата создания |

### tasks

| Поле | Тип | Описание |
|------|-----|----------|
| id | SERIAL | Primary key |
| title | TEXT | Название задачи |
| description | TEXT | Описание |
| status | TEXT | Статус (pending/in_progress/done) |
| user_id | INTEGER | Foreign key → users(id) |
| created_at | TIMESTAMP | Дата создания |

**Связь:** один пользователь — много задач.

## Как запустить

```bash
docker compose up --build
```

Сервер будет доступен на: **http://localhost:8080**

## Аутентификация

### Регистрация

**POST** `/auth/register`

```json
{
  "email": "test@mail.com",
  "password": "123456"
}
```

### Логин

**POST** `/auth/login`

**Ответ:**

```json
{
  "token": "JWT_TOKEN"
}
```


## Работа с задачами

Все запросы требуют заголовок:

```Authorization: Bearer JWT_TOKEN```

### Создать задачу

**POST** `/tasks`

```json
{
  "title": "Learn Go",
  "description": "Finish pet project"
}
```



### Получить задачи

**GET** `/tasks`

### Обновить задачу

**PUT** `/tasks/{id}`

### Удалить задачу

**DELETE** `/tasks/{id}`


## Возможные ошибки

| Код | Ответ | Описание |
|-----|-------|----------|
| 400 | `{"error": "invalid request"}` | Неверный формат JSON |
| 401 | `{"error": "missing token"}` | Нет токена или он неверный |
| 404 | `{"error": "task not found"}` | Задача с таким ID не найдена |
| 500 | `{"error": "internal server error"}` | Ошибка на сервере |

## Что реализовано

- ✅ JWT авторизация
- ✅ Хеширование паролей (bcrypt)
- ✅ CRUD операции
- ✅ PostgreSQL
- ✅ Docker
- ✅ Индексы и foreign keys
- ✅ Чистая архитектура (layered architecture)
