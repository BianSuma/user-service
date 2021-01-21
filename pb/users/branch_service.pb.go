// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: users/branch_service.proto

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

type ListBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	RegionId   string      `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *ListBranchRequest) Reset() {
	*x = ListBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchRequest) ProtoMessage() {}

func (x *ListBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchRequest.ProtoReflect.Descriptor instead.
func (*ListBranchRequest) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListBranchRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListBranchRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type BranchPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	RegionId   string      `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Count      uint32      `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BranchPaginationResponse) Reset() {
	*x = BranchPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchPaginationResponse) ProtoMessage() {}

func (x *BranchPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchPaginationResponse.ProtoReflect.Descriptor instead.
func (*BranchPaginationResponse) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{1}
}

func (x *BranchPaginationResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *BranchPaginationResponse) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *BranchPaginationResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *BranchPaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Branch     *Branch                   `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *ListBranchResponse) Reset() {
	*x = ListBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchResponse) ProtoMessage() {}

func (x *ListBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchResponse.ProtoReflect.Descriptor instead.
func (*ListBranchResponse) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBranchResponse) GetPagination() *BranchPaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListBranchResponse) GetBranch() *Branch {
	if x != nil {
		return x.Branch
	}
	return nil
}

var File_users_branch_service_proto protoreflect.FileDescriptor

var file_users_branch_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1a, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x18, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x8e, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x32, 0xc9, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x17,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_branch_service_proto_rawDescOnce sync.Once
	file_users_branch_service_proto_rawDescData = file_users_branch_service_proto_rawDesc
)

func file_users_branch_service_proto_rawDescGZIP() []byte {
	file_users_branch_service_proto_rawDescOnce.Do(func() {
		file_users_branch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_branch_service_proto_rawDescData)
	})
	return file_users_branch_service_proto_rawDescData
}

var file_users_branch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_users_branch_service_proto_goTypes = []interface{}{
	(*ListBranchRequest)(nil),        // 0: wiradata.users.ListBranchRequest
	(*BranchPaginationResponse)(nil), // 1: wiradata.users.BranchPaginationResponse
	(*ListBranchResponse)(nil),       // 2: wiradata.users.ListBranchResponse
	(*Pagination)(nil),               // 3: wiradata.users.Pagination
	(*Branch)(nil),                   // 4: wiradata.users.Branch
	(*Id)(nil),                       // 5: wiradata.users.Id
	(*Boolean)(nil),                  // 6: wiradata.users.Boolean
}
var file_users_branch_service_proto_depIdxs = []int32{
	3, // 0: wiradata.users.ListBranchRequest.pagination:type_name -> wiradata.users.Pagination
	3, // 1: wiradata.users.BranchPaginationResponse.pagination:type_name -> wiradata.users.Pagination
	1, // 2: wiradata.users.ListBranchResponse.pagination:type_name -> wiradata.users.BranchPaginationResponse
	4, // 3: wiradata.users.ListBranchResponse.branch:type_name -> wiradata.users.Branch
	4, // 4: wiradata.users.BranchService.Create:input_type -> wiradata.users.Branch
	4, // 5: wiradata.users.BranchService.Update:input_type -> wiradata.users.Branch
	5, // 6: wiradata.users.BranchService.View:input_type -> wiradata.users.Id
	5, // 7: wiradata.users.BranchService.Delete:input_type -> wiradata.users.Id
	0, // 8: wiradata.users.BranchService.List:input_type -> wiradata.users.ListBranchRequest
	4, // 9: wiradata.users.BranchService.Create:output_type -> wiradata.users.Branch
	4, // 10: wiradata.users.BranchService.Update:output_type -> wiradata.users.Branch
	4, // 11: wiradata.users.BranchService.View:output_type -> wiradata.users.Branch
	6, // 12: wiradata.users.BranchService.Delete:output_type -> wiradata.users.Boolean
	2, // 13: wiradata.users.BranchService.List:output_type -> wiradata.users.ListBranchResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_users_branch_service_proto_init() }
func file_users_branch_service_proto_init() {
	if File_users_branch_service_proto != nil {
		return
	}
	file_users_branch_message_proto_init()
	file_users_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_branch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchRequest); i {
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
		file_users_branch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchPaginationResponse); i {
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
		file_users_branch_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchResponse); i {
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
			RawDescriptor: file_users_branch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_branch_service_proto_goTypes,
		DependencyIndexes: file_users_branch_service_proto_depIdxs,
		MessageInfos:      file_users_branch_service_proto_msgTypes,
	}.Build()
	File_users_branch_service_proto = out.File
	file_users_branch_service_proto_rawDesc = nil
	file_users_branch_service_proto_goTypes = nil
	file_users_branch_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BranchServiceClient interface {
	Create(ctx context.Context, in *Branch, opts ...grpc.CallOption) (*Branch, error)
	Update(ctx context.Context, in *Branch, opts ...grpc.CallOption) (*Branch, error)
	View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Branch, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Boolean, error)
	List(ctx context.Context, in *ListBranchRequest, opts ...grpc.CallOption) (BranchService_ListClient, error)
}

type branchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchServiceClient(cc grpc.ClientConnInterface) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) Create(ctx context.Context, in *Branch, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/wiradata.users.BranchService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) Update(ctx context.Context, in *Branch, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/wiradata.users.BranchService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/wiradata.users.BranchService/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Boolean, error) {
	out := new(Boolean)
	err := c.cc.Invoke(ctx, "/wiradata.users.BranchService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) List(ctx context.Context, in *ListBranchRequest, opts ...grpc.CallOption) (BranchService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BranchService_serviceDesc.Streams[0], "/wiradata.users.BranchService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &branchServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BranchService_ListClient interface {
	Recv() (*ListBranchResponse, error)
	grpc.ClientStream
}

type branchServiceListClient struct {
	grpc.ClientStream
}

func (x *branchServiceListClient) Recv() (*ListBranchResponse, error) {
	m := new(ListBranchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BranchServiceServer is the server API for BranchService service.
type BranchServiceServer interface {
	Create(context.Context, *Branch) (*Branch, error)
	Update(context.Context, *Branch) (*Branch, error)
	View(context.Context, *Id) (*Branch, error)
	Delete(context.Context, *Id) (*Boolean, error)
	List(*ListBranchRequest, BranchService_ListServer) error
}

// UnimplementedBranchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBranchServiceServer struct {
}

func (*UnimplementedBranchServiceServer) Create(context.Context, *Branch) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBranchServiceServer) Update(context.Context, *Branch) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedBranchServiceServer) View(context.Context, *Id) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (*UnimplementedBranchServiceServer) Delete(context.Context, *Id) (*Boolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedBranchServiceServer) List(*ListBranchRequest, BranchService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterBranchServiceServer(s *grpc.Server, srv BranchServiceServer) {
	s.RegisterService(&_BranchService_serviceDesc, srv)
}

func _BranchService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Branch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.users.BranchService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Create(ctx, req.(*Branch))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Branch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.users.BranchService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Update(ctx, req.(*Branch))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.users.BranchService/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).View(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.users.BranchService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBranchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BranchServiceServer).List(m, &branchServiceListServer{stream})
}

type BranchService_ListServer interface {
	Send(*ListBranchResponse) error
	grpc.ServerStream
}

type branchServiceListServer struct {
	grpc.ServerStream
}

func (x *branchServiceListServer) Send(m *ListBranchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BranchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.users.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BranchService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BranchService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _BranchService_View_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BranchService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _BranchService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/branch_service.proto",
}