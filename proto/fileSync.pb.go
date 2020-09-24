// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileSync.proto

package fileSync

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

//message for file sync request
type FileSyncReq struct {
	SubDir               string   `protobuf:"bytes,1,opt,name=subDir,proto3" json:"subDir,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize             int64    `protobuf:"varint,3,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	FileContent          []byte   `protobuf:"bytes,4,opt,name=fileContent,proto3" json:"fileContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSyncReq) Reset()         { *m = FileSyncReq{} }
func (m *FileSyncReq) String() string { return proto.CompactTextString(m) }
func (*FileSyncReq) ProtoMessage()    {}
func (*FileSyncReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f500c2c5ea8ca846, []int{0}
}

func (m *FileSyncReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSyncReq.Unmarshal(m, b)
}
func (m *FileSyncReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSyncReq.Marshal(b, m, deterministic)
}
func (m *FileSyncReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSyncReq.Merge(m, src)
}
func (m *FileSyncReq) XXX_Size() int {
	return xxx_messageInfo_FileSyncReq.Size(m)
}
func (m *FileSyncReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSyncReq.DiscardUnknown(m)
}

var xxx_messageInfo_FileSyncReq proto.InternalMessageInfo

func (m *FileSyncReq) GetSubDir() string {
	if m != nil {
		return m.SubDir
	}
	return ""
}

func (m *FileSyncReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileSyncReq) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FileSyncReq) GetFileContent() []byte {
	if m != nil {
		return m.FileContent
	}
	return nil
}

