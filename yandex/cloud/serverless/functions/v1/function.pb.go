// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/serverless/functions/v1/function.proto

package functions

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Function_Status int32

const (
	Function_STATUS_UNSPECIFIED Function_Status = 0
	// Function is being created.
	Function_CREATING Function_Status = 1
	// Function is ready to be invoked.
	Function_ACTIVE Function_Status = 2
	// Function is being deleted.
	Function_DELETING Function_Status = 3
	// Function failed.
	Function_ERROR Function_Status = 4
)

// Enum value maps for Function_Status.
var (
	Function_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
		4: "ERROR",
	}
	Function_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"DELETING":           3,
		"ERROR":              4,
	}
)

func (x Function_Status) Enum() *Function_Status {
	p := new(Function_Status)
	*p = x
	return p
}

func (x Function_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Function_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes[0].Descriptor()
}

func (Function_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes[0]
}

func (x Function_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Function_Status.Descriptor instead.
func (Function_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{0, 0}
}

type Version_Status int32

const (
	Version_STATUS_UNSPECIFIED Version_Status = 0
	// Version is being created.
	Version_CREATING Version_Status = 1
	// Version is ready to use.
	Version_ACTIVE Version_Status = 2
)

// Enum value maps for Version_Status.
var (
	Version_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
	}
	Version_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
	}
)

func (x Version_Status) Enum() *Version_Status {
	p := new(Version_Status)
	*p = x
	return p
}

func (x Version_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes[1].Descriptor()
}

func (Version_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes[1]
}

func (x Version_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_Status.Descriptor instead.
func (Version_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{1, 0}
}

// A serverless function. For details about the concept, see [Functions](/docs/functions/concepts/function).
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the function. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the function belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp for the function.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the function. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the function.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Function labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the log group for the function.
	LogGroupId string `protobuf:"bytes,7,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// URL that needs to be requested to invoke the function.
	HttpInvokeUrl string `protobuf:"bytes,8,opt,name=http_invoke_url,json=httpInvokeUrl,proto3" json:"http_invoke_url,omitempty"`
	// Status of the function.
	Status Function_Status `protobuf:"varint,9,opt,name=status,proto3,enum=yandex.cloud.serverless.functions.v1.Function_Status" json:"status,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{0}
}

func (x *Function) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Function) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Function) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Function) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Function) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *Function) GetHttpInvokeUrl() string {
	if x != nil {
		return x.HttpInvokeUrl
	}
	return ""
}

func (x *Function) GetStatus() Function_Status {
	if x != nil {
		return x.Status
	}
	return Function_STATUS_UNSPECIFIED
}

// Version of a function. For details about the concept, see [Function versions](/docs/functions/concepts/function#version).
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the version.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the function that the version belongs to.
	FunctionId string `protobuf:"bytes,2,opt,name=function_id,json=functionId,proto3" json:"function_id,omitempty"`
	// Description of the version.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Creation timestamp for the version.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the runtime environment for the function.
	//
	// Supported environments and their identifiers are listed in the [Runtime environments](/docs/functions/concepts/runtime).
	Runtime string `protobuf:"bytes,6,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Entrypoint for the function: the name of the function to be called as the handler.
	//
	// Specified in the format `<function file name>.<handler name>`, for example, `index.myFunction`.
	Entrypoint string `protobuf:"bytes,7,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	// Resources allocated to the version.
	Resources *Resources `protobuf:"bytes,8,opt,name=resources,proto3" json:"resources,omitempty"`
	// Timeout for the execution of the version.
	//
	// If the timeout is exceeded, Cloud Functions responds with a 504 HTTP code.
	ExecutionTimeout *duration.Duration `protobuf:"bytes,9,opt,name=execution_timeout,json=executionTimeout,proto3" json:"execution_timeout,omitempty"`
	// ID of the service account associated with the version.
	ServiceAccountId string `protobuf:"bytes,10,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Final size of the deployment package after unpacking.
	ImageSize int64 `protobuf:"varint,12,opt,name=image_size,json=imageSize,proto3" json:"image_size,omitempty"`
	// Status of the version.
	Status Version_Status `protobuf:"varint,13,opt,name=status,proto3,enum=yandex.cloud.serverless.functions.v1.Version_Status" json:"status,omitempty"`
	// Version tags. For details, see [Version tag](/docs/functions/concepts/function#tag).
	Tags []string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	// ID of the log group for the version.
	LogGroupId string `protobuf:"bytes,15,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// Environment settings for the version.
	Environment map[string]string `protobuf:"bytes,16,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Network access. If specified the version will be attached to specified network/subnet(s).
	Connectivity *Connectivity `protobuf:"bytes,17,opt,name=connectivity,proto3" json:"connectivity,omitempty"`
	// Additional service accounts to be used by the version.
	NamedServiceAccounts map[string]string `protobuf:"bytes,18,rep,name=named_service_accounts,json=namedServiceAccounts,proto3" json:"named_service_accounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{1}
}

