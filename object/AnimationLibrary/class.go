package AnimationLibrary

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
An animation library stores a set of animations accessible through [StringName] keys, for use with [AnimationPlayer] nodes.

*/
type Simple [1]classdb.AnimationLibrary
func (self Simple) AddAnimation(name string, animation [1]classdb.Animation) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AddAnimation(gc.StringName(name), animation))
}
func (self Simple) RemoveAnimation(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveAnimation(gc.StringName(name))
}
func (self Simple) RenameAnimation(name string, newname string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RenameAnimation(gc.StringName(name), gc.StringName(newname))
}
func (self Simple) HasAnimation(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAnimation(gc.StringName(name)))
}
func (self Simple) GetAnimation(name string) [1]classdb.Animation {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Animation(Expert(self).GetAnimation(gc, gc.StringName(name)))
}
func (self Simple) GetAnimationList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetAnimationList(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationLibrary
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
//go:nosplit
func (self class) AddAnimation(name gd.StringName, animation object.Animation) int64 {
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
func (self class) GetAnimation(ctx gd.Lifetime, name gd.StringName) object.Animation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationLibrary.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Animation
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

//go:nosplit
func (self class) AsAnimationLibrary() Expert { return self[0].AsAnimationLibrary() }


//go:nosplit
func (self Simple) AsAnimationLibrary() Simple { return self[0].AsAnimationLibrary() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AnimationLibrary", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
