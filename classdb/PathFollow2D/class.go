// Package PathFollow2D provides methods for working with PathFollow2D object instances.
package PathFollow2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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

/*
This node takes its parent [Path2D], and returns the coordinates of a point within it, given a distance from the first vertex.
It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be children of this node. The descendant nodes will then move accordingly when setting the [member progress] in this node.
*/
type Instance [1]gdclass.PathFollow2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPathFollow2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PathFollow2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PathFollow2D"))
	casted := Instance{*(*gdclass.PathFollow2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Progress() Float.X {
	return Float.X(Float.X(class(self).GetProgress()))
}

func (self Instance) SetProgress(value Float.X) {
	class(self).SetProgress(gd.Float(value))
}

func (self Instance) ProgressRatio() Float.X {
	return Float.X(Float.X(class(self).GetProgressRatio()))
}

func (self Instance) SetProgressRatio(value Float.X) {
	class(self).SetProgressRatio(gd.Float(value))
}

func (self Instance) HOffset() Float.X {
	return Float.X(Float.X(class(self).GetHOffset()))
}

func (self Instance) SetHOffset(value Float.X) {
	class(self).SetHOffset(gd.Float(value))
}

func (self Instance) VOffset() Float.X {
	return Float.X(Float.X(class(self).GetVOffset()))
}

func (self Instance) SetVOffset(value Float.X) {
	class(self).SetVOffset(gd.Float(value))
}

func (self Instance) Rotates() bool {
	return bool(class(self).IsRotating())
}

func (self Instance) SetRotates(value bool) {
	class(self).SetRotates(value)
}

func (self Instance) CubicInterp() bool {
	return bool(class(self).GetCubicInterpolation())
}

func (self Instance) SetCubicInterp(value bool) {
	class(self).SetCubicInterpolation(value)
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

//go:nosplit
func (self class) SetProgress(progress gd.Float) { //gd:PathFollow2D.set_progress
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgress() gd.Float { //gd:PathFollow2D.get_progress
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_get_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHOffset(h_offset gd.Float) { //gd:PathFollow2D.set_h_offset
	var frame = callframe.New()
	callframe.Arg(frame, h_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHOffset() gd.Float { //gd:PathFollow2D.get_h_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVOffset(v_offset gd.Float) { //gd:PathFollow2D.set_v_offset
	var frame = callframe.New()
	callframe.Arg(frame, v_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVOffset() gd.Float { //gd:PathFollow2D.get_v_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProgressRatio(ratio gd.Float) { //gd:PathFollow2D.set_progress_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgressRatio() gd.Float { //gd:PathFollow2D.get_progress_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_get_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotates(enabled bool) { //gd:PathFollow2D.set_rotates
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_rotates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRotating() bool { //gd:PathFollow2D.is_rotating
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_is_rotating, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCubicInterpolation(enabled bool) { //gd:PathFollow2D.set_cubic_interpolation
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCubicInterpolation() bool { //gd:PathFollow2D.get_cubic_interpolation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_get_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) { //gd:PathFollow2D.set_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool { //gd:PathFollow2D.has_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow2D.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPathFollow2D() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPathFollow2D() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("PathFollow2D", func(ptr gd.Object) any {
		return [1]gdclass.PathFollow2D{*(*gdclass.PathFollow2D)(unsafe.Pointer(&ptr))}
	})
}
