package Timer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [Timer] node is a countdown timer and is the simplest way to handle time-based logic in the engine. When a timer reaches the end of its [member wait_time], it will emit the [signal timeout] signal.
After a timer enters the tree, it can be manually started with [method start]. A timer node is also started automatically if [member autostart] is [code]true[/code].
Without requiring much code, a timer node can be added and configured in the editor. The [signal timeout] signal it emits can also be connected through the Node dock in the editor:
[codeblock]
func _on_timer_timeout():

	print("Time to attack!")

[/codeblock]
[b]Note:[/b] To create a one-shot timer without instantiating a node, use [method SceneTree.create_timer].
[b]Note:[/b] Timers are affected by [member Engine.time_scale]. The higher the time scale, the sooner timers will end. How often a timer processes may depend on the framerate or [member Engine.physics_ticks_per_second].
*/
type Instance [1]classdb.Timer
type Any interface {
	gd.IsClass
	AsTimer() Instance
}

/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
func (self Instance) Start() {
	class(self).Start(gd.Float(-1))
}

/*
Stops the timer.
*/
func (self Instance) Stop() {
	class(self).Stop()
}

/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
func (self Instance) IsStopped() bool {
	return bool(class(self).IsStopped())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Timer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Timer"))
	return Instance{*(*classdb.Timer)(unsafe.Pointer(&object))}
}

func (self Instance) ProcessCallback() classdb.TimerTimerProcessCallback {
	return classdb.TimerTimerProcessCallback(class(self).GetTimerProcessCallback())
}

func (self Instance) SetProcessCallback(value classdb.TimerTimerProcessCallback) {
	class(self).SetTimerProcessCallback(value)
}

func (self Instance) WaitTime() Float.X {
	return Float.X(Float.X(class(self).GetWaitTime()))
}

func (self Instance) SetWaitTime(value Float.X) {
	class(self).SetWaitTime(gd.Float(value))
}

func (self Instance) OneShot() bool {
	return bool(class(self).IsOneShot())
}

func (self Instance) SetOneShot(value bool) {
	class(self).SetOneShot(value)
}

func (self Instance) Autostart() bool {
	return bool(class(self).HasAutostart())
}

func (self Instance) SetAutostart(value bool) {
	class(self).SetAutostart(value)
}

func (self Instance) Paused() bool {
	return bool(class(self).IsPaused())
}

func (self Instance) SetPaused(value bool) {
	class(self).SetPaused(value)
}

func (self Instance) TimeLeft() Float.X {
	return Float.X(Float.X(class(self).GetTimeLeft()))
}

//go:nosplit
func (self class) SetWaitTime(time_sec gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_wait_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWaitTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_wait_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOneShot(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsOneShot() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutostart(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_autostart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutostart() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_has_autostart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
//go:nosplit
func (self class) Start(time_sec gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the timer.
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPaused(paused bool) {
	var frame = callframe.New()
	callframe.Arg(frame, paused)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPaused() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
//go:nosplit
func (self class) IsStopped() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_stopped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTimeLeft() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_time_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimerProcessCallback(callback classdb.TimerTimerProcessCallback) {
	var frame = callframe.New()
	callframe.Arg(frame, callback)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimerProcessCallback() classdb.TimerTimerProcessCallback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TimerTimerProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTimeout(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("timeout"), gd.NewCallable(cb), 0)
}

func (self class) AsTimer() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTimer() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	classdb.Register("Timer", func(ptr gd.Object) any { return [1]classdb.Timer{*(*classdb.Timer)(unsafe.Pointer(&ptr))} })
}

type TimerProcessCallback = classdb.TimerTimerProcessCallback

const (
	/*Update the timer every physics process frame (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	TimerProcessPhysics TimerProcessCallback = 0
	/*Update the timer every process (rendered) frame (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	TimerProcessIdle TimerProcessCallback = 1
)
