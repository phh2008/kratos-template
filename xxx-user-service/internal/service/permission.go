package service

import (
    "context"
    pb "example.com/xxx/user-service/api/user/v1"
    "example.com/xxx/user-service/internal/biz"
    "example.com/xxx/user-service/internal/model"
    "github.com/jinzhu/copier"
)

type PermissionService struct {
    pb.UnimplementedPermissionServer
    permissionUseCase *biz.PermissionUseCase
}

func NewPermissionService(permissionUseCase *biz.PermissionUseCase) *PermissionService {
    return &PermissionService{
        permissionUseCase: permissionUseCase,
    }
}

func (s *PermissionService) ListPage(ctx context.Context, req *pb.PermListRequest) (*pb.PermListReply, error) {
    var request model.PermissionListReq
    _ = copier.Copy(&request, &req)
    result, err := s.permissionUseCase.ListPage(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.PermListReply
    _ = copier.Copy(&reply, &result)
    _ = copier.Copy(&reply.PermList, &result.Data)
    return &reply, nil
}

func (s *PermissionService) Add(ctx context.Context, req *pb.PermSaveRequest) (*pb.PermReply, error) {
    var request model.PermissionModel
    _ = copier.Copy(&request, &req)
    result, err := s.permissionUseCase.Add(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.PermReply
    _ = copier.Copy(&reply, result)
    return &reply, nil
}

func (s *PermissionService) Update(ctx context.Context, req *pb.PermSaveRequest) (*pb.PermReply, error) {
    var request model.PermissionModel
    _ = copier.Copy(&request, &req)
    result, err := s.permissionUseCase.Update(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.PermReply
    _ = copier.Copy(&reply, result)
    return &reply, nil
}
