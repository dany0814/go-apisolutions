package mongodb

import "time"

const (
	userCollection = "users"
)

type DocUser struct {
	ID        string     `bson:"id"`
	Name      string     `bson:"name"`
	Lastname  string     `bson:"lastname"`
	Email     string     `bson:"email"`
	Password  string     `bson:"password"`
	UserName  string     `bson:"username"`
	Phone     string     `bson:"phone,omitempty"`
	State     string     `bson:"state,omitempty"`
	CreatedAt time.Time  `bson:"created_at,omitempty"`
	UpdatedAt time.Time  `bson:"updated_at,omitempty"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
