package AnimationNodeBlendSpace1D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationRootNode"
import "grow.graphics/gd/gdclass/AnimationNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource used by [AnimationNodeBlendTree].
[AnimationNodeBlendSpace1D] represents a virtual axis on which any type of [AnimationRootNode]s can be added using [method add_blend_point]. Outputs the linear blend of the two [AnimationRootNode]s adjacent to the current value.
You can set the extents of the axis with [member min_space] and [member max_space].

*/
type Go [1]classdb.AnimationNodeBlendSpace1D

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Go) AddBlendPoint(node gdclass.AnimationRootNode, pos float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddBlendPoint(node, gd.Float(pos), gd.Int(-1))
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
func (self Go) SetBlendPointPosition(point int, pos float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendPointPosition(gd.Int(point), gd.Float(pos))
}

/*
Returns the position of the point at index [param point].
*/
func (self Go) GetBlendPointPosition(point int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetBlendPointPosition(gd.Int(point))))
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
func (self Go) SetBlendPointNode(point int, node gdclass.AnimationRootNode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendPointNode(gd.Int(point), node)
}

/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
func (self Go) GetBlendPointNode(point int) gdclass.AnimationRootNode {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AnimationRootNode(class(self).GetBlendPointNode(gc, gd.Int(point)))
}

/*
Removes the point at index [param point] from the blend axis.
*/
func (self Go) RemoveBlendPoint(point int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveBlendPoint(gd.Int(point))
}

/*
Returns the number of points on the blend axis.
*/
func (self Go) GetBlendPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBlendPointCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeBlendSpace1D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationNodeBlendSpace1D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MinSpace() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinSpace()))
}

func (self Go) SetMinSpace(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinSpace(gd.Float(value))
}

func (self Go) MaxSpace() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxSpace()))
}

func (self Go) SetMaxSpace(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSpace(gd.Float(value))
}

func (self Go) Snap() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSnap()))
}

func (self Go) SetSnap(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSnap(gd.Float(value))
}

func (self Go) ValueLabel() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetValueLabel(gc).String())
}

func (self Go) SetValueLabel(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetValueLabel(gc.String(value))
}

func (self Go) BlendMode() classdb.AnimationNodeBlendSpace1DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AnimationNodeBlendSpace1DBlendMode(class(self).GetBlendMode())
}

func (self Go) SetBlendMode(value classdb.AnimationNodeBlendSpace1DBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendMode(value)
}

func (self Go) Sync() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingSync())
}

func (self Go) SetSync(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseSync(value)
}

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddBlendPoint(node gdclass.AnimationRootNode, pos gd.Float, at_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_add_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Updates the position of the point at index [param point] on the blend axis.
*/
//go:nosplit
func (self class) SetBlendPointPosition(point gd.Int, pos gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pos)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointPosition(point gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) SetBlendPointNode(point gd.Int, node gdclass.AnimationRootNode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointNode(ctx gd.Lifetime, point gd.Int) gdclass.AnimationRootNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AnimationRootNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes the point at index [param point] from the blend axis.
*/
//go:nosplit
func (self class) RemoveBlendPoint(point gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_remove_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of points on the blend axis.
*/
//go:nosplit
func (self class) GetBlendPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinSpace(min_space gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinSpace() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSpace(max_space gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxSpace() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnap(snap gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSnap() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetValueLabel(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_value_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetValueLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_value_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendMode(mode classdb.AnimationNodeBlendSpace1DBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.AnimationNodeBlendSpace1DBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeBlendSpace1DBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseSync(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_set_use_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingSync() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace1D.Bind_is_using_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeBlendSpace1D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeBlendSpace1D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationRootNode() AnimationRootNode.GD { return *((*AnimationRootNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationRootNode() AnimationRootNode.Go { return *((*AnimationRootNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.GD { return *((*AnimationNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() AnimationNode.Go { return *((*AnimationNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}
func init() {classdb.Register("AnimationNodeBlendSpace1D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type BlendMode = classdb.AnimationNodeBlendSpace1DBlendMode

const (
/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
