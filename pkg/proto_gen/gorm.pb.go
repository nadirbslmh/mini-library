// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: gorm.proto

package proto_gen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type DeletedAt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Valid bool                 `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *DeletedAt) Reset() {
	*x = DeletedAt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletedAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletedAt) ProtoMessage() {}

func (x *DeletedAt) ProtoReflect() protoreflect.Message {
	mi := &file_gorm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletedAt.ProtoReflect.Descriptor instead.
func (*DeletedAt) Descriptor() ([]byte, []int) {
	return file_gorm_proto_rawDescGZIP(), []int{0}
}

func (x *DeletedAt) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeletedAt) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_gorm_proto protoreflect.FileDescriptor

var file_gorm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gorm_proto_rawDescOnce sync.Once
	file_gorm_proto_rawDescData = file_gorm_proto_rawDesc
)

func file_gorm_proto_rawDescGZIP() []byte {
	file_gorm_proto_rawDescOnce.Do(func() {
		file_gorm_proto_rawDescData = protoimpl.X.CompressGZIP(file_gorm_proto_rawDescData)
	})
	return file_gorm_proto_rawDescData
}

var file_gorm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gorm_proto_goTypes = []interface{}{
	(*DeletedAt)(nil),           // 0: DeletedAt
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_gorm_proto_depIdxs = []int32{
	1, // 0: DeletedAt.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gorm_proto_init() }
func file_gorm_proto_init() {
	if File_gorm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gorm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletedAt); i {
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
			RawDescriptor: file_gorm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gorm_proto_goTypes,
		DependencyIndexes: file_gorm_proto_depIdxs,
		MessageInfos:      file_gorm_proto_msgTypes,
	}.Build()
	File_gorm_proto = out.File
	file_gorm_proto_rawDesc = nil
	file_gorm_proto_goTypes = nil
	file_gorm_proto_depIdxs = nil
}