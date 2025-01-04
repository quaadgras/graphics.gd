// Package Mutex provides methods for working with Mutex object instances.
package Mutex

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A synchronization mutex (mutual exclusion). This is used to synchronize multiple [Thread]s, and is equivalent to a binary [Semaphore]. It guarantees that only one thread can access a critical section at a time.
This is a reentrant mutex, meaning that it can be locked multiple times by one thread, provided it also unlocks it as many times.
[b]Warning:[/b] Mutexes must be used carefully to avoid deadlocks.
[b]Warning:[/b] To ensure proper cleanup without crashes or deadlocks, the following conditions must be met:
- When a [Mutex]'s reference count reaches zero and it is therefore destroyed, no threads (including the one on which the destruction will happen) must have it locked.
- When a [Thread]'s reference count reaches zero and it is therefore destroyed, it must not have any mutex locked.
*/
type Instance [1]gdclass.Mutex
type Any interface {
	gd.IsClass
	AsMutex() Instance
}

/*
Locks this [Mutex], blocks until it is unlocked by the current owner.
[b]Note:[/b] This function returns without blocking if the thread already has ownership of the mutex.
*/
func (self Instance) Lock() {
	class(self).Lock()
}

/*
Tries locking this [Mutex], but does not block. Returns [code]true[/code] on success, [code]false[/code] otherwise.
[b]Note:[/b] This function returns [code]true[/code] if the thread already has ownership of the mutex.
*/
func (self Instance) TryLock() bool {
	return bool(class(self).TryLock())
}

/*
Unlocks this [Mutex], leaving it to other threads.
[b]Note:[/b] If a thread called [method lock] or [method try_lock] multiple times while already having ownership of the mutex, it must also call [method unlock] the same number of times in order to unlock it correctly.
[b]Warning:[/b] Calling [method unlock] more times that [method lock] on a given thread, thus ending up trying to unlock a non-locked mutex, is wrong and may causes crashes or deadlocks.
*/
func (self Instance) Unlock() {
	class(self).Unlock()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Mutex

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Mutex"))
	return Instance{*(*gdclass.Mutex)(unsafe.Pointer(&object))}
}

/*
Locks this [Mutex], blocks until it is unlocked by the current owner.
[b]Note:[/b] This function returns without blocking if the thread already has ownership of the mutex.
*/
//go:nosplit
func (self class) Lock() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tries locking this [Mutex], but does not block. Returns [code]true[/code] on success, [code]false[/code] otherwise.
[b]Note:[/b] This function returns [code]true[/code] if the thread already has ownership of the mutex.
*/
//go:nosplit
func (self class) TryLock() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_try_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unlocks this [Mutex], leaving it to other threads.
[b]Note:[/b] If a thread called [method lock] or [method try_lock] multiple times while already having ownership of the mutex, it must also call [method unlock] the same number of times in order to unlock it correctly.
[b]Warning:[/b] Calling [method unlock] more times that [method lock] on a given thread, thus ending up trying to unlock a non-locked mutex, is wrong and may causes crashes or deadlocks.
*/
//go:nosplit
func (self class) Unlock() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_unlock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsMutex() Advanced              { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMutex() Instance           { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	gdclass.Register("Mutex", func(ptr gd.Object) any { return [1]gdclass.Mutex{*(*gdclass.Mutex)(unsafe.Pointer(&ptr))} })
}
