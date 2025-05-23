// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/service/endpoint/v3/leds.proto

package endpointv3

import (
	context "context"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type LedsDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LedsDummy) Reset() {
	*x = LedsDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_endpoint_v3_leds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedsDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedsDummy) ProtoMessage() {}

func (x *LedsDummy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_endpoint_v3_leds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedsDummy.ProtoReflect.Descriptor instead.
func (*LedsDummy) Descriptor() ([]byte, []int) {
	return file_envoy_service_endpoint_v3_leds_proto_rawDescGZIP(), []int{0}
}

var File_envoy_service_endpoint_v3_leds_proto protoreflect.FileDescriptor

var file_envoy_service_endpoint_v3_leds_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x65, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76,
	0x33, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b,
	0x0a, 0x09, 0x4c, 0x65, 0x64, 0x73, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0xd7, 0x01, 0x0a, 0x20,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x85, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x1a, 0x2b, 0x8a, 0xa4, 0x96, 0xf3, 0x07, 0x25,
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x62, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x8b, 0x01, 0x0a, 0x27, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76,
	0x33, 0x42, 0x09, 0x4c, 0x65, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x33,
	0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_endpoint_v3_leds_proto_rawDescOnce sync.Once
	file_envoy_service_endpoint_v3_leds_proto_rawDescData = file_envoy_service_endpoint_v3_leds_proto_rawDesc
)

func file_envoy_service_endpoint_v3_leds_proto_rawDescGZIP() []byte {
	file_envoy_service_endpoint_v3_leds_proto_rawDescOnce.Do(func() {
		file_envoy_service_endpoint_v3_leds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_endpoint_v3_leds_proto_rawDescData)
	})
	return file_envoy_service_endpoint_v3_leds_proto_rawDescData
}

var file_envoy_service_endpoint_v3_leds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_service_endpoint_v3_leds_proto_goTypes = []interface{}{
	(*LedsDummy)(nil),                 // 0: envoy.service.endpoint.v3.LedsDummy
	(*v3.DeltaDiscoveryRequest)(nil),  // 1: envoy.service.discovery.v3.DeltaDiscoveryRequest
	(*v3.DeltaDiscoveryResponse)(nil), // 2: envoy.service.discovery.v3.DeltaDiscoveryResponse
}
var file_envoy_service_endpoint_v3_leds_proto_depIdxs = []int32{
	1, // 0: envoy.service.endpoint.v3.LocalityEndpointDiscoveryService.DeltaLocalityEndpoints:input_type -> envoy.service.discovery.v3.DeltaDiscoveryRequest
	2, // 1: envoy.service.endpoint.v3.LocalityEndpointDiscoveryService.DeltaLocalityEndpoints:output_type -> envoy.service.discovery.v3.DeltaDiscoveryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_service_endpoint_v3_leds_proto_init() }
func file_envoy_service_endpoint_v3_leds_proto_init() {
	if File_envoy_service_endpoint_v3_leds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_endpoint_v3_leds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedsDummy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_endpoint_v3_leds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_endpoint_v3_leds_proto_goTypes,
		DependencyIndexes: file_envoy_service_endpoint_v3_leds_proto_depIdxs,
		MessageInfos:      file_envoy_service_endpoint_v3_leds_proto_msgTypes,
	}.Build()
	File_envoy_service_endpoint_v3_leds_proto = out.File
	file_envoy_service_endpoint_v3_leds_proto_rawDesc = nil
	file_envoy_service_endpoint_v3_leds_proto_goTypes = nil
	file_envoy_service_endpoint_v3_leds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocalityEndpointDiscoveryServiceClient is the client API for LocalityEndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocalityEndpointDiscoveryServiceClient interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(ctx context.Context, opts ...grpc.CallOption) (LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient, error)
}

type localityEndpointDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalityEndpointDiscoveryServiceClient(cc grpc.ClientConnInterface) LocalityEndpointDiscoveryServiceClient {
	return &localityEndpointDiscoveryServiceClient{cc}
}

func (c *localityEndpointDiscoveryServiceClient) DeltaLocalityEndpoints(ctx context.Context, opts ...grpc.CallOption) (LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LocalityEndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.service.endpoint.v3.LocalityEndpointDiscoveryService/DeltaLocalityEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient{stream}
	return x, nil
}

type LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient struct {
	grpc.ClientStream
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocalityEndpointDiscoveryServiceServer is the server API for LocalityEndpointDiscoveryService service.
type LocalityEndpointDiscoveryServiceServer interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer) error
}

// UnimplementedLocalityEndpointDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLocalityEndpointDiscoveryServiceServer struct {
}

func (*UnimplementedLocalityEndpointDiscoveryServiceServer) DeltaLocalityEndpoints(LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaLocalityEndpoints not implemented")
}

func RegisterLocalityEndpointDiscoveryServiceServer(s *grpc.Server, srv LocalityEndpointDiscoveryServiceServer) {
	s.RegisterService(&_LocalityEndpointDiscoveryService_serviceDesc, srv)
}

func _LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocalityEndpointDiscoveryServiceServer).DeltaLocalityEndpoints(&localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer{stream})
}

type LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer struct {
	grpc.ServerStream
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LocalityEndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.endpoint.v3.LocalityEndpointDiscoveryService",
	HandlerType: (*LocalityEndpointDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaLocalityEndpoints",
			Handler:       _LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/endpoint/v3/leds.proto",
}
