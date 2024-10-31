package CameraTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This texture gives access to the camera texture provided by a [CameraFeed].
[b]Note:[/b] Many cameras supply YCbCr images which need to be converted in a shader.
*/
type Instance [1]classdb.CameraTexture

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CameraTexture

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraTexture"))
	return Instance{classdb.CameraTexture(object)}
}

func (self Instance) CameraFeedId() int {
	return int(int(class(self).GetCameraFeedId()))
}

func (self Instance) SetCameraFeedId(value int) {
	class(self).SetCameraFeedId(gd.Int(value))
}

func (self Instance) WhichFeed() classdb.CameraServerFeedImage {
	return classdb.CameraServerFeedImage(class(self).GetWhichFeed())
}

func (self Instance) SetWhichFeed(value classdb.CameraServerFeedImage) {
	class(self).SetWhichFeed(value)
}

func (self Instance) CameraIsActive() bool {
	return bool(class(self).GetCameraActive())
}

func (self Instance) SetCameraIsActive(value bool) {
	class(self).SetCameraActive(value)
}

//go:nosplit
func (self class) SetCameraFeedId(feed_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, feed_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWhichFeed(which_feed classdb.CameraServerFeedImage) {
	var frame = callframe.New()
	callframe.Arg(frame, which_feed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_set_which_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWhichFeed() classdb.CameraServerFeedImage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraServerFeedImage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_get_which_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraActive(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_set_camera_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_get_camera_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCameraTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("CameraTexture", func(ptr gd.Object) any { return classdb.CameraTexture(ptr) })
}
