package middleware

import (
    "context"
    "example.com/xxx/common-lib/consts"
    userv1 "example.com/xxx/user-service/api/user/v1"
    mapset "github.com/deckarep/golang-set/v2"
    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
    "log/slog"
    "path/filepath"
    "strings"
)

var NoneAuthOperation = mapset.NewSet[string](
    // 无需登录览权限的API
)

// NewAuthenticate 认证校验
func NewAuthenticate(uc userv1.UserClient) middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            tr, ok := transport.FromServerContext(ctx)
            if !ok {
                return nil, userv1.ErrorNoLogin("未登录")
            }
            token := tr.RequestHeader().Get(consts.AuthTokenHeaderKey)
            if token == "" {
                return nil, userv1.ErrorNoLogin("未登录")
            }
            result, err := uc.VerifyToken(ctx, &userv1.UserVerifyTokenRequest{
                Token: token,
            })
            if err != nil {
                return nil, err
            }
            ctx = context.WithValue(ctx, consts.UserCtxKey{}, result)
            return handler(ctx, req)
        }
    }
}

// NewAuthorization 权限校验
func NewAuthorization(uc userv1.UserClient) middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            user, ok := ctx.Value(consts.UserCtxKey{}).(userv1.UserClaimsReply)
            if !ok {
                return nil, userv1.ErrorUnauthorized("无权限")
            }
            tr, ok := transport.FromServerContext(ctx)
            if !ok {
                return nil, userv1.ErrorUnauthorized("无权限")
            }
            // 获取接口包名与方法名称
            dir, file := filepath.Split(tr.Operation())
            obj := strings.TrimSuffix(dir, "/")
            act := file
            slog.Info("权限校验", "sub", user.Sub, "obj", obj, "act", act)
            _, err = uc.CheckPermission(ctx, &userv1.CheckPermissionRequest{
                Sub: user.Sub,
                Obj: obj,
                Act: act,
            })
            if err != nil {
                // 鉴权出错了
                return nil, err
            }
            return handler(ctx, req)
        }
    }
}
