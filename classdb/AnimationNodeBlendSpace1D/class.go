// Package AnimationNodeBlendSpace1D provides methods for working with AnimationNodeBlendSpace1D object instances.
package AnimationNodeBlendSpace1D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/AnimationRootNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A resource used by [AnimationNodeBlendTree].
[AnimationNodeBlendSpace1D] represents a virtual axis on which any type of [AnimationRootNode]s can be added using [method add_blend_point]. Outputs the linear blend of the two [AnimationRootNode]s adjacent to the current value.
You can set the extents of the axis with [member min_space] and [member max_space].
*/
type Instance [1]gdclass.AnimationNodeBlendSpace1D
type Expanded [1]gdclass.AnimationNodeBlendSpace1D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeBlendSpace1D() Instance
}

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Instance) AddBlendPoint(node [1]gdclass.AnimationRootNode, pos Float.X) { //gd:AnimationNodeBlendSpace1D.add_blend_point
	Advanced(self).AddBlendPoint(node, float64(pos), int64(-1))
}

/*
Adds a new point that represents a [param node] on the virtual axis at a given position set by [param pos]. You can insert it at a specific index using the [param at_index] argument. If you use the default value for [param at_index], the point is inserted at the end of the blend points array.
*/
func (self Expanded) AddBlendPoint(node [1]gdclass.AnimationRootNode, pos Float.X, at_index int) { //gd:AnimationNodeBlendSpace1D.add_blend_point
	Advanced(self).AddBlendPoint(node, float64(pos), int64(at_index))
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
func (self Instance) SetBlendPointPosition(point int, pos Float.X) { //gd:AnimationNodeBlendSpace1D.set_blend_point_position
	Advanced(self).SetBlendPointPosition(int64(point), float64(pos))
}

/*
Returns the position of the point at index [param point].
*/
func (self Instance) GetBlendPointPosition(point int) Float.X { //gd:AnimationNodeBlendSpace1D.get_blend_point_position
	return Float.X(Float.X(Advanced(self).GetBlendPointPosition(int64(point))))
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
func (self Instance) SetBlendPointNode(point int, node [1]gdclass.AnimationRootNode) { //gd:AnimationNodeBlendSpace1D.set_blend_point_node
	Advanced(self).SetBlendPointNode(int64(point), node)
}

/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
func (self Instance) GetBlendPointNode(point int) [1]gdclass.AnimationRootNode { //gd:AnimationNodeBlendSpace1D.get_blend_point_node
	return [1]gdclass.AnimationRootNode(Advanced(self).GetBlendPointNode(int64(point)))
}

/*
Removes the point at index [param point] from the blend axis.
*/
func (self Instance) RemoveBlendPoint(point int) { //gd:AnimationNodeBlendSpace1D.remove_blend_point
	Advanced(self).RemoveBlendPoint(int64(point))
}

/*
Returns the number of points on the blend axis.
*/
func (self Instance) GetBlendPointCount() int { //gd:AnimationNodeBlendSpace1D.get_blend_point_count
	return int(int(Advanced(self).GetBlendPointCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeBlendSpace1D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendSpace1D"))
	casted := Instance{*(*gdclass.AnimationNodeBlendSpace1D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MinSpace() Float.X {
	return Float.X(Float.X(class(self).GetMinSpace()))
}

func (self Instance) SetMinSpace(value Float.X) {
	class(self).SetMinSpace(float64(value))
}

func (self Instance) MaxSpace() Float.X {
	return Float.X(Float.X(class(self).GetMaxSpace()))
}

func (self Instance) SetMaxSpace(value Float.X) {
	class(self).SetMaxSpace(float64(value))
}

func (self Instance) Snap() Float.X {
	return Float.X(Float.X(class(self).GetSnap()))
}

func (self Instance) SetSnap(value Float.X) {
	class(self).SetSnap(float64(value))
}

func (self Instance) ValueLabel() string {
	return string(class(self).GetValueLabel().String())
}

func (self Instance) SetValueLabel(value string) {
	class(self).SetValueLabel(String.New(value))
}

func (self Instance) BlendMode() gdclass.AnimationNodeBlendSpace1DBlendMode {
	return gdclass.AnimationNodeBlendSpace1DBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value gdclass.AnimationNodeBlendSpace1DBlendMode) {
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
func (self class) AddBlendPoint(node [1]gdclass.AnimationRootNode, pos float64, at_index int64) { //gd:AnimationNodeBlendSpace1D.add_blend_point
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, at_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_add_blend_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Updates the position of the point at index [param point] on the blend axis.
*/
//go:nosplit
func (self class) SetBlendPointPosition(point int64, pos float64) { //gd:AnimationNodeBlendSpace1D.set_blend_point_position
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointPosition(point int64) float64 { //gd:AnimationNodeBlendSpace1D.get_blend_point_position
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) SetBlendPointNode(point int64, node [1]gdclass.AnimationRootNode) { //gd:AnimationNodeBlendSpace1D.set_blend_point_node
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [AnimationNode] referenced by the point at index [param point].
*/
//go:nosplit
func (self class) GetBlendPointNode(point int64) [1]gdclass.AnimationRootNode { //gd:AnimationNodeBlendSpace1D.get_blend_point_node
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationRootNode{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationRootNode](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes the point at index [param point] from the blend axis.
*/
//go:nosplit
func (self class) RemoveBlendPoint(point int64) { //gd:AnimationNodeBlendSpace1D.remove_blend_point
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_remove_blend_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of points on the blend axis.
*/
//go:nosplit
func (self class) GetBlendPointCount() int64 { //gd:AnimationNodeBlendSpace1D.get_blend_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSpace(min_space float64) { //gd:AnimationNodeBlendSpace1D.set_min_space
	var frame = callframe.New()
	callframe.Arg(frame, min_space)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_min_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSpace() float64 { //gd:AnimationNodeBlendSpace1D.get_min_space
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_min_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSpace(max_space float64) { //gd:AnimationNodeBlendSpace1D.set_max_space
	var frame = callframe.New()
	callframe.Arg(frame, max_space)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_max_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxSpace() float64 { //gd:AnimationNodeBlendSpace1D.get_max_space
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_max_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap(snap float64) { //gd:AnimationNodeBlendSpace1D.set_snap
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnap() float64 { //gd:AnimationNodeBlendSpace1D.get_snap
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetValueLabel(text String.Readable) { //gd:AnimationNodeBlendSpace1D.set_value_label
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_value_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetValueLabel() String.Readable { //gd:AnimationNodeBlendSpace1D.get_value_label
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_value_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendMode(mode gdclass.AnimationNodeBlendSpace1DBlendMode) { //gd:AnimationNodeBlendSpace1D.set_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() gdclass.AnimationNodeBlendSpace1DBlendMode { //gd:AnimationNodeBlendSpace1D.get_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationNodeBlendSpace1DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseSync(enable bool) { //gd:AnimationNodeBlendSpace1D.set_use_sync
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_set_use_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingSync() bool { //gd:AnimationNodeBlendSpace1D.is_using_sync
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendSpace1D.Bind_is_using_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationRootNode.Advanced(self.AsAnimationRootNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationRootNode.Instance(self.AsAnimationRootNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeBlendSpace1D", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeBlendSpace1D{*(*gdclass.AnimationNodeBlendSpace1D)(unsafe.Pointer(&ptr))}
	})
}

type BlendMode = gdclass.AnimationNodeBlendSpace1DBlendMode //gd:AnimationNodeBlendSpace1D.BlendMode

const (
	/*The interpolation between animations is linear.*/
	BlendModeInterpolated BlendMode = 0
	/*The blend space plays the animation of the animation node which blending position is closest to. Useful for frame-by-frame 2D animations.*/
	BlendModeDiscrete BlendMode = 1
	/*Similar to [constant BLEND_MODE_DISCRETE], but starts the new animation at the last animation's playback position.*/
	BlendModeDiscreteCarry BlendMode = 2
)
