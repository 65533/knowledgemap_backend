// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/knowledgemap.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Client API for KnowledegeMap service

type KnowledegeMapService interface {
	GetKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMapBySubject, opts ...client.CallOption) (*KnowledegeMapInfo, error)
	GetMyKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMyMapBySubject, opts ...client.CallOption) (*KnowledegeMapInfo, error)
	CreateKnowledege(ctx context.Context, in *CreateKnowledegeReq, opts ...client.CallOption) (*KnowledegeInfoReply, error)
	QueryKnowledegeInfo(ctx context.Context, in *QueryKnowledegeInfoReq, opts ...client.CallOption) (*KnowledegeInfoReply, error)
}

type knowledegeMapService struct {
	c    client.Client
	name string
}

func NewKnowledegeMapService(name string, c client.Client) KnowledegeMapService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &knowledegeMapService{
		c:    c,
		name: name,
	}
}

func (c *knowledegeMapService) GetKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMapBySubject, opts ...client.CallOption) (*KnowledegeMapInfo, error) {
	req := c.c.NewRequest(c.name, "KnowledegeMap.GetKnowledegeMapBySubject", in)
	out := new(KnowledegeMapInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledegeMapService) GetMyKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMyMapBySubject, opts ...client.CallOption) (*KnowledegeMapInfo, error) {
	req := c.c.NewRequest(c.name, "KnowledegeMap.GetMyKnowledegeMapBySubject", in)
	out := new(KnowledegeMapInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledegeMapService) CreateKnowledege(ctx context.Context, in *CreateKnowledegeReq, opts ...client.CallOption) (*KnowledegeInfoReply, error) {
	req := c.c.NewRequest(c.name, "KnowledegeMap.CreateKnowledege", in)
	out := new(KnowledegeInfoReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledegeMapService) QueryKnowledegeInfo(ctx context.Context, in *QueryKnowledegeInfoReq, opts ...client.CallOption) (*KnowledegeInfoReply, error) {
	req := c.c.NewRequest(c.name, "KnowledegeMap.QueryKnowledegeInfo", in)
	out := new(KnowledegeInfoReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KnowledegeMap service

type KnowledegeMapHandler interface {
	GetKnowledegeMapBySubject(context.Context, *CRqQueryMapBySubject, *KnowledegeMapInfo) error
	GetMyKnowledegeMapBySubject(context.Context, *CRqQueryMyMapBySubject, *KnowledegeMapInfo) error
	CreateKnowledege(context.Context, *CreateKnowledegeReq, *KnowledegeInfoReply) error
	QueryKnowledegeInfo(context.Context, *QueryKnowledegeInfoReq, *KnowledegeInfoReply) error
}

func RegisterKnowledegeMapHandler(s server.Server, hdlr KnowledegeMapHandler, opts ...server.HandlerOption) error {
	type knowledegeMap interface {
		GetKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMapBySubject, out *KnowledegeMapInfo) error
		GetMyKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMyMapBySubject, out *KnowledegeMapInfo) error
		CreateKnowledege(ctx context.Context, in *CreateKnowledegeReq, out *KnowledegeInfoReply) error
		QueryKnowledegeInfo(ctx context.Context, in *QueryKnowledegeInfoReq, out *KnowledegeInfoReply) error
	}
	type KnowledegeMap struct {
		knowledegeMap
	}
	h := &knowledegeMapHandler{hdlr}
	return s.Handle(s.NewHandler(&KnowledegeMap{h}, opts...))
}

type knowledegeMapHandler struct {
	KnowledegeMapHandler
}

func (h *knowledegeMapHandler) GetKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMapBySubject, out *KnowledegeMapInfo) error {
	return h.KnowledegeMapHandler.GetKnowledegeMapBySubject(ctx, in, out)
}

func (h *knowledegeMapHandler) GetMyKnowledegeMapBySubject(ctx context.Context, in *CRqQueryMyMapBySubject, out *KnowledegeMapInfo) error {
	return h.KnowledegeMapHandler.GetMyKnowledegeMapBySubject(ctx, in, out)
}

func (h *knowledegeMapHandler) CreateKnowledege(ctx context.Context, in *CreateKnowledegeReq, out *KnowledegeInfoReply) error {
	return h.KnowledegeMapHandler.CreateKnowledege(ctx, in, out)
}

func (h *knowledegeMapHandler) QueryKnowledegeInfo(ctx context.Context, in *QueryKnowledegeInfoReq, out *KnowledegeInfoReply) error {
	return h.KnowledegeMapHandler.QueryKnowledegeInfo(ctx, in, out)
}
