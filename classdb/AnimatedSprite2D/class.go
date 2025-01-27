// Package AnimatedSprite2D provides methods for working with AnimatedSprite2D object instances.
package AnimatedSprite2D

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

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

/*
[AnimatedSprite2D] is similar to the [Sprite2D] node, except it carries multiple textures as animation frames. Animations are created using a [SpriteFrames] resource, which allows you to import image files (or a folder containing said files) to provide the animation frames for the sprite. The [SpriteFrames] resource can be configured in the editor via the SpriteFrames bottom panel.
*/
type Instance [1]gdclass.AnimatedSprite2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimatedSprite2D() Instance
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
func (self Instance) IsPlaying() bool { //gd:AnimatedSprite2D.is_playing
	return bool(class(self).IsPlaying())
}

/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
func (self Instance) Play() { //gd:AnimatedSprite2D.play
	class(self).Play(String.Name(String.New("")), gd.Float(1.0), false)
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
func (self Instance) PlayBackwards() { //gd:AnimatedSprite2D.play_backwards
	class(self).PlayBackwards(String.Name(String.New("")))
}

/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
func (self Instance) Pause() { //gd:AnimatedSprite2D.pause
	class(self).Pause()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
func (self Instance) Stop() { //gd:AnimatedSprite2D.stop
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
func (self Instance) SetFrameAndProgress(frame_ int, progress Float.X) { //gd:AnimatedSprite2D.set_frame_and_progress
	class(self).SetFrameAndProgress(gd.Int(frame_), gd.Float(progress))
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
func (self Instance) GetPlayingSpeed() Float.X { //gd:AnimatedSprite2D.get_playing_speed
	return Float.X(Float.X(class(self).GetPlayingSpeed()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimatedSprite2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatedSprite2D"))
	casted := Instance{*(*gdclass.AnimatedSprite2D)(unsafe.Pointer(&object))}
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
	class(self).SetAnimation(String.Name(String.New(value)))
}

func (self Instance) Autoplay() string {
	return string(class(self).GetAutoplay().String())
}

func (self Instance) SetAutoplay(value string) {
	class(self).SetAutoplay(String.New(value))
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

func (self Instance) Centered() bool {
	return bool(class(self).IsCentered())
}

func (self Instance) SetCentered(value bool) {
	class(self).SetCentered(value)
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) FlipH() bool {
	return bool(class(self).IsFlippedH())
}

func (self Instance) SetFlipH(value bool) {
	class(self).SetFlipH(value)
}

func (self Instance) FlipV() bool {
	return bool(class(self).IsFlippedV())
}

func (self Instance) SetFlipV(value bool) {
	class(self).SetFlipV(value)
}

//go:nosplit
func (self class) SetSpriteFrames(sprite_frames [1]gdclass.SpriteFrames) { //gd:AnimatedSprite2D.set_sprite_frames
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sprite_frames[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpriteFrames() [1]gdclass.SpriteFrames { //gd:AnimatedSprite2D.get_sprite_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SpriteFrames{gd.PointerWithOwnershipTransferredToGo[gdclass.SpriteFrames](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimation(name String.Name) { //gd:AnimatedSprite2D.set_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimation() String.Name { //gd:AnimatedSprite2D.get_animation
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(name String.Readable) { //gd:AnimatedSprite2D.set_autoplay
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoplay() String.Readable { //gd:AnimatedSprite2D.get_autoplay
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
//go:nosplit
func (self class) IsPlaying() bool { //gd:AnimatedSprite2D.is_playing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
//go:nosplit
func (self class) Play(name String.Name, custom_speed gd.Float, from_end bool) { //gd:AnimatedSprite2D.play
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
//go:nosplit
func (self class) PlayBackwards(name String.Name) { //gd:AnimatedSprite2D.play_backwards
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_play_backwards, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
//go:nosplit
func (self class) Pause() { //gd:AnimatedSprite2D.pause
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
//go:nosplit
func (self class) Stop() { //gd:AnimatedSprite2D.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCentered(centered bool) { //gd:AnimatedSprite2D.set_centered
	var frame = callframe.New()
	callframe.Arg(frame, centered)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCentered() bool { //gd:AnimatedSprite2D.is_centered
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_is_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) { //gd:AnimatedSprite2D.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:AnimatedSprite2D.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipH(flip_h bool) { //gd:AnimatedSprite2D.set_flip_h
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool { //gd:AnimatedSprite2D.is_flipped_h
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(flip_v bool) { //gd:AnimatedSprite2D.set_flip_v
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool { //gd:AnimatedSprite2D.is_flipped_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrame(frame_ gd.Int) { //gd:AnimatedSprite2D.set_frame
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrame() gd.Int { //gd:AnimatedSprite2D.get_frame
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrameProgress(progress gd.Float) { //gd:AnimatedSprite2D.set_frame_progress
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_frame_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrameProgress() gd.Float { //gd:AnimatedSprite2D.get_frame_progress
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_frame_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetFrameAndProgress(frame_ gd.Int, progress gd.Float) { //gd:AnimatedSprite2D.set_frame_and_progress
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_frame_and_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(speed_scale gd.Float) { //gd:AnimatedSprite2D.set_speed_scale
	var frame = callframe.New()
	callframe.Arg(frame, speed_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float { //gd:AnimatedSprite2D.get_speed_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
//go:nosplit
func (self class) GetPlayingSpeed() gd.Float { //gd:AnimatedSprite2D.get_playing_speed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite2D.Bind_get_playing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
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

func (self class) AsAnimatedSprite2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatedSprite2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced       { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance    { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("AnimatedSprite2D", func(ptr gd.Object) any {
		return [1]gdclass.AnimatedSprite2D{*(*gdclass.AnimatedSprite2D)(unsafe.Pointer(&ptr))}
	})
}
