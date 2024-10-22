package CollisionObject3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract base class for 3D physics objects. [CollisionObject3D] can hold any number of [Shape3D]s for collision. Each shape must be assigned to a [i]shape owner[/i]. Shape owners are not nodes and do not appear in the editor, but are accessible through code using the [code]shape_owner_*[/code] methods.
[b]Warning:[/b] With a non-uniform scale, this node will likely not behave as expected. It is advised to keep its scale the same on all axes and adjust its collision shape(s) instead.
	// CollisionObject3D methods that can be overridden by a [Class] that extends it.
	type CollisionObject3D interface {
		//Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
		//[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
		InputEvent(camera gdclass.Camera3D, event gdclass.InputEvent, event_position gd.Vector3, normal gd.Vector3, shape_idx int) 
		//Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
		MouseEnter() 
		//Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
		MouseExit() 
	}

*/
type Go [1]classdb.CollisionObject3D

/*
Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (Go) _input_event(impl func(ptr unsafe.Pointer, camera gdclass.Camera3D, event gdclass.InputEvent, event_position gd.Vector3, normal gd.Vector3, shape_idx int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var camera gdclass.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var event gdclass.InputEvent
		event[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var event_position = gd.UnsafeGet[gd.Vector3](p_args,2)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,3)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, camera, event, event_position, normal, int(shape_idx))
		gc.End()
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (Go) _mouse_enter(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (Go) _mouse_exit(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Returns the object's [RID].
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
Creates a new shape owner for the given object. Returns [code]owner_id[/code] of the new owner for future reference.
*/
func (self Go) CreateShapeOwner(owner gd.Object) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).CreateShapeOwner(owner)))
}

/*
Removes the given shape owner.
*/
func (self Go) RemoveShapeOwner(owner_id int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveShapeOwner(gd.Int(owner_id))
}

/*
Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
*/
func (self Go) GetShapeOwners() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetShapeOwners(gc).AsSlice())
}

/*
Sets the [Transform3D] of the given shape owner.
*/
func (self Go) ShapeOwnerSetTransform(owner_id int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapeOwnerSetTransform(gd.Int(owner_id), transform)
}

/*
Returns the shape owner's [Transform3D].
*/
func (self Go) ShapeOwnerGetTransform(owner_id int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).ShapeOwnerGetTransform(gd.Int(owner_id)))
}

/*
Returns the parent object of the given shape owner.
*/
func (self Go) ShapeOwnerGetOwner(owner_id int) gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).ShapeOwnerGetOwner(gc, gd.Int(owner_id)))
}

/*
If [code]true[/code], disables the given shape owner.
*/
func (self Go) ShapeOwnerSetDisabled(owner_id int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapeOwnerSetDisabled(gd.Int(owner_id), disabled)
}

/*
If [code]true[/code], the shape owner and its shapes are disabled.
*/
func (self Go) IsShapeOwnerDisabled(owner_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsShapeOwnerDisabled(gd.Int(owner_id)))
}

/*
Adds a [Shape3D] to the shape owner.
*/
func (self Go) ShapeOwnerAddShape(owner_id int, shape gdclass.Shape3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapeOwnerAddShape(gd.Int(owner_id), shape)
}

/*
Returns the number of shapes the given shape owner contains.
*/
func (self Go) ShapeOwnerGetShapeCount(owner_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapeOwnerGetShapeCount(gd.Int(owner_id))))
}

/*
Returns the [Shape3D] with the given ID from the given shape owner.
*/
func (self Go) ShapeOwnerGetShape(owner_id int, shape_id int) gdclass.Shape3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Shape3D(class(self).ShapeOwnerGetShape(gc, gd.Int(owner_id), gd.Int(shape_id)))
}

/*
Returns the child index of the [Shape3D] with the given ID from the given shape owner.
*/
func (self Go) ShapeOwnerGetShapeIndex(owner_id int, shape_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapeOwnerGetShapeIndex(gd.Int(owner_id), gd.Int(shape_id))))
}

/*
Removes a shape from the given shape owner.
*/
func (self Go) ShapeOwnerRemoveShape(owner_id int, shape_id int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapeOwnerRemoveShape(gd.Int(owner_id), gd.Int(shape_id))
}

/*
Removes all shapes from the shape owner.
*/
func (self Go) ShapeOwnerClearShapes(owner_id int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapeOwnerClearShapes(gd.Int(owner_id))
}

/*
Returns the [code]owner_id[/code] of the given shape.
*/
func (self Go) ShapeFindOwner(shape_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapeFindOwner(gd.Int(shape_index))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CollisionObject3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CollisionObject3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DisableMode() classdb.CollisionObject3DDisableMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CollisionObject3DDisableMode(class(self).GetDisableMode())
}

func (self Go) SetDisableMode(value classdb.CollisionObject3DDisableMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableMode(value)
}

