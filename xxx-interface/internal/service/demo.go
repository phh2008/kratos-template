package service

import (
    "context"
    "example.com/xxx/interface/internal/biz"
    "example.com/xxx/interface/internal/model"
    "github.com/jinzhu/copier"

    pb "example.com/xxx/interface/api/interface/v1"
)

type DemoService struct {
    pb.UnimplementedDemoServer
    demoUseCase *biz.DemoUseCase
}

func NewDemoService(demoUseCase *biz.DemoUseCase) *DemoService {
    return &DemoService{
        demoUseCase: demoUseCase,
    }
}

func (a *DemoService) ListPage(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListReply, error) {
    var request model.RoleListReq
    err := copier.Copy(&request, req)
    if err != nil {
        return nil, err
    }
    result, err := a.demoUseCase.ListRole(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.RoleListReply
    err = copier.Copy(&reply, result)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(&reply.RoleList, result.Data)
    if err != nil {
        return nil, err
    }
    return &reply, nil
}