func (x *Version) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Version) GetFunctionId() string {
	if x != nil {
		return x.FunctionId
	}
	return ""
}

func (x *Version) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Version) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Version) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Version) GetEntrypoint() string {
	if x != nil {
		return x.Entrypoint
	}
	return ""
}

func (x *Version) GetResources() *Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Version) GetExecutionTimeout() *duration.Duration {
	if x != nil {
		return x.ExecutionTimeout
	}
	return nil
}

func (x *Version) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Version) GetImageSize() int64 {
	if x != nil {
		return x.ImageSize
	}
	return 0
}

func (x *Version) GetStatus() Version_Status {
	if x != nil {
		return x.Status
	}
	return Version_STATUS_UNSPECIFIED
}

func (x *Version) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Version) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *Version) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *Version) GetConnectivity() *Connectivity {
	if x != nil {
		return x.Connectivity
	}
	return nil
}

func (x *Version) GetNamedServiceAccounts() map[string]string {
	if x != nil {
		return x.NamedServiceAccounts
	}
	return nil
}

// Resources allocated to a version.
type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount of memory available to the version, specified in bytes.
	Memory int64 `protobuf:"varint,1,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{2}
}

func (x *Resources) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

// Version deployment package.
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the bucket that stores the code for the version.
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	// Name of the object in the bucket that stores the code for the version.
	ObjectName string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	// SHA256 hash of the version deployment package.
	Sha256 string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{3}
}

func (x *Package) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *Package) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *Package) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

// Version connectivity specification.
type Connectivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Network the version will have access to.
	// It's essential to specify network with subnets in all availability zones.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Complete list of subnets (from the same network) the version can be attached to.
	// It's essential to specify at least one subnet for each availability zones.
	SubnetId []string `protobuf:"bytes,2,rep,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
}

func (x *Connectivity) Reset() {
	*x = Connectivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connectivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connectivity) ProtoMessage() {}

func (x *Connectivity) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connectivity.ProtoReflect.Descriptor instead.
func (*Connectivity) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP(), []int{4}
}

func (x *Connectivity) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Connectivity) GetSubnetId() []string {
	if x != nil {
		return x.SubnetId
	}
	return nil
}

var File_yandex_cloud_serverless_functions_v1_function_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_functions_v1_function_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x04, 0x0a, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0x8a, 0xc8, 0x31, 0x04, 0x33, 0x2d, 0x36, 0x33, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x30, 0x2d, 0x32, 0x35, 0x36, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0x82, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36,
	0x34, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x22, 0xc2, 0x08, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x30, 0x2d, 0x32, 0x35, 0x36, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x12, 0x7d, 0x0a, 0x16, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x47, 0x0a, 0x19, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x22, 0x3d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x18, 0xfa, 0xc7, 0x31, 0x14, 0x31, 0x33, 0x34, 0x32, 0x31, 0x37,
	0x37, 0x32, 0x38, 0x2d, 0x32, 0x31, 0x34, 0x37, 0x34, 0x38, 0x33, 0x36, 0x34, 0x38, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x6f, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8,
	0xc7, 0x31, 0x01, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x49, 0x64, 0x42, 0x7e, 0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_serverless_functions_v1_function_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_functions_v1_function_proto_rawDescData = file_yandex_cloud_serverless_functions_v1_function_proto_rawDesc
)

