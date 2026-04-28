# Blog UZ — REST API

Backend для блог-платформы на Go с использованием Echo, GORM и PostgreSQL.

## Стек технологий

- **Go** 1.24
- **Echo** v4 — веб-фреймворк
- **GORM** — ORM для работы с БД
- **PostgreSQL** 15 — база данных
- **Swagger** — документация API
- **Docker / Docker Compose** — контейнеризация

## Структура проекта

```
backend/
├── cmd/
│   └── main.go                  # Точка входа
├── src/
│   ├── core/
│   │   └── domain/
│   │       ├── entity/          # Доменные сущности
│   │       └── application/
│   │           └── usecases/    # Бизнес-логика
│   ├── entrypoint/
│   │   ├── groups/              # Группировка маршрутов
│   │   └── http/handlers/       # HTTP обработчики
│   └── infrastructure/
│       ├── config/              # Подключение к БД
│       └── persistence/         # Репозитории, маперы, модели
├── env/
│   └── .env                     # Переменные окружения
├── docs/                        # Swagger (генерируется автоматически)
├── Dockerfile
└── docker-compose.yml
```

## Быстрый старт

### 1. Клонируй репозиторий

```bash
git clone <repo-url>
cd blog_uz/backend
```

### 2. Создай файл переменных окружения

Создай файл `env/.env`:

```env
DB_HOST=db
DB_USER=bloguser
DB_PASSWORD=blogpassword
DB_NAME=blogdb
DB_PORT=5432
JWT_SECRET=supersecretkey123

POSTGRES_USER=bloguser
POSTGRES_PASSWORD=blogpassword
POSTGRES_DB=blogdb
```

### 3. Запусти через Docker Compose

```bash
docker-compose up --build
```

Приложение поднимется на `http://localhost:8080`

### 4. Swagger документация

```
http://localhost:8080/swagger/index.html
```

## API Эндпоинты

### Auth

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/auth/register` | Регистрация пользователя |
| POST | `/auth/login` | Вход в систему |

**Регистрация — тело запроса:**
```json
{
  "username": "maruf",
  "email": "maruf@example.com",
  "password": "secret123"
}
```

**Вход — тело запроса:**
```json
{
  "email": "maruf@example.com",
  "password": "secret123"
}
```

---

### Posts

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/posts` | Создать пост |
| GET | `/posts` | Получить все посты |
| PUT | `/posts/:id` | Обновить пост |
| DELETE | `/posts/:id` | Удалить пост |

**Создание поста — тело запроса:**
```json
{
  "user_id": 1,
  "title": "Мой первый пост",
  "content": "Содержимое поста"
}
```

---

### Comments

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/comments` | Создать комментарий |
| DELETE | `/comments/:id` | Удалить комментарий |

**Создание комментария — тело запроса:**
```json
{
  "post_id": 1,
  "user_id": 1,
  "content": "Отличный пост!"
}
```

---

### Likes

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/likes` | Поставить лайк |
| DELETE | `/likes/:id` | Убрать лайк |

**Лайк — тело запроса:**
```json
{
  "user_id": 1,
  "post_id": 1
}
```

---

### Profile

| Метод | URL | Описание |
|-------|-----|----------|
| PUT | `/profile/:id` | Обновить профиль |

**Обновление профиля — тело запроса:**
```json
{
  "full_name": "Маруф Каримов",
  "bio": "Go разработчик",
  "avatar": "https://example.com/avatar.png"
}
```

## Переменные окружения

| Переменная | Описание |
|------------|----------|
| `DB_HOST` | Хост базы данных |
| `DB_USER` | Пользователь БД |
| `DB_PASSWORD` | Пароль БД |
| `DB_NAME` | Название БД |
| `DB_PORT` | Порт БД (обычно 5432) |
| `JWT_SECRET` | Секретный ключ для JWT |
| `POSTGRES_USER` | Пользователь для инициализации PostgreSQL |
| `POSTGRES_PASSWORD` | Пароль для инициализации PostgreSQL |
| `POSTGRES_DB` | БД для инициализации PostgreSQL |

## Полезные команды

```bash
# Запуск
docker-compose up --build

# Остановка
docker-compose down

# Остановка с удалением данных БД
docker-compose down -v

# Логи приложения
docker-compose logs -f app

# Логи базы данных
docker-compose logs -f db
```
