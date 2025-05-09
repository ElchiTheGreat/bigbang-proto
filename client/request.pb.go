// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: client/request.proto

package client

import (
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

// Deploy related messages
type RequestDeploy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ipv4      string `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Port      uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Version   string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Bootstrap []byte `protobuf:"bytes,5,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
}

func (x *RequestDeploy) Reset() {
	*x = RequestDeploy{}
	mi := &file_client_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestDeploy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDeploy) ProtoMessage() {}

func (x *RequestDeploy) ProtoReflect() protoreflect.Message {
	mi := &file_client_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDeploy.ProtoReflect.Descriptor instead.
func (*RequestDeploy) Descriptor() ([]byte, []int) {
	return file_client_request_proto_rawDescGZIP(), []int{0}
}

func (x *RequestDeploy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestDeploy) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *RequestDeploy) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RequestDeploy) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RequestDeploy) GetBootstrap() []byte {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

var File_client_request_proto protoreflect.FileDescriptor

var file_client_request_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x83,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x63, 0x68, 0x69, 0x54, 0x68, 0x65, 0x47, 0x72, 0x65, 0x61, 0x74,
	0x2f, 0x62, 0x69, 0x67, 0x62, 0x61, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_request_proto_rawDescOnce sync.Once
	file_client_request_proto_rawDescData = file_client_request_proto_rawDesc
)

func file_client_request_proto_rawDescGZIP() []byte {
	file_client_request_proto_rawDescOnce.Do(func() {
		file_client_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_request_proto_rawDescData)
	})
	return file_client_request_proto_rawDescData
}

var file_client_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_request_proto_goTypes = []any{
	(*RequestDeploy)(nil), // 0: client.RequestDeploy
}
var file_client_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_request_proto_init() }
func file_client_request_proto_init() {
	if File_client_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_request_proto_goTypes,
		DependencyIndexes: file_client_request_proto_depIdxs,
		MessageInfos:      file_client_request_proto_msgTypes,
	}.Build()
	File_client_request_proto = out.File
	file_client_request_proto_rawDesc = nil
	file_client_request_proto_goTypes = nil
	file_client_request_proto_depIdxs = nil
}
