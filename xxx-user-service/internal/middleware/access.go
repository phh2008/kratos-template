package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	v1 "example.com/xxx/user-service/api/user/v1"
	"example.com/xxx/user-service/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"log/slog"
)

func NewAccess(accessRepo biz.AccessKeyRepo) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, v1.ErrorSysError("get server context error")
			}
			xAccessID := tr.RequestHeader().Get("x-access-id")
			if xAccessID == "" {
				return nil, v1.ErrorUnauthorized("缺少请求头 x-access-id")
			}
			xTimestamp := tr.RequestHeader().Get("x-timestamp")
			if xTimestamp == "" {
				return nil, v1.ErrorUnauthorized("缺少请求头 x-timestamp")
			}
			xRequestID := tr.RequestHeader().Get("x-request-id")
			if xRequestID == "" {
				return nil, v1.ErrorUnauthorized("缺少请求头 x-request-id")
			}
			xSignature := tr.RequestHeader().Get("x-signature")
			if xSignature == "" {
				return nil, v1.ErrorUnauthorized("缺少请求头 x-signature")
			}
			access, err := accessRepo.GetByAccessID(context.TODO(), xAccessID)
			if err != nil {
				slog.Error("查询AccessKey出错", "error", err)
				return nil, v1.ErrorSysError("DB系统错误")
			}
			if access == nil || access.AccessID == "" {
				return nil, v1.ErrorUnauthorized("x-access-id 无效")
			}
			message := xAccessID + xRequestID + xTimestamp
			expectedSignature := generateSignature(access.AccessKey, message)
			if !validateSignature(xSignature, expectedSignature) {
				return nil, v1.ErrorUnauthorized("签名验证失败")
			}
			return handler(ctx, req)
		}
	}
}

func generateSignature(key, message string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func validateSignature(reqSignature, expectedSignature string) bool {
	return hmac.Equal([]byte(reqSignature), []byte(expectedSignature))
}
