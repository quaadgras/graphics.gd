// Package CollisionObject2D provides methods for working with CollisionObject2D object instances.
package CollisionObject2D

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
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Abstract base class for 2D physics objects. [CollisionObject2D] can hold any number of [Shape2D]s for collision. Each shape must be assigned to a [i]shape owner[/i]. Shape owners are not nodes and do not appear in the editor, but are accessible through code using the [code]shape_owner_*[/code] methods.
[b]Note:[/b] Only collisions between objects within the same canvas ([Viewport] canvas or [CanvasLayer]) are supported. The behavior of collisions between objects in different canvases is undefined.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=CollisionObject2D)
*/
type Instance [1]gdclass.CollisionObject2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCollisionObject2D() Instance
}
type Interface interface {
	//Accepts unhandled [InputEvent]s. [param shape_idx] is the child index of the clicked [Shape2D]. Connect to [signal input_event] to easily pick up these events.
	//[b]Note:[/b] [method _input_event] requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
	InputEvent(viewport [1]gdclass.Viewport, event [1]gdclass.InputEvent, shape_idx int)
	//Called when the mouse pointer enters any of this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
	MouseEnter()
	//Called when the mouse pointer exits all this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
	MouseExit()
	//Called when the mouse pointer enters any of this object's shapes or moves from one shape to another. [param shape_idx] is the child index of the newly entered [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
	MouseShapeEnter(shape_idx int)
	//Called when the mouse pointer exits any of this object's shapes. [param shape_idx] is the child index of the exited [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
	MouseShapeExit(shape_idx int)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) InputEvent(viewport [1]gdclass.Viewport, event [1]gdclass.InputEvent, shape_idx int) {
	return
}
func (self implementation) MouseEnter()                   { return }
func (self implementation) MouseExit()                    { return }
func (self implementation) MouseShapeEnter(shape_idx int) { return }
func (self implementation) MouseShapeExit(shape_idx int)  { return }

/*
Accepts unhandled [InputEvent]s. [param shape_idx] is the child index of the clicked [Shape2D]. Connect to [signal input_event] to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (Instance) _input_event(impl func(ptr unsafe.Pointer, viewport [1]gdclass.Viewport, event [1]gdclass.InputEvent, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport = [1]gdclass.Viewport{pointers.New[gdclass.Viewport]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(event[0])
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport, event, int(shape_idx))
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
*/
func (Instance) _mouse_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
*/
func (Instance) _mouse_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer enters any of this object's shapes or moves from one shape to another. [param shape_idx] is the child index of the newly entered [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
*/
func (Instance) _mouse_shape_enter(impl func(ptr unsafe.Pointer, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(shape_idx))
	}
}

/*
Called when the mouse pointer exits any of this object's shapes. [param shape_idx] is the child index of the exited [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
*/
func (Instance) _mouse_shape_exit(impl func(ptr unsafe.Pointer, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(shape_idx))
	}
}

/*
Returns the object's [RID].
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
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
Sets the [Transform2D] of the given shape owner.
*/
func (self Instance) ShapeOwnerSetTransform(owner_id int, transform Transform2D.OriginXY) {
	class(self).ShapeOwnerSetTransform(gd.Int(owner_id), gd.Transform2D(transform))
}

/*
Returns the shape owner's [Transform2D].
*/
func (self Instance) ShapeOwnerGetTransform(owner_id int) Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).ShapeOwnerGetTransform(gd.Int(owner_id)))
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
If [param enable] is [code]true[/code], collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
*/
func (self Instance) ShapeOwnerSetOneWayCollision(owner_id int, enable bool) {
	class(self).ShapeOwnerSetOneWayCollision(gd.Int(owner_id), enable)
}

/*
Returns [code]true[/code] if collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
*/
func (self Instance) IsShapeOwnerOneWayCollisionEnabled(owner_id int) bool {
	return bool(class(self).IsShapeOwnerOneWayCollisionEnabled(gd.Int(owner_id)))
}

/*
Sets the [code]one_way_collision_margin[/code] of the shape owner identified by given [param owner_id] to [param margin] pixels.
*/
func (self Instance) ShapeOwnerSetOneWayCollisionMargin(owner_id int, margin Float.X) {
	class(self).ShapeOwnerSetOneWayCollisionMargin(gd.Int(owner_id), gd.Float(margin))
}

/*
Returns the [code]one_way_collision_margin[/code] of the shape owner identified by given [param owner_id].
*/
func (self Instance) GetShapeOwnerOneWayCollisionMargin(owner_id int) Float.X {
	return Float.X(Float.X(class(self).GetShapeOwnerOneWayCollisionMargin(gd.Int(owner_id))))
}

/*
Adds a [Shape2D] to the shape owner.
*/
func (self Instance) ShapeOwnerAddShape(owner_id int, shape [1]gdclass.Shape2D) {
	class(self).ShapeOwnerAddShape(gd.Int(owner_id), shape)
}

/*
Returns the number of shapes the given shape owner contains.
*/
func (self Instance) ShapeOwnerGetShapeCount(owner_id int) int {
	return int(int(class(self).ShapeOwnerGetShapeCount(gd.Int(owner_id))))
}

/*
Returns the [Shape2D] with the given ID from the given shape owner.
*/
func (self Instance) ShapeOwnerGetShape(owner_id int, shape_id int) [1]gdclass.Shape2D {
	return [1]gdclass.Shape2D(class(self).ShapeOwnerGetShape(gd.Int(owner_id), gd.Int(shape_id)))
}

