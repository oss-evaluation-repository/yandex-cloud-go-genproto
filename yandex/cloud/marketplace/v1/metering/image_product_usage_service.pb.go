// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/marketplace/v1/metering/image_product_usage_service.proto

package metering

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type WriteImageProductUsageRequest struct {
	// Checks whether you have the access required for the emit usage
	ValidateOnly bool `protobuf:"varint,1,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Marketplace Product's ID
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// List of product usage records (up to 25 pet request)
	UsageRecords         []*UsageRecord `protobuf:"bytes,3,rep,name=usage_records,json=usageRecords,proto3" json:"usage_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WriteImageProductUsageRequest) Reset()         { *m = WriteImageProductUsageRequest{} }
func (m *WriteImageProductUsageRequest) String() string { return proto.CompactTextString(m) }
func (*WriteImageProductUsageRequest) ProtoMessage()    {}
func (*WriteImageProductUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_613da1456bc6e886, []int{0}
}

func (m *WriteImageProductUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteImageProductUsageRequest.Unmarshal(m, b)
}
func (m *WriteImageProductUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteImageProductUsageRequest.Marshal(b, m, deterministic)
}
func (m *WriteImageProductUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteImageProductUsageRequest.Merge(m, src)
}
func (m *WriteImageProductUsageRequest) XXX_Size() int {
	return xxx_messageInfo_WriteImageProductUsageRequest.Size(m)
}
func (m *WriteImageProductUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteImageProductUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteImageProductUsageRequest proto.InternalMessageInfo

func (m *WriteImageProductUsageRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

func (m *WriteImageProductUsageRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *WriteImageProductUsageRequest) GetUsageRecords() []*UsageRecord {
	if m != nil {
		return m.UsageRecords
	}
	return nil
}

type WriteImageProductUsageResponse struct {
	// List of accepted product usage records
	Accepted []*AcceptedUsageRecord `protobuf:"bytes,1,rep,name=accepted,proto3" json:"accepted,omitempty"`
	// List of rejected product usage records (with reason)
	Rejected             []*RejectedUsageRecord `protobuf:"bytes,2,rep,name=rejected,proto3" json:"rejected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WriteImageProductUsageResponse) Reset()         { *m = WriteImageProductUsageResponse{} }
func (m *WriteImageProductUsageResponse) String() string { return proto.CompactTextString(m) }
func (*WriteImageProductUsageResponse) ProtoMessage()    {}
func (*WriteImageProductUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_613da1456bc6e886, []int{1}
}

func (m *WriteImageProductUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteImageProductUsageResponse.Unmarshal(m, b)
}
func (m *WriteImageProductUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteImageProductUsageResponse.Marshal(b, m, deterministic)
}
func (m *WriteImageProductUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteImageProductUsageResponse.Merge(m, src)
}
func (m *WriteImageProductUsageResponse) XXX_Size() int {
	return xxx_messageInfo_WriteImageProductUsageResponse.Size(m)
}
func (m *WriteImageProductUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteImageProductUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteImageProductUsageResponse proto.InternalMessageInfo

func (m *WriteImageProductUsageResponse) GetAccepted() []*AcceptedUsageRecord {
	if m != nil {
		return m.Accepted
	}
	return nil
}

func (m *WriteImageProductUsageResponse) GetRejected() []*RejectedUsageRecord {
	if m != nil {
		return m.Rejected
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteImageProductUsageRequest)(nil), "yandex.cloud.marketplace.v1.metering.WriteImageProductUsageRequest")
	proto.RegisterType((*WriteImageProductUsageResponse)(nil), "yandex.cloud.marketplace.v1.metering.WriteImageProductUsageResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/marketplace/v1/metering/image_product_usage_service.proto", fileDescriptor_613da1456bc6e886)
}

var fileDescriptor_613da1456bc6e886 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0xde, 0x81, 0x7a, 0xa6, 0xb7, 0x64, 0x8a, 0x2a, 0x0e, 0x55, 0x85, 0xa1, 0x02,
	0xd5, 0xbe, 0x14, 0x55, 0x08, 0x0e, 0x06, 0x0a, 0x42, 0xba, 0x09, 0x08, 0x3a, 0x21, 0xb1, 0x14,
	0x9f, 0xfd, 0x14, 0x0c, 0xa9, 0x1d, 0x6c, 0xa7, 0xd0, 0x81, 0xe5, 0x46, 0x56, 0x3e, 0x14, 0x37,
	0xc1, 0x70, 0x5f, 0x81, 0x81, 0x81, 0x0f, 0x81, 0x12, 0x27, 0x90, 0x0a, 0x5d, 0x15, 0x74, 0x5b,
	0xe2, 0x97, 0xff, 0xef, 0xfd, 0xff, 0x2f, 0x7e, 0xf8, 0xc9, 0x8a, 0x29, 0x01, 0x1f, 0x29, 0x4f,
	0x75, 0x2e, 0xe8, 0x82, 0x99, 0x77, 0xe0, 0xb2, 0x94, 0x71, 0xa0, 0xcb, 0x88, 0x2e, 0xc0, 0x81,
	0x91, 0x2a, 0xa1, 0x72, 0xc1, 0x12, 0x98, 0x67, 0x46, 0x8b, 0x9c, 0xbb, 0x79, 0x6e, 0x8b, 0x37,
	0x0b, 0x66, 0x29, 0x39, 0x90, 0xcc, 0x68, 0xa7, 0x83, 0x1b, 0x9e, 0x43, 0x4a, 0x0e, 0x69, 0x70,
	0xc8, 0x32, 0x22, 0x35, 0xa7, 0x7f, 0x35, 0xd1, 0x3a, 0x49, 0x81, 0xb2, 0x4c, 0x52, 0xa6, 0x94,
	0x76, 0xcc, 0x49, 0xad, 0xac, 0x67, 0xf4, 0xf7, 0xd6, 0xbc, 0x2c, 0x59, 0x2a, 0x45, 0x59, 0xaf,
	0xca, 0x77, 0x5a, 0x59, 0xf5, 0xe6, 0x0c, 0x70, 0x6d, 0x84, 0x17, 0x0e, 0xcf, 0x10, 0xde, 0x7b,
	0x69, 0xa4, 0x83, 0xc3, 0x22, 0xc6, 0x33, 0x9f, 0xe2, 0xa8, 0xf8, 0x2e, 0x86, 0xf7, 0x39, 0x58,
	0x17, 0x5c, 0xc7, 0xbb, 0x55, 0x3b, 0x98, 0x6b, 0x95, 0xae, 0x42, 0x34, 0x40, 0xa3, 0x6e, 0xdc,
	0xab, 0x0f, 0x9f, 0xaa, 0x74, 0x15, 0xdc, 0xc2, 0xb8, 0x9e, 0x80, 0x14, 0x61, 0x67, 0x80, 0x46,
	0x3b, 0xb3, 0xde, 0xcf, 0xaf, 0x11, 0xfa, 0x7c, 0x1a, 0x6d, 0xdf, 0x7f, 0x30, 0xdd, 0x8f, 0x77,
	0xaa, 0xfa, 0xa1, 0x08, 0x5e, 0xe3, 0xdd, 0xa6, 0x13, 0x1b, 0x6e, 0x0d, 0xb6, 0x46, 0x57, 0x26,
	0x11, 0x69, 0x33, 0x27, 0x52, 0x99, 0x2b, 0x94, 0xb3, 0xee, 0xc9, 0x69, 0xb4, 0x1d, 0x8d, 0x27,
	0xd3, 0xb8, 0x97, 0xff, 0x3d, 0xb6, 0xc3, 0x6f, 0x08, 0x5f, 0x3b, 0x2f, 0x95, 0xcd, 0xb4, 0xb2,
	0x10, 0x1c, 0xe1, 0x2e, 0xe3, 0x1c, 0x32, 0x07, 0x22, 0x44, 0x65, 0xff, 0xbb, 0xed, 0xfa, 0x3f,
	0xac, 0x54, 0x0d, 0x1f, 0xf1, 0x1f, 0x54, 0x81, 0x35, 0xf0, 0x16, 0x78, 0x81, 0xed, 0xfc, 0x0f,
	0x36, 0xae, 0x54, 0x6b, 0xd8, 0x1a, 0x35, 0xf9, 0x85, 0x70, 0xf8, 0x4f, 0x96, 0x17, 0xfe, 0x96,
	0x05, 0xdf, 0x11, 0xbe, 0x54, 0xa6, 0x0d, 0x1e, 0xb5, 0xeb, 0xb5, 0xf1, 0x87, 0xf7, 0x1f, 0x5f,
	0x0c, 0xe2, 0xe7, 0x3b, 0x3c, 0x38, 0x39, 0xfb, 0xf1, 0xa5, 0x33, 0x1d, 0xee, 0x6f, 0x5e, 0x9c,
	0xa6, 0x96, 0x7e, 0x28, 0x90, 0xf7, 0xd0, 0xcd, 0xd9, 0x27, 0x3c, 0x5a, 0xf3, 0xc0, 0x32, 0x79,
	0x9e, 0x8f, 0x57, 0xcf, 0x13, 0xe9, 0xde, 0xe4, 0xc7, 0x84, 0xeb, 0x05, 0xf5, 0xa2, 0xb1, 0xdf,
	0x82, 0x44, 0x8f, 0x13, 0x50, 0xe5, 0x35, 0xa7, 0x6d, 0xd6, 0xe3, 0xa0, 0x7e, 0x38, 0xbe, 0x5c,
	0x8a, 0x6e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x7d, 0x1b, 0xcc, 0x01, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageProductUsageServiceClient is the client API for ImageProductUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageProductUsageServiceClient interface {
	// Writes image product's usage (authenticated by user's service account)
	Write(ctx context.Context, in *WriteImageProductUsageRequest, opts ...grpc.CallOption) (*WriteImageProductUsageResponse, error)
}

type imageProductUsageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageProductUsageServiceClient(cc *grpc.ClientConn) ImageProductUsageServiceClient {
	return &imageProductUsageServiceClient{cc}
}

func (c *imageProductUsageServiceClient) Write(ctx context.Context, in *WriteImageProductUsageRequest, opts ...grpc.CallOption) (*WriteImageProductUsageResponse, error) {
	out := new(WriteImageProductUsageResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.marketplace.v1.metering.ImageProductUsageService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageProductUsageServiceServer is the server API for ImageProductUsageService service.
type ImageProductUsageServiceServer interface {
	// Writes image product's usage (authenticated by user's service account)
	Write(context.Context, *WriteImageProductUsageRequest) (*WriteImageProductUsageResponse, error)
}

// UnimplementedImageProductUsageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageProductUsageServiceServer struct {
}

func (*UnimplementedImageProductUsageServiceServer) Write(ctx context.Context, req *WriteImageProductUsageRequest) (*WriteImageProductUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterImageProductUsageServiceServer(s *grpc.Server, srv ImageProductUsageServiceServer) {
	s.RegisterService(&_ImageProductUsageService_serviceDesc, srv)
}

func _ImageProductUsageService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteImageProductUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProductUsageServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.marketplace.v1.metering.ImageProductUsageService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProductUsageServiceServer).Write(ctx, req.(*WriteImageProductUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageProductUsageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.marketplace.v1.metering.ImageProductUsageService",
	HandlerType: (*ImageProductUsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _ImageProductUsageService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/marketplace/v1/metering/image_product_usage_service.proto",
}
