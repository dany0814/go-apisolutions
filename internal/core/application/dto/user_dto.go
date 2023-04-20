package dto

import "time"

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	UserName  string     `json:"username"`
	Phone     string     `json:"phone,omitempty"`
	State     string     `json:"state,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Phone    string `json:"phone,omitempty"`
	Token    string `json:"token"`
}
