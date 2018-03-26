package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "frrpc/protoFile"
	"frrpc/services/redisCache"
)

const (
	port = "114.55.248.175:50051"
)

type server struct {}

/**
 * 写入redis缓存服务
 */
func (s *server) RedisCache(ctx context.Context, in *pb.RedisCacheRequest) (reply *pb.RedisCacheReply , err error) {
	reply , err = redisCache.RedisCache(in.Name , in.Express , in.Value)
	return
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}