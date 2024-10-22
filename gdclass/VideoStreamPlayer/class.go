package VideoStreamPlayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A control used for playback of [VideoStream] resources.
Supported video formats are [url=https://www.theora.org/]Ogg Theora[/url] ([code].ogv[/code], [VideoStreamTheora]) and any format exposed via a GDExtension plugin.
[b]Warning:[/b] On Web, video playback [i]will[/i] perform poorly due to missing architecture-specific assembly optimizations.

*/
type Go [1]classdb.VideoStreamPlayer

/*
Starts the video playback from the beginning. If the video is paused, this will not unpause the video.
*/
func (self Go) Play() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Play()
}

/*
Stops the video playback and sets the stream position to 0.
[b]Note:[/b] Although the stream position will be set to 0, the first frame of the video stream won't become the current frame.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}

/*
Returns [code]true[/code] if the video is playing.
[b]Note:[/b] The video is still considered playing if paused during playback.
*/
func (self Go) IsPlaying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPlaying())
}

/*
Returns the video stream's name, or [code]"<No Stream>"[/code] if no video stream is assigned.
*/
func (self Go) GetStreamName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetStreamName(gc).String())
}

/*
The length of the current stream, in seconds.
[b]Note:[/b] For [VideoStreamTheora] streams (the built-in format supported by Godot), this value will always be zero, as getting the stream length is not implemented yet. The feature may be supported by video formats implemented by a GDExtension add-on.
*/
func (self Go) GetStreamLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetStreamLength()))
}

/*
Returns the current frame as a [Texture2D].
*/
func (self Go) GetVideoTexture() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetVideoTexture(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VideoStreamPlayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VideoStreamPlayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) AudioTrack() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetAudioTrack()))
}

func (self Go) SetAudioTrack(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAudioTrack(gd.Int(value))
}

func (self Go) Stream() gdclass.VideoStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.VideoStream(class(self).GetStream(gc))
}

func (self Go) SetStream(value gdclass.VideoStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStream(value)
}

func (self Go) VolumeDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumeDb()))
}

func (self Go) SetVolumeDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumeDb(gd.Float(value))
}

func (self Go) Volume() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolume()))
}

func (self Go) SetVolume(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolume(gd.Float(value))
}

func (self Go) Autoplay() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasAutoplay())
}

func (self Go) SetAutoplay(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoplay(value)
}

func (self Go) Paused() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPaused())
}

func (self Go) SetPaused(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPaused(value)
}

func (self Go) Expand() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasExpand())
}

func (self Go) SetExpand(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpand(value)
}

func (self Go) Loop() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasLoop())
}

func (self Go) SetLoop(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoop(value)
}

func (self Go) BufferingMsec() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBufferingMsec()))
}

func (self Go) SetBufferingMsec(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferingMsec(gd.Int(value))
}

func (self Go) StreamPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetStreamPosition()))
}

func (self Go) SetStreamPosition(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStreamPosition(gd.Float(value))
}

func (self Go) Bus() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBus(gc).String())
}

func (self Go) SetBus(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBus(gc.StringName(value))
}

//go:nosplit
func (self class) SetStream(stream gdclass.VideoStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime) gdclass.VideoStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.VideoStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Starts the video playback from the beginning. If the video is paused, this will not unpause the video.
*/
//go:nosplit
func (self class) Play()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the video playback and sets the stream position to 0.
[b]Note:[/b] Although the stream position will be set to 0, the first frame of the video stream won't become the current frame.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the video is playing.
[b]Note:[/b] The video is still considered playing if paused during playback.
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPaused(paused bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paused)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolume(volume gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, volume)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolume() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumeDb(db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumeDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioTrack(track gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_audio_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAudioTrack() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_audio_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the video stream's name, or [code]"<No Stream>"[/code] if no video stream is assigned.
*/
//go:nosplit
func (self class) GetStreamName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_stream_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
The length of the current stream, in seconds.
[b]Note:[/b] For [VideoStreamTheora] streams (the built-in format supported by Godot), this value will always be zero, as getting the stream length is not implemented yet. The feature may be supported by video formats implemented by a GDExtension add-on.
*/
//go:nosplit
func (self class) GetStreamLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_stream_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStreamPosition(position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_stream_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamPosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_stream_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoplay(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasAutoplay() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_has_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExpand(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_expand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasExpand() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_has_expand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBufferingMsec(msec gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_buffering_msec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBufferingMsec() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_buffering_msec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBus(bus gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bus))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBus(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current frame as a [Texture2D].
*/
//go:nosplit
func (self class) GetVideoTexture(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayer.Bind_get_video_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsVideoStreamPlayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVideoStreamPlayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("VideoStreamPlayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
