// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/cdn/v1/rule.proto

package cdn

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

// Resource rule.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rule ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Rule name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Rule pattern.
	// Must be a valid regular expression.
	RulePattern string           `protobuf:"bytes,3,opt,name=rule_pattern,json=rulePattern,proto3" json:"rule_pattern,omitempty"`
	Options     *ResourceOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetRulePattern() string {
	if x != nil {
		return x.RulePattern
	}
	return ""
}

func (x *Rule) GetOptions() *ResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_yandex_cloud_cdn_v1_rule_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_rule_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa,
	0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30,
	0x52, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x3e, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x56, 0x0a,
	0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x64, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_rule_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_rule_proto_rawDescData = file_yandex_cloud_cdn_v1_rule_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_rule_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_rule_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_rule_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_rule_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_cdn_v1_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),            // 0: yandex.cloud.cdn.v1.Rule
	(*ResourceOptions)(nil), // 1: yandex.cloud.cdn.v1.ResourceOptions
}
var file_yandex_cloud_cdn_v1_rule_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.cdn.v1.Rule.options:type_name -> yandex.cloud.cdn.v1.ResourceOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_rule_proto_init() }
func file_yandex_cloud_cdn_v1_rule_proto_init() {
	if File_yandex_cloud_cdn_v1_rule_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
			RawDescriptor: file_yandex_cloud_cdn_v1_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_rule_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_rule_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_rule_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_rule_proto = out.File
	file_yandex_cloud_cdn_v1_rule_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_rule_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_rule_proto_depIdxs = nil
}
