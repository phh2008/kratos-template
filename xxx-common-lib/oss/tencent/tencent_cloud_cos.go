package tencent

import (
    "context"
    "example.com/xxx/common-lib/oss/storage"
    "fmt"
    "github.com/tencentyun/cos-go-sdk-v5"
    "io"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "time"
)

type Client struct {
    config *storage.Config
    client *cos.Client
}

// Config 腾讯云 cos client config

func NewClient(conf *storage.Config) *Client {
    u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", conf.Bucket, conf.Region))
    su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", conf.Region))
    b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  conf.AccessID,
            SecretKey: conf.AccessKey,
        },
    })
    return &Client{
        config: conf,
        client: client,
    }
}

// Get 从COS下载到本地临时文件返回
func (a *Client) Get(path string) (file *os.File, err error) {
    // 创建临时文件
    if file, err = os.CreateTemp("/tmp", "tencent"); err == nil {
        readCloser, err := a.GetStream(path)
        if err != nil {
            return nil, err
        }
        defer readCloser.Close()
        _, err = io.Copy(file, readCloser)
        _, _ = file.Seek(0, 0)
    }
    return file, err
}

// GetStream 从COS下载到流，用完记得关闭流
func (a *Client) GetStream(path string) (io.ReadCloser, error) {
    // 腾讯云 cos 简单接口下载文件
    resp, err := a.client.Object.Get(context.Background(), path, nil)
    if err != nil {
        return nil, err
    }
    return resp.Body, nil
}

func (a *Client) Put(path string, body io.Reader) (*storage.Object, error) {
    _, err := a.client.Object.Put(context.Background(), path, body, nil)
    if err != nil {
        return nil, err
    }
    return &storage.Object{
        Path:         path,
        Name:         filepath.Base(path),
        LastModified: time.Now(),
    }, nil
}

// Delete 删除对象
func (a *Client) Delete(path string) error {
    _, err := a.client.Object.Delete(context.Background(), path)
    return err
}

// List 列出对象列表
func (a *Client) List(path string) ([]*storage.Object, error) {
    var objects []*storage.Object
    opt := &cos.BucketGetOptions{
        Prefix:    path,
        Delimiter: "/",
        MaxKeys:   100,
    }
    for {
        result, _, err := a.client.Bucket.Get(context.Background(), opt)
        if err != nil {
            return nil, err
        }
        for _, obj := range result.Contents {
            modified, _ := time.Parse(time.RFC3339, obj.LastModified)
            objects = append(objects, &storage.Object{
                Path:         obj.Key,
                Name:         filepath.Base(obj.Key),
                LastModified: modified,
                Size:         obj.Size,
            })
        }
        if result.IsTruncated && result.NextMarker != "" {
            opt.Marker = result.NextMarker
            continue
        }
        break
    }
    return objects, nil
}

// GetEndpoint 获取端点，e.g.: bucket-1000000.cos.ap-guangzhou.myqcloud.com
func (a *Client) GetEndpoint() string {
    return fmt.Sprintf("%s.cos.%s.myqcloud.com", a.config.Bucket, a.config.Region)
}

// GetURL 获取文件URL
func (a *Client) GetURL(path string) (string, error) {
    ourl := a.client.Object.GetObjectURL(path)
    return ourl.String(), nil
}

func (a *Client) GetType() string {
    return storage.TypeTencentCos
}
