package AnimationLibrary

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
An animation library stores a set of animations accessible through [StringName] keys, for use with [AnimationPlayer] nodes.

*/
type Go [1]classdb.AnimationLibrary

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
func (self Go) AddAnimation(name string, animation gdclass.Animation) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AddAnimation(gc.StringName(name), animation))
}

/*
Removes the [Animation] with the key [param name].
*/
func (self Go) RemoveAnimation(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveAnimation(gc.StringName(name))
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
func (self Go) RenameAnimation(name string, newname string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameAnimation(gc.StringName(name), gc.StringName(newname))
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
func (self Go) HasAnimation(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasAnimation(gc.StringName(name)))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Go) GetAnimation(name string) gdclass.Animation {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Animation(class(self).GetAnimation(gc, gc.StringName(name)))
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
func (self Go) GetAnimationList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](class(self).GetAnimationList(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationLibrary
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationLibrary"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
//go:nosplit
func (self class) AddAnimation(name gd.StringName, animation gdclass.Animation) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(animation[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_add_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the [Animation] with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimation(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(name gd.StringName, newname gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(newname))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_rename_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
//go:nosplit
func (self class) HasAnimation(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
//go:nosplit
func (self class) GetAnimation(ctx gd.Lifetime, name gd.StringName) gdclass.Animation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Animation
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the keys for the [Animation]s stored in the library.
*/
//go:nosplit
func (self class) GetAnimationList(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
func (self Go) OnAnimationAdded(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_added"), gc.Callable(cb), 0)
}


func (self Go) OnAnimationRemoved(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_removed"), gc.Callable(cb), 0)
}


func (self Go) OnAnimationRenamed(cb func(name string, to_name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_renamed"), gc.Callable(cb), 0)
}


func (self Go) OnAnimationChanged(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_changed"), gc.Callable(cb), 0)
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
func init() {classdb.Register("AnimationLibrary", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
