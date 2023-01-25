package server

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

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
	ErrorChan   chan error
}

func NewServerWithConfiguration(conf config.Conf) *Server {
	return &Server{conf: conf}
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) error {

	grpcServerEndpoint := flag.String("grpc-server-endpoint", s.conf.AddrGrpc, "gRPC server endpoint")

	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gw.RegisterServiceNameServer(grpcServer, s.GrpcHadles)
	go func(ctx context.Context) {
		err = grpcServer.Serve(lis)
		if err != nil {
			s.ErrorChan <- err
		}
		wg.Done()
	}(ctx)
	go func(ctx context.Context) {
		err = http.ListenAndServe(s.conf.AddrHttp, s.Mux)
		if err != nil {
			s.ErrorChan <- err
		}
		wg.Done()
	}(ctx)

	log.Println("[SERVER] - launched")
	return nil
}
