package server

import (
	"burmachine/LinkGenerator/internal/config"
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	gw "burmachine/LinkGenerator/gen/go/protos"
)

type Server struct {
	conf config.Conf
}

func NewServerWithConfiguration(conf config.Conf) *Server {
	return &Server{conf: conf}
}

func (s Server) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServerEndpoint := flag.String("grpc-server-endpoint", s.conf.AddrGrpc, "gRPC server endpoint")

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterServiceNameHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(s.conf.AddrHttp, mux)
}
