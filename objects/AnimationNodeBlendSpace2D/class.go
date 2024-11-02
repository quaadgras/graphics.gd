package AnimationNodeBlendSpace2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AnimationRootNode"
import "grow.graphics/gd/objects/AnimationNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A resource used by [AnimationNodeBlendTree].
[AnimationNodeBlendSpace1D] represents a virtual 2D space on which [AnimationRootNode]s are placed. Outputs the linear blend of the three adjacent animations using a [Vector2] weight. Adjacent in this context means the three [AnimationRootNode]s making up the triangle that contains the current value.
You can add vertices to the blend space with [method add_blend_point] and automatically triangulate it by setting [member auto_triangles] to [code]true[/code]. Otherwise, use [method add_triangle] and [method remove_triangle] to triangulate the blend space by hand.
*/
type Instance [1]classdb.AnimationNodeBlendSpace2D

/*
Adds a new point that represents a [param node] at the position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Instance) AddBlendPoint(node objects.AnimationRootNode, pos gd.Vector2) {
	class(self).AddBlendPoint(node, pos, gd.Int(-1))
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
func (self Instance) SetBlendPointPosition(point int, pos gd.Vector2) {
	class(self).SetBlendPointPosition(gd.Int(point), pos)
}

/*
Returns the position of the point at index [param point].
*/
func (self Instance) GetBlendPointPosition(point int) gd.Vector2 {
	return gd.Vector2(class(self).GetBlendPointPosition(gd.Int(point)))
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
func (self Instance) SetBlendPointNode(point int, node objects.AnimationRootNode) {
	class(self).SetBlendPointNode(gd.Int(point), node)
}

/*
Returns the [AnimationRootNode] referenced by the point at index [param point].
*/
func (self Instance) GetBlendPointNode(point int) objects.AnimationRootNode {
	return objects.AnimationRootNode(class(self).GetBlendPointNode(gd.Int(point)))
}

/*
Removes the point at index [param point] from the blend space.
*/
func (self Instance) RemoveBlendPoint(point int) {
	class(self).RemoveBlendPoint(gd.Int(point))
}

/*
Returns the number of points in the blend space.
*/
func (self Instance) GetBlendPointCount() int {
	return int(int(class(self).GetBlendPointCount()))
}

/*
Creates a new triangle using three points [param x], [param y], and [param z]. Triangles can overlap. You can insert the triangle at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Instance) AddTriangle(x int, y int, z int) {
	class(self).AddTriangle(gd.Int(x), gd.Int(y), gd.Int(z), gd.Int(-1))
}

/*
Returns the position of the point at index [param point] in the triangle of index [param triangle].
*/
func (self Instance) GetTrianglePoint(triangle int, point int) int {
	return int(int(class(self).GetTrianglePoint(gd.Int(triangle), gd.Int(point))))
}

/*
Removes the triangle at index [param triangle] from the blend space.
*/
func (self Instance) RemoveTriangle(triangle int) {
	class(self).RemoveTriangle(gd.Int(triangle))
}

/*
Returns the number of triangles in the blend space.
*/
func (self Instance) GetTriangleCount() int {
	return int(int(class(self).GetTriangleCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeBlendSpace2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendSpace2D"))
	return Instance{classdb.AnimationNodeBlendSpace2D(object)}
}

func (self Instance) AutoTriangles() bool {
	return bool(class(self).GetAutoTriangles())
}

func (self Instance) SetAutoTriangles(value bool) {
	class(self).SetAutoTriangles(value)
}

func (self Instance) MinSpace() gd.Vector2 {
	return gd.Vector2(class(self).GetMinSpace())
}

func (self Instance) SetMinSpace(value gd.Vector2) {
	class(self).SetMinSpace(value)
}

func (self Instance) MaxSpace() gd.Vector2 {
	return gd.Vector2(class(self).GetMaxSpace())
}

func (self Instance) SetMaxSpace(value gd.Vector2) {
	class(self).SetMaxSpace(value)
}

func (self Instance) Snap() gd.Vector2 {
	return gd.Vector2(class(self).GetSnap())
}

func (self Instance) SetSnap(value gd.Vector2) {
	class(self).SetSnap(value)
}

func (self Instance) XLabel() string {
	return string(class(self).GetXLabel().String())
}

func (self Instance) SetXLabel(value string) {
	class(self).SetXLabel(gd.NewString(value))
}

func (self Instance) YLabel() string {
	return string(class(self).GetYLabel().String())
}

func (self Instance) SetYLabel(value string) {
	class(self).SetYLabel(gd.NewString(value))
}

func (self Instance) BlendMode() classdb.AnimationNodeBlendSpace2DBlendMode {
	return classdb.AnimationNodeBlendSpace2DBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value classdb.AnimationNodeBlendSpace2DBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) Sync() bool {
	return bool(class(self).IsUsingSync())
}

func (self Instance) SetSync(value bool) {
	class(self).SetUseSync(value)
}

/*
Adds a new point that represents a [param node] at the position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddBlendPoint(node objects.AnimationRootNode, pos gd.Vector2, at_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_add_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
//go:nosplit
func (self class) SetBlendPointPosition(point gd.Int, pos gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pos)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointPosition(point gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) SetBlendPointNode(point gd.Int, node objects.AnimationRootNode) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [AnimationRootNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointNode(point gd.Int) objects.AnimationRootNode {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AnimationRootNode{classdb.AnimationRootNode(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Removes the point at index [param point] from the blend space.
*/
//go:nosplit
func (self class) RemoveBlendPoint(point gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_remove_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of points in the blend space.
*/
//go:nosplit
func (self class) GetBlendPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_blend_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new triangle using three points [param x], [param y], and [param z]. Triangles can overlap. You can insert the triangle at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddTriangle(x gd.Int, y gd.Int, z gd.Int, at_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, z)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_add_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the point at index [param point] in the triangle of index [param triangle].
*/
//go:nosplit
func (self class) GetTrianglePoint(triangle gd.Int, point gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, triangle)
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_triangle_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the triangle at index [param triangle] from the blend space.
*/
//go:nosplit
func (self class) RemoveTriangle(triangle gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, triangle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_remove_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of triangles in the blend space.
*/
//go:nosplit
func (self class) GetTriangleCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_triangle_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSpace(min_space gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, min_space)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSpace() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSpace(max_space gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, max_space)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxSpace() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap(snap gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnap() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXLabel(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_x_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_x_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetYLabel(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_y_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetYLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_y_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoTriangles(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_auto_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoTriangles() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_auto_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendMode(mode classdb.AnimationNodeBlendSpace2DBlendMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() classdb.AnimationNodeBlendSpace2DBlendMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeBlendSpace2DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseSync(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_set_use_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingSync() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace2D.Bind_is_using_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTrianglesUpdated(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("triangles_updated"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationNodeBlendSpace2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeBlendSpace2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationRootNode() AnimationRootNode.Advanced {
	return *((*AnimationRootNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationRootNode() AnimationRootNode.Instance {
	return *((*AnimationRootNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}
func init() {
	classdb.Register("AnimationNodeBlendSpace2D", func(ptr gd.Object) any { return classdb.AnimationNodeBlendSpace2D(ptr) })
}

type BlendMode = classdb.AnimationNodeBlendSpace2DBlendMode

const (
	/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
	/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
	/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
