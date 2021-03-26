// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/mysql/v1/database.proto

package mysql

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// A MySQL database. For more information, see
// the [documentation](/docs/managed-mysql/concepts).
type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the MySQL cluster that the database belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type DatabaseSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the MySQL database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DatabaseSpec) Reset() {
	*x = DatabaseSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseSpec) ProtoMessage() {}

func (x *DatabaseSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseSpec.ProtoReflect.Descriptor instead.
func (*DatabaseSpec) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_yandex_cloud_mdb_mysql_v1_database_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mysql_v1_database_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73,
	0x71, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x33, 0xf2,
	0xc7, 0x31, 0x0e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d,
	0x2a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x64, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescData = file_yandex_cloud_mdb_mysql_v1_database_proto_rawDesc
)

func file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mysql_v1_database_proto_rawDescData
}

var file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_mysql_v1_database_proto_goTypes = []interface{}{
	(*Database)(nil),     // 0: yandex.cloud.mdb.mysql.v1.Database
	(*DatabaseSpec)(nil), // 1: yandex.cloud.mdb.mysql.v1.DatabaseSpec
}
var file_yandex_cloud_mdb_mysql_v1_database_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mysql_v1_database_proto_init() }
func file_yandex_cloud_mdb_mysql_v1_database_proto_init() {
	if File_yandex_cloud_mdb_mysql_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
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
		file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseSpec); i {
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
			RawDescriptor: file_yandex_cloud_mdb_mysql_v1_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_mysql_v1_database_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mysql_v1_database_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_mysql_v1_database_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mysql_v1_database_proto = out.File
	file_yandex_cloud_mdb_mysql_v1_database_proto_rawDesc = nil
	file_yandex_cloud_mdb_mysql_v1_database_proto_goTypes = nil
	file_yandex_cloud_mdb_mysql_v1_database_proto_depIdxs = nil
}
