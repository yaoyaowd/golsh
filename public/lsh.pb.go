// Code generated by protoc-gen-go.
// source: lsh.proto
// DO NOT EDIT!

/*
Package lsh is a generated protocol buffer package.

It is generated from these files:
	lsh.proto

It has these top-level messages:
	SearchRequest
	SearchDoc
	SearchResult
*/
package lsh

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

type SearchRequest struct {
	Cid string `protobuf:"bytes,1,opt,name=cid" json:"cid,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type SearchDoc struct {
	Id      string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Score   float32 `protobuf:"fixed32,2,opt,name=score" json:"score,omitempty"`
	Explain string  `protobuf:"bytes,3,opt,name=explain" json:"explain,omitempty"`
}

func (m *SearchDoc) Reset()                    { *m = SearchDoc{} }
func (m *SearchDoc) String() string            { return proto.CompactTextString(m) }
func (*SearchDoc) ProtoMessage()               {}
func (*SearchDoc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchDoc) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchDoc) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *SearchDoc) GetExplain() string {
	if m != nil {
		return m.Explain
	}
	return ""
}

type SearchResult struct {
	Docs []*SearchDoc `protobuf:"bytes,1,rep,name=docs" json:"docs,omitempty"`
	Hits int32        `protobuf:"varint,2,opt,name=hits" json:"hits,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchResult) GetDocs() []*SearchDoc {
	if m != nil {
		return m.Docs
	}
	return nil
}

func (m *SearchResult) GetHits() int32 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "lsh.SearchRequest")
	proto.RegisterType((*SearchDoc)(nil), "lsh.SearchDoc")
	proto.RegisterType((*SearchResult)(nil), "lsh.SearchResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LSH service

type LSHClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error)
}

type lSHClient struct {
	cc *grpc.ClientConn
}

func NewLSHClient(cc *grpc.ClientConn) LSHClient {
	return &lSHClient{cc}
}

func (c *lSHClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := grpc.Invoke(ctx, "/lsh.LSH/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LSH service

type LSHServer interface {
	Search(context.Context, *SearchRequest) (*SearchResult, error)
}

func RegisterLSHServer(s *grpc.Server, srv LSHServer) {
	s.RegisterService(&_LSH_serviceDesc, srv)
}

func _LSH_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LSHServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsh.LSH/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LSHServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lsh.LSH",
	HandlerType: (*LSHServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _LSH_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lsh.proto",
}

func init() { proto.RegisterFile("lsh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4a, 0xc6, 0x30,
	0x14, 0x84, 0x4d, 0xd2, 0x56, 0xfa, 0xd4, 0xa2, 0x0f, 0x17, 0xc1, 0x55, 0xcd, 0xaa, 0xab, 0x82,
	0x75, 0xe3, 0x01, 0x44, 0x04, 0x5d, 0xa5, 0x27, 0xa8, 0x69, 0x20, 0x81, 0x60, 0x6a, 0x93, 0x82,
	0xc7, 0x97, 0xa6, 0xb6, 0xfc, 0xff, 0x6e, 0x26, 0x33, 0xe1, 0x7b, 0x03, 0xa5, 0x0b, 0xa6, 0x9d,
	0x66, 0x1f, 0x3d, 0x32, 0x17, 0x8c, 0x78, 0x84, 0x9b, 0x5e, 0x0f, 0xb3, 0x32, 0x52, 0xff, 0x2c,
	0x3a, 0x44, 0xbc, 0x05, 0xa6, 0xec, 0xc8, 0x49, 0x4d, 0x9a, 0x52, 0xae, 0x52, 0x7c, 0x40, 0xb9,
	0x55, 0x5e, 0xbd, 0xc2, 0x0a, 0xe8, 0x91, 0x52, 0x3b, 0xe2, 0x3d, 0xe4, 0x41, 0xf9, 0x59, 0x73,
	0x5a, 0x93, 0x86, 0xca, 0xcd, 0x20, 0x87, 0x4b, 0xfd, 0x3b, 0xb9, 0xc1, 0x7e, 0x73, 0x96, 0xaa,
	0xbb, 0x15, 0x6f, 0x70, 0xbd, 0xf3, 0xc2, 0xe2, 0x22, 0x0a, 0xc8, 0x46, 0xaf, 0x02, 0x27, 0x35,
	0x6b, 0xae, 0xba, 0xaa, 0x5d, 0xcf, 0x3b, 0x68, 0x32, 0x65, 0x88, 0x90, 0x19, 0x1b, 0x43, 0x42,
	0xe4, 0x32, 0xe9, 0xee, 0x05, 0xd8, 0x67, 0xff, 0x8e, 0x4f, 0x50, 0x6c, 0x6d, 0xc4, 0x93, 0xaf,
	0xff, 0x5b, 0x1e, 0xee, 0xce, 0xde, 0x56, 0x9e, 0xb8, 0xf8, 0x2a, 0xd2, 0xfa, 0xe7, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x96, 0xf4, 0x79, 0x0a, 0x01, 0x00, 0x00,
}