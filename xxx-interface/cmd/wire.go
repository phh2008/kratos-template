//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "example.com/xxx/interface/internal/biz"
    "example.com/xxx/interface/internal/conf"
    "example.com/xxx/interface/internal/data"
    "example.com/xxx/interface/internal/server"
    "example.com/xxx/interface/internal/service"
    "github.com/go-kratos/kratos/v2"
    "github.com/google/wire"
    "go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *zap.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(server.ProviderSet,
        data.ProviderSet,
        biz.ProviderSet,
        service.ProviderSet,
        newApp,
    ))
}
