package app

import (
	"context"

	"example.com/hasher/gen/oapi"
	"example.com/hasher/gen/proto"
	"example.com/hasher/internal/config"
	"example.com/hasher/internal/domain"
	"example.com/hasher/internal/repositories"
	"example.com/hasher/internal/services"
	"example.com/hasher/pkg/servers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Main application
type App struct {
	stop       context.CancelFunc
	httpServer *servers.HttpServer
	grpcServer *servers.GrpcServer
}

// Hasher dependency interface
type HasherService interface {
	GetHash() *domain.Hash
}

// Init app
func New() *App {
	cfg := config.Init()

	hasherRepository := repositories.NewHasher()
	hasherService := services.NewHasher(
		cfg.HashTTL,
		hasherRepository,
	)

	httpHandler := NewHttpHandler(hasherService)
	grpcHandler := NewGrpcHandler(hasherService)

	grpcServer := servers.NewGrpcServer(servers.GrpcOptions{
		Port: cfg.GrpcPort,
		Register: func(registrar grpc.ServiceRegistrar) {
			proto.RegisterHasherServer(
				registrar,
				grpcHandler,
			)
		},
	})

	httpServer := servers.NewHttpServer(servers.HttpOptions{
		Port: cfg.HttpPort,
		Register: func(r *mux.Router) {
			oapi.HandlerFromMux(httpHandler, r)
		},
	})

	return &App{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

// Start app
func (a *App) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	a.stop = cancel

	go a.grpcServer.Start(ctx)
	go a.httpServer.Start(ctx)

	<-ctx.Done()
}

// Stop app
func (a *App) Stop() {
	a.stop()
}
