package Thread

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A unit of execution in a process. Can run methods on [Object]s simultaneously. The use of synchronization via [Mutex] or [Semaphore] is advised if working with shared objects.
[b]Warning:[/b]
To ensure proper cleanup without crashes or deadlocks, when a [Thread]'s reference count reaches zero and it is therefore destroyed, the following conditions must be met:
- It must not have any [Mutex] objects locked.
- It must not be waiting on any [Semaphore] objects.
- [method wait_to_finish] should have been called on it.

*/
type Simple [1]classdb.Thread
func (self Simple) Start(callable gd.Callable, priority classdb.ThreadPriority) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Start(callable, priority))
}
func (self Simple) GetId() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetId(gc).String())
}
func (self Simple) IsStarted() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsStarted())
}
func (self Simple) IsAlive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAlive())
}
func (self Simple) WaitToFinish() gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).WaitToFinish(gc))
}
func (self Simple) SetThreadSafetyChecksEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetThreadSafetyChecksEnabled(gc, enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Thread
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
//go:nosplit
func (self class) Start(callable gd.Callable, priority classdb.ThreadPriority) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	callframe.Arg(frame, priority)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Thread.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
//go:nosplit
func (self class) GetId(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Thread.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
//go:nosplit
func (self class) IsStarted() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Thread.Bind_is_started, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Thread.Bind_is_alive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) WaitToFinish(ctx gd.Lifetime) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Thread.Bind_wait_to_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) SetThreadSafetyChecksEnabled(ctx gd.Lifetime, enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Thread.Bind_set_thread_safety_checks_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsThread() Expert { return self[0].AsThread() }


//go:nosplit
func (self Simple) AsThread() Simple { return self[0].AsThread() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Thread", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Priority = classdb.ThreadPriority

const (
/*A thread running with lower priority than normally.*/
	PriorityLow Priority = 0
/*A thread with a standard priority.*/
	PriorityNormal Priority = 1
/*A thread running with higher priority than normally.*/
	PriorityHigh Priority = 2
)
