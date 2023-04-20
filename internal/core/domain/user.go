package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/dany0814/go-apisolutions/pkg/uidgen"
)

var ErrInvalidUserID = errors.New("invalid User ID")
var ErrEmptyUserName = errors.New("the field username is required")
var ErrEmptyName = errors.New("the field name is required")
var ErrUserNotFound = errors.New("user not found")
var ErrUserConflict = errors.New("user already exists")
var ErrInvalidPassword = errors.New("user invalid password")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uidgen.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}
	return UserID{
		value: v,
	}, nil
}

func (id UserID) String() string {
	return id.value
}

var ErrInvalidUserEmail = errors.New("invalid Email")

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserEmail{}, fmt.Errorf("%w: %s", ErrInvalidUserEmail, value)
	}
	return UserEmail{
		value: value,
	}, nil
}

func (email UserEmail) String() string {
	return email.value
}

var ErrInvalidUserPassword = errors.New("invalid Password")

type UserPassword struct {
	value string
}

func NewUserPassword(value string) (UserPassword, error) {
	if value == "" {
		return UserPassword{}, fmt.Errorf("%w: %s", ErrInvalidUserPassword, value)
	}
	return UserPassword{
		value: value,
	}, nil
}

func (pass UserPassword) String() string {
	return pass.value
}

var ErrEmptyUserUsername = errors.New("the field Username is required")

type UserUsername struct {
	value string
}

func NewUserUsername(value string) (UserUsername, error) {
	if value == "" {
		return UserUsername{}, fmt.Errorf("%w: %s", ErrEmptyUserUsername, value)
	}
	return UserUsername{
		value: value,
	}, nil
}

func (usrname UserUsername) String() string {
	return usrname.value
}

type User struct {
	ID        UserID
	Name      string
	Lastname  string
	Email     UserEmail
	Password  UserPassword
	UserName  UserUsername
	Phone     string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(userID, name, lastname, email, password, username, phone, state string) (User, error) {
	idVo, err := NewUserID(userID)
	if err != nil {
		return User{}, err
	}

	if name == "" {
		return User{}, fmt.Errorf("%w: %s", ErrEmptyName, name)
	}

	emailVo, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	passwordVo, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	userNameVo, err := NewUserUsername(username)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       idVo,
		Name:     name,
		Lastname: lastname,
		Email:    emailVo,
		Password: passwordVo,
		UserName: userNameVo,
		Phone:    phone,
		State:    state,
	}, nil
}
