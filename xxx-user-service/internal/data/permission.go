package data

import (
    "context"
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/common-lib/orm"
    "example.com/xxx/user-service/internal/biz"
    "example.com/xxx/user-service/internal/model"
    "github.com/jinzhu/copier"
)

type PermissionEntity struct {
    orm.BaseEntity
    PermName string // 权限名称
    Url      string // URL路径
    Action   string // 权限动作：比如get、post、delete等
    PermType uint8  `gorm:"default:1"` // 权限类型：1-菜单、2-按钮
    ParentId int64  `gorm:"default:0"` // 父级ID：资源层级关系（0表示顶级）
}

func (PermissionEntity) TableName() string {
    return "sys_permission"
}

var _ biz.PermissionRepo = (*permissionRepo)(nil)

type permissionRepo struct {
    *orm.BaseRepo[PermissionEntity]
    data *Data
}

// NewPermissionRepo 创建dao
func NewPermissionRepo(data *Data) biz.PermissionRepo {
    return &permissionRepo{
        BaseRepo: orm.NewBaseRepo[PermissionEntity](data.db),
        data:     data,
    }
}

func (a *permissionRepo) ListPage(ctx context.Context, req model.PermissionListReq) (*page.PageData[*model.PermissionModel], error) {
    db := a.GetDB(ctx)
    db = db.Model(&PermissionEntity{})
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
    var entity PermissionEntity
    _ = copier.Copy(&entity, permission)
    err := a.Insert_(ctx, &entity)
    if err != nil {
        return nil, err
    }
    _ = copier.Copy(&permission, &entity)
    return &permission, nil
}

func (a *permissionRepo) Update(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error) {
    var entity PermissionEntity
    _ = copier.Copy(&entity, &permission)
    err := a.Update_(ctx, &entity)
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
    list, err := a.ListByIds_(ctx, idList)
    if err != nil {
        return nil, err
    }
    var result []model.PermissionModel
    _ = copier.Copy(&result, &list)
    return result, nil
}

func (a *permissionRepo) GetById(ctx context.Context, id int64) (*model.PermissionModel, error) {
    entity, err := a.GetByID_(ctx, id)
    if err != nil {
        return nil, err
    }
    var permission model.PermissionModel
    _ = copier.Copy(&permission, entity)
    return &permission, nil
}
