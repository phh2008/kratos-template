package data

import (
	"context"
	"example.com/xxx/common-lib/oss"
	"example.com/xxx/common-lib/oss/storage"
	"example.com/xxx/interface/internal/conf"
	userv1 "example.com/xxx/user-service/api/user/v1"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"log/slog"
	
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewUserServiceClient,
	NewRoleServiceClient,
	NewFileServiceClient,
	NewFileRepo,
	NewDemoRepo,
	NewStorage,
)

// Data .
type Data struct {
	uc userv1.UserClient
	rc userv1.RoleClient
	fc userv1.FileClient
	sc storage.Storage
}

// NewData .
func NewData(
	c *conf.Bootstrap,
	uc userv1.UserClient,
	rc userv1.RoleClient,
	fc userv1.FileClient,
) (*Data, func(), error) {
	cleanup := func() {
		slog.Info("closing the data resources")
	}
	return &Data{uc: uc, rc: rc, fc: fc}, cleanup, nil
}

func NewStorage() storage.Storage {
	// 本地文件存储
	return oss.NewStorage(&storage.Config{OssType: storage.TypeFileSystem, BaseFolder: "/data/files"})
}

func NewDiscovery(conf *conf.Bootstrap) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = conf.Registry.Consul.Address
	c.Scheme = conf.Registry.Consul.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Bootstrap) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = conf.Registry.Consul.Address
	c.Scheme = conf.Registry.Consul.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

func NewRoleServiceClient(r registry.Discovery) userv1.RoleClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewRoleClient(conn)
	return c
}

func NewFileServiceClient(r registry.Discovery) userv1.FileClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewFileClient(conn)
	return c
}
