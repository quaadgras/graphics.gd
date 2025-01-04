package Node

// IsQueuedForDeletion returns true if the [Instance.QueueFree] method was called for the object.
func (self Instance) IsQueuedForDeletion() bool {
	return bool(self[0].AsObject().IsQueuedForDeletion())
}
