// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/version/v1/version.proto

package version // import "github.com/ehazlett/stellar/api/services/version/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Revision             string   `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	ContainerdVersion    string   `protobuf:"bytes,3,opt,name=containerd_version,json=containerdVersion,proto3" json:"containerd_version,omitempty"`
	ContainerdRevision   string   `protobuf:"bytes,4,opt,name=containerd_revision,json=containerdRevision,proto3" json:"containerd_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_version_5f6fe1fb74a4ce9e, []int{0}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (dst *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(dst, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *VersionResponse) GetContainerdVersion() string {
	if m != nil {
		return m.ContainerdVersion
	}
	return ""
}

func (m *VersionResponse) GetContainerdRevision() string {
	if m != nil {
		return m.ContainerdRevision
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "stellar.services.version.v1.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VersionClient is the client API for Version service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionClient interface {
	Version(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type versionClient struct {
	cc *grpc.ClientConn
}

func NewVersionClient(cc *grpc.ClientConn) VersionClient {
	return &versionClient{cc}
}

func (c *versionClient) Version(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.version.v1.Version/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServer is the server API for Version service.
type VersionServer interface {
	Version(context.Context, *types.Empty) (*VersionResponse, error)
}

func RegisterVersionServer(s *grpc.Server, srv VersionServer) {
	s.RegisterService(&_Version_serviceDesc, srv)
}

func _Version_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.version.v1.Version/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).Version(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Version_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.version.v1.Version",
	HandlerType: (*VersionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Version_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/version/v1/version.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/version/v1/version.proto", fileDescriptor_version_5f6fe1fb74a4ce9e)
}

var fileDescriptor_version_5f6fe1fb74a4ce9e = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x8a, 0xd5, 0xbd, 0x88, 0xab, 0x48, 0x49, 0x2f, 0xe2, 0xc9, 0x83, 0xce, 0x50,
	0x3d, 0x16, 0x0f, 0x0a, 0x9e, 0x85, 0x1c, 0x44, 0x7a, 0x91, 0x24, 0x8e, 0xdb, 0x85, 0x34, 0x13,
	0x76, 0xb7, 0x0b, 0xfa, 0x87, 0xfc, 0x9b, 0xd2, 0xcd, 0x6e, 0x1b, 0x3c, 0x78, 0xe8, 0xed, 0x4d,
	0xde, 0x37, 0x6f, 0x78, 0x59, 0xf1, 0xa8, 0xb4, 0x5b, 0xac, 0x2a, 0xa8, 0x79, 0x89, 0xb4, 0x28,
	0xbf, 0x1b, 0x72, 0x0e, 0xad, 0xa3, 0xa6, 0x29, 0x0d, 0x96, 0x9d, 0x46, 0x4b, 0xc6, 0xeb, 0x9a,
	0x2c, 0x7a, 0x32, 0x56, 0x73, 0x8b, 0x7e, 0x9a, 0x24, 0x74, 0x86, 0x1d, 0xcb, 0x49, 0xc4, 0x21,
	0xa1, 0x90, 0x7c, 0x3f, 0xcd, 0x27, 0x8a, 0x59, 0x35, 0x84, 0x01, 0xad, 0x56, 0x9f, 0x48, 0xcb,
	0xce, 0x7d, 0xf5, 0x9b, 0xf9, 0xb9, 0x62, 0xc5, 0x41, 0xe2, 0x5a, 0xf5, 0x5f, 0xaf, 0x7e, 0x32,
	0x71, 0xf2, 0xda, 0x27, 0x14, 0x64, 0x3b, 0x6e, 0x2d, 0xc9, 0xb1, 0x18, 0xc5, 0xd0, 0x71, 0x76,
	0x99, 0x5d, 0x1f, 0x17, 0x69, 0x94, 0xb9, 0x38, 0x32, 0xe4, 0x75, 0xb0, 0xf6, 0x82, 0xb5, 0x99,
	0xe5, 0xad, 0x90, 0x35, 0xb7, 0xae, 0xd4, 0x2d, 0x99, 0x8f, 0xf7, 0x14, 0xb0, 0x1f, 0xa8, 0xd3,
	0xad, 0x13, 0x8f, 0x49, 0x14, 0x67, 0x03, 0x7c, 0x93, 0x7a, 0x10, 0xf8, 0x41, 0x52, 0x11, 0x9d,
	0xbb, 0xb9, 0x18, 0xa5, 0xdd, 0x97, 0xad, 0xbc, 0x80, 0xbe, 0x33, 0xa4, 0xce, 0xf0, 0xbc, 0xee,
	0x9c, 0xdf, 0xc0, 0x3f, 0x3f, 0x0a, 0xfe, 0x34, 0x7e, 0x7a, 0x98, 0xcf, 0x76, 0x78, 0x9a, 0x59,
	0x94, 0x6f, 0x59, 0x75, 0x18, 0xce, 0xdf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd0, 0xea,
	0x3f, 0xe2, 0x01, 0x00, 0x00,
}
