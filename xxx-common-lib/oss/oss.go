package oss

import (
    "example.com/xxx/common-lib/oss/filesystem"
    "example.com/xxx/common-lib/oss/storage"
    "example.com/xxx/common-lib/oss/tencent"
)

func NewStorage(conf *storage.Config) storage.Storage {
    switch conf.OssType {
    case storage.TypeFileSystem:
        return filesystem.NewLocalFileSystem(conf)
    case storage.TypeTencentCos:
        return tencent.NewClient(conf)
    default:
        return nil
    }
}
