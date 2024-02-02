// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/datatransfer/v1/endpoint/parsers.proto

package endpoint

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

type Parser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Parser:
	//
	//	*Parser_JsonParser
	//	*Parser_AuditTrailsV1Parser
	//	*Parser_CloudLoggingParser
	//	*Parser_TskvParser
	Parser isParser_Parser `protobuf_oneof:"parser"`
}

func (x *Parser) Reset() {
	*x = Parser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parser) ProtoMessage() {}

func (x *Parser) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parser.ProtoReflect.Descriptor instead.
func (*Parser) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{0}
}

func (m *Parser) GetParser() isParser_Parser {
	if m != nil {
		return m.Parser
	}
	return nil
}

func (x *Parser) GetJsonParser() *GenericParserCommon {
	if x, ok := x.GetParser().(*Parser_JsonParser); ok {
		return x.JsonParser
	}
	return nil
}

func (x *Parser) GetAuditTrailsV1Parser() *AuditTrailsV1Parser {
	if x, ok := x.GetParser().(*Parser_AuditTrailsV1Parser); ok {
		return x.AuditTrailsV1Parser
	}
	return nil
}

func (x *Parser) GetCloudLoggingParser() *CloudLoggingParser {
	if x, ok := x.GetParser().(*Parser_CloudLoggingParser); ok {
		return x.CloudLoggingParser
	}
	return nil
}

func (x *Parser) GetTskvParser() *GenericParserCommon {
	if x, ok := x.GetParser().(*Parser_TskvParser); ok {
		return x.TskvParser
	}
	return nil
}

type isParser_Parser interface {
	isParser_Parser()
}

type Parser_JsonParser struct {
	JsonParser *GenericParserCommon `protobuf:"bytes,1,opt,name=json_parser,json=jsonParser,proto3,oneof"`
}

type Parser_AuditTrailsV1Parser struct {
	AuditTrailsV1Parser *AuditTrailsV1Parser `protobuf:"bytes,2,opt,name=audit_trails_v1_parser,json=auditTrailsV1Parser,proto3,oneof"`
}

type Parser_CloudLoggingParser struct {
	CloudLoggingParser *CloudLoggingParser `protobuf:"bytes,4,opt,name=cloud_logging_parser,json=cloudLoggingParser,proto3,oneof"`
}

type Parser_TskvParser struct {
	TskvParser *GenericParserCommon `protobuf:"bytes,6,opt,name=tskv_parser,json=tskvParser,proto3,oneof"`
}

func (*Parser_JsonParser) isParser_Parser() {}

func (*Parser_AuditTrailsV1Parser) isParser_Parser() {}

func (*Parser_CloudLoggingParser) isParser_Parser() {}

func (*Parser_TskvParser) isParser_Parser() {}

type GenericParserCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSchema *DataSchema `protobuf:"bytes,1,opt,name=data_schema,json=dataSchema,proto3" json:"data_schema,omitempty"`
	// Allow null keys, if no - null keys will be putted to unparsed data
	NullKeysAllowed bool `protobuf:"varint,2,opt,name=null_keys_allowed,json=nullKeysAllowed,proto3" json:"null_keys_allowed,omitempty"`
	// Will add _rest column for all unknown fields
	AddRestColumn bool `protobuf:"varint,3,opt,name=add_rest_column,json=addRestColumn,proto3" json:"add_rest_column,omitempty"`
	// Unescape string values
	UnescapeStringValues bool `protobuf:"varint,7,opt,name=unescape_string_values,json=unescapeStringValues,proto3" json:"unescape_string_values,omitempty"`
}

func (x *GenericParserCommon) Reset() {
	*x = GenericParserCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericParserCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericParserCommon) ProtoMessage() {}

func (x *GenericParserCommon) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericParserCommon.ProtoReflect.Descriptor instead.
func (*GenericParserCommon) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{1}
}

func (x *GenericParserCommon) GetDataSchema() *DataSchema {
	if x != nil {
		return x.DataSchema
	}
	return nil
}

