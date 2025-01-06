// Package Semaphore provides methods for working with Semaphore object instances.
package Semaphore

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A synchronization semaphore that can be used to synchronize multiple [Thread]s. Initialized to zero on creation. For a binary version, see [Mutex].
[b]Warning:[/b] Semaphores must be used carefully to avoid deadlocks.
[b]Warning:[/b] To guarantee that the operating system is able to perform proper cleanup (no crashes, no deadlocks), these conditions must be met:
- When a [Semaphore]'s reference count reaches zero and it is therefore destroyed, no threads must be waiting on it.
- When a [Thread]'s reference count reaches zero and it is therefore destroyed, it must not be waiting on any semaphore.
*/
type Instance [1]gdclass.Semaphore
type Any interface {
	gd.IsClass
	AsSemaphore() Instance
}

/*
Waits for the [Semaphore], if its value is zero, blocks until non-zero.
*/
func (self Instance) Wait() {
	class(self).Wait()
}

/*
Like [method wait], but won't block, so if the value is zero, fails immediately and returns [code]false[/code]. If non-zero, it returns [code]true[/code] to report success.
*/
func (self Instance) TryWait() bool {
	return bool(class(self).TryWait())
}

/*
Lowers the [Semaphore], allowing one more thread in.
*/
func (self Instance) Post() {
	class(self).Post()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Semaphore

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Semaphore"))
	casted := Instance{*(*gdclass.Semaphore)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Waits for the [Semaphore], if its value is zero, blocks until non-zero.
*/
//go:nosplit
func (self class) Wait() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Semaphore.Bind_wait, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Like [method wait], but won't block, so if the value is zero, fails immediately and returns [code]false[/code]. If non-zero, it returns [code]true[/code] to report success.
*/
//go:nosplit
func (self class) TryWait() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Semaphore.Bind_try_wait, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Lowers the [Semaphore], allowing one more thread in.
*/
//go:nosplit
func (self class) Post() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Semaphore.Bind_post, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsSemaphore() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSemaphore() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("Semaphore", func(ptr gd.Object) any { return [1]gdclass.Semaphore{*(*gdclass.Semaphore)(unsafe.Pointer(&ptr))} })
}
