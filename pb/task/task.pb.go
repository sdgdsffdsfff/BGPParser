// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package task

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type FilePath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePath) Reset()         { *m = FilePath{} }
func (m *FilePath) String() string { return proto.CompactTextString(m) }
func (*FilePath) ProtoMessage()    {}
func (*FilePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *FilePath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePath.Unmarshal(m, b)
}
func (m *FilePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePath.Marshal(b, m, deterministic)
}
func (m *FilePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePath.Merge(m, src)
}
func (m *FilePath) XXX_Size() int {
	return xxx_messageInfo_FilePath.Size(m)
}
func (m *FilePath) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePath.DiscardUnknown(m)
}

var xxx_messageInfo_FilePath proto.InternalMessageInfo

func (m *FilePath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type TaskReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskReply) Reset()         { *m = TaskReply{} }
func (m *TaskReply) String() string { return proto.CompactTextString(m) }
func (*TaskReply) ProtoMessage()    {}
func (*TaskReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskReply.Unmarshal(m, b)
}
func (m *TaskReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskReply.Marshal(b, m, deterministic)
}
func (m *TaskReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReply.Merge(m, src)
}
func (m *TaskReply) XXX_Size() int {
	return xxx_messageInfo_TaskReply.Size(m)
}
func (m *TaskReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReply.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReply proto.InternalMessageInfo

func (m *TaskReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type IPAddr struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPAddr) Reset()         { *m = IPAddr{} }
func (m *IPAddr) String() string { return proto.CompactTextString(m) }
func (*IPAddr) ProtoMessage()    {}
func (*IPAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *IPAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddr.Unmarshal(m, b)
}
func (m *IPAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddr.Marshal(b, m, deterministic)
}
func (m *IPAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddr.Merge(m, src)
}
func (m *IPAddr) XXX_Size() int {
	return xxx_messageInfo_IPAddr.Size(m)
}
func (m *IPAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddr.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddr proto.InternalMessageInfo

func (m *IPAddr) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type SearchReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReply) Reset()         { *m = SearchReply{} }
func (m *SearchReply) String() string { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()    {}
func (*SearchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *SearchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReply.Unmarshal(m, b)
}
func (m *SearchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReply.Marshal(b, m, deterministic)
}
func (m *SearchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReply.Merge(m, src)
}
func (m *SearchReply) XXX_Size() int {
	return xxx_messageInfo_SearchReply.Size(m)
}
func (m *SearchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReply proto.InternalMessageInfo

func (m *SearchReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type SyncReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncReply) Reset()         { *m = SyncReply{} }
func (m *SyncReply) String() string { return proto.CompactTextString(m) }
func (*SyncReply) ProtoMessage()    {}
func (*SyncReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{4}
}

func (m *SyncReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncReply.Unmarshal(m, b)
}
func (m *SyncReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncReply.Marshal(b, m, deterministic)
}
func (m *SyncReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncReply.Merge(m, src)
}
func (m *SyncReply) XXX_Size() int {
	return xxx_messageInfo_SyncReply.Size(m)
}
func (m *SyncReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncReply.DiscardUnknown(m)
}

var xxx_messageInfo_SyncReply proto.InternalMessageInfo

func (m *SyncReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *SyncReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*FilePath)(nil), "task.FilePath")
	proto.RegisterType((*TaskReply)(nil), "task.TaskReply")
	proto.RegisterType((*IPAddr)(nil), "task.IPAddr")
	proto.RegisterType((*SearchReply)(nil), "task.SearchReply")
	proto.RegisterType((*SyncReply)(nil), "task.SyncReply")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0xd2, 0x12, 0x9b, 0xa9, 0x54, 0x9d, 0x83, 0x84, 0x1e, 0x44, 0x16, 0x0a, 0x5e, 0xac,
	0x52, 0xf5, 0x03, 0xd2, 0x83, 0x25, 0xe0, 0x61, 0x49, 0xfa, 0x03, 0x63, 0x33, 0x98, 0xd0, 0x68,
	0xc2, 0xee, 0xaa, 0xf4, 0x33, 0xfd, 0x23, 0x49, 0x76, 0x2b, 0x9e, 0x4a, 0x6e, 0xef, 0xed, 0xbe,
	0x37, 0xcc, 0x7b, 0x03, 0x60, 0x48, 0xef, 0x16, 0x8d, 0xaa, 0x4d, 0x8d, 0xa3, 0x16, 0x8b, 0x2b,
	0x18, 0x3f, 0x97, 0x15, 0x4b, 0x32, 0x05, 0x22, 0x8c, 0x1a, 0x32, 0x45, 0xe4, 0x5d, 0x7b, 0x37,
	0x61, 0xda, 0x61, 0x31, 0x87, 0x70, 0x43, 0x7a, 0x97, 0x72, 0x53, 0xed, 0x31, 0x82, 0x93, 0x77,
	0xd6, 0x9a, 0xde, 0xd8, 0x69, 0x0e, 0x54, 0x44, 0x10, 0x24, 0x32, 0xce, 0x73, 0x85, 0x53, 0xf0,
	0xcb, 0xc6, 0x7d, 0xfb, 0x65, 0x23, 0xe6, 0x30, 0xc9, 0x98, 0xd4, 0xb6, 0xb0, 0x23, 0x2e, 0x21,
	0x50, 0xac, 0x3f, 0x2b, 0xe3, 0x24, 0x8e, 0x89, 0x27, 0x08, 0xb3, 0xfd, 0xc7, 0xf6, 0xa8, 0x08,
	0xcf, 0x61, 0xc8, 0x4a, 0x45, 0x7e, 0xf7, 0xd8, 0xc2, 0xe5, 0x8f, 0x07, 0xe1, 0x6a, 0x2d, 0xdb,
	0x15, 0x59, 0xe1, 0x3d, 0x4c, 0xe2, 0x3c, 0x4f, 0xe9, 0x5b, 0x92, 0xd2, 0x8c, 0xd3, 0x45, 0x17,
	0xf7, 0x90, 0x6f, 0x76, 0x66, 0xf9, 0x5f, 0x1e, 0x31, 0x70, 0x8e, 0xd5, 0x5a, 0xf6, 0x76, 0xdc,
	0x01, 0xbc, 0xd4, 0x94, 0x27, 0x72, 0xa3, 0xb8, 0xaf, 0x21, 0xa3, 0x2f, 0xee, 0x6d, 0x58, 0x3e,
	0xc2, 0x30, 0x96, 0x09, 0xde, 0xc2, 0xd8, 0x16, 0x97, 0x48, 0x3c, 0xb5, 0x2a, 0x5b, 0xf1, 0xec,
	0xc2, 0xb2, 0x7f, 0xb5, 0x8a, 0xc1, 0x6b, 0xd0, 0x5d, 0xf5, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x81, 0xa4, 0xcb, 0xaa, 0xe3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BGPTaskerClient is the client API for BGPTasker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BGPTaskerClient interface {
	AddRawParse(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error)
	AddBGPParse(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error)
	LoadIPTree(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error)
	SaveIPTree(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error)
}

type bGPTaskerClient struct {
	cc grpc.ClientConnInterface
}

func NewBGPTaskerClient(cc grpc.ClientConnInterface) BGPTaskerClient {
	return &bGPTaskerClient{cc}
}

func (c *bGPTaskerClient) AddRawParse(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.BGPTasker/AddRawParse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGPTaskerClient) AddBGPParse(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.BGPTasker/AddBGPParse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGPTaskerClient) LoadIPTree(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.BGPTasker/LoadIPTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGPTaskerClient) SaveIPTree(ctx context.Context, in *FilePath, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.BGPTasker/SaveIPTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BGPTaskerServer is the server API for BGPTasker service.
type BGPTaskerServer interface {
	AddRawParse(context.Context, *FilePath) (*TaskReply, error)
	AddBGPParse(context.Context, *FilePath) (*TaskReply, error)
	LoadIPTree(context.Context, *FilePath) (*TaskReply, error)
	SaveIPTree(context.Context, *FilePath) (*TaskReply, error)
}

// UnimplementedBGPTaskerServer can be embedded to have forward compatible implementations.
type UnimplementedBGPTaskerServer struct {
}

func (*UnimplementedBGPTaskerServer) AddRawParse(ctx context.Context, req *FilePath) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRawParse not implemented")
}
func (*UnimplementedBGPTaskerServer) AddBGPParse(ctx context.Context, req *FilePath) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBGPParse not implemented")
}
func (*UnimplementedBGPTaskerServer) LoadIPTree(ctx context.Context, req *FilePath) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadIPTree not implemented")
}
func (*UnimplementedBGPTaskerServer) SaveIPTree(ctx context.Context, req *FilePath) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveIPTree not implemented")
}

