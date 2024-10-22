package VisualInstance3D

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
The [VisualInstance3D] is used to connect a resource to a visual representation. All visual 3D nodes inherit from the [VisualInstance3D]. In general, you should not access the [VisualInstance3D] properties directly as they are accessed and managed by the nodes that inherit from [VisualInstance3D]. [VisualInstance3D] is the node representation of the [RenderingServer] instance.
	// VisualInstance3D methods that can be overridden by a [Class] that extends it.
	type VisualInstance3D interface {
		GetAabb() gd.AABB
	}

*/
type Go [1]classdb.VisualInstance3D
func (Go) _get_aabb(impl func(ptr unsafe.Pointer) gd.AABB, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Sets the resource that is instantiated by this [VisualInstance3D], which changes how the engine handles the [VisualInstance3D] under the hood. Equivalent to [method RenderingServer.instance_set_base].
*/
func (self Go) SetBase(base gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBase(base)
}

/*
Returns the RID of the resource associated with this [VisualInstance3D]. For example, if the Node is a [MeshInstance3D], this will return the RID of the associated [Mesh].
*/
func (self Go) GetBase() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetBase())
}

/*
Returns the RID of this instance. This RID is the same as the RID returned by [method RenderingServer.instance_create]. This RID is needed if you want to call [RenderingServer] functions directly on this [VisualInstance3D].
*/
func (self Go) GetInstance() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetInstance())
}

/*
Based on [param value], enables or disables the specified layer in the [member layers], given a [param layer_number] between 1 and 20.
*/
func (self Go) SetLayerMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member layers] is enabled, given a [param layer_number] between 1 and 20.
*/
func (self Go) GetLayerMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetLayerMaskValue(gd.Int(layer_number)))
}

/*
Returns the [AABB] (also known as the bounding box) for this [VisualInstance3D].
*/
func (self Go) GetAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(class(self).GetAabb())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualInstance3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualInstance3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Layers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLayerMask()))
}

func (self Go) SetLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerMask(gd.Int(value))
}

func (self Go) SortingOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSortingOffset()))
}

func (self Go) SetSortingOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSortingOffset(gd.Float(value))
}

func (self Go) SortingUseAabbCenter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSortingUseAabbCenter())
}

func (self Go) SetSortingUseAabbCenter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSortingUseAabbCenter(value)
}

func (class) _get_aabb(impl func(ptr unsafe.Pointer) gd.AABB, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Sets the resource that is instantiated by this [VisualInstance3D], which changes how the engine handles the [VisualInstance3D] under the hood. Equivalent to [method RenderingServer.instance_set_base].
*/
//go:nosplit
func (self class) SetBase(base gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_set_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the RID of the resource associated with this [VisualInstance3D]. For example, if the Node is a [MeshInstance3D], this will return the RID of the associated [Mesh].
*/
//go:nosplit
func (self class) GetBase() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the RID of this instance. This RID is the same as the RID returned by [method RenderingServer.instance_create]. This RID is needed if you want to call [RenderingServer] functions directly on this [VisualInstance3D].
*/
//go:nosplit
func (self class) GetInstance() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLayerMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member layers], given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) SetLayerMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_set_layer_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member layers] is enabled, given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) GetLayerMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_layer_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSortingOffset(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_set_sorting_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSortingOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_sorting_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSortingUseAabbCenter(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_set_sorting_use_aabb_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSortingUseAabbCenter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_is_sorting_use_aabb_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [AABB] (also known as the bounding box) for this [VisualInstance3D].
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualInstance3D.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualInstance3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_aabb": return reflect.ValueOf(self._get_aabb);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_aabb": return reflect.ValueOf(self._get_aabb);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("VisualInstance3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
