// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/config/mysql5_7.proto

package mysql // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1alpha/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Options and structure of `MysqlConfig5_7` reflects MySQL 5.7 configuration file
type MysqlConfig5_7 struct {
	// Size of the InnoDB buffer pool used for caching table and index data.
	//
	// For details, see [MySQL documentation for the parameter](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).
	InnodbBufferPoolSize *wrappers.Int64Value `protobuf:"bytes,1,opt,name=innodb_buffer_pool_size,json=innodbBufferPoolSize,proto3" json:"innodb_buffer_pool_size,omitempty"`
	// The maximum permitted number of simultaneous client connections.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_max_connections).
	MaxConnections *wrappers.Int64Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// Time that it takes to process a query before it is considered slow.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_long_query_time).
	LongQueryTime        *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=long_query_time,json=longQueryTime,proto3" json:"long_query_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MysqlConfig5_7) Reset()         { *m = MysqlConfig5_7{} }
func (m *MysqlConfig5_7) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig5_7) ProtoMessage()    {}
func (*MysqlConfig5_7) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysql5_7_1f2771aee9a9a4ee, []int{0}
}
func (m *MysqlConfig5_7) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig5_7.Unmarshal(m, b)
}
func (m *MysqlConfig5_7) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig5_7.Marshal(b, m, deterministic)
}
func (dst *MysqlConfig5_7) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig5_7.Merge(dst, src)
}
func (m *MysqlConfig5_7) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig5_7.Size(m)
}
func (m *MysqlConfig5_7) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig5_7.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig5_7 proto.InternalMessageInfo

func (m *MysqlConfig5_7) GetInnodbBufferPoolSize() *wrappers.Int64Value {
	if m != nil {
		return m.InnodbBufferPoolSize
	}
	return nil
}

func (m *MysqlConfig5_7) GetMaxConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *MysqlConfig5_7) GetLongQueryTime() *wrappers.DoubleValue {
	if m != nil {
		return m.LongQueryTime
	}
	return nil
}

type MysqlConfigSet5_7 struct {
	// Effective settings for a MySQL 5.7 cluster (a combination of settings defined
	// in [user_config] and [default_config]).
	EffectiveConfig *MysqlConfig5_7 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a MySQL 5.7 cluster.
	UserConfig *MysqlConfig5_7 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a MySQL 5.7 cluster.
	DefaultConfig        *MysqlConfig5_7 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MysqlConfigSet5_7) Reset()         { *m = MysqlConfigSet5_7{} }
func (m *MysqlConfigSet5_7) String() string { return proto.CompactTextString(m) }
func (*MysqlConfigSet5_7) ProtoMessage()    {}
func (*MysqlConfigSet5_7) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysql5_7_1f2771aee9a9a4ee, []int{1}
}
func (m *MysqlConfigSet5_7) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfigSet5_7.Unmarshal(m, b)
}
func (m *MysqlConfigSet5_7) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfigSet5_7.Marshal(b, m, deterministic)
}
func (dst *MysqlConfigSet5_7) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfigSet5_7.Merge(dst, src)
}
func (m *MysqlConfigSet5_7) XXX_Size() int {
	return xxx_messageInfo_MysqlConfigSet5_7.Size(m)
}
func (m *MysqlConfigSet5_7) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfigSet5_7.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfigSet5_7 proto.InternalMessageInfo

func (m *MysqlConfigSet5_7) GetEffectiveConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.EffectiveConfig
	}
	return nil
}

func (m *MysqlConfigSet5_7) GetUserConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.UserConfig
	}
	return nil
}

func (m *MysqlConfigSet5_7) GetDefaultConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*MysqlConfig5_7)(nil), "yandex.cloud.mdb.mysql.v1alpha.config.MysqlConfig5_7")
	proto.RegisterType((*MysqlConfigSet5_7)(nil), "yandex.cloud.mdb.mysql.v1alpha.config.MysqlConfigSet5_7")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/config/mysql5_7.proto", fileDescriptor_mysql5_7_1f2771aee9a9a4ee)
}

