package data

import (
    "context"
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/common-lib/orm"
    "example.com/xxx/interface/internal/biz"
    "example.com/xxx/interface/internal/model"
    userv1 "example.com/xxx/user-service/api/user/v1"
    "github.com/jinzhu/copier"
)

var _ biz.DemoRepo = (*demoRepo)(nil)

type demoRepo struct {
    *orm.BaseRepo[any]
    data *Data
}

// NewDemoRepo 创建 dao
func NewDemoRepo(data *Data) biz.DemoRepo {
    return &demoRepo{
        BaseRepo: orm.NewBaseRepo[any](nil),
        data:     data,
    }
}

func (a *demoRepo) ListRole(ctx context.Context, req model.RoleListReq) (*page.PageData[*model.RoleModel], error) {
    var request userv1.RoleListRequest
    err := copier.Copy(&request, &req)
    if err != nil {
        return nil, err
    }
    resp, err := a.data.rc.ListPage(ctx, &request)
    if err != nil {
        return nil, err
    }
    var pageData page.PageData[*model.RoleModel]
    err = copier.Copy(&pageData, resp)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&pageData.Data, resp.RoleList)
    if err != nil {
        return nil, err
    }
    return &pageData, nil
}
