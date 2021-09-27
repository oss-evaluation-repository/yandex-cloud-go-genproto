// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/greenplum/v1/host.proto

package greenplum

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

type Host_Type int32

const (
	Host_TYPE_UNSPECIFIED Host_Type = 0
	// Greenplum master host.
	Host_MASTER Host_Type = 1
	// Greenplum master host.
	Host_REPLICA Host_Type = 2
	// Greenplum segment host.
	Host_SEGMENT Host_Type = 3
)

// Enum value maps for Host_Type.
var (
	Host_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "MASTER",
		2: "REPLICA",
		3: "SEGMENT",
	}
	Host_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"MASTER":           1,
		"REPLICA":          2,
		"SEGMENT":          3,
	}
)

func (x Host_Type) Enum() *Host_Type {
	p := new(Host_Type)
	*p = x
	return p
}

func (x Host_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Host_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes[0].Descriptor()
}

func (Host_Type) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes[0]
}

func (x Host_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Host_Type.Descriptor instead.
func (Host_Type) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescGZIP(), []int{0, 0}
}

type Host_Health int32

const (
	// Health of the host is unknown.
	Host_UNKNOWN Host_Health = 0
	// The host is performing all its functions normally.
	Host_ALIVE Host_Health = 1
	// The host is inoperable, and cannot perform any of its essential functions.
	Host_DEAD Host_Health = 2
	// The host is working below capacity or not fully functional.
	Host_DEGRADED Host_Health = 3
	// One or more segments are not in prefer role.
	Host_UNBALANCED Host_Health = 4
)

// Enum value maps for Host_Health.
var (
	Host_Health_name = map[int32]string{
		0: "UNKNOWN",
		1: "ALIVE",
		2: "DEAD",
		3: "DEGRADED",
		4: "UNBALANCED",
	}
	Host_Health_value = map[string]int32{
		"UNKNOWN":    0,
		"ALIVE":      1,
		"DEAD":       2,
		"DEGRADED":   3,
		"UNBALANCED": 4,
	}
)

func (x Host_Health) Enum() *Host_Health {
	p := new(Host_Health)
	*p = x
	return p
}

func (x Host_Health) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Host_Health) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes[1].Descriptor()
}

func (Host_Health) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes[1]
}

func (x Host_Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Host_Health.Descriptor instead.
func (Host_Health) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescGZIP(), []int{0, 1}
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the Greenplum host. The host name is assigned by MDB at creation time, and cannot be changed.
	// 1-63 characters long.
	//
	// The name is unique across all existing MDB hosts in Yandex.Cloud, as it defines the FQDN of the host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the Greenplum cluster. The ID is assigned by MDB at creation time.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the availability zone where the Greenplum host resides.
	ZoneId string `protobuf:"bytes,3,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Type of the host.
	Type Host_Type `protobuf:"varint,4,opt,name=type,proto3,enum=yandex.cloud.mdb.greenplum.v1.Host_Type" json:"type,omitempty"`
	// Resources allocated to the Greenplum host.
	Resources *Resources `protobuf:"bytes,5,opt,name=resources,proto3" json:"resources,omitempty"`
	// Status code of the aggregated health of the host.
	Health Host_Health `protobuf:"varint,6,opt,name=health,proto3,enum=yandex.cloud.mdb.greenplum.v1.Host_Health" json:"health,omitempty"`
	// ID of the subnet that the host belongs to.
	SubnetId string `protobuf:"bytes,7,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Flag showing public IP assignment status to this host.
	AssignPublicIp bool `protobuf:"varint,8,opt,name=assign_public_ip,json=assignPublicIp,proto3" json:"assign_public_ip,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Host) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Host) GetType() Host_Type {
	if x != nil {
		return x.Type
	}
	return Host_TYPE_UNSPECIFIED
}

func (x *Host) GetResources() *Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Host) GetHealth() Host_Health {
	if x != nil {
		return x.Health
	}
	return Host_UNKNOWN
}

func (x *Host) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *Host) GetAssignPublicIp() bool {
	if x != nil {
		return x.AssignPublicIp
	}
	return false
}

var File_yandex_cloud_mdb_greenplum_v1_host_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65,
	0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x03, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x33, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x22, 0x42, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x22, 0x48, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x47, 0x52,
	0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x42, 0x41, 0x4c, 0x41,
	0x4e, 0x43, 0x45, 0x44, 0x10, 0x04, 0x42, 0x75, 0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x03, 0x47, 0x50, 0x48,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescData = file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDesc
)

func file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDescData
}

var file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_mdb_greenplum_v1_host_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_mdb_greenplum_v1_host_proto_goTypes = []interface{}{
	(Host_Type)(0),    // 0: yandex.cloud.mdb.greenplum.v1.Host.Type
	(Host_Health)(0),  // 1: yandex.cloud.mdb.greenplum.v1.Host.Health
	(*Host)(nil),      // 2: yandex.cloud.mdb.greenplum.v1.Host
	(*Resources)(nil), // 3: yandex.cloud.mdb.greenplum.v1.Resources
}
var file_yandex_cloud_mdb_greenplum_v1_host_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.mdb.greenplum.v1.Host.type:type_name -> yandex.cloud.mdb.greenplum.v1.Host.Type
	3, // 1: yandex.cloud.mdb.greenplum.v1.Host.resources:type_name -> yandex.cloud.mdb.greenplum.v1.Resources
	1, // 2: yandex.cloud.mdb.greenplum.v1.Host.health:type_name -> yandex.cloud.mdb.greenplum.v1.Host.Health
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_greenplum_v1_host_proto_init() }
func file_yandex_cloud_mdb_greenplum_v1_host_proto_init() {
	if File_yandex_cloud_mdb_greenplum_v1_host_proto != nil {
		return
	}
	file_yandex_cloud_mdb_greenplum_v1_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_greenplum_v1_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
			RawDescriptor: file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_greenplum_v1_host_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_greenplum_v1_host_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_greenplum_v1_host_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_greenplum_v1_host_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_greenplum_v1_host_proto = out.File
	file_yandex_cloud_mdb_greenplum_v1_host_proto_rawDesc = nil
	file_yandex_cloud_mdb_greenplum_v1_host_proto_goTypes = nil
	file_yandex_cloud_mdb_greenplum_v1_host_proto_depIdxs = nil
}
