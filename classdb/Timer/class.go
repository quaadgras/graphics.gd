// Package Timer provides methods for working with Timer object instances.
package Timer

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
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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
var _ String.Readable

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
type Instance [1]gdclass.Timer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTimer() Instance
}

/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
func (self Instance) Start() { //gd:Timer.start
	class(self).Start(gd.Float(-1))
}

/*
Stops the timer.
*/
func (self Instance) Stop() { //gd:Timer.stop
	class(self).Stop()
}

/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
func (self Instance) IsStopped() bool { //gd:Timer.is_stopped
	return bool(class(self).IsStopped())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Timer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Timer"))
	casted := Instance{*(*gdclass.Timer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) ProcessCallback() gdclass.TimerTimerProcessCallback {
	return gdclass.TimerTimerProcessCallback(class(self).GetTimerProcessCallback())
}

func (self Instance) SetProcessCallback(value gdclass.TimerTimerProcessCallback) {
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
func (self class) SetWaitTime(time_sec gd.Float) { //gd:Timer.set_wait_time
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_wait_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWaitTime() gd.Float { //gd:Timer.get_wait_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_wait_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOneShot(enable bool) { //gd:Timer.set_one_shot
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOneShot() bool { //gd:Timer.is_one_shot
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutostart(enable bool) { //gd:Timer.set_autostart
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_autostart, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutostart() bool { //gd:Timer.has_autostart
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_has_autostart, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
//go:nosplit
func (self class) Start(time_sec gd.Float) { //gd:Timer.start
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the timer.
*/
//go:nosplit
func (self class) Stop() { //gd:Timer.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPaused(paused bool) { //gd:Timer.set_paused
	var frame = callframe.New()
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPaused() bool { //gd:Timer.is_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
//go:nosplit
func (self class) IsStopped() bool { //gd:Timer.is_stopped
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_is_stopped, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTimeLeft() gd.Float { //gd:Timer.get_time_left
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_time_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimerProcessCallback(callback gdclass.TimerTimerProcessCallback) { //gd:Timer.set_timer_process_callback
	var frame = callframe.New()
	callframe.Arg(frame, callback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_set_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimerProcessCallback() gdclass.TimerTimerProcessCallback { //gd:Timer.get_timer_process_callback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TimerTimerProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Timer.Bind_get_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTimeout(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("timeout"), gd.NewCallable(cb), 0)
}

func (self class) AsTimer() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTimer() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("Timer", func(ptr gd.Object) any { return [1]gdclass.Timer{*(*gdclass.Timer)(unsafe.Pointer(&ptr))} })
}

type TimerProcessCallback = gdclass.TimerTimerProcessCallback //gd:Timer.TimerProcessCallback

const (
	/*Update the timer every physics process frame (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	TimerProcessPhysics TimerProcessCallback = 0
	/*Update the timer every process (rendered) frame (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	TimerProcessIdle TimerProcessCallback = 1
)
