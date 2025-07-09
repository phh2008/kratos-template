package data

import (
	"context"
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/jinzhu/copier"
)

type RolePermissionPO struct {
	Id     int64 // 主键id
	RoleId int64 // 角色id
	PermId int64 // 权限id
}

func (RolePermissionPO) TableName() string {
	return "sys_role_permission"
}

var _ biz.RolePermissionRepo = (*rolePermissionRepo)(nil)

type rolePermissionRepo struct {
	*orm.BaseRepo
	data *Data
}

// NewRolePermissionRepo 创建 repo
func NewRolePermissionRepo(data *Data, baseRepo *orm.BaseRepo) biz.RolePermissionRepo {
	return &rolePermissionRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *rolePermissionRepo) DeleteByRoleId(ctx context.Context, roleId int64) error {
	db := a.GetDB(ctx).Where("role_id=?", roleId).Delete(&RolePermissionPO{})
	return db.Error
}

func (a *rolePermissionRepo) BatchAdd(ctx context.Context, list []model.RolePermission) error {
	if len(list) == 0 {
		return nil
	}
	var listEntity []RolePermissionPO
	_ = copier.Copy(&listEntity, list)
	err := a.GetDB(ctx).Create(&listEntity).Error
	return err
}

func (a *rolePermissionRepo) ListRoleIdByPermId(ctx context.Context, permId int64) []int64 {
	var roleIds []int64
	a.GetDB(ctx).Model(&RolePermissionPO{}).
		Where("perm_id=?", permId).
		Pluck("role_id", &roleIds)
	return roleIds
}
