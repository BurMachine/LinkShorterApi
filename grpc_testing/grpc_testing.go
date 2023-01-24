package main

import (
	pb "burmachine/LinkGenerator/gen/go/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds := credentials.NewTLS(nil)
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewServiceNameClient(conn)
	req := &pb.RequestLink{Link: "google.com"}
	res, err := client.GenerateShortLink(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling MyMethod: %v", err)
	}
	fmt.Println(res.Link)
}
