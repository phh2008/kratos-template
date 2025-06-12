package server

import (
    "context"
    "example.com/xxx/interface/api/interface/v1"
    "example.com/xxx/interface/internal/conf"
    "example.com/xxx/interface/internal/middleware"
    "example.com/xxx/interface/internal/service"
    userv1 "example.com/xxx/user-service/api/user/v1"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/middleware/selector"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "log/slog"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
    cf *conf.Bootstrap,
    uc userv1.UserClient,
    demoService *service.DemoService,
) *grpc.Server {
    c := cf.Server
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            recovery.Recovery(),
            selector.Server(
                middleware.NewAuthenticate(uc),
                middleware.NewAuthorization(uc),
            ).Match(func(ctx context.Context, operation string) bool {
                slog.Info("operation: " + operation)
                if middleware.NoneAuthOperation.Contains(operation) {
                    return false
                }
                return true
            }).Build(),
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
