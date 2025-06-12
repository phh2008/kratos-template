package service

import (
    "context"
    pb "example.com/xxx/user-service/api/user/v1"
    "example.com/xxx/user-service/internal/biz"
    "example.com/xxx/user-service/internal/model"
    "github.com/jinzhu/copier"
)

type UserService struct {
    pb.UnimplementedUserServer
    userUseCase *biz.UserUseCase
}

func NewUserService(userUseCase *biz.UserUseCase) *UserService {
    return &UserService{
        userUseCase: userUseCase,
    }
}

func (s *UserService) ListPage(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
    var request model.UserListReq
    _ = copier.Copy(&request, &req)
    result, err := s.userUseCase.ListPage(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.UserListReply
    _ = copier.Copy(&reply, &result)
    _ = copier.Copy(&reply.UserList, &result.Data)
    return &reply, nil
}

func (s *UserService) CreateByEmail(ctx context.Context, req *pb.UserEmailRequest) (*pb.UserReply, error) {
    var request model.UserEmailRegister
    _ = copier.Copy(&request, &req)
    result, err := s.userUseCase.CreateByEmail(ctx, request)
    if err != nil {
        return nil, err
    }
    var reply pb.UserReply
    _ = copier.Copy(&reply, &result)
    return &reply, nil
}

func (s *UserService) LoginByEmail(ctx context.Context, req *pb.UserEmailLoginRequest) (*pb.LoginReply, error) {
    var request model.UserLoginModel
    _ = copier.Copy(&request, &req)
    result, err := s.userUseCase.LoginByEmail(ctx, request)
    if err != nil {
        return nil, err
    }
    return &pb.LoginReply{Token: result}, nil
}

func (s *UserService) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.UserOk, error) {
    var request model.AssignRoleModel
    _ = copier.Copy(&request, &req)
    err := s.userUseCase.AssignRole(ctx, request)
    if err != nil {
        return nil, err
    }
    return &pb.UserOk{Success: true}, nil
}

func (s *UserService) DeleteById(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserOk, error) {
    err := s.userUseCase.DeleteById(ctx, req.Id)
    if err != nil {
        return nil, err
    }
    return &pb.UserOk{Success: true}, nil
}

func (s *UserService) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.UserOk, error) {
    err := s.userUseCase.CheckPermission(ctx, model.CheckPermReq{
        Sub: req.Sub,
        Obj: req.Obj,
        Act: req.Act,
    })
    if err != nil {
        return nil, err
    }
    return &pb.UserOk{Success: true}, nil
}

func (s *UserService) VerifyToken(ctx context.Context, req *pb.UserVerifyTokenRequest) (*pb.UserClaimsReply, error) {
    claims, err := s.userUseCase.VerifyToken(ctx, req.Token)
    if err != nil {
        return nil, err
    }
    return &pb.UserClaimsReply{
        Id:   claims.ID,
        Sub:  claims.Subject,
        Role: claims.Role,
    }, nil
}
