package grpcHandlers

import (
	pb "burmachine/LinkGenerator/gen/go/protos"
	"context"
)

func (s *GrpcHandlers) GetOriginalLink(context.Context, *pb.RequestLink) (*pb.ResponseLink, error) {
	println("qwe")
	return nil, nil
}
