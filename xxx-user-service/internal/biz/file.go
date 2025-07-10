package biz

import (
	"context"
	"example.com/xxx/common-lib/base"
	"example.com/xxx/user-service/internal/model"
)

type FileRepo interface {
	base.IBaseRepo
	// Add 添加文件信息
	Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error)
}

type FileUseCase struct {
	fileRepo FileRepo
}

func NewFileUseCase(
	fileRepo FileRepo,
) *FileUseCase {
	return &FileUseCase{
		fileRepo: fileRepo,
	}
}

func (a *FileUseCase) Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error) {
	return a.fileRepo.Add(ctx, req)
}
