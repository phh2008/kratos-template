package service

import (
    "context"
    pb "example.com/xxx/user-service/api/user/v1"
    "example.com/xxx/user-service/internal/biz"
    "example.com/xxx/user-service/internal/model"
    "github.com/jinzhu/copier"
)

type RoleService struct {
    pb.UnimplementedRoleServer
    roleUseCase *biz.RoleUseCase
}

func NewRoleService(roleUseCase *biz.RoleUseCase) *RoleService {
    return &RoleService{
        roleUseCase: roleUseCase,
    }
}

func (s *RoleService) ListPage(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListReply, error) {
    var request model.RoleListReq
    _ = copier.Copy(&request, &req)
    result, err := s.roleUseCase.ListPage(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.RoleListReply
    _ = copier.Copy(&reply, &result)
    _ = copier.Copy(&reply.RoleList, &result.Data)
    return &reply, nil
}

func (s *RoleService) Add(ctx context.Context, req *pb.RoleSaveRequest) (*pb.RoleReply, error) {
    var request model.RoleModel
    _ = copier.Copy(&request, &req)
    result, err := s.roleUseCase.Add(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.RoleReply
    _ = copier.Copy(&reply, result)
    return &reply, nil
}
func (s *RoleService) GetByCode(ctx context.Context, req *pb.RoleCodeRequest) (*pb.RoleReply, error) {
    result, err := s.roleUseCase.GetByCode(ctx, req.RoleCode)
    if err != nil {
        return nil, err
    }
    var reply pb.RoleReply
    _ = copier.Copy(&reply, result)
    return &reply, nil
}

func (s *RoleService) AssignPermission(ctx context.Context, req *pb.RoleAssignPermRequest) (*pb.RoleOk, error) {
    var request model.RoleAssignPermModel
    _ = copier.Copy(&request, &req)
    err := s.roleUseCase.AssignPermission(ctx, request)
    if err != nil {
        return nil, err
    }
    return &pb.RoleOk{Success: true}, nil
}

func (s *RoleService) DeleteById(ctx context.Context, req *pb.RoleDeleteRequest) (*pb.RoleOk, error) {
    err := s.roleUseCase.DeleteById(ctx, req.Id)
    if err != nil {
        return nil, err
    }
    return &pb.RoleOk{Success: true}, nil
}
