package data

import (
	"context"
	"example.com/xxx/common-lib/model/page"
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
)

type RolePO struct {
	orm.BasePO
	RoleCode string // 角色编码
	RoleName string // 角色名称
}

func (RolePO) TableName() string {
	return "sys_role"
}

var _ biz.RoleRepo = (*roleRepo)(nil)

type roleRepo struct {
	*orm.BaseRepo
	data *Data
}

// NewRoleRepo 创建 dao
func NewRoleRepo(data *Data, baseRepo *orm.BaseRepo) biz.RoleRepo {
	return &roleRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *roleRepo) ListPage(ctx context.Context, req model.RoleListReq) (*page.PageData[*model.RoleModel], error) {
	db := a.GetDB(ctx)
	db = db.Model(&RolePO{})
	if req.RoleCode != "" {
		db = db.Where("role_code like ?", "%"+req.RoleCode+"%")
	}
	if req.RoleName != "" {
		db = db.Where("role_name like ?", "%"+req.RoleName+"%")
	}
	pageData, err := orm.QueryPage[*model.RoleModel](db, req.GetPageNo(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	return pageData, nil
}

// Add 添加角色
func (a *roleRepo) Add(ctx context.Context, req model.RoleModel) (*model.RoleModel, error) {
	// 检查角色是否存在
	role, err := a.GetByCode(ctx, req.RoleCode)
	if err != nil {
		return nil, err
	}
	if role.Id > 0 {
		return nil, errors.New(500, "role_exist", "角色已存在")
	}
	var entity RolePO
	_ = copier.Copy(&entity, req)
	err = a.GetDB(ctx).Create(&entity).Error
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&req, &entity)
	return &req, nil
}

func (a *roleRepo) GetById(ctx context.Context, id int64) (*model.RoleModel, error) {
	var domain RolePO
	err := a.GetDB(ctx).Limit(1).Find(&domain, id).Error
	if err != nil {
		return nil, err
	}
	var role model.RoleModel
	_ = copier.Copy(&role, domain)
	return &role, nil
}

// GetByCode 根据角色编号获取角色
func (a *roleRepo) GetByCode(ctx context.Context, code string) (*model.RoleModel, error) {
	var entity RolePO
	err := a.GetDB(ctx).Where("role_code=?", code).Limit(1).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	var role model.RoleModel
	_ = copier.Copy(&role, &entity)
	return &role, nil
}

// DeleteById 删除角色
func (a *roleRepo) DeleteById(ctx context.Context, id int64) error {
	return a.GetDB(ctx).Delete(&RolePO{}, id).Error
}

// ListByIds 根据角色ID集合查询角色列表
func (a *roleRepo) ListByIds(ctx context.Context, ids []int64) ([]model.RoleModel, error) {
	var list []RolePO
	err := a.GetDB(ctx).Find(&list, ids).Error
	if err != nil {
		return nil, err
	}
	var roleList []model.RoleModel
	_ = copier.Copy(&roleList, list)
	return roleList, nil
}
