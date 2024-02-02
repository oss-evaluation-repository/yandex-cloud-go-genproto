// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/datatransfer/v1/endpoint.proto

package datatransfer

import (
	endpoint "github.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1/endpoint"
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

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId    string            `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	Name        string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Settings    *EndpointSettings `protobuf:"bytes,52,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Endpoint) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Endpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Endpoint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Endpoint) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Endpoint) GetSettings() *EndpointSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type EndpointSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Settings:
	//
	//	*EndpointSettings_MysqlSource
	//	*EndpointSettings_PostgresSource
	//	*EndpointSettings_YdbSource
	//	*EndpointSettings_YdsSource
	//	*EndpointSettings_KafkaSource
	//	*EndpointSettings_MongoSource
	//	*EndpointSettings_ClickhouseSource
	//	*EndpointSettings_MysqlTarget
	//	*EndpointSettings_PostgresTarget
	//	*EndpointSettings_ClickhouseTarget
	//	*EndpointSettings_YdbTarget
	//	*EndpointSettings_KafkaTarget
	//	*EndpointSettings_MongoTarget
	//	*EndpointSettings_YdsTarget
	Settings isEndpointSettings_Settings `protobuf_oneof:"settings"`
}

func (x *EndpointSettings) Reset() {
	*x = EndpointSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSettings) ProtoMessage() {}

func (x *EndpointSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSettings.ProtoReflect.Descriptor instead.
func (*EndpointSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescGZIP(), []int{1}
}

func (m *EndpointSettings) GetSettings() isEndpointSettings_Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (x *EndpointSettings) GetMysqlSource() *endpoint.MysqlSource {
	if x, ok := x.GetSettings().(*EndpointSettings_MysqlSource); ok {
		return x.MysqlSource
	}
	return nil
}

func (x *EndpointSettings) GetPostgresSource() *endpoint.PostgresSource {
	if x, ok := x.GetSettings().(*EndpointSettings_PostgresSource); ok {
		return x.PostgresSource
	}
	return nil
}

func (x *EndpointSettings) GetYdbSource() *endpoint.YdbSource {
	if x, ok := x.GetSettings().(*EndpointSettings_YdbSource); ok {
		return x.YdbSource
	}
	return nil
}

func (x *EndpointSettings) GetYdsSource() *endpoint.YDSSource {
	if x, ok := x.GetSettings().(*EndpointSettings_YdsSource); ok {
		return x.YdsSource
	}
	return nil
}

func (x *EndpointSettings) GetKafkaSource() *endpoint.KafkaSource {
	if x, ok := x.GetSettings().(*EndpointSettings_KafkaSource); ok {
		return x.KafkaSource
	}
	return nil
}

func (x *EndpointSettings) GetMongoSource() *endpoint.MongoSource {
	if x, ok := x.GetSettings().(*EndpointSettings_MongoSource); ok {
		return x.MongoSource
	}
	return nil
}

func (x *EndpointSettings) GetClickhouseSource() *endpoint.ClickhouseSource {
	if x, ok := x.GetSettings().(*EndpointSettings_ClickhouseSource); ok {
		return x.ClickhouseSource
	}
	return nil
}

func (x *EndpointSettings) GetMysqlTarget() *endpoint.MysqlTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_MysqlTarget); ok {
		return x.MysqlTarget
	}
	return nil
}

func (x *EndpointSettings) GetPostgresTarget() *endpoint.PostgresTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_PostgresTarget); ok {
		return x.PostgresTarget
	}
	return nil
}

func (x *EndpointSettings) GetClickhouseTarget() *endpoint.ClickhouseTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_ClickhouseTarget); ok {
		return x.ClickhouseTarget
	}
	return nil
}

func (x *EndpointSettings) GetYdbTarget() *endpoint.YdbTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_YdbTarget); ok {
		return x.YdbTarget
	}
	return nil
}

func (x *EndpointSettings) GetKafkaTarget() *endpoint.KafkaTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_KafkaTarget); ok {
		return x.KafkaTarget
	}
	return nil
}

func (x *EndpointSettings) GetMongoTarget() *endpoint.MongoTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_MongoTarget); ok {
		return x.MongoTarget
	}
	return nil
}

func (x *EndpointSettings) GetYdsTarget() *endpoint.YDSTarget {
	if x, ok := x.GetSettings().(*EndpointSettings_YdsTarget); ok {
		return x.YdsTarget
	}
	return nil
}

type isEndpointSettings_Settings interface {
	isEndpointSettings_Settings()
}

type EndpointSettings_MysqlSource struct {
	MysqlSource *endpoint.MysqlSource `protobuf:"bytes,1,opt,name=mysql_source,json=mysqlSource,proto3,oneof"`
}

type EndpointSettings_PostgresSource struct {
	PostgresSource *endpoint.PostgresSource `protobuf:"bytes,2,opt,name=postgres_source,json=postgresSource,proto3,oneof"`
}

