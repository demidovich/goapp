package profile

import (
	"context"

	"github.com/demidovich/failure"
)

// type SearchFilters struct {
// 	Email string `validate:"omitempty,email"`
// }

// func (c *SearchFilters) Validate() error {
// 	return validation.ValidateStruct(c)
// }

func (u *Usecases) searchByEmail(ctx context.Context, email string) (Profile, error) {
	item := Profile{}
	if email == "" {
		return item, failure.New("передан пустой email")
	}

	err := u.db.GetContext(ctx, &item, "select * from profile where email = lower($1) limit 1", email)

	return item, err
}
