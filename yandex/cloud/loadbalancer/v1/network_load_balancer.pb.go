// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/loadbalancer/v1/network_load_balancer.proto

package loadbalancer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
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

// IP version of the addresses that the load balancer works with.
// Only IPv4 is currently available.
type IpVersion int32

const (
	IpVersion_IP_VERSION_UNSPECIFIED IpVersion = 0
	// IPv4
	IpVersion_IPV4 IpVersion = 1
	// IPv6
	IpVersion_IPV6 IpVersion = 2
)

var IpVersion_name = map[int32]string{
	0: "IP_VERSION_UNSPECIFIED",
	1: "IPV4",
	2: "IPV6",
}

var IpVersion_value = map[string]int32{
	"IP_VERSION_UNSPECIFIED": 0,
	"IPV4":                   1,
	"IPV6":                   2,
}

func (x IpVersion) String() string {
	return proto.EnumName(IpVersion_name, int32(x))
}

func (IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{0}
}

type NetworkLoadBalancer_Status int32

const (
	NetworkLoadBalancer_STATUS_UNSPECIFIED NetworkLoadBalancer_Status = 0
	// Network load balancer is being created.
	NetworkLoadBalancer_CREATING NetworkLoadBalancer_Status = 1
	// Network load balancer is being started.
	NetworkLoadBalancer_STARTING NetworkLoadBalancer_Status = 2
	// Network load balancer is active and sends traffic to the targets.
	NetworkLoadBalancer_ACTIVE NetworkLoadBalancer_Status = 3
	// Network load balancer is being stopped.
	NetworkLoadBalancer_STOPPING NetworkLoadBalancer_Status = 4
	// Network load balancer is stopped and doesn't send traffic to the targets.
	NetworkLoadBalancer_STOPPED NetworkLoadBalancer_Status = 5
	// Network load balancer is being deleted.
	NetworkLoadBalancer_DELETING NetworkLoadBalancer_Status = 6
	// The load balancer doesn't have any listeners or target groups, or
	// attached target groups are empty. The load balancer doesn't perform any health checks or
	// send traffic in this state.
	NetworkLoadBalancer_INACTIVE NetworkLoadBalancer_Status = 7
)

var NetworkLoadBalancer_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "STARTING",
	3: "ACTIVE",
	4: "STOPPING",
	5: "STOPPED",
	6: "DELETING",
	7: "INACTIVE",
}

var NetworkLoadBalancer_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"STARTING":           2,
	"ACTIVE":             3,
	"STOPPING":           4,
	"STOPPED":            5,
	"DELETING":           6,
	"INACTIVE":           7,
}

func (x NetworkLoadBalancer_Status) String() string {
	return proto.EnumName(NetworkLoadBalancer_Status_name, int32(x))
}

func (NetworkLoadBalancer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{0, 0}
}

// Type of the load balancer. Only external load balancers are currently available.
type NetworkLoadBalancer_Type int32

const (
	NetworkLoadBalancer_TYPE_UNSPECIFIED NetworkLoadBalancer_Type = 0
	// External network load balancer.
	NetworkLoadBalancer_EXTERNAL NetworkLoadBalancer_Type = 1
)

var NetworkLoadBalancer_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "EXTERNAL",
}

var NetworkLoadBalancer_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"EXTERNAL":         1,
}

func (x NetworkLoadBalancer_Type) String() string {
	return proto.EnumName(NetworkLoadBalancer_Type_name, int32(x))
}

func (NetworkLoadBalancer_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{0, 1}
}

// Type of session affinity. Only 5-tuple affinity is currently available.
// For more information, see [Load Balancer concepts](/load-balancer/concepts/).
type NetworkLoadBalancer_SessionAffinity int32

const (
	NetworkLoadBalancer_SESSION_AFFINITY_UNSPECIFIED NetworkLoadBalancer_SessionAffinity = 0
	// 5-tuple affinity.
	NetworkLoadBalancer_CLIENT_IP_PORT_PROTO NetworkLoadBalancer_SessionAffinity = 1
)

