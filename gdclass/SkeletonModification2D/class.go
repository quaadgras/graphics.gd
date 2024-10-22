package SkeletonModification2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource provides an interface that can be expanded so code that operates on [Bone2D] nodes in a [Skeleton2D] can be mixed and matched together to create complex interactions.
This is used to provide Godot with a flexible and powerful Inverse Kinematics solution that can be adapted for many different uses.
	// SkeletonModification2D methods that can be overridden by a [Class] that extends it.
	type SkeletonModification2D interface {
		//Executes the given modification. This is where the modification performs whatever function it is designed to do.
		Execute(delta float64) 
		//Called when the modification is setup. This is where the modification performs initialization.
		SetupModification(modification_stack gdclass.SkeletonModificationStack2D) 
		//Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
		//[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
		DrawEditorGizmo() 
	}

*/
type Go [1]classdb.SkeletonModification2D

/*
Executes the given modification. This is where the modification performs whatever function it is designed to do.
*/
func (Go) _execute(impl func(ptr unsafe.Pointer, delta float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(delta))
		gc.End()
	}
}

/*
Called when the modification is setup. This is where the modification performs initialization.
*/
func (Go) _setup_modification(impl func(ptr unsafe.Pointer, modification_stack gdclass.SkeletonModificationStack2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var modification_stack gdclass.SkeletonModificationStack2D
		modification_stack[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, modification_stack)
		gc.End()
	}
}

/*
Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
*/
func (Go) _draw_editor_gizmo(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Returns the [SkeletonModificationStack2D] that this modification is bound to. Through the modification stack, you can access the Skeleton2D the modification is operating on.
*/
func (self Go) GetModificationStack() gdclass.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.SkeletonModificationStack2D(class(self).GetModificationStack(gc))
}

/*
Manually allows you to set the setup state of the modification. This function should only rarely be used, as the [SkeletonModificationStack2D] the modification is bound to should handle setting the modification up.
*/
func (self Go) SetIsSetup(is_setup bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIsSetup(is_setup)
}

/*
Returns whether this modification has been successfully setup or not.
*/
func (self Go) GetIsSetup() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetIsSetup())
}

/*
Takes an angle and clamps it so it is within the passed-in [param min] and [param max] range. [param invert] will inversely clamp the angle, clamping it to the range outside of the given bounds.
*/
func (self Go) ClampAngle(angle float64, min float64, max float64, invert bool) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ClampAngle(gd.Float(angle), gd.Float(min), gd.Float(max), invert)))
}

/*
Sets whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
func (self Go) SetEditorDrawGizmo(draw_gizmo bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditorDrawGizmo(draw_gizmo)
}

/*
Returns whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
func (self Go) GetEditorDrawGizmo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetEditorDrawGizmo())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Enabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnabled())
}

func (self Go) SetEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnabled(value)
}

func (self Go) ExecutionMode() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetExecutionMode()))
}

func (self Go) SetExecutionMode(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExecutionMode(gd.Int(value))
}

/*
Executes the given modification. This is where the modification performs whatever function it is designed to do.
*/
func (class) _execute(impl func(ptr unsafe.Pointer, delta gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, delta)
		ctx.End()
	}
}

/*
Called when the modification is setup. This is where the modification performs initialization.
*/
func (class) _setup_modification(impl func(ptr unsafe.Pointer, modification_stack gdclass.SkeletonModificationStack2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var modification_stack gdclass.SkeletonModificationStack2D
		modification_stack[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, modification_stack)
		ctx.End()
	}
}

/*
Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
*/
func (class) _draw_editor_gizmo(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [SkeletonModificationStack2D] that this modification is bound to. Through the modification stack, you can access the Skeleton2D the modification is operating on.
*/
//go:nosplit
func (self class) GetModificationStack(ctx gd.Lifetime) gdclass.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SkeletonModificationStack2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Manually allows you to set the setup state of the modification. This function should only rarely be used, as the [SkeletonModificationStack2D] the modification is bound to should handle setting the modification up.
*/
//go:nosplit
func (self class) SetIsSetup(is_setup bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, is_setup)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_set_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether this modification has been successfully setup or not.
*/
//go:nosplit
func (self class) GetIsSetup() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExecutionMode(execution_mode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_set_execution_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExecutionMode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_execution_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Takes an angle and clamps it so it is within the passed-in [param min] and [param max] range. [param invert] will inversely clamp the angle, clamping it to the range outside of the given bounds.
*/
//go:nosplit
func (self class) ClampAngle(angle gd.Float, min gd.Float, max gd.Float, invert bool) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, invert)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_clamp_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
//go:nosplit
func (self class) SetEditorDrawGizmo(draw_gizmo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_gizmo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_set_editor_draw_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether this modification will call [method _draw_editor_gizmo] in the Godot editor to draw modification-specific gizmos.
*/
//go:nosplit
func (self class) GetEditorDrawGizmo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_editor_draw_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_execute": return reflect.ValueOf(self._execute);
	case "_setup_modification": return reflect.ValueOf(self._setup_modification);
	case "_draw_editor_gizmo": return reflect.ValueOf(self._draw_editor_gizmo);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_execute": return reflect.ValueOf(self._execute);
	case "_setup_modification": return reflect.ValueOf(self._setup_modification);
	case "_draw_editor_gizmo": return reflect.ValueOf(self._draw_editor_gizmo);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("SkeletonModification2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
