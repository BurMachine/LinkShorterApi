package grpcHandlers

import (
	pb "burmachine/LinkGenerator/gen/go/protos"
	"burmachine/LinkGenerator/internal/storage"
)

type GrpcHandlers struct {
	pb.UnimplementedServiceNameServer
	Storage *storage.ServiceStorage
}
