package biz

import (
    "context"
    "example.com/xxx/common-lib/base"
    "example.com/xxx/common-lib/model/page"
    "example.com/xxx/user-service/internal/model"
    xjwt "example.com/xxx/user-service/internal/pkg/jwt"
    "github.com/casbin/casbin/v2"
    "github.com/cristalhq/jwt/v5"
    "github.com/go-kratos/kratos/v2/errors"
    "golang.org/x/crypto/bcrypt"
    "log/slog"
    "net/http"
    "strconv"
    "time"
)

type UserRepo interface {
    base.IBaseRepo
    ListPage(ctx context.Context, req model.UserListReq) (*page.PageData[*model.User], error)
    // GetByEmail 根据 email 查询
    GetByEmail(ctx context.Context, email string) (*model.User, error)
    // Add 添加用户
    Add(ctx context.Context, user model.User) (*model.User, error)
    // SetRole 设置角色
    SetRole(ctx context.Context, userId int64, role string) error
    // DeleteById 删除用户
    DeleteById(ctx context.Context, id int64) error
    // CancelRole 撤销用户角色
    CancelRole(ctx context.Context, roleCode string) error
    // CheckPermission 验证权限
    CheckPermission(ctx context.Context, req model.CheckPermReq) error
    // VerifyToken 验证token
    VerifyToken(ctx context.Context, token string) (*xjwt.UserClaims, error)
}

// UserUseCase 用户业务逻辑封装
type UserUseCase struct {
    userRepo UserRepo
    jwt      *xjwt.JwtHelper
    enforcer *casbin.Enforcer
}

// NewUserUseCase 构造业务结构体
func NewUserUseCase(userRepo UserRepo, jwt *xjwt.JwtHelper, enforcer *casbin.Enforcer) *UserUseCase {
    return &UserUseCase{
        userRepo: userRepo,
        jwt:      jwt,
        enforcer: enforcer,
    }
}

// ListPage 用户列表
func (a *UserUseCase) ListPage(ctx context.Context, req model.UserListReq) (*page.PageData[*model.User], error) {
    return a.userRepo.ListPage(ctx, req)
}

// CreateByEmail 根据邮箱创建用户
func (a *UserUseCase) CreateByEmail(ctx context.Context, email model.UserEmailRegister) (*model.User, error) {
    user, err := a.userRepo.GetByEmail(ctx, email.Email)
    if err != nil {
        return nil, err
    }
    if user.Id > 0 {
        return nil, errors.New(http.StatusConflict, "EMAIL_EXISTS", "邮箱已存在")
    }
    pwd, err := bcrypt.GenerateFromPassword([]byte(email.Password), 1)
    if err != nil {
        slog.Error("生成密码出错", "error", err.Error())
        return nil, err
    }
    user = &model.User{
        Email:    email.Email,
        RealName: email.Email,
        UserName: email.Email,
        Password: string(pwd),
        Status:   1,
        RoleCode: "",
    }
    ret, err := a.userRepo.Add(ctx, *user)
    if err != nil {
        slog.Error("创建用户出错", "error", err.Error())
        return nil, err
    }
    return ret, nil
}

// LoginByEmail 邮箱登录
func (a *UserUseCase) LoginByEmail(ctx context.Context, loginModel model.UserLoginModel) (string, error) {
    user, err := a.userRepo.GetByEmail(ctx, loginModel.Email)
    if err != nil {
        return "", err
    }
    if user.Id == 0 {
        return "", errors.New(http.StatusBadRequest, "PARAM_ERROR", "用户或密码错误")
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password))
    if err != nil {
        return "", errors.New(http.StatusBadRequest, "PARAM_ERROR", "用户或密码错误")
    }
    // 生成token
    userClaims := xjwt.UserClaims{}
    userClaims.ID = strconv.FormatInt(user.Id, 10)
    userClaims.Role = user.RoleCode
    userClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))
    token, err := a.jwt.CreateToken(userClaims)
    if err != nil {
        slog.Error("生成token错误", "error", err.Error())
        return "", err
    }
    return token.String(), nil
}

// AssignRole 给用户分配角色
func (a *UserUseCase) AssignRole(ctx context.Context, userRole model.AssignRoleModel) error {
    err := a.userRepo.SetRole(ctx, userRole.UserId, userRole.RoleCode)
    if err != nil {
        slog.Error("db update error", "error", err.Error())
        return errors.New(http.StatusInternalServerError, "DB_ERROR", "db update error")
    }
    // 更新casbin中的用户与角色关系
    uid := strconv.FormatInt(userRole.UserId, 10)
    _, _ = a.enforcer.DeleteRolesForUser(uid)
    // 角色为空，表示清除此用户的角色,无需添加
    if userRole.RoleCode != "" {
        _, _ = a.enforcer.AddGroupingPolicy(uid, userRole.RoleCode)
    }
    return nil
}

// DeleteById 根据ID删除
func (a *UserUseCase) DeleteById(ctx context.Context, id int64) error {
    err := a.userRepo.DeleteById(ctx, id)
    if err != nil {
        slog.Error("delete error", "error", err.Error())
        return errors.New(http.StatusInternalServerError, "DB_ERROR", "db update error")
    }
    // 清除 casbin 中用户信息
    _, err = a.enforcer.DeleteRolesForUser(strconv.FormatInt(id, 10))
    if err != nil {
        slog.Error("Enforcer.DeleteRolesForUser error", "error", err)
    }
    return nil
}

func (a *UserUseCase) CheckPermission(ctx context.Context, req model.CheckPermReq) error {
    return a.userRepo.CheckPermission(ctx, req)
}

func (a *UserUseCase) VerifyToken(ctx context.Context, token string) (*xjwt.UserClaims, error) {
    return a.userRepo.VerifyToken(ctx, token)
}
