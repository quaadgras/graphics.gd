package SkeletonModification2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This resource provides an interface that can be expanded so code that operates on [Bone2D] nodes in a [Skeleton2D] can be mixed and matched together to create complex interactions.
This is used to provide Godot with a flexible and powerful Inverse Kinematics solution that can be adapted for many different uses.

	// SkeletonModification2D methods that can be overridden by a [Class] that extends it.
	type SkeletonModification2D interface {
		//Executes the given modification. This is where the modification performs whatever function it is designed to do.
		Execute(delta float64)
		//Called when the modification is setup. This is where the modification performs initialization.
		SetupModification(modification_stack objects.SkeletonModificationStack2D)
		//Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
		//[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
		DrawEditorGizmo()
	}
*/
type Instance [1]classdb.SkeletonModification2D

/*
Executes the given modification. This is where the modification performs whatever function it is designed to do.
*/
func (Instance) _execute(impl func(ptr unsafe.Pointer, delta float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, float64(delta))
	}
}

/*
Called when the modification is setup. This is where the modification performs initialization.
*/
func (Instance) _setup_modification(impl func(ptr unsafe.Pointer, modification_stack objects.SkeletonModificationStack2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var modification_stack = objects.SkeletonModificationStack2D{pointers.New[classdb.SkeletonModificationStack2D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(modification_stack[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, modification_stack)
	}
}

/*
Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
*/
func (Instance) _draw_editor_gizmo(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns the [SkeletonModificationStack2D] that this modification is bound to. Through the modification stack, you can access the Skeleton2D the modification is operating on.
*/
func (self Instance) GetModificationStack() objects.SkeletonModificationStack2D {
	return objects.SkeletonModificationStack2D(class(self).GetModificationStack())
}

/*
Manually allows you to set the setup state of the modification. This function should only rarely be used, as the [SkeletonModificationStack2D] the modification is bound to should handle setting the modification up.
*/
func (self Instance) SetIsSetup(is_setup bool) {
	class(self).SetIsSetup(is_setup)
}

/*
Returns whether this modification has been successfully setup or not.
*/
func (self Instance) GetIsSetup() bool {
	return bool(class(self).GetIsSetup())
}

/*
Takes an angle and clamps it so it is within the passed-in [param min] and [param max] range. [param invert] will inversely clamp the angle, clamping it to the range outside of the given bounds.
*/
func (self Instance) ClampAngle(angle float64, min float64, max float64, invert bool) float64 {
	return float64(float64(class(self).ClampAngle(gd.Float(angle), gd.Float(min), gd.Float(max), invert)))
}

/*
Sets whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
func (self Instance) SetEditorDrawGizmo(draw_gizmo bool) {
	class(self).SetEditorDrawGizmo(draw_gizmo)
}

/*
Returns whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
func (self Instance) GetEditorDrawGizmo() bool {
	return bool(class(self).GetEditorDrawGizmo())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SkeletonModification2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2D"))
	return Instance{classdb.SkeletonModification2D(object)}
}

func (self Instance) Enabled() bool {
	return bool(class(self).GetEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) ExecutionMode() int {
	return int(int(class(self).GetExecutionMode()))
}

func (self Instance) SetExecutionMode(value int) {
	class(self).SetExecutionMode(gd.Int(value))
}

/*
Executes the given modification. This is where the modification performs whatever function it is designed to do.
*/
func (class) _execute(impl func(ptr unsafe.Pointer, delta gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, delta)
	}
}

/*
Called when the modification is setup. This is where the modification performs initialization.
*/
func (class) _setup_modification(impl func(ptr unsafe.Pointer, modification_stack objects.SkeletonModificationStack2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var modification_stack = objects.SkeletonModificationStack2D{pointers.New[classdb.SkeletonModificationStack2D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(modification_stack[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, modification_stack)
	}
}

/*
Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
*/
func (class) _draw_editor_gizmo(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [SkeletonModificationStack2D] that this modification is bound to. Through the modification stack, you can access the Skeleton2D the modification is operating on.
*/
//go:nosplit
func (self class) GetModificationStack() objects.SkeletonModificationStack2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.SkeletonModificationStack2D{classdb.SkeletonModificationStack2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Manually allows you to set the setup state of the modification. This function should only rarely be used, as the [SkeletonModificationStack2D] the modification is bound to should handle setting the modification up.
*/
//go:nosplit
func (self class) SetIsSetup(is_setup bool) {
	var frame = callframe.New()
	callframe.Arg(frame, is_setup)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_set_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether this modification has been successfully setup or not.
*/
//go:nosplit
func (self class) GetIsSetup() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_get_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExecutionMode(execution_mode gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_set_execution_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExecutionMode() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_get_execution_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Takes an angle and clamps it so it is within the passed-in [param min] and [param max] range. [param invert] will inversely clamp the angle, clamping it to the range outside of the given bounds.
*/
//go:nosplit
func (self class) ClampAngle(angle gd.Float, min gd.Float, max gd.Float, invert bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, invert)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_clamp_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
//go:nosplit
func (self class) SetEditorDrawGizmo(draw_gizmo bool) {
	var frame = callframe.New()
	callframe.Arg(frame, draw_gizmo)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_set_editor_draw_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
//go:nosplit
func (self class) GetEditorDrawGizmo() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2D.Bind_get_editor_draw_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeletonModification2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
	case "_execute":
		return reflect.ValueOf(self._execute)
	case "_setup_modification":
		return reflect.ValueOf(self._setup_modification)
	case "_draw_editor_gizmo":
		return reflect.ValueOf(self._draw_editor_gizmo)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_execute":
		return reflect.ValueOf(self._execute)
	case "_setup_modification":
		return reflect.ValueOf(self._setup_modification)
	case "_draw_editor_gizmo":
		return reflect.ValueOf(self._draw_editor_gizmo)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("SkeletonModification2D", func(ptr gd.Object) any { return classdb.SkeletonModification2D(ptr) })
}
