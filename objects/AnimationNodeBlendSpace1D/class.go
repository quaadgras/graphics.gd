package AnimationNodeBlendSpace1D

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
[AnimationNodeBlendSpace1D] represents a virtual axis on which any type of [AnimationRootNode]s can be added using [method add_blend_point]. Outputs the linear blend of the two [AnimationRootNode]s adjacent to the current value.
You can set the extents of the axis with [member min_space] and [member max_space].
*/
type Instance [1]classdb.AnimationNodeBlendSpace1D

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Instance) AddBlendPoint(node objects.AnimationRootNode, pos float64) {
	class(self).AddBlendPoint(node, gd.Float(pos), gd.Int(-1))
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
func (self Instance) SetBlendPointPosition(point int, pos float64) {
	class(self).SetBlendPointPosition(gd.Int(point), gd.Float(pos))
}

/*
Returns the position of the point at index [param point].
*/
func (self Instance) GetBlendPointPosition(point int) float64 {
	return float64(float64(class(self).GetBlendPointPosition(gd.Int(point))))
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
func (self Instance) SetBlendPointNode(point int, node objects.AnimationRootNode) {
	class(self).SetBlendPointNode(gd.Int(point), node)
}

/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
func (self Instance) GetBlendPointNode(point int) objects.AnimationRootNode {
	return objects.AnimationRootNode(class(self).GetBlendPointNode(gd.Int(point)))
}

/*
Removes the point at index [param point] from the blend axis.
*/
func (self Instance) RemoveBlendPoint(point int) {
	class(self).RemoveBlendPoint(gd.Int(point))
}

/*
Returns the number of points on the blend axis.
*/
func (self Instance) GetBlendPointCount() int {
	return int(int(class(self).GetBlendPointCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeBlendSpace1D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendSpace1D"))
	return Instance{classdb.AnimationNodeBlendSpace1D(object)}
}

func (self Instance) MinSpace() float64 {
	return float64(float64(class(self).GetMinSpace()))
}

func (self Instance) SetMinSpace(value float64) {
	class(self).SetMinSpace(gd.Float(value))
}

func (self Instance) MaxSpace() float64 {
	return float64(float64(class(self).GetMaxSpace()))
}

func (self Instance) SetMaxSpace(value float64) {
	class(self).SetMaxSpace(gd.Float(value))
}

func (self Instance) Snap() float64 {
	return float64(float64(class(self).GetSnap()))
}

func (self Instance) SetSnap(value float64) {
	class(self).SetSnap(gd.Float(value))
}

func (self Instance) ValueLabel() string {
	return string(class(self).GetValueLabel().String())
}

func (self Instance) SetValueLabel(value string) {
	class(self).SetValueLabel(gd.NewString(value))
}

func (self Instance) BlendMode() classdb.AnimationNodeBlendSpace1DBlendMode {
	return classdb.AnimationNodeBlendSpace1DBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value classdb.AnimationNodeBlendSpace1DBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) Sync() bool {
	return bool(class(self).IsUsingSync())
}

func (self Instance) SetSync(value bool) {
	class(self).SetUseSync(value)
}

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
//go:nosplit
func (self class) AddBlendPoint(node objects.AnimationRootNode, pos gd.Float, at_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, at_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_add_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
//go:nosplit
func (self class) SetBlendPointPosition(point gd.Int, pos gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pos)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointPosition(point gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointNode(point gd.Int) objects.AnimationRootNode {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AnimationRootNode{classdb.AnimationRootNode(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Removes the point at index [param point] from the blend axis.
*/
//go:nosplit
func (self class) RemoveBlendPoint(point gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_remove_blend_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of points on the blend axis.
*/
//go:nosplit
func (self class) GetBlendPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSpace(min_space gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, min_space)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSpace() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_min_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSpace(max_space gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, max_space)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxSpace() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_max_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap(snap gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnap() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetValueLabel(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_value_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetValueLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_value_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendMode(mode classdb.AnimationNodeBlendSpace1DBlendMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() classdb.AnimationNodeBlendSpace1DBlendMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeBlendSpace1DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseSync(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_use_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingSync() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_is_using_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeBlendSpace1D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeBlendSpace1D() Instance {
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
	classdb.Register("AnimationNodeBlendSpace1D", func(ptr gd.Object) any { return classdb.AnimationNodeBlendSpace1D(ptr) })
}

type BlendMode = classdb.AnimationNodeBlendSpace1DBlendMode

const (
	/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
	/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
	/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
