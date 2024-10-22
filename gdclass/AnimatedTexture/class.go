package AnimatedTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[AnimatedTexture] is a resource format for frame-based animations, where multiple textures can be chained automatically with a predefined delay for each frame. Unlike [AnimationPlayer] or [AnimatedSprite2D], it isn't a [Node], but has the advantage of being usable anywhere a [Texture2D] resource can be used, e.g. in a [TileSet].
The playback of the animation is controlled by the [member speed_scale] property, as well as each frame's duration (see [method set_frame_duration]). The animation loops, i.e. it will restart at frame 0 automatically after playing the last frame.
[AnimatedTexture] currently requires all frame textures to have the same size, otherwise the bigger ones will be cropped to match the smallest one.
[b]Note:[/b] AnimatedTexture doesn't support using [AtlasTexture]s. Each frame needs to be a separate [Texture2D].
[b]Warning:[/b] The current implementation is not efficient for the modern renderers.

*/
type Go [1]classdb.AnimatedTexture

/*
Assigns a [Texture2D] to the given frame. Frame IDs start at 0, so the first frame has ID 0, and the last frame of the animation has ID [member frames] - 1.
You can define any number of textures up to [constant MAX_FRAMES], but keep in mind that only frames from 0 to [member frames] - 1 will be part of the animation.
*/
func (self Go) SetFrameTexture(frame_ int, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrameTexture(gd.Int(frame_), texture)
}

/*
Returns the given frame's [Texture2D].
*/
func (self Go) GetFrameTexture(frame_ int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetFrameTexture(gc, gd.Int(frame_)))
}

/*
Sets the duration of any given [param frame]. The final duration is affected by the [member speed_scale]. If set to [code]0[/code], the frame is skipped during playback.
*/
func (self Go) SetFrameDuration(frame_ int, duration float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrameDuration(gd.Int(frame_), gd.Float(duration))
}

/*
Returns the given [param frame]'s duration, in seconds.
*/
func (self Go) GetFrameDuration(frame_ int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetFrameDuration(gd.Int(frame_))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimatedTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimatedTexture"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Frames() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFrames()))
}

func (self Go) SetFrames(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrames(gd.Int(value))
}

func (self Go) CurrentFrame() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCurrentFrame()))
}

func (self Go) SetCurrentFrame(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentFrame(gd.Int(value))
}

func (self Go) Pause() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetPause())
}

func (self Go) SetPause(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPause(value)
}

func (self Go) OneShot() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetOneShot())
}

func (self Go) SetOneShot(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOneShot(value)
}

func (self Go) SpeedScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSpeedScale()))
}

func (self Go) SetSpeedScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpeedScale(gd.Float(value))
}

//go:nosplit
func (self class) SetFrames(frames gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentFrame(frame_ gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_current_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_current_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPause(pause bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPause() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOneShot(one_shot bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, one_shot)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOneShot() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpeedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns a [Texture2D] to the given frame. Frame IDs start at 0, so the first frame has ID 0, and the last frame of the animation has ID [member frames] - 1.
You can define any number of textures up to [constant MAX_FRAMES], but keep in mind that only frames from 0 to [member frames] - 1 will be part of the animation.
*/
//go:nosplit
func (self class) SetFrameTexture(frame_ gd.Int, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_frame_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given frame's [Texture2D].
*/
//go:nosplit
func (self class) GetFrameTexture(ctx gd.Lifetime, frame_ gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_frame_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the duration of any given [param frame]. The final duration is affected by the [member speed_scale]. If set to [code]0[/code], the frame is skipped during playback.
*/
//go:nosplit
func (self class) SetFrameDuration(frame_ gd.Int, duration gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, duration)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_set_frame_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given [param frame]'s duration, in seconds.
*/
//go:nosplit
func (self class) GetFrameDuration(frame_ gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimatedTexture.Bind_get_frame_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimatedTexture() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimatedTexture() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("AnimatedTexture", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
