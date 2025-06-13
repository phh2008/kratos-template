package storage

import (
    "io"
    "os"
    "time"
)

const (
    TypeFileSystem = "fileSystem"
    TypeTencentCos = "tencentCos"
)

type Storage interface {
    Get(path string) (*os.File, error)
    GetStream(path string) (io.ReadCloser, error)
    Put(path string, reader io.Reader) (*Object, error)
    Delete(path string) error
    List(path string) ([]*Object, error)
    GetURL(path string) (string, error)
    GetEndpoint() string
    GetType() string
}

type Config struct {
    OssType    string `json:"ossType"`    // 存储类型： fileSystem, tencentCos
    BaseFolder string `json:"baseFolder"` // 文件系统存储-存储的根目录
    AccessID   string `json:"accessID"`   // cos-访问ID
    AccessKey  string `json:"accessKey"`  // cos-访问密钥
    Region     string `json:"region"`     // cos-地域
    Bucket     string `json:"bucket"`     // cos-存储桶
}

type Object struct {
    Path         string    `json:"path"`         // 文件路径包含文件名(不包括根目录    )
    Name         string    `json:"name"`         // 文件名
    LastModified time.Time `json:"lastModified"` // 最后修改时间
    Size         int64     `json:"size"`         // 文件大小
}
