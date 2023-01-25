package server

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	gw "burmachine/LinkGenerator/gen/go/protos"
	"burmachine/LinkGenerator/internal/config"
	grpcHandlers "burmachine/LinkGenerator/internal/handlers/grpc"
	httpHandlers "burmachine/LinkGenerator/internal/handlers/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	httpHandles *httpHandlers.HttpHandlers
	conf        config.Conf
	Mux         *runtime.ServeMux
	GrpcServ    *grpc.Server
	GrpcHadles  *grpcHandlers.GrpcHandlers
}

func NewServerWithConfiguration(conf config.Conf) *Server {
	return &Server{conf: conf}
}

func (s *Server) Run(ctx context.Context) error {
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	grpcServerEndpoint := flag.String("grpc-server-endpoint", s.conf.AddrGrpc, "gRPC server endpoint")

	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gw.RegisterServiceNameServer(grpcServer, s.GrpcHadles)
	grpcServer.Serve(lis)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(s.conf.AddrHttp, s.Mux)
}
