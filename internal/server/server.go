package server

import (
	"burmachine/LinkGenerator/internal/config"
	httpHandlers "burmachine/LinkGenerator/internal/handlers/http"
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	gw "burmachine/LinkGenerator/gen/go/protos"
)

type Server struct {
	httpHandles *httpHandlers.HttpHandlers
	conf        config.Conf
	Mux         *runtime.ServeMux
}

func NewServerWithConfiguration(conf config.Conf) *Server {
	return &Server{conf: conf}
}

func (s *Server) Run(ctx context.Context) error {
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	grpcServerEndpoint := flag.String("grpc-server-endpoint", s.conf.AddrGrpc, "gRPC server endpoint")

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterServiceNameHandlerFromEndpoint(ctx, s.Mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(s.conf.AddrHttp, s.Mux)
}
