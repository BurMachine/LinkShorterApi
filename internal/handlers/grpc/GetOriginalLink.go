package grpcHandlers

import (
	"context"
	"log"

	pb "burmachine/LinkGenerator/gen/go/protos"
)

func (s *GrpcHandlers) GetOriginalLink(ctx context.Context, req *pb.RequestLink) (*pb.ResponseLink, error) {
	log.Println("GetOriginalLink un use")

	shortLink := req.Link
	originalLink, err := (*s.Storage).GetFullLink(shortLink)
	if err != nil {
		log.Println("Link does not exist: ", err)
		return &pb.ResponseLink{Link: "Incorrect link"}, nil
	}

	return &pb.ResponseLink{Link: originalLink}, nil
}
