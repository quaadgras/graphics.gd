package SkeletonModification2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource provides an interface that can be expanded so code that operates on [Bone2D] nodes in a [Skeleton2D] can be mixed and matched together to create complex interactions.
This is used to provide Godot with a flexible and powerful Inverse Kinematics solution that can be adapted for many different uses.
	// SkeletonModification2D methods that can be overridden by a [Class] that extends it.
	type SkeletonModification2D interface {
		//Executes the given modification. This is where the modification performs whatever function it is designed to do.
		Execute(delta gd.Float) 
		//Called when the modification is setup. This is where the modification performs initialization.
		SetupModification(modification_stack object.SkeletonModificationStack2D) 
		//Used for drawing [b]editor-only[/b] modification gizmos. This function will only be called in the Godot editor and can be overridden to draw custom gizmos.
		//[b]Note:[/b] You will need to use the Skeleton2D from [method SkeletonModificationStack2D.get_skeleton] and it's draw functions, as the [SkeletonModification2D] resource cannot draw on its own.
		DrawEditorGizmo() 
	}

*/
type Simple [1]classdb.SkeletonModification2D
func (self Simple) SetEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnabled(enabled)
}
func (self Simple) GetEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnabled())
}
func (self Simple) GetModificationStack() [1]classdb.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkeletonModificationStack2D(Expert(self).GetModificationStack(gc))
}
func (self Simple) SetIsSetup(is_setup bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIsSetup(is_setup)
}
func (self Simple) GetIsSetup() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetIsSetup())
}
func (self Simple) SetExecutionMode(execution_mode int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExecutionMode(gd.Int(execution_mode))
}
func (self Simple) GetExecutionMode() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetExecutionMode()))
}
func (self Simple) ClampAngle(angle float64, min float64, max float64, invert bool) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ClampAngle(gd.Float(angle), gd.Float(min), gd.Float(max), invert)))
}
func (self Simple) SetEditorDrawGizmo(draw_gizmo bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditorDrawGizmo(draw_gizmo)
}
func (self Simple) GetEditorDrawGizmo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEditorDrawGizmo())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (class) _setup_modification(impl func(ptr unsafe.Pointer, modification_stack object.SkeletonModificationStack2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var modification_stack object.SkeletonModificationStack2D
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
func (self class) GetModificationStack(ctx gd.Lifetime) object.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkeletonModificationStack2D
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

//go:nosplit
func (self class) AsSkeletonModification2D() Expert { return self[0].AsSkeletonModification2D() }


//go:nosplit
func (self Simple) AsSkeletonModification2D() Simple { return self[0].AsSkeletonModification2D() }


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
	case "_execute": return reflect.ValueOf(self._execute);
	case "_setup_modification": return reflect.ValueOf(self._setup_modification);
	case "_draw_editor_gizmo": return reflect.ValueOf(self._draw_editor_gizmo);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SkeletonModification2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
