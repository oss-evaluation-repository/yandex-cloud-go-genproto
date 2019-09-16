// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/database_service.proto

package mysql

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetDatabaseRequest struct {
	// ID of the MySQL cluster that the database belongs to.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MySQL database to return.
	// To get the name of the database use a [DatabaseService.List] request.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatabaseRequest) Reset()         { *m = GetDatabaseRequest{} }
func (m *GetDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseRequest) ProtoMessage()    {}
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{0}
}

func (m *GetDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatabaseRequest.Unmarshal(m, b)
}
func (m *GetDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatabaseRequest.Marshal(b, m, deterministic)
}
func (m *GetDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatabaseRequest.Merge(m, src)
}
func (m *GetDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatabaseRequest.Size(m)
}
func (m *GetDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatabaseRequest proto.InternalMessageInfo

func (m *GetDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *GetDatabaseRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type ListDatabasesRequest struct {
	// ID of the MySQL cluster to list databases in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListDatabasesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, Set [page_token] to the [ListDatabasesResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDatabasesRequest) Reset()         { *m = ListDatabasesRequest{} }
func (m *ListDatabasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesRequest) ProtoMessage()    {}
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{1}
}

func (m *ListDatabasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDatabasesRequest.Unmarshal(m, b)
}
func (m *ListDatabasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDatabasesRequest.Marshal(b, m, deterministic)
}
func (m *ListDatabasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDatabasesRequest.Merge(m, src)
}
func (m *ListDatabasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDatabasesRequest.Size(m)
}
func (m *ListDatabasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDatabasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDatabasesRequest proto.InternalMessageInfo

func (m *ListDatabasesRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListDatabasesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDatabasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListDatabasesResponse struct {
	// List of MySQL databases.
	Databases []*Database `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListDatabasesRequest.page_size], use the [next_page_token] as the value
	// for the [ListDatabasesRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDatabasesResponse) Reset()         { *m = ListDatabasesResponse{} }
func (m *ListDatabasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesResponse) ProtoMessage()    {}
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{2}
}

func (m *ListDatabasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDatabasesResponse.Unmarshal(m, b)
}
func (m *ListDatabasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDatabasesResponse.Marshal(b, m, deterministic)
}
func (m *ListDatabasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDatabasesResponse.Merge(m, src)
}
func (m *ListDatabasesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDatabasesResponse.Size(m)
}
func (m *ListDatabasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDatabasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDatabasesResponse proto.InternalMessageInfo

func (m *ListDatabasesResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

func (m *ListDatabasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateDatabaseRequest struct {
	// ID of the MySQL cluster to create a database in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Configuration of the database to create.
	DatabaseSpec         *DatabaseSpec `protobuf:"bytes,2,opt,name=database_spec,json=databaseSpec,proto3" json:"database_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateDatabaseRequest) Reset()         { *m = CreateDatabaseRequest{} }
func (m *CreateDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseRequest) ProtoMessage()    {}
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{3}
}

func (m *CreateDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatabaseRequest.Unmarshal(m, b)
}
func (m *CreateDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatabaseRequest.Marshal(b, m, deterministic)
}
func (m *CreateDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatabaseRequest.Merge(m, src)
}
func (m *CreateDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDatabaseRequest.Size(m)
}
func (m *CreateDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatabaseRequest proto.InternalMessageInfo

func (m *CreateDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateDatabaseRequest) GetDatabaseSpec() *DatabaseSpec {
	if m != nil {
		return m.DatabaseSpec
	}
	return nil
}

type CreateDatabaseMetadata struct {
	// ID of the MySQL cluster where a database is being created.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MySQL database that is being created.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDatabaseMetadata) Reset()         { *m = CreateDatabaseMetadata{} }
func (m *CreateDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseMetadata) ProtoMessage()    {}
func (*CreateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{4}
}

func (m *CreateDatabaseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatabaseMetadata.Unmarshal(m, b)
}
func (m *CreateDatabaseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatabaseMetadata.Marshal(b, m, deterministic)
}
func (m *CreateDatabaseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatabaseMetadata.Merge(m, src)
}
func (m *CreateDatabaseMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateDatabaseMetadata.Size(m)
}
func (m *CreateDatabaseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatabaseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatabaseMetadata proto.InternalMessageInfo

func (m *CreateDatabaseMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateDatabaseMetadata) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type DeleteDatabaseRequest struct {
	// ID of the MySQL cluster to delete a database in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the database to delete.
	// To get the name of the database, use a [DatabaseService.List] request.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDatabaseRequest) Reset()         { *m = DeleteDatabaseRequest{} }
func (m *DeleteDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDatabaseRequest) ProtoMessage()    {}
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{5}
}

func (m *DeleteDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDatabaseRequest.Unmarshal(m, b)
}
func (m *DeleteDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDatabaseRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDatabaseRequest.Merge(m, src)
}
func (m *DeleteDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDatabaseRequest.Size(m)
}
func (m *DeleteDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDatabaseRequest proto.InternalMessageInfo

func (m *DeleteDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteDatabaseRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type DeleteDatabaseMetadata struct {
	// ID of the MySQL cluster where a database is being deleted.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MySQL database that is being deleted.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDatabaseMetadata) Reset()         { *m = DeleteDatabaseMetadata{} }
func (m *DeleteDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteDatabaseMetadata) ProtoMessage()    {}
func (*DeleteDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e9c55fa14d67b3, []int{6}
}

func (m *DeleteDatabaseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDatabaseMetadata.Unmarshal(m, b)
}
func (m *DeleteDatabaseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDatabaseMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteDatabaseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDatabaseMetadata.Merge(m, src)
}
func (m *DeleteDatabaseMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteDatabaseMetadata.Size(m)
}
func (m *DeleteDatabaseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDatabaseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDatabaseMetadata proto.InternalMessageInfo

func (m *DeleteDatabaseMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteDatabaseMetadata) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.GetDatabaseRequest")
	proto.RegisterType((*ListDatabasesRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListDatabasesRequest")
	proto.RegisterType((*ListDatabasesResponse)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListDatabasesResponse")
	proto.RegisterType((*CreateDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.CreateDatabaseRequest")
	proto.RegisterType((*CreateDatabaseMetadata)(nil), "yandex.cloud.mdb.mysql.v1alpha.CreateDatabaseMetadata")
	proto.RegisterType((*DeleteDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.DeleteDatabaseRequest")
	proto.RegisterType((*DeleteDatabaseMetadata)(nil), "yandex.cloud.mdb.mysql.v1alpha.DeleteDatabaseMetadata")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/database_service.proto", fileDescriptor_55e9c55fa14d67b3)
}

var fileDescriptor_55e9c55fa14d67b3 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x4f, 0x13, 0x4f,
	0x1c, 0xce, 0x50, 0xfe, 0x0d, 0x1d, 0xe0, 0x4f, 0x32, 0xb1, 0xa4, 0x69, 0x84, 0xe0, 0x9a, 0x60,
	0x53, 0xdd, 0xb7, 0x22, 0x44, 0x05, 0x4c, 0x2c, 0x08, 0x68, 0x14, 0x4d, 0x31, 0x31, 0x41, 0x4c,
	0x33, 0xed, 0x8e, 0xcb, 0xc6, 0x7d, 0xa3, 0x33, 0x25, 0xbc, 0x84, 0x83, 0x1e, 0x34, 0x72, 0x35,
	0xf1, 0xe6, 0x97, 0x40, 0xbf, 0x03, 0x24, 0xde, 0xf0, 0x2b, 0x18, 0xe3, 0xd9, 0xa3, 0x27, 0xb3,
	0x33, 0xed, 0xb6, 0x0b, 0x85, 0x56, 0xe0, 0xb6, 0x3b, 0xbf, 0xe7, 0x99, 0x79, 0x9e, 0x99, 0xe7,
	0x37, 0x03, 0xc7, 0x37, 0xb1, 0x6b, 0x90, 0x0d, 0xb5, 0x6c, 0x7b, 0x55, 0x43, 0x75, 0x8c, 0x92,
	0xea, 0x6c, 0xd2, 0x35, 0x5b, 0x5d, 0xd7, 0xb1, 0xed, 0xaf, 0x62, 0xd5, 0xc0, 0x0c, 0x97, 0x30,
	0x25, 0x45, 0x4a, 0x2a, 0xeb, 0x56, 0x99, 0x28, 0x7e, 0xc5, 0x63, 0x1e, 0x1a, 0x16, 0x34, 0x85,
	0xd3, 0x14, 0xc7, 0x28, 0x29, 0x9c, 0xa6, 0xd4, 0x68, 0xe9, 0xcb, 0xa6, 0xe7, 0x99, 0x36, 0x51,
	0xb1, 0x6f, 0xa9, 0xd8, 0x75, 0x3d, 0x86, 0x99, 0xe5, 0xb9, 0x54, 0xb0, 0xd3, 0x23, 0x91, 0x45,
	0x03, 0x8c, 0xe7, 0x93, 0x0a, 0x87, 0xd4, 0x10, 0xa3, 0x11, 0x44, 0x58, 0x3d, 0x86, 0x1b, 0x8a,
	0xe0, 0xd6, 0xb1, 0x6d, 0x19, 0xcd, 0x65, 0xb9, 0x43, 0x77, 0x02, 0x2e, 0xbd, 0x03, 0x10, 0xcd,
	0x13, 0x36, 0x5b, 0x1b, 0x2d, 0x90, 0xb5, 0x2a, 0xa1, 0x0c, 0x5d, 0x87, 0xb0, 0x6c, 0x57, 0x29,
	0x23, 0x95, 0xa2, 0x65, 0xa4, 0xc0, 0x08, 0xc8, 0x24, 0xf2, 0x7d, 0xbf, 0xf6, 0x75, 0xb0, 0x7b,
	0xa0, 0x77, 0x4f, 0x4d, 0x8f, 0x6b, 0x85, 0x44, 0xad, 0xfe, 0xc0, 0x40, 0x33, 0xb0, 0x3f, 0xdc,
	0x33, 0x17, 0x3b, 0x24, 0xd5, 0xc5, 0xf1, 0xc3, 0x01, 0xfe, 0xf7, 0xbe, 0xfe, 0xff, 0x0b, 0x2c,
	0x6f, 0xdd, 0x93, 0x97, 0x35, 0xf9, 0x76, 0x51, 0x7e, 0x99, 0x15, 0x33, 0x4c, 0x8c, 0x15, 0xfa,
	0xea, 0xa4, 0x45, 0xec, 0x10, 0xe9, 0x13, 0x80, 0x97, 0x1e, 0x59, 0x34, 0x54, 0x42, 0xcf, 0x24,
	0xe5, 0x1a, 0x4c, 0xf8, 0xd8, 0x24, 0x45, 0x6a, 0x6d, 0x09, 0x19, 0xb1, 0x3c, 0xfc, 0xb3, 0xaf,
	0xc7, 0x35, 0x59, 0xd7, 0x34, 0xad, 0xd0, 0x13, 0x14, 0x97, 0xac, 0x2d, 0x82, 0x32, 0x10, 0x72,
	0x20, 0xf3, 0x5e, 0x13, 0x37, 0x15, 0xe3, 0xb3, 0x26, 0x76, 0x0f, 0xf4, 0xff, 0xa6, 0xa6, 0x75,
	0x4d, 0x2b, 0xf0, 0x59, 0x9e, 0x05, 0x35, 0xe9, 0x3d, 0x80, 0xc9, 0x23, 0xc2, 0xa8, 0xef, 0xb9,
	0x94, 0xa0, 0x39, 0x98, 0xa8, 0x5b, 0xa0, 0x29, 0x30, 0x12, 0xcb, 0xf4, 0xe6, 0x32, 0xca, 0xe9,
	0x29, 0x51, 0xc2, 0x8d, 0x6e, 0x50, 0xd1, 0x28, 0x1c, 0x70, 0xc9, 0x06, 0x2b, 0x36, 0x09, 0xe2,
	0x3b, 0x58, 0xe8, 0x0f, 0x86, 0x9f, 0x86, 0x4a, 0x3e, 0x03, 0x98, 0x9c, 0xa9, 0x10, 0xcc, 0xc8,
	0xb9, 0x8e, 0xeb, 0x79, 0xd3, 0x71, 0x51, 0x9f, 0x94, 0xf9, 0x62, 0xbd, 0xb9, 0x1b, 0x9d, 0x4a,
	0x5f, 0xf2, 0x49, 0x39, 0xdf, 0x1d, 0xcc, 0xde, 0x38, 0xc2, 0x60, 0x4c, 0x5a, 0x81, 0x83, 0x51,
	0x79, 0x8f, 0x09, 0xc3, 0x01, 0x02, 0x0d, 0x1d, 0xd7, 0xd7, 0xac, 0xe8, 0x6a, 0xcb, 0x00, 0x1d,
	0x09, 0xc8, 0x07, 0x00, 0x93, 0xb3, 0xc4, 0x26, 0xe7, 0x74, 0x7f, 0x21, 0x61, 0x5d, 0x81, 0x83,
	0x51, 0x29, 0x17, 0xe9, 0x34, 0xf7, 0x35, 0x0e, 0x07, 0xc2, 0xcd, 0x16, 0x77, 0x10, 0xfa, 0x02,
	0x60, 0x6c, 0x9e, 0x30, 0x94, 0x6b, 0x77, 0x4a, 0xc7, 0x9b, 0x39, 0xdd, 0x71, 0x28, 0xa5, 0xc5,
	0xb7, 0xdf, 0x7f, 0x7c, 0xec, 0x5a, 0x40, 0x73, 0xaa, 0x83, 0x5d, 0x6c, 0x12, 0x43, 0x8e, 0x5e,
	0x1e, 0x35, 0x23, 0x54, 0xdd, 0x6e, 0x98, 0xdc, 0x09, 0xaf, 0x14, 0xaa, 0x6e, 0x47, 0xcc, 0xed,
	0x04, 0xaa, 0xbb, 0x83, 0xde, 0x41, 0x37, 0xdb, 0x49, 0x68, 0xd5, 0xfa, 0xe9, 0xf1, 0x7f, 0x64,
	0x89, 0xbe, 0x94, 0xee, 0x72, 0x17, 0xb7, 0xd0, 0xc4, 0xd9, 0x5c, 0xa0, 0x6f, 0x00, 0xc6, 0x45,
	0x90, 0x51, 0x5b, 0x05, 0x2d, 0xfb, 0x31, 0x7d, 0x25, 0x4a, 0x6b, 0x5c, 0xe1, 0x4f, 0xea, 0x5f,
	0x92, 0xb9, 0x77, 0x98, 0x95, 0x4e, 0x6c, 0x98, 0x9e, 0xfa, 0x08, 0xb7, 0x32, 0x29, 0x9d, 0xd1,
	0xca, 0x1d, 0x90, 0x45, 0x3f, 0x01, 0x8c, 0x8b, 0xb0, 0xb6, 0x77, 0xd3, 0xb2, 0xbf, 0x3a, 0x71,
	0xf3, 0x06, 0xec, 0x1d, 0x66, 0xd5, 0x13, 0xbb, 0x22, 0x29, 0xde, 0x46, 0xf1, 0xe6, 0x94, 0xaa,
	0xaf, 0x94, 0xfb, 0x8e, 0xcf, 0x36, 0x45, 0xd8, 0xb2, 0x17, 0x14, 0xb6, 0xfc, 0xc3, 0xe5, 0x05,
	0xd3, 0x62, 0xab, 0xd5, 0x92, 0x52, 0xf6, 0x1c, 0x55, 0x48, 0x96, 0xc5, 0x33, 0x68, 0x7a, 0xb2,
	0x49, 0x5c, 0xbe, 0xba, 0x7a, 0xfa, 0xfb, 0x38, 0xc9, 0xff, 0x4a, 0x71, 0x8e, 0x1d, 0xfb, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x08, 0xc0, 0x8d, 0x2c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	// Returns the specified MySQL database.
	//
	// To get the list of available MySQL databases, make a [List] request.
	Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error)
	// Retrieves the list of MySQL databases in the specified cluster.
	List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error)
	// Creates a new MySQL database in the specified cluster.
	Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified MySQL database.
	Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type databaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseServiceClient(cc *grpc.ClientConn) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error) {
	out := new(Database)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error) {
	out := new(ListDatabasesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
type DatabaseServiceServer interface {
	// Returns the specified MySQL database.
	//
	// To get the list of available MySQL databases, make a [List] request.
	Get(context.Context, *GetDatabaseRequest) (*Database, error)
	// Retrieves the list of MySQL databases in the specified cluster.
	List(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
	// Creates a new MySQL database in the specified cluster.
	Create(context.Context, *CreateDatabaseRequest) (*operation.Operation, error)
	// Deletes the specified MySQL database.
	Delete(context.Context, *DeleteDatabaseRequest) (*operation.Operation, error)
}

// UnimplementedDatabaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (*UnimplementedDatabaseServiceServer) Get(ctx context.Context, req *GetDatabaseRequest) (*Database, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDatabaseServiceServer) List(ctx context.Context, req *ListDatabasesRequest) (*ListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDatabaseServiceServer) Create(ctx context.Context, req *CreateDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDatabaseServiceServer) Delete(ctx context.Context, req *DeleteDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterDatabaseServiceServer(s *grpc.Server, srv DatabaseServiceServer) {
	s.RegisterService(&_DatabaseService_serviceDesc, srv)
}

func _DatabaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Get(ctx, req.(*GetDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).List(ctx, req.(*ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Create(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Delete(ctx, req.(*DeleteDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mysql.v1alpha.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DatabaseService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DatabaseService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DatabaseService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatabaseService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mysql/v1alpha/database_service.proto",
}