type EndpointSettings_YdbSource struct {
	YdbSource *endpoint.YdbSource `protobuf:"bytes,3,opt,name=ydb_source,json=ydbSource,proto3,oneof"`
}

type EndpointSettings_YdsSource struct {
	YdsSource *endpoint.YDSSource `protobuf:"bytes,7,opt,name=yds_source,json=ydsSource,proto3,oneof"`
}

type EndpointSettings_KafkaSource struct {
	KafkaSource *endpoint.KafkaSource `protobuf:"bytes,8,opt,name=kafka_source,json=kafkaSource,proto3,oneof"`
}

type EndpointSettings_MongoSource struct {
	MongoSource *endpoint.MongoSource `protobuf:"bytes,9,opt,name=mongo_source,json=mongoSource,proto3,oneof"`
}

type EndpointSettings_ClickhouseSource struct {
	ClickhouseSource *endpoint.ClickhouseSource `protobuf:"bytes,16,opt,name=clickhouse_source,json=clickhouseSource,proto3,oneof"`
}

type EndpointSettings_MysqlTarget struct {
	MysqlTarget *endpoint.MysqlTarget `protobuf:"bytes,101,opt,name=mysql_target,json=mysqlTarget,proto3,oneof"`
}

type EndpointSettings_PostgresTarget struct {
	PostgresTarget *endpoint.PostgresTarget `protobuf:"bytes,102,opt,name=postgres_target,json=postgresTarget,proto3,oneof"`
}

type EndpointSettings_ClickhouseTarget struct {
	ClickhouseTarget *endpoint.ClickhouseTarget `protobuf:"bytes,104,opt,name=clickhouse_target,json=clickhouseTarget,proto3,oneof"`
}

type EndpointSettings_YdbTarget struct {
	YdbTarget *endpoint.YdbTarget `protobuf:"bytes,105,opt,name=ydb_target,json=ydbTarget,proto3,oneof"`
}

type EndpointSettings_KafkaTarget struct {
	KafkaTarget *endpoint.KafkaTarget `protobuf:"bytes,110,opt,name=kafka_target,json=kafkaTarget,proto3,oneof"`
}

type EndpointSettings_MongoTarget struct {
	MongoTarget *endpoint.MongoTarget `protobuf:"bytes,111,opt,name=mongo_target,json=mongoTarget,proto3,oneof"`
}

type EndpointSettings_YdsTarget struct {
	YdsTarget *endpoint.YDSTarget `protobuf:"bytes,150,opt,name=yds_target,json=ydsTarget,proto3,oneof"`
}

func (*EndpointSettings_MysqlSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_PostgresSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_YdbSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_YdsSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_KafkaSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_MongoSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_ClickhouseSource) isEndpointSettings_Settings() {}

func (*EndpointSettings_MysqlTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_PostgresTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_ClickhouseTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_YdbTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_KafkaTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_MongoTarget) isEndpointSettings_Settings() {}

func (*EndpointSettings_YdsTarget) isEndpointSettings_Settings() {}

var File_yandex_cloud_datatransfer_v1_endpoint_proto protoreflect.FileDescriptor

var file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x79, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x79,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x34, 0x22, 0xba, 0x0a, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x0c,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x79, 0x64, 0x62, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x59, 0x64, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x79, 0x64, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x79, 0x64,
	0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x59, 0x44, 0x53, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x79, 0x64, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x0c, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x66, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x60, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x66, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x51, 0x0a, 0x0a, 0x79, 0x64,
	0x62, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x59, 0x64, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x79, 0x64, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x57, 0x0a,
	0x0c, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x52, 0x0a, 0x0a, 0x79, 0x64, 0x73, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x96, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x59, 0x44, 0x53,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x09, 0x79, 0x64, 0x73, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x10, 0x4a, 0x04, 0x08, 0x11, 0x10,
	0x65, 0x4a, 0x04, 0x08, 0x67, 0x10, 0x68, 0x4a, 0x04, 0x08, 0x6a, 0x10, 0x6e, 0x4a, 0x05, 0x08,
	0x70, 0x10, 0x96, 0x01, 0x42, 0x71, 0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescData = file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDesc
)

