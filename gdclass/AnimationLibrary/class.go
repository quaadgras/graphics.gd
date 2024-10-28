package AnimationLibrary

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
An animation library stores a set of animations accessible through [StringName] keys, for use with [AnimationPlayer] nodes.

*/
type Go [1]classdb.AnimationLibrary

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
func (self Go) AddAnimation(name string, animation gdclass.Animation) gd.Error {
	return gd.Error(class(self).AddAnimation(gd.NewStringName(name), animation))
}

/*
Removes the [Animation] with the key [param name].
*/
func (self Go) RemoveAnimation(name string) {
	class(self).RemoveAnimation(gd.NewStringName(name))
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
func (self Go) RenameAnimation(name string, newname string) {
	class(self).RenameAnimation(gd.NewStringName(name), gd.NewStringName(newname))
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
func (self Go) HasAnimation(name string) bool {
	return bool(class(self).HasAnimation(gd.NewStringName(name)))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Go) GetAnimation(name string) gdclass.Animation {
	return gdclass.Animation(class(self).GetAnimation(gd.NewStringName(name)))
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
func (self Go) GetAnimationList() gd.Array {
	return gd.Array(class(self).GetAnimationList())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationLibrary
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationLibrary"))
	return Go{classdb.AnimationLibrary(object)}
}

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
//go:nosplit
func (self class) AddAnimation(name gd.StringName, animation gdclass.Animation) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(animation[0])[0])
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
func (self class) RemoveAnimation(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(name gd.StringName, newname gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(newname))
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
	callframe.Arg(frame, discreet.Get(name))
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
func (self class) GetAnimation(name gd.StringName) gdclass.Animation {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Animation{classdb.Animation(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnAnimationAdded(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_added"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationRemoved(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_removed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationRenamed(cb func(name string, to_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_renamed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationChanged(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsAnimationLibrary() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationLibrary() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("AnimationLibrary", func(ptr gd.Object) any { return classdb.AnimationLibrary(ptr) })}
