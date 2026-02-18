package Node

import gd "graphics.gd/internal"

// IsQueuedForDeletion returns true if the [Instance.QueueFree] method was called for the object.
func (self Instance) IsQueuedForDeletion() bool {
	return bool(gd.ObjectIsQueuedForDeletion(self.AsObject()[0]))
}
