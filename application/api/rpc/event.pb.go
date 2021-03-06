// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/api/rpc/event.proto

package rpc

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

type Goods struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Img                  string   `protobuf:"bytes,3,opt,name=img,proto3" json:"img,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	EventPrice           string   `protobuf:"bytes,5,opt,name=event_price,json=eventPrice,proto3" json:"event_price,omitempty"`
	EventStock           int32    `protobuf:"varint,6,opt,name=event_stock,json=eventStock,proto3" json:"event_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Goods) Reset()         { *m = Goods{} }
func (m *Goods) String() string { return proto.CompactTextString(m) }
func (*Goods) ProtoMessage()    {}
func (*Goods) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae9c5481ac0209, []int{0}
}

func (m *Goods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Goods.Unmarshal(m, b)
}
func (m *Goods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Goods.Marshal(b, m, deterministic)
}
func (m *Goods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Goods.Merge(m, src)
}
func (m *Goods) XXX_Size() int {
	return xxx_messageInfo_Goods.Size(m)
}
func (m *Goods) XXX_DiscardUnknown() {
	xxx_messageInfo_Goods.DiscardUnknown(m)
}

var xxx_messageInfo_Goods proto.InternalMessageInfo

func (m *Goods) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Goods) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Goods) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *Goods) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Goods) GetEventPrice() string {
	if m != nil {
		return m.EventPrice
	}
	return ""
}

func (m *Goods) GetEventStock() int32 {
	if m != nil {
		return m.EventStock
	}
	return 0
}

type Event struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId              int32    `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	GoodsList            []*Goods `protobuf:"bytes,6,rep,name=goods_list,json=goodsList,proto3" json:"goods_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae9c5481ac0209, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetTopicId() int32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *Event) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Event) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Event) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Event) GetGoodsList() []*Goods {
	if m != nil {
		return m.GoodsList
	}
	return nil
}

