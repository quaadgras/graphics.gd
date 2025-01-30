// Package SpriteFrames provides methods for working with SpriteFrames object instances.
package SpriteFrames

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Sprite frame library for an [AnimatedSprite2D] or [AnimatedSprite3D] node. Contains frames and animation data for playback.
*/
type Instance [1]gdclass.SpriteFrames

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSpriteFrames() Instance
}

/*
Adds a new [param anim] animation to the library.
*/
func (self Instance) AddAnimation(anim string) { //gd:SpriteFrames.add_animation
	class(self).AddAnimation(String.Name(String.New(anim)))
}

/*
Returns [code]true[/code] if the [param anim] animation exists.
*/
func (self Instance) HasAnimation(anim string) bool { //gd:SpriteFrames.has_animation
	return bool(class(self).HasAnimation(String.Name(String.New(anim))))
}

/*
Removes the [param anim] animation.
*/
func (self Instance) RemoveAnimation(anim string) { //gd:SpriteFrames.remove_animation
	class(self).RemoveAnimation(String.Name(String.New(anim)))
}

/*
Changes the [param anim] animation's name to [param newname].
*/
func (self Instance) RenameAnimation(anim string, newname string) { //gd:SpriteFrames.rename_animation
	class(self).RenameAnimation(String.Name(String.New(anim)), String.Name(String.New(newname)))
}

/*
Returns an array containing the names associated to each animation. Values are placed in alphabetical order.
*/
func (self Instance) GetAnimationNames() []string { //gd:SpriteFrames.get_animation_names
	return []string(class(self).GetAnimationNames().Strings())
}

/*
Sets the speed for the [param anim] animation in frames per second.
*/
func (self Instance) SetAnimationSpeed(anim string, fps Float.X) { //gd:SpriteFrames.set_animation_speed
	class(self).SetAnimationSpeed(String.Name(String.New(anim)), float64(fps))
}

/*
Returns the speed in frames per second for the [param anim] animation.
*/
func (self Instance) GetAnimationSpeed(anim string) Float.X { //gd:SpriteFrames.get_animation_speed
	return Float.X(Float.X(class(self).GetAnimationSpeed(String.Name(String.New(anim)))))
}

/*
If [param loop] is [code]true[/code], the [param anim] animation will loop when it reaches the end, or the start if it is played in reverse.
*/
func (self Instance) SetAnimationLoop(anim string, loop bool) { //gd:SpriteFrames.set_animation_loop
	class(self).SetAnimationLoop(String.Name(String.New(anim)), loop)
}

/*
Returns [code]true[/code] if the given animation is configured to loop when it finishes playing. Otherwise, returns [code]false[/code].
*/
func (self Instance) GetAnimationLoop(anim string) bool { //gd:SpriteFrames.get_animation_loop
	return bool(class(self).GetAnimationLoop(String.Name(String.New(anim))))
}

/*
Adds a frame to the [param anim] animation. If [param at_position] is [code]-1[/code], the frame will be added to the end of the animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
func (self Instance) AddFrame(anim string, texture [1]gdclass.Texture2D) { //gd:SpriteFrames.add_frame
	class(self).AddFrame(String.Name(String.New(anim)), texture, float64(1.0), int64(-1))
}

/*
Sets the [param texture] and the [param duration] of the frame [param idx] in the [param anim] animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
func (self Instance) SetFrame(anim string, idx int, texture [1]gdclass.Texture2D) { //gd:SpriteFrames.set_frame
	class(self).SetFrame(String.Name(String.New(anim)), int64(idx), texture, float64(1.0))
}

/*
Removes the [param anim] animation's frame [param idx].
*/
func (self Instance) RemoveFrame(anim string, idx int) { //gd:SpriteFrames.remove_frame
	class(self).RemoveFrame(String.Name(String.New(anim)), int64(idx))
}

/*
Returns the number of frames for the [param anim] animation.
*/
func (self Instance) GetFrameCount(anim string) int { //gd:SpriteFrames.get_frame_count
	return int(int(class(self).GetFrameCount(String.Name(String.New(anim)))))
}

/*
Returns the texture of the frame [param idx] in the [param anim] animation.
*/
func (self Instance) GetFrameTexture(anim string, idx int) [1]gdclass.Texture2D { //gd:SpriteFrames.get_frame_texture
	return [1]gdclass.Texture2D(class(self).GetFrameTexture(String.Name(String.New(anim)), int64(idx)))
}

