package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "frrpc/protoFile"
	"fmt"
)

const (
	address     = "114.55.248.175:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	request := &pb.RedisCacheRequest{
		Name:"wangyu",
		Express:200,
		Value:"hello",
	}
	r , err := c.RedisCache(context.Background(), request)
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	fmt.Println(r)
}