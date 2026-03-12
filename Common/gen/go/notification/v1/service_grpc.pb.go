// Code generated manually (protoc not available). DO NOT EDIT.
// source: notification/v1/service.proto
// go_package: go-link/common/gen/go/notification/v1;notificationv1

package notificationv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion9

const (
	NotificationService_SendNotification_FullMethodName  = "/notification.v1.NotificationService/SendNotification"
	NotificationService_GetUnreadCount_FullMethodName    = "/notification.v1.NotificationService/GetUnreadCount"
	NotificationService_GetNotifications_FullMethodName  = "/notification.v1.NotificationService/GetNotifications"
)

// NotificationServiceClient is the client API for NotificationService.
type NotificationServiceClient interface {
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error)
	GetUnreadCount(ctx context.Context, in *GetUnreadCountRequest, opts ...grpc.CallOption) (*GetUnreadCountResponse, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_SendNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetUnreadCount(ctx context.Context, in *GetUnreadCountRequest, opts ...grpc.CallOption) (*GetUnreadCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUnreadCountResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetUnreadCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer for forward compatibility.
type NotificationServiceServer interface {
	SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error)
	GetUnreadCount(context.Context, *GetUnreadCountRequest) (*GetUnreadCountResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SendNotification not implemented")
}

func (UnimplementedNotificationServiceServer) GetUnreadCount(context.Context, *GetUnreadCountRequest) (*GetUnreadCountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUnreadCount not implemented")
}

func (UnimplementedNotificationServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetNotifications not implemented")
}

func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}
func (UnimplementedNotificationServiceServer) testEmbeddedByValue()                             {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

// RegisterNotificationServiceServer registers the server implementation with the gRPC server.
func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_SendNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetUnreadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnreadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetUnreadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetUnreadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetUnreadCount(ctx, req.(*GetUnreadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotification",
			Handler:    _NotificationService_SendNotification_Handler,
		},
		{
			MethodName: "GetUnreadCount",
			Handler:    _NotificationService_GetUnreadCount_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _NotificationService_GetNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/v1/service.proto",
}
