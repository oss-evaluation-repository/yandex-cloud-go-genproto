// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/disk_service.proto

package compute

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
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
	DiskService_Get_FullMethodName                   = "/yandex.cloud.compute.v1.DiskService/Get"
	DiskService_List_FullMethodName                  = "/yandex.cloud.compute.v1.DiskService/List"
	DiskService_Create_FullMethodName                = "/yandex.cloud.compute.v1.DiskService/Create"
	DiskService_Update_FullMethodName                = "/yandex.cloud.compute.v1.DiskService/Update"
	DiskService_Delete_FullMethodName                = "/yandex.cloud.compute.v1.DiskService/Delete"
	DiskService_ListOperations_FullMethodName        = "/yandex.cloud.compute.v1.DiskService/ListOperations"
	DiskService_Move_FullMethodName                  = "/yandex.cloud.compute.v1.DiskService/Move"
	DiskService_Relocate_FullMethodName              = "/yandex.cloud.compute.v1.DiskService/Relocate"
	DiskService_ListSnapshotSchedules_FullMethodName = "/yandex.cloud.compute.v1.DiskService/ListSnapshotSchedules"
	DiskService_ListAccessBindings_FullMethodName    = "/yandex.cloud.compute.v1.DiskService/ListAccessBindings"
	DiskService_SetAccessBindings_FullMethodName     = "/yandex.cloud.compute.v1.DiskService/SetAccessBindings"
	DiskService_UpdateAccessBindings_FullMethodName  = "/yandex.cloud.compute.v1.DiskService/UpdateAccessBindings"
)

