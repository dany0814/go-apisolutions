package repository

import (
	"context"

	"github.com/dany0814/go-apisolutions/internal/platform/storage/mongodb"
)

type UserRepository interface {
	Save(ctx context.Context, docUser mongodb.DocUser) error
	FindByEmail(ctx context.Context, email string) (*mongodb.DocUser, error)
	FindById(ctx context.Context, id string) (*mongodb.DocUser, error)
	FindAll(ctx context.Context) ([]*mongodb.DocUser, error)
}
