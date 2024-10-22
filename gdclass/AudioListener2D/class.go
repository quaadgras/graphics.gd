package AudioListener2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Once added to the scene tree and enabled using [method make_current], this node will override the location sounds are heard from. Only one [AudioListener2D] can be current. Using [method make_current] will disable the previous [AudioListener2D].
If there is no active [AudioListener2D] in the current [Viewport], center of the screen will be used as a hearing point for the audio. [AudioListener2D] needs to be inside [SceneTree] to function.

*/
type Go [1]classdb.AudioListener2D

/*
Makes the [AudioListener2D] active, setting it as the hearing point for the sounds. If there is already another active [AudioListener2D], it will be disabled.
This method will have no effect if the [AudioListener2D] is not added to [SceneTree].
*/
func (self Go) MakeCurrent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakeCurrent()
}

/*
Disables the [AudioListener2D]. If it's not set as current, this method will have no effect.
*/
func (self Go) ClearCurrent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearCurrent()
}

/*
Returns [code]true[/code] if this [AudioListener2D] is currently active.
*/
func (self Go) IsCurrent() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCurrent())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioListener2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioListener2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Makes the [AudioListener2D] active, setting it as the hearing point for the sounds. If there is already another active [AudioListener2D], it will be disabled.
This method will have no effect if the [AudioListener2D] is not added to [SceneTree].
*/
//go:nosplit
func (self class) MakeCurrent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioListener2D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables the [AudioListener2D]. If it's not set as current, this method will have no effect.
*/
//go:nosplit
func (self class) ClearCurrent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioListener2D.Bind_clear_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if this [AudioListener2D] is currently active.
*/
//go:nosplit
func (self class) IsCurrent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioListener2D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioListener2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioListener2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("AudioListener2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
