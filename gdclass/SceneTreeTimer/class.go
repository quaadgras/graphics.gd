package SceneTreeTimer

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
A one-shot timer managed by the scene tree, which emits [signal timeout] on completion. See also [method SceneTree.create_timer].
As opposed to [Timer], it does not require the instantiation of a node. Commonly used to create a one-shot delay timer as in the following example:
[codeblocks]
[gdscript]
func some_function():
    print("Timer started.")
    await get_tree().create_timer(1.0).timeout
    print("Timer ended.")
[/gdscript]
[csharp]
public async Task SomeFunction()
{
    GD.Print("Timer started.");
    await ToSignal(GetTree().CreateTimer(1.0f), SceneTreeTimer.SignalName.Timeout);
    GD.Print("Timer ended.");
}
[/csharp]
[/codeblocks]
The timer will be dereferenced after its time elapses. To preserve the timer, you can keep a reference to it. See [RefCounted].
[b]Note:[/b] The timer is processed after all of the nodes in the current frame, i.e. node's [method Node._process] method would be called before the timer (or [method Node._physics_process] if [code]process_in_physics[/code] in [method SceneTree.create_timer] has been set to [code]true[/code]).

*/
type Go [1]classdb.SceneTreeTimer
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SceneTreeTimer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SceneTreeTimer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TimeLeft() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTimeLeft()))
}

func (self Go) SetTimeLeft(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimeLeft(gd.Float(value))
}

//go:nosplit
func (self class) SetTimeLeft(time gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTreeTimer.Bind_set_time_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimeLeft() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTreeTimer.Bind_get_time_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTimeout(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("timeout"), gc.Callable(cb), 0)
}


func (self class) AsSceneTreeTimer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSceneTreeTimer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SceneTreeTimer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
