package Semaphore

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A synchronization semaphore that can be used to synchronize multiple [Thread]s. Initialized to zero on creation. For a binary version, see [Mutex].
[b]Warning:[/b] Semaphores must be used carefully to avoid deadlocks.
[b]Warning:[/b] To guarantee that the operating system is able to perform proper cleanup (no crashes, no deadlocks), these conditions must be met:
- When a [Semaphore]'s reference count reaches zero and it is therefore destroyed, no threads must be waiting on it.
- When a [Thread]'s reference count reaches zero and it is therefore destroyed, it must not be waiting on any semaphore.

*/
type Go [1]classdb.Semaphore

/*
Waits for the [Semaphore], if its value is zero, blocks until non-zero.
*/
func (self Go) Wait() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Wait()
}

/*
Like [method wait], but won't block, so if the value is zero, fails immediately and returns [code]false[/code]. If non-zero, it returns [code]true[/code] to report success.
*/
func (self Go) TryWait() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TryWait())
}

/*
Lowers the [Semaphore], allowing one more thread in.
*/
func (self Go) Post() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Post()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Semaphore
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Semaphore"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Waits for the [Semaphore], if its value is zero, blocks until non-zero.
*/
//go:nosplit
func (self class) Wait()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Semaphore.Bind_wait, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Like [method wait], but won't block, so if the value is zero, fails immediately and returns [code]false[/code]. If non-zero, it returns [code]true[/code] to report success.
*/
//go:nosplit
func (self class) TryWait() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Semaphore.Bind_try_wait, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Lowers the [Semaphore], allowing one more thread in.
*/
//go:nosplit
func (self class) Post()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Semaphore.Bind_post, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsSemaphore() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSemaphore() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("Semaphore", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