/*
Returns the child index of the [Shape2D] with the given ID from the given shape owner.
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
type class [1]gdclass.CollisionObject2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionObject2D"))
	casted := Instance{*(*gdclass.CollisionObject2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) DisableMode() gdclass.CollisionObject2DDisableMode {
	return gdclass.CollisionObject2DDisableMode(class(self).GetDisableMode())
}

func (self Instance) SetDisableMode(value gdclass.CollisionObject2DDisableMode) {
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

func (self Instance) InputPickable() bool {
	return bool(class(self).IsPickable())
}

func (self Instance) SetInputPickable(value bool) {
	class(self).SetPickable(value)
}

/*
Accepts unhandled [InputEvent]s. [param shape_idx] is the child index of the clicked [Shape2D]. Connect to [signal input_event] to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (class) _input_event(impl func(ptr unsafe.Pointer, viewport [1]gdclass.Viewport, event [1]gdclass.InputEvent, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport = [1]gdclass.Viewport{pointers.New[gdclass.Viewport]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(event[0])
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport, event, shape_idx)
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
*/
func (class) _mouse_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject2D] won't cause this function to be called.
*/
func (class) _mouse_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the mouse pointer enters any of this object's shapes or moves from one shape to another. [param shape_idx] is the child index of the newly entered [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
*/
func (class) _mouse_shape_enter(impl func(ptr unsafe.Pointer, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape_idx)
	}
}

/*
Called when the mouse pointer exits any of this object's shapes. [param shape_idx] is the child index of the exited [Shape2D]. Requires [member input_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be called.
*/
func (class) _mouse_shape_exit(impl func(ptr unsafe.Pointer, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape_idx)
	}
}

/*
Returns the object's [RID].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableMode(mode gdclass.CollisionObject2DDisableMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_disable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableMode() gdclass.CollisionObject2DDisableMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CollisionObject2DDisableMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_disable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPickable(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_set_pickable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPickable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_is_pickable, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_create_shape_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_remove_shape_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
*/
//go:nosplit
func (self class) GetShapeOwners() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_shape_owners, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [Transform2D] of the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerSetTransform(owner_id gd.Int, transform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the shape owner's [Transform2D].
*/
//go:nosplit
func (self class) ShapeOwnerGetTransform(owner_id gd.Int) gd.Transform2D {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_get_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_set_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_is_shape_owner_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enable] is [code]true[/code], collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
*/
//go:nosplit
func (self class) ShapeOwnerSetOneWayCollision(owner_id gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_set_one_way_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
*/
//go:nosplit
func (self class) IsShapeOwnerOneWayCollisionEnabled(owner_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_is_shape_owner_one_way_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [code]one_way_collision_margin[/code] of the shape owner identified by given [param owner_id] to [param margin] pixels.
*/
//go:nosplit
func (self class) ShapeOwnerSetOneWayCollisionMargin(owner_id gd.Int, margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_set_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]one_way_collision_margin[/code] of the shape owner identified by given [param owner_id].
*/
//go:nosplit
func (self class) GetShapeOwnerOneWayCollisionMargin(owner_id gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_get_shape_owner_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a [Shape2D] to the shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerAddShape(owner_id gd.Int, shape [1]gdclass.Shape2D) {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_add_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Shape2D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShape(owner_id gd.Int, shape_id gd.Int) [1]gdclass.Shape2D {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shape2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the child index of the [Shape2D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShapeIndex(owner_id gd.Int, shape_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_get_shape_index, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_remove_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_owner_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionObject2D.Bind_shape_find_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnInputEvent(cb func(viewport [1]gdclass.Node, event [1]gdclass.InputEvent, shape_idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("input_event"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseShapeEntered(cb func(shape_idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseShapeExited(cb func(shape_idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_shape_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsCollisionObject2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCollisionObject2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced        { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance     { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_input_event":
		return reflect.ValueOf(self._input_event)
	case "_mouse_enter":
		return reflect.ValueOf(self._mouse_enter)
	case "_mouse_exit":
		return reflect.ValueOf(self._mouse_exit)
	case "_mouse_shape_enter":
		return reflect.ValueOf(self._mouse_shape_enter)
	case "_mouse_shape_exit":
		return reflect.ValueOf(self._mouse_shape_exit)
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
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
	case "_mouse_shape_enter":
		return reflect.ValueOf(self._mouse_shape_enter)
	case "_mouse_shape_exit":
		return reflect.ValueOf(self._mouse_shape_exit)
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("CollisionObject2D", func(ptr gd.Object) any {
		return [1]gdclass.CollisionObject2D{*(*gdclass.CollisionObject2D)(unsafe.Pointer(&ptr))}
	})
}

type DisableMode = gdclass.CollisionObject2DDisableMode

const (
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], remove from the physics simulation to stop all physics interactions with this [CollisionObject2D].
	  Automatically re-added to the physics simulation when the [Node] is processed again.*/
	DisableModeRemove DisableMode = 0
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], make the body static. Doesn't affect [Area2D]. [PhysicsBody2D] can't be affected by forces or other bodies while static.
	  Automatically set [PhysicsBody2D] back to its original mode when the [Node] is processed again.*/
	DisableModeMakeStatic DisableMode = 1
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], do not affect the physics simulation.*/
	DisableModeKeepActive DisableMode = 2
)