type Topic struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Banner               string   `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner,omitempty"`
	StartTime            int64    `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae9c5481ac0209, []int{2}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Topic) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Topic) GetBanner() string {
	if m != nil {
		return m.Banner
	}
	return ""
}

func (m *Topic) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Topic) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae9c5481ac0209, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Goods)(nil), "rpc.Goods")
	proto.RegisterType((*Event)(nil), "rpc.Event")
	proto.RegisterType((*Topic)(nil), "rpc.Topic")
	proto.RegisterType((*Response)(nil), "rpc.Response")
}

func init() { proto.RegisterFile("application/api/rpc/event.proto", fileDescriptor_83ae9c5481ac0209) }

var fileDescriptor_83ae9c5481ac0209 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5f, 0xab, 0xd3, 0x30,
	0x14, 0xa7, 0xeb, 0xd2, 0xbb, 0x9d, 0xaa, 0x48, 0x18, 0x52, 0x05, 0xb9, 0xa3, 0x4f, 0xbd, 0x2f,
	0x9b, 0x5c, 0x3f, 0x82, 0x88, 0x08, 0x82, 0x97, 0x78, 0xdf, 0x4b, 0x97, 0x64, 0xe3, 0x60, 0x9b,
	0x84, 0x26, 0xf8, 0x45, 0xc4, 0x8f, 0xe1, 0x9b, 0x1f, 0x50, 0x72, 0x52, 0x86, 0x75, 0x43, 0xdf,
	0xce, 0xef, 0xcf, 0x4e, 0x7e, 0xe7, 0xb7, 0xc2, 0x6d, 0xe7, 0x5c, 0x8f, 0xb2, 0x0b, 0x68, 0xcd,
	0xbe, 0x73, 0xb8, 0x1f, 0x9d, 0xdc, 0xeb, 0x6f, 0xda, 0x84, 0x9d, 0x1b, 0x6d, 0xb0, 0x3c, 0x1f,
	0x9d, 0xac, 0x7f, 0x64, 0xc0, 0x3e, 0x58, 0xab, 0x3c, 0x7f, 0x06, 0x0b, 0x54, 0x55, 0xb6, 0xcd,
	0x1a, 0x26, 0x16, 0xa8, 0x38, 0x87, 0xa5, 0xd2, 0x5e, 0x56, 0x8b, 0x6d, 0xd6, 0xac, 0x05, 0xcd,
	0xfc, 0x39, 0xe4, 0x38, 0x9c, 0xaa, 0x9c, 0xa8, 0x38, 0xf2, 0x0d, 0x30, 0x37, 0xa2, 0xd4, 0xd5,
	0x92, 0xb8, 0x04, 0xf8, 0x2d, 0x94, 0xf4, 0x52, 0x9b, 0x34, 0x46, 0x1a, 0x10, 0xf5, 0x30, 0x37,
	0xf8, 0x60, 0xe5, 0xd7, 0xaa, 0xa0, 0x57, 0x93, 0xe1, 0x4b, 0x64, 0xea, 0x9f, 0x19, 0xb0, 0xf7,
	0x11, 0x5e, 0xe4, 0x7a, 0x09, 0xab, 0x60, 0x1d, 0xca, 0x16, 0x15, 0x65, 0x63, 0xe2, 0x86, 0xf0,
	0x47, 0xc5, 0x5f, 0x03, 0xf8, 0xd0, 0x8d, 0xa1, 0x0d, 0x38, 0x68, 0x4a, 0x99, 0x8b, 0x35, 0x31,
	0x8f, 0x38, 0xe8, 0xf8, 0x4b, 0x6d, 0x54, 0x12, 0x97, 0x24, 0xde, 0x68, 0xa3, 0x48, 0xda, 0x00,
	0xeb, 0x71, 0xc0, 0x40, 0x51, 0x99, 0x48, 0x80, 0xdf, 0x01, 0x9c, 0x62, 0x37, 0x6d, 0x8f, 0x3e,
	0x54, 0xc5, 0x36, 0x6f, 0xca, 0x7b, 0xd8, 0x8d, 0x4e, 0xee, 0xa8, 0x32, 0xb1, 0x26, 0xf5, 0x13,
	0xfa, 0x50, 0x7f, 0xcf, 0x80, 0x3d, 0xc6, 0x18, 0x17, 0x79, 0x37, 0xc0, 0x02, 0x86, 0x5e, 0x4f,
	0x45, 0x26, 0x70, 0x6e, 0x37, 0xff, 0xa3, 0xdd, 0x17, 0x50, 0x1c, 0x3a, 0x63, 0xf4, 0x38, 0x95,
	0x39, 0xa1, 0xbf, 0xce, 0x62, 0xff, 0x3a, 0xab, 0x98, 0x9d, 0x55, 0xbf, 0x81, 0x95, 0xd0, 0xde,
	0x59, 0xe3, 0xe9, 0x45, 0x69, 0x95, 0x9e, 0x92, 0xd1, 0x1c, 0xff, 0xcf, 0xc1, 0x9f, 0xa6, 0x64,
	0x71, 0xbc, 0xff, 0x95, 0xc1, 0x8a, 0x7a, 0x17, 0x0f, 0xef, 0x78, 0x03, 0x25, 0xcd, 0x9f, 0x4d,
	0x8f, 0x46, 0xf3, 0x74, 0x3a, 0x31, 0xaf, 0x9e, 0xd2, 0x7c, 0x5e, 0x7e, 0x07, 0x4f, 0x92, 0xf3,
	0x78, 0xfc, 0x9f, 0xb5, 0x81, 0x92, 0x8a, 0x9a, 0x2d, 0x25, 0xe6, 0xca, 0xd2, 0xe4, 0x9c, 0x2d,
	0xbd, 0x66, 0x3d, 0x14, 0xf4, 0x49, 0xbf, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x20, 0x74,
	0x1f, 0xf5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventRPCClient is the client API for EventRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventRPCClient interface {
	EventOnline(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
	EventOffline(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
	TopicOnline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error)
	TopicOffline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error)
}

type eventRPCClient struct {
	cc *grpc.ClientConn
}

func NewEventRPCClient(cc *grpc.ClientConn) EventRPCClient {
	return &eventRPCClient{cc}
}

func (c *eventRPCClient) EventOnline(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.EventRPC/EventOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRPCClient) EventOffline(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.EventRPC/EventOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRPCClient) TopicOnline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.EventRPC/TopicOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRPCClient) TopicOffline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.EventRPC/TopicOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventRPCServer is the server API for EventRPC service.
type EventRPCServer interface {
	EventOnline(context.Context, *Event) (*Response, error)
	EventOffline(context.Context, *Event) (*Response, error)
	TopicOnline(context.Context, *Topic) (*Response, error)
	TopicOffline(context.Context, *Topic) (*Response, error)
}

// UnimplementedEventRPCServer can be embedded to have forward compatible implementations.
type UnimplementedEventRPCServer struct {
}

func (*UnimplementedEventRPCServer) EventOnline(ctx context.Context, req *Event) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventOnline not implemented")
}
func (*UnimplementedEventRPCServer) EventOffline(ctx context.Context, req *Event) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventOffline not implemented")
}
func (*UnimplementedEventRPCServer) TopicOnline(ctx context.Context, req *Topic) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicOnline not implemented")
}
func (*UnimplementedEventRPCServer) TopicOffline(ctx context.Context, req *Topic) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicOffline not implemented")
}

func RegisterEventRPCServer(s *grpc.Server, srv EventRPCServer) {
	s.RegisterService(&_EventRPC_serviceDesc, srv)
}

func _EventRPC_EventOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRPCServer).EventOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.EventRPC/EventOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRPCServer).EventOnline(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRPC_EventOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRPCServer).EventOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.EventRPC/EventOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRPCServer).EventOffline(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRPC_TopicOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRPCServer).TopicOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.EventRPC/TopicOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRPCServer).TopicOnline(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRPC_TopicOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRPCServer).TopicOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.EventRPC/TopicOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRPCServer).TopicOffline(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.EventRPC",
	HandlerType: (*EventRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventOnline",
			Handler:    _EventRPC_EventOnline_Handler,
		},
		{
			MethodName: "EventOffline",
			Handler:    _EventRPC_EventOffline_Handler,
		},
		{
			MethodName: "TopicOnline",
			Handler:    _EventRPC_TopicOnline_Handler,
		},
		{
			MethodName: "TopicOffline",
			Handler:    _EventRPC_TopicOffline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/api/rpc/event.proto",
}
