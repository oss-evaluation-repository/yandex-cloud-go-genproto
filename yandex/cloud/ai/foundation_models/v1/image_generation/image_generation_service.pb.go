// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/ai/foundation_models/v1/image_generation/image_generation_service.proto

package image_generation

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request for the service to generate an image.
type ImageGenerationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The [ID of the model](/docs/foundation-models/concepts/yandexart/models) to be used for image generation.
	ModelUri string `protobuf:"bytes,1,opt,name=model_uri,json=modelUri,proto3" json:"model_uri,omitempty"`
	// A list of messages representing the context for the image generation model.
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	// Image generation options.
	GenerationOptions *ImageGenerationOptions `protobuf:"bytes,3,opt,name=generation_options,json=generationOptions,proto3" json:"generation_options,omitempty"`
}

func (x *ImageGenerationRequest) Reset() {
	*x = ImageGenerationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageGenerationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageGenerationRequest) ProtoMessage() {}

func (x *ImageGenerationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageGenerationRequest.ProtoReflect.Descriptor instead.
func (*ImageGenerationRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescGZIP(), []int{0}
}

func (x *ImageGenerationRequest) GetModelUri() string {
	if x != nil {
		return x.ModelUri
	}
	return ""
}

func (x *ImageGenerationRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ImageGenerationRequest) GetGenerationOptions() *ImageGenerationOptions {
	if x != nil {
		return x.GenerationOptions
	}
	return nil
}

// Response containing generated image.
type ImageGenerationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image is serialized as an array of bytes encoded in base64.
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// The model version changes with each new releases.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
}

func (x *ImageGenerationResponse) Reset() {
	*x = ImageGenerationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageGenerationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageGenerationResponse) ProtoMessage() {}

func (x *ImageGenerationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageGenerationResponse.ProtoReflect.Descriptor instead.
func (*ImageGenerationResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescGZIP(), []int{1}
}

func (x *ImageGenerationResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ImageGenerationResponse) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

var File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDesc = []byte{
	0x0a, 0x54, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x16, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x5a, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x7c, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x54, 0x0a, 0x17, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xef, 0x01, 0x0a, 0x1b,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x73, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcf, 0x01, 0x0a, 0x08,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0xb2, 0xd2, 0x2a, 0x19,
	0x12, 0x17, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a,
	0x01, 0x2a, 0x22, 0x29, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x42, 0xa7, 0x01,
	0x0a, 0x39, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x6a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69,
	0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescData = file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDesc
)

func file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescData)
	})
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDescData
}

var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_goTypes = []interface{}{
	(*ImageGenerationRequest)(nil),  // 0: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationRequest
	(*ImageGenerationResponse)(nil), // 1: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationResponse
	(*Message)(nil),                 // 2: yandex.cloud.ai.foundation_models.v1.image_generation.Message
	(*ImageGenerationOptions)(nil),  // 3: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationOptions
	(*operation.Operation)(nil),     // 4: yandex.cloud.operation.Operation
}
var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationRequest.messages:type_name -> yandex.cloud.ai.foundation_models.v1.image_generation.Message
	3, // 1: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationRequest.generation_options:type_name -> yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationOptions
	0, // 2: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationAsyncService.Generate:input_type -> yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationRequest
	4, // 3: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationAsyncService.Generate:output_type -> yandex.cloud.operation.Operation
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_init()
}
func file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_init() {
	if File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageGenerationRequest); i {
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
		file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageGenerationResponse); i {
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
			RawDescriptor: file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto = out.File
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_rawDesc = nil
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_goTypes = nil
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_service_proto_depIdxs = nil
}
