package di

// Container holds all dependency containers for the Notification service.
type Container struct {
	NotificationContainer NotificationContainer
	PreferenceContainer   PreferenceContainer
	WebhookContainer      WebhookContainer
	SSEContainer          SSEContainer
}

// GlobalContainer is the global instance of Container.
var GlobalContainer *Container
