package grpcHandlers

import (
	pb "burmachine/LinkGenerator/gen/go/protos"
	"golang.org/x/mod/sumdb/storage"
)

type GrpcHandlers struct {
	pb.UnimplementedServiceNameServer
	// server  *server2.Server
	storage *storage.Storage
}
