package AnimationNodeBlendSpace2D

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
[AnimationNodeBlendSpace1D] represents a virtual 2D space on which [AnimationRootNode]s are placed. Outputs the linear blend of the three adjacent animations using a [Vector2] weight. Adjacent in this context means the three [AnimationRootNode]s making up the triangle that contains the current value.
You can add vertices to the blend space with [method add_blend_point] and automatically triangulate it by setting [member auto_triangles] to [code]true[/code]. Otherwise, use [method add_triangle] and [method remove_triangle] to triangulate the blend space by hand.

*/
type Go [1]classdb.AnimationNodeBlendSpace2D

/*
Adds a new point that represents a [param node] at the position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Go) AddBlendPoint(node gdclass.AnimationRootNode, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddBlendPoint(node, pos, gd.Int(-1))
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
func (self Go) SetBlendPointPosition(point int, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendPointPosition(gd.Int(point), pos)
}

/*
Returns the position of the point at index [param point].
*/
func (self Go) GetBlendPointPosition(point int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetBlendPointPosition(gd.Int(point)))
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
func (self Go) SetBlendPointNode(point int, node gdclass.AnimationRootNode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendPointNode(gd.Int(point), node)
}

/*
Returns the [AnimationRootNode] referenced by the point at index [param point].
*/
func (self Go) GetBlendPointNode(point int) gdclass.AnimationRootNode {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AnimationRootNode(class(self).GetBlendPointNode(gc, gd.Int(point)))
}

/*
Removes the point at index [param point] from the blend space.
*/
func (self Go) RemoveBlendPoint(point int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveBlendPoint(gd.Int(point))
}

/*
Returns the number of points in the blend space.
*/
func (self Go) GetBlendPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBlendPointCount()))
}

/*
Creates a new triangle using three points [param x], [param y], and [param z]. Triangles can overlap. You can insert the triangle at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Go) AddTriangle(x int, y int, z int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddTriangle(gd.Int(x), gd.Int(y), gd.Int(z), gd.Int(-1))
}

/*
Returns the position of the point at index [param point] in the triangle of index [param triangle].
*/
func (self Go) GetTrianglePoint(triangle int, point int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTrianglePoint(gd.Int(triangle), gd.Int(point))))
}

/*
Removes the triangle at index [param triangle] from the blend space.
*/
func (self Go) RemoveTriangle(triangle int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveTriangle(gd.Int(triangle))
}

/*
Returns the number of triangles in the blend space.
*/
func (self Go) GetTriangleCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTriangleCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeBlendSpace2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationNodeBlendSpace2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) AutoTriangles() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAutoTriangles())
}

func (self Go) SetAutoTriangles(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoTriangles(value)
}

func (self Go) MinSpace() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetMinSpace())
}

func (self Go) SetMinSpace(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinSpace(value)
}

func (self Go) MaxSpace() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetMaxSpace())
}

func (self Go) SetMaxSpace(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSpace(value)
}

func (self Go) Snap() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetSnap())
}

func (self Go) SetSnap(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSnap(value)
}

func (self Go) XLabel() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetXLabel(gc).String())
}

func (self Go) SetXLabel(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetXLabel(gc.String(value))
}

func (self Go) YLabel() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetYLabel(gc).String())
}

func (self Go) SetYLabel(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetYLabel(gc.String(value))
}

func (self Go) BlendMode() classdb.AnimationNodeBlendSpace2DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AnimationNodeBlendSpace2DBlendMode(class(self).GetBlendMode())
}

func (self Go) SetBlendMode(value classdb.AnimationNodeBlendSpace2DBlendMode) {
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
Adds a new point that represents a [param node] at the position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddBlendPoint(node gdclass.AnimationRootNode, pos gd.Vector2, at_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_add_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Updates the position of the point at index [param point] on the blend axis.
*/
//go:nosplit
func (self class) SetBlendPointPosition(point gd.Int, pos gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pos)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointPosition(point gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [AnimationRootNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointNode(ctx gd.Lifetime, point gd.Int) gdclass.AnimationRootNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AnimationRootNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes the point at index [param point] from the blend space.
*/
//go:nosplit
func (self class) RemoveBlendPoint(point gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_remove_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of points in the blend space.
*/
//go:nosplit
func (self class) GetBlendPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new triangle using three points [param x], [param y], and [param z]. Triangles can overlap. You can insert the triangle at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddTriangle(x gd.Int, y gd.Int, z gd.Int, at_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, z)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_add_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the point at index [param point] in the triangle of index [param triangle].
*/
//go:nosplit
func (self class) GetTrianglePoint(triangle gd.Int, point gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, triangle)
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_triangle_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the triangle at index [param triangle] from the blend space.
*/
//go:nosplit
func (self class) RemoveTriangle(triangle gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, triangle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_remove_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of triangles in the blend space.
*/
//go:nosplit
func (self class) GetTriangleCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_triangle_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinSpace(min_space gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinSpace() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSpace(max_space gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxSpace() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnap(snap gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSnap() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetXLabel(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_x_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_x_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetYLabel(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_y_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetYLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_y_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoTriangles(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_auto_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoTriangles() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_auto_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendMode(mode classdb.AnimationNodeBlendSpace2DBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.AnimationNodeBlendSpace2DBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeBlendSpace2DBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_set_use_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingSync() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_is_using_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTrianglesUpdated(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("triangles_updated"), gc.Callable(cb), 0)
}


func (self class) AsAnimationNodeBlendSpace2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeBlendSpace2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AnimationNodeBlendSpace2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type BlendMode = classdb.AnimationNodeBlendSpace2DBlendMode

const (
/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
