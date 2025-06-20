// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.23.4
// source: user/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserAssignRole = "/api.user.v1.User/AssignRole"
const OperationUserCheckPermission = "/api.user.v1.User/CheckPermission"
const OperationUserCreateByEmail = "/api.user.v1.User/CreateByEmail"
const OperationUserDeleteById = "/api.user.v1.User/DeleteById"
const OperationUserListPage = "/api.user.v1.User/ListPage"
const OperationUserLoginByEmail = "/api.user.v1.User/LoginByEmail"
const OperationUserVerifyToken = "/api.user.v1.User/VerifyToken"

type UserHTTPServer interface {
	// AssignRole 分配角色
	AssignRole(context.Context, *AssignRoleRequest) (*UserOk, error)
	// CheckPermission 校验权限
	CheckPermission(context.Context, *CheckPermissionRequest) (*UserOk, error)
	// CreateByEmail 邮箱注册
	CreateByEmail(context.Context, *UserEmailRequest) (*UserReply, error)
	// DeleteById 删除用户
	DeleteById(context.Context, *UserDeleteRequest) (*UserOk, error)
	// ListPage 用户例表
	ListPage(context.Context, *UserListRequest) (*UserListReply, error)
	// LoginByEmail 邮箱登录
	LoginByEmail(context.Context, *UserEmailLoginRequest) (*LoginReply, error)
	// VerifyToken 验证token
	VerifyToken(context.Context, *UserVerifyTokenRequest) (*UserClaimsReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/user/list", _User_ListPage2_HTTP_Handler(srv))
	r.POST("/v1/user/createByEmail", _User_CreateByEmail0_HTTP_Handler(srv))
	r.POST("/v1/user/login", _User_LoginByEmail0_HTTP_Handler(srv))
	r.POST("/v1/user/assignRole", _User_AssignRole0_HTTP_Handler(srv))
	r.POST("/v1/user/delete", _User_DeleteById1_HTTP_Handler(srv))
	r.POST("/v1/user/checkPermission", _User_CheckPermission0_HTTP_Handler(srv))
	r.POST("/v1/user/verifyToken", _User_VerifyToken0_HTTP_Handler(srv))
}

func _User_ListPage2_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPage(ctx, req.(*UserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserListReply)
		return ctx.Result(200, reply)
	}
}

func _User_CreateByEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateByEmail(ctx, req.(*UserEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginByEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserEmailLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByEmail(ctx, req.(*UserEmailLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_AssignRole0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssignRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAssignRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AssignRole(ctx, req.(*AssignRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserOk)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteById1_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteById(ctx, req.(*UserDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserOk)
		return ctx.Result(200, reply)
	}
}

func _User_CheckPermission0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCheckPermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckPermission(ctx, req.(*CheckPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserOk)
		return ctx.Result(200, reply)
	}
}

func _User_VerifyToken0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserVerifyTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserVerifyToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyToken(ctx, req.(*UserVerifyTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserClaimsReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AssignRole(ctx context.Context, req *AssignRoleRequest, opts ...http.CallOption) (rsp *UserOk, err error)
	CheckPermission(ctx context.Context, req *CheckPermissionRequest, opts ...http.CallOption) (rsp *UserOk, err error)
	CreateByEmail(ctx context.Context, req *UserEmailRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	DeleteById(ctx context.Context, req *UserDeleteRequest, opts ...http.CallOption) (rsp *UserOk, err error)
	ListPage(ctx context.Context, req *UserListRequest, opts ...http.CallOption) (rsp *UserListReply, err error)
	LoginByEmail(ctx context.Context, req *UserEmailLoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	VerifyToken(ctx context.Context, req *UserVerifyTokenRequest, opts ...http.CallOption) (rsp *UserClaimsReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AssignRole(ctx context.Context, in *AssignRoleRequest, opts ...http.CallOption) (*UserOk, error) {
	var out UserOk
	pattern := "/v1/user/assignRole"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAssignRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...http.CallOption) (*UserOk, error) {
	var out UserOk
	pattern := "/v1/user/checkPermission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCheckPermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) CreateByEmail(ctx context.Context, in *UserEmailRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/user/createByEmail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) DeleteById(ctx context.Context, in *UserDeleteRequest, opts ...http.CallOption) (*UserOk, error) {
	var out UserOk
	pattern := "/v1/user/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) ListPage(ctx context.Context, in *UserListRequest, opts ...http.CallOption) (*UserListReply, error) {
	var out UserListReply
	pattern := "/v1/user/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserListPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) LoginByEmail(ctx context.Context, in *UserEmailLoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) VerifyToken(ctx context.Context, in *UserVerifyTokenRequest, opts ...http.CallOption) (*UserClaimsReply, error) {
	var out UserClaimsReply
	pattern := "/v1/user/verifyToken"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserVerifyToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
