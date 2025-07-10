package data

import (
	"context"
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/jinzhu/copier"
)

type AccessKeyPO struct {
	ID        int64  `gorm:"column:id"`
	AccessID  string `gorm:"column:access_id"`
	AccessKey string `gorm:"column:access_key"`
	Remark    string `gorm:"column:remark"`
}

func (AccessKeyPO) TableName() string {
	return "sys_access_key"
}

var _ biz.AccessKeyRepo = (*accessKeyRepo)(nil)

type accessKeyRepo struct {
	*orm.BaseRepo
	data *Data
}

func NewAccessKeyRepo(data *Data, baseRepo *orm.BaseRepo) biz.AccessKeyRepo {
	return &accessKeyRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *accessKeyRepo) GetByAccessID(ctx context.Context, accessID string) (*model.AccessKey, error) {
	var entity AccessKeyPO
	err := a.GetDB(ctx).Model(AccessKeyPO{}).
		Where("access_id = ?", accessID).
		Limit(1).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	var reply model.AccessKey
	err = copier.Copy(&reply, entity)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
