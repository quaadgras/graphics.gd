// Package AnimatedSprite3D provides methods for working with AnimatedSprite3D object instances.
package AnimatedSprite3D

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
import "graphics.gd/classdb/SpriteBase3D"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
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
[AnimatedSprite3D] is similar to the [Sprite3D] node, except it carries multiple textures as animation [member sprite_frames]. Animations are created using a [SpriteFrames] resource, which allows you to import image files (or a folder containing said files) to provide the animation frames for the sprite. The [SpriteFrames] resource can be configured in the editor via the SpriteFrames bottom panel.
*/
type Instance [1]gdclass.AnimatedSprite3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimatedSprite3D() Instance
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
func (self Instance) IsPlaying() bool {
	return bool(class(self).IsPlaying())
}

/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
func (self Instance) Play() {
	class(self).Play(gd.NewStringName(""), gd.Float(1.0), false)
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
func (self Instance) PlayBackwards() {
	class(self).PlayBackwards(gd.NewStringName(""))
}

/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
func (self Instance) Pause() {
	class(self).Pause()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
func (self Instance) Stop() {
	class(self).Stop()
}

/*
The setter of [member frame] resets the [member frame_progress] to [code]0.0[/code] implicitly, but this method avoids that.
This is useful when you want to carry over the current [member frame_progress] to another [member frame].
[b]Example:[/b]
[codeblocks]
[gdscript]
# Change the animation with keeping the frame index and progress.
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
[/codeblocks]
*/
func (self Instance) SetFrameAndProgress(frame_ int, progress Float.X) {
	class(self).SetFrameAndProgress(gd.Int(frame_), gd.Float(progress))
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
func (self Instance) GetPlayingSpeed() Float.X {
	return Float.X(Float.X(class(self).GetPlayingSpeed()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimatedSprite3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatedSprite3D"))
	casted := Instance{*(*gdclass.AnimatedSprite3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) SpriteFrames() [1]gdclass.SpriteFrames {
	return [1]gdclass.SpriteFrames(class(self).GetSpriteFrames())
}

func (self Instance) SetSpriteFrames(value [1]gdclass.SpriteFrames) {
	class(self).SetSpriteFrames(value)
}

func (self Instance) Animation() string {
	return string(class(self).GetAnimation().String())
}

func (self Instance) SetAnimation(value string) {
	class(self).SetAnimation(gd.NewStringName(value))
}

func (self Instance) Autoplay() string {
	return string(class(self).GetAutoplay().String())
}

func (self Instance) SetAutoplay(value string) {
	class(self).SetAutoplay(gd.NewString(value))
}

func (self Instance) Frame() int {
	return int(int(class(self).GetFrame()))
}

func (self Instance) SetFrame(value int) {
	class(self).SetFrame(gd.Int(value))
}

func (self Instance) FrameProgress() Float.X {
	return Float.X(Float.X(class(self).GetFrameProgress()))
}

func (self Instance) SetFrameProgress(value Float.X) {
	class(self).SetFrameProgress(gd.Float(value))
}

func (self Instance) SpeedScale() Float.X {
	return Float.X(Float.X(class(self).GetSpeedScale()))
}

func (self Instance) SetSpeedScale(value Float.X) {
	class(self).SetSpeedScale(gd.Float(value))
}

//go:nosplit
func (self class) SetSpriteFrames(sprite_frames [1]gdclass.SpriteFrames) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sprite_frames[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpriteFrames() [1]gdclass.SpriteFrames {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SpriteFrames{gd.PointerWithOwnershipTransferredToGo[gdclass.SpriteFrames](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimation(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimation() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoplay() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
//go:nosplit
func (self class) Play(name gd.StringName, custom_speed gd.Float, from_end bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
//go:nosplit
func (self class) PlayBackwards(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_play_backwards, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
//go:nosplit
func (self class) Pause() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFrame(frame_ gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrame() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrameProgress(progress gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrameProgress() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_frame_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The setter of [member frame] resets the [member frame_progress] to [code]0.0[/code] implicitly, but this method avoids that.
This is useful when you want to carry over the current [member frame_progress] to another [member frame].
[b]Example:[/b]
[codeblocks]
[gdscript]
# Change the animation with keeping the frame index and progress.
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) SetFrameAndProgress(frame_ gd.Int, progress gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame_and_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(speed_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, speed_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
//go:nosplit
func (self class) GetPlayingSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_playing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnSpriteFramesChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("sprite_frames_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFrameChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("frame_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationLooped(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_looped"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_finished"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimatedSprite3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatedSprite3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSpriteBase3D() SpriteBase3D.Advanced {
	return *((*SpriteBase3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSpriteBase3D() SpriteBase3D.Instance {
	return *((*SpriteBase3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SpriteBase3D.Advanced(self.AsSpriteBase3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SpriteBase3D.Instance(self.AsSpriteBase3D()), name)
	}
}
func init() {
	gdclass.Register("AnimatedSprite3D", func(ptr gd.Object) any {
		return [1]gdclass.AnimatedSprite3D{*(*gdclass.AnimatedSprite3D)(unsafe.Pointer(&ptr))}
	})
}
