package model

import "example.com/xxx/common-lib/types"

type FileAddReq struct {
	FilePath     string              // 文件路径(保存后路径)
	FileSize     int64               // 文件大小(字节)
	FileMd5      string              // 文件md5
	MediaType    string              // 媒介类型：image/jpeg 等
	OriginalName string              // 原始文件名
	CreatedAt    types.LocalDateTime // 创建时间
	UpdatedAt    types.LocalDateTime // 更新时间
	CreatedBy    string              // 创建人
	UpdatedBy    string              // 更新人
}

type FileAddResp struct {
	ID           int64  `json:"id"`           // 文件ID
	FilePath     string `json:"filePath"`     // 文件路径(保存后路径)
	FileSize     int64  `json:"fileSize"`     // 文件大小(字节)
	FileMd5      string `json:"fileMd5"`      // 文件md5
	MediaType    string `json:"mediaType"`    // 媒介类型：image/jpeg 等
	OriginalName string `json:"originalName"` // 原始文件名
}

type File struct {
	ID           int64               // 文件ID
	FilePath     string              // 文件路径(保存后路径)
	FileSize     int64               // 文件大小(字节)
	FileMd5      string              // 文件md5
	MediaType    string              // 媒介类型：image/jpeg 等
	OriginalName string              // 原始文件名
	CreatedAt    types.LocalDateTime // 创建时间
	UpdatedAt    types.LocalDateTime // 更新时间
	CreatedBy    string              // 创建人
	UpdatedBy    string              // 更新人
}
