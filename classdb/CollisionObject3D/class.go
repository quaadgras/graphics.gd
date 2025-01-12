// Package CollisionObject3D provides methods for working with CollisionObject3D object instances.
package CollisionObject3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class for 3D physics objects. [CollisionObject3D] can hold any number of [Shape3D]s for collision. Each shape must be assigned to a [i]shape owner[/i]. Shape owners are not nodes and do not appear in the editor, but are accessible through code using the [code]shape_owner_*[/code] methods.
[b]Warning:[/b] With a non-uniform scale, this node will likely not behave as expected. It is advised to keep its scale the same on all axes and adjust its collision shape(s) instead.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=CollisionObject3D)
*/
type Instance [1]gdclass.CollisionObject3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCollisionObject3D() Instance
}
type Interface interface {
	//Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
	//[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
	InputEvent(camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent, event_position Vector3.XYZ, normal Vector3.XYZ, shape_idx int)
	//Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
	MouseEnter()
	//Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
	MouseExit()
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) InputEvent(camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent, event_position Vector3.XYZ, normal Vector3.XYZ, shape_idx int) {
	return
}
func (self Implementation) MouseEnter() { return }
func (self Implementation) MouseExit()  { return }

/*
Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (Instance) _input_event(impl func(ptr unsafe.Pointer, camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent, event_position Vector3.XYZ, normal Vector3.XYZ, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(camera[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(event[0])
		var event_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var normal = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, camera, event, event_position, normal, int(shape_idx))
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (Instance) _mouse_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (Instance) _mouse_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionLayerValue(layer_number int, value bool) {
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionLayerValue(layer_number int) bool {
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Returns the object's [RID].
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
}

/*
Creates a new shape owner for the given object. Returns [code]owner_id[/code] of the new owner for future reference.
*/
func (self Instance) CreateShapeOwner(owner Object.Instance) int {
	return int(int(class(self).CreateShapeOwner(owner)))
}

/*
Removes the given shape owner.
*/
func (self Instance) RemoveShapeOwner(owner_id int) {
	class(self).RemoveShapeOwner(gd.Int(owner_id))
}

/*
Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
*/
func (self Instance) GetShapeOwners() []int32 {
	return []int32(class(self).GetShapeOwners().AsSlice())
}

/*
Sets the [Transform3D] of the given shape owner.
*/
func (self Instance) ShapeOwnerSetTransform(owner_id int, transform Transform3D.BasisOrigin) {
	class(self).ShapeOwnerSetTransform(gd.Int(owner_id), gd.Transform3D(transform))
}

/*
Returns the shape owner's [Transform3D].
*/
func (self Instance) ShapeOwnerGetTransform(owner_id int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).ShapeOwnerGetTransform(gd.Int(owner_id)))
}

/*
Returns the parent object of the given shape owner.
*/
func (self Instance) ShapeOwnerGetOwner(owner_id int) Object.Instance {
	return Object.Instance(class(self).ShapeOwnerGetOwner(gd.Int(owner_id)))
}

/*
If [code]true[/code], disables the given shape owner.
*/
func (self Instance) ShapeOwnerSetDisabled(owner_id int, disabled bool) {
	class(self).ShapeOwnerSetDisabled(gd.Int(owner_id), disabled)
}

/*
If [code]true[/code], the shape owner and its shapes are disabled.
*/
func (self Instance) IsShapeOwnerDisabled(owner_id int) bool {
	return bool(class(self).IsShapeOwnerDisabled(gd.Int(owner_id)))
}

/*
Adds a [Shape3D] to the shape owner.
*/
func (self Instance) ShapeOwnerAddShape(owner_id int, shape [1]gdclass.Shape3D) {
	class(self).ShapeOwnerAddShape(gd.Int(owner_id), shape)
}

/*
Returns the number of shapes the given shape owner contains.
*/
func (self Instance) ShapeOwnerGetShapeCount(owner_id int) int {
	return int(int(class(self).ShapeOwnerGetShapeCount(gd.Int(owner_id))))
}

/*
Returns the [Shape3D] with the given ID from the given shape owner.
*/
func (self Instance) ShapeOwnerGetShape(owner_id int, shape_id int) [1]gdclass.Shape3D {
	return [1]gdclass.Shape3D(class(self).ShapeOwnerGetShape(gd.Int(owner_id), gd.Int(shape_id)))
}

/*
Returns the child index of the [Shape3D] with the given ID from the given shape owner.
*/
func (self Instance) ShapeOwnerGetShapeIndex(owner_id int, shape_id int) int {
	return int(int(class(self).ShapeOwnerGetShapeIndex(gd.Int(owner_id), gd.Int(shape_id))))
}

/*
Removes a shape from the given shape owner.
*/
func (self Instance) ShapeOwnerRemoveShape(owner_id int, shape_id int) {
	class(self).ShapeOwnerRemoveShape(gd.Int(owner_id), gd.Int(shape_id))
}

/*
Removes all shapes from the shape owner.
*/
func (self Instance) ShapeOwnerClearShapes(owner_id int) {
	class(self).ShapeOwnerClearShapes(gd.Int(owner_id))
}