func RegisterBGPTaskerServer(s *grpc.Server, srv BGPTaskerServer) {
	s.RegisterService(&_BGPTasker_serviceDesc, srv)
}

func _BGPTasker_AddRawParse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPTaskerServer).AddRawParse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.BGPTasker/AddRawParse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPTaskerServer).AddRawParse(ctx, req.(*FilePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGPTasker_AddBGPParse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPTaskerServer).AddBGPParse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.BGPTasker/AddBGPParse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPTaskerServer).AddBGPParse(ctx, req.(*FilePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGPTasker_LoadIPTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPTaskerServer).LoadIPTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.BGPTasker/LoadIPTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPTaskerServer).LoadIPTree(ctx, req.(*FilePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGPTasker_SaveIPTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPTaskerServer).SaveIPTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.BGPTasker/SaveIPTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPTaskerServer).SaveIPTree(ctx, req.(*FilePath))
	}
	return interceptor(ctx, in, info, handler)
}

var _BGPTasker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.BGPTasker",
	HandlerType: (*BGPTaskerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRawParse",
			Handler:    _BGPTasker_AddRawParse_Handler,
		},
		{
			MethodName: "AddBGPParse",
			Handler:    _BGPTasker_AddBGPParse_Handler,
		},
		{
			MethodName: "LoadIPTree",
			Handler:    _BGPTasker_LoadIPTree_Handler,
		},
		{
			MethodName: "SaveIPTree",
			Handler:    _BGPTasker_SaveIPTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	SearchIP(ctx context.Context, in *IPAddr, opts ...grpc.CallOption) (*SearchReply, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) SearchIP(ctx context.Context, in *IPAddr, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/task.API/SearchIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	SearchIP(context.Context, *IPAddr) (*SearchReply, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) SearchIP(ctx context.Context, req *IPAddr) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIP not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_SearchIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SearchIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.API/SearchIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SearchIP(ctx, req.(*IPAddr))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchIP",
			Handler:    _API_SearchIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}
