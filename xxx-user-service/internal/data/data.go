package data

import (
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/user-service/internal/conf"
	xcasbin "example.com/xxx/user-service/internal/pkg/casbin"
	xjwt "example.com/xxx/user-service/internal/pkg/jwt"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"log/slog"
	
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDb,
	NewData,
	orm.NewBaseRepo,
	xcasbin.NewCasbin,
	xjwt.NewJwtHelper,
	NewDiscovery,
	NewRegistrar,
	NewUserRepo,
	NewRolePermissionRepo,
	NewRoleRepo,
	NewPermissionRepo,
)

// Data .
type Data struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
	jwt      *xjwt.JwtHelper
}

func NewDb(c *conf.Bootstrap) *gorm.DB {
	return orm.NewDB(c.Data.Database.Source)
}

// NewData .
func NewData(c *conf.Bootstrap, db *gorm.DB, enforcer *casbin.Enforcer, jwt *xjwt.JwtHelper) (*Data, func(), error) {
	cleanup := func() {
		slog.Info("closing the data resources")
	}
	return &Data{db: db, enforcer: enforcer, jwt: jwt}, cleanup, nil
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
