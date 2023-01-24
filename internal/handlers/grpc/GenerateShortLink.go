package grpcHandlers

import "context"
import pb "burmachine/LinkGenerator/gen/go/protos"

func (s *GrpcHandlers) MyMethod(ctx context.Context, req *pb.RequestLink) (*pb.ResponseLink, error) {
	println("ПРивет жрпц")
	return &pb.ResponseLink{Link: "Привет"}, nil
}
