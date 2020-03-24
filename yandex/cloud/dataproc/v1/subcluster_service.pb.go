// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/v1/subcluster_service.proto

package dataproc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type GetSubclusterRequest struct {
	// ID of the Data Proc cluster that the subcluster belongs to.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster to return.
	//
	// To get a subcluster ID make a [SubclusterService.List] request.
	SubclusterId         string   `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubclusterRequest) Reset()         { *m = GetSubclusterRequest{} }
func (m *GetSubclusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetSubclusterRequest) ProtoMessage()    {}
func (*GetSubclusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{0}
}

func (m *GetSubclusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubclusterRequest.Unmarshal(m, b)
}
func (m *GetSubclusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubclusterRequest.Marshal(b, m, deterministic)
}
func (m *GetSubclusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubclusterRequest.Merge(m, src)
}
func (m *GetSubclusterRequest) XXX_Size() int {
	return xxx_messageInfo_GetSubclusterRequest.Size(m)
}
func (m *GetSubclusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubclusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubclusterRequest proto.InternalMessageInfo

func (m *GetSubclusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *GetSubclusterRequest) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

type ListSubclustersRequest struct {
	// ID of the Data Proc cluster to list subclusters in.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListSubclustersResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListSubclustersResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters subclusters listed in the response.
	//
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Subcluster.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. The value. Must be 3-63 characters long and match the regular expression `^[a-z][-a-z0-9]{1,61}[a-z0-9].
	// Example of a filter: `name=dataproc123_subcluster456`.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSubclustersRequest) Reset()         { *m = ListSubclustersRequest{} }
func (m *ListSubclustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListSubclustersRequest) ProtoMessage()    {}
func (*ListSubclustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{1}
}

func (m *ListSubclustersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubclustersRequest.Unmarshal(m, b)
}
func (m *ListSubclustersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubclustersRequest.Marshal(b, m, deterministic)
}
func (m *ListSubclustersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubclustersRequest.Merge(m, src)
}
func (m *ListSubclustersRequest) XXX_Size() int {
	return xxx_messageInfo_ListSubclustersRequest.Size(m)
}
func (m *ListSubclustersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubclustersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubclustersRequest proto.InternalMessageInfo

func (m *ListSubclustersRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListSubclustersRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSubclustersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListSubclustersRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListSubclustersResponse struct {
	// List of subclusters in the specified cluster.
	Subclusters []*Subcluster `protobuf:"bytes,1,rep,name=subclusters,proto3" json:"subclusters,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListSubclustersRequest.page_size], use `next_page_token` as the value
	// for the [ListSubclustersRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSubclustersResponse) Reset()         { *m = ListSubclustersResponse{} }
func (m *ListSubclustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListSubclustersResponse) ProtoMessage()    {}
func (*ListSubclustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{2}
}

func (m *ListSubclustersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubclustersResponse.Unmarshal(m, b)
}
func (m *ListSubclustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubclustersResponse.Marshal(b, m, deterministic)
}
func (m *ListSubclustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubclustersResponse.Merge(m, src)
}
func (m *ListSubclustersResponse) XXX_Size() int {
	return xxx_messageInfo_ListSubclustersResponse.Size(m)
}
func (m *ListSubclustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubclustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubclustersResponse proto.InternalMessageInfo

func (m *ListSubclustersResponse) GetSubclusters() []*Subcluster {
	if m != nil {
		return m.Subclusters
	}
	return nil
}

func (m *ListSubclustersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateSubclusterRequest struct {
	// ID of the Data Proc cluster to create a subcluster in.
	//
	// To get a cluster ID, make a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the subcluster. The name must be unique within the cluster. The name can’t be
	// changed when the subcluster is created.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Role that is fulfilled by hosts of the subcluster.
	Role Role `protobuf:"varint,3,opt,name=role,proto3,enum=yandex.cloud.dataproc.v1.Role" json:"role,omitempty"`
	// Resources allocated for each host in the subcluster.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// ID of the VPC subnet used for hosts in the subcluster.
	SubnetId string `protobuf:"bytes,5,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Number of hosts in the subcluster.
	HostsCount           int64    `protobuf:"varint,6,opt,name=hosts_count,json=hostsCount,proto3" json:"hosts_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSubclusterRequest) Reset()         { *m = CreateSubclusterRequest{} }
func (m *CreateSubclusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSubclusterRequest) ProtoMessage()    {}
func (*CreateSubclusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{3}
}

func (m *CreateSubclusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSubclusterRequest.Unmarshal(m, b)
}
func (m *CreateSubclusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSubclusterRequest.Marshal(b, m, deterministic)
}
func (m *CreateSubclusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSubclusterRequest.Merge(m, src)
}
func (m *CreateSubclusterRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSubclusterRequest.Size(m)
}
func (m *CreateSubclusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSubclusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSubclusterRequest proto.InternalMessageInfo

func (m *CreateSubclusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateSubclusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSubclusterRequest) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (m *CreateSubclusterRequest) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *CreateSubclusterRequest) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *CreateSubclusterRequest) GetHostsCount() int64 {
	if m != nil {
		return m.HostsCount
	}
	return 0
}

type CreateSubclusterMetadata struct {
	// ID of the cluster that the subcluster is being added to.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster that is being created.
	SubclusterId         string   `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSubclusterMetadata) Reset()         { *m = CreateSubclusterMetadata{} }
func (m *CreateSubclusterMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateSubclusterMetadata) ProtoMessage()    {}
func (*CreateSubclusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{4}
}

func (m *CreateSubclusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSubclusterMetadata.Unmarshal(m, b)
}
func (m *CreateSubclusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSubclusterMetadata.Marshal(b, m, deterministic)
}
func (m *CreateSubclusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSubclusterMetadata.Merge(m, src)
}
func (m *CreateSubclusterMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateSubclusterMetadata.Size(m)
}
func (m *CreateSubclusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSubclusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSubclusterMetadata proto.InternalMessageInfo

func (m *CreateSubclusterMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateSubclusterMetadata) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

type UpdateSubclusterRequest struct {
	// ID of the cluster to update a subcluster in.
	//
	// To get a cluster ID, make a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster to update.
	//
	// To get a subcluster ID, make a [SubclusterService.List] request.
	SubclusterId string `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	// Field mask that specifies which attributes of the subcluster should be updated.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// New configuration of resources that should be allocated for each host in the subcluster.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// New name for the subcluster. The name must be unique within the cluster.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// New number of hosts in the subcluster.
	HostsCount           int64    `protobuf:"varint,6,opt,name=hosts_count,json=hostsCount,proto3" json:"hosts_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSubclusterRequest) Reset()         { *m = UpdateSubclusterRequest{} }
func (m *UpdateSubclusterRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSubclusterRequest) ProtoMessage()    {}
func (*UpdateSubclusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{5}
}

func (m *UpdateSubclusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSubclusterRequest.Unmarshal(m, b)
}
func (m *UpdateSubclusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSubclusterRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSubclusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSubclusterRequest.Merge(m, src)
}
func (m *UpdateSubclusterRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSubclusterRequest.Size(m)
}
func (m *UpdateSubclusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSubclusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSubclusterRequest proto.InternalMessageInfo

func (m *UpdateSubclusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *UpdateSubclusterRequest) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

func (m *UpdateSubclusterRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateSubclusterRequest) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *UpdateSubclusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSubclusterRequest) GetHostsCount() int64 {
	if m != nil {
		return m.HostsCount
	}
	return 0
}

type UpdateSubclusterMetadata struct {
	// ID of the cluster whose subcluster is being updated.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster that is being updated.
	SubclusterId         string   `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSubclusterMetadata) Reset()         { *m = UpdateSubclusterMetadata{} }
func (m *UpdateSubclusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateSubclusterMetadata) ProtoMessage()    {}
func (*UpdateSubclusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{6}
}

func (m *UpdateSubclusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSubclusterMetadata.Unmarshal(m, b)
}
func (m *UpdateSubclusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSubclusterMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateSubclusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSubclusterMetadata.Merge(m, src)
}
func (m *UpdateSubclusterMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateSubclusterMetadata.Size(m)
}
func (m *UpdateSubclusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSubclusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSubclusterMetadata proto.InternalMessageInfo

func (m *UpdateSubclusterMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *UpdateSubclusterMetadata) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

type DeleteSubclusterRequest struct {
	// ID of the cluster to remove a subcluster from.
	//
	// To get a cluster ID, make a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster to delete.
	SubclusterId         string   `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSubclusterRequest) Reset()         { *m = DeleteSubclusterRequest{} }
func (m *DeleteSubclusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSubclusterRequest) ProtoMessage()    {}
func (*DeleteSubclusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{7}
}

func (m *DeleteSubclusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubclusterRequest.Unmarshal(m, b)
}
func (m *DeleteSubclusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubclusterRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSubclusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubclusterRequest.Merge(m, src)
}
func (m *DeleteSubclusterRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSubclusterRequest.Size(m)
}
func (m *DeleteSubclusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubclusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubclusterRequest proto.InternalMessageInfo

func (m *DeleteSubclusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteSubclusterRequest) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

type DeleteSubclusterMetadata struct {
	// ID of the cluster whose subcluster is being deleted.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the subcluster that is being deleted.
	SubclusterId         string   `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSubclusterMetadata) Reset()         { *m = DeleteSubclusterMetadata{} }
func (m *DeleteSubclusterMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteSubclusterMetadata) ProtoMessage()    {}
func (*DeleteSubclusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f77207766937a6, []int{8}
}

func (m *DeleteSubclusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubclusterMetadata.Unmarshal(m, b)
}
func (m *DeleteSubclusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubclusterMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteSubclusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubclusterMetadata.Merge(m, src)
}
func (m *DeleteSubclusterMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteSubclusterMetadata.Size(m)
}
func (m *DeleteSubclusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubclusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubclusterMetadata proto.InternalMessageInfo

func (m *DeleteSubclusterMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteSubclusterMetadata) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSubclusterRequest)(nil), "yandex.cloud.dataproc.v1.GetSubclusterRequest")
	proto.RegisterType((*ListSubclustersRequest)(nil), "yandex.cloud.dataproc.v1.ListSubclustersRequest")
	proto.RegisterType((*ListSubclustersResponse)(nil), "yandex.cloud.dataproc.v1.ListSubclustersResponse")
	proto.RegisterType((*CreateSubclusterRequest)(nil), "yandex.cloud.dataproc.v1.CreateSubclusterRequest")
	proto.RegisterType((*CreateSubclusterMetadata)(nil), "yandex.cloud.dataproc.v1.CreateSubclusterMetadata")
	proto.RegisterType((*UpdateSubclusterRequest)(nil), "yandex.cloud.dataproc.v1.UpdateSubclusterRequest")
	proto.RegisterType((*UpdateSubclusterMetadata)(nil), "yandex.cloud.dataproc.v1.UpdateSubclusterMetadata")
	proto.RegisterType((*DeleteSubclusterRequest)(nil), "yandex.cloud.dataproc.v1.DeleteSubclusterRequest")
	proto.RegisterType((*DeleteSubclusterMetadata)(nil), "yandex.cloud.dataproc.v1.DeleteSubclusterMetadata")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/v1/subcluster_service.proto", fileDescriptor_27f77207766937a6)
}

var fileDescriptor_27f77207766937a6 = []byte{
	// 905 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xc6, 0x8e, 0x89, 0x9f, 0x5b, 0x10, 0x23, 0x20, 0x2b, 0xab, 0x45, 0xe9, 0x02, 0x69,
	0x1a, 0xea, 0x5d, 0x8f, 0xab, 0x56, 0x40, 0x9b, 0x0a, 0x12, 0x9a, 0x28, 0x12, 0x15, 0x68, 0x03,
	0x17, 0xa2, 0xca, 0x1a, 0xef, 0x4e, 0xdc, 0x55, 0xd6, 0x3b, 0xcb, 0xce, 0xac, 0xd5, 0x24, 0xe4,
	0xc2, 0x8d, 0x5c, 0xf9, 0x09, 0xfc, 0x83, 0x9e, 0xb9, 0x71, 0x70, 0xce, 0xed, 0xa9, 0x42, 0xe2,
	0xc4, 0x81, 0x33, 0xc7, 0x9e, 0xd0, 0xce, 0xee, 0x66, 0xd7, 0xc6, 0x6b, 0xd9, 0xb5, 0x2a, 0x6e,
	0x63, 0xbf, 0xef, 0xcd, 0xfb, 0xbe, 0xf7, 0xe6, 0xbd, 0xb7, 0x80, 0x8f, 0x88, 0x67, 0xd3, 0x27,
	0x86, 0xe5, 0xb2, 0xd0, 0x36, 0x6c, 0x22, 0x88, 0x1f, 0x30, 0xcb, 0xe8, 0x63, 0x83, 0x87, 0x1d,
	0xcb, 0x0d, 0xb9, 0xa0, 0x41, 0x9b, 0xd3, 0xa0, 0xef, 0x58, 0x54, 0xf7, 0x03, 0x26, 0x18, 0x52,
	0x63, 0x17, 0x5d, 0xba, 0xe8, 0xa9, 0x8b, 0xde, 0xc7, 0xf5, 0x2b, 0x5d, 0xc6, 0xba, 0x2e, 0x35,
	0x88, 0xef, 0x18, 0xc4, 0xf3, 0x98, 0x20, 0xc2, 0x61, 0x1e, 0x8f, 0xfd, 0xea, 0x2b, 0x89, 0x55,
	0xfe, 0xea, 0x84, 0x07, 0xc6, 0x81, 0x43, 0x5d, 0xbb, 0xdd, 0x23, 0xfc, 0x30, 0x41, 0xac, 0x0e,
	0x91, 0x61, 0x3e, 0x0d, 0xe4, 0x05, 0xd9, 0x29, 0xc1, 0x7d, 0x54, 0x48, 0xda, 0x62, 0xbd, 0xde,
	0x05, 0xec, 0xc6, 0x14, 0xda, 0x12, 0xe8, 0xd5, 0x21, 0x68, 0x9f, 0xb8, 0x8e, 0x9d, 0x0f, 0xb8,
	0x32, 0x64, 0x8e, 0xe4, 0x8d, 0x50, 0xd2, 0xfa, 0xf0, 0xce, 0x0e, 0x15, 0x7b, 0x17, 0xf7, 0x9a,
	0xf4, 0x87, 0x90, 0x72, 0x81, 0x3e, 0x06, 0x48, 0xb3, 0xe8, 0xd8, 0xaa, 0xb2, 0xa2, 0xac, 0x55,
	0x37, 0x2f, 0xfd, 0x3d, 0xc0, 0xca, 0xd9, 0x39, 0x2e, 0xdf, 0xdb, 0xb8, 0xdd, 0x34, 0xab, 0x89,
	0x7d, 0xd7, 0x46, 0x18, 0x2e, 0xe7, 0xb2, 0xee, 0xd8, 0xea, 0xc2, 0x18, 0xfc, 0xa5, 0x0c, 0xb2,
	0x6b, 0x6b, 0xbf, 0x29, 0xf0, 0xde, 0x57, 0x0e, 0xcf, 0x45, 0xe6, 0xaf, 0x14, 0xfa, 0x3a, 0x54,
	0x7d, 0xd2, 0xa5, 0x6d, 0xee, 0x1c, 0x53, 0x19, 0xb6, 0xb4, 0x09, 0x2f, 0x07, 0xb8, 0x72, 0x6f,
	0x03, 0x37, 0x9b, 0x4d, 0x73, 0x29, 0x32, 0xee, 0x39, 0xc7, 0x14, 0xad, 0x01, 0x48, 0xa0, 0x60,
	0x87, 0xd4, 0x53, 0x4b, 0xf2, 0xd6, 0xea, 0xd9, 0x39, 0x5e, 0x94, 0x48, 0x53, 0xde, 0xf2, 0x6d,
	0x64, 0x43, 0x1a, 0x54, 0x0e, 0x1c, 0x57, 0xd0, 0x40, 0x2d, 0x4b, 0x14, 0x9c, 0x9d, 0x5f, 0xdc,
	0x97, 0x58, 0xb4, 0x9f, 0x15, 0x58, 0xfe, 0x0f, 0x7d, 0xee, 0x33, 0x8f, 0x53, 0xb4, 0x0d, 0xb5,
	0x4c, 0x2a, 0x57, 0x95, 0x95, 0xd2, 0x5a, 0xad, 0xf5, 0xa1, 0x5e, 0xf4, 0xfa, 0xf4, 0x5c, 0xf2,
	0xf3, 0x8e, 0x68, 0x15, 0xde, 0xf2, 0xe8, 0x13, 0xd1, 0xce, 0xd1, 0x96, 0x79, 0x35, 0x2f, 0x47,
	0x7f, 0x7f, 0x93, 0xf2, 0xd5, 0x5e, 0x2c, 0xc0, 0xf2, 0x56, 0x40, 0x89, 0xa0, 0x73, 0x96, 0xf1,
	0x36, 0x94, 0x3d, 0xd2, 0xa3, 0x49, 0xf5, 0xae, 0xfd, 0x33, 0xc0, 0x57, 0x7f, 0xdc, 0x27, 0x8d,
	0xe3, 0x47, 0xfb, 0x0d, 0xd2, 0x38, 0x6e, 0x36, 0x3e, 0x7d, 0x74, 0x82, 0x6f, 0xde, 0xc1, 0xa7,
	0xfb, 0xc9, 0x2f, 0x53, 0xc2, 0xd1, 0x27, 0x50, 0x0e, 0x98, 0x4b, 0x65, 0x4e, 0xdf, 0x6c, 0xbd,
	0x5f, 0x2c, 0xd4, 0x64, 0x2e, 0xdd, 0x2c, 0x47, 0xd1, 0x4d, 0xe9, 0x81, 0x76, 0xa0, 0x1a, 0x50,
	0xce, 0xc2, 0xc0, 0xa2, 0x5c, 0x26, 0xbb, 0xd6, 0xfa, 0x60, 0x82, 0x7b, 0x0a, 0x4d, 0xee, 0xc8,
	0x7c, 0xd1, 0x0d, 0xa8, 0xf2, 0xb0, 0xe3, 0x51, 0x11, 0xa9, 0x5c, 0x1c, 0xa3, 0x72, 0x29, 0x36,
	0xef, 0xda, 0xe8, 0x26, 0xd4, 0x1e, 0x33, 0x2e, 0x78, 0xdb, 0x62, 0xa1, 0x27, 0xd4, 0x8a, 0x7c,
	0x32, 0xb5, 0x08, 0xfc, 0x72, 0x80, 0x4b, 0xf7, 0x37, 0xb0, 0x09, 0xd2, 0xbe, 0x15, 0x99, 0xb5,
	0x00, 0xd4, 0xd1, 0xd4, 0x3e, 0xa4, 0x82, 0x44, 0xb4, 0xd0, 0xf5, 0x31, 0xb9, 0x5d, 0x1a, 0x97,
	0xd7, 0xc6, 0xf8, 0xf6, 0x58, 0x2a, 0x68, 0x8d, 0x3f, 0x17, 0x60, 0xf9, 0x3b, 0xdf, 0x9e, 0xbf,
	0x9e, 0xb3, 0xb7, 0x25, 0xba, 0x0b, 0xb5, 0x50, 0x86, 0x96, 0xe3, 0x4d, 0x96, 0xb4, 0xd6, 0xaa,
	0xeb, 0xf1, 0x04, 0xd4, 0xd3, 0x09, 0xa8, 0x6f, 0x47, 0x13, 0xf0, 0x21, 0xe1, 0x87, 0x26, 0xc4,
	0xf0, 0xe8, 0x8c, 0xbe, 0x78, 0xb5, 0x72, 0xe6, 0x0b, 0x99, 0x3e, 0xc1, 0xc5, 0xd9, 0x9e, 0xe0,
	0xcc, 0x45, 0x1d, 0xcd, 0xef, 0x6b, 0x2f, 0xea, 0x11, 0x2c, 0x7f, 0x49, 0x5d, 0xfa, 0x3f, 0xd4,
	0x34, 0x92, 0x3b, 0x1a, 0xfa, 0x75, 0xcb, 0x6d, 0xbd, 0x78, 0x03, 0xde, 0xce, 0xc2, 0xed, 0xc5,
	0x7b, 0x18, 0xfd, 0xaa, 0x40, 0x69, 0x87, 0x0a, 0xa4, 0x17, 0xbf, 0x8a, 0x71, 0xcb, 0xa8, 0x3e,
	0xd5, 0xf0, 0xd4, 0xb6, 0x7f, 0x7a, 0xfe, 0xd7, 0x2f, 0x0b, 0x9f, 0xa3, 0xfb, 0xc3, 0x9b, 0x35,
	0x19, 0xa7, 0xc6, 0x49, 0x46, 0xfe, 0x34, 0xb7, 0x48, 0xb9, 0x71, 0x32, 0x24, 0xec, 0x34, 0x62,
	0x59, 0x8e, 0x66, 0x3b, 0x6a, 0x16, 0x87, 0x1d, 0xbf, 0xba, 0xea, 0x78, 0x06, 0x8f, 0x78, 0x5b,
	0x68, 0x77, 0x24, 0xeb, 0x26, 0xd2, 0x67, 0x63, 0x8d, 0x7e, 0x57, 0xa0, 0x12, 0x8f, 0x26, 0x34,
	0x21, 0x6a, 0xc1, 0x5e, 0xa8, 0x5f, 0x1b, 0x76, 0xc9, 0xbe, 0x0a, 0xbe, 0x4e, 0x4f, 0x9a, 0xf5,
	0xf4, 0xd9, 0xfa, 0xea, 0x84, 0xf1, 0x07, 0xd9, 0x7f, 0x52, 0xc2, 0x2d, 0x6d, 0x46, 0x09, 0x9f,
	0x29, 0xeb, 0xe8, 0xb9, 0x02, 0x95, 0xb8, 0x17, 0x27, 0xa9, 0x28, 0x98, 0x86, 0xd3, 0xa8, 0xe0,
	0xb1, 0x8a, 0xc2, 0x7e, 0x1f, 0x55, 0xb1, 0xd5, 0x9a, 0xf3, 0xf9, 0x44, 0xaa, 0xfe, 0x50, 0xa0,
	0x12, 0xb7, 0xdc, 0x24, 0x55, 0x05, 0xf3, 0x60, 0x1a, 0x55, 0x27, 0x4f, 0x9f, 0xad, 0xe3, 0x09,
	0x6d, 0xfd, 0xee, 0xe8, 0xc0, 0x7e, 0xd0, 0xf3, 0xc5, 0x51, 0xdc, 0x1f, 0xeb, 0x73, 0x0a, 0xdc,
	0xa4, 0x70, 0x65, 0x88, 0x20, 0xf1, 0x9d, 0xbc, 0xae, 0xef, 0x1f, 0x74, 0x1d, 0xf1, 0x38, 0xec,
	0xe8, 0x16, 0xeb, 0x19, 0x31, 0xb0, 0x11, 0x7f, 0x7f, 0x76, 0x59, 0xa3, 0x4b, 0x3d, 0xc9, 0xc9,
	0x28, 0xfa, 0xc4, 0xbd, 0x9b, 0x9e, 0x3b, 0x15, 0x09, 0xbc, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x71, 0xf0, 0x10, 0xea, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubclusterServiceClient is the client API for SubclusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubclusterServiceClient interface {
	// Returns the specified subcluster.
	//
	// To get the list of all available subclusters, make a [SubclusterService.List] request.
	Get(ctx context.Context, in *GetSubclusterRequest, opts ...grpc.CallOption) (*Subcluster, error)
	// Retrieves a list of subclusters in the specified cluster.
	List(ctx context.Context, in *ListSubclustersRequest, opts ...grpc.CallOption) (*ListSubclustersResponse, error)
	// Creates a subcluster in the specified cluster.
	Create(ctx context.Context, in *CreateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified subcluster.
	Update(ctx context.Context, in *UpdateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified subcluster.
	Delete(ctx context.Context, in *DeleteSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type subclusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubclusterServiceClient(cc grpc.ClientConnInterface) SubclusterServiceClient {
	return &subclusterServiceClient{cc}
}

func (c *subclusterServiceClient) Get(ctx context.Context, in *GetSubclusterRequest, opts ...grpc.CallOption) (*Subcluster, error) {
	out := new(Subcluster)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.SubclusterService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subclusterServiceClient) List(ctx context.Context, in *ListSubclustersRequest, opts ...grpc.CallOption) (*ListSubclustersResponse, error) {
	out := new(ListSubclustersResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.SubclusterService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subclusterServiceClient) Create(ctx context.Context, in *CreateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.SubclusterService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subclusterServiceClient) Update(ctx context.Context, in *UpdateSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.SubclusterService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subclusterServiceClient) Delete(ctx context.Context, in *DeleteSubclusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.SubclusterService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubclusterServiceServer is the server API for SubclusterService service.
type SubclusterServiceServer interface {
	// Returns the specified subcluster.
	//
	// To get the list of all available subclusters, make a [SubclusterService.List] request.
	Get(context.Context, *GetSubclusterRequest) (*Subcluster, error)
	// Retrieves a list of subclusters in the specified cluster.
	List(context.Context, *ListSubclustersRequest) (*ListSubclustersResponse, error)
	// Creates a subcluster in the specified cluster.
	Create(context.Context, *CreateSubclusterRequest) (*operation.Operation, error)
	// Updates the specified subcluster.
	Update(context.Context, *UpdateSubclusterRequest) (*operation.Operation, error)
	// Deletes the specified subcluster.
	Delete(context.Context, *DeleteSubclusterRequest) (*operation.Operation, error)
}

// UnimplementedSubclusterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubclusterServiceServer struct {
}

func (*UnimplementedSubclusterServiceServer) Get(ctx context.Context, req *GetSubclusterRequest) (*Subcluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSubclusterServiceServer) List(ctx context.Context, req *ListSubclustersRequest) (*ListSubclustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedSubclusterServiceServer) Create(ctx context.Context, req *CreateSubclusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSubclusterServiceServer) Update(ctx context.Context, req *UpdateSubclusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedSubclusterServiceServer) Delete(ctx context.Context, req *DeleteSubclusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterSubclusterServiceServer(s *grpc.Server, srv SubclusterServiceServer) {
	s.RegisterService(&_SubclusterService_serviceDesc, srv)
}

func _SubclusterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubclusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubclusterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.SubclusterService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubclusterServiceServer).Get(ctx, req.(*GetSubclusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubclusterService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubclustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubclusterServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.SubclusterService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubclusterServiceServer).List(ctx, req.(*ListSubclustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubclusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubclusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubclusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.SubclusterService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubclusterServiceServer).Create(ctx, req.(*CreateSubclusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubclusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubclusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubclusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.SubclusterService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubclusterServiceServer).Update(ctx, req.(*UpdateSubclusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubclusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubclusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubclusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.SubclusterService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubclusterServiceServer).Delete(ctx, req.(*DeleteSubclusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubclusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.v1.SubclusterService",
	HandlerType: (*SubclusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SubclusterService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SubclusterService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SubclusterService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SubclusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SubclusterService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/v1/subcluster_service.proto",
}
