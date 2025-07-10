package service

import (
	"context"
	"example.com/xxx/common-lib/oss/storage"
	"example.com/xxx/interface/internal/biz"
	"github.com/gabriel-vasile/mimetype"
	"github.com/jinzhu/copier"
	"github.com/rs/xid"
	"mime"
	"mime/multipart"
	"path/filepath"
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
	MimeType     string    `json:"mimeType"`     // 文件类型
}

func (a *AttachmentService) Upload(ctx context.Context, req *UploadReq) (*UploadResp, error) {
	// 原始文件名称
	originalName := req.FileHeader.Filename
	file := req.File
	defer file.Close()
	fileName := xid.New().String() + filepath.Ext(originalName)
	// 保存文件的路径
	filePath := filepath.Join("compatTest", fileName)
	obj, err := a.store.Put(filePath, file)
	if err != nil {
		return nil, err
	}
	// 获取文件类型
	var mediaType string
	in, err := a.store.GetStream(filePath)
	if err == nil {
		defer in.Close()
		mtype, err := mimetype.DetectReader(in)
		if err == nil {
			mediaType = mtype.String()
		}
	}
	if mediaType == "" {
		mediaType = mime.TypeByExtension(fileName)
	}
	var resp UploadResp
	_ = copier.Copy(&resp, obj)
	resp.MimeType = mediaType
	return &resp, nil
}
