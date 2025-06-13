package filesystem

import (
	"example.com/xxx/common-lib/oss/storage"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LocalFileSystem 本地文件存储
type LocalFileSystem struct {
	config *storage.Config
}

// NewLocalFileSystem 初始化本地文件存储
func NewLocalFileSystem(conf *storage.Config) *LocalFileSystem {
	if conf.BaseFolder == "" {
		conf.BaseFolder = "files"
	}
	var err error
	conf.BaseFolder, err = filepath.Abs(conf.BaseFolder)
	if err != nil {
		panic(err)
	}
	return &LocalFileSystem{config: conf}
}

// GetFullPath 获取完整路径
func (a *LocalFileSystem) GetFullPath(path string) string {
	fullPath := path
	if !strings.HasPrefix(path, a.config.BaseFolder) {
		fullPath, _ = filepath.Abs(filepath.Join(a.config.BaseFolder, path))
	}
	return fullPath
}

// Get 根据 path 获取文件
func (a *LocalFileSystem) Get(path string) (*os.File, error) {
	return os.Open(a.GetFullPath(path))
}

// GetStream 获取文件流
func (a *LocalFileSystem) GetStream(path string) (io.ReadCloser, error) {
	return os.Open(a.GetFullPath(path))
}

// Put 存储文件
// path: 存储的文件路径
// reader: 要存储的文件流
func (a *LocalFileSystem) Put(path string, reader io.Reader) (*storage.Object, error) {
	fullPath := a.GetFullPath(path)
	// 创建目录
	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("创建文件目录失败: %s,error: %s", filepath.Dir(fullPath), err.Error())
	}
	// 创建文件
	dst, err := os.Create(filepath.Clean(fullPath))
	if err == nil {
		defer dst.Close()
		if seeker, ok := reader.(io.ReadSeeker); ok {
			seeker.Seek(0, 0)
		}
		_, err = io.Copy(dst, reader)
	}
	return &storage.Object{Path: path, Name: filepath.Base(path), LastModified: time.Now()}, err
}

// Delete 删除文件
func (a *LocalFileSystem) Delete(path string) error {
	return os.Remove(a.GetFullPath(path))
}

// List 获取文件列表
func (a *LocalFileSystem) List(path string) ([]*storage.Object, error) {
	var objects []*storage.Object
	fullPath := a.GetFullPath(path)
	err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if path == fullPath {
			return nil
		}
		if err == nil && !info.IsDir() {
			objects = append(objects, &storage.Object{
				Path:         strings.TrimPrefix(path, a.config.BaseFolder),
				Name:         info.Name(),
				LastModified: info.ModTime(),
				Size:         info.Size(),
			})
		}
		return nil
	})
	return objects, err
}

// GetEndpoint 获取端点，本地文件就是 /
func (a *LocalFileSystem) GetEndpoint() string {
	return "/"
}

// GetURL 获取文件URL(本地文件原样返回 path)
func (a *LocalFileSystem) GetURL(path string) (url string, err error) {
	return path, nil
}

func (a *LocalFileSystem) GetType() string {
	return storage.TypeFileSystem
}
