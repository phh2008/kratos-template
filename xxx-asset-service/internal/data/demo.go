package data

import (
	"context"
	"example.com/xxx/asset-service/internal/biz"
	"example.com/xxx/asset-service/internal/model"
	"example.com/xxx/common-lib/orm"
	"github.com/jinzhu/copier"
)

type DemoPO struct {
	orm.BasePO
	// other field ...
}

var _ biz.DemoRepo = (*demoRepo)(nil)

type demoRepo struct {
	*orm.BaseRepo
	data *Data
}

// NewDemoRepo 创建 dao
func NewDemoRepo(data *Data, baseRepo *orm.BaseRepo) biz.DemoRepo {
	return &demoRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *demoRepo) GetByID(ctx context.Context, id int64) (*model.Demo, error) {
	// TODO 调用 db 获取数据
	var domain DemoPO
	err := a.GetDB(ctx).Limit(1).Find(&domain, id).Error
	if err != nil {
		return nil, err
	}
	var resp model.Demo
	err = copier.Copy(&resp, domain)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
