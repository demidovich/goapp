# Go Application Boilerplate

### Запуск

Docker
```shell
make up
```

Локальный Go
```shell
make up-local
go run ./cmd/rest/main.go
```

### Логирование

Используется `log/slog`. Параметры логгера:
- **Encoding**
- **Level**

Encoding:
- **json** - ECS схема
- **text** - ECS схема key-value
- **pretty** - вывод через [lmittmann/tint](https://github.com/lmittmann/tint), цветной читаемый вывод для локальной разработки

Вывод JSON

![logger json](/docs/assets/logger_json.png)

Вывод Pretty

![logger pretty](/docs/assets/logger_pretty.png)

### Миграции

Используется [pressly/goose](https://github.com/pressly/goose)

### Консольный интерфейс

```shell
go run cmd/cli/main.go
```

Команды

```shell
make:migration Создать файл миграции
migrate        Применить миграции базы данных
migrate:down   Откатить миграции базы данных
migrate:status Состояние миграций базы данных
```

### e2e

Используется [steinfletcher/apitest](https://github.com/steinfletcher/apitest)

Для запуска тестов необходимо создать тестовую базу данных
```shell
make up-local
make test-db
```

Запуск тестов
```shell
go test -v ./e2e
```

Пример теста

```go
func TestHealth(t *testing.T) {
    body := `{
        "status": "UP",
        "details": {
            "database": {
                "status": "UP"
            }
        }
    }`

    app.API().
        Get("/health").
        // Debug(). // Дебаг ответа
        Expect(t).
        Status(http.StatusOK).
        Body(body).
        End()
}
```
