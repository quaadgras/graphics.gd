// Package CSGShape3D provides methods for working with CSGShape3D object instances.
package CSGShape3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
This is the CSG base class that provides CSG operation support to the various CSG nodes in Godot.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.
*/
type Instance [1]gdclass.CSGShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCSGShape3D() Instance
}

/*
Returns [code]true[/code] if this is a root shape and is thus the object that is rendered.
*/
func (self Instance) IsRootShape() bool {
	return bool(class(self).IsRootShape())
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
Returns an [Array] with two elements, the first is the [Transform3D] of this node and the second is the root [Mesh] of this node. Only works when this node is the root shape.
*/
func (self Instance) GetMeshes() Array.Any {
	return Array.Any(class(self).GetMeshes())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CSGShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CSGShape3D"))
	casted := Instance{*(*gdclass.CSGShape3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Operation() gdclass.CSGShape3DOperation {
	return gdclass.CSGShape3DOperation(class(self).GetOperation())
}

func (self Instance) SetOperation(value gdclass.CSGShape3DOperation) {
	class(self).SetOperation(value)
}

func (self Instance) Snap() Float.X {
	return Float.X(Float.X(class(self).GetSnap()))
}

func (self Instance) SetSnap(value Float.X) {
	class(self).SetSnap(gd.Float(value))
}

func (self Instance) CalculateTangents() bool {
	return bool(class(self).IsCalculatingTangents())
}

func (self Instance) SetCalculateTangents(value bool) {
	class(self).SetCalculateTangents(value)
}

func (self Instance) UseCollision() bool {
	return bool(class(self).IsUsingCollision())
}

func (self Instance) SetUseCollision(value bool) {
	class(self).SetUseCollision(value)
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

/*
Returns [code]true[/code] if this is a root shape and is thus the object that is rendered.
*/
//go:nosplit
func (self class) IsRootShape() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_root_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOperation(operation gdclass.CSGShape3DOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, operation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperation() gdclass.CSGShape3DOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CSGShape3DOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap(snap gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnap() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCollision(operation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, operation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_use_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCollision() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_using_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCalculateTangents(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_calculate_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCalculatingTangents() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_calculating_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Array] with two elements, the first is the [Transform3D] of this node and the second is the root [Mesh] of this node. Only works when this node is the root shape.
*/
//go:nosplit
func (self class) GetMeshes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCSGShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCSGShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Advanced(self.AsGeometryInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Instance(self.AsGeometryInstance3D()), name)
	}
}
func init() {
	gdclass.Register("CSGShape3D", func(ptr gd.Object) any { return [1]gdclass.CSGShape3D{*(*gdclass.CSGShape3D)(unsafe.Pointer(&ptr))} })
}

type Operation = gdclass.CSGShape3DOperation

const (
	/*Geometry of both primitives is merged, intersecting geometry is removed.*/
	OperationUnion Operation = 0
	/*Only intersecting geometry remains, the rest is removed.*/
	OperationIntersection Operation = 1
	/*The second shape is subtracted from the first, leaving a dent with its shape.*/
	OperationSubtraction Operation = 2
)