var fileDescriptor_mysql5_7_1f2771aee9a9a4ee = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd2, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0x06, 0x70, 0xba, 0x05, 0xd1, 0xa9, 0xbb, 0xab, 0x41, 0x70, 0xa9, 0x7f, 0x90, 0x82, 0xe0,
	0xcd, 0xce, 0x24, 0xeb, 0xae, 0x2d, 0x88, 0x5e, 0x6c, 0x7b, 0xe3, 0x85, 0x68, 0x53, 0xe9, 0x85,
	0x08, 0xe3, 0x24, 0x39, 0x49, 0x07, 0x66, 0xe6, 0xa4, 0x49, 0x66, 0xdd, 0xed, 0x3b, 0xf8, 0x24,
	0x3e, 0x48, 0xdf, 0xa9, 0x57, 0x92, 0x99, 0xac, 0x5a, 0xbc, 0x68, 0x61, 0x6f, 0x4f, 0xf2, 0xfd,
	0x4e, 0xf8, 0x72, 0xc8, 0x74, 0x25, 0x4c, 0x06, 0x4b, 0x96, 0x2a, 0xb4, 0x19, 0xd3, 0x59, 0xc2,
	0xf4, 0xaa, 0x3e, 0x57, 0x6c, 0x11, 0x09, 0x55, 0x9e, 0x09, 0x96, 0xa2, 0xc9, 0x65, 0xe1, 0x87,
	0x33, 0xbe, 0x4f, 0xcb, 0x0a, 0x1b, 0x0c, 0x5e, 0xfa, 0x14, 0x75, 0x29, 0xaa, 0xb3, 0x84, 0xba,
	0x17, 0x68, 0x97, 0xa2, 0x3e, 0xb5, 0xfb, 0xbc, 0x40, 0x2c, 0x14, 0x30, 0x17, 0x4a, 0x6c, 0xce,
	0x7e, 0x54, 0xa2, 0x2c, 0xa1, 0xaa, 0x3d, 0xb3, 0xfb, 0xec, 0xda, 0xf2, 0x85, 0x50, 0x32, 0x13,
	0x8d, 0x44, 0xe3, 0x1f, 0xef, 0xfd, 0xec, 0x91, 0xc1, 0xc7, 0xd6, 0x3d, 0x74, 0xdc, 0x8c, 0xef,
	0x07, 0x82, 0x3c, 0x96, 0xc6, 0x60, 0x96, 0xf0, 0xc4, 0xe6, 0x39, 0x54, 0xbc, 0x44, 0x54, 0xbc,
	0x96, 0x17, 0x30, 0xda, 0x7a, 0xb1, 0xf5, 0x6a, 0x67, 0xf2, 0x84, 0xfa, 0x9d, 0x74, 0xbd, 0x93,
	0x7e, 0x30, 0xcd, 0x9b, 0xe9, 0xa9, 0x50, 0x16, 0xe6, 0xfd, 0xab, 0xcb, 0xe8, 0xde, 0xfb, 0x77,
	0xb3, 0xc9, 0x74, 0x72, 0x70, 0x10, 0xc6, 0x8f, 0x3c, 0x35, 0x77, 0xd2, 0x67, 0x44, 0x75, 0x22,
	0x2f, 0x20, 0x88, 0xc9, 0x50, 0x8b, 0x25, 0x4f, 0xd1, 0x18, 0x48, 0xdb, 0xaf, 0xa9, 0x47, 0xbd,
	0x9b, 0xe9, 0xfb, 0x57, 0x97, 0xd1, 0xdd, 0x28, 0x1c, 0x47, 0x61, 0x18, 0x86, 0xf1, 0x40, 0x8b,
	0xe5, 0xe1, 0x5f, 0x20, 0x38, 0x22, 0x43, 0x85, 0xa6, 0xe0, 0xe7, 0x16, 0xaa, 0x15, 0x6f, 0xa4,
	0x86, 0xd1, 0xb6, 0x33, 0x9f, 0xfe, 0x67, 0x1e, 0xa1, 0x4d, 0x14, 0x38, 0x34, 0xee, 0xb7, 0xa1,
	0xe3, 0x36, 0xf3, 0x45, 0x6a, 0xd8, 0xfb, 0xd5, 0x23, 0x0f, 0xff, 0xe9, 0xe3, 0x04, 0x9a, 0xb6,
	0x92, 0xef, 0xe4, 0x01, 0xe4, 0x79, 0xbb, 0x69, 0x01, 0xdc, 0x17, 0xdf, 0x75, 0x31, 0xa3, 0xb7,
	0xfa, 0x4d, 0xf4, 0x7a, 0xc7, 0xf1, 0xf0, 0x0f, 0xe7, 0x67, 0xc1, 0x29, 0xd9, 0xb1, 0x35, 0x54,
	0x6b, 0xbc, 0xb7, 0x09, 0x4e, 0x5a, 0xa9, 0x73, 0xbf, 0x91, 0x41, 0x06, 0xb9, 0xb0, 0xaa, 0x59,
	0xd3, 0xdb, 0x9b, 0xd0, 0xfd, 0x0e, 0xf3, 0x93, 0xf9, 0xf1, 0xd7, 0x4f, 0x85, 0x6c, 0xce, 0x6c,
	0x42, 0x53, 0xd4, 0xcc, 0x8b, 0x63, 0x7f, 0x69, 0x05, 0x8e, 0x0b, 0x30, 0xae, 0x72, 0x76, 0xab,
	0xfb, 0x7f, 0xeb, 0x86, 0xc9, 0x1d, 0x17, 0x79, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x89, 0x30,
	0x81, 0x5e, 0x35, 0x03, 0x00, 0x00,
}
