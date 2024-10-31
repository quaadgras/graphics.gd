package PathFollow3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This node takes its parent [Path3D], and returns the coordinates of a point within it, given a distance from the first vertex.
It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be children of this node. The descendant nodes will then move accordingly when setting the [member progress] in this node.
*/
type Instance [1]classdb.PathFollow3D

/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
func (self Instance) CorrectPosture(transform gd.Transform3D, rotation_mode classdb.PathFollow3DRotationMode) gd.Transform3D {
	return gd.Transform3D(class(self).CorrectPosture(transform, rotation_mode))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PathFollow3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PathFollow3D"))
	return Instance{classdb.PathFollow3D(object)}
}

func (self Instance) Progress() float64 {
	return float64(float64(class(self).GetProgress()))
}

func (self Instance) SetProgress(value float64) {
	class(self).SetProgress(gd.Float(value))
}

func (self Instance) ProgressRatio() float64 {
	return float64(float64(class(self).GetProgressRatio()))
}

func (self Instance) SetProgressRatio(value float64) {
	class(self).SetProgressRatio(gd.Float(value))
}

func (self Instance) HOffset() float64 {
	return float64(float64(class(self).GetHOffset()))
}

func (self Instance) SetHOffset(value float64) {
	class(self).SetHOffset(gd.Float(value))
}

func (self Instance) VOffset() float64 {
	return float64(float64(class(self).GetVOffset()))
}

func (self Instance) SetVOffset(value float64) {
	class(self).SetVOffset(gd.Float(value))
}

func (self Instance) RotationMode() classdb.PathFollow3DRotationMode {
	return classdb.PathFollow3DRotationMode(class(self).GetRotationMode())
}

func (self Instance) SetRotationMode(value classdb.PathFollow3DRotationMode) {
	class(self).SetRotationMode(value)
}

func (self Instance) UseModelFront() bool {
	return bool(class(self).IsUsingModelFront())
}

func (self Instance) SetUseModelFront(value bool) {
	class(self).SetUseModelFront(value)
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

func (self Instance) TiltEnabled() bool {
	return bool(class(self).IsTiltEnabled())
}

func (self Instance) SetTiltEnabled(value bool) {
	class(self).SetTiltEnabled(value)
}

//go:nosplit
func (self class) SetProgress(progress gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgress() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHOffset(h_offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, h_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVOffset(v_offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, v_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProgressRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgressRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationMode(rotation_mode classdb.PathFollow3DRotationMode) {
	var frame = callframe.New()
	callframe.Arg(frame, rotation_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationMode() classdb.PathFollow3DRotationMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PathFollow3DRotationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCubicInterpolation(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCubicInterpolation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseModelFront(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_use_model_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingModelFront() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_is_using_model_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) {
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTiltEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsTiltEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_is_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
//go:nosplit
func (self class) CorrectPosture(transform gd.Transform3D, rotation_mode classdb.PathFollow3DRotationMode) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	callframe.Arg(frame, rotation_mode)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_correct_posture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPathFollow3D() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPathFollow3D() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	classdb.Register("PathFollow3D", func(ptr gd.Object) any { return classdb.PathFollow3D(ptr) })
}

type RotationMode = classdb.PathFollow3DRotationMode

const (
	/*Forbids the PathFollow3D to rotate.*/
	RotationNone RotationMode = 0
	/*Allows the PathFollow3D to rotate in the Y axis only.*/
	RotationY RotationMode = 1
	/*Allows the PathFollow3D to rotate in both the X, and Y axes.*/
	RotationXy RotationMode = 2
	/*Allows the PathFollow3D to rotate in any axis.*/
	RotationXyz RotationMode = 3
	/*Uses the up vector information in a [Curve3D] to enforce orientation. This rotation mode requires the [Path3D]'s [member Curve3D.up_vector_enabled] property to be set to [code]true[/code].*/
	RotationOriented RotationMode = 4
)
