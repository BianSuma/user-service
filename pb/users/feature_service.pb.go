// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: users/feature_service.proto

package users

import (
	context "context"
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

type ListFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *Feature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *ListFeatureResponse) Reset() {
	*x = ListFeatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_feature_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureResponse) ProtoMessage() {}

func (x *ListFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_feature_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureResponse.ProtoReflect.Descriptor instead.
func (*ListFeatureResponse) Descriptor() ([]byte, []int) {
	return file_users_feature_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListFeatureResponse) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type ListPackageFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageOfFeature *PackageOfFeature `protobuf:"bytes,1,opt,name=package_of_feature,json=packageOfFeature,proto3" json:"package_of_feature,omitempty"`
}

func (x *ListPackageFeatureResponse) Reset() {
	*x = ListPackageFeatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_feature_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackageFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackageFeatureResponse) ProtoMessage() {}

func (x *ListPackageFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_feature_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackageFeatureResponse.ProtoReflect.Descriptor instead.
func (*ListPackageFeatureResponse) Descriptor() ([]byte, []int) {
	return file_users_feature_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPackageFeatureResponse) GetPackageOfFeature() *PackageOfFeature {
	if x != nil {
		return x.PackageOfFeature
	}
	return nil
}

var File_users_feature_service_proto protoreflect.FileDescriptor

var file_users_feature_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1b, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x6c, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x12, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x10, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32,
	0x58, 0x0a, 0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x23, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0xa6, 0x01, 0x0a, 0x15, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x12, 0x2e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a,
	0x20, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_feature_service_proto_rawDescOnce sync.Once
	file_users_feature_service_proto_rawDescData = file_users_feature_service_proto_rawDesc
)

func file_users_feature_service_proto_rawDescGZIP() []byte {
	file_users_feature_service_proto_rawDescOnce.Do(func() {
		file_users_feature_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_feature_service_proto_rawDescData)
	})
	return file_users_feature_service_proto_rawDescData
}

var file_users_feature_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_users_feature_service_proto_goTypes = []interface{}{
	(*ListFeatureResponse)(nil),        // 0: wiradata.users.ListFeatureResponse
	(*ListPackageFeatureResponse)(nil), // 1: wiradata.users.ListPackageFeatureResponse
	(*Feature)(nil),                    // 2: wiradata.users.Feature
	(*PackageOfFeature)(nil),           // 3: wiradata.users.PackageOfFeature
	(*Empty)(nil),                      // 4: wiradata.users.Empty
	(*Id)(nil),                         // 5: wiradata.users.Id
}
var file_users_feature_service_proto_depIdxs = []int32{
	2, // 0: wiradata.users.ListFeatureResponse.feature:type_name -> wiradata.users.Feature
	3, // 1: wiradata.users.ListPackageFeatureResponse.package_of_feature:type_name -> wiradata.users.PackageOfFeature
	4, // 2: wiradata.users.FeatureService.List:input_type -> wiradata.users.Empty
	5, // 3: wiradata.users.PackageFeatureService.View:input_type -> wiradata.users.Id
	4, // 4: wiradata.users.PackageFeatureService.List:input_type -> wiradata.users.Empty
	0, // 5: wiradata.users.FeatureService.List:output_type -> wiradata.users.ListFeatureResponse
	3, // 6: wiradata.users.PackageFeatureService.View:output_type -> wiradata.users.PackageOfFeature
	1, // 7: wiradata.users.PackageFeatureService.List:output_type -> wiradata.users.ListPackageFeatureResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_users_feature_service_proto_init() }
func file_users_feature_service_proto_init() {
	if File_users_feature_service_proto != nil {
		return
	}
	file_users_feature_message_proto_init()
	file_users_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_feature_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeatureResponse); i {
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
		file_users_feature_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackageFeatureResponse); i {
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
			RawDescriptor: file_users_feature_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_users_feature_service_proto_goTypes,
		DependencyIndexes: file_users_feature_service_proto_depIdxs,
		MessageInfos:      file_users_feature_service_proto_msgTypes,
	}.Build()
	File_users_feature_service_proto = out.File
	file_users_feature_service_proto_rawDesc = nil
	file_users_feature_service_proto_goTypes = nil
	file_users_feature_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeatureServiceClient is the client API for FeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeatureServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (FeatureService_ListClient, error)
}

type featureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureServiceClient(cc grpc.ClientConnInterface) FeatureServiceClient {
	return &featureServiceClient{cc}
}

func (c *featureServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (FeatureService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FeatureService_serviceDesc.Streams[0], "/wiradata.users.FeatureService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &featureServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeatureService_ListClient interface {
	Recv() (*ListFeatureResponse, error)
	grpc.ClientStream
}

type featureServiceListClient struct {
	grpc.ClientStream
}

func (x *featureServiceListClient) Recv() (*ListFeatureResponse, error) {
	m := new(ListFeatureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FeatureServiceServer is the server API for FeatureService service.
type FeatureServiceServer interface {
	List(*Empty, FeatureService_ListServer) error
}

// UnimplementedFeatureServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeatureServiceServer struct {
}

func (*UnimplementedFeatureServiceServer) List(*Empty, FeatureService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterFeatureServiceServer(s *grpc.Server, srv FeatureServiceServer) {
	s.RegisterService(&_FeatureService_serviceDesc, srv)
}

func _FeatureService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeatureServiceServer).List(m, &featureServiceListServer{stream})
}

type FeatureService_ListServer interface {
	Send(*ListFeatureResponse) error
	grpc.ServerStream
}

type featureServiceListServer struct {
	grpc.ServerStream
}

func (x *featureServiceListServer) Send(m *ListFeatureResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _FeatureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.users.FeatureService",
	HandlerType: (*FeatureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _FeatureService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/feature_service.proto",
}

// PackageFeatureServiceClient is the client API for PackageFeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PackageFeatureServiceClient interface {
	View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PackageOfFeature, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PackageFeatureService_ListClient, error)
}

type packageFeatureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageFeatureServiceClient(cc grpc.ClientConnInterface) PackageFeatureServiceClient {
	return &packageFeatureServiceClient{cc}
}

func (c *packageFeatureServiceClient) View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PackageOfFeature, error) {
	out := new(PackageOfFeature)
	err := c.cc.Invoke(ctx, "/wiradata.users.PackageFeatureService/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageFeatureServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PackageFeatureService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PackageFeatureService_serviceDesc.Streams[0], "/wiradata.users.PackageFeatureService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &packageFeatureServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PackageFeatureService_ListClient interface {
	Recv() (*ListPackageFeatureResponse, error)
	grpc.ClientStream
}

type packageFeatureServiceListClient struct {
	grpc.ClientStream
}

func (x *packageFeatureServiceListClient) Recv() (*ListPackageFeatureResponse, error) {
	m := new(ListPackageFeatureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PackageFeatureServiceServer is the server API for PackageFeatureService service.
type PackageFeatureServiceServer interface {
	View(context.Context, *Id) (*PackageOfFeature, error)
	List(*Empty, PackageFeatureService_ListServer) error
}

// UnimplementedPackageFeatureServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPackageFeatureServiceServer struct {
}

func (*UnimplementedPackageFeatureServiceServer) View(context.Context, *Id) (*PackageOfFeature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (*UnimplementedPackageFeatureServiceServer) List(*Empty, PackageFeatureService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterPackageFeatureServiceServer(s *grpc.Server, srv PackageFeatureServiceServer) {
	s.RegisterService(&_PackageFeatureService_serviceDesc, srv)
}

func _PackageFeatureService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageFeatureServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.users.PackageFeatureService/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageFeatureServiceServer).View(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageFeatureService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PackageFeatureServiceServer).List(m, &packageFeatureServiceListServer{stream})
}

type PackageFeatureService_ListServer interface {
	Send(*ListPackageFeatureResponse) error
	grpc.ServerStream
}

type packageFeatureServiceListServer struct {
	grpc.ServerStream
}

func (x *packageFeatureServiceListServer) Send(m *ListPackageFeatureResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PackageFeatureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.users.PackageFeatureService",
	HandlerType: (*PackageFeatureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "View",
			Handler:    _PackageFeatureService_View_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _PackageFeatureService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/feature_service.proto",
}
