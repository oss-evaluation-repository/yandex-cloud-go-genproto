// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/ydb/v1/backup_service.proto

package ydb

import (
	context "context"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListPathsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the YDB backup.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListPaths requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListPaths
	// request to get the next page of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPathsRequest) Reset() {
	*x = ListPathsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPathsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPathsRequest) ProtoMessage() {}

func (x *ListPathsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPathsRequest.ProtoReflect.Descriptor instead.
func (*ListPathsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPathsRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *ListPathsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPathsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPathsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths         []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPathsResponse) Reset() {
	*x = ListPathsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPathsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPathsResponse) ProtoMessage() {}

func (x *ListPathsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPathsResponse.ProtoReflect.Descriptor instead.
func (*ListPathsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPathsResponse) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *ListPathsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the YDB backup.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *GetBackupRequest) Reset() {
	*x = GetBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackupRequest) ProtoMessage() {}

func (x *GetBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackupRequest.ProtoReflect.Descriptor instead.
func (*GetBackupRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

type ListBackupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListBackups requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListBackups
	// request to get the next page of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBackupsRequest) Reset() {
	*x = ListBackupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsRequest) ProtoMessage() {}

func (x *ListBackupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsRequest.ProtoReflect.Descriptor instead.
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBackupsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListBackupsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBackupsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBackupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backups []*Backup `protobuf:"bytes,1,rep,name=backups,proto3" json:"backups,omitempty"`
	// This token allows you to get the next page of results for ListBackups requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListBackups request. Subsequent ListBackups
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBackupsResponse) Reset() {
	*x = ListBackupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsResponse) ProtoMessage() {}

func (x *ListBackupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsResponse.ProtoReflect.Descriptor instead.
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListBackupsResponse) GetBackups() []*Backup {
	if x != nil {
		return x.Backups
	}
	return nil
}

func (x *ListBackupsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *DeleteBackupRequest) Reset() {
	*x = DeleteBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupRequest) ProtoMessage() {}

func (x *DeleteBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupRequest.ProtoReflect.Descriptor instead.
func (*DeleteBackupRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

type DeleteBackupMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackupId   string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	DatabaseId string `protobuf:"bytes,2,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
}

func (x *DeleteBackupMetadata) Reset() {
	*x = DeleteBackupMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupMetadata) ProtoMessage() {}

func (x *DeleteBackupMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupMetadata.ProtoReflect.Descriptor instead.
func (*DeleteBackupMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBackupMetadata) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *DeleteBackupMetadata) GetDatabaseId() string {
	if x != nil {
		return x.DatabaseId
	}
	return ""
}

var File_yandex_cloud_ydb_v1_backup_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ydb_v1_backup_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79,
	0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31,
	0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x22, 0x84, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x74, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x49, 0x64, 0x32, 0xa9, 0x04, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x79, 0x64, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x79, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x72,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x73, 0x12, 0xab, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0xb2,
	0xd2, 0x2a, 0x2d, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x3b, 0x79, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ydb_v1_backup_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ydb_v1_backup_service_proto_rawDescData = file_yandex_cloud_ydb_v1_backup_service_proto_rawDesc
)

func file_yandex_cloud_ydb_v1_backup_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ydb_v1_backup_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ydb_v1_backup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ydb_v1_backup_service_proto_rawDescData)
	})
	return file_yandex_cloud_ydb_v1_backup_service_proto_rawDescData
}

var file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_ydb_v1_backup_service_proto_goTypes = []interface{}{
	(*ListPathsRequest)(nil),     // 0: yandex.cloud.ydb.v1.ListPathsRequest
	(*ListPathsResponse)(nil),    // 1: yandex.cloud.ydb.v1.ListPathsResponse
	(*GetBackupRequest)(nil),     // 2: yandex.cloud.ydb.v1.GetBackupRequest
	(*ListBackupsRequest)(nil),   // 3: yandex.cloud.ydb.v1.ListBackupsRequest
	(*ListBackupsResponse)(nil),  // 4: yandex.cloud.ydb.v1.ListBackupsResponse
	(*DeleteBackupRequest)(nil),  // 5: yandex.cloud.ydb.v1.DeleteBackupRequest
	(*DeleteBackupMetadata)(nil), // 6: yandex.cloud.ydb.v1.DeleteBackupMetadata
	(*Backup)(nil),               // 7: yandex.cloud.ydb.v1.Backup
	(*operation.Operation)(nil),  // 8: yandex.cloud.operation.Operation
}
var file_yandex_cloud_ydb_v1_backup_service_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.ydb.v1.ListBackupsResponse.backups:type_name -> yandex.cloud.ydb.v1.Backup
	2, // 1: yandex.cloud.ydb.v1.BackupService.Get:input_type -> yandex.cloud.ydb.v1.GetBackupRequest
	0, // 2: yandex.cloud.ydb.v1.BackupService.ListPaths:input_type -> yandex.cloud.ydb.v1.ListPathsRequest
	3, // 3: yandex.cloud.ydb.v1.BackupService.List:input_type -> yandex.cloud.ydb.v1.ListBackupsRequest
	5, // 4: yandex.cloud.ydb.v1.BackupService.Delete:input_type -> yandex.cloud.ydb.v1.DeleteBackupRequest
	7, // 5: yandex.cloud.ydb.v1.BackupService.Get:output_type -> yandex.cloud.ydb.v1.Backup
	1, // 6: yandex.cloud.ydb.v1.BackupService.ListPaths:output_type -> yandex.cloud.ydb.v1.ListPathsResponse
	4, // 7: yandex.cloud.ydb.v1.BackupService.List:output_type -> yandex.cloud.ydb.v1.ListBackupsResponse
	8, // 8: yandex.cloud.ydb.v1.BackupService.Delete:output_type -> yandex.cloud.operation.Operation
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ydb_v1_backup_service_proto_init() }
func file_yandex_cloud_ydb_v1_backup_service_proto_init() {
	if File_yandex_cloud_ydb_v1_backup_service_proto != nil {
		return
	}
	file_yandex_cloud_ydb_v1_backup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPathsRequest); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPathsResponse); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackupRequest); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsRequest); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsResponse); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupRequest); i {
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
		file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupMetadata); i {
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
			RawDescriptor: file_yandex_cloud_ydb_v1_backup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ydb_v1_backup_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ydb_v1_backup_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ydb_v1_backup_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ydb_v1_backup_service_proto = out.File
	file_yandex_cloud_ydb_v1_backup_service_proto_rawDesc = nil
	file_yandex_cloud_ydb_v1_backup_service_proto_goTypes = nil
	file_yandex_cloud_ydb_v1_backup_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Returns the specified backup.
	Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error)
	ListPaths(ctx context.Context, in *ListPathsRequest, opts ...grpc.CallOption) (*ListPathsResponse, error)
	// Retrieves a list of backups.
	List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
	// Deletes the specified backup.
	Delete(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error) {
	out := new(Backup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.BackupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListPaths(ctx context.Context, in *ListPathsRequest, opts ...grpc.CallOption) (*ListPathsResponse, error) {
	out := new(ListPathsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.BackupService/ListPaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.BackupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) Delete(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.BackupService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	// Returns the specified backup.
	Get(context.Context, *GetBackupRequest) (*Backup, error)
	ListPaths(context.Context, *ListPathsRequest) (*ListPathsResponse, error)
	// Retrieves a list of backups.
	List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
	// Deletes the specified backup.
	Delete(context.Context, *DeleteBackupRequest) (*operation.Operation, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) Get(context.Context, *GetBackupRequest) (*Backup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBackupServiceServer) ListPaths(context.Context, *ListPathsRequest) (*ListPathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaths not implemented")
}
func (*UnimplementedBackupServiceServer) List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedBackupServiceServer) Delete(context.Context, *DeleteBackupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.BackupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Get(ctx, req.(*GetBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListPaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPathsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListPaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.BackupService/ListPaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListPaths(ctx, req.(*ListPathsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.BackupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).List(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.BackupService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Delete(ctx, req.(*DeleteBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ydb.v1.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BackupService_Get_Handler,
		},
		{
			MethodName: "ListPaths",
			Handler:    _BackupService_ListPaths_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BackupService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BackupService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ydb/v1/backup_service.proto",
}
