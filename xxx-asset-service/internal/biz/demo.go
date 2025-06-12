package biz

import (
    "context"
    "example.com/xxx/asset-service/internal/model"
    "example.com/xxx/common-lib/base"
)

type DemoRepo interface {
    base.IBaseRepo
    GetByID(ctx context.Context, id int64) (*model.Demo, error)
}

type DemoUseCase struct {
    demoRepo DemoRepo
}

func NewDemoUseCase(
    demoRepo DemoRepo,
) *DemoUseCase {
    return &DemoUseCase{
        demoRepo: demoRepo,
    }
}

// GetByID 根据ID获取
func (a *DemoUseCase) GetByID(ctx context.Context, id int64) (*model.Demo, error) {
    return a.demoRepo.GetByID(ctx, id)
}
