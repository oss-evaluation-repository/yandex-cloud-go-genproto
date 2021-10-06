// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/cdn/v1/origin.proto

package cdn

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

// An origin. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the origin.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the parent origin group.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
	// IP address or Domain name of your origin and the port (if custom).
	// Used if [meta] variant is `common`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origin.
	// False - The origin is disabled and the CDN is not using it to pull content.
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Specifies whether the origin is used in its origin group as backup.
	// A backup origin is used when one of active origins becomes unavailable.
	Backup bool `protobuf:"varint,5,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up origin of the content.
	Meta *OriginMeta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{0}
}

func (x *Origin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Origin) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

func (x *Origin) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Origin) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Origin) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *Origin) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Origin parameters. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source: IP address or Domain name of your origin and the port (if custom).
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origins. False - The origin is disabled
	// and the CDN is not using it to pull content.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// backup option has two possible values:
	//
	//   True - The option is active. The origin will not be used until one of
	//          active origins become unavailable.
	//   False - The option is disabled.
	Backup bool `protobuf:"varint,3,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up origin of the content.
	Meta *OriginMeta `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *OriginParams) Reset() {
	*x = OriginParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginParams) ProtoMessage() {}

func (x *OriginParams) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginParams.ProtoReflect.Descriptor instead.
func (*OriginParams) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{1}
}

func (x *OriginParams) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OriginParams) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *OriginParams) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *OriginParams) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Origin type. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the origin.
	//
	// Types that are assignable to OriginMetaVariant:
	//	*OriginMeta_Common
	//	*OriginMeta_Bucket
	//	*OriginMeta_Website
	//	*OriginMeta_Balancer
	OriginMetaVariant isOriginMeta_OriginMetaVariant `protobuf_oneof:"origin_meta_variant"`
}

func (x *OriginMeta) Reset() {
	*x = OriginMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginMeta) ProtoMessage() {}

func (x *OriginMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginMeta.ProtoReflect.Descriptor instead.
func (*OriginMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{2}
}

func (m *OriginMeta) GetOriginMetaVariant() isOriginMeta_OriginMetaVariant {
	if m != nil {
		return m.OriginMetaVariant
	}
	return nil
}

func (x *OriginMeta) GetCommon() *OriginNamedMeta {
	if x, ok := x.GetOriginMetaVariant().(*OriginMeta_Common); ok {
		return x.Common
	}
	return nil
}

func (x *OriginMeta) GetBucket() *OriginNamedMeta {
	if x, ok := x.GetOriginMetaVariant().(*OriginMeta_Bucket); ok {
		return x.Bucket
	}
	return nil
}

func (x *OriginMeta) GetWebsite() *OriginNamedMeta {
	if x, ok := x.GetOriginMetaVariant().(*OriginMeta_Website); ok {
		return x.Website
	}
	return nil
}

func (x *OriginMeta) GetBalancer() *OriginBalancerMeta {
	if x, ok := x.GetOriginMetaVariant().(*OriginMeta_Balancer); ok {
		return x.Balancer
	}
	return nil
}

type isOriginMeta_OriginMetaVariant interface {
	isOriginMeta_OriginMetaVariant()
}

type OriginMeta_Common struct {
	// A server with a domain name linked to it
	Common *OriginNamedMeta `protobuf:"bytes,1,opt,name=common,proto3,oneof"`
}

type OriginMeta_Bucket struct {
	// A Yandex Object Storage bucket not configured as a static site hosting.
	Bucket *OriginNamedMeta `protobuf:"bytes,2,opt,name=bucket,proto3,oneof"`
}

type OriginMeta_Website struct {
	// A Yandex Object Storage bucket configured as a static site hosting.
	Website *OriginNamedMeta `protobuf:"bytes,3,opt,name=website,proto3,oneof"`
}

type OriginMeta_Balancer struct {
	// An L7 load balancer from Yandex Application Load Balancer.
	// CDN servers will access the load balancer at one of its IP addresses that must be selected in the origin settings.
	Balancer *OriginBalancerMeta `protobuf:"bytes,4,opt,name=balancer,proto3,oneof"`
}

func (*OriginMeta_Common) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Bucket) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Website) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Balancer) isOriginMeta_OriginMetaVariant() {}

// Origin info. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginNamedMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the origin.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OriginNamedMeta) Reset() {
	*x = OriginNamedMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginNamedMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginNamedMeta) ProtoMessage() {}

func (x *OriginNamedMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginNamedMeta.ProtoReflect.Descriptor instead.
func (*OriginNamedMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{3}
}

func (x *OriginNamedMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Yandex Application Load Balancer origin info. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginBalancerMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the origin.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OriginBalancerMeta) Reset() {
	*x = OriginBalancerMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginBalancerMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginBalancerMeta) ProtoMessage() {}

func (x *OriginBalancerMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginBalancerMeta.ProtoReflect.Descriptor instead.
func (*OriginBalancerMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{4}
}

func (x *OriginBalancerMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_yandex_cloud_cdn_v1_origin_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_origin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xac, 0x02, 0x0a, 0x0a, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x42, 0x15, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x24, 0x0a, 0x12, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x64, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_origin_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_origin_proto_rawDescData = file_yandex_cloud_cdn_v1_origin_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_origin_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_origin_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_origin_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_origin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_cdn_v1_origin_proto_goTypes = []interface{}{
	(*Origin)(nil),             // 0: yandex.cloud.cdn.v1.Origin
	(*OriginParams)(nil),       // 1: yandex.cloud.cdn.v1.OriginParams
	(*OriginMeta)(nil),         // 2: yandex.cloud.cdn.v1.OriginMeta
	(*OriginNamedMeta)(nil),    // 3: yandex.cloud.cdn.v1.OriginNamedMeta
	(*OriginBalancerMeta)(nil), // 4: yandex.cloud.cdn.v1.OriginBalancerMeta
}
var file_yandex_cloud_cdn_v1_origin_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.cdn.v1.Origin.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	2, // 1: yandex.cloud.cdn.v1.OriginParams.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	3, // 2: yandex.cloud.cdn.v1.OriginMeta.common:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	3, // 3: yandex.cloud.cdn.v1.OriginMeta.bucket:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	3, // 4: yandex.cloud.cdn.v1.OriginMeta.website:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	4, // 5: yandex.cloud.cdn.v1.OriginMeta.balancer:type_name -> yandex.cloud.cdn.v1.OriginBalancerMeta
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_origin_proto_init() }
func file_yandex_cloud_cdn_v1_origin_proto_init() {
	if File_yandex_cloud_cdn_v1_origin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_origin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_yandex_cloud_cdn_v1_origin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginParams); i {
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
		file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginMeta); i {
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
		file_yandex_cloud_cdn_v1_origin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginNamedMeta); i {
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
		file_yandex_cloud_cdn_v1_origin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginBalancerMeta); i {
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
	file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*OriginMeta_Common)(nil),
		(*OriginMeta_Bucket)(nil),
		(*OriginMeta_Website)(nil),
		(*OriginMeta_Balancer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_cdn_v1_origin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_origin_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_origin_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_origin_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_origin_proto = out.File
	file_yandex_cloud_cdn_v1_origin_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_origin_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_origin_proto_depIdxs = nil
}
