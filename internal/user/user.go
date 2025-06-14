package user

import (
	"goapp/pkg/logger"

	"github.com/demidovich/failure"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Find(log *logger.Logger, id int) (User, error) {
	log.Infof("Поиск пользователя с ID %d", id)
	u := User{}

	if id < 10 {
		return u, failure.New("user not found")
	}

	u.Name = "Test User"
	u.Email = "test@mail.ru"

	return u, nil
}
