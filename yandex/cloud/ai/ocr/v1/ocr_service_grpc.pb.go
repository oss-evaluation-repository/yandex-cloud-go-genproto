// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/ai/ocr/v1/ocr_service.proto

package ocr

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TextRecognitionService_Recognize_FullMethodName = "/yandex.cloud.ai.ocr.v1.TextRecognitionService/Recognize"
)

// TextRecognitionServiceClient is the client API for TextRecognitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextRecognitionServiceClient interface {
	// To send the image for text recognition.
	Recognize(ctx context.Context, in *RecognizeTextRequest, opts ...grpc.CallOption) (TextRecognitionService_RecognizeClient, error)
}

type textRecognitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextRecognitionServiceClient(cc grpc.ClientConnInterface) TextRecognitionServiceClient {
	return &textRecognitionServiceClient{cc}
}

func (c *textRecognitionServiceClient) Recognize(ctx context.Context, in *RecognizeTextRequest, opts ...grpc.CallOption) (TextRecognitionService_RecognizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TextRecognitionService_ServiceDesc.Streams[0], TextRecognitionService_Recognize_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &textRecognitionServiceRecognizeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextRecognitionService_RecognizeClient interface {
	Recv() (*RecognizeTextResponse, error)
	grpc.ClientStream
}

type textRecognitionServiceRecognizeClient struct {
	grpc.ClientStream
}

func (x *textRecognitionServiceRecognizeClient) Recv() (*RecognizeTextResponse, error) {
	m := new(RecognizeTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextRecognitionServiceServer is the server API for TextRecognitionService service.
// All implementations should embed UnimplementedTextRecognitionServiceServer
// for forward compatibility
type TextRecognitionServiceServer interface {
	// To send the image for text recognition.
	Recognize(*RecognizeTextRequest, TextRecognitionService_RecognizeServer) error
}

// UnimplementedTextRecognitionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTextRecognitionServiceServer struct {
}

func (UnimplementedTextRecognitionServiceServer) Recognize(*RecognizeTextRequest, TextRecognitionService_RecognizeServer) error {
	return status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}

// UnsafeTextRecognitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextRecognitionServiceServer will
// result in compilation errors.
type UnsafeTextRecognitionServiceServer interface {
	mustEmbedUnimplementedTextRecognitionServiceServer()
}

func RegisterTextRecognitionServiceServer(s grpc.ServiceRegistrar, srv TextRecognitionServiceServer) {
	s.RegisterService(&TextRecognitionService_ServiceDesc, srv)
}

func _TextRecognitionService_Recognize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecognizeTextRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextRecognitionServiceServer).Recognize(m, &textRecognitionServiceRecognizeServer{stream})
}

type TextRecognitionService_RecognizeServer interface {
	Send(*RecognizeTextResponse) error
	grpc.ServerStream
}

type textRecognitionServiceRecognizeServer struct {
	grpc.ServerStream
}

