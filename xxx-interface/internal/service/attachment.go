package service

import (
    "context"
    "example.com/xxx/interface/internal/biz"
    "mime/multipart"
)

type AttachmentService struct {
    attachmentUseCase *biz.AttachmentUseCase
}

func NewAttachmentService(attachmentUseCase *biz.AttachmentUseCase) *AttachmentService {
    return &AttachmentService{attachmentUseCase: attachmentUseCase}
}

type UploadReq struct {
    File       multipart.File
    FileHeader *multipart.FileHeader
}

type UploadResp struct {
}

func (a *AttachmentService) Upload(ctx context.Context, req *UploadReq) (*UploadResp, error) {
    // TODO 保存文件
    return nil, nil
}
