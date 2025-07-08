package model

import (
	"example.com/xxx/common-lib/model/page"
	"example.com/xxx/common-lib/types"
)

type User struct {
	Id        int64               `json:"id"`        // 主键id
	CreatedAt types.LocalDateTime `json:"createdAt"` // 创建时间
	UpdatedAt types.LocalDateTime `json:"updatedAt"` // 更新时间
	CreatedBy string              `json:"createdBy"` // 创建人
	UpdatedBy string              `json:"updatedBy"` // 更新人
	RealName  string              `json:"realName"`  // 姓名
	UserName  string              `json:"userName"`  // 用户名
	Email     string              `json:"email"`     // 邮箱
	Password  string              `json:"password"`  // 密码
	Status    int                 `json:"status"`    //状态: 1-启用，2-禁用
	RoleCode  string              `json:"roleCode"`  // 角色编号
}

type UserListReq struct {
	page.QueryPage
	RealName string `json:"realName" form:"realName"` // 姓名
	Email    string `json:"email" form:"email"`       // 邮箱
	Status   int    `json:"status" form:"status"`     // 状态: 1-启用，2-禁用
}

type UserEmailRegister struct {
	Email    string `json:"email" binding:"required"`    // 邮箱
	Password string `json:"password" binding:"required"` // 密码
}

type UserLoginModel struct {
	Email    string `json:"email" binding:"required"`    // 邮箱
	Password string `json:"password" binding:"required"` // 密码
}

type AssignRoleModel struct {
	UserId   int64  `json:"userId" binding:"required"`   // 用户ID
	RoleCode string `json:"roleCode" binding:"required"` // 角色编码
}

type CheckPermReq struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}