func (self Go) CollisionLayer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionLayer()))
}

func (self Go) SetCollisionLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) CollisionPriority() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCollisionPriority()))
}

func (self Go) SetCollisionPriority(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionPriority(gd.Float(value))
}

func (self Go) InputRayPickable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRayPickable())
}

func (self Go) SetInputRayPickable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRayPickable(value)
}

func (self Go) InputCaptureOnDrag() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCaptureInputOnDrag())
}

func (self Go) SetInputCaptureOnDrag(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaptureInputOnDrag(value)
}

/*
Receives unhandled [InputEvent]s. [param event_position] is the location in world space of the mouse pointer on the surface of the shape with index [param shape_idx] and [param normal] is the normal vector of the surface at that point. Connect to the [signal input_event] signal to easily pick up these events.
[b]Note:[/b] [method _input_event] requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set.
*/
func (class) _input_event(impl func(ptr unsafe.Pointer, camera gdclass.Camera3D, event gdclass.InputEvent, event_position gd.Vector3, normal gd.Vector3, shape_idx gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var camera gdclass.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var event gdclass.InputEvent
		event[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var event_position = gd.UnsafeGet[gd.Vector3](p_args,2)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,3)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, camera, event, event_position, normal, shape_idx)
		ctx.End()
	}
}

/*
Called when the mouse pointer enters any of this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (class) _mouse_enter(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the mouse pointer exits all this object's shapes. Requires [member input_ray_pickable] to be [code]true[/code] and at least one [member collision_layer] bit to be set. Note that moving between different shapes within a single [CollisionObject3D] won't cause this function to be called.
*/
func (class) _mouse_exit(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableMode(mode classdb.CollisionObject3DDisableMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisableMode() classdb.CollisionObject3DDisableMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CollisionObject3DDisableMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRayPickable(ray_pickable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ray_pickable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRayPickable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_is_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaptureInputOnDrag(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_set_capture_input_on_drag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCaptureInputOnDrag() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_capture_input_on_drag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the object's [RID].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new shape owner for the given object. Returns [code]owner_id[/code] of the new owner for future reference.
*/
//go:nosplit
func (self class) CreateShapeOwner(owner gd.Object) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(owner.AsPointer())[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_create_shape_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the given shape owner.
*/
//go:nosplit
func (self class) RemoveShapeOwner(owner_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_remove_shape_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
*/
//go:nosplit
func (self class) GetShapeOwners(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_get_shape_owners, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the [Transform3D] of the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerSetTransform(owner_id gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the shape owner's [Transform3D].
*/
//go:nosplit
func (self class) ShapeOwnerGetTransform(owner_id gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the parent object of the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetOwner(ctx gd.Lifetime, owner_id gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_get_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
If [code]true[/code], disables the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerSetDisabled(owner_id gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the shape owner and its shapes are disabled.
*/
//go:nosplit
func (self class) IsShapeOwnerDisabled(owner_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_is_shape_owner_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a [Shape3D] to the shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerAddShape(owner_id gd.Int, shape gdclass.Shape3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, mmm.Get(shape[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_add_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of shapes the given shape owner contains.
*/
//go:nosplit
func (self class) ShapeOwnerGetShapeCount(owner_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Shape3D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShape(ctx gd.Lifetime, owner_id gd.Int, shape_id gd.Int) gdclass.Shape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the child index of the [Shape3D] with the given ID from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerGetShapeIndex(owner_id gd.Int, shape_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_get_shape_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a shape from the given shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerRemoveShape(owner_id gd.Int, shape_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	callframe.Arg(frame, shape_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_remove_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all shapes from the shape owner.
*/
//go:nosplit
func (self class) ShapeOwnerClearShapes(owner_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, owner_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_owner_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]owner_id[/code] of the given shape.
*/
//go:nosplit
func (self class) ShapeFindOwner(shape_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionObject3D.Bind_shape_find_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnInputEvent(cb func(camera gdclass.Node, event gdclass.InputEvent, event_position gd.Vector3, normal gd.Vector3, shape_idx int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("input_event"), gc.Callable(cb), 0)
}


func (self Go) OnMouseEntered(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("mouse_entered"), gc.Callable(cb), 0)
}


func (self Go) OnMouseExited(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("mouse_exited"), gc.Callable(cb), 0)
}


func (self class) AsCollisionObject3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_input_event": return reflect.ValueOf(self._input_event);
	case "_mouse_enter": return reflect.ValueOf(self._mouse_enter);
	case "_mouse_exit": return reflect.ValueOf(self._mouse_exit);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_input_event": return reflect.ValueOf(self._input_event);
	case "_mouse_enter": return reflect.ValueOf(self._mouse_enter);
	case "_mouse_exit": return reflect.ValueOf(self._mouse_exit);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("CollisionObject3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DisableMode = classdb.CollisionObject3DDisableMode

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
