// Package SkeletonModificationStack2D provides methods for working with SkeletonModificationStack2D object instances.
package SkeletonModificationStack2D

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
This resource is used by the Skeleton and holds a stack of [SkeletonModification2D]s.
This controls the order of the modifications and how they are applied. Modification order is especially important for full-body IK setups, as you need to execute the modifications in the correct order to get the desired results. For example, you want to execute a modification on the spine [i]before[/i] the arms on a humanoid skeleton.
This resource also controls how strongly all of the modifications are applied to the [Skeleton2D].
*/
type Instance [1]gdclass.SkeletonModificationStack2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModificationStack2D() Instance
}

/*
Sets up the modification stack so it can execute. This function should be called by [Skeleton2D] and shouldn't be manually called unless you know what you are doing.
*/
func (self Instance) Setup() {
	class(self).Setup()
}

/*
Executes all of the [SkeletonModification2D]s in the stack that use the same execution mode as the passed-in [param execution_mode], starting from index [code]0[/code] to [member modification_count].
[b]Note:[/b] The order of the modifications can matter depending on the modifications. For example, modifications on a spine should operate before modifications on the arms in order to get proper results.
*/
func (self Instance) Execute(delta Float.X, execution_mode int) {
	class(self).Execute(gd.Float(delta), gd.Int(execution_mode))
}

/*
Enables all [SkeletonModification2D]s in the stack.
*/
func (self Instance) EnableAllModifications(enabled bool) {
	class(self).EnableAllModifications(enabled)
}

/*
Returns the [SkeletonModification2D] at the passed-in index, [param mod_idx].
*/
func (self Instance) GetModification(mod_idx int) [1]gdclass.SkeletonModification2D {
	return [1]gdclass.SkeletonModification2D(class(self).GetModification(gd.Int(mod_idx)))
}

/*
Adds the passed-in [SkeletonModification2D] to the stack.
*/
func (self Instance) AddModification(modification [1]gdclass.SkeletonModification2D) {
	class(self).AddModification(modification)
}

/*
Deletes the [SkeletonModification2D] at the index position [param mod_idx], if it exists.
*/
func (self Instance) DeleteModification(mod_idx int) {
	class(self).DeleteModification(gd.Int(mod_idx))
}

/*
Sets the modification at [param mod_idx] to the passed-in modification, [param modification].
*/
func (self Instance) SetModification(mod_idx int, modification [1]gdclass.SkeletonModification2D) {
	class(self).SetModification(gd.Int(mod_idx), modification)
}

/*
Returns a boolean that indicates whether the modification stack is setup and can execute.
*/
func (self Instance) GetIsSetup() bool {
	return bool(class(self).GetIsSetup())
}

/*
Returns the [Skeleton2D] node that the SkeletonModificationStack2D is bound to.
*/
func (self Instance) GetSkeleton() [1]gdclass.Skeleton2D {
	return [1]gdclass.Skeleton2D(class(self).GetSkeleton())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModificationStack2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModificationStack2D"))
	casted := Instance{*(*gdclass.SkeletonModificationStack2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Enabled() bool {
	return bool(class(self).GetEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) Strength() Float.X {
	return Float.X(Float.X(class(self).GetStrength()))
}

func (self Instance) SetStrength(value Float.X) {
	class(self).SetStrength(gd.Float(value))
}

func (self Instance) ModificationCount() int {
	return int(int(class(self).GetModificationCount()))
}

func (self Instance) SetModificationCount(value int) {
	class(self).SetModificationCount(gd.Int(value))
}

/*
Sets up the modification stack so it can execute. This function should be called by [Skeleton2D] and shouldn't be manually called unless you know what you are doing.
*/
//go:nosplit
func (self class) Setup() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Executes all of the [SkeletonModification2D]s in the stack that use the same execution mode as the passed-in [param execution_mode], starting from index [code]0[/code] to [member modification_count].
[b]Note:[/b] The order of the modifications can matter depending on the modifications. For example, modifications on a spine should operate before modifications on the arms in order to get proper results.
*/
//go:nosplit
func (self class) Execute(delta gd.Float, execution_mode gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables all [SkeletonModification2D]s in the stack.
*/
//go:nosplit
func (self class) EnableAllModifications(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_enable_all_modifications, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [SkeletonModification2D] at the passed-in index, [param mod_idx].
*/
//go:nosplit
func (self class) GetModification(mod_idx gd.Int) [1]gdclass.SkeletonModification2D {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_modification, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkeletonModification2D{gd.PointerWithOwnershipTransferredToGo[gdclass.SkeletonModification2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds the passed-in [SkeletonModification2D] to the stack.
*/
//go:nosplit
func (self class) AddModification(modification [1]gdclass.SkeletonModification2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(modification[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_add_modification, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Deletes the [SkeletonModification2D] at the index position [param mod_idx], if it exists.
*/
//go:nosplit
func (self class) DeleteModification(mod_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_delete_modification, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the modification at [param mod_idx] to the passed-in modification, [param modification].
*/
//go:nosplit
func (self class) SetModification(mod_idx gd.Int, modification [1]gdclass.SkeletonModification2D) {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	callframe.Arg(frame, pointers.Get(modification[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_modification, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetModificationCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_modification_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModificationCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_modification_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a boolean that indicates whether the modification stack is setup and can execute.
*/
//go:nosplit
func (self class) GetIsSetup() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_is_setup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Skeleton2D] node that the SkeletonModificationStack2D is bound to.
*/
//go:nosplit
func (self class) GetSkeleton() [1]gdclass.Skeleton2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skeleton2D{gd.PointerMustAssertInstanceID[gdclass.Skeleton2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSkeletonModificationStack2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModificationStack2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("SkeletonModificationStack2D", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModificationStack2D{*(*gdclass.SkeletonModificationStack2D)(unsafe.Pointer(&ptr))}
	})
}
