package data

import (
	"context"
	"example.com/xxx/common-lib/consts"
	"example.com/xxx/common-lib/orm"
	"example.com/xxx/common-lib/types"
	"example.com/xxx/user-service/internal/biz"
	"example.com/xxx/user-service/internal/model"
	"github.com/jinzhu/copier"
	"log/slog"
	"time"
)

type FilePO struct {
	orm.BasePO
	FilePath     string `gorm:"column:file_path" json:"filePath"`         // 文件路径(保存后路径)
	FileSize     int64  `gorm:"column:file_size" json:"fileSize"`         // 文件大小(字节)
	FileMd5      string `gorm:"column:file_md5" json:"fileMd5"`           // 文件md5
	MediaType    string `gorm:"column:media_type" json:"mediaType"`       // 媒介类型：image/jpeg 等
	OriginalName string `gorm:"column:original_name" json:"originalName"` // 原始文件名
}

func (FilePO) TableName() string {
	return "sys_file"
}

var _ biz.FileRepo = (*fileRepo)(nil)

type fileRepo struct {
	*orm.BaseRepo
	data *Data
}

func NewFileRepo(data *Data, baseRepo *orm.BaseRepo) biz.FileRepo {
	return &fileRepo{
		BaseRepo: baseRepo,
		data:     data,
	}
}

func (a *fileRepo) Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error) {
	var file FilePO
	err := copier.Copy(&file, req)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	if file.CreatedAt.IsZero() {
		file.CreatedAt = types.LocalDateTime{Time: time.Now()}
	}
	if file.UpdatedAt.IsZero() {
		file.UpdatedAt = types.LocalDateTime{Time: time.Now()}
	}
	file.Deleted = consts.DeleteNot
	err = a.GetDB(ctx).Create(&file).Error
	if err != nil {
		slog.Error("insert error", "error", err)
		return nil, err
	}
	var resp model.FileAddResp
	err = copier.Copy(&resp, file)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	return &resp, nil
}
