package biz

import (
    "context"
    "example.com/xxx/common-lib/base"
    "example.com/xxx/user-service/internal/model"
)

type RolePermissionRepo interface {
    base.IBaseRepo
    DeleteByRoleId(ctx context.Context, roleId int64) error
    BatchAdd(ctx context.Context, list []model.RolePermission) error
    ListRoleIdByPermId(ctx context.Context, permId int64) []int64
}

type RolePermissionUseCase struct {
    rolePermissionRepo RolePermissionRepo
}

func NewRolePermissionUseCase(rolePermissionRepo RolePermissionRepo) *RolePermissionUseCase {
    return &RolePermissionUseCase{
        rolePermissionRepo: rolePermissionRepo,
    }
}
