package Thread

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A unit of execution in a process. Can run methods on [Object]s simultaneously. The use of synchronization via [Mutex] or [Semaphore] is advised if working with shared objects.
[b]Warning:[/b]
To ensure proper cleanup without crashes or deadlocks, when a [Thread]'s reference count reaches zero and it is therefore destroyed, the following conditions must be met:
- It must not have any [Mutex] objects locked.
- It must not be waiting on any [Semaphore] objects.
- [method wait_to_finish] should have been called on it.
*/
type Instance [1]classdb.Thread

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
func (self Instance) Start(callable gd.Callable) gd.Error {
	return gd.Error(class(self).Start(callable, 1))
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
func (self Instance) GetId() string {
	return string(class(self).GetId().String())
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
func (self Instance) IsStarted() bool {
	return bool(class(self).IsStarted())
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
func (self Instance) IsAlive() bool {
	return bool(class(self).IsAlive())
}

/*
Joins the [Thread] and waits for it to finish. Returns the output of the [Callable] passed to [method start].
Should either be used when you want to retrieve the value returned from the method called by the [Thread] or before freeing the instance that contains the [Thread].
To determine if this can be called without blocking the calling thread, check if [method is_alive] is [code]false[/code].
*/
func (self Instance) WaitToFinish() gd.Variant {
	return gd.Variant(class(self).WaitToFinish())
}

/*
Sets whether the thread safety checks the engine normally performs in methods of certain classes (e.g., [Node]) should happen [b]on the current thread[/b].
The default, for every thread, is that they are enabled (as if called with [param enabled] being [code]true[/code]).
Those checks are conservative. That means that they will only succeed in considering a call thread-safe (and therefore allow it to happen) if the engine can guarantee such safety.
Because of that, there may be cases where the user may want to disable them ([param enabled] being [code]false[/code]) to make certain operations allowed again. By doing so, it becomes the user's responsibility to ensure thread safety (e.g., by using [Mutex]) for those objects that are otherwise protected by the engine.
[b]Note:[/b] This is an advanced usage of the engine. You are advised to use it only if you know what you are doing and there is no safer way.
[b]Note:[/b] This is useful for scripts running on either arbitrary [Thread] objects or tasks submitted to the [WorkerThreadPool]. It doesn't apply to code running during [Node] group processing, where the checks will be always performed.
[b]Note:[/b] Even in the case of having disabled the checks in a [WorkerThreadPool] task, there's no need to re-enable them at the end. The engine will do so.
*/
func (self Instance) SetThreadSafetyChecksEnabled(enabled bool) {
	class(self).SetThreadSafetyChecksEnabled(enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Thread

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Thread"))
	return Instance{classdb.Thread(object)}
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
//go:nosplit
func (self class) Start(callable gd.Callable, priority classdb.ThreadPriority) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callable))
	callframe.Arg(frame, priority)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
//go:nosplit
func (self class) GetId() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
//go:nosplit
func (self class) IsStarted() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_started, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
//go:nosplit
func (self class) IsAlive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_alive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Joins the [Thread] and waits for it to finish. Returns the output of the [Callable] passed to [method start].
Should either be used when you want to retrieve the value returned from the method called by the [Thread] or before freeing the instance that contains the [Thread].
To determine if this can be called without blocking the calling thread, check if [method is_alive] is [code]false[/code].
*/
//go:nosplit
func (self class) WaitToFinish() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_wait_to_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets whether the thread safety checks the engine normally performs in methods of certain classes (e.g., [Node]) should happen [b]on the current thread[/b].
The default, for every thread, is that they are enabled (as if called with [param enabled] being [code]true[/code]).
Those checks are conservative. That means that they will only succeed in considering a call thread-safe (and therefore allow it to happen) if the engine can guarantee such safety.
Because of that, there may be cases where the user may want to disable them ([param enabled] being [code]false[/code]) to make certain operations allowed again. By doing so, it becomes the user's responsibility to ensure thread safety (e.g., by using [Mutex]) for those objects that are otherwise protected by the engine.
[b]Note:[/b] This is an advanced usage of the engine. You are advised to use it only if you know what you are doing and there is no safer way.
[b]Note:[/b] This is useful for scripts running on either arbitrary [Thread] objects or tasks submitted to the [WorkerThreadPool]. It doesn't apply to code running during [Node] group processing, where the checks will be always performed.
[b]Note:[/b] Even in the case of having disabled the checks in a [WorkerThreadPool] task, there's no need to re-enable them at the end. The engine will do so.
*/
//go:nosplit
func (self class) SetThreadSafetyChecksEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_set_thread_safety_checks_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsThread() Advanced             { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsThread() Instance          { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("Thread", func(ptr gd.Object) any { return classdb.Thread(ptr) }) }

type Priority = classdb.ThreadPriority

const (
	/*A thread running with lower priority than normally.*/
	PriorityLow Priority = 0
	/*A thread with a standard priority.*/
	PriorityNormal Priority = 1
	/*A thread running with higher priority than normally.*/
	PriorityHigh Priority = 2
)
