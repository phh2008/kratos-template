package data

import (
	"context"
	"example.com/xxx/interface/internal/biz"
	"example.com/xxx/interface/internal/model"
	userv1 "example.com/xxx/user-service/api/user/v1"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

var _ biz.FileRepo = (*fileRepo)(nil)

type fileRepo struct {
	data *Data
}

// NewFileRepo 创建 dao
func NewFileRepo(data *Data) biz.FileRepo {
	return &fileRepo{
		data: data,
	}
}

func (a *fileRepo) Add(ctx context.Context, req *model.FileAddReq) (*model.FileAddResp, error) {
	var in userv1.FileAddRequest
	err := copier.Copy(&in, req)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	if !req.CreatedAt.IsZero() {
		in.CreatedAt = timestamppb.New(req.CreatedAt.Time)
	}
	if !req.UpdatedAt.IsZero() {
		in.UpdatedAt = timestamppb.New(req.UpdatedAt.Time)
	}
	out, err := a.data.fc.Add(ctx, &in)
	if err != nil {
		slog.Error("fileClient error", "error", err)
		return nil, err
	}
	var resp model.FileAddResp
	err = copier.Copy(&resp, out)
	if err != nil {
		slog.Error("copier error", "error", err)
		return nil, err
	}
	return &resp, nil
}
