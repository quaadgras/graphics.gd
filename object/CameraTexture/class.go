package CameraTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This texture gives access to the camera texture provided by a [CameraFeed].
[b]Note:[/b] Many cameras supply YCbCr images which need to be converted in a shader.

*/
type Simple [1]classdb.CameraTexture
func (self Simple) SetCameraFeedId(feed_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameraFeedId(gd.Int(feed_id))
}
func (self Simple) GetCameraFeedId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCameraFeedId()))
}
func (self Simple) SetWhichFeed(which_feed classdb.CameraServerFeedImage) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWhichFeed(which_feed)
}
func (self Simple) GetWhichFeed() classdb.CameraServerFeedImage {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CameraServerFeedImage(Expert(self).GetWhichFeed())
}
func (self Simple) SetCameraActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameraActive(active)
}
func (self Simple) GetCameraActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCameraActive())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CameraTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetCameraFeedId(feed_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feed_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWhichFeed(which_feed classdb.CameraServerFeedImage)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, which_feed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_set_which_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWhichFeed() classdb.CameraServerFeedImage {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraServerFeedImage](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_get_which_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_set_camera_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraTexture.Bind_get_camera_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCameraTexture() Expert { return self[0].AsCameraTexture() }


//go:nosplit
func (self Simple) AsCameraTexture() Simple { return self[0].AsCameraTexture() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CameraTexture", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
