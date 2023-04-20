package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dany0814/go-apisolutions/internal/core/application/dto"
	"github.com/dany0814/go-apisolutions/internal/core/domain"
	outdb "github.com/dany0814/go-apisolutions/internal/core/ports/driven"
	"github.com/dany0814/go-apisolutions/pkg/appjwt"
	"github.com/dany0814/go-apisolutions/pkg/encryption"
	"github.com/dany0814/go-apisolutions/pkg/uidgen"
)

type UserService struct {
	userDB outdb.UserDB
}

func NewUserService(userDB outdb.UserDB) UserService {
	return UserService{
		userDB: userDB,
	}
}

func (usrv UserService) Register(ctx context.Context, user dto.User) (*dto.User, error) {
	id := uidgen.New().New()

	newuser, err := domain.NewUser(id, user.Name, user.Lastname, user.Email, user.Password, user.UserName, user.Phone, "Active")

	if err != nil {
		return nil, err
	}

	found, err := usrv.userDB.FindByEmail(ctx, user.Email)

	if !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}

	if found != nil {
		return nil, fmt.Errorf("%w: %s", domain.ErrUserConflict, err)
	}

	pass, err := encryption.HashAndSalt(user.Password)

	if err != nil {
		return nil, err
	}

	vopass, _ := domain.NewUserPassword(pass)

	newuser.Password = vopass
	newuser.CreatedAt = time.Now()
	newuser.UpdatedAt = time.Now()

	if err := usrv.userDB.Create(ctx, newuser); err != nil {
		return nil, err
	}

	user.Password = ""
	user.ID = id

	return &user, nil
}

func (usrv UserService) Login(ctx context.Context, login dto.LoginReq) (*dto.LoginRes, error) {
	found, err := usrv.userDB.FindByEmail(ctx, login.Email)

	if err != nil {
		return nil, err
	}

	if !encryption.ComparePasswords(found.Password, login.Password) {
		return nil, fmt.Errorf("%w: %s", domain.ErrInvalidPassword, err)
	}

	token, err := appjwt.CreateToken(&found.ID)
	if err != nil {
		return nil, err
	}

	res := dto.LoginRes{
		ID:       found.ID,
		Name:     found.Name,
		Lastname: found.Lastname,
		Email:    found.Email,
		UserName: found.UserName,
		Phone:    found.Phone,
		Token:    token,
	}

	return &res, nil
}
