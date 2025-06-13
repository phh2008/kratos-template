package biz

import (
    "context"
    "example.com/xxx/common-lib/base"
    "example.com/xxx/interface/internal/model"
)

type AttachmentRepo interface {
    base.IBaseRepo // 附件
    Save(ctx context.Context, req model.AttachmentSaveReq) (*model.AttachmentSaveResp, error)
}

type AttachmentUseCase struct {
    attachmentRepo AttachmentRepo
}

func NewAttachmentUseCase(attachmentRepo AttachmentRepo) *AttachmentUseCase {
    return &AttachmentUseCase{attachmentRepo: attachmentRepo}
}