func file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescData)
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_datatransfer_v1_endpoint_proto_goTypes = []interface{}{
	(*Endpoint)(nil),                  // 0: yandex.cloud.datatransfer.v1.Endpoint
	(*EndpointSettings)(nil),          // 1: yandex.cloud.datatransfer.v1.EndpointSettings
	nil,                               // 2: yandex.cloud.datatransfer.v1.Endpoint.LabelsEntry
	(*endpoint.MysqlSource)(nil),      // 3: yandex.cloud.datatransfer.v1.endpoint.MysqlSource
	(*endpoint.PostgresSource)(nil),   // 4: yandex.cloud.datatransfer.v1.endpoint.PostgresSource
	(*endpoint.YdbSource)(nil),        // 5: yandex.cloud.datatransfer.v1.endpoint.YdbSource
	(*endpoint.YDSSource)(nil),        // 6: yandex.cloud.datatransfer.v1.endpoint.YDSSource
	(*endpoint.KafkaSource)(nil),      // 7: yandex.cloud.datatransfer.v1.endpoint.KafkaSource
	(*endpoint.MongoSource)(nil),      // 8: yandex.cloud.datatransfer.v1.endpoint.MongoSource
	(*endpoint.ClickhouseSource)(nil), // 9: yandex.cloud.datatransfer.v1.endpoint.ClickhouseSource
	(*endpoint.MysqlTarget)(nil),      // 10: yandex.cloud.datatransfer.v1.endpoint.MysqlTarget
	(*endpoint.PostgresTarget)(nil),   // 11: yandex.cloud.datatransfer.v1.endpoint.PostgresTarget
	(*endpoint.ClickhouseTarget)(nil), // 12: yandex.cloud.datatransfer.v1.endpoint.ClickhouseTarget
	(*endpoint.YdbTarget)(nil),        // 13: yandex.cloud.datatransfer.v1.endpoint.YdbTarget
	(*endpoint.KafkaTarget)(nil),      // 14: yandex.cloud.datatransfer.v1.endpoint.KafkaTarget
	(*endpoint.MongoTarget)(nil),      // 15: yandex.cloud.datatransfer.v1.endpoint.MongoTarget
	(*endpoint.YDSTarget)(nil),        // 16: yandex.cloud.datatransfer.v1.endpoint.YDSTarget
}
var file_yandex_cloud_datatransfer_v1_endpoint_proto_depIdxs = []int32{
	2,  // 0: yandex.cloud.datatransfer.v1.Endpoint.labels:type_name -> yandex.cloud.datatransfer.v1.Endpoint.LabelsEntry
	1,  // 1: yandex.cloud.datatransfer.v1.Endpoint.settings:type_name -> yandex.cloud.datatransfer.v1.EndpointSettings
	3,  // 2: yandex.cloud.datatransfer.v1.EndpointSettings.mysql_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.MysqlSource
	4,  // 3: yandex.cloud.datatransfer.v1.EndpointSettings.postgres_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.PostgresSource
	5,  // 4: yandex.cloud.datatransfer.v1.EndpointSettings.ydb_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdbSource
	6,  // 5: yandex.cloud.datatransfer.v1.EndpointSettings.yds_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.YDSSource
	7,  // 6: yandex.cloud.datatransfer.v1.EndpointSettings.kafka_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.KafkaSource
	8,  // 7: yandex.cloud.datatransfer.v1.EndpointSettings.mongo_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoSource
	9,  // 8: yandex.cloud.datatransfer.v1.EndpointSettings.clickhouse_source:type_name -> yandex.cloud.datatransfer.v1.endpoint.ClickhouseSource
	10, // 9: yandex.cloud.datatransfer.v1.EndpointSettings.mysql_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.MysqlTarget
	11, // 10: yandex.cloud.datatransfer.v1.EndpointSettings.postgres_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.PostgresTarget
	12, // 11: yandex.cloud.datatransfer.v1.EndpointSettings.clickhouse_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.ClickhouseTarget
	13, // 12: yandex.cloud.datatransfer.v1.EndpointSettings.ydb_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdbTarget
	14, // 13: yandex.cloud.datatransfer.v1.EndpointSettings.kafka_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.KafkaTarget
	15, // 14: yandex.cloud.datatransfer.v1.EndpointSettings.mongo_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoTarget
	16, // 15: yandex.cloud.datatransfer.v1.EndpointSettings.yds_target:type_name -> yandex.cloud.datatransfer.v1.endpoint.YDSTarget
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointSettings); i {
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
	file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EndpointSettings_MysqlSource)(nil),
		(*EndpointSettings_PostgresSource)(nil),
		(*EndpointSettings_YdbSource)(nil),
		(*EndpointSettings_YdsSource)(nil),
		(*EndpointSettings_KafkaSource)(nil),
		(*EndpointSettings_MongoSource)(nil),
		(*EndpointSettings_ClickhouseSource)(nil),
		(*EndpointSettings_MysqlTarget)(nil),
		(*EndpointSettings_PostgresTarget)(nil),
		(*EndpointSettings_ClickhouseTarget)(nil),
		(*EndpointSettings_YdbTarget)(nil),
		(*EndpointSettings_KafkaTarget)(nil),
		(*EndpointSettings_MongoTarget)(nil),
		(*EndpointSettings_YdsTarget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_proto_rawDesc = nil
	file_yandex_cloud_datatransfer_v1_endpoint_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_proto_depIdxs = nil
}
