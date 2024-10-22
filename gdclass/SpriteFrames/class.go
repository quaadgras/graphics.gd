package SpriteFrames

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
Sprite frame library for an [AnimatedSprite2D] or [AnimatedSprite3D] node. Contains frames and animation data for playback.

*/
type Go [1]classdb.SpriteFrames

/*
Adds a new [param anim] animation to the library.
*/
func (self Go) AddAnimation(anim string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddAnimation(gc.StringName(anim))
}

/*
Returns [code]true[/code] if the [param anim] animation exists.
*/
func (self Go) HasAnimation(anim string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasAnimation(gc.StringName(anim)))
}

/*
Removes the [param anim] animation.
*/
func (self Go) RemoveAnimation(anim string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveAnimation(gc.StringName(anim))
}

/*
Changes the [param anim] animation's name to [param newname].
*/
func (self Go) RenameAnimation(anim string, newname string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameAnimation(gc.StringName(anim), gc.StringName(newname))
}

/*
Returns an array containing the names associated to each animation. Values are placed in alphabetical order.
*/
func (self Go) GetAnimationNames() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetAnimationNames(gc).Strings(gc))
}

/*
Sets the speed for the [param anim] animation in frames per second.
*/
func (self Go) SetAnimationSpeed(anim string, fps float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnimationSpeed(gc.StringName(anim), gd.Float(fps))
}

/*
Returns the speed in frames per second for the [param anim] animation.
*/
func (self Go) GetAnimationSpeed(anim string) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetAnimationSpeed(gc.StringName(anim))))
}

/*
If [param loop] is [code]true[/code], the [param anim] animation will loop when it reaches the end, or the start if it is played in reverse.
*/
func (self Go) SetAnimationLoop(anim string, loop bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnimationLoop(gc.StringName(anim), loop)
}

/*
Returns [code]true[/code] if the given animation is configured to loop when it finishes playing. Otherwise, returns [code]false[/code].
*/
func (self Go) GetAnimationLoop(anim string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetAnimationLoop(gc.StringName(anim)))
}

/*
Adds a frame to the [param anim] animation. If [param at_position] is [code]-1[/code], the frame will be added to the end of the animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
func (self Go) AddFrame(anim string, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddFrame(gc.StringName(anim), texture, gd.Float(1.0), gd.Int(-1))
}

/*
Sets the [param texture] and the [param duration] of the frame [param idx] in the [param anim] animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
func (self Go) SetFrame(anim string, idx int, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrame(gc.StringName(anim), gd.Int(idx), texture, gd.Float(1.0))
}

/*
Removes the [param anim] animation's frame [param idx].
*/
func (self Go) RemoveFrame(anim string, idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveFrame(gc.StringName(anim), gd.Int(idx))
}

/*
Returns the number of frames for the [param anim] animation.
*/
func (self Go) GetFrameCount(anim string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFrameCount(gc.StringName(anim))))
}

/*
Returns the texture of the frame [param idx] in the [param anim] animation.
*/
func (self Go) GetFrameTexture(anim string, idx int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetFrameTexture(gc, gc.StringName(anim), gd.Int(idx)))
}

/*
Returns a relative duration of the frame [param idx] in the [param anim] animation (defaults to [code]1.0[/code]). For example, a frame with a duration of [code]2.0[/code] is displayed twice as long as a frame with a duration of [code]1.0[/code]. You can calculate the absolute duration (in seconds) of a frame using the following formula:
[codeblock]
absolute_duration = relative_duration / (animation_fps * abs(playing_speed))
[/codeblock]
In this example, [code]playing_speed[/code] refers to either [method AnimatedSprite2D.get_playing_speed] or [method AnimatedSprite3D.get_playing_speed].
*/
func (self Go) GetFrameDuration(anim string, idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetFrameDuration(gc.StringName(anim), gd.Int(idx))))
}

/*
Removes all frames from the [param anim] animation.
*/
func (self Go) Clear(anim string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear(gc.StringName(anim))
}

/*
Removes all animations. An empty [code]default[/code] animation will be created.
*/
func (self Go) ClearAll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearAll()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SpriteFrames
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SpriteFrames"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Adds a new [param anim] animation to the library.
*/
//go:nosplit
func (self class) AddAnimation(anim gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_add_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [param anim] animation exists.
*/
//go:nosplit
func (self class) HasAnimation(anim gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the [param anim] animation.
*/
//go:nosplit
func (self class) RemoveAnimation(anim gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the [param anim] animation's name to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(anim gd.StringName, newname gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, mmm.Get(newname))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_rename_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array containing the names associated to each animation. Values are placed in alphabetical order.
*/
//go:nosplit
func (self class) GetAnimationNames(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_animation_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the speed for the [param anim] animation in frames per second.
*/
//go:nosplit
func (self class) SetAnimationSpeed(anim gd.StringName, fps gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_set_animation_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the speed in frames per second for the [param anim] animation.
*/
//go:nosplit
func (self class) GetAnimationSpeed(anim gd.StringName) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_animation_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param loop] is [code]true[/code], the [param anim] animation will loop when it reaches the end, or the start if it is played in reverse.
*/
//go:nosplit
func (self class) SetAnimationLoop(anim gd.StringName, loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_set_animation_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given animation is configured to loop when it finishes playing. Otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) GetAnimationLoop(anim gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_animation_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a frame to the [param anim] animation. If [param at_position] is [code]-1[/code], the frame will be added to the end of the animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
//go:nosplit
func (self class) AddFrame(anim gd.StringName, texture gdclass.Texture2D, duration gd.Float, at_position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, duration)
	callframe.Arg(frame, at_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_add_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param texture] and the [param duration] of the frame [param idx] in the [param anim] animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
//go:nosplit
func (self class) SetFrame(anim gd.StringName, idx gd.Int, texture gdclass.Texture2D, duration gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, duration)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the [param anim] animation's frame [param idx].
*/
//go:nosplit
func (self class) RemoveFrame(anim gd.StringName, idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_remove_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of frames for the [param anim] animation.
*/
//go:nosplit
func (self class) GetFrameCount(anim gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_frame_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the texture of the frame [param idx] in the [param anim] animation.
*/
//go:nosplit
func (self class) GetFrameTexture(ctx gd.Lifetime, anim gd.StringName, idx gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_frame_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a relative duration of the frame [param idx] in the [param anim] animation (defaults to [code]1.0[/code]). For example, a frame with a duration of [code]2.0[/code] is displayed twice as long as a frame with a duration of [code]1.0[/code]. You can calculate the absolute duration (in seconds) of a frame using the following formula:
[codeblock]
absolute_duration = relative_duration / (animation_fps * abs(playing_speed))
[/codeblock]
In this example, [code]playing_speed[/code] refers to either [method AnimatedSprite2D.get_playing_speed] or [method AnimatedSprite3D.get_playing_speed].
*/
//go:nosplit
func (self class) GetFrameDuration(anim gd.StringName, idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_get_frame_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all frames from the [param anim] animation.
*/
//go:nosplit
func (self class) Clear(anim gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(anim))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all animations. An empty [code]default[/code] animation will be created.
*/
//go:nosplit
func (self class) ClearAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteFrames.Bind_clear_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsSpriteFrames() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSpriteFrames() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SpriteFrames", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
