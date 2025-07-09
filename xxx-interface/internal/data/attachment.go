package data

import (
	"context"
	"example.com/xxx/interface/internal/biz"
	"example.com/xxx/interface/internal/model"
)

var _ biz.AttachmentRepo = (*attachmentRepo)(nil)

type attachmentRepo struct {
	data *Data
}

// NewAttachmentRepo 创建 dao
func NewAttachmentRepo(data *Data) biz.AttachmentRepo {
	return &attachmentRepo{
		data: data,
	}
}

func (a *attachmentRepo) Save(ctx context.Context, req model.AttachmentSaveReq) (*model.AttachmentSaveResp, error) {
	// 保存上传的附件信息
	return nil, nil
}
