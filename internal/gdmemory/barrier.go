package gdmemory

// barrier is the current barrier revision. Engine.Pointer values created with
// [WrapPointer] carry the barrier revision at creation time; a mismatch with
// the current value means the pointer has been invalidated.
//
// Starts at 2 so that a zero-value revision (inside a zero-value Engine.Pointer)
// is always stale.
var barrier uint64 = 2

// Barrier increments the barrier, invalidating all borrowed pointers
// created with [WrapPointer].
func Barrier() { barrier++ }
