// Package VisualInstance3D provides methods for working with VisualInstance3D object instances.
package VisualInstance3D

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
import "graphics.gd/variant/AABB"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [VisualInstance3D] is used to connect a resource to a visual representation. All visual 3D nodes inherit from the [VisualInstance3D]. In general, you should not access the [VisualInstance3D] properties directly as they are accessed and managed by the nodes that inherit from [VisualInstance3D]. [VisualInstance3D] is the node representation of the [RenderingServer] instance.

	// VisualInstance3D methods that can be overridden by a [Class] that extends it.
	type VisualInstance3D interface {
		GetAabb() AABB.PositionSize
	}
*/
type Instance [1]gdclass.VisualInstance3D
type Any interface {
	gd.IsClass
	AsVisualInstance3D() Instance
}

func (Instance) _get_aabb(impl func(ptr unsafe.Pointer) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.AABB(ret))
	}
}

/*
Sets the resource that is instantiated by this [VisualInstance3D], which changes how the engine handles the [VisualInstance3D] under the hood. Equivalent to [method RenderingServer.instance_set_base].
*/
func (self Instance) SetBase(base Resource.ID) {
	class(self).SetBase(base)
}

/*
Returns the RID of the resource associated with this [VisualInstance3D]. For example, if the Node is a [MeshInstance3D], this will return the RID of the associated [Mesh].
*/
func (self Instance) GetBase() Resource.ID {
	return Resource.ID(class(self).GetBase())
}

/*
Returns the RID of this instance. This RID is the same as the RID returned by [method RenderingServer.instance_create]. This RID is needed if you want to call [RenderingServer] functions directly on this [VisualInstance3D].
*/
func (self Instance) GetInstance() Resource.ID {
	return Resource.ID(class(self).GetInstance())
}

/*
Based on [param value], enables or disables the specified layer in the [member layers], given a [param layer_number] between 1 and 20.
*/
func (self Instance) SetLayerMaskValue(layer_number int, value bool) {
	class(self).SetLayerMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member layers] is enabled, given a [param layer_number] between 1 and 20.
*/
func (self Instance) GetLayerMaskValue(layer_number int) bool {
	return bool(class(self).GetLayerMaskValue(gd.Int(layer_number)))
}

/*
Returns the [AABB] (also known as the bounding box) for this [VisualInstance3D].
*/
func (self Instance) GetAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetAabb())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualInstance3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualInstance3D"))
	return Instance{*(*gdclass.VisualInstance3D)(unsafe.Pointer(&object))}
}

func (self Instance) Layers() int {
	return int(int(class(self).GetLayerMask()))
}

func (self Instance) SetLayers(value int) {
	class(self).SetLayerMask(gd.Int(value))
}

func (self Instance) SortingOffset() Float.X {
	return Float.X(Float.X(class(self).GetSortingOffset()))
}

func (self Instance) SetSortingOffset(value Float.X) {
	class(self).SetSortingOffset(gd.Float(value))
}

func (self Instance) SortingUseAabbCenter() bool {
	return bool(class(self).IsSortingUseAabbCenter())
}

func (self Instance) SetSortingUseAabbCenter(value bool) {
	class(self).SetSortingUseAabbCenter(value)
}

func (class) _get_aabb(impl func(ptr unsafe.Pointer) gd.AABB) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets the resource that is instantiated by this [VisualInstance3D], which changes how the engine handles the [VisualInstance3D] under the hood. Equivalent to [method RenderingServer.instance_set_base].
*/
//go:nosplit
func (self class) SetBase(base gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, base)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_set_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the RID of the resource associated with this [VisualInstance3D]. For example, if the Node is a [MeshInstance3D], this will return the RID of the associated [Mesh].
*/
//go:nosplit
func (self class) GetBase() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the RID of this instance. This RID is the same as the RID returned by [method RenderingServer.instance_create]. This RID is needed if you want to call [RenderingServer] functions directly on this [VisualInstance3D].
*/
//go:nosplit
func (self class) GetInstance() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLayerMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member layers], given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) SetLayerMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_set_layer_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member layers] is enabled, given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) GetLayerMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_layer_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSortingOffset(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_set_sorting_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSortingOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_sorting_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSortingUseAabbCenter(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_set_sorting_use_aabb_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSortingUseAabbCenter() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_is_sorting_use_aabb_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AABB] (also known as the bounding box) for this [VisualInstance3D].
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualInstance3D.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualInstance3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualInstance3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced       { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance    { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("VisualInstance3D", func(ptr gd.Object) any {
		return [1]gdclass.VisualInstance3D{*(*gdclass.VisualInstance3D)(unsafe.Pointer(&ptr))}
	})
}