var NetworkLoadBalancer_SessionAffinity_name = map[int32]string{
	0: "SESSION_AFFINITY_UNSPECIFIED",
	1: "CLIENT_IP_PORT_PROTO",
}

var NetworkLoadBalancer_SessionAffinity_value = map[string]int32{
	"SESSION_AFFINITY_UNSPECIFIED": 0,
	"CLIENT_IP_PORT_PROTO":         1,
}

func (x NetworkLoadBalancer_SessionAffinity) String() string {
	return proto.EnumName(NetworkLoadBalancer_SessionAffinity_name, int32(x))
}

func (NetworkLoadBalancer_SessionAffinity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{0, 2}
}

// Network protocol to use.
type Listener_Protocol int32

const (
	Listener_PROTOCOL_UNSPECIFIED Listener_Protocol = 0
	Listener_TCP                  Listener_Protocol = 1
)

var Listener_Protocol_name = map[int32]string{
	0: "PROTOCOL_UNSPECIFIED",
	1: "TCP",
}

var Listener_Protocol_value = map[string]int32{
	"PROTOCOL_UNSPECIFIED": 0,
	"TCP":                  1,
}

func (x Listener_Protocol) String() string {
	return proto.EnumName(Listener_Protocol_name, int32(x))
}

func (Listener_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{2, 0}
}

// Status of the target.
type TargetState_Status int32

const (
	TargetState_STATUS_UNSPECIFIED TargetState_Status = 0
	// The network load balancer is setting up health checks for this target.
	TargetState_INITIAL TargetState_Status = 1
	// Health check passed and the target is ready to receive traffic.
	TargetState_HEALTHY TargetState_Status = 2
	// Health check failed and the target is not receiving traffic.
	TargetState_UNHEALTHY TargetState_Status = 3
	// Target is being deleted and the network load balancer is no longer sending traffic to this target.
	TargetState_DRAINING TargetState_Status = 4
	// The network load balancer is stopped and not performing health checks on this target.
	TargetState_INACTIVE TargetState_Status = 5
)

var TargetState_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "INITIAL",
	2: "HEALTHY",
	3: "UNHEALTHY",
	4: "DRAINING",
	5: "INACTIVE",
}

var TargetState_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"INITIAL":            1,
	"HEALTHY":            2,
	"UNHEALTHY":          3,
	"DRAINING":           4,
	"INACTIVE":           5,
}

func (x TargetState_Status) String() string {
	return proto.EnumName(TargetState_Status_name, int32(x))
}

func (TargetState_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{3, 0}
}

// A NetworkLoadBalancer resource. For more information, see [Network Load Balancer](/docs/load-balancer/concepts).
type NetworkLoadBalancer struct {
	// ID of the network load balancer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the network load balancer belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the network load balancer. The name is unique within the folder. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the network load balancer. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the region that the network load balancer belongs to.
	RegionId string `protobuf:"bytes,7,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Status of the network load balancer.
	Status NetworkLoadBalancer_Status `protobuf:"varint,9,opt,name=status,proto3,enum=yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_Status" json:"status,omitempty"`
	// Type of the network load balancer. Only external network load balancers are available now.
	Type NetworkLoadBalancer_Type `protobuf:"varint,10,opt,name=type,proto3,enum=yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_Type" json:"type,omitempty"`
	// Type of the session affinity. Only 5-tuple affinity is available now.
	SessionAffinity NetworkLoadBalancer_SessionAffinity `protobuf:"varint,11,opt,name=session_affinity,json=sessionAffinity,proto3,enum=yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_SessionAffinity" json:"session_affinity,omitempty"`
	// List of listeners for the network load balancer.
	Listeners []*Listener `protobuf:"bytes,12,rep,name=listeners,proto3" json:"listeners,omitempty"`
	// List of target groups attached to the network load balancer.
	AttachedTargetGroups []*AttachedTargetGroup `protobuf:"bytes,13,rep,name=attached_target_groups,json=attachedTargetGroups,proto3" json:"attached_target_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NetworkLoadBalancer) Reset()         { *m = NetworkLoadBalancer{} }
func (m *NetworkLoadBalancer) String() string { return proto.CompactTextString(m) }
func (*NetworkLoadBalancer) ProtoMessage()    {}
func (*NetworkLoadBalancer) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{0}
}

