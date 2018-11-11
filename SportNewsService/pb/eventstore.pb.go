// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventstore.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	eventstore.proto
	news.proto
	user_data.proto

It has these top-level messages:
	Event
	Response
	EventFilter
	EventResponse
	News
	PostNewsRequest
	PostNewsResponse
	GetNewsRequest
	GetNewsResponse
	GetAllNewsRequest
	GetAllNewsResponse
	UserDataRequest
	UserDataResponse
*/
package pb

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

type Event struct {
	EventId       string `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	EventType     string `protobuf:"bytes,2,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	AggregateId   string `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId" json:"aggregate_id,omitempty"`
	AggregateType string `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType" json:"aggregate_type,omitempty"`
	EventData     string `protobuf:"bytes,5,opt,name=event_data,json=eventData" json:"event_data,omitempty"`
	Channel       string `protobuf:"bytes,6,opt,name=channel" json:"channel,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

func (m *Event) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type Response struct {
	IsSuccess bool   `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventFilter struct {
	EventId     string `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	AggregateId string `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId" json:"aggregate_id,omitempty"`
}

func (m *EventFilter) Reset()                    { *m = EventFilter{} }
func (m *EventFilter) String() string            { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()               {}
func (*EventFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventFilter) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventFilter) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

type EventResponse struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EventResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "pb.Event")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*EventFilter)(nil), "pb.EventFilter")
	proto.RegisterType((*EventResponse)(nil), "pb.EventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventStore service

type EventStoreClient interface {
	// Get all events for the given aggregate and event
	GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error)
	// Create a new event to the event store
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/pb.EventStore/GetEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pb.EventStore/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventStore service

type EventStoreServer interface {
	// Get all events for the given aggregate and event
	GetEvents(context.Context, *EventFilter) (*EventResponse, error)
	// Create a new event to the event store
	CreateEvent(context.Context, *Event) (*Response, error)
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EventStore/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).GetEvents(ctx, req.(*EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _EventStore_GetEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventstore.proto",
}

func init() { proto.RegisterFile("eventstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xdd, 0x6a, 0xb3, 0x40,
	0x10, 0xfd, 0x34, 0x9f, 0x46, 0xc7, 0xa4, 0x3f, 0x4b, 0x2f, 0xb6, 0x81, 0x42, 0x22, 0x14, 0xbc,
	0x12, 0x6a, 0x1f, 0xa0, 0x17, 0xfd, 0x23, 0xf4, 0xce, 0xf4, 0x3e, 0xac, 0x3a, 0x58, 0x21, 0xe8,
	0xb2, 0xbb, 0x2d, 0xe4, 0xfd, 0xfa, 0x60, 0xc5, 0x59, 0xd1, 0xd2, 0x42, 0xef, 0x76, 0xce, 0x9c,
	0x3d, 0x67, 0xce, 0x0c, 0x9c, 0xe1, 0x07, 0xb6, 0x46, 0x9b, 0x4e, 0x61, 0x2a, 0x55, 0x67, 0x3a,
	0xe6, 0xca, 0x22, 0xfe, 0x74, 0xc0, 0x7b, 0xec, 0x1b, 0xec, 0x12, 0x02, 0x62, 0xec, 0x9b, 0x8a,
	0x3b, 0x6b, 0x27, 0x09, 0xf3, 0x39, 0xd5, 0xdb, 0x8a, 0x5d, 0x01, 0xd8, 0x96, 0x39, 0x4a, 0xe4,
	0x2e, 0x35, 0x43, 0x42, 0x5e, 0x8f, 0x12, 0xd9, 0x06, 0x16, 0xa2, 0xae, 0x15, 0xd6, 0xc2, 0x60,
	0xff, 0x7b, 0x46, 0x84, 0x68, 0xc4, 0xb6, 0x15, 0xbb, 0x86, 0x93, 0x89, 0x42, 0x2a, 0xff, 0x89,
	0xb4, 0x1c, 0x51, 0x52, 0x1a, 0x8d, 0x2a, 0x61, 0x04, 0xf7, 0xbe, 0x19, 0x3d, 0x08, 0x23, 0x18,
	0x87, 0x79, 0xf9, 0x26, 0xda, 0x16, 0x0f, 0xdc, 0xb7, 0x13, 0x0e, 0x65, 0x7c, 0x07, 0x41, 0x8e,
	0x5a, 0x76, 0xad, 0x26, 0x91, 0x46, 0xef, 0xf5, 0x7b, 0x59, 0xa2, 0xd6, 0x14, 0x25, 0xc8, 0xc3,
	0x46, 0xef, 0x2c, 0xc0, 0x2e, 0xc0, 0x43, 0xa5, 0x3a, 0x35, 0xe4, 0xb0, 0x45, 0xfc, 0x02, 0x11,
	0xad, 0xe1, 0xa9, 0x39, 0x18, 0x54, 0x7f, 0x2d, 0xe3, 0x67, 0x5a, 0xf7, 0x57, 0xda, 0x38, 0x83,
	0x25, 0x89, 0x8d, 0x23, 0x6d, 0xc0, 0xb7, 0xdb, 0xe7, 0xce, 0x7a, 0x96, 0x44, 0x59, 0x98, 0xca,
	0x22, 0xb5, 0x94, 0xa1, 0x91, 0x35, 0x00, 0x04, 0xec, 0xfa, 0x03, 0xb1, 0x1b, 0x08, 0x9f, 0xd1,
	0x10, 0xa0, 0xd9, 0xe9, 0xc8, 0xb6, 0xd3, 0xad, 0xce, 0xa7, 0xef, 0x83, 0x43, 0xfc, 0x8f, 0x25,
	0x10, 0xdd, 0x2b, 0x14, 0x06, 0xed, 0x39, 0x27, 0x8b, 0xd5, 0xa2, 0x7f, 0x4e, 0xcc, 0xc2, 0xa7,
	0xf3, 0xdf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xcd, 0xb3, 0x6c, 0x12, 0x02, 0x00, 0x00,
}
