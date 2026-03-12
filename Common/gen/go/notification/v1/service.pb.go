// Code generated manually (protoc not available). DO NOT EDIT.
// source: notification/v1/service.proto

package notificationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Recipient holds target user information for a notification.
type Recipient struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Recipient) Reset()         { *x = Recipient{} }
func (x *Recipient) String() string { return protoimpl.X.MessageStringOf(x) }
func (*Recipient) ProtoMessage()    {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Recipient) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Recipient) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Recipient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Recipient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SendNotificationRequest is the gRPC request for sending a notification.
type SendNotificationRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	IdempotencyKey string                 `protobuf:"bytes,1,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	Type           string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Channel        string                 `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Priority       string                 `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Recipient      *Recipient             `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TemplateData   map[string]string      `protobuf:"bytes,6,rep,name=template_data,json=templateData,proto3" json:"template_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SendNotificationRequest) Reset()         { *x = SendNotificationRequest{} }
func (x *SendNotificationRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*SendNotificationRequest) ProtoMessage()    {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendNotificationRequest) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (x *SendNotificationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendNotificationRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SendNotificationRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *SendNotificationRequest) GetRecipient() *Recipient {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *SendNotificationRequest) GetTemplateData() map[string]string {
	if x != nil {
		return x.TemplateData
	}
	return nil
}

// SendNotificationResponse is the gRPC response after sending a notification.
type SendNotificationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendNotificationResponse) Reset()         { *x = SendNotificationResponse{} }
func (x *SendNotificationResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*SendNotificationResponse) ProtoMessage()    {}

func (x *SendNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*SendNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *SendNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetUnreadCountRequest requests the unread notification count for a user.
type GetUnreadCountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUnreadCountRequest) Reset()         { *x = GetUnreadCountRequest{} }
func (x *GetUnreadCountRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetUnreadCountRequest) ProtoMessage()    {}

func (x *GetUnreadCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetUnreadCountRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUnreadCountRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// GetUnreadCountResponse returns the unread notification count.
type GetUnreadCountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUnreadCountResponse) Reset()         { *x = GetUnreadCountResponse{} }
func (x *GetUnreadCountResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetUnreadCountResponse) ProtoMessage()    {}

func (x *GetUnreadCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetUnreadCountResponse) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUnreadCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// GetNotificationsRequest requests a paginated list of notifications for a user.
type GetNotificationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsRequest) Reset()         { *x = GetNotificationsRequest{} }
func (x *GetNotificationsRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetNotificationsRequest) ProtoMessage()    {}

func (x *GetNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetNotificationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetNotificationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// NotificationItem represents a single notification in a list response.
type NotificationItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Channel       string                 `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Subject       string                 `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Body          string                 `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	IsRead        bool                   `protobuf:"varint,7,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationItem) Reset()         { *x = NotificationItem{} }
func (x *NotificationItem) String() string { return protoimpl.X.MessageStringOf(x) }
func (*NotificationItem) ProtoMessage()    {}

func (x *NotificationItem) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*NotificationItem) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *NotificationItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotificationItem) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *NotificationItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NotificationItem) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *NotificationItem) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *NotificationItem) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *NotificationItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// GetNotificationsResponse returns paginated notifications.
type GetNotificationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notifications []*NotificationItem    `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsResponse) Reset()         { *x = GetNotificationsResponse{} }
func (x *GetNotificationsResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetNotificationsResponse) ProtoMessage()    {}

func (x *GetNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notification_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetNotificationsResponse) GetNotifications() []*NotificationItem {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *GetNotificationsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_notification_v1_service_proto protoreflect.FileDescriptor

// Minimal raw descriptor encoding for the notification service proto.
const file_notification_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x1enotification/v1/service.proto\x12\x0fnotification.v1" +
	"\"D\n" +
	"\tRecipient\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name" +
	"B5Z3go-link/common/gen/notification/v1;notificationv1b\x06proto3"

var (
	file_notification_v1_service_proto_rawDescOnce sync.Once
	file_notification_v1_service_proto_rawDescData []byte
)

func file_notification_v1_service_proto_rawDescGZIP() []byte {
	file_notification_v1_service_proto_rawDescOnce.Do(func() {
		file_notification_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notification_v1_service_proto_rawDesc), len(file_notification_v1_service_proto_rawDesc)))
	})
	return file_notification_v1_service_proto_rawDescData
}

var file_notification_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_notification_v1_service_proto_goTypes = []any{
	(*Recipient)(nil),                // 0: notification.v1.Recipient
	(*SendNotificationRequest)(nil),  // 1: notification.v1.SendNotificationRequest
	(*SendNotificationResponse)(nil), // 2: notification.v1.SendNotificationResponse
	(*GetUnreadCountRequest)(nil),    // 3: notification.v1.GetUnreadCountRequest
	(*GetUnreadCountResponse)(nil),   // 4: notification.v1.GetUnreadCountResponse
	(*GetNotificationsRequest)(nil),  // 5: notification.v1.GetNotificationsRequest
	(*NotificationItem)(nil),         // 6: notification.v1.NotificationItem
	(*GetNotificationsResponse)(nil), // 7: notification.v1.GetNotificationsResponse
}

var file_notification_v1_service_proto_depIdxs = []int32{
	0, // 0: notification.v1.SendNotificationRequest.recipient:type_name -> notification.v1.Recipient
	6, // 1: notification.v1.GetNotificationsResponse.notifications:type_name -> notification.v1.NotificationItem
	1, // 2: notification.v1.NotificationService.SendNotification:input_type -> notification.v1.SendNotificationRequest
	3, // 3: notification.v1.NotificationService.GetUnreadCount:input_type -> notification.v1.GetUnreadCountRequest
	5, // 4: notification.v1.NotificationService.GetNotifications:input_type -> notification.v1.GetNotificationsRequest
	2, // 5: notification.v1.NotificationService.SendNotification:output_type -> notification.v1.SendNotificationResponse
	4, // 6: notification.v1.NotificationService.GetUnreadCount:output_type -> notification.v1.GetUnreadCountResponse
	7, // 7: notification.v1.NotificationService.GetNotifications:output_type -> notification.v1.GetNotificationsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_v1_service_proto_init() }

func file_notification_v1_service_proto_init() {
	if File_notification_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notification_v1_service_proto_rawDesc), len(file_notification_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_v1_service_proto_goTypes,
		DependencyIndexes: file_notification_v1_service_proto_depIdxs,
		MessageInfos:      file_notification_v1_service_proto_msgTypes,
	}.Build()
	File_notification_v1_service_proto = out.File
	file_notification_v1_service_proto_goTypes = nil
	file_notification_v1_service_proto_depIdxs = nil
}
