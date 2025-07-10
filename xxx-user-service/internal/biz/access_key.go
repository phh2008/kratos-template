package biz

import (
	"context"
	"example.com/xxx/common-lib/base"
	"example.com/xxx/user-service/internal/model"
)

type AccessKeyRepo interface {
	base.IBaseRepo
	GetByAccessID(ctx context.Context, accessID string) (*model.AccessKey, error)
}
