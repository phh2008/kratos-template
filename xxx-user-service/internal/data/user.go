package data

import (
    "context"
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/common-lib/orm"
    v1 "example.com/xxx/user-service/api/user/v1"
    "example.com/xxx/user-service/internal/biz"
    "example.com/xxx/user-service/internal/model"
    xjwt "example.com/xxx/user-service/internal/pkg/jwt"
    "github.com/jinzhu/copier"
    "log/slog"
    "time"
)

type UserEntity struct {
    orm.BaseEntity
    RealName string `json:"realName"`                // 姓名
    UserName string `json:"userName"`                // 用户名
    Email    string `json:"email"`                   // 邮箱
    Password string `json:"password"`                // 密码
    Status   int    `gorm:"default:1" json:"status"` //状态: 1-启用，2-禁用
    RoleCode string `json:"roleCode"`                // 角色编号
}

func (UserEntity) TableName() string {
    return "sys_user"
}

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
    *orm.BaseRepo[UserEntity]
    data *Data
}

// NewUserRepo 创建 userRepo
func NewUserRepo(data *Data) biz.UserRepo {
    return &userRepo{
        BaseRepo: orm.NewBaseRepo[UserEntity](data.db),
        data:     data,
    }
}

func (a *userRepo) ListPage(ctx context.Context, req model.UserListReq) (*page.PageData[*model.User], error) {
    db := a.GetDB(ctx)
    db = db.Model(&UserEntity{})
    if req.RealName != "" {
        db = db.Where("real_name like ?", "%"+req.RealName+"%")
    }
    if req.Email != "" {
        db = db.Where("email like ?", "%"+req.Email+"%")
    }
    if req.Status != 0 {
        db = db.Where("status=?", req.Status)
    }
    var pageData page.PageData[*model.User]
    pageData.PageNo = req.GetPageNo()
    pageData.PageSize = req.GetPageSize()
    err := db.Count(&pageData.Count).
        Scopes(orm.OrderScope(req.Sort, req.Direction), orm.PageScope(req.GetPageNo(), req.GetPageSize())).
        Find(&pageData.Data).Error
    return &pageData, err
}

// GetByEmail 根据 email 查询
func (a *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
    var entity UserEntity
    err := a.GetDB(ctx).Model(UserEntity{}).Where("email=?", email).First(&entity).Error
    if err != nil {
        return nil, err
    }
    var user model.User
    _ = copier.Copy(&user, &entity)
    return &user, nil
}

// Add 添加用户
func (a *userRepo) Add(ctx context.Context, user model.User) (*model.User, error) {
    var entity UserEntity
    _ = copier.Copy(&entity, &user)
    err := a.GetDB(ctx).Create(&entity).Error
    _ = copier.Copy(&user, &entity)
    return &user, err
}

// SetRole 设置角色
func (a *userRepo) SetRole(ctx context.Context, userId int64, role string) error {
    db := a.GetDB(ctx).Model(&UserEntity{}).Where("id=?", userId).Update("role_code", role)
    return db.Error
}

// DeleteById 删除用户
func (a *userRepo) DeleteById(ctx context.Context, id int64) error {
    db := a.GetDB(ctx).Delete(&UserEntity{}, id)
    return db.Error
}

// CancelRole 撤销用户角色
func (a *userRepo) CancelRole(ctx context.Context, roleCode string) error {
    ret := a.GetDB(ctx).Model(&UserEntity{}).Where("role_code=?", roleCode).Update("role_code", "")
    return ret.Error
}

func (a *userRepo) CheckPermission(ctx context.Context, req model.CheckPermReq) error {
    ok, err := a.data.enforcer.Enforce(req.Sub, req.Obj, req.Act)
    if err != nil {
        slog.Error("Enforcer.Enforce", "error", err)
        // 鉴权出错了
        return v1.ErrorSysError("鉴权出错了,请联系管理员")
    }
    if !ok {
        // 无权限
        return v1.ErrorUnauthorized("无权限")
    }
    return nil
}

func (a *userRepo) VerifyToken(ctx context.Context, token string) (*xjwt.UserClaims, error) {
    if token == "" {
        return nil, v1.ErrorNoLogin("未登录")
    }
    jwtToken, err := a.data.jwt.VerifyToken(token)
    if err != nil {
        return nil, v1.ErrorNoLogin("未登录")
    }
    user, err := a.data.jwt.ParseToken(jwtToken)
    if err != nil {
        return nil, v1.ErrorNoLogin("未登录")
    }
    if !user.IsValidExpiresAt(time.Now()) {
        return nil, v1.ErrorNoLogin("未登录")
    }
    return &user, nil
}
