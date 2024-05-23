// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/video/v1/stream.proto

package video

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Stream status.
type Stream_StreamStatus int32

const (
	// Stream status unspecified.
	Stream_STREAM_STATUS_UNSPECIFIED Stream_StreamStatus = 0
	// Stream offline.
	Stream_OFFLINE Stream_StreamStatus = 1
	// Preparing the infrastructure for receiving video signal.
	Stream_PREPARING Stream_StreamStatus = 2
	// Everything is ready to launch stream.
	Stream_READY Stream_StreamStatus = 3
	// Stream onair.
	Stream_ONAIR Stream_StreamStatus = 4
	// Stream finished.
	Stream_FINISHED Stream_StreamStatus = 5
)

// Enum value maps for Stream_StreamStatus.
var (
	Stream_StreamStatus_name = map[int32]string{
		0: "STREAM_STATUS_UNSPECIFIED",
		1: "OFFLINE",
		2: "PREPARING",
		3: "READY",
		4: "ONAIR",
		5: "FINISHED",
	}
	Stream_StreamStatus_value = map[string]int32{
		"STREAM_STATUS_UNSPECIFIED": 0,
		"OFFLINE":                   1,
		"PREPARING":                 2,
		"READY":                     3,
		"ONAIR":                     4,
		"FINISHED":                  5,
	}
)

func (x Stream_StreamStatus) Enum() *Stream_StreamStatus {
	p := new(Stream_StreamStatus)
	*p = x
	return p
}

func (x Stream_StreamStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stream_StreamStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_video_v1_stream_proto_enumTypes[0].Descriptor()
}

func (Stream_StreamStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_video_v1_stream_proto_enumTypes[0]
}

