package data

import (
	"context"
	"example.com/xxx/common-lib/model/page"
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/jinzhu/copier"
)

type PermissionPO struct {
	orm.BasePO
	PermName string // 权限名称
	Url      string // URL路径
	Action   string // 权限动作：比如get、post、delete等
	PermType uint8  `gorm:"default:1"` // 权限类型：1-菜单、2-按钮
	ParentId int64  `gorm:"default:0"` // 父级ID：资源层级关系（0表示顶级）
}

func (PermissionPO) TableName() string {
	return "sys_permission"
}

var _ biz.PermissionRepo = (*permissionRepo)(nil)

type permissionRepo struct {
	*orm.BaseRepo
	data *Data
}

// NewPermissionRepo 创建dao
func NewPermissionRepo(data *Data, baseRepo *orm.BaseRepo) biz.PermissionRepo {
	return &permissionRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *permissionRepo) ListPage(ctx context.Context, req model.PermissionListReq) (*page.PageData[*model.PermissionModel], error) {
	db := a.GetDB(ctx)
	db = db.Model(&PermissionPO{})
	if req.PermName != "" {
		db = db.Where("perm_name like ?", "%"+req.PermName+"%")
	}
	if req.Url != "" {
		db = db.Where("url=?", req.Url)
	}
	if req.Action != "" {
		db = db.Where("action=?", req.Action)
	}
	if req.PermType != 0 {
		db = db.Where("perm_type=?", req.PermType)
	}
	pageData, err := orm.QueryPage[*model.PermissionModel](db, req.GetPageNo(), req.GetPageSize())
	return pageData, err
}

func (a *permissionRepo) Add(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error) {
	var entity PermissionPO
	_ = copier.Copy(&entity, permission)
	err := a.GetDB(ctx).Create(&entity).Error
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&permission, &entity)
	return &permission, nil
}

func (a *permissionRepo) Update(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error) {
	var entity PermissionPO
	_ = copier.Copy(&entity, &permission)
	err := a.GetDB(ctx).Model(&entity).Updates(entity).Error
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&permission, &entity)
	return &permission, nil
}

func (a *permissionRepo) FindByIdList(ctx context.Context, idList []int64) ([]model.PermissionModel, error) {
	if len(idList) == 0 {
		return nil, nil
	}
	var list []PermissionPO
	err := a.GetDB(ctx).Find(&list, idList).Error
	if err != nil {
		return nil, err
	}
	var result []model.PermissionModel
	_ = copier.Copy(&result, &list)
	return result, nil
}

func (a *permissionRepo) GetById(ctx context.Context, id int64) (*model.PermissionModel, error) {
	var domain PermissionPO
	err := a.GetDB(ctx).Limit(1).Find(&domain, id).Error
	if err != nil {
		return nil, err
	}
	var permission model.PermissionModel
	_ = copier.Copy(&permission, domain)
	return &permission, nil
}
