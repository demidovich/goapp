package e2e2

import (
	"goapp/e2e/app"
	factory "goapp/internal/testfactory"
	"net/http"
	"testing"
)

func Test_ProfileCreate(t *testing.T) {
	payload := params{
		"email":    factory.UniqueEmail(),
		"password": "123456789",
	}

	app.API().
		Post("/profile").
		JSON(payload).
		Expect(t).
		Status(http.StatusOK).
		End()

	app.ExpectDatabaseHas(t, "profile", params{
		"email": payload["email"],
	})
}

func Test_ProfileCreateExistiong(t *testing.T) {
	// Необходимо переделать ошибки
	// Здесь должно возвращаться ErrValidation или что-то в этом духе
	t.Skip()

	profile := app.Factory().Profile().New(t)

	payload := params{
		"email":    profile.Email,
		"password": "123456789",
	}

	app.API().
		Debug().
		Post("/profile").
		JSON(payload).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