func (x Stream_StreamStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stream_StreamStatus.Descriptor instead.
func (Stream_StreamStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_proto_rawDescGZIP(), []int{0, 0}
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the stream.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the channel where the stream was created.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// ID of the line to which stream is linked.
	LineId string `protobuf:"bytes,3,opt,name=line_id,json=lineId,proto3" json:"line_id,omitempty"`
	// Stream title.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Stream description.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,6,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
	// Stream status.
	Status Stream_StreamStatus `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.video.v1.Stream_StreamStatus" json:"status,omitempty"`
	// Stream start time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Stream publish time. Time when stream switched to ONAIR status.
	PublishTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	// Stream finish time.
	FinishTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// Stream type.
	//
	// Types that are assignable to StreamType:
	//
	//	*Stream_OnDemand
	//	*Stream_Schedule
	StreamType isStream_StreamType `protobuf_oneof:"stream_type"`
	// Time when stream was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time of last stream update.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Custom labels as “ key:value “ pairs. Maximum 64 per resource.
	Labels map[string]string `protobuf:"bytes,200,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Stream) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stream) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Stream) GetLineId() string {
	if x != nil {
		return x.LineId
	}
	return ""
}

func (x *Stream) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Stream) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Stream) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

func (x *Stream) GetStatus() Stream_StreamStatus {
	if x != nil {
		return x.Status
	}
	return Stream_STREAM_STATUS_UNSPECIFIED
}

func (x *Stream) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Stream) GetPublishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishTime
	}
	return nil
}

func (x *Stream) GetFinishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishTime
	}
	return nil
}

func (m *Stream) GetStreamType() isStream_StreamType {
	if m != nil {
		return m.StreamType
	}
	return nil
}

func (x *Stream) GetOnDemand() *OnDemand {
	if x, ok := x.GetStreamType().(*Stream_OnDemand); ok {
		return x.OnDemand
	}
	return nil
}

func (x *Stream) GetSchedule() *Schedule {
	if x, ok := x.GetStreamType().(*Stream_Schedule); ok {
		return x.Schedule
	}
	return nil
}

func (x *Stream) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Stream) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Stream) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isStream_StreamType interface {
	isStream_StreamType()
}

type Stream_OnDemand struct {
	// On demand stream. It starts immediately when a signal appears.
	OnDemand *OnDemand `protobuf:"bytes,1000,opt,name=on_demand,json=onDemand,proto3,oneof"`
}

type Stream_Schedule struct {
	// Schedule stream. Determines when to start receiving the signal or finish time.
	Schedule *Schedule `protobuf:"bytes,1001,opt,name=schedule,proto3,oneof"`
}

func (*Stream_OnDemand) isStream_StreamType() {}

func (*Stream_Schedule) isStream_StreamType() {}

// If "OnDemand" is used, client should start and finish explicitly.
type OnDemand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnDemand) Reset() {
	*x = OnDemand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnDemand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDemand) ProtoMessage() {}

func (x *OnDemand) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDemand.ProtoReflect.Descriptor instead.
func (*OnDemand) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_proto_rawDescGZIP(), []int{1}
}

// If "Schedule" is used, stream automatically start and finish at this time.
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_proto_rawDescGZIP(), []int{2}
}

func (x *Schedule) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Schedule) GetFinishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishTime
	}
	return nil
}

var File_yandex_cloud_video_v1_stream_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x07, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6e,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x53,
	0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46,
	0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x45, 0x50, 0x41,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x4e, 0x41, 0x49, 0x52, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x05, 0x08, 0x66, 0x10, 0xc8, 0x01,
	0x4a, 0x06, 0x08, 0xc9, 0x01, 0x10, 0xe8, 0x07, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x64, 0x4a, 0x04,
	0x08, 0x07, 0x10, 0x08, 0x22, 0x0a, 0x0a, 0x08, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x82, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x5c, 0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x76, 0x31, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_stream_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_stream_proto_rawDescData = file_yandex_cloud_video_v1_stream_proto_rawDesc
)

func file_yandex_cloud_video_v1_stream_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_stream_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_stream_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_stream_proto_rawDescData
}

var file_yandex_cloud_video_v1_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_video_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_video_v1_stream_proto_goTypes = []interface{}{
	(Stream_StreamStatus)(0),      // 0: yandex.cloud.video.v1.Stream.StreamStatus
	(*Stream)(nil),                // 1: yandex.cloud.video.v1.Stream
	(*OnDemand)(nil),              // 2: yandex.cloud.video.v1.OnDemand
	(*Schedule)(nil),              // 3: yandex.cloud.video.v1.Schedule
	nil,                           // 4: yandex.cloud.video.v1.Stream.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_video_v1_stream_proto_depIdxs = []int32{
	0,  // 0: yandex.cloud.video.v1.Stream.status:type_name -> yandex.cloud.video.v1.Stream.StreamStatus
	5,  // 1: yandex.cloud.video.v1.Stream.start_time:type_name -> google.protobuf.Timestamp
	5,  // 2: yandex.cloud.video.v1.Stream.publish_time:type_name -> google.protobuf.Timestamp
	5,  // 3: yandex.cloud.video.v1.Stream.finish_time:type_name -> google.protobuf.Timestamp
	2,  // 4: yandex.cloud.video.v1.Stream.on_demand:type_name -> yandex.cloud.video.v1.OnDemand
	3,  // 5: yandex.cloud.video.v1.Stream.schedule:type_name -> yandex.cloud.video.v1.Schedule
	5,  // 6: yandex.cloud.video.v1.Stream.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: yandex.cloud.video.v1.Stream.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 8: yandex.cloud.video.v1.Stream.labels:type_name -> yandex.cloud.video.v1.Stream.LabelsEntry
	5,  // 9: yandex.cloud.video.v1.Schedule.start_time:type_name -> google.protobuf.Timestamp
	5,  // 10: yandex.cloud.video.v1.Schedule.finish_time:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_stream_proto_init() }
func file_yandex_cloud_video_v1_stream_proto_init() {
	if File_yandex_cloud_video_v1_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_yandex_cloud_video_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnDemand); i {
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
		file_yandex_cloud_video_v1_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
	file_yandex_cloud_video_v1_stream_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Stream_OnDemand)(nil),
		(*Stream_Schedule)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_video_v1_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_video_v1_stream_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_stream_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_video_v1_stream_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_video_v1_stream_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_stream_proto = out.File
	file_yandex_cloud_video_v1_stream_proto_rawDesc = nil
	file_yandex_cloud_video_v1_stream_proto_goTypes = nil
	file_yandex_cloud_video_v1_stream_proto_depIdxs = nil
}
