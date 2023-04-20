package bootstrap

import (
	"context"
	"log"

	database "github.com/dany0814/go-apisolutions/internal/adapters/driven"
	"github.com/dany0814/go-apisolutions/internal/core/application"
	"github.com/dany0814/go-apisolutions/internal/platform/server"
	"github.com/dany0814/go-apisolutions/internal/platform/storage/mongodb"
	"github.com/dany0814/go-apisolutions/pkg/config"
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	ctx := context.Background()
	db, err := config.ConfigDb(ctx)

	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	userRepository := mongodb.NewUserRepository(db, config.Cfg.DbTimeout)
	userAdapter := database.NewUserAdapter(userRepository)
	userService := application.NewUserService(userAdapter)

	appService := server.AppService{
		UserService: userService,
	}

	ctx, srv := server.NewServer(context.Background(), config.Cfg.Host, config.Cfg.Port, config.Cfg.ShutdownTimeout, server.Application{
		Service: appService,
	})

	return srv.Run(ctx)
}
