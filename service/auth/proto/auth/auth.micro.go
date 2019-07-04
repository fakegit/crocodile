// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package crocodile_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	ChangeUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	LoginUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	LogoutUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "crocodile.srv.auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ChangeUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ChangeUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.GetUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) LoginUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.LoginUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) LogoutUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.LogoutUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	CreateUser(context.Context, *User, *Response) error
	ChangeUser(context.Context, *User, *Response) error
	DeleteUser(context.Context, *User, *Response) error
	GetUser(context.Context, *User, *Response) error
	LoginUser(context.Context, *User, *Response) error
	LogoutUser(context.Context, *User, *Response) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		CreateUser(ctx context.Context, in *User, out *Response) error
		ChangeUser(ctx context.Context, in *User, out *Response) error
		DeleteUser(ctx context.Context, in *User, out *Response) error
		GetUser(ctx context.Context, in *User, out *Response) error
		LoginUser(ctx context.Context, in *User, out *Response) error
		LogoutUser(ctx context.Context, in *User, out *Response) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) CreateUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.CreateUser(ctx, in, out)
}

func (h *authHandler) ChangeUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.ChangeUser(ctx, in, out)
}

func (h *authHandler) DeleteUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.DeleteUser(ctx, in, out)
}

func (h *authHandler) GetUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.GetUser(ctx, in, out)
}

func (h *authHandler) LoginUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.LoginUser(ctx, in, out)
}

func (h *authHandler) LogoutUser(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.LogoutUser(ctx, in, out)
}