func (m *NetworkLoadBalancer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkLoadBalancer.Unmarshal(m, b)
}
func (m *NetworkLoadBalancer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkLoadBalancer.Marshal(b, m, deterministic)
}
func (m *NetworkLoadBalancer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkLoadBalancer.Merge(m, src)
}
func (m *NetworkLoadBalancer) XXX_Size() int {
	return xxx_messageInfo_NetworkLoadBalancer.Size(m)
}
func (m *NetworkLoadBalancer) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkLoadBalancer.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkLoadBalancer proto.InternalMessageInfo

func (m *NetworkLoadBalancer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkLoadBalancer) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *NetworkLoadBalancer) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NetworkLoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkLoadBalancer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkLoadBalancer) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NetworkLoadBalancer) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *NetworkLoadBalancer) GetStatus() NetworkLoadBalancer_Status {
	if m != nil {
		return m.Status
	}
	return NetworkLoadBalancer_STATUS_UNSPECIFIED
}

func (m *NetworkLoadBalancer) GetType() NetworkLoadBalancer_Type {
	if m != nil {
		return m.Type
	}
	return NetworkLoadBalancer_TYPE_UNSPECIFIED
}

func (m *NetworkLoadBalancer) GetSessionAffinity() NetworkLoadBalancer_SessionAffinity {
	if m != nil {
		return m.SessionAffinity
	}
	return NetworkLoadBalancer_SESSION_AFFINITY_UNSPECIFIED
}

func (m *NetworkLoadBalancer) GetListeners() []*Listener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *NetworkLoadBalancer) GetAttachedTargetGroups() []*AttachedTargetGroup {
	if m != nil {
		return m.AttachedTargetGroups
	}
	return nil
}

// An AttachedTargetGroup resource. For more information, see [Targets and groups](/docs/load-balancer/concepts/target-resources).
type AttachedTargetGroup struct {
	// ID of the target group.
	TargetGroupId string `protobuf:"bytes,1,opt,name=target_group_id,json=targetGroupId,proto3" json:"target_group_id,omitempty"`
	// A health check to perform on the target group.
	// For now we accept only one health check per AttachedTargetGroup.
	HealthChecks         []*HealthCheck `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AttachedTargetGroup) Reset()         { *m = AttachedTargetGroup{} }
func (m *AttachedTargetGroup) String() string { return proto.CompactTextString(m) }
func (*AttachedTargetGroup) ProtoMessage()    {}
func (*AttachedTargetGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{1}
}

func (m *AttachedTargetGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachedTargetGroup.Unmarshal(m, b)
}
func (m *AttachedTargetGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachedTargetGroup.Marshal(b, m, deterministic)
}
func (m *AttachedTargetGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachedTargetGroup.Merge(m, src)
}
func (m *AttachedTargetGroup) XXX_Size() int {
	return xxx_messageInfo_AttachedTargetGroup.Size(m)
}
func (m *AttachedTargetGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachedTargetGroup.DiscardUnknown(m)
}

var xxx_messageInfo_AttachedTargetGroup proto.InternalMessageInfo

func (m *AttachedTargetGroup) GetTargetGroupId() string {
	if m != nil {
		return m.TargetGroupId
	}
	return ""
}

func (m *AttachedTargetGroup) GetHealthChecks() []*HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

// A Listener resource. For more information, see [Listener](/docs/load-balancer/concepts/listener)
type Listener struct {
	// Name of the listener. The name must be unique for each listener on a single load balancer. 3-63 characters long.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// IP address for the listener.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Port.
	Port int64 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Network protocol for incoming traffic.
	Protocol Listener_Protocol `protobuf:"varint,4,opt,name=protocol,proto3,enum=yandex.cloud.loadbalancer.v1.Listener_Protocol" json:"protocol,omitempty"`
	// Port of a target.
	TargetPort           int64    `protobuf:"varint,5,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{2}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Listener) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Listener) GetProtocol() Listener_Protocol {
	if m != nil {
		return m.Protocol
	}
	return Listener_PROTOCOL_UNSPECIFIED
}

