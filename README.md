# Go Application Boilerplate

### e2e

Создание тестовой базы данных

```shell
make test-db
```

Запуск тестов

```shell
go test -v ./e2e
```

Пример теста

```go
package e2e

import (
    "goapp/e2e/app"
    "goapp/e2e/expect"
    "net/http"
    "testing"
)

func TestHealth(t *testing.T) {
    response := app.HTTPGet("/health")
    expect.HTTPStatusCode(t, response, http.StatusOK)
    expect.HTTPBodyJSONFragment(t, response, `
        {
            "status": "UP",
            "details": {
                "database": {
                    "status": "UP"
                }
            }
        }
    `)
}
```

Дебаг ответа

```go
func TestHealth(t *testing.T) {
    // response := app.HTTPGet("/health")
    // assert.HTTPStatus(t, response, http.StatusOK)
    app.HTTPGetDebug("/health")
}
```
```shell
=== RUN   TestHealth

Caller   : /mnt/hdata/code/go/goapp/e2e/health_test.go:12
Request  : GET /health
Payload  :
Status   : 200 OK
Response :
{
    "status": "UP",
    "details": {
        "database": {
            "status": "UP"
        }
    }
}

FAIL    goapp/e2e       0.015s
FAIL
```
