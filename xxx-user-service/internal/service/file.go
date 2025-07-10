package service

import (
	"context"
	"example.com/xxx/common-lib/types"
	pb "example.com/xxx/user-service/api/user/v1"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/jinzhu/copier"
	"log/slog"
)

type FileService struct {
	pb.UnimplementedFileServer
	fileUseCase *biz.FileUseCase
}

func NewFileService(fileUseCase *biz.FileUseCase) *FileService {
	return &FileService{fileUseCase: fileUseCase}
}

func (a *FileService) Add(ctx context.Context, req *pb.FileAddRequest) (*pb.FileAddReply, error) {
	var request model.FileAddReq
	err := copier.Copy(&request, req)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	if req.CreatedAt != nil {
		request.CreatedAt = types.LocalDateTime{Time: req.CreatedAt.AsTime()}
	}
	if req.UpdatedAt != nil {
		request.UpdatedAt = types.LocalDateTime{Time: req.UpdatedAt.AsTime()}
	}
	resp, err := a.fileUseCase.Add(ctx, &request)
	if err != nil {
		return nil, err
	}
	var reply pb.FileAddReply
	err = copier.Copy(&reply, resp)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	return &reply, nil
}
