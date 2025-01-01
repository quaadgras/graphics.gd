package XRController3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/XRNode3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This is a helper spatial node that is linked to the tracking of controllers. It also offers several handy passthroughs to the state of buttons and such on the controllers.
Controllers are linked by their ID. You can create controller nodes before the controllers are available. If your game always uses two controllers (one for each hand), you can predefine the controllers with ID 1 and 2; they will become active as soon as the controllers are identified. If you expect additional controllers to be used, you should react to the signals and add XRController3D nodes to your scene.
The position of the controller node is automatically updated by the [XRServer]. This makes this node ideal to add child nodes to visualize the controller.
As many XR runtimes now use a configurable action map all inputs are named.
*/
type Instance [1]classdb.XRController3D
type Any interface {
	gd.IsClass
	AsXRController3D() Instance
}

/*
Returns [code]true[/code] if the button with the given [param name] is pressed.
*/
func (self Instance) IsButtonPressed(name string) bool {
	return bool(class(self).IsButtonPressed(gd.NewStringName(name)))
}

/*
Returns a [Variant] for the input with the given [param name]. This works for any input type, the variant will be typed according to the actions configuration.
*/
func (self Instance) GetInput(name string) any {
	return any(class(self).GetInput(gd.NewStringName(name)).Interface())
}

/*
Returns a numeric value for the input with the given [param name]. This is used for triggers and grip sensors.
*/
func (self Instance) GetFloat(name string) Float.X {
	return Float.X(Float.X(class(self).GetFloat(gd.NewStringName(name))))
}

/*
Returns a [Vector2] for the input with the given [param name]. This is used for thumbsticks and thumbpads found on many controllers.
*/
func (self Instance) GetVector2(name string) Vector2.XY {
	return Vector2.XY(class(self).GetVector2(gd.NewStringName(name)))
}

/*
Returns the hand holding this controller, if known. See [enum XRPositionalTracker.TrackerHand].
*/
func (self Instance) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	return classdb.XRPositionalTrackerTrackerHand(class(self).GetTrackerHand())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRController3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRController3D"))
	return Instance{classdb.XRController3D(object)}
}

/*
Returns [code]true[/code] if the button with the given [param name] is pressed.
*/
//go:nosplit
func (self class) IsButtonPressed(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRController3D.Bind_is_button_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Variant] for the input with the given [param name]. This works for any input type, the variant will be typed according to the actions configuration.
*/
//go:nosplit
func (self class) GetInput(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRController3D.Bind_get_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a numeric value for the input with the given [param name]. This is used for triggers and grip sensors.
*/
//go:nosplit
func (self class) GetFloat(name gd.StringName) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRController3D.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Vector2] for the input with the given [param name]. This is used for thumbsticks and thumbpads found on many controllers.
*/
//go:nosplit
func (self class) GetVector2(name gd.StringName) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRController3D.Bind_get_vector2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the hand holding this controller, if known. See [enum XRPositionalTracker.TrackerHand].
*/
//go:nosplit
func (self class) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRPositionalTrackerTrackerHand](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRController3D.Bind_get_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnButtonPressed(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonReleased(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("button_released"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputFloatChanged(cb func(name string, value Float.X)) {
	self[0].AsObject().Connect(gd.NewStringName("input_float_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputVector2Changed(cb func(name string, value Vector2.XY)) {
	self[0].AsObject().Connect(gd.NewStringName("input_vector2_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProfileChanged(cb func(role string)) {
	self[0].AsObject().Connect(gd.NewStringName("profile_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsXRController3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRController3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRNode3D() XRNode3D.Advanced {
	return *((*XRNode3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRNode3D() XRNode3D.Instance {
	return *((*XRNode3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRNode3D(), name)
	}
}
func init() {
	classdb.Register("XRController3D", func(ptr gd.Object) any { return [1]classdb.XRController3D{classdb.XRController3D(ptr)} })
}
