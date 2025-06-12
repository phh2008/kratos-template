package biz

import (
    "context"
    "example.com/xxx/common-lib/base"
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/user-service/internal/model"
    "github.com/casbin/casbin/v2"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/jinzhu/copier"
    "log/slog"
    "net/http"
)

type PermissionRepo interface {
    base.IBaseRepo
    ListPage(ctx context.Context, req model.PermissionListReq) (*page.PageData[*model.PermissionModel], error)
    Add(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error)
    Update(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error)
    FindByIdList(ctx context.Context, idList []int64) ([]model.PermissionModel, error)
    GetById(ctx context.Context, id int64) (*model.PermissionModel, error)
}

type PermissionUseCase struct {
    permissionRepo PermissionRepo
    enforcer       *casbin.Enforcer
}

func NewPermissionUseCase(permissionRepo PermissionRepo, enforcer *casbin.Enforcer) *PermissionUseCase {
    return &PermissionUseCase{
        permissionRepo: permissionRepo,
        enforcer:       enforcer,
    }
}

// ListPage 权限资源列表
func (a *PermissionUseCase) ListPage(ctx context.Context, req model.PermissionListReq) (*page.PageData[*model.PermissionModel], error) {
    return a.permissionRepo.ListPage(ctx, req)
}

// Add 添加权限资源
func (a *PermissionUseCase) Add(ctx context.Context, perm model.PermissionModel) (*model.PermissionModel, error) {
    var permission model.PermissionModel
    err := copier.Copy(&permission, &perm)
    if err != nil {
        return nil, err
    }
    return a.permissionRepo.Add(ctx, permission)
}

// Update 更新权限资源
func (a *PermissionUseCase) Update(ctx context.Context, perm model.PermissionModel) (*model.PermissionModel, error) {
    oldPerm, err := a.permissionRepo.GetById(ctx, perm.Id)
    if err != nil {
        return nil, err
    }
    if oldPerm.Id == 0 {
        return nil, errors.New(http.StatusNotFound, "NOT_FOUND", "权限不存在")
    }
    var permission model.PermissionModel
    err = copier.Copy(&permission, &perm)
    if err != nil {
        return nil, err
    }
    // 更新权限资源表
    res, err := a.permissionRepo.Update(ctx, permission)
    if err != nil {
        return nil, err
    }
    // 获取角色与资源列表,比如：[[systemAdmin /api/v1/user/list get] [guest /api/v1/user/list get]]
    perms, _ := a.enforcer.GetFilteredPolicy(1, oldPerm.Url, oldPerm.Action)
    // 更新casbin中的数据
    if len(perms) > 0 {
        for i, v := range perms {
            item := v
            item[1] = res.Url
            item[2] = res.Action
            perms[i] = item
        }
        _, err = a.enforcer.UpdateFilteredPolicies(perms, 1, oldPerm.Url, oldPerm.Action)
        if err != nil {
            slog.Error("更新casbin中的权限错误", "error", err)
        }
    }
    return res, nil
}
