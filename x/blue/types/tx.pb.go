// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blue/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreatePool struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Question    string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	AnswerRange uint64 `protobuf:"varint,4,opt,name=answerRange,proto3" json:"answerRange,omitempty"`
}

func (m *MsgCreatePool) Reset()         { *m = MsgCreatePool{} }
func (m *MsgCreatePool) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePool) ProtoMessage()    {}
func (*MsgCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_44496d1efbf19795, []int{0}
}
func (m *MsgCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePool.Merge(m, src)
}
func (m *MsgCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePool proto.InternalMessageInfo

func (m *MsgCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePool) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgCreatePool) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *MsgCreatePool) GetAnswerRange() uint64 {
	if m != nil {
		return m.AnswerRange
	}
	return 0
}

type MsgCreatePoolResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreatePoolResponse) Reset()         { *m = MsgCreatePoolResponse{} }
func (m *MsgCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoolResponse) ProtoMessage()    {}
func (*MsgCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44496d1efbf19795, []int{1}
}
func (m *MsgCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoolResponse.Merge(m, src)
}
func (m *MsgCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoolResponse proto.InternalMessageInfo

func (m *MsgCreatePoolResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgVote struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PoolId  string `protobuf:"bytes,2,opt,name=poolId,proto3" json:"poolId,omitempty"`
	Answer  string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (m *MsgVote) Reset()         { *m = MsgVote{} }
func (m *MsgVote) String() string { return proto.CompactTextString(m) }
func (*MsgVote) ProtoMessage()    {}
func (*MsgVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_44496d1efbf19795, []int{2}
}
func (m *MsgVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVote.Merge(m, src)
}
func (m *MsgVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVote proto.InternalMessageInfo

func (m *MsgVote) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgVote) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *MsgVote) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type MsgVoteResponse struct {
	Pool *Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (m *MsgVoteResponse) Reset()         { *m = MsgVoteResponse{} }
func (m *MsgVoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVoteResponse) ProtoMessage()    {}
func (*MsgVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44496d1efbf19795, []int{3}
}
func (m *MsgVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteResponse.Merge(m, src)
}
func (m *MsgVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteResponse proto.InternalMessageInfo

func (m *MsgVoteResponse) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreatePool)(nil), "blue.blue.MsgCreatePool")
	proto.RegisterType((*MsgCreatePoolResponse)(nil), "blue.blue.MsgCreatePoolResponse")
	proto.RegisterType((*MsgVote)(nil), "blue.blue.MsgVote")
	proto.RegisterType((*MsgVoteResponse)(nil), "blue.blue.MsgVoteResponse")
}

func init() { proto.RegisterFile("blue/tx.proto", fileDescriptor_44496d1efbf19795) }

var fileDescriptor_44496d1efbf19795 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xcd, 0xa6, 0xb1, 0xb5, 0x53, 0x6a, 0x70, 0x51, 0x09, 0x39, 0x2c, 0x21, 0x1e, 0x2c, 0x08,
	0x29, 0x54, 0xe8, 0x07, 0xe8, 0x45, 0x0f, 0x01, 0x89, 0xe0, 0xc1, 0x5b, 0x6a, 0x96, 0x10, 0x08,
	0xd9, 0x98, 0xdd, 0x62, 0x3d, 0x78, 0xf6, 0xea, 0x67, 0x79, 0xec, 0xd1, 0xa3, 0x24, 0x3f, 0x22,
	0xbb, 0x49, 0x43, 0x82, 0xd8, 0x4b, 0xe0, 0xbd, 0xb7, 0x33, 0xef, 0xcd, 0x4c, 0x60, 0xba, 0x4a,
	0xd7, 0x74, 0x2e, 0x36, 0x5e, 0x5e, 0x30, 0xc1, 0xf0, 0x58, 0x42, 0x4f, 0x7e, 0x6c, 0x53, 0x29,
	0x39, 0x63, 0x69, 0xad, 0xb9, 0xef, 0x30, 0xf5, 0x79, 0x7c, 0x53, 0xd0, 0x50, 0xd0, 0x7b, 0xc6,
	0x52, 0x6c, 0xc1, 0xe8, 0x59, 0x22, 0x56, 0x58, 0xc8, 0x41, 0xb3, 0x71, 0xb0, 0x83, 0xf8, 0x04,
	0x0e, 0x44, 0x22, 0x52, 0x6a, 0xe9, 0x8a, 0xaf, 0x01, 0xb6, 0xe1, 0xf0, 0x65, 0x4d, 0xb9, 0x48,
	0x58, 0x66, 0x0d, 0x94, 0xd0, 0x62, 0xec, 0xc0, 0x24, 0xcc, 0xf8, 0x2b, 0x2d, 0x82, 0x30, 0x8b,
	0xa9, 0x65, 0x38, 0x68, 0x66, 0x04, 0x5d, 0xca, 0xbd, 0x80, 0xd3, 0x9e, 0x7d, 0x40, 0x79, 0xce,
	0x32, 0x4e, 0xf1, 0x11, 0xe8, 0x49, 0xa4, 0x12, 0x18, 0x81, 0x9e, 0x44, 0xee, 0x03, 0x8c, 0x7c,
	0x1e, 0x3f, 0x32, 0x41, 0xf7, 0x24, 0x3c, 0x83, 0xa1, 0x1c, 0xed, 0x2e, 0x6a, 0x22, 0x36, 0x48,
	0xf2, 0xb5, 0x69, 0x93, 0xb0, 0x41, 0xee, 0x12, 0xcc, 0xa6, 0x69, 0xeb, 0x7b, 0x0e, 0x86, 0x2c,
	0x52, 0x9d, 0x27, 0x0b, 0xd3, 0x6b, 0x57, 0xe7, 0xa9, 0x78, 0x4a, 0x5c, 0x7c, 0x20, 0x18, 0xf8,
	0x3c, 0xc6, 0xb7, 0x00, 0xdd, 0xcd, 0x75, 0x1e, 0xf7, 0x86, 0xb2, 0x9d, 0xff, 0x94, 0xd6, 0x76,
	0x09, 0x86, 0x9a, 0x0d, 0xf7, 0x5f, 0x4a, 0xce, 0xb6, 0xff, 0x72, 0xbb, 0xba, 0xeb, 0xcb, 0xaf,
	0x92, 0xa0, 0x6d, 0x49, 0xd0, 0x4f, 0x49, 0xd0, 0x67, 0x45, 0xb4, 0x6d, 0x45, 0xb4, 0xef, 0x8a,
	0x68, 0x4f, 0xc7, 0xea, 0xd2, 0x9b, 0x79, 0xfd, 0x2b, 0xbc, 0xe5, 0x94, 0xaf, 0x86, 0xea, 0xe4,
	0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x59, 0x99, 0x00, 0x1f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/blue.blue.Msg/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/blue.blue.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePool(ctx context.Context, req *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (*UnimplementedMsgServer) Vote(ctx context.Context, req *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blue.blue.Msg/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blue.blue.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blue.blue.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blue/tx.proto",
}

func (m *MsgCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AnswerRange != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AnswerRange))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Answer) > 0 {
		i -= len(m.Answer)
		copy(dAtA[i:], m.Answer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Answer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pool != nil {
		{
			size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AnswerRange != 0 {
		n += 1 + sovTx(uint64(m.AnswerRange))
	}
	return n
}

func (m *MsgCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Answer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pool != nil {
		l = m.Pool.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnswerRange", wireType)
			}
			m.AnswerRange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnswerRange |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreatePoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pool == nil {
				m.Pool = &Pool{}
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
