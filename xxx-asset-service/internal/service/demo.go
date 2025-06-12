package service

import (
    "context"
    "example.com/xxx/asset-service/internal/biz"

    pb "example.com/xxx/asset-service/api/asset/v1"
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

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
    return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
    return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
    return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
    return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
    return &pb.ListDemoReply{}, nil
}
