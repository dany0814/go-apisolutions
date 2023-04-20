package drivenadapt

import (
	"context"

	"github.com/dany0814/go-apisolutions/internal/adapters/driven/repository"
	"github.com/dany0814/go-apisolutions/internal/core/application/dto"
	"github.com/dany0814/go-apisolutions/internal/core/domain"
	"github.com/dany0814/go-apisolutions/internal/platform/storage/mongodb"
)

type UserAdapter struct {
	userRepository repository.UserRepository
}

func NewUserAdapter(ur repository.UserRepository) UserAdapter {
	return UserAdapter{
		userRepository: ur,
	}
}

func (uad UserAdapter) Create(ctx context.Context, user domain.User) error {
	docUser := mongodb.DocUser{
		ID:        user.ID.String(),
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email.String(),
		Password:  user.Password.String(),
		UserName:  user.UserName.String(),
		Phone:     user.Phone,
		State:     user.State,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return uad.userRepository.Save(ctx, docUser)
}

func (uad UserAdapter) FindByEmail(ctx context.Context, email string) (*dto.User, error) {
	foudUser, err := uad.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	resUser := dto.User{
		ID:        foudUser.ID,
		Name:      foudUser.Name,
		Lastname:  foudUser.Lastname,
		Email:     foudUser.Email,
		Password:  foudUser.Password,
		UserName:  foudUser.UserName,
		Phone:     foudUser.Phone,
		State:     foudUser.State,
		CreatedAt: foudUser.CreatedAt,
		UpdatedAt: foudUser.UpdatedAt,
		DeletedAt: foudUser.DeletedAt,
	}

	return &resUser, nil
}
