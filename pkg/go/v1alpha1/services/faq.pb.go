// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1alpha1/services/faq.proto

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	v1alpha1/services/faq.proto

It has these top-level messages:
	GetFaqRequest
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1_types "github.com/andrewhowdencom/mtm-apis/pkg/go/v1alpha1/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Takes a subset of the fields of "Faq"
type GetFaqRequest struct {
	FaqId string `protobuf:"bytes,1,opt,name=faq_id,json=faqId" json:"faq_id,omitempty"`
}

func (m *GetFaqRequest) Reset()                    { *m = GetFaqRequest{} }
func (m *GetFaqRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFaqRequest) ProtoMessage()               {}
func (*GetFaqRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetFaqRequest) GetFaqId() string {
	if m != nil {
		return m.FaqId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFaqRequest)(nil), "v1alpha1.services.GetFaqRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FaqService service

type FaqServiceClient interface {
	GetFaq(ctx context.Context, in *GetFaqRequest, opts ...grpc.CallOption) (*v1alpha1_types.Faq, error)
}

type faqServiceClient struct {
	cc *grpc.ClientConn
}

func NewFaqServiceClient(cc *grpc.ClientConn) FaqServiceClient {
	return &faqServiceClient{cc}
}

func (c *faqServiceClient) GetFaq(ctx context.Context, in *GetFaqRequest, opts ...grpc.CallOption) (*v1alpha1_types.Faq, error) {
	out := new(v1alpha1_types.Faq)
	err := grpc.Invoke(ctx, "/v1alpha1.services.FaqService/GetFaq", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FaqService service

type FaqServiceServer interface {
	GetFaq(context.Context, *GetFaqRequest) (*v1alpha1_types.Faq, error)
}

func RegisterFaqServiceServer(s *grpc.Server, srv FaqServiceServer) {
	s.RegisterService(&_FaqService_serviceDesc, srv)
}

func _FaqService_GetFaq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFaqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaqServiceServer).GetFaq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.services.FaqService/GetFaq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaqServiceServer).GetFaq(ctx, req.(*GetFaqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FaqService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.services.FaqService",
	HandlerType: (*FaqServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFaq",
			Handler:    _FaqService_GetFaq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha1/services/faq.proto",
}

func init() { proto.RegisterFile("v1alpha1/services/faq.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x33, 0x4c, 0xcc,
	0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x4b,
	0x2c, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x49, 0xea, 0xc1, 0x24, 0xa5, 0x24,
	0xe0, 0xea, 0x4b, 0x2a, 0x0b, 0x90, 0x15, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43,
	0x64, 0x95, 0xd4, 0xb8, 0x78, 0xdd, 0x53, 0x4b, 0xdc, 0x12, 0x0b, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x44, 0xb9, 0xd8, 0xd2, 0x12, 0x0b, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x58, 0xd3, 0x12, 0x0b, 0x3d, 0x53, 0x8c, 0x72, 0xb9, 0xb8, 0xdc, 0x12, 0x0b,
	0x83, 0x21, 0xd6, 0x09, 0xc5, 0x73, 0xb1, 0x41, 0x74, 0x09, 0x29, 0xe8, 0x61, 0xb8, 0x45, 0x0f,
	0xc5, 0x40, 0x29, 0x61, 0x84, 0x0a, 0xb0, 0xd3, 0xf4, 0xdc, 0x12, 0x0b, 0x95, 0xe4, 0x9a, 0x2e,
	0x3f, 0x99, 0xcc, 0x24, 0x21, 0x24, 0xa6, 0x0f, 0x77, 0x77, 0x5a, 0x62, 0xa1, 0x7e, 0x35, 0xc4,
	0xea, 0x5a, 0x27, 0xbb, 0x28, 0x9b, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xc4, 0xbc, 0x94, 0xa2, 0xd4, 0xf2, 0x8c, 0xfc, 0xf2, 0x94, 0xd4, 0x3c, 0x10, 0x3f, 0xb7,
	0x24, 0x57, 0x37, 0xb1, 0x20, 0xb3, 0x58, 0xbf, 0x20, 0x3b, 0x5d, 0x3f, 0x3d, 0x5f, 0x1f, 0x23,
	0xac, 0x92, 0xd8, 0xc0, 0xbe, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x67, 0xfd, 0x85,
	0x47, 0x01, 0x00, 0x00,
}