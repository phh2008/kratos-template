package server

import (
    "example.com/xxx/asset-service/api/asset/v1"
    "example.com/xxx/asset-service/internal/conf"
    "example.com/xxx/asset-service/internal/service"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
    cf *conf.Bootstrap,
    demoService *service.DemoService,
) *grpc.Server {
    c := cf.Server
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            recovery.Recovery(),
        ),
    }
    if c.Grpc.Network != "" {
        opts = append(opts, grpc.Network(c.Grpc.Network))
    }
    if c.Grpc.Addr != "" {
        opts = append(opts, grpc.Address(c.Grpc.Addr))
    }
    if c.Grpc.Timeout != nil {
        opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
    }
    srv := grpc.NewServer(opts...)
    v1.RegisterDemoServer(srv, demoService)
    return srv
}
