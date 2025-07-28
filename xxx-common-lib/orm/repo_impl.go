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

var _ base.IBaseRepo = (*BaseRepo)(nil)

type dbTxKey struct{}

type BaseRepo struct {
	database *gorm.DB
}

// NewBaseRepo 创建 baseRepo
func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{database: db}
}

// Transaction 开启事务
func (a *BaseRepo) Transaction(c context.Context, handler func(tx context.Context) error) error {
	db := a.database
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return handler(context.WithValue(c, dbTxKey{}, tx))
	})
}

// GetDB 获取事务的db连接
func (a *BaseRepo) GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbTxKey{}).(*gorm.DB)
	if !ok {
		db = a.database
		return db.WithContext(ctx)
	}
	return db
}

type BasePO struct {
	ID        int64                 `gorm:"column:id;primaryKey" json:"id"`                                                   // 主键ID
	CreatedAt types.LocalDateTime   `gorm:"column:created_at;autoCreateTime" json:"createdAt"`                                // 创建时间
	UpdatedAt types.LocalDateTime   `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`                                // 更新时间
	CreatedBy string                `gorm:"column:created_by" json:"createdBy"`                                               // 创建人
	UpdatedBy string                `gorm:"column:updated_by" json:"updatedBy"`                                               // 更新人
	Deleted   soft_delete.DeletedAt `gorm:"column:deleted;softDelete:flag,DeletedAtField:UpdatedAt;default:1" json:"deleted"` // 是否删除 1-否，2-是
}
