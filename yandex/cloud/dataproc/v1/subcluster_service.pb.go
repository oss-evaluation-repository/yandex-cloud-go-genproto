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
	HostsCount int64 `protobuf:"varint,6,opt,name=hosts_count,json=hostsCount,proto3" json:"hosts_count,omitempty"`
	// Timeout to gracefully decommission nodes. In seconds. Default value: 0
	DecommissionTimeout  int64    `protobuf:"varint,7,opt,name=decommission_timeout,json=decommissionTimeout,proto3" json:"decommission_timeout,omitempty"`
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

func (m *UpdateSubclusterRequest) GetDecommissionTimeout() int64 {
	if m != nil {
		return m.DecommissionTimeout
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
	SubclusterId string `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	// Timeout to gracefully decommission nodes. In seconds. Default value: 0
	DecommissionTimeout  int64    `protobuf:"varint,3,opt,name=decommission_timeout,json=decommissionTimeout,proto3" json:"decommission_timeout,omitempty"`
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

func (m *DeleteSubclusterRequest) GetDecommissionTimeout() int64 {
	if m != nil {
		return m.DecommissionTimeout
	}
	return 0
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
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xbf, 0x6f, 0xdb, 0x46,
	0x14, 0x06, 0x2d, 0x59, 0xb6, 0x9e, 0x92, 0x16, 0xbd, 0xa6, 0x35, 0x21, 0x24, 0x85, 0xc3, 0xb6,
	0x8e, 0xe3, 0x46, 0x24, 0x4f, 0x69, 0x8c, 0xb4, 0x89, 0x8d, 0xd6, 0x6e, 0x6c, 0x18, 0x68, 0xd0,
	0x82, 0x4e, 0x97, 0x1a, 0x81, 0x40, 0x91, 0x67, 0x85, 0x30, 0xc5, 0x63, 0x79, 0x47, 0x21, 0xb1,
	0xeb, 0xa5, 0x5b, 0xbd, 0x76, 0xef, 0xd2, 0xff, 0x20, 0x5b, 0x81, 0x6e, 0x1d, 0xe4, 0x39, 0x99,
	0x82, 0xae, 0x1d, 0x3a, 0x77, 0xf4, 0x54, 0xf0, 0x48, 0x99, 0x94, 0x2a, 0x0a, 0x52, 0x04, 0x23,
	0xdb, 0x91, 0xef, 0x7b, 0xf7, 0xde, 0xf7, 0x7e, 0x1e, 0xe0, 0x67, 0xa6, 0x67, 0x93, 0xa7, 0x9a,
	0xe5, 0xd2, 0xd0, 0xd6, 0x6c, 0x93, 0x9b, 0x7e, 0x40, 0x2d, 0xad, 0x83, 0x35, 0x16, 0x36, 0x2d,
	0x37, 0x64, 0x9c, 0x04, 0x0d, 0x46, 0x82, 0x8e, 0x63, 0x11, 0xd5, 0x0f, 0x28, 0xa7, 0x48, 0x8e,
	0x55, 0x54, 0xa1, 0xa2, 0xf6, 0x54, 0xd4, 0x0e, 0xae, 0x5e, 0x6d, 0x51, 0xda, 0x72, 0x89, 0x66,
	0xfa, 0x8e, 0x66, 0x7a, 0x1e, 0xe5, 0x26, 0x77, 0xa8, 0xc7, 0x62, 0xbd, 0xea, 0x62, 0x22, 0x15,
	0x5f, 0xcd, 0x70, 0x5f, 0xdb, 0x77, 0x88, 0x6b, 0x37, 0xda, 0x26, 0x3b, 0x48, 0x10, 0x4b, 0x7d,
	0xce, 0x50, 0x9f, 0x04, 0xe2, 0x82, 0xf4, 0x94, 0xe0, 0x3e, 0xce, 0x75, 0xda, 0xa2, 0xed, 0xf6,
	0x39, 0xec, 0xe6, 0x18, 0xdc, 0x12, 0xe8, 0xb5, 0x3e, 0x68, 0xc7, 0x74, 0x1d, 0x3b, 0x6b, 0x70,
	0xb1, 0x4f, 0x1c, 0xd1, 0x1b, 0x70, 0x49, 0xe9, 0xc0, 0x95, 0x6d, 0xc2, 0x77, 0xcf, 0xef, 0x35,
	0xc8, 0x0f, 0x21, 0x61, 0x1c, 0x7d, 0x02, 0xd0, 0x8b, 0xa2, 0x63, 0xcb, 0xd2, 0xa2, 0xb4, 0x5c,
	0xde, 0xb8, 0xf4, 0x4f, 0x17, 0x4b, 0x27, 0xa7, 0xb8, 0x78, 0x7f, 0xed, 0x8e, 0x6e, 0x94, 0x13,
	0xf9, 0x8e, 0x8d, 0x30, 0x5c, 0xce, 0x44, 0xdd, 0xb1, 0xe5, 0x99, 0x21, 0xf8, 0x4b, 0x29, 0x64,
	0xc7, 0x56, 0xfe, 0x90, 0xe0, 0xfd, 0xaf, 0x1d, 0x96, 0xb1, 0xcc, 0x5e, 0xcb, 0xf4, 0x0d, 0x28,
	0xfb, 0x66, 0x8b, 0x34, 0x98, 0x73, 0x48, 0x84, 0xd9, 0xc2, 0x06, 0x9c, 0x75, 0x71, 0xe9, 0xfe,
	0x1a, 0xd6, 0x75, 0xdd, 0x98, 0x8f, 0x84, 0xbb, 0xce, 0x21, 0x41, 0xcb, 0x00, 0x02, 0xc8, 0xe9,
	0x01, 0xf1, 0xe4, 0x82, 0xb8, 0xb5, 0x7c, 0x72, 0x8a, 0x67, 0x05, 0xd2, 0x10, 0xb7, 0x3c, 0x8a,
	0x64, 0x48, 0x81, 0xd2, 0xbe, 0xe3, 0x72, 0x12, 0xc8, 0x45, 0x81, 0x82, 0x93, 0xd3, 0xf3, 0xfb,
	0x12, 0x89, 0xf2, 0xb3, 0x04, 0x0b, 0xff, 0x73, 0x9f, 0xf9, 0xd4, 0x63, 0x04, 0x6d, 0x41, 0x25,
	0xa5, 0xca, 0x64, 0x69, 0xb1, 0xb0, 0x5c, 0xa9, 0x7f, 0xa4, 0xe6, 0x55, 0x9f, 0x9a, 0x09, 0x7e,
	0x56, 0x11, 0x2d, 0xc1, 0xdb, 0x1e, 0x79, 0xca, 0x1b, 0x19, 0xb7, 0x45, 0x5c, 0x8d, 0xcb, 0xd1,
	0xef, 0x6f, 0x7b, 0xfe, 0x2a, 0xaf, 0x66, 0x60, 0x61, 0x33, 0x20, 0x26, 0x27, 0x53, 0xa6, 0xf1,
	0x0e, 0x14, 0x3d, 0xb3, 0x4d, 0x92, 0xec, 0x5d, 0xff, 0xb7, 0x8b, 0xaf, 0xfd, 0xb8, 0x67, 0xd6,
	0x0e, 0x1f, 0xef, 0xd5, 0xcc, 0xda, 0xa1, 0x5e, 0xfb, 0xec, 0xf1, 0x11, 0xbe, 0xb5, 0x8a, 0x8f,
	0xf7, 0x92, 0x2f, 0x43, 0xc0, 0xd1, 0x5d, 0x28, 0x06, 0xd4, 0x25, 0x22, 0xa6, 0x6f, 0xd5, 0x3f,
	0xc8, 0x27, 0x6a, 0x50, 0x97, 0x6c, 0x14, 0x23, 0xeb, 0x86, 0xd0, 0x40, 0xdb, 0x50, 0x0e, 0x08,
	0xa3, 0x61, 0x60, 0x11, 0x26, 0x82, 0x5d, 0xa9, 0x7f, 0x38, 0x42, 0xbd, 0x07, 0x4d, 0xee, 0x48,
	0x75, 0xd1, 0x4d, 0x28, 0xb3, 0xb0, 0xe9, 0x11, 0x1e, 0xb1, 0x9c, 0x1d, 0xc2, 0x72, 0x3e, 0x16,
	0xef, 0xd8, 0xe8, 0x16, 0x54, 0x9e, 0x50, 0xc6, 0x59, 0xc3, 0xa2, 0xa1, 0xc7, 0xe5, 0x92, 0x28,
	0x99, 0x4a, 0x04, 0x3e, 0xeb, 0xe2, 0xc2, 0xfa, 0x1a, 0x36, 0x40, 0xc8, 0x37, 0x23, 0xb1, 0x12,
	0x80, 0x3c, 0x18, 0xda, 0x87, 0x84, 0x9b, 0x91, 0x5b, 0xe8, 0xc6, 0x90, 0xd8, 0xce, 0x0f, 0x8b,
	0x6b, 0x6d, 0x78, 0x7b, 0xcc, 0xe7, 0xb4, 0xc6, 0xaf, 0x05, 0x58, 0xf8, 0xce, 0xb7, 0xa7, 0xcf,
	0xe7, 0xe4, 0x6d, 0x89, 0xee, 0x41, 0x25, 0x14, 0xa6, 0xc5, 0x78, 0x13, 0x29, 0xad, 0xd4, 0xab,
	0x6a, 0x3c, 0x01, 0xd5, 0xde, 0x04, 0x54, 0xb7, 0xa2, 0x09, 0xf8, 0xd0, 0x64, 0x07, 0x06, 0xc4,
	0xf0, 0xe8, 0x8c, 0xbe, 0x7c, 0xbd, 0x74, 0x66, 0x13, 0xd9, 0x2b, 0xc1, 0xd9, 0xc9, 0x4a, 0x70,
	0xa2, 0xa4, 0xa2, 0x75, 0xb8, 0x62, 0x93, 0x68, 0xe2, 0x3a, 0x8c, 0x39, 0xd4, 0x6b, 0x70, 0xa7,
	0x4d, 0x68, 0xc8, 0xe5, 0xb9, 0x58, 0xed, 0xac, 0x8b, 0xe7, 0xf4, 0xda, 0xdd, 0xd5, 0x4f, 0x75,
	0xdd, 0x78, 0x37, 0x0b, 0x7c, 0x14, 0xe3, 0xa2, 0xa2, 0x18, 0xcc, 0xcf, 0x85, 0x17, 0xc5, 0xef,
	0x12, 0x2c, 0x7c, 0x45, 0x5c, 0xf2, 0x26, 0x8a, 0x22, 0x2f, 0x5e, 0x85, 0xf1, 0xe3, 0x35, 0xe8,
	0xfa, 0x45, 0xc7, 0xab, 0xfe, 0x6a, 0x0e, 0xde, 0x49, 0xcd, 0xed, 0xc6, 0x0f, 0x01, 0xf4, 0x9b,
	0x04, 0x85, 0x6d, 0xc2, 0x91, 0x9a, 0x5f, 0x96, 0xc3, 0xb6, 0x61, 0x75, 0xac, 0xe9, 0xad, 0x6c,
	0xfd, 0xf4, 0xf2, 0xef, 0x5f, 0x66, 0xbe, 0x40, 0xeb, 0xfd, 0xab, 0x3d, 0x99, 0xe7, 0xda, 0x51,
	0xea, 0xfc, 0x71, 0x66, 0x93, 0x33, 0xed, 0xa8, 0x8f, 0xd8, 0x71, 0xe4, 0x65, 0x31, 0x5a, 0x2e,
	0x48, 0xcf, 0x37, 0x3b, 0x7c, 0x77, 0x56, 0xf1, 0x04, 0x1a, 0xf1, 0xba, 0x52, 0x56, 0x85, 0xd7,
	0x3a, 0x52, 0x27, 0xf3, 0x1a, 0xfd, 0x29, 0x41, 0x29, 0x9e, 0x8d, 0x68, 0x84, 0xd5, 0x9c, 0xc5,
	0x54, 0xbd, 0xde, 0xaf, 0x92, 0x3e, 0x4b, 0xbe, 0xe9, 0x9d, 0x14, 0xeb, 0xf9, 0x8b, 0x95, 0xa5,
	0x11, 0xf3, 0x17, 0xd2, 0x7f, 0x82, 0xc2, 0x6d, 0x65, 0x42, 0x0a, 0x9f, 0x4b, 0x2b, 0xe8, 0xa5,
	0x04, 0xa5, 0xb8, 0x99, 0x47, 0xb1, 0xc8, 0x19, 0xc7, 0xe3, 0xb0, 0x60, 0x31, 0x8b, 0xdc, 0x81,
	0x31, 0xc8, 0x62, 0xb3, 0x3e, 0x65, 0xf9, 0x44, 0xac, 0xfe, 0x92, 0xa0, 0x14, 0xb7, 0xdc, 0x28,
	0x56, 0x39, 0xf3, 0x64, 0x1c, 0x56, 0x47, 0xcf, 0x5f, 0xac, 0xe0, 0x11, 0x6d, 0xfd, 0xde, 0xe0,
	0xc6, 0x78, 0xd0, 0xf6, 0xf9, 0xb3, 0xb8, 0x3f, 0x56, 0xa6, 0x24, 0xb8, 0x41, 0xe0, 0x6a, 0x9f,
	0x83, 0xa6, 0xef, 0x64, 0x79, 0x7d, 0xff, 0xa0, 0xe5, 0xf0, 0x27, 0x61, 0x53, 0xb5, 0x68, 0x5b,
	0x8b, 0x81, 0xb5, 0xf8, 0x01, 0xdc, 0xa2, 0xb5, 0x16, 0xf1, 0x84, 0x4f, 0x5a, 0xde, 0x1b, 0xfb,
	0x5e, 0xef, 0xdc, 0x2c, 0x09, 0xe0, 0xed, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe1, 0x19,
	0x0f, 0x6b, 0x0c, 0x00, 0x00,
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
