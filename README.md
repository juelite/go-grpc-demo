# grpc微服务 写redis缓存 demo
### 依赖包
- golang.org/x/net/context
- google.golang.org/grpc
- github.com/Unknwon/goconfig //为了读取配置文件
- google.golang.org/grpc
- github.com/golang/protobuf/proto

### protoc 工具需先安装 
```
    https://github.com/google/protobuf/releases/download/v2.4.1/protobuf-2.4.1.tar.gz
```
  
### 目录结构
```
    ├── README.md
    ├── client
    │   └── client.go 客户端
    ├── conf
    │   ├── env.conf 运行环境
    │   └── setting.conf 配置文件
    ├── const
    │   └── statusCode.go 常量定义
    ├── genPb.sh 根据proto文件生成pb文件脚本
    ├── protoFile
    │   ├── defaults.pb.go 执行genPb.sh脚本生成的文件
    │   └── defaults.proto 服务定义
    ├── rpc-srv 服务程序，startServe.sh编译生成
    ├── runtime
    │   ├── 2018-03-30
    │   │   └── index 服务运行日志
    │   ├── rpc-srv.out 异常终止日志
    │   └── rpc-srv.pid 当前服务pid
    ├── serve
    │   └── serve.go 服务端
    ├── services
    │   ├── BaseService.go 服务公共方法
    │   ├── frLog
    │   │   └── writeLog.go 写日志实现
    │   └── redisCache
    │       └── redisCache.go 缓存实现
    └── startServe.sh 起服务
```

###实现
定义proto文件
```
    syntax = "proto3";
    option java_package = "io.grpc.examples";
    package protoFile;
    //service 写在这里
    service Greeter {
      //写redis缓存
      rpc RedisCache (RedisCacheRequest) returns (RedisCacheReply) {}
    }
    //写redis缓存 传入参数
    message RedisCacheRequest {
        string name = 1;
        string value = 2;
        int64 express = 3;
    }
    //写redis缓存 返回信息
    message RedisCacheReply {
        int32 code = 4;
        string message = 5;
        map<string , string> data = 6;
    }
```
protoful类型映射参见：https://blog.csdn.net/superbfly/article/details/17920383
生成pb文件
```
    protoc --go_out=plugins=grpc:. protoFile/defaults.proto
```
需预先安装 protoc，见下文
服务端实现
```
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
```
起服务 go run serve/serve.go