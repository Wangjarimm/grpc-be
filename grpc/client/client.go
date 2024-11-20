package client

import (
	"google.golang.org/grpc"
	"log"
	"grpc1/grpc/proto"  // Sesuaikan path-nya
)

// NewItemServiceClient membuat koneksi ke server gRPC
func NewItemServiceClient() (proto.ItemServiceClient, error) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		return nil, err
	}
	return proto.NewItemServiceClient(conn), nil
}
