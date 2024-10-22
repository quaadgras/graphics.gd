package Timer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.Timer

/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
func (self Go) Start() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Start(gd.Float(-1))
}

/*
Stops the timer.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}

/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
func (self Go) IsStopped() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsStopped())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Timer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Timer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ProcessCallback() classdb.TimerTimerProcessCallback {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TimerTimerProcessCallback(class(self).GetTimerProcessCallback())
}

func (self Go) SetProcessCallback(value classdb.TimerTimerProcessCallback) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimerProcessCallback(value)
}

func (self Go) WaitTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWaitTime()))
}

func (self Go) SetWaitTime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWaitTime(gd.Float(value))
}

func (self Go) OneShot() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsOneShot())
}

func (self Go) SetOneShot(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOneShot(value)
}

func (self Go) Autostart() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasAutostart())
}

func (self Go) SetAutostart(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutostart(value)
}

func (self Go) Paused() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPaused())
}

func (self Go) SetPaused(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPaused(value)
}

func (self Go) TimeLeft() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTimeLeft()))
}

//go:nosplit
func (self class) SetWaitTime(time_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_set_wait_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWaitTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_get_wait_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOneShot(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOneShot() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_is_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutostart(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_set_autostart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasAutostart() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_has_autostart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts the timer, if it was not started already. Fails if the timer is not inside the tree. If [param time_sec] is greater than [code]0[/code], this value is used for the [member wait_time].
[b]Note:[/b] This method does not resume a paused timer. See [member paused].
*/
//go:nosplit
func (self class) Start(time_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the timer.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPaused(paused bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paused)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_set_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the timer is stopped or has not started.
*/
//go:nosplit
func (self class) IsStopped() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_is_stopped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTimeLeft() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_get_time_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimerProcessCallback(callback classdb.TimerTimerProcessCallback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, callback)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_set_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimerProcessCallback() classdb.TimerTimerProcessCallback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TimerTimerProcessCallback](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Timer.Bind_get_timer_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTimeout(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("timeout"), gc.Callable(cb), 0)
}


func (self class) AsTimer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTimer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("Timer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TimerProcessCallback = classdb.TimerTimerProcessCallback

const (
/*Update the timer every physics process frame (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	TimerProcessPhysics TimerProcessCallback = 0
/*Update the timer every process (rendered) frame (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	TimerProcessIdle TimerProcessCallback = 1
)
