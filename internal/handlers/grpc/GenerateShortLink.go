package grpcHandlers

import (
	"context"
	"log"

	"burmachine/LinkGenerator/internal/service"
)
import pb "burmachine/LinkGenerator/gen/go/protos"

func (s *GrpcHandlers) GenerateShortLink(ctx context.Context, req *pb.RequestLink) (*pb.ResponseLink, error) {

	link := req.Link

	shortLink, err := service.GenerateLink(link)
	if err != nil {
		log.Println("Link generation error: ", err)
		return &pb.ResponseLink{Link: "Incorrect link"}, nil
	}
	err = (*s.Storage).AddShortLink(link, shortLink)
	if err != nil {
		log.Println("Link add error: ", err)
		return &pb.ResponseLink{Link: "Link already exist"}, nil
	}

	return &pb.ResponseLink{Link: shortLink}, nil
}
