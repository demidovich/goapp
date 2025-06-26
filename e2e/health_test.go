package e2e2

import (
	"goapp/e2e/app"
	"net/http"
	"testing"
)

func Test_Health(t *testing.T) {
	body := `{
		"status": "UP",
		"details": {
			"database": {
				"status": "UP"
			}
		}
	}`

	app.API().
		// Debug().
		// Report(apitest.SequenceDiagram()).
		Get("/health").
		Expect(t).
		Status(http.StatusOK).
		Body(body).
		End()
}
