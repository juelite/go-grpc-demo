package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "frrpc/protoFile"
	"frrpc/services/redisCache"
	"frrpc/services/frLog"
	"frrpc/services"
)



type server struct {}

/**
 * 写入redis缓存服务
 */
func (s *server) RedisCache(ctx context.Context, in *pb.RedisCacheRequest) (reply *pb.RedisCacheReply , err error) {
	reply , err = redisCache.RedisCache(in.Name , in.Express , in.Value)
	return
}

func (s *server) GetCache(ctx context.Context, in *pb.GetCacheRequest) (reply *pb.GetCacheReply , err error) {
	reply , err = redisCache.GetCache(in.Name)
	return
}

/**
 * 写kibana日志服务
 */
func (s *server) FrLog(ctx context.Context , in *pb.FrLogRequest) (reply *pb.FrLogReply , err error) {
	reply , err = frLog.WriteLog(in.Tag , in.Info , in.Level)
	return
}

func main() {
	base := services.BaseService{}
	port := base.GetVal("rpcserve")
	base.LogInfo("index" , port)
	lis , err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}