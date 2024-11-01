package AnimationLibrary

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
An animation library stores a set of animations accessible through [StringName] keys, for use with [AnimationPlayer] nodes.
*/
type Instance [1]classdb.AnimationLibrary

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
func (self Instance) AddAnimation(name string, animation objects.Animation) gd.Error {
	return gd.Error(class(self).AddAnimation(gd.NewStringName(name), animation))
}

/*
Removes the [Animation] with the key [param name].
*/
func (self Instance) RemoveAnimation(name string) {
	class(self).RemoveAnimation(gd.NewStringName(name))
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
func (self Instance) RenameAnimation(name string, newname string) {
	class(self).RenameAnimation(gd.NewStringName(name), gd.NewStringName(newname))
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
func (self Instance) HasAnimation(name string) bool {
	return bool(class(self).HasAnimation(gd.NewStringName(name)))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Instance) GetAnimation(name string) objects.Animation {
	return objects.Animation(class(self).GetAnimation(gd.NewStringName(name)))
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
func (self Instance) GetAnimationList() gd.Array {
	return gd.Array(class(self).GetAnimationList())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationLibrary

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationLibrary"))
	return Instance{classdb.AnimationLibrary(object)}
}

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
//go:nosplit
func (self class) AddAnimation(name gd.StringName, animation objects.Animation) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_add_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the [Animation] with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimation(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(name gd.StringName, newname gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(newname))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_rename_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
//go:nosplit
func (self class) HasAnimation(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
//go:nosplit
func (self class) GetAnimation(name gd.StringName) objects.Animation {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Animation{classdb.Animation(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
//go:nosplit
func (self class) GetAnimationList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnAnimationAdded(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_added"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationRemoved(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationRenamed(cb func(name string, to_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_renamed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationChanged(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationLibrary() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationLibrary() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("AnimationLibrary", func(ptr gd.Object) any { return classdb.AnimationLibrary(ptr) })
}