func (x *GenericParserCommon) GetNullKeysAllowed() bool {
	if x != nil {
		return x.NullKeysAllowed
	}
	return false
}

func (x *GenericParserCommon) GetAddRestColumn() bool {
	if x != nil {
		return x.AddRestColumn
	}
	return false
}

func (x *GenericParserCommon) GetUnescapeStringValues() bool {
	if x != nil {
		return x.UnescapeStringValues
	}
	return false
}

type AuditTrailsV1Parser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuditTrailsV1Parser) Reset() {
	*x = AuditTrailsV1Parser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditTrailsV1Parser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditTrailsV1Parser) ProtoMessage() {}

func (x *AuditTrailsV1Parser) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditTrailsV1Parser.ProtoReflect.Descriptor instead.
func (*AuditTrailsV1Parser) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{2}
}

type CloudLoggingParser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloudLoggingParser) Reset() {
	*x = CloudLoggingParser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudLoggingParser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudLoggingParser) ProtoMessage() {}

func (x *CloudLoggingParser) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudLoggingParser.ProtoReflect.Descriptor instead.
func (*CloudLoggingParser) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{3}
}

var File_yandex_cloud_datatransfer_v1_endpoint_parsers_proto protoreflect.FileDescriptor

var file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x32, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x03, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0b, 0x6a,
	0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a,
	0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x16, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x76, 0x31, 0x5f, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x56, 0x31,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x13, 0x61, 0x75, 0x64, 0x69, 0x74, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x73, 0x56, 0x31, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x6d, 0x0a,
	0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0b,
	0x74, 0x73, 0x6b, 0x76, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x0a, 0x74, 0x73, 0x6b, 0x76, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x05, 0x10,
	0x06, 0x22, 0xf9, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2a, 0x0a,
	0x11, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e, 0x75, 0x6c, 0x6c, 0x4b, 0x65,
	0x79, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x64, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x6e, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x14, 0x75, 0x6e, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x07, 0x22, 0x15, 0x0a,
	0x13, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x56, 0x31, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x42, 0xa7, 0x01, 0x0a, 0x29, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0xaa, 0x02, 0x25, 0x59,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescData = file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDesc
)

func file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescData)
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_goTypes = []interface{}{
	(*Parser)(nil),              // 0: yandex.cloud.datatransfer.v1.endpoint.Parser
	(*GenericParserCommon)(nil), // 1: yandex.cloud.datatransfer.v1.endpoint.GenericParserCommon
	(*AuditTrailsV1Parser)(nil), // 2: yandex.cloud.datatransfer.v1.endpoint.AuditTrailsV1Parser
	(*CloudLoggingParser)(nil),  // 3: yandex.cloud.datatransfer.v1.endpoint.CloudLoggingParser
	(*DataSchema)(nil),          // 4: yandex.cloud.datatransfer.v1.endpoint.DataSchema
}
var file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.datatransfer.v1.endpoint.Parser.json_parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.GenericParserCommon
	2, // 1: yandex.cloud.datatransfer.v1.endpoint.Parser.audit_trails_v1_parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.AuditTrailsV1Parser
	3, // 2: yandex.cloud.datatransfer.v1.endpoint.Parser.cloud_logging_parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.CloudLoggingParser
	1, // 3: yandex.cloud.datatransfer.v1.endpoint.Parser.tskv_parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.GenericParserCommon
	4, // 4: yandex.cloud.datatransfer.v1.endpoint.GenericParserCommon.data_schema:type_name -> yandex.cloud.datatransfer.v1.endpoint.DataSchema
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_parsers_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parser); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericParserCommon); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditTrailsV1Parser); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudLoggingParser); i {
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
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Parser_JsonParser)(nil),
		(*Parser_AuditTrailsV1Parser)(nil),
		(*Parser_CloudLoggingParser)(nil),
		(*Parser_TskvParser)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_parsers_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_rawDesc = nil
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_depIdxs = nil
}