/*
Returns a relative duration of the frame [param idx] in the [param anim] animation (defaults to [code]1.0[/code]). For example, a frame with a duration of [code]2.0[/code] is displayed twice as long as a frame with a duration of [code]1.0[/code]. You can calculate the absolute duration (in seconds) of a frame using the following formula:
[codeblock]
absolute_duration = relative_duration / (animation_fps * abs(playing_speed))
[/codeblock]
In this example, [code]playing_speed[/code] refers to either [method AnimatedSprite2D.get_playing_speed] or [method AnimatedSprite3D.get_playing_speed].
*/
func (self Instance) GetFrameDuration(anim string, idx int) Float.X { //gd:SpriteFrames.get_frame_duration
	return Float.X(Float.X(class(self).GetFrameDuration(String.Name(String.New(anim)), int64(idx))))
}

/*
Removes all frames from the [param anim] animation.
*/
func (self Instance) Clear(anim string) { //gd:SpriteFrames.clear
	class(self).Clear(String.Name(String.New(anim)))
}

/*
Removes all animations. An empty [code]default[/code] animation will be created.
*/
func (self Instance) ClearAll() { //gd:SpriteFrames.clear_all
	class(self).ClearAll()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SpriteFrames

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SpriteFrames"))
	casted := Instance{*(*gdclass.SpriteFrames)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Adds a new [param anim] animation to the library.
*/
//go:nosplit
func (self class) AddAnimation(anim String.Name) { //gd:SpriteFrames.add_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_add_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [param anim] animation exists.
*/
//go:nosplit
func (self class) HasAnimation(anim String.Name) bool { //gd:SpriteFrames.has_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the [param anim] animation.
*/
//go:nosplit
func (self class) RemoveAnimation(anim String.Name) { //gd:SpriteFrames.remove_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the [param anim] animation's name to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(anim String.Name, newname String.Name) { //gd:SpriteFrames.rename_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(newname)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_rename_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array containing the names associated to each animation. Values are placed in alphabetical order.
*/
//go:nosplit
func (self class) GetAnimationNames() Packed.Strings { //gd:SpriteFrames.get_animation_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_animation_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the speed for the [param anim] animation in frames per second.
*/
//go:nosplit
func (self class) SetAnimationSpeed(anim String.Name, fps float64) { //gd:SpriteFrames.set_animation_speed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, fps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_set_animation_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the speed in frames per second for the [param anim] animation.
*/
//go:nosplit
func (self class) GetAnimationSpeed(anim String.Name) float64 { //gd:SpriteFrames.get_animation_speed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_animation_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param loop] is [code]true[/code], the [param anim] animation will loop when it reaches the end, or the start if it is played in reverse.
*/
//go:nosplit
func (self class) SetAnimationLoop(anim String.Name, loop bool) { //gd:SpriteFrames.set_animation_loop
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_set_animation_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given animation is configured to loop when it finishes playing. Otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) GetAnimationLoop(anim String.Name) bool { //gd:SpriteFrames.get_animation_loop
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_animation_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a frame to the [param anim] animation. If [param at_position] is [code]-1[/code], the frame will be added to the end of the animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
//go:nosplit
func (self class) AddFrame(anim String.Name, texture [1]gdclass.Texture2D, duration float64, at_position int64) { //gd:SpriteFrames.add_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, duration)
	callframe.Arg(frame, at_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_add_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param texture] and the [param duration] of the frame [param idx] in the [param anim] animation. [param duration] specifies the relative duration, see [method get_frame_duration] for details.
*/
//go:nosplit
func (self class) SetFrame(anim String.Name, idx int64, texture [1]gdclass.Texture2D, duration float64) { //gd:SpriteFrames.set_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, duration)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the [param anim] animation's frame [param idx].
*/
//go:nosplit
func (self class) RemoveFrame(anim String.Name, idx int64) { //gd:SpriteFrames.remove_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_remove_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of frames for the [param anim] animation.
*/
//go:nosplit
func (self class) GetFrameCount(anim String.Name) int64 { //gd:SpriteFrames.get_frame_count
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_frame_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture of the frame [param idx] in the [param anim] animation.
*/
//go:nosplit
func (self class) GetFrameTexture(anim String.Name, idx int64) [1]gdclass.Texture2D { //gd:SpriteFrames.get_frame_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_frame_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
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
func (self class) GetFrameDuration(anim String.Name, idx int64) float64 { //gd:SpriteFrames.get_frame_duration
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_get_frame_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all frames from the [param anim] animation.
*/
//go:nosplit
func (self class) Clear(anim String.Name) { //gd:SpriteFrames.clear
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(anim)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all animations. An empty [code]default[/code] animation will be created.
*/
//go:nosplit
func (self class) ClearAll() { //gd:SpriteFrames.clear_all
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteFrames.Bind_clear_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsSpriteFrames() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSpriteFrames() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("SpriteFrames", func(ptr gd.Object) any {
		return [1]gdclass.SpriteFrames{*(*gdclass.SpriteFrames)(unsafe.Pointer(&ptr))}
	})
}