func (m *Listener) GetTargetPort() int64 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

// State of the target that was returned after the last health check.
type TargetState struct {
	// ID of the subnet that the target is connected to.
	SubnetId string `protobuf:"bytes,1,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// IP address of the target.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Status of the target.
	Status               TargetState_Status `protobuf:"varint,3,opt,name=status,proto3,enum=yandex.cloud.loadbalancer.v1.TargetState_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TargetState) Reset()         { *m = TargetState{} }
func (m *TargetState) String() string { return proto.CompactTextString(m) }
func (*TargetState) ProtoMessage()    {}
func (*TargetState) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d34a5ec4bd7fd4, []int{3}
}

func (m *TargetState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetState.Unmarshal(m, b)
}
func (m *TargetState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetState.Marshal(b, m, deterministic)
}
func (m *TargetState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetState.Merge(m, src)
}
func (m *TargetState) XXX_Size() int {
	return xxx_messageInfo_TargetState.Size(m)
}
func (m *TargetState) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetState.DiscardUnknown(m)
}

var xxx_messageInfo_TargetState proto.InternalMessageInfo

func (m *TargetState) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *TargetState) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TargetState) GetStatus() TargetState_Status {
	if m != nil {
		return m.Status
	}
	return TargetState_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.IpVersion", IpVersion_name, IpVersion_value)
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_Status", NetworkLoadBalancer_Status_name, NetworkLoadBalancer_Status_value)
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_Type", NetworkLoadBalancer_Type_name, NetworkLoadBalancer_Type_value)
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.NetworkLoadBalancer_SessionAffinity", NetworkLoadBalancer_SessionAffinity_name, NetworkLoadBalancer_SessionAffinity_value)
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.Listener_Protocol", Listener_Protocol_name, Listener_Protocol_value)
	proto.RegisterEnum("yandex.cloud.loadbalancer.v1.TargetState_Status", TargetState_Status_name, TargetState_Status_value)
	proto.RegisterType((*NetworkLoadBalancer)(nil), "yandex.cloud.loadbalancer.v1.NetworkLoadBalancer")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.loadbalancer.v1.NetworkLoadBalancer.LabelsEntry")
	proto.RegisterType((*AttachedTargetGroup)(nil), "yandex.cloud.loadbalancer.v1.AttachedTargetGroup")
	proto.RegisterType((*Listener)(nil), "yandex.cloud.loadbalancer.v1.Listener")
	proto.RegisterType((*TargetState)(nil), "yandex.cloud.loadbalancer.v1.TargetState")
}

func init() {
	proto.RegisterFile("yandex/cloud/loadbalancer/v1/network_load_balancer.proto", fileDescriptor_33d34a5ec4bd7fd4)
}

