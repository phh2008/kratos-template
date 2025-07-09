package biz

import (
	"context"
	"example.com/xxx/common-lib/model/page"
	"example.com/xxx/interface/internal/model"
)

type DemoRepo interface {
	ListRole(ctx context.Context, req model.RoleListReq) (*page.PageData[*model.RoleModel], error)
}

type DemoUseCase struct {
	demoRepo DemoRepo
}

func NewDemoUseCase(demoRepo DemoRepo) *DemoUseCase {
	return &DemoUseCase{demoRepo: demoRepo}
}

func (a *DemoUseCase) ListRole(ctx context.Context, req model.RoleListReq) (*page.PageData[*model.RoleModel], error) {
	return a.demoRepo.ListRole(ctx, req)
}
