// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/job/job.proto

package crocodile_srv_job

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Client API for Job service

type JobService interface {
	// 增
	CreateJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
	// 删
	DeleteJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
	// 改 停止 恢复
	ChangeJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
	// 查
	GetJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
	// 强杀
	KillJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
	// 运行
	RunJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error)
}

type jobService struct {
	c    client.Client
	name string
}

func NewJobService(name string, c client.Client) JobService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "crocodile.srv.job"
	}
	return &jobService{
		c:    c,
		name: name,
	}
}

func (c *jobService) CreateJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.CreateJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) DeleteJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.DeleteJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) ChangeJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.ChangeJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) GetJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.GetJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) KillJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.KillJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) RunJob(ctx context.Context, in *Task, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Job.RunJob", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobHandler interface {
	// 增
	CreateJob(context.Context, *Task, *Response) error
	// 删
	DeleteJob(context.Context, *Task, *Response) error
	// 改 停止 恢复
	ChangeJob(context.Context, *Task, *Response) error
	// 查
	GetJob(context.Context, *Task, *Response) error
	// 强杀
	KillJob(context.Context, *Task, *Response) error
	// 运行
	RunJob(context.Context, *Task, *Response) error
}

func RegisterJobHandler(s server.Server, hdlr JobHandler, opts ...server.HandlerOption) error {
	type job interface {
		CreateJob(ctx context.Context, in *Task, out *Response) error
		DeleteJob(ctx context.Context, in *Task, out *Response) error
		ChangeJob(ctx context.Context, in *Task, out *Response) error
		GetJob(ctx context.Context, in *Task, out *Response) error
		KillJob(ctx context.Context, in *Task, out *Response) error
		RunJob(ctx context.Context, in *Task, out *Response) error
	}
	type Job struct {
		job
	}
	h := &jobHandler{hdlr}
	return s.Handle(s.NewHandler(&Job{h}, opts...))
}

type jobHandler struct {
	JobHandler
}

func (h *jobHandler) CreateJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.CreateJob(ctx, in, out)
}

func (h *jobHandler) DeleteJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.DeleteJob(ctx, in, out)
}

func (h *jobHandler) ChangeJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.ChangeJob(ctx, in, out)
}

func (h *jobHandler) GetJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.GetJob(ctx, in, out)
}

func (h *jobHandler) KillJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.KillJob(ctx, in, out)
}

func (h *jobHandler) RunJob(ctx context.Context, in *Task, out *Response) error {
	return h.JobHandler.RunJob(ctx, in, out)
}
