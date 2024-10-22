package VideoStream

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
Base resource type for all video streams. Classes that derive from [VideoStream] can all be used as resource types to play back videos in [VideoStreamPlayer].
	// VideoStream methods that can be overridden by a [Class] that extends it.
	type VideoStream interface {
		//Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
		InstantiatePlayback() gdclass.VideoStreamPlayback
	}

*/
type Go [1]classdb.VideoStream

/*
Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
*/
func (Go) _instantiate_playback(impl func(ptr unsafe.Pointer) gdclass.VideoStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VideoStream
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VideoStream"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) File() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetFile(gc).String())
}

func (self Go) SetFile(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFile(gc.String(value))
}

/*
Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
*/
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) gdclass.VideoStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

//go:nosplit
func (self class) SetFile(file gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStream.Bind_set_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFile(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStream.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVideoStream() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVideoStream() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("VideoStream", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
