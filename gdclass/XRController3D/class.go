package XRController3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRNode3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This is a helper spatial node that is linked to the tracking of controllers. It also offers several handy passthroughs to the state of buttons and such on the controllers.
Controllers are linked by their ID. You can create controller nodes before the controllers are available. If your game always uses two controllers (one for each hand), you can predefine the controllers with ID 1 and 2; they will become active as soon as the controllers are identified. If you expect additional controllers to be used, you should react to the signals and add XRController3D nodes to your scene.
The position of the controller node is automatically updated by the [XRServer]. This makes this node ideal to add child nodes to visualize the controller.
As many XR runtimes now use a configurable action map all inputs are named.

*/
type Go [1]classdb.XRController3D

/*
Returns [code]true[/code] if the button with the given [param name] is pressed.
*/
func (self Go) IsButtonPressed(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsButtonPressed(gc.StringName(name)))
}

/*
Returns a [Variant] for the input with the given [param name]. This works for any input type, the variant will be typed according to the actions configuration.
*/
func (self Go) GetInput(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetInput(gc, gc.StringName(name)))
}

/*
Returns a numeric value for the input with the given [param name]. This is used for triggers and grip sensors.
*/
func (self Go) GetFloat(name string) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetFloat(gc.StringName(name))))
}

/*
Returns a [Vector2] for the input with the given [param name]. This is used for thumbsticks and thumbpads found on many controllers.
*/
func (self Go) GetVector2(name string) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetVector2(gc.StringName(name)))
}

/*
Returns the hand holding this controller, if known. See [enum XRPositionalTracker.TrackerHand].
*/
func (self Go) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRPositionalTrackerTrackerHand(class(self).GetTrackerHand())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRController3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRController3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns [code]true[/code] if the button with the given [param name] is pressed.
*/
//go:nosplit
func (self class) IsButtonPressed(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRController3D.Bind_is_button_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Variant] for the input with the given [param name]. This works for any input type, the variant will be typed according to the actions configuration.
*/
//go:nosplit
func (self class) GetInput(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRController3D.Bind_get_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a numeric value for the input with the given [param name]. This is used for triggers and grip sensors.
*/
//go:nosplit
func (self class) GetFloat(name gd.StringName) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRController3D.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Vector2] for the input with the given [param name]. This is used for thumbsticks and thumbpads found on many controllers.
*/
//go:nosplit
func (self class) GetVector2(name gd.StringName) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRController3D.Bind_get_vector2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the hand holding this controller, if known. See [enum XRPositionalTracker.TrackerHand].
*/
//go:nosplit
func (self class) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRPositionalTrackerTrackerHand](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRController3D.Bind_get_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnButtonPressed(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_pressed"), gc.Callable(cb), 0)
}


func (self Go) OnButtonReleased(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_released"), gc.Callable(cb), 0)
}


func (self Go) OnInputFloatChanged(cb func(name string, value float64)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("input_float_changed"), gc.Callable(cb), 0)
}


func (self Go) OnInputVector2Changed(cb func(name string, value gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("input_vector2_changed"), gc.Callable(cb), 0)
}


func (self Go) OnProfileChanged(cb func(role string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("profile_changed"), gc.Callable(cb), 0)
}


func (self class) AsXRController3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRController3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRNode3D() XRNode3D.GD { return *((*XRNode3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRNode3D() XRNode3D.Go { return *((*XRNode3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRNode3D(), name)
	}
}
func init() {classdb.Register("XRController3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