//message for file remove request
type FileRemoveReq struct {
	SubDir               string   `protobuf:"bytes,1,opt,name=subDir,proto3" json:"subDir,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRemoveReq) Reset()         { *m = FileRemoveReq{} }
func (m *FileRemoveReq) String() string { return proto.CompactTextString(m) }
func (*FileRemoveReq) ProtoMessage()    {}
func (*FileRemoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f500c2c5ea8ca846, []int{1}
}

func (m *FileRemoveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRemoveReq.Unmarshal(m, b)
}
func (m *FileRemoveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRemoveReq.Marshal(b, m, deterministic)
}
func (m *FileRemoveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRemoveReq.Merge(m, src)
}
func (m *FileRemoveReq) XXX_Size() int {
	return xxx_messageInfo_FileRemoveReq.Size(m)
}
func (m *FileRemoveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRemoveReq.DiscardUnknown(m)
}

var xxx_messageInfo_FileRemoveReq proto.InternalMessageInfo

func (m *FileRemoveReq) GetSubDir() string {
	if m != nil {
		return m.SubDir
	}
	return ""
}

func (m *FileRemoveReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

//message for sync response
type SyncResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncResp) Reset()         { *m = SyncResp{} }
func (m *SyncResp) String() string { return proto.CompactTextString(m) }
func (*SyncResp) ProtoMessage()    {}
func (*SyncResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f500c2c5ea8ca846, []int{2}
}

func (m *SyncResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResp.Unmarshal(m, b)
}
func (m *SyncResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResp.Marshal(b, m, deterministic)
}
func (m *SyncResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResp.Merge(m, src)
}
func (m *SyncResp) XXX_Size() int {
	return xxx_messageInfo_SyncResp.Size(m)
}
func (m *SyncResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResp.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResp proto.InternalMessageInfo

func (m *SyncResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//message for dir sync request
type DirSyncReq struct {
	SubDir               string   `protobuf:"bytes,1,opt,name=subDir,proto3" json:"subDir,omitempty"`
	NewSubDir            string   `protobuf:"bytes,2,opt,name=newSubDir,proto3" json:"newSubDir,omitempty"`
	IsRemove             bool     `protobuf:"varint,3,opt,name=isRemove,proto3" json:"isRemove,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirSyncReq) Reset()         { *m = DirSyncReq{} }
func (m *DirSyncReq) String() string { return proto.CompactTextString(m) }
func (*DirSyncReq) ProtoMessage()    {}
func (*DirSyncReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f500c2c5ea8ca846, []int{3}
}

func (m *DirSyncReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirSyncReq.Unmarshal(m, b)
}
func (m *DirSyncReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirSyncReq.Marshal(b, m, deterministic)
}
func (m *DirSyncReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirSyncReq.Merge(m, src)
}
func (m *DirSyncReq) XXX_Size() int {
	return xxx_messageInfo_DirSyncReq.Size(m)
}
func (m *DirSyncReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DirSyncReq.DiscardUnknown(m)
}

var xxx_messageInfo_DirSyncReq proto.InternalMessageInfo

func (m *DirSyncReq) GetSubDir() string {
	if m != nil {
		return m.SubDir
	}
	return ""
}

func (m *DirSyncReq) GetNewSubDir() string {
	if m != nil {
		return m.NewSubDir
	}
	return ""
}

func (m *DirSyncReq) GetIsRemove() bool {
	if m != nil {
		return m.IsRemove
	}
	return false
}

func init() {
	proto.RegisterType((*FileSyncReq)(nil), "fileSync.FileSyncReq")
	proto.RegisterType((*FileRemoveReq)(nil), "fileSync.FileRemoveReq")
	proto.RegisterType((*SyncResp)(nil), "fileSync.SyncResp")
	proto.RegisterType((*DirSyncReq)(nil), "fileSync.DirSyncReq")
}

func init() { proto.RegisterFile("fileSync.proto", fileDescriptor_f500c2c5ea8ca846) }

var fileDescriptor_f500c2c5ea8ca846 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xc9, 0xdb, 0x97, 0x76, 0x7a, 0xeb, 0x07, 0x5e, 0xbf, 0x86, 0xe2, 0x22, 0x04, 0x17,
	0x59, 0x65, 0x61, 0x71, 0xe1, 0xd6, 0x16, 0x97, 0x2e, 0x26, 0x7b, 0xc1, 0x0e, 0x57, 0x18, 0x30,
	0x33, 0x35, 0x93, 0x56, 0xea, 0xc6, 0xbf, 0xe5, 0xcf, 0x93, 0x4c, 0x32, 0x33, 0x16, 0x0a, 0x82,
	0xbb, 0x7b, 0xee, 0xe1, 0xdc, 0x9c, 0x27, 0x09, 0x1c, 0xbd, 0xa8, 0x57, 0x2a, 0xb7, 0x5a, 0x16,
	0xab, 0xda, 0x34, 0x06, 0x99, 0xd7, 0xd9, 0x27, 0x4c, 0x1e, 0xfa, 0x59, 0xd0, 0x1b, 0x5e, 0xc0,
	0xd0, 0xae, 0x97, 0x0b, 0x55, 0xf3, 0x24, 0x4d, 0xf2, 0xb1, 0xe8, 0x15, 0x4e, 0xc1, 0x45, 0x1e,
	0x9f, 0x2b, 0xe2, 0xff, 0x9c, 0x13, 0xb4, 0xf7, 0x4a, 0xf5, 0x41, 0x7c, 0x90, 0x26, 0xf9, 0x40,
	0x04, 0x8d, 0x29, 0x4c, 0xda, 0x79, 0x6e, 0x74, 0x43, 0xba, 0xe1, 0xff, 0xd3, 0x24, 0x3f, 0x10,
	0x3f, 0x57, 0xd9, 0x1c, 0x0e, 0xdb, 0x02, 0x82, 0x2a, 0xb3, 0xa1, 0x3f, 0x56, 0xc8, 0xae, 0x81,
	0x75, 0x04, 0x76, 0x85, 0x1c, 0x46, 0x76, 0x2d, 0x25, 0x59, 0xeb, 0x0e, 0x30, 0xe1, 0x65, 0xf6,
	0x04, 0xb0, 0x50, 0xf5, 0x6f, 0xa8, 0x57, 0x30, 0xd6, 0xf4, 0x5e, 0x76, 0x56, 0xf7, 0xa0, 0xb8,
	0x68, 0x5b, 0x28, 0xdb, 0x95, 0x75, 0xb0, 0x4c, 0x04, 0x7d, 0xf3, 0x95, 0xc0, 0xb1, 0x7f, 0x99,
	0x25, 0xd5, 0x1b, 0x25, 0x09, 0xef, 0x00, 0x22, 0x1e, 0x5e, 0x16, 0xe1, 0x43, 0xec, 0x40, 0x4f,
	0x31, 0x1a, 0x01, 0xe4, 0x16, 0x98, 0xbf, 0x86, 0xe7, 0xbb, 0xc1, 0x9e, 0x61, 0x6f, 0x6c, 0x06,
	0xa3, 0x9e, 0x12, 0xcf, 0xa2, 0x1d, 0xc1, 0xf7, 0x85, 0xee, 0x4f, 0xe1, 0x44, 0x9a, 0xaa, 0x68,
	0x94, 0xde, 0x06, 0x77, 0x39, 0x74, 0x3f, 0xcb, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x71,
	0xb2, 0x5c, 0x3e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileSyncServiceClient is the client API for FileSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileSyncServiceClient interface {
	//file remove
	FileRemove(ctx context.Context, in *FileRemoveReq, opts ...grpc.CallOption) (*SyncResp, error)
	//file sync
	FileSync(ctx context.Context, in *FileSyncReq, opts ...grpc.CallOption) (*SyncResp, error)
	//dir sync
	DirSync(ctx context.Context, in *DirSyncReq, opts ...grpc.CallOption) (*SyncResp, error)
}

type fileSyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileSyncServiceClient(cc *grpc.ClientConn) FileSyncServiceClient {
	return &fileSyncServiceClient{cc}
}

func (c *fileSyncServiceClient) FileRemove(ctx context.Context, in *FileRemoveReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/fileSync.FileSyncService/FileRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) FileSync(ctx context.Context, in *FileSyncReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/fileSync.FileSyncService/FileSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) DirSync(ctx context.Context, in *DirSyncReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/fileSync.FileSyncService/DirSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileSyncServiceServer is the server API for FileSyncService service.
type FileSyncServiceServer interface {
	//file remove
	FileRemove(context.Context, *FileRemoveReq) (*SyncResp, error)
	//file sync
	FileSync(context.Context, *FileSyncReq) (*SyncResp, error)
	//dir sync
	DirSync(context.Context, *DirSyncReq) (*SyncResp, error)
}

// UnimplementedFileSyncServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileSyncServiceServer struct {
}

func (*UnimplementedFileSyncServiceServer) FileRemove(ctx context.Context, req *FileRemoveReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileRemove not implemented")
}
func (*UnimplementedFileSyncServiceServer) FileSync(ctx context.Context, req *FileSyncReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileSync not implemented")
}
func (*UnimplementedFileSyncServiceServer) DirSync(ctx context.Context, req *DirSyncReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirSync not implemented")
}

func RegisterFileSyncServiceServer(s *grpc.Server, srv FileSyncServiceServer) {
	s.RegisterService(&_FileSyncService_serviceDesc, srv)
}

func _FileSyncService_FileRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).FileRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSync.FileSyncService/FileRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).FileRemove(ctx, req.(*FileRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_FileSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).FileSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSync.FileSyncService/FileSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).FileSync(ctx, req.(*FileSyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_DirSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirSyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).DirSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSync.FileSyncService/DirSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).DirSync(ctx, req.(*DirSyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileSyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileSync.FileSyncService",
	HandlerType: (*FileSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileRemove",
			Handler:    _FileSyncService_FileRemove_Handler,
		},
		{
			MethodName: "FileSync",
			Handler:    _FileSyncService_FileSync_Handler,
		},
		{
			MethodName: "DirSync",
			Handler:    _FileSyncService_DirSync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fileSync.proto",
}
