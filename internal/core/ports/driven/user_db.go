package drivenport

import (
	"context"

	"github.com/dany0814/go-apisolutions/internal/core/application/dto"
	"github.com/dany0814/go-apisolutions/internal/core/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (*dto.User, error)
}
