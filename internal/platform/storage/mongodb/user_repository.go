package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dany0814/go-apisolutions/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db        *mongo.Database
	dbTimeout time.Duration
}

func NewUserRepository(db *mongo.Database, dbTimeout time.Duration) UserRepository {
	return UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r UserRepository) Save(ctx context.Context, user DocUser) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	found, err := r.FindById(ctx, user.ID)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	coll := r.db.Collection(userCollection)

	if found == nil {
		if _, err := coll.InsertOne(ctxTimeout, user); err != nil {
			return err
		}
		return nil
	}

	if _, err := coll.UpdateOne(ctxTimeout, bson.M{"id": user.ID}, bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"lastname":   user.Lastname,
			"email":      user.Email,
			"password":   user.Password,
			"phone":      user.Phone,
			"state":      user.State,
			"updated_at": user.UpdatedAt,
		},
	}); err != nil {
		return err
	}

	return nil
}

func (r UserRepository) FindByEmail(ctx context.Context, email string) (*DocUser, error) {
	var foudUser DocUser

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.Collection(userCollection).FindOne(ctxTimeout, bson.M{"email": email}).Decode(&foudUser)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("%w: %s", domain.ErrUserNotFound, err)
	}

	if err != nil {
		return nil, err
	}

	return &foudUser, nil
}

func (r UserRepository) FindById(ctx context.Context, id string) (*DocUser, error) {
	var foudUser DocUser

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	err := r.db.Collection(userCollection).FindOne(ctxTimeout, bson.M{"id": id}).Decode(&foudUser)

	if err != nil {
		return nil, err
	}

	return &foudUser, nil
}

func (r UserRepository) FindAll(ctx context.Context) ([]*DocUser, error) {
	var allUser []*DocUser

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	results, err := r.db.Collection(userCollection).Find(ctxTimeout, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctxTimeout)

	for results.Next(ctxTimeout) {
		var docuser DocUser
		if err := results.Decode(&docuser); err != nil {
			return nil, err
		}
		docuser.Password = ""
		allUser = append(allUser, &docuser)
	}

	return allUser, nil
}