var fileDescriptor_33d34a5ec4bd7fd4 = []byte{
	// 952 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0xf5, 0xaf, 0x91, 0x7f, 0x88, 0x8d, 0x61, 0x10, 0x6a, 0x8a, 0x08, 0x3a, 0x14, 0x6e,
	0x00, 0x53, 0x91, 0x9b, 0x06, 0x4e, 0xd3, 0x1c, 0x68, 0x89, 0x8e, 0xb7, 0x55, 0x28, 0x62, 0x45,
	0x19, 0x75, 0x2f, 0xc4, 0x4a, 0x5c, 0x4b, 0x84, 0x69, 0x52, 0x20, 0x57, 0x6e, 0x75, 0x2a, 0xd0,
	0x63, 0xdf, 0xa0, 0xd7, 0xbe, 0x4c, 0xf3, 0x28, 0xed, 0x03, 0xf4, 0x5e, 0xec, 0x92, 0xb2, 0x69,
	0x27, 0x50, 0x9b, 0xde, 0x76, 0x66, 0xf6, 0xfb, 0x66, 0x67, 0xbf, 0xd9, 0x59, 0x38, 0x5e, 0xd1,
	0xd0, 0x63, 0x3f, 0x75, 0xa6, 0x41, 0xb4, 0xf4, 0x3a, 0x41, 0x44, 0xbd, 0x09, 0x0d, 0x68, 0x38,
	0x65, 0x71, 0xe7, 0xa6, 0xdb, 0x09, 0x19, 0xff, 0x31, 0x8a, 0xaf, 0x5c, 0xe1, 0x77, 0xd7, 0x01,
	0x7d, 0x11, 0x47, 0x3c, 0x42, 0x8f, 0x53, 0xa4, 0x2e, 0x91, 0x7a, 0x1e, 0xa9, 0xdf, 0x74, 0x9b,
	0x4f, 0x66, 0x51, 0x34, 0x0b, 0x58, 0x47, 0xee, 0x9d, 0x2c, 0x2f, 0x3b, 0xdc, 0xbf, 0x66, 0x09,
	0xa7, 0xd7, 0x8b, 0x14, 0xde, 0xfc, 0xec, 0x5e, 0xe2, 0x1b, 0x1a, 0xf8, 0x1e, 0xe5, 0x7e, 0x14,
	0x66, 0xe1, 0xce, 0xc6, 0x73, 0xcd, 0x19, 0x0d, 0xf8, 0xdc, 0x9d, 0xce, 0xd9, 0xf4, 0x2a, 0x05,
	0xb4, 0x7f, 0xab, 0xc1, 0x23, 0x2b, 0x3d, 0xee, 0x20, 0xa2, 0xde, 0x49, 0xb6, 0x1b, 0xed, 0x40,
	0xc1, 0xf7, 0x34, 0xa5, 0xa5, 0x1c, 0xd4, 0x49, 0xc1, 0xf7, 0xd0, 0xa7, 0x50, 0xbf, 0x8c, 0x02,
	0x8f, 0xc5, 0xae, 0xef, 0x69, 0x05, 0xe9, 0xae, 0xa5, 0x0e, 0xec, 0xa1, 0x97, 0x00, 0xd3, 0x98,
	0x51, 0xce, 0x3c, 0x97, 0x72, 0xad, 0xd8, 0x52, 0x0e, 0x1a, 0x47, 0x4d, 0x3d, 0x2d, 0x45, 0x5f,
	0x97, 0xa2, 0x3b, 0xeb, 0x52, 0x48, 0x3d, 0xdb, 0x6d, 0x70, 0x84, 0xa0, 0x14, 0xd2, 0x6b, 0xa6,
	0x95, 0x24, 0xa5, 0x5c, 0xa3, 0x16, 0x34, 0x3c, 0x96, 0x4c, 0x63, 0x7f, 0x21, 0x2a, 0xd3, 0xca,
	0x32, 0x94, 0x77, 0xa1, 0x31, 0x54, 0x02, 0x3a, 0x61, 0x41, 0xa2, 0x55, 0x5a, 0xc5, 0x83, 0xc6,
	0xd1, 0x6b, 0x7d, 0xd3, 0xad, 0xea, 0x1f, 0x28, 0x50, 0x1f, 0x48, 0xbc, 0x19, 0xf2, 0x78, 0x45,
	0x32, 0x32, 0x51, 0x64, 0xcc, 0x66, 0x7e, 0x14, 0x8a, 0x22, 0xab, 0x69, 0x91, 0xa9, 0x03, 0x7b,
	0xc8, 0x86, 0x4a, 0xc2, 0x29, 0x5f, 0x26, 0x5a, 0xbd, 0xa5, 0x1c, 0xec, 0x1c, 0x1d, 0x7f, 0x7c,
	0xce, 0x91, 0xc4, 0x93, 0x8c, 0x07, 0x7d, 0x0b, 0x25, 0xbe, 0x5a, 0x30, 0x0d, 0x24, 0xdf, 0x8b,
	0x8f, 0xe7, 0x73, 0x56, 0x0b, 0x46, 0x24, 0x07, 0x0a, 0x40, 0x4d, 0x58, 0x92, 0x88, 0xb3, 0xd3,
	0xcb, 0x4b, 0x3f, 0xf4, 0xf9, 0x4a, 0x6b, 0x48, 0x5e, 0xe3, 0x7f, 0x9c, 0x33, 0x65, 0x32, 0x32,
	0x22, 0xb2, 0x9b, 0xdc, 0x77, 0xa0, 0x3e, 0xd4, 0x03, 0x3f, 0xe1, 0x2c, 0x64, 0x71, 0xa2, 0x6d,
	0x49, 0x09, 0x3e, 0xdf, 0x9c, 0x66, 0x90, 0x6d, 0x27, 0x77, 0x40, 0x34, 0x83, 0x7d, 0xca, 0x39,
	0x9d, 0xce, 0x99, 0xe7, 0x72, 0x1a, 0xcf, 0x18, 0x77, 0x67, 0x71, 0xb4, 0x5c, 0x24, 0xda, 0xb6,
	0xa4, 0xec, 0x6e, 0xa6, 0x34, 0x32, 0xac, 0x23, 0xa1, 0x6f, 0x04, 0x92, 0xec, 0xd1, 0xf7, 0x9d,
	0x49, 0xf3, 0x25, 0x34, 0x72, 0x72, 0x23, 0x15, 0x8a, 0x57, 0x6c, 0x95, 0x35, 0xb7, 0x58, 0xa2,
	0x3d, 0x28, 0xdf, 0xd0, 0x60, 0xc9, 0xb2, 0xce, 0x4e, 0x8d, 0xaf, 0x0b, 0xc7, 0x4a, 0xfb, 0x67,
	0xa8, 0xa4, 0xaa, 0xa1, 0x7d, 0x40, 0x23, 0xc7, 0x70, 0xc6, 0x23, 0x77, 0x6c, 0x8d, 0x6c, 0xb3,
	0x87, 0x4f, 0xb1, 0xd9, 0x57, 0x3f, 0x41, 0x5b, 0x50, 0xeb, 0x11, 0xd3, 0x70, 0xb0, 0xf5, 0x46,
	0x55, 0x84, 0x35, 0x72, 0x0c, 0x22, 0xad, 0x02, 0x02, 0xa8, 0x18, 0x3d, 0x07, 0x9f, 0x9b, 0x6a,
	0x31, 0x8d, 0x0c, 0x6d, 0x5b, 0x44, 0x4a, 0xa8, 0x01, 0x55, 0x69, 0x99, 0x7d, 0xb5, 0x2c, 0x42,
	0x7d, 0x73, 0x60, 0x4a, 0x50, 0x45, 0x58, 0xd8, 0xca, 0x60, 0xd5, 0xf6, 0x53, 0x28, 0x09, 0x99,
	0xd1, 0x1e, 0xa8, 0xce, 0x85, 0x6d, 0xbe, 0x9f, 0xdc, 0xfc, 0xde, 0x31, 0x89, 0x65, 0x0c, 0x54,
	0xa5, 0xfd, 0x16, 0x76, 0x1f, 0x48, 0x87, 0x5a, 0xf0, 0x78, 0x64, 0x8e, 0x46, 0x78, 0x68, 0xb9,
	0xc6, 0xe9, 0x29, 0xb6, 0xb0, 0x73, 0xf1, 0x80, 0x42, 0x83, 0xbd, 0xde, 0x00, 0x9b, 0x96, 0xe3,
	0x62, 0xdb, 0xb5, 0x87, 0xc4, 0x71, 0x6d, 0x32, 0x74, 0x86, 0xaa, 0xd2, 0xfe, 0x5d, 0x81, 0x47,
	0x1f, 0xb8, 0x64, 0xf4, 0x1c, 0x76, 0xf3, 0x72, 0xb9, 0xeb, 0x41, 0x71, 0xb2, 0xf5, 0xe7, 0x1f,
	0x5d, 0xe5, 0xd7, 0x77, 0xdd, 0xd2, 0x37, 0xaf, 0xbf, 0x7a, 0x46, 0xb6, 0xf9, 0x1d, 0x06, 0x7b,
	0x68, 0x0c, 0xdb, 0xf9, 0xf9, 0x93, 0x68, 0x05, 0x29, 0xf2, 0x17, 0x9b, 0x45, 0x3e, 0x93, 0x90,
	0x9e, 0x40, 0x9c, 0x94, 0x7f, 0x79, 0xd7, 0x55, 0xba, 0x64, 0x6b, 0x7e, 0xe7, 0x4b, 0xda, 0x7f,
	0x29, 0x50, 0x5b, 0x37, 0xd7, 0xed, 0x34, 0x51, 0x72, 0xd3, 0x44, 0x83, 0x2a, 0xf5, 0xbc, 0x98,
	0x25, 0x49, 0xa6, 0xee, 0xda, 0x14, 0xbb, 0x17, 0x51, 0x9c, 0x0e, 0xac, 0x22, 0x91, 0x6b, 0xf4,
	0x1d, 0xd4, 0xe4, 0xc0, 0x9a, 0x46, 0x81, 0x9c, 0x49, 0x3b, 0x47, 0x9d, 0xff, 0xd6, 0xd8, 0xba,
	0x9d, 0xc1, 0xc8, 0x2d, 0x01, 0x7a, 0x02, 0x8d, 0xec, 0xa2, 0x64, 0x9e, 0xb2, 0xcc, 0x03, 0xa9,
	0xcb, 0x8e, 0x62, 0xde, 0x3e, 0x84, 0xda, 0x1a, 0x26, 0x74, 0x90, 0x17, 0xdf, 0x1b, 0x0e, 0x1e,
	0x28, 0x54, 0x85, 0xa2, 0xd3, 0xb3, 0x55, 0xa5, 0xfd, 0xb7, 0x02, 0x8d, 0x54, 0x08, 0xd1, 0x93,
	0x4c, 0xcc, 0xab, 0x64, 0x39, 0x09, 0x19, 0xbf, 0x95, 0x80, 0xd4, 0x52, 0x07, 0xf6, 0x36, 0xd4,
	0x7d, 0x76, 0x3b, 0xc9, 0x8a, 0xb2, 0xc2, 0x67, 0x9b, 0x2b, 0xcc, 0x65, 0x7c, 0x30, 0xc1, 0xda,
	0xec, 0x5f, 0x5f, 0x47, 0x03, 0xaa, 0xa2, 0xe9, 0xb0, 0xe8, 0x4f, 0x61, 0x9c, 0x99, 0xc6, 0xc0,
	0x39, 0xbb, 0x50, 0x0b, 0x68, 0x1b, 0xea, 0x63, 0x6b, 0x6d, 0xca, 0xe7, 0xd1, 0x27, 0x06, 0xb6,
	0xd2, 0xe7, 0x91, 0x7f, 0x03, 0xe5, 0xa7, 0xaf, 0xa0, 0x8e, 0x17, 0xe7, 0x2c, 0x16, 0x9d, 0x8d,
	0x9a, 0xb0, 0x8f, 0x6d, 0xf7, 0xdc, 0x24, 0xb2, 0xa9, 0xef, 0x67, 0xab, 0x41, 0x09, 0xdb, 0xe7,
	0xcf, 0x55, 0x25, 0x5b, 0xbd, 0x50, 0x0b, 0x27, 0xc3, 0x1f, 0xde, 0xce, 0x7c, 0x3e, 0x5f, 0x4e,
	0xf4, 0x69, 0x74, 0x9d, 0xfd, 0x8f, 0x87, 0xe9, 0xff, 0x38, 0x8b, 0x0e, 0x67, 0x2c, 0x94, 0x72,
	0x6d, 0xfc, 0x38, 0x5f, 0xe5, 0xed, 0x49, 0x45, 0x02, 0xbe, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x7c, 0x99, 0x24, 0x04, 0x08, 0x00, 0x00,
}
