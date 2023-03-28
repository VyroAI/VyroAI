package repo

import "context"

type BcryptRepo interface {
	CompareHashAndPassword(ctx context.Context, hashedPassword, password string) error
	GenerateFromPassword(ctx context.Context, password string) (string, error)
}
