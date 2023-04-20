package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	web "github.com/dany0814/go-apisolutions/internal/adapters/driver"
	"github.com/dany0814/go-apisolutions/internal/core/application"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Service AppService
	Channel AppChannel
}

type AppService struct {
	UserService application.UserService
}

type AppChannel struct {
	ChannelOut chan []byte
	ChannelIn  []chan []byte
}

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	ShutdownTimeout time.Duration
	app             Application
}

func NewServer(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, app Application) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		ShutdownTimeout: shutdownTimeout,
		app:             app,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	// Routes user
	uh := web.NewUserHandler(s.app.Service.UserService)
	s.engine.POST("/user/sigin", uh.SignInHandler())
	s.engine.POST("/user/login", uh.LoginHandler())
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