func file_yandex_cloud_serverless_functions_v1_function_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_functions_v1_function_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_functions_v1_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_functions_v1_function_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_functions_v1_function_proto_rawDescData
}

var file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_yandex_cloud_serverless_functions_v1_function_proto_goTypes = []interface{}{
	(Function_Status)(0),        // 0: yandex.cloud.serverless.functions.v1.Function.Status
	(Version_Status)(0),         // 1: yandex.cloud.serverless.functions.v1.Version.Status
	(*Function)(nil),            // 2: yandex.cloud.serverless.functions.v1.Function
	(*Version)(nil),             // 3: yandex.cloud.serverless.functions.v1.Version
	(*Resources)(nil),           // 4: yandex.cloud.serverless.functions.v1.Resources
	(*Package)(nil),             // 5: yandex.cloud.serverless.functions.v1.Package
	(*Connectivity)(nil),        // 6: yandex.cloud.serverless.functions.v1.Connectivity
	nil,                         // 7: yandex.cloud.serverless.functions.v1.Function.LabelsEntry
	nil,                         // 8: yandex.cloud.serverless.functions.v1.Version.EnvironmentEntry
	nil,                         // 9: yandex.cloud.serverless.functions.v1.Version.NamedServiceAccountsEntry
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 11: google.protobuf.Duration
}
var file_yandex_cloud_serverless_functions_v1_function_proto_depIdxs = []int32{
	10, // 0: yandex.cloud.serverless.functions.v1.Function.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: yandex.cloud.serverless.functions.v1.Function.labels:type_name -> yandex.cloud.serverless.functions.v1.Function.LabelsEntry
	0,  // 2: yandex.cloud.serverless.functions.v1.Function.status:type_name -> yandex.cloud.serverless.functions.v1.Function.Status
	10, // 3: yandex.cloud.serverless.functions.v1.Version.created_at:type_name -> google.protobuf.Timestamp
	4,  // 4: yandex.cloud.serverless.functions.v1.Version.resources:type_name -> yandex.cloud.serverless.functions.v1.Resources
	11, // 5: yandex.cloud.serverless.functions.v1.Version.execution_timeout:type_name -> google.protobuf.Duration
	1,  // 6: yandex.cloud.serverless.functions.v1.Version.status:type_name -> yandex.cloud.serverless.functions.v1.Version.Status
	8,  // 7: yandex.cloud.serverless.functions.v1.Version.environment:type_name -> yandex.cloud.serverless.functions.v1.Version.EnvironmentEntry
	6,  // 8: yandex.cloud.serverless.functions.v1.Version.connectivity:type_name -> yandex.cloud.serverless.functions.v1.Connectivity
	9,  // 9: yandex.cloud.serverless.functions.v1.Version.named_service_accounts:type_name -> yandex.cloud.serverless.functions.v1.Version.NamedServiceAccountsEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_functions_v1_function_proto_init() }
func file_yandex_cloud_serverless_functions_v1_function_proto_init() {
	if File_yandex_cloud_serverless_functions_v1_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
		file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connectivity); i {
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
			RawDescriptor: file_yandex_cloud_serverless_functions_v1_function_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_functions_v1_function_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_functions_v1_function_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_serverless_functions_v1_function_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_serverless_functions_v1_function_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_functions_v1_function_proto = out.File
	file_yandex_cloud_serverless_functions_v1_function_proto_rawDesc = nil
	file_yandex_cloud_serverless_functions_v1_function_proto_goTypes = nil
	file_yandex_cloud_serverless_functions_v1_function_proto_depIdxs = nil
}
