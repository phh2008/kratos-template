package main

import (
    "example.com/xxx/asset-service/internal/conf"
    "example.com/xxx/common-lib/logger"
    "flag"
    "github.com/go-kratos/kratos/v2/registry"
    "go.uber.org/zap"
    "go.uber.org/zap/exp/zapslog"
    "log/slog"
    "os"

    kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    _ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
    // Name is the name of the compiled software.
    Name = "asset.service"
    // Version is the version of the compiled software.
    Version = "v1.0.0"
    // flagconf is the config flag.
    flagconf string

    id, _ = os.Hostname()

    active string
)

func init() {
    flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf ./config")
    flag.StringVar(&active, "active", "dev", "active environment, eg: -active dev")
}

func newApp(gs *grpc.Server, zapLog *zap.Logger, rr registry.Registrar) *kratos.App {
    // 包装 zap logger
    zlog := kratoszap.NewLogger(zapLog.WithOptions(zap.AddCallerSkip(3)))
    wrapLog := log.With(zlog,
        "service.name", Name,
        "service.version", Version,
        "ts", log.DefaultTimestamp,
    )
    // slog 全局日志
    sl := slog.New(zapslog.NewHandler(zapLog.Core(), zapslog.WithCaller(true)))
    sl = sl.With(
        slog.String("service.name", Name),
        slog.String("service.version", Version),
    )
    slog.SetDefault(sl)
    // kratos 全局日志
    log.SetLogger(wrapLog)
    return kratos.New(
        kratos.ID(id+"_"+Name),
        kratos.Name(Name),
        kratos.Version(Version),
        kratos.Metadata(map[string]string{}),
        kratos.Logger(wrapLog),
        kratos.Server(
            gs,
        ),
        kratos.Registrar(rr),
    )
}

func main() {
    flag.Parse()
    // 加载配置
    conf.Active = active
    var bc = conf.NewConfig(flagconf)
    // 初始化日志
    zapLog := logger.NewLogger(&logger.Config{
        Level:      bc.Log.Level,
        Filename:   bc.Log.Filename,
        MaxSize:    bc.Log.MaxSize,
        MaxBackups: bc.Log.MaxBackups,
        MaxAge:     bc.Log.MaxAge,
        Compress:   bc.Log.Compress,
        LocalTime:  bc.Log.LocalTime,
    })
    // 依赖注入
    app, cleanup, err := wireApp(bc, zapLog)
    if err != nil {
        panic(err)
    }
    defer cleanup()

    // start and wait for stop signal
    if err := app.Run(); err != nil {
        panic(err)
    }
}
