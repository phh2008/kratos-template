package biz

import (
	"context"
	"example.com/xxx/interface/internal/model"
)

type FileRepo interface {
	Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error)
}

type FileUseCase struct {
	fileRepo FileRepo
}

func NewFileUseCase(fileRepo FileRepo) *FileUseCase {
	return &FileUseCase{fileRepo: fileRepo}
}

func (a *FileUseCase) Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error) {
	return a.fileRepo.Add(ctx, req)
}
