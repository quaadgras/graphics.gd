// Package Thread provides methods for working with Thread object instances.
package Thread

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
A unit of execution in a process. Can run methods on [Object]s simultaneously. The use of synchronization via [Mutex] or [Semaphore] is advised if working with shared objects.
[b]Warning:[/b]
To ensure proper cleanup without crashes or deadlocks, when a [Thread]'s reference count reaches zero and it is therefore destroyed, the following conditions must be met:
- It must not have any [Mutex] objects locked.
- It must not be waiting on any [Semaphore] objects.
- [method wait_to_finish] should have been called on it.
*/
type Instance [1]gdclass.Thread

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsThread() Instance
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
func (self Instance) Start(callable func()) error { //gd:Thread.start
	return error(gd.ToError(class(self).Start(Callable.New(callable), 1)))
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
func (self Instance) GetId() string { //gd:Thread.get_id
	return string(class(self).GetId().String())
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
func (self Instance) IsStarted() bool { //gd:Thread.is_started
	return bool(class(self).IsStarted())
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
func (self Instance) IsAlive() bool { //gd:Thread.is_alive
	return bool(class(self).IsAlive())
}

/*
Joins the [Thread] and waits for it to finish. Returns the output of the [Callable] passed to [method start].
Should either be used when you want to retrieve the value returned from the method called by the [Thread] or before freeing the instance that contains the [Thread].
To determine if this can be called without blocking the calling thread, check if [method is_alive] is [code]false[/code].
*/
func (self Instance) WaitToFinish() any { //gd:Thread.wait_to_finish
	return any(class(self).WaitToFinish().Interface())
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
func SetThreadSafetyChecksEnabled(enabled bool) { //gd:Thread.set_thread_safety_checks_enabled
	self := Instance{}
	class(self).SetThreadSafetyChecksEnabled(enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Thread

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Thread"))
	casted := Instance{*(*gdclass.Thread)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
//go:nosplit
func (self class) Start(callable Callable.Function, priority gdclass.ThreadPriority) gd.Error { //gd:Thread.start
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	callframe.Arg(frame, priority)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
//go:nosplit
func (self class) GetId() gd.String { //gd:Thread.get_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
//go:nosplit
func (self class) IsStarted() bool { //gd:Thread.is_started
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_started, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
//go:nosplit
func (self class) IsAlive() bool { //gd:Thread.is_alive
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_alive, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) WaitToFinish() gd.Variant { //gd:Thread.wait_to_finish
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_wait_to_finish, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetThreadSafetyChecksEnabled(enabled bool) { //gd:Thread.set_thread_safety_checks_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_set_thread_safety_checks_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsThread() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsThread() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Thread", func(ptr gd.Object) any { return [1]gdclass.Thread{*(*gdclass.Thread)(unsafe.Pointer(&ptr))} })
}

type Priority = gdclass.ThreadPriority //gd:Thread.Priority

const (
	/*A thread running with lower priority than normally.*/
	PriorityLow Priority = 0
	/*A thread with a standard priority.*/
	PriorityNormal Priority = 1
	/*A thread running with higher priority than normally.*/
	PriorityHigh Priority = 2
)

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
