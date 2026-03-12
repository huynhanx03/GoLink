package grpc

import (
	notificationv1 "go-link/common/gen/go/notification/v1"
	"go-link/notification/internal/ports"

	"google.golang.org/grpc"
)

// V1Routes returns a function that registers the NotificationService gRPC routes.
func V1Routes(
	notificationService ports.NotificationService,
	notificationRepo ports.NotificationRepository,
	unreadCounter ports.UnreadCounter,
) func(srv *grpc.Server) {
	return func(srv *grpc.Server) {
		notificationv1.RegisterNotificationServiceServer(srv, NewNotificationServer(notificationService, notificationRepo, unreadCounter))
	}
}
