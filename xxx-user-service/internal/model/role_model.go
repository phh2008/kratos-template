package model

import (
	"example.com/xxx/common-lib/model/page"
	"time"
)

type RoleModel struct {
	Id        int64     `json:"id"`                          // 主键id
	RoleCode  string    `json:"roleCode" binding:"required"` // 角色编号
	RoleName  string    `json:"roleName" binding:"required"` // 角色名称
	CreatedAt time.Time `json:"createdAt"`                   // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                   // 更新时间
	CreatedBy string    `json:"createdBy"`                   // 创建人
	UpdatedBy string    `json:"updatedBy"`                   // 更新人
}

type RoleListReq struct {
	page.QueryPage
	RoleCode string `json:"roleCode" form:"roleCode"` // 角色编号
	RoleName string `json:"roleName" form:"roleName"` // 角色名称
}

type RoleAssignPermModel struct {
	RoleId     int64   `json:"roleId" binding:"required"`     // 角色ID
	PermIdList []int64 `json:"permIdList" binding:"required"` // 权限ID列表
}

type RolePermission struct {
	Id     int64 // 主键id
	RoleId int64 // 角色id
	PermId int64 // 权限id
}
