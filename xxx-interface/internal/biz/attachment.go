package biz

import (
	"context"
	"example.com/xxx/interface/internal/model"
)

type AttachmentRepo interface {
	Save(ctx context.Context, req model.AttachmentSaveReq) (*model.AttachmentSaveResp, error)
}

type AttachmentUseCase struct {
	attachmentRepo AttachmentRepo
}

func NewAttachmentUseCase(attachmentRepo AttachmentRepo) *AttachmentUseCase {
	return &AttachmentUseCase{attachmentRepo: attachmentRepo}
}
