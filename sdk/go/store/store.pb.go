// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/store.proto

package go_micro_store

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

type Record struct {
	// key of the record
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value in the record
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp in unix seconds
	Expiry               int64    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type ReadRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type WriteRequest struct {
	Records              []*Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{3}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{4}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	// optional key
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{7}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ListResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{8}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "go.micro.store.Record")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.store.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.store.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "go.micro.store.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "go.micro.store.WriteResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.store.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.store.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.store.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.store.ListResponse")
}

func init() {
	proto.RegisterFile("store/store.proto", fileDescriptor_8549980b097f750b)
}

var fileDescriptor_8549980b097f750b = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xeb, 0xa6, 0xcd, 0xaf, 0x4e, 0x2f, 0x7f, 0x19, 0xa1, 0x2a, 0x0a, 0xb7, 0x10, 0x36,
	0x59, 0x85, 0xaa, 0xbc, 0x00, 0x12, 0x17, 0x81, 0xc4, 0xca, 0x2c, 0x58, 0x97, 0x76, 0x84, 0xa2,
	0x16, 0x1c, 0xec, 0x14, 0x91, 0x17, 0xe2, 0x39, 0x91, 0xed, 0x04, 0x52, 0xd2, 0x4a, 0x88, 0x4d,
	0x64, 0x67, 0xce, 0x7c, 0x73, 0x32, 0x27, 0xb0, 0xa3, 0x32, 0x21, 0xe9, 0xd4, 0x3c, 0xe3, 0x54,
	0x8a, 0x4c, 0xe0, 0xe0, 0x49, 0xc4, 0xcf, 0xc9, 0x4c, 0x8a, 0xd8, 0xbc, 0x0d, 0x6f, 0xc0, 0xe5,
	0x34, 0x13, 0x72, 0x8e, 0x43, 0x70, 0x16, 0x94, 0x7b, 0x2c, 0x60, 0x51, 0x87, 0xeb, 0x23, 0xee,
	0x42, 0xfb, 0x6d, 0xba, 0x5c, 0x91, 0xd7, 0x0c, 0x58, 0xd4, 0xe3, 0xf6, 0x82, 0x23, 0x70, 0xe9,
	0x3d, 0x4d, 0x64, 0xee, 0x39, 0x01, 0x8b, 0x1c, 0x5e, 0xdc, 0xc2, 0x63, 0xe8, 0x72, 0x9a, 0xce,
	0x39, 0xbd, 0xae, 0x48, 0x65, 0x88, 0xd0, 0x5a, 0x50, 0xae, 0x3c, 0x16, 0x38, 0x51, 0x87, 0x9b,
	0x73, 0x78, 0x0e, 0x3d, 0x2b, 0x51, 0xa9, 0x78, 0x51, 0x84, 0x63, 0xf8, 0x27, 0xcd, 0x70, 0x2b,
	0xeb, 0x4e, 0x46, 0xf1, 0xba, 0xbd, 0xd8, 0x7a, 0xe3, 0xa5, 0x4c, 0x13, 0x1e, 0x64, 0x92, 0x51,
	0x39, 0xa5, 0x42, 0x68, 0xfe, 0x8e, 0xf0, 0x1f, 0xfa, 0x05, 0xc1, 0x9a, 0x08, 0x4f, 0xa0, 0x7f,
	0x49, 0x4b, 0xfa, 0x66, 0x6e, 0x72, 0x3e, 0x84, 0x41, 0x29, 0x2a, 0xda, 0x8e, 0xa0, 0x7b, 0x97,
	0xa8, 0xac, 0x6c, 0xaa, 0x6d, 0x4f, 0x5b, 0xb5, 0x82, 0xbf, 0x7e, 0xec, 0xe4, 0xa3, 0x09, 0xed,
	0x7b, 0x5d, 0xc1, 0x2b, 0x68, 0x69, 0x16, 0xee, 0xfd, 0x6c, 0xa9, 0x58, 0xf0, 0xf7, 0x37, 0x17,
	0x0b, 0xbf, 0x8d, 0x31, 0xc3, 0x0b, 0x68, 0xe9, 0xfd, 0xd7, 0x31, 0x95, 0xe0, 0xea, 0x98, 0x6a,
	0x64, 0x61, 0x03, 0xaf, 0xa1, 0x6d, 0x16, 0x88, 0x35, 0x61, 0x35, 0x19, 0xff, 0x60, 0x4b, 0xf5,
	0x8b, 0x73, 0x0b, 0xae, 0x5d, 0x29, 0xd6, 0xa4, 0x6b, 0x79, 0xf8, 0x87, 0xdb, 0xca, 0x25, 0xea,
	0xd1, 0x35, 0xff, 0xf6, 0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x97, 0xd0, 0x5f, 0xf0,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Store_ListClient, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Store_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Store_serviceDesc.Streams[0], "/go.micro.store.Store/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Store_ListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type storeListClient struct {
	grpc.ClientStream
}

func (x *storeListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/go.micro.store.Store/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/go.micro.store.Store/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/go.micro.store.Store/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	List(*ListRequest, Store_ListServer) error
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) List(req *ListRequest, srv Store_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStoreServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStoreServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedStoreServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreServer).List(m, &storeListServer{stream})
}

type Store_ListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type storeListServer struct {
	grpc.ServerStream
}

func (x *storeListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Store_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.store.Store/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.store.Store/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.store.Store/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.store.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Store_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Store_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Store_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Store_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "store/store.proto",
}