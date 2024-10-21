package AnimationNodeBlendSpace2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationRootNode"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource used by [AnimationNodeBlendTree].
[AnimationNodeBlendSpace1D] represents a virtual 2D space on which [AnimationRootNode]s are placed. Outputs the linear blend of the three adjacent animations using a [Vector2] weight. Adjacent in this context means the three [AnimationRootNode]s making up the triangle that contains the current value.
You can add vertices to the blend space with [method add_blend_point] and automatically triangulate it by setting [member auto_triangles] to [code]true[/code]. Otherwise, use [method add_triangle] and [method remove_triangle] to triangulate the blend space by hand.

*/
type Simple [1]classdb.AnimationNodeBlendSpace2D
func (self Simple) AddBlendPoint(node [1]classdb.AnimationRootNode, pos gd.Vector2, at_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddBlendPoint(node, pos, gd.Int(at_index))
}
func (self Simple) SetBlendPointPosition(point int, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendPointPosition(gd.Int(point), pos)
}
func (self Simple) GetBlendPointPosition(point int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetBlendPointPosition(gd.Int(point)))
}
func (self Simple) SetBlendPointNode(point int, node [1]classdb.AnimationRootNode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendPointNode(gd.Int(point), node)
}
func (self Simple) GetBlendPointNode(point int) [1]classdb.AnimationRootNode {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationRootNode(Expert(self).GetBlendPointNode(gc, gd.Int(point)))
}
func (self Simple) RemoveBlendPoint(point int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveBlendPoint(gd.Int(point))
}
func (self Simple) GetBlendPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBlendPointCount()))
}
func (self Simple) AddTriangle(x int, y int, z int, at_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTriangle(gd.Int(x), gd.Int(y), gd.Int(z), gd.Int(at_index))
}
func (self Simple) GetTrianglePoint(triangle int, point int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTrianglePoint(gd.Int(triangle), gd.Int(point))))
}
func (self Simple) RemoveTriangle(triangle int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTriangle(gd.Int(triangle))
}
func (self Simple) GetTriangleCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTriangleCount()))
}
func (self Simple) SetMinSpace(min_space gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinSpace(min_space)
}
func (self Simple) GetMinSpace() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMinSpace())
}
func (self Simple) SetMaxSpace(max_space gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxSpace(max_space)
}
func (self Simple) GetMaxSpace() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMaxSpace())
}
func (self Simple) SetSnap(snap gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnap(snap)
}
func (self Simple) GetSnap() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetSnap())
}
func (self Simple) SetXLabel(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetXLabel(gc.String(text))
}
func (self Simple) GetXLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetXLabel(gc).String())
}
func (self Simple) SetYLabel(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetYLabel(gc.String(text))
}
func (self Simple) GetYLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetYLabel(gc).String())
}
func (self Simple) SetAutoTriangles(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoTriangles(enable)
}
func (self Simple) GetAutoTriangles() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAutoTriangles())
}
func (self Simple) SetBlendMode(mode classdb.AnimationNodeBlendSpace2DBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendMode(mode)
}
func (self Simple) GetBlendMode() classdb.AnimationNodeBlendSpace2DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationNodeBlendSpace2DBlendMode(Expert(self).GetBlendMode())
}
func (self Simple) SetUseSync(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseSync(enable)
}
func (self Simple) IsUsingSync() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingSync())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeBlendSpace2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds a new point that represents a [param node] at the position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddBlendPoint(node object.AnimationRootNode, pos gd.Vector2, at_index gd.Int)  {
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
func (self class) SetBlendPointNode(point gd.Int, node object.AnimationRootNode)  {
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
func (self class) GetBlendPointNode(ctx gd.Lifetime, point gd.Int) object.AnimationRootNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationRootNode
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

//go:nosplit
func (self class) AsAnimationNodeBlendSpace2D() Expert { return self[0].AsAnimationNodeBlendSpace2D() }


//go:nosplit
func (self Simple) AsAnimationNodeBlendSpace2D() Simple { return self[0].AsAnimationNodeBlendSpace2D() }


//go:nosplit
func (self class) AsAnimationRootNode() AnimationRootNode.Expert { return self[0].AsAnimationRootNode() }


//go:nosplit
func (self Simple) AsAnimationRootNode() AnimationRootNode.Simple { return self[0].AsAnimationRootNode() }


//go:nosplit
func (self class) AsAnimationNode() AnimationNode.Expert { return self[0].AsAnimationNode() }


//go:nosplit
func (self Simple) AsAnimationNode() AnimationNode.Simple { return self[0].AsAnimationNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AnimationNodeBlendSpace2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BlendMode = classdb.AnimationNodeBlendSpace2DBlendMode

const (
/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
