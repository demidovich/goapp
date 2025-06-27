package profile2

import "context"

type Repository interface {
	NextID(ctx context.Context) (int, error)
	Create(ctx context.Context, p *Profile) error
	SearchByEmail(ctx context.Context, email string) (*Profile, error)
}
