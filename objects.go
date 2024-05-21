package gd

// NotificationType sent to an Object's notification method.
type NotificationType int32

const (
	// Notification received when the object is initialized, before its script is attached. Used internally.
	NotificationPostInitialize NotificationType = iota
	// Notification received when the object is about to be deleted. Can act as the deconstructor of some programming languages.
	NotificationPreDelete
	// Notification received when the object finishes hot reloading. This notification is only sent for extensions classes and derived.
	NotificationExtensionReloaded
)
