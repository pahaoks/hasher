package servers

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Wrapper for grpc server. Incapsulates common logic
type GrpcServer struct {
	options GrpcOptions
	server  *grpc.Server
	cancel  context.CancelFunc
}

// Options
type GrpcOptions struct {
	Port     string
	Register func(s grpc.ServiceRegistrar)
}

// Instantiate server
func NewGrpcServer(options GrpcOptions) *GrpcServer {
	return &GrpcServer{options, grpc.NewServer(), nil}
}

// Start server
func (s *GrpcServer) Start(ctx context.Context) {
	if s.options.Register == nil {
		log.Fatal("grpc: register function is nil")
	}

	s.options.Register(s.server)

	listener := initTcpListener(s.options.Port)

	ctx, s.cancel = context.WithCancel(ctx)

	go onShutdown(
		ctx,
		func() {
			log.Println("grpc: server has been stopped")

			s.server.Stop()
		},
	)

	log.Printf("grpc: server is listening at %v", listener.Addr())
	if err := s.server.Serve(listener); err != nil {
		log.Fatalf("grpc: failed to serve: %v", err)
	}
}

// Stop server
func (s *GrpcServer) Stop() {
	s.cancel()
}

// Init tcp listener
func initTcpListener(port string) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("grpc: failed to listen: %v", err)
	}

	return lis
}
