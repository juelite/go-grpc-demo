# grpc微服务 写redis缓存 demo
### 依赖包
- golang.org/x/net/context
- google.golang.org/grpc
- github.com/Unknwon/goconfig //为了读取配置文件

### protoc 工具需先安装 
  ```https://github.com/google/protobuf/releases/download/v2.4.1/protobuf-2.4.1.tar.gz```
  
### 服务端 go run serve/serve.go

### 客户端 client/client.go
```
├── client
│   └── client.go //客户端
├── conf
│   └── setting.conf //配置文件
├── const
│   └── statusCode.go //常量定义
├── genPb.sh  //生成proto 脚本
├── protoFile
│   ├── defaults.pb.go //上面脚本生成
│   └── defaults.proto //gRPC服务定义
├── serve
│   └── serve.go //服务端
└── services
    ├── BaseService.go //公共方法定义
    └── redisCache
        ├── redis.go //redis 操作包
        └── redisCache.go //写redis缓存 logic
```
