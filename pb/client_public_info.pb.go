// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: client_public_info.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientPublicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname         string                 `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Email           string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordUpdated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=password_updated,json=passwordUpdated,proto3" json:"password_updated,omitempty"`
	Created         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ClientPublicInfo) Reset() {
	*x = ClientPublicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_public_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPublicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPublicInfo) ProtoMessage() {}

func (x *ClientPublicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_public_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPublicInfo.ProtoReflect.Descriptor instead.
func (*ClientPublicInfo) Descriptor() ([]byte, []int) {
	return file_client_public_info_proto_rawDescGZIP(), []int{0}
}

func (x *ClientPublicInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientPublicInfo) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *ClientPublicInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ClientPublicInfo) GetPasswordUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordUpdated
	}
	return nil
}

func (x *ClientPublicInfo) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ClientPublicInfo) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_client_public_info_proto protoreflect.FileDescriptor

var file_client_public_info_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x02, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x45, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x64, 0x76, 0x69, 0x6e, 0x2f,
	0x67, 0x6f, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_public_info_proto_rawDescOnce sync.Once
	file_client_public_info_proto_rawDescData = file_client_public_info_proto_rawDesc
)

func file_client_public_info_proto_rawDescGZIP() []byte {
	file_client_public_info_proto_rawDescOnce.Do(func() {
		file_client_public_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_public_info_proto_rawDescData)
	})
	return file_client_public_info_proto_rawDescData
}

var file_client_public_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_public_info_proto_goTypes = []interface{}{
	(*ClientPublicInfo)(nil),      // 0: pb.ClientPublicInfo
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_client_public_info_proto_depIdxs = []int32{
	1, // 0: pb.ClientPublicInfo.password_updated:type_name -> google.protobuf.Timestamp
	1, // 1: pb.ClientPublicInfo.created:type_name -> google.protobuf.Timestamp
	1, // 2: pb.ClientPublicInfo.updated:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_public_info_proto_init() }
func file_client_public_info_proto_init() {
	if File_client_public_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_public_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPublicInfo); i {
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
			RawDescriptor: file_client_public_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_public_info_proto_goTypes,
		DependencyIndexes: file_client_public_info_proto_depIdxs,
		MessageInfos:      file_client_public_info_proto_msgTypes,
	}.Build()
	File_client_public_info_proto = out.File
	file_client_public_info_proto_rawDesc = nil
	file_client_public_info_proto_goTypes = nil
	file_client_public_info_proto_depIdxs = nil
}
