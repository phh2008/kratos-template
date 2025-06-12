package xcasbin

import (
    "example.com/xxx/user-service/internal/conf"
    "example.com/xxx/user-service/internal/pkg/constant"
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "gorm.io/gorm"
    "log/slog"
    "path/filepath"
)

func NewCasbin(db *gorm.DB, c *conf.Bootstrap) *casbin.Enforcer {
    adapter, err := gormadapter.NewAdapterByDB(db)
    if err != nil {
        slog.Error("casbin gorm 适配器创建失败", "error", err.Error())
        panic(err)
    }
    configFile := filepath.Join(c.Folder, "rbac_model.conf")
    rbacEnforcer, err := casbin.NewEnforcer(configFile, adapter)
    if err != nil {
        slog.Error("casbin.NewEnforcer 错误", "error", err.Error())
        panic(err)
    }
    // 超级管理员角色处理
    rbacEnforcer.AddFunction("checkSuperAdmin", func(args ...interface{}) (interface{}, error) {
        username := args[0].(string)
        return rbacEnforcer.HasRoleForUser(username, constant.SuperAdmin)
    })
    rbacEnforcer.EnableAutoSave(true)
    // Load the policy from DB.
    _ = rbacEnforcer.LoadPolicy()
    return rbacEnforcer
}
