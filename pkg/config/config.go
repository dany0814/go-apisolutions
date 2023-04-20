package config

import (
	"context"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	// Server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
	// Database config
	DbUser    string        `default:"apisolutions"`
	DbPass    string        `default:"apisolutions"`
	DbHost    string        `default:"database"`
	DbPort    uint          `default:"27017"`
	DbName    string        `default:"apisolutions"`
	DbTimeout time.Duration `default:"10s"`
	// Security token
	Secret      string `default:"apisolutions"`
	ExpiredHour int64  `efault:"2"`
}

var Cfg config

// load environment variables
func LoadConfig() error {
	err := envconfig.Process("ENV", &Cfg)
	if err != nil {
		return err
	}
	return nil
}

// Init a mongodb client using a environment variables
func ConfigDb(ctx context.Context) (*mongo.Database, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", Cfg.DbUser, Cfg.DbPass, Cfg.DbHost, Cfg.DbPort, Cfg.DbName)
	fmt.Println("Config environment variables: ", mongoURI)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(Cfg.DbName)

	return db, nil
}
