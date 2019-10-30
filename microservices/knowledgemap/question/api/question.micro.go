// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/question.proto

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

// Client API for Question service

type QuestionService interface {
	GetMyQuestionInfo(ctx context.Context, in *CRqQueryMyQuestionInfoBySubject, opts ...client.CallOption) (*CRpMyQuestionInfoBySubject, error)
}

type questionService struct {
	c    client.Client
	name string
}

func NewQuestionService(name string, c client.Client) QuestionService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &questionService{
		c:    c,
		name: name,
	}
}

func (c *questionService) GetMyQuestionInfo(ctx context.Context, in *CRqQueryMyQuestionInfoBySubject, opts ...client.CallOption) (*CRpMyQuestionInfoBySubject, error) {
	req := c.c.NewRequest(c.name, "Question.GetMyQuestionInfo", in)
	out := new(CRpMyQuestionInfoBySubject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Question service

type QuestionHandler interface {
	GetMyQuestionInfo(context.Context, *CRqQueryMyQuestionInfoBySubject, *CRpMyQuestionInfoBySubject) error
}

func RegisterQuestionHandler(s server.Server, hdlr QuestionHandler, opts ...server.HandlerOption) error {
	type question interface {
		GetMyQuestionInfo(ctx context.Context, in *CRqQueryMyQuestionInfoBySubject, out *CRpMyQuestionInfoBySubject) error
	}
	type Question struct {
		question
	}
	h := &questionHandler{hdlr}
	return s.Handle(s.NewHandler(&Question{h}, opts...))
}

type questionHandler struct {
	QuestionHandler
}

func (h *questionHandler) GetMyQuestionInfo(ctx context.Context, in *CRqQueryMyQuestionInfoBySubject, out *CRpMyQuestionInfoBySubject) error {
	return h.QuestionHandler.GetMyQuestionInfo(ctx, in, out)
}
