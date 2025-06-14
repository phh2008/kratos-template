package service

import (
	"context"
	"example.com/xxx/common-lib/oss/storage"
	"example.com/xxx/interface/internal/biz"
	"github.com/jinzhu/copier"
	"mime/multipart"
	"time"
)

type AttachmentService struct {
	attachmentUseCase *biz.AttachmentUseCase
	store             storage.Storage
}

func NewAttachmentService(attachmentUseCase *biz.AttachmentUseCase, store storage.Storage) *AttachmentService {
	return &AttachmentService{attachmentUseCase: attachmentUseCase, store: store}
}

type UploadReq struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadResp struct {
	Path         string    `json:"path"`         // 文件路径包含文件名(不包括根目录    )
	Name         string    `json:"name"`         // 文件名
	LastModified time.Time `json:"lastModified"` // 最后修改时间
	Size         int64     `json:"size"`         // 文件大小
}

func (a *AttachmentService) Upload(ctx context.Context, req *UploadReq) (*UploadResp, error) {
	// 保存文件
	fileName := req.FileHeader.Filename
	file := req.File
	defer file.Close()
	obj, err := a.store.Put("test/"+fileName, file)
	if err != nil {
		return nil, err
	}
	var resp UploadResp
	_ = copier.Copy(&resp, obj)
	return &resp, nil
}
