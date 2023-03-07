package grpc

import (
	"log"

	"google.golang.org/grpc"
)

func GetConnGrpc(url string) grpc.ClientConnInterface {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	return conn
}
