// Package CameraTexture provides methods for working with CameraTexture object instances.
package CameraTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This texture gives access to the camera texture provided by a [CameraFeed].
[b]Note:[/b] Many cameras supply YCbCr images which need to be converted in a shader.
*/
type Instance [1]gdclass.CameraTexture
type Any interface {
	gd.IsClass
	AsCameraTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CameraTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraTexture"))
	return Instance{*(*gdclass.CameraTexture)(unsafe.Pointer(&object))}
}

func (self Instance) CameraFeedId() int {
	return int(int(class(self).GetCameraFeedId()))
}

func (self Instance) SetCameraFeedId(value int) {
	class(self).SetCameraFeedId(gd.Int(value))
}

func (self Instance) WhichFeed() gdclass.CameraServerFeedImage {
	return gdclass.CameraServerFeedImage(class(self).GetWhichFeed())
}

func (self Instance) SetWhichFeed(value gdclass.CameraServerFeedImage) {
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
func (self class) SetWhichFeed(which_feed gdclass.CameraServerFeedImage) {
	var frame = callframe.New()
	callframe.Arg(frame, which_feed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraTexture.Bind_set_which_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWhichFeed() gdclass.CameraServerFeedImage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CameraServerFeedImage](frame)
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
	gdclass.Register("CameraTexture", func(ptr gd.Object) any {
		return [1]gdclass.CameraTexture{*(*gdclass.CameraTexture)(unsafe.Pointer(&ptr))}
	})
}
