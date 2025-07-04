// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.23.4
// source: user/v1/permission.proto

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

const OperationPermissionAdd = "/api.user.v1.Permission/Add"
const OperationPermissionListPage = "/api.user.v1.Permission/ListPage"
const OperationPermissionUpdate = "/api.user.v1.Permission/Update"

type PermissionHTTPServer interface {
	Add(context.Context, *PermSaveRequest) (*PermReply, error)
	ListPage(context.Context, *PermListRequest) (*PermListReply, error)
	Update(context.Context, *PermSaveRequest) (*PermReply, error)
}

func RegisterPermissionHTTPServer(s *http.Server, srv PermissionHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/permission/list", _Permission_ListPage0_HTTP_Handler(srv))
	r.POST("/v1/permission/add", _Permission_Add0_HTTP_Handler(srv))
	r.POST("/v1/permission/update", _Permission_Update0_HTTP_Handler(srv))
}

func _Permission_ListPage0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermissionListPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPage(ctx, req.(*PermListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermListReply)
		return ctx.Result(200, reply)
	}
}

func _Permission_Add0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermSaveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermissionAdd)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Add(ctx, req.(*PermSaveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermReply)
		return ctx.Result(200, reply)
	}
}

func _Permission_Update0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermSaveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPermissionUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*PermSaveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermReply)
		return ctx.Result(200, reply)
	}
}

type PermissionHTTPClient interface {
	Add(ctx context.Context, req *PermSaveRequest, opts ...http.CallOption) (rsp *PermReply, err error)
	ListPage(ctx context.Context, req *PermListRequest, opts ...http.CallOption) (rsp *PermListReply, err error)
	Update(ctx context.Context, req *PermSaveRequest, opts ...http.CallOption) (rsp *PermReply, err error)
}

type PermissionHTTPClientImpl struct {
	cc *http.Client
}

func NewPermissionHTTPClient(client *http.Client) PermissionHTTPClient {
	return &PermissionHTTPClientImpl{client}
}

func (c *PermissionHTTPClientImpl) Add(ctx context.Context, in *PermSaveRequest, opts ...http.CallOption) (*PermReply, error) {
	var out PermReply
	pattern := "/v1/permission/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPermissionAdd))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PermissionHTTPClientImpl) ListPage(ctx context.Context, in *PermListRequest, opts ...http.CallOption) (*PermListReply, error) {
	var out PermListReply
	pattern := "/v1/permission/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPermissionListPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PermissionHTTPClientImpl) Update(ctx context.Context, in *PermSaveRequest, opts ...http.CallOption) (*PermReply, error) {
	var out PermReply
	pattern := "/v1/permission/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPermissionUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