func (x *textRecognitionServiceRecognizeServer) Send(m *RecognizeTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TextRecognitionService_ServiceDesc is the grpc.ServiceDesc for TextRecognitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextRecognitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.ocr.v1.TextRecognitionService",
	HandlerType: (*TextRecognitionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Recognize",
			Handler:       _TextRecognitionService_Recognize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yandex/cloud/ai/ocr/v1/ocr_service.proto",
}

const (
	TextRecognitionAsyncService_Recognize_FullMethodName      = "/yandex.cloud.ai.ocr.v1.TextRecognitionAsyncService/Recognize"
	TextRecognitionAsyncService_GetRecognition_FullMethodName = "/yandex.cloud.ai.ocr.v1.TextRecognitionAsyncService/GetRecognition"
)

// TextRecognitionAsyncServiceClient is the client API for TextRecognitionAsyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextRecognitionAsyncServiceClient interface {
	// To send the image for asynchronous text recognition.
	Recognize(ctx context.Context, in *RecognizeTextRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// To get recognition results.
	GetRecognition(ctx context.Context, in *GetRecognitionRequest, opts ...grpc.CallOption) (TextRecognitionAsyncService_GetRecognitionClient, error)
}

type textRecognitionAsyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextRecognitionAsyncServiceClient(cc grpc.ClientConnInterface) TextRecognitionAsyncServiceClient {
	return &textRecognitionAsyncServiceClient{cc}
}

func (c *textRecognitionAsyncServiceClient) Recognize(ctx context.Context, in *RecognizeTextRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TextRecognitionAsyncService_Recognize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textRecognitionAsyncServiceClient) GetRecognition(ctx context.Context, in *GetRecognitionRequest, opts ...grpc.CallOption) (TextRecognitionAsyncService_GetRecognitionClient, error) {
	stream, err := c.cc.NewStream(ctx, &TextRecognitionAsyncService_ServiceDesc.Streams[0], TextRecognitionAsyncService_GetRecognition_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &textRecognitionAsyncServiceGetRecognitionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextRecognitionAsyncService_GetRecognitionClient interface {
	Recv() (*RecognizeTextResponse, error)
	grpc.ClientStream
}

type textRecognitionAsyncServiceGetRecognitionClient struct {
	grpc.ClientStream
}

func (x *textRecognitionAsyncServiceGetRecognitionClient) Recv() (*RecognizeTextResponse, error) {
	m := new(RecognizeTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextRecognitionAsyncServiceServer is the server API for TextRecognitionAsyncService service.
// All implementations should embed UnimplementedTextRecognitionAsyncServiceServer
// for forward compatibility
type TextRecognitionAsyncServiceServer interface {
	// To send the image for asynchronous text recognition.
	Recognize(context.Context, *RecognizeTextRequest) (*operation.Operation, error)
	// To get recognition results.
	GetRecognition(*GetRecognitionRequest, TextRecognitionAsyncService_GetRecognitionServer) error
}

// UnimplementedTextRecognitionAsyncServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTextRecognitionAsyncServiceServer struct {
}

func (UnimplementedTextRecognitionAsyncServiceServer) Recognize(context.Context, *RecognizeTextRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}
func (UnimplementedTextRecognitionAsyncServiceServer) GetRecognition(*GetRecognitionRequest, TextRecognitionAsyncService_GetRecognitionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecognition not implemented")
}

// UnsafeTextRecognitionAsyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextRecognitionAsyncServiceServer will
// result in compilation errors.
type UnsafeTextRecognitionAsyncServiceServer interface {
	mustEmbedUnimplementedTextRecognitionAsyncServiceServer()
}

func RegisterTextRecognitionAsyncServiceServer(s grpc.ServiceRegistrar, srv TextRecognitionAsyncServiceServer) {
	s.RegisterService(&TextRecognitionAsyncService_ServiceDesc, srv)
}

func _TextRecognitionAsyncService_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizeTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextRecognitionAsyncServiceServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextRecognitionAsyncService_Recognize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextRecognitionAsyncServiceServer).Recognize(ctx, req.(*RecognizeTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextRecognitionAsyncService_GetRecognition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecognitionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextRecognitionAsyncServiceServer).GetRecognition(m, &textRecognitionAsyncServiceGetRecognitionServer{stream})
}

type TextRecognitionAsyncService_GetRecognitionServer interface {
	Send(*RecognizeTextResponse) error
	grpc.ServerStream
}

type textRecognitionAsyncServiceGetRecognitionServer struct {
	grpc.ServerStream
}

func (x *textRecognitionAsyncServiceGetRecognitionServer) Send(m *RecognizeTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TextRecognitionAsyncService_ServiceDesc is the grpc.ServiceDesc for TextRecognitionAsyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextRecognitionAsyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.ocr.v1.TextRecognitionAsyncService",
	HandlerType: (*TextRecognitionAsyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _TextRecognitionAsyncService_Recognize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRecognition",
			Handler:       _TextRecognitionAsyncService_GetRecognition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yandex/cloud/ai/ocr/v1/ocr_service.proto",
}
