package orm

import (
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/common-lib/util"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func NewDB(dsn string) *gorm.DB {
    var gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        panic(err)
    }
    return gdb
}

func emptyScope() func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db
    }
}

func OrderScope(field string, direct string) func(db *gorm.DB) *gorm.DB {
    if field == "" {
        return emptyScope()
    }
    if !util.ColumnReg.MatchString(field) {
        // 非法字段
        return emptyScope()
    }
    if direct == "" || !util.DirectReg.MatchString(direct) {
        direct = "asc"
    }
    sort := fmt.Sprintf("%s %s", field, direct)
    return func(db *gorm.DB) *gorm.DB {
        return db.Order(sort)
    }
}

func PageScope(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if pageNo <= 0 {
            pageNo = 1
        }
        if pageSize <= 0 {
            pageSize = 10
        }
        offset := (pageNo - 1) * pageSize
        return db.Offset(offset).Limit(pageSize)
    }
}

func OrderPageScope[T any](page page.QueryPage, data *page.PageData[T]) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        pageNo := page.GetPageNo()
        pageSize := page.GetPageSize()
        data.PageNo = pageNo
        data.PageSize = pageSize
        offset := (pageNo - 1) * pageSize
        return db.Scopes(OrderScope(page.Sort, page.Direction)).Offset(offset).Limit(pageSize)
    }
}

// QueryPage 分页查询
func QueryPage[T any](db *gorm.DB, pageNo, pageSize int) (*page.PageData[T], error) {
    return QueryOrderPage[T](db, "", "", pageNo, pageSize)
}

// QueryOrderPage 排序分页查询
func QueryOrderPage[T any](db *gorm.DB, sortField string, direct string, pageNo, pageSize int) (*page.PageData[T], error) {
    if pageNo <= 0 {
        pageNo = 1
    }
    if pageSize <= 0 {
        pageSize = 10
    }
    var pageData page.PageData[T]
    pageData.PageNo = pageNo
    pageData.PageSize = pageSize
    offset := (pageNo - 1) * pageSize
    err := db.Count(&pageData.Count).Scopes(OrderScope(sortField, direct)).Offset(offset).Limit(pageSize).Find(&pageData.Data).Error
    return &pageData, err
}
