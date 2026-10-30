# ПЗ №9 — Регистрация и вход пользователей, bcrypt

Студент: Абдулатифзода Акмал, группа ЭФМО-0126

Сервис `pz9-auth` на Go (chi + GORM + PostgreSQL): регистрация с хэшированием пароля через bcrypt и вход с проверкой пароля.

## Структура
```
cmd/api/main.go                     — конфиг, подключение к БД, AutoMigrate, маршруты
internal/platform/config/config.go  — переменные окружения DB_DSN, BCRYPT_COST, APP_ADDR
internal/repo/postgres.go           — подключение GORM к PostgreSQL
internal/repo/user_repo.go          — AutoMigrate, Create, ByEmail
internal/core/user.go               — модель User (PasswordHash не отдаётся в JSON)
internal/http/handlers/auth.go      — Register / Login, bcrypt
```

## Запуск
```bash
# PostgreSQL: база pz9 и пользователь
sudo -u postgres psql -c "CREATE USER akmal WITH PASSWORD 'akmal'"
sudo -u postgres psql -c "CREATE DATABASE pz9 OWNER akmal"

export DB_DSN="postgres://akmal:akmal@localhost:5432/pz9?sslmode=disable"
export BCRYPT_COST=12
export APP_ADDR=:8080
go mod tidy
go run ./cmd/api
```
Windows PowerShell: `$env:DB_DSN="postgres://akmal:akmal@localhost:5432/pz9?sslmode=disable"`.

Таблица `users` создаётся автоматически (GORM AutoMigrate).

## Проверка
```bash
# регистрация -> 201, повтор -> 409
curl -i -X POST http://localhost:8080/auth/register -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Secret123!"}'
# вход -> 200, неверный пароль -> 401
curl -i -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Secret123!"}'
```

| Сценарий | Код | Ответ |
|---|---|---|
| Успешная регистрация | 201 | `{"status":"ok","user":{...}}` |
| Email уже занят | 409 | `{"error":"email_taken"}` |
| Неверный email / пароль < 8 | 400 | `{"error":"email_required_and_password_min_8"}` |
| Успешный вход | 200 | `{"status":"ok","user":{...}}` |
| Неверный пароль или email | 401 | `{"error":"invalid_credentials"}` |
