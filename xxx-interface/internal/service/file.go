package service

import (
	"context"
	"example.com/xxx/common-lib/oss/storage"
	"example.com/xxx/interface/internal/biz"
	"example.com/xxx/interface/internal/model"
	"github.com/gabriel-vasile/mimetype"
	"github.com/rs/xid"
	"log/slog"
	"mime"
	"mime/multipart"
	"path/filepath"
)

type FileService struct {
	fileUseCase *biz.FileUseCase
	store       storage.Storage
}

func NewFileService(fileUseCase *biz.FileUseCase, store storage.Storage) *FileService {
	return &FileService{fileUseCase: fileUseCase, store: store}
}

type UploadReq struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	// 其它参数
}

func (a *FileService) Upload(ctx context.Context, req *UploadReq) (any, error) {
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
	// 调用文件服务保存文件信息
	resp, err := a.fileUseCase.Add(ctx, &model.FileAddReq{
		FilePath:     filepath.ToSlash(filePath),
		FileSize:     obj.Size,
		FileMd5:      "",
		MediaType:    mediaType,
		OriginalName: originalName,
	})
	if err != nil {
		slog.Error("fileUseCase.Add error", "error", err)
		return nil, err
	}
	return &resp, nil
}
