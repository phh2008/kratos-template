package orm

import (
    "context"
    "example.com/xxx/common-lib/base"
    "example.com/xxx/common-lib/types"
    "gorm.io/gorm"
    "gorm.io/plugin/soft_delete"
)

func init() {
    soft_delete.FlagDeleted = 2
    soft_delete.FlagActived = 1
}

var _ base.IBaseRepo = (*BaseRepo[any])(nil)

type dbTxKey struct{}

type BaseRepo[T any] struct {
    database *gorm.DB
}

// NewBaseRepo 创建 baseRepo
func NewBaseRepo[T any](db *gorm.DB) *BaseRepo[T] {
    return &BaseRepo[T]{database: db}
}

// Transaction 开启事务
func (a *BaseRepo[T]) Transaction(c context.Context, handler func(tx context.Context) error) error {
    db := a.database
    return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
        return handler(context.WithValue(c, dbTxKey{}, tx))
    })
}

// GetDB 获取事务的db连接
func (a *BaseRepo[T]) GetDB(ctx context.Context) *gorm.DB {
    db, ok := ctx.Value(dbTxKey{}).(*gorm.DB)
    if !ok {
        db = a.database
        return db.WithContext(ctx)
    }
    return db
}

// GetByID_ 根据ID查询
func (a *BaseRepo[T]) GetByID_(ctx context.Context, id int64) (*T, error) {
    var domain T
    err := a.GetDB(ctx).Limit(1).Find(&domain, id).Error
    return &domain, err
}

// Insert_ 新增
func (a *BaseRepo[T]) Insert_(ctx context.Context, entity *T) error {
    return a.GetDB(ctx).Create(entity).Error
}

// Update_ 更新
func (a *BaseRepo[T]) Update_(ctx context.Context, entity *T) error {
    return a.GetDB(ctx).Model(entity).Updates(*entity).Error
}

// DeleteByID_ 根据ID删除
func (a *BaseRepo[T]) DeleteByID_(ctx context.Context, id int64) error {
    return a.GetDB(ctx).Delete(new(T), id).Error
}

// ListByIds_ 根据ID集合查询
func (a *BaseRepo[T]) ListByIds_(ctx context.Context, ids []int64) ([]T, error) {
    var list []T
    db := a.GetDB(ctx).Find(&list, ids)
    return list, db.Error
}

type BaseEntity struct {
    ID       int64                 `gorm:"column:id;primaryKey" json:"id"`                                                  // 主键ID
    CreateAt types.LocalTime       `gorm:"column:create_at;autoCreateTime" json:"createAt"`                                 // 创建时间
    UpdateAt types.LocalTime       `gorm:"column:update_at;autoUpdateTime" json:"updateAt"`                                 // 更新时间
    CreateBy int64                 `gorm:"column:create_by" json:"createBy"`                                                // 创建人
    UpdateBy int64                 `gorm:"column:update_by" json:"updateBy"`                                                // 更新人
    Deleted  soft_delete.DeletedAt `gorm:"column:deleted;softDelete:flag,DeletedAtField:UpdateAt,default:1" json:"deleted"` // 是否删除 1-否，2-是
}
