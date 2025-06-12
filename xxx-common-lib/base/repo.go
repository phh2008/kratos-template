package base

import (
    "context"
)

type IBaseRepo interface {
    // Transaction 开启事务
    Transaction(c context.Context, handler func(tx context.Context) error) error
}