// DiskServiceClient is the client API for DiskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiskServiceClient interface {
	// Returns the specified Disk resource.
	//
	// To get the list of available Disk resources, make a [List] request.
	Get(ctx context.Context, in *GetDiskRequest, opts ...grpc.CallOption) (*Disk, error)
	// Retrieves the list of Disk resources in the specified folder.
	List(ctx context.Context, in *ListDisksRequest, opts ...grpc.CallOption) (*ListDisksResponse, error)
	// Creates a disk in the specified folder.
	//
	// You can create an empty disk or restore it from a snapshot or an image.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(ctx context.Context, in *CreateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified disk.
	Update(ctx context.Context, in *UpdateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified disk.
	//
	// Deleting a disk removes its data permanently and is irreversible. However, deleting a disk does not delete
	// any snapshots or images previously made from the disk. You must delete snapshots and images separately.
	//
	// It is not possible to delete a disk that is attached to an instance.
	Delete(ctx context.Context, in *DeleteDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified disk.
	ListOperations(ctx context.Context, in *ListDiskOperationsRequest, opts ...grpc.CallOption) (*ListDiskOperationsResponse, error)
	// Moves the specified disk to another folder of the same cloud.
	Move(ctx context.Context, in *MoveDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Moves the specified disk to another availability zone
	//
	// Disk must be detached from instances. To move attached
	// disk use [InstanceService.Relocate] request.
	Relocate(ctx context.Context, in *RelocateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of snapshot schedules the specified disk is attached to.
	ListSnapshotSchedules(ctx context.Context, in *ListDiskSnapshotSchedulesRequest, opts ...grpc.CallOption) (*ListDiskSnapshotSchedulesResponse, error)
	// Lists access bindings for the disk.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the disk.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the disk.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type diskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiskServiceClient(cc grpc.ClientConnInterface) DiskServiceClient {
	return &diskServiceClient{cc}
}

func (c *diskServiceClient) Get(ctx context.Context, in *GetDiskRequest, opts ...grpc.CallOption) (*Disk, error) {
	out := new(Disk)
	err := c.cc.Invoke(ctx, DiskService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) List(ctx context.Context, in *ListDisksRequest, opts ...grpc.CallOption) (*ListDisksResponse, error) {
	out := new(ListDisksResponse)
	err := c.cc.Invoke(ctx, DiskService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) Create(ctx context.Context, in *CreateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) Update(ctx context.Context, in *UpdateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) Delete(ctx context.Context, in *DeleteDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) ListOperations(ctx context.Context, in *ListDiskOperationsRequest, opts ...grpc.CallOption) (*ListDiskOperationsResponse, error) {
	out := new(ListDiskOperationsResponse)
	err := c.cc.Invoke(ctx, DiskService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) Move(ctx context.Context, in *MoveDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_Move_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) Relocate(ctx context.Context, in *RelocateDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_Relocate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) ListSnapshotSchedules(ctx context.Context, in *ListDiskSnapshotSchedulesRequest, opts ...grpc.CallOption) (*ListDiskSnapshotSchedulesResponse, error) {
	out := new(ListDiskSnapshotSchedulesResponse)
	err := c.cc.Invoke(ctx, DiskService_ListSnapshotSchedules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, DiskService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DiskService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskServiceServer is the server API for DiskService service.
// All implementations should embed UnimplementedDiskServiceServer
// for forward compatibility
type DiskServiceServer interface {
	// Returns the specified Disk resource.
	//
	// To get the list of available Disk resources, make a [List] request.
	Get(context.Context, *GetDiskRequest) (*Disk, error)
	// Retrieves the list of Disk resources in the specified folder.
	List(context.Context, *ListDisksRequest) (*ListDisksResponse, error)
	// Creates a disk in the specified folder.
	//
	// You can create an empty disk or restore it from a snapshot or an image.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(context.Context, *CreateDiskRequest) (*operation.Operation, error)
	// Updates the specified disk.
	Update(context.Context, *UpdateDiskRequest) (*operation.Operation, error)
	// Deletes the specified disk.
	//
	// Deleting a disk removes its data permanently and is irreversible. However, deleting a disk does not delete
	// any snapshots or images previously made from the disk. You must delete snapshots and images separately.
	//
	// It is not possible to delete a disk that is attached to an instance.
	Delete(context.Context, *DeleteDiskRequest) (*operation.Operation, error)
	// Lists operations for the specified disk.
	ListOperations(context.Context, *ListDiskOperationsRequest) (*ListDiskOperationsResponse, error)
	// Moves the specified disk to another folder of the same cloud.
	Move(context.Context, *MoveDiskRequest) (*operation.Operation, error)
	// Moves the specified disk to another availability zone
	//
	// Disk must be detached from instances. To move attached
	// disk use [InstanceService.Relocate] request.
	Relocate(context.Context, *RelocateDiskRequest) (*operation.Operation, error)
	// Retrieves the list of snapshot schedules the specified disk is attached to.
	ListSnapshotSchedules(context.Context, *ListDiskSnapshotSchedulesRequest) (*ListDiskSnapshotSchedulesResponse, error)
	// Lists access bindings for the disk.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the disk.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the disk.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedDiskServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiskServiceServer struct {
}

func (UnimplementedDiskServiceServer) Get(context.Context, *GetDiskRequest) (*Disk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDiskServiceServer) List(context.Context, *ListDisksRequest) (*ListDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDiskServiceServer) Create(context.Context, *CreateDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDiskServiceServer) Update(context.Context, *UpdateDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDiskServiceServer) Delete(context.Context, *DeleteDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDiskServiceServer) ListOperations(context.Context, *ListDiskOperationsRequest) (*ListDiskOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedDiskServiceServer) Move(context.Context, *MoveDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedDiskServiceServer) Relocate(context.Context, *RelocateDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relocate not implemented")
}
func (UnimplementedDiskServiceServer) ListSnapshotSchedules(context.Context, *ListDiskSnapshotSchedulesRequest) (*ListDiskSnapshotSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshotSchedules not implemented")
}
func (UnimplementedDiskServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedDiskServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedDiskServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeDiskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiskServiceServer will
// result in compilation errors.
type UnsafeDiskServiceServer interface {
	mustEmbedUnimplementedDiskServiceServer()
}

func RegisterDiskServiceServer(s grpc.ServiceRegistrar, srv DiskServiceServer) {
	s.RegisterService(&DiskService_ServiceDesc, srv)
}

func _DiskService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Get(ctx, req.(*GetDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).List(ctx, req.(*ListDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Create(ctx, req.(*CreateDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Update(ctx, req.(*UpdateDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Delete(ctx, req.(*DeleteDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).ListOperations(ctx, req.(*ListDiskOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Move(ctx, req.(*MoveDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_Relocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelocateDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).Relocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_Relocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).Relocate(ctx, req.(*RelocateDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_ListSnapshotSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskSnapshotSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).ListSnapshotSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_ListSnapshotSchedules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).ListSnapshotSchedules(ctx, req.(*ListDiskSnapshotSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiskService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiskService_ServiceDesc is the grpc.ServiceDesc for DiskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.DiskService",
	HandlerType: (*DiskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DiskService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiskService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DiskService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DiskService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DiskService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _DiskService_ListOperations_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _DiskService_Move_Handler,
		},
		{
			MethodName: "Relocate",
			Handler:    _DiskService_Relocate_Handler,
		},
		{
			MethodName: "ListSnapshotSchedules",
			Handler:    _DiskService_ListSnapshotSchedules_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _DiskService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _DiskService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _DiskService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/disk_service.proto",
}
