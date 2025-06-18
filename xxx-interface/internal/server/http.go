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
    "github.com/go-kratos/kratos/v2/middleware/validate"
    "github.com/go-kratos/kratos/v2/transport/http"
    "github.com/gorilla/handlers"
    "log/slog"
    "strings"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
    cf *conf.Bootstrap,
    uc userv1.UserClient,
    demoService *service.DemoService,
    attachmentService *service.AttachmentService,
) *http.Server {
    c := cf.Server
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
            validate.Validator(),
            selector.Server(
                middleware.NewAuthenticate(uc),
                middleware.NewAuthorization(uc),
            ).Match(func(ctx context.Context, operation string) bool {
                slog.Info("operation: " + operation)
                if middleware.NoneAuthOperation.Contains(operation) {
                    return false
                }
                // TODO 测试，放行所有API修改返回为false
                return false
            }).Build(),
        ),
    }
    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }
    corsCfg := cf.Cors
    // 跨域配置
    corsOption := []handlers.CORSOption{
        handlers.AllowedOrigins(corsCfg.AllowedOriginPatterns),
        handlers.AllowedMethods(strings.Split(corsCfg.AllowedMethods, ",")),
        handlers.AllowedHeaders(strings.Split(corsCfg.AllowedHeaders, ",")),
        handlers.ExposedHeaders(strings.Split(corsCfg.ExposeHeaders, ",")),
        handlers.MaxAge(int(corsCfg.MaxAge)),
    }
    if corsCfg.AllowCredentials {
        corsOption = append(corsOption, handlers.AllowCredentials())
    }
    opts = append(opts, http.Filter(handlers.CORS(corsOption...)))
    opts = append(opts, http.PathPrefix("/api"))
    srv := http.NewServer(opts...)
    v1.RegisterDemoHTTPServer(srv, demoService)

    // 上传文件处理
    route := srv.Route("/")
    route.POST("/v1/upload", uploadHandler(attachmentService))
    return srv
}

func uploadHandler(srv *service.AttachmentService) func(ctx http.Context) error {
    return func(ctx http.Context) error {
        request := ctx.Request()
        f, header, err := request.FormFile("file")
        if err != nil {
            return err
        }
        in := service.UploadReq{
            File:       f,
            FileHeader: header,
        }
        http.SetOperation(ctx, "/v1.upload")
        h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
            return srv.Upload(ctx, req.(*service.UploadReq))
        })
        out, err := h(ctx, &in)
        if err != nil {
            return err
        }
        reply := out.(*service.UploadResp)
        return ctx.Result(200, reply)
    }
}
