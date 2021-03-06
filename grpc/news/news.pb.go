// Code generated by protoc-gen-go.
// source: news.proto
// DO NOT EDIT!

/*
Package news is a generated protocol buffer package.

It is generated from these files:
	news.proto

It has these top-level messages:
	NewsInfo
	GetNewsRequest
	GetNewsReply
	PostNewsRequest
	PostNewsReply
	RecallNewsRequest
	RecallNewsReply
	LikeNewsRequest
	LikeNewsReply
*/
package news

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewsInfo struct {
	Uid      string `protobuf:"bytes,1,opt,name=Uid,json=uid" json:"Uid,omitempty"`
	Likes    int32  `protobuf:"varint,2,opt,name=Likes,json=likes" json:"Likes,omitempty"`
	Fowards  int32  `protobuf:"varint,3,opt,name=Fowards,json=fowards" json:"Fowards,omitempty"`
	MeipaiID string `protobuf:"bytes,4,opt,name=MeipaiID,json=meipaiID" json:"MeipaiID,omitempty"`
	Values   []byte `protobuf:"bytes,5,opt,name=Values,json=values,proto3" json:"Values,omitempty"`
}

func (m *NewsInfo) Reset()                    { *m = NewsInfo{} }
func (m *NewsInfo) String() string            { return proto.CompactTextString(m) }
func (*NewsInfo) ProtoMessage()               {}
func (*NewsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetNewsRequest struct {
	Uid   string `protobuf:"bytes,1,opt,name=Uid,json=uid" json:"Uid,omitempty"`
	Index uint64 `protobuf:"varint,2,opt,name=Index,json=index" json:"Index,omitempty"`
}

func (m *GetNewsRequest) Reset()                    { *m = GetNewsRequest{} }
func (m *GetNewsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNewsRequest) ProtoMessage()               {}
func (*GetNewsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetNewsReply struct {
	Status int32       `protobuf:"varint,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Index  uint64      `protobuf:"varint,2,opt,name=Index,json=index" json:"Index,omitempty"`
	News   []*NewsInfo `protobuf:"bytes,3,rep,name=News,json=news" json:"News,omitempty"`
}

func (m *GetNewsReply) Reset()                    { *m = GetNewsReply{} }
func (m *GetNewsReply) String() string            { return proto.CompactTextString(m) }
func (*GetNewsReply) ProtoMessage()               {}
func (*GetNewsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetNewsReply) GetNews() []*NewsInfo {
	if m != nil {
		return m.News
	}
	return nil
}

type PostNewsRequest struct {
	Uid       string `protobuf:"bytes,1,opt,name=Uid,json=uid" json:"Uid,omitempty"`
	Devid     string `protobuf:"bytes,2,opt,name=Devid,json=devid" json:"Devid,omitempty"`
	TimeStamp string `protobuf:"bytes,3,opt,name=TimeStamp,json=timeStamp" json:"TimeStamp,omitempty"`
	MeipaiID  string `protobuf:"bytes,4,opt,name=MeipaiID,json=meipaiID" json:"MeipaiID,omitempty"`
	Values    []byte `protobuf:"bytes,5,opt,name=Values,json=values,proto3" json:"Values,omitempty"`
}

func (m *PostNewsRequest) Reset()                    { *m = PostNewsRequest{} }
func (m *PostNewsRequest) String() string            { return proto.CompactTextString(m) }
func (*PostNewsRequest) ProtoMessage()               {}
func (*PostNewsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PostNewsReply struct {
	Status int32  `protobuf:"varint,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Newsid string `protobuf:"bytes,2,opt,name=Newsid,json=newsid" json:"Newsid,omitempty"`
}

func (m *PostNewsReply) Reset()                    { *m = PostNewsReply{} }
func (m *PostNewsReply) String() string            { return proto.CompactTextString(m) }
func (*PostNewsReply) ProtoMessage()               {}
func (*PostNewsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RecallNewsRequest struct {
	Newsid string `protobuf:"bytes,1,opt,name=Newsid,json=newsid" json:"Newsid,omitempty"`
}

func (m *RecallNewsRequest) Reset()                    { *m = RecallNewsRequest{} }
func (m *RecallNewsRequest) String() string            { return proto.CompactTextString(m) }
func (*RecallNewsRequest) ProtoMessage()               {}
func (*RecallNewsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RecallNewsReply struct {
	Status int32 `protobuf:"varint,1,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *RecallNewsReply) Reset()                    { *m = RecallNewsReply{} }
func (m *RecallNewsReply) String() string            { return proto.CompactTextString(m) }
func (*RecallNewsReply) ProtoMessage()               {}
func (*RecallNewsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LikeNewsRequest struct {
	Newsid string `protobuf:"bytes,1,opt,name=Newsid,json=newsid" json:"Newsid,omitempty"`
}

func (m *LikeNewsRequest) Reset()                    { *m = LikeNewsRequest{} }
func (m *LikeNewsRequest) String() string            { return proto.CompactTextString(m) }
func (*LikeNewsRequest) ProtoMessage()               {}
func (*LikeNewsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type LikeNewsReply struct {
	Status int32 `protobuf:"varint,1,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *LikeNewsReply) Reset()                    { *m = LikeNewsReply{} }
func (m *LikeNewsReply) String() string            { return proto.CompactTextString(m) }
func (*LikeNewsReply) ProtoMessage()               {}
func (*LikeNewsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*NewsInfo)(nil), "news.NewsInfo")
	proto.RegisterType((*GetNewsRequest)(nil), "news.GetNewsRequest")
	proto.RegisterType((*GetNewsReply)(nil), "news.GetNewsReply")
	proto.RegisterType((*PostNewsRequest)(nil), "news.PostNewsRequest")
	proto.RegisterType((*PostNewsReply)(nil), "news.PostNewsReply")
	proto.RegisterType((*RecallNewsRequest)(nil), "news.RecallNewsRequest")
	proto.RegisterType((*RecallNewsReply)(nil), "news.RecallNewsReply")
	proto.RegisterType((*LikeNewsRequest)(nil), "news.LikeNewsRequest")
	proto.RegisterType((*LikeNewsReply)(nil), "news.LikeNewsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for News service

type NewsClient interface {
	GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsReply, error)
	GetMyNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsReply, error)
	PostNews(ctx context.Context, in *PostNewsRequest, opts ...grpc.CallOption) (*PostNewsReply, error)
	RecallNews(ctx context.Context, in *RecallNewsRequest, opts ...grpc.CallOption) (*RecallNewsReply, error)
	LikeNews(ctx context.Context, in *LikeNewsRequest, opts ...grpc.CallOption) (*LikeNewsReply, error)
}

type newsClient struct {
	cc *grpc.ClientConn
}

func NewNewsClient(cc *grpc.ClientConn) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsReply, error) {
	out := new(GetNewsReply)
	err := grpc.Invoke(ctx, "/news.News/GetNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetMyNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsReply, error) {
	out := new(GetNewsReply)
	err := grpc.Invoke(ctx, "/news.News/GetMyNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) PostNews(ctx context.Context, in *PostNewsRequest, opts ...grpc.CallOption) (*PostNewsReply, error) {
	out := new(PostNewsReply)
	err := grpc.Invoke(ctx, "/news.News/PostNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) RecallNews(ctx context.Context, in *RecallNewsRequest, opts ...grpc.CallOption) (*RecallNewsReply, error) {
	out := new(RecallNewsReply)
	err := grpc.Invoke(ctx, "/news.News/RecallNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) LikeNews(ctx context.Context, in *LikeNewsRequest, opts ...grpc.CallOption) (*LikeNewsReply, error) {
	out := new(LikeNewsReply)
	err := grpc.Invoke(ctx, "/news.News/LikeNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for News service

type NewsServer interface {
	GetNews(context.Context, *GetNewsRequest) (*GetNewsReply, error)
	GetMyNews(context.Context, *GetNewsRequest) (*GetNewsReply, error)
	PostNews(context.Context, *PostNewsRequest) (*PostNewsReply, error)
	RecallNews(context.Context, *RecallNewsRequest) (*RecallNewsReply, error)
	LikeNews(context.Context, *LikeNewsRequest) (*LikeNewsReply, error)
}

func RegisterNewsServer(s *grpc.Server, srv NewsServer) {
	s.RegisterService(&_News_serviceDesc, srv)
}

func _News_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetMyNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetMyNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/GetMyNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetMyNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_PostNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).PostNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/PostNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).PostNews(ctx, req.(*PostNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_RecallNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecallNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).RecallNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/RecallNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).RecallNews(ctx, req.(*RecallNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_LikeNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).LikeNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/LikeNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).LikeNews(ctx, req.(*LikeNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _News_serviceDesc = grpc.ServiceDesc{
	ServiceName: "news.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _News_GetNews_Handler,
		},
		{
			MethodName: "GetMyNews",
			Handler:    _News_GetMyNews_Handler,
		},
		{
			MethodName: "PostNews",
			Handler:    _News_PostNews_Handler,
		},
		{
			MethodName: "RecallNews",
			Handler:    _News_RecallNews_Handler,
		},
		{
			MethodName: "LikeNews",
			Handler:    _News_LikeNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("news.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0xa5, 0xf4, 0x83, 0xf6, 0xca, 0x87, 0x8e, 0x80, 0x4d, 0xe3, 0xa2, 0xe9, 0xc6, 0x1a, 0x13,
	0x16, 0x18, 0x23, 0x2b, 0xdd, 0x10, 0x09, 0x89, 0x18, 0x33, 0xa8, 0x6b, 0xab, 0x1d, 0x92, 0x89,
	0xa5, 0xad, 0xcc, 0x14, 0x1e, 0xbb, 0xb7, 0x7f, 0xff, 0xe4, 0xfd, 0xca, 0x97, 0x99, 0x52, 0x0a,
	0x85, 0xf0, 0x3e, 0x76, 0x9c, 0xcb, 0x39, 0xf7, 0x9e, 0x39, 0xf7, 0x16, 0x20, 0x26, 0x1b, 0x36,
	0x48, 0x57, 0x09, 0x4f, 0x90, 0x26, 0x7e, 0x7b, 0xd7, 0x0a, 0x98, 0xdf, 0xc8, 0x86, 0x4d, 0xe3,
	0x45, 0x82, 0x9e, 0x83, 0xfa, 0x93, 0x86, 0xb6, 0xe2, 0x2a, 0xbe, 0x85, 0xd5, 0x8c, 0x86, 0xa8,
	0x0b, 0xfa, 0x57, 0xfa, 0x8f, 0x30, 0xbb, 0xee, 0x2a, 0xbe, 0x8e, 0xf5, 0x48, 0x00, 0x64, 0x43,
	0xe3, 0x4b, 0xb2, 0x09, 0x56, 0x21, 0xb3, 0x55, 0x59, 0x6f, 0x2c, 0x72, 0x88, 0x1c, 0x30, 0x67,
	0x84, 0xa6, 0x01, 0x9d, 0x8e, 0x6d, 0x4d, 0xb6, 0x31, 0x97, 0x3b, 0x8c, 0xfa, 0x60, 0xfc, 0x0a,
	0xa2, 0x8c, 0x30, 0x5b, 0x77, 0x15, 0xbf, 0x89, 0x8d, 0xb5, 0x44, 0xde, 0x08, 0xda, 0x13, 0xc2,
	0x85, 0x09, 0x4c, 0xfe, 0x67, 0x84, 0xf1, 0xf3, 0x3e, 0xa6, 0x71, 0x48, 0xae, 0xa4, 0x0f, 0x0d,
	0xeb, 0x54, 0x00, 0xef, 0x37, 0x34, 0xf7, 0xca, 0x34, 0xda, 0x8a, 0x09, 0x73, 0x1e, 0xf0, 0x8c,
	0x49, 0xa9, 0x8e, 0x0d, 0x26, 0xd1, 0x79, 0x35, 0xf2, 0x40, 0x13, 0x52, 0x5b, 0x75, 0x55, 0xff,
	0xd9, 0xb0, 0x3d, 0x90, 0xd9, 0x14, 0x59, 0xe0, 0x3c, 0x9e, 0x1b, 0x05, 0x3a, 0xdf, 0x13, 0x76,
	0xbf, 0xbb, 0x31, 0x59, 0xd3, 0x50, 0xf6, 0xb7, 0xb0, 0x1e, 0x0a, 0x80, 0x5e, 0x83, 0xf5, 0x83,
	0x2e, 0xc9, 0x9c, 0x07, 0xcb, 0x54, 0xe6, 0x64, 0x61, 0x8b, 0x17, 0x85, 0x27, 0x25, 0xf5, 0x19,
	0x5a, 0xa5, 0x99, 0x4b, 0x0f, 0xee, 0x83, 0x21, 0x48, 0x7b, 0x47, 0x46, 0x2c, 0x91, 0xf7, 0x0e,
	0x5e, 0x60, 0xf2, 0x37, 0x88, 0xa2, 0xc3, 0xf7, 0x94, 0x64, 0xe5, 0x88, 0xfc, 0x16, 0x3a, 0x87,
	0xe4, 0x0b, 0xf3, 0x04, 0x55, 0x9c, 0xc9, 0x43, 0xba, 0xbe, 0x81, 0x56, 0x49, 0xbd, 0xd0, 0x73,
	0x78, 0x5b, 0xcf, 0xf7, 0x83, 0x3e, 0x40, 0x63, 0xb7, 0x65, 0xd4, 0xcd, 0x97, 0x74, 0x7c, 0x2e,
	0x0e, 0xaa, 0x54, 0xd3, 0x68, 0xeb, 0xd5, 0xd0, 0x47, 0xb0, 0x26, 0x84, 0xcf, 0xb6, 0x8f, 0x16,
	0x8e, 0xc0, 0x2c, 0x52, 0x46, 0xbd, 0x9c, 0x51, 0x39, 0x01, 0xe7, 0x65, 0xb5, 0x9c, 0x2b, 0x3f,
	0x01, 0x94, 0x89, 0xa1, 0x57, 0x39, 0xe9, 0x24, 0x70, 0xa7, 0x77, 0xfa, 0xc7, 0x7e, 0x72, 0x91,
	0x4d, 0x31, 0xb9, 0x12, 0x6b, 0x31, 0xf9, 0x28, 0x42, 0xaf, 0xf6, 0xc7, 0x90, 0xdf, 0xf4, 0xfb,
	0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x7c, 0x76, 0xf0, 0xe1, 0x03, 0x00, 0x00,
}