/*
Returns the [code]owner_id[/code] of the given shape.
*/
func (self Instance) ShapeFindOwner(shape_index int) int {
	return int(int(class(self).ShapeFindOwner(gd.Int(shape_index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CollisionObject3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionObject3D"))
	casted := Instance{*(*gdclass.CollisionObject3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) DisableMode() gdclass.CollisionObject3DDisableMode {
	return gdclass.CollisionObject3DDisableMode(class(self).GetDisableMode())
}

func (self Instance) SetDisableMode(value gdclass.CollisionObject3DDisableMode) {
	class(self).SetDisableMode(value)
}

func (self Instance) CollisionLayer() int {
	return int(int(class(self).GetCollisionLayer()))
}

func (self Instance) SetCollisionLayer(value int) {
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) CollisionPriority() Float.X {
	return Float.X(Float.X(class(self).GetCollisionPriority()))
}

func (self Instance) SetCollisionPriority(value Float.X) {
	class(self).SetCollisionPriority(gd.Float(value))
}

func (self Instance) InputRayPickable() bool {
	return bool(class(self).IsRayPickable())
}

func (self Instance) SetInputRayPickable(value bool) {
	class(self).SetRayPickable(value)
}

func (self Instance) InputCaptureOnDrag() bool {
	return bool(class(self).GetCaptureInputOnDrag())
}

func (self Instance) SetInputCaptureOnDrag(value bool) {
	class(self).SetCaptureInputOnDrag(value)
}

/*
Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (class) _input_event(impl func(ptr unsafe.Pointer, camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent, event_position gd.Vector3, normal gd.Vector3, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(camera[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(event[0])
		var event_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var normal = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, camera, event, event_position, normal, shape_idx)
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (class) _mouse_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (class) _mouse_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableMode(mode gdclass.CollisionObject3DDisableMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_disable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableMode() gdclass.CollisionObject3DDisableMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CollisionObject3DDisableMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_disable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRayPickable(ray_pickable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ray_pickable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRayPickable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_is_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaptureInputOnDrag(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_set_capture_input_on_drag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCaptureInputOnDrag() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_capture_input_on_drag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the object's [RID].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new shape owner for the given object. Returns [code]owner_id[/code] of the new owner for future reference.
*/
//go:nosplit
func (self class) CreateShapeOwner(owner [1]gd.Object) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(owner[0])[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_create_shape_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the given shape owner.
*/
//go:nosplit
func (self class) RemoveShapeOwner(owner_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_remove_shape_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
*/
//go:nosplit
func (self class) GetShapeOwners() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_get_shape_owners, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [Transform3D] of the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerSetTransform(owner_id gd.Int, transform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the shape owner's [Transform3D].
*/
//go:nosplit
func (self class) ShapeOwnerGetTransform(owner_id gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the parent object of the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetOwner(owner_id gd.Int) [1]gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_get_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}

/*
If [code]true[/code], disables the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerSetDisabled(owner_id gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_set_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [code]true[/code], the shape owner and its shapes are disabled.
*/
//go:nosplit
func (self class) IsShapeOwnerDisabled(owner_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_is_shape_owner_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a [Shape3D] to the shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerAddShape(owner_id gd.Int, shape [1]gdclass.Shape3D) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_add_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of shapes the given shape owner contains.
*/
//go:nosplit
func (self class) ShapeOwnerGetShapeCount(owner_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Shape3D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShape(owner_id gd.Int, shape_id gd.Int) [1]gdclass.Shape3D {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the child index of the [Shape3D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShapeIndex(owner_id gd.Int, shape_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_get_shape_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes a shape from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerRemoveShape(owner_id gd.Int, shape_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_remove_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all shapes from the shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerClearShapes(owner_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_owner_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]owner_id[/code] of the given shape.
*/
//go:nosplit
func (self class) ShapeFindOwner(shape_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, shape_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject3D.Bind_shape_find_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnInputEvent(cb func(camera [1]gdclass.Node, event [1]gdclass.InputEvent, event_position Vector3.XYZ, normal Vector3.XYZ, shape_idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("input_event"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsCollisionObject3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCollisionObject3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced        { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance     { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced            { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance         { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_input_event":
		return reflect.ValueOf(self._input_event)
	case "_mouse_enter":
		return reflect.ValueOf(self._mouse_enter)
	case "_mouse_exit":
		return reflect.ValueOf(self._mouse_exit)
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_input_event":
		return reflect.ValueOf(self._input_event)
	case "_mouse_enter":
		return reflect.ValueOf(self._mouse_enter)
	case "_mouse_exit":
		return reflect.ValueOf(self._mouse_exit)
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("CollisionObject3D", func(ptr gd.Object) any {
		return [1]gdclass.CollisionObject3D{*(*gdclass.CollisionObject3D)(unsafe.Pointer(&ptr))}
	})
}

type DisableMode = gdclass.CollisionObject3DDisableMode

const (
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], remove from the physics simulation to stop all physics interactions with this [CollisionObject3D].
	  Automatically re-added to the physics simulation when the [Node] is processed again.*/
	DisableModeRemove DisableMode = 0
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], make the body static. Doesn't affect [Area3D]. [PhysicsBody3D] can't be affected by forces or other bodies while static.
	  Automatically set [PhysicsBody3D] back to its original mode when the [Node] is processed again.*/
	DisableModeMakeStatic DisableMode = 1
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], do not affect the physics simulation.*/
	DisableModeKeepActive DisableMode = 2
)
