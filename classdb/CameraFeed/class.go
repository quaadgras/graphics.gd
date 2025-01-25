// Package CameraFeed provides methods for working with CameraFeed object instances.
package CameraFeed

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
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A camera feed gives you access to a single physical camera attached to your device. When enabled, Godot will start capturing frames from the camera which can then be used. See also [CameraServer].
[b]Note:[/b] Many cameras will return YCbCr images which are split into two textures and need to be combined in a shader. Godot does this automatically for you if you set the environment to show the camera image in the background.
*/
type Instance [1]gdclass.CameraFeed

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCameraFeed() Instance
}

/*
Returns the unique ID for this feed.
*/
func (self Instance) GetId() int {
	return int(int(class(self).GetId()))
}

/*
Returns the camera's name.
*/
func (self Instance) GetName() string {
	return string(class(self).GetName().String())
}

/*
Returns the position of camera on the device.
*/
func (self Instance) GetPosition() gdclass.CameraFeedFeedPosition {
	return gdclass.CameraFeedFeedPosition(class(self).GetPosition())
}

/*
Returns feed image data type.
*/
func (self Instance) GetDatatype() gdclass.CameraFeedFeedDataType {
	return gdclass.CameraFeedFeedDataType(class(self).GetDatatype())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CameraFeed

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraFeed"))
	casted := Instance{*(*gdclass.CameraFeed)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) FeedIsActive() bool {
	return bool(class(self).IsActive())
}

func (self Instance) SetFeedIsActive(value bool) {
	class(self).SetActive(value)
}

func (self Instance) FeedTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetTransform())
}

func (self Instance) SetFeedTransform(value Transform2D.OriginXY) {
	class(self).SetTransform(gd.Transform2D(value))
}

/*
Returns the unique ID for this feed.
*/
//go:nosplit
func (self class) GetId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActive(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the camera's name.
*/
//go:nosplit
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the position of camera on the device.
*/
//go:nosplit
func (self class) GetPosition() gdclass.CameraFeedFeedPosition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CameraFeedFeedPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns feed image data type.
*/
//go:nosplit
func (self class) GetDatatype() gdclass.CameraFeedFeedDataType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CameraFeedFeedDataType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_datatype, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraFeed() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCameraFeed() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("CameraFeed", func(ptr gd.Object) any { return [1]gdclass.CameraFeed{*(*gdclass.CameraFeed)(unsafe.Pointer(&ptr))} })
}

type FeedDataType = gdclass.CameraFeedFeedDataType

const (
	/*No image set for the feed.*/
	FeedNoimage FeedDataType = 0
	/*Feed supplies RGB images.*/
	FeedRgb FeedDataType = 1
	/*Feed supplies YCbCr images that need to be converted to RGB.*/
	FeedYcbcr FeedDataType = 2
	/*Feed supplies separate Y and CbCr images that need to be combined and converted to RGB.*/
	FeedYcbcrSep FeedDataType = 3
)

type FeedPosition = gdclass.CameraFeedFeedPosition

const (
	/*Unspecified position.*/
	FeedUnspecified FeedPosition = 0
	/*Camera is mounted at the front of the device.*/
	FeedFront FeedPosition = 1
	/*Camera is mounted at the back of the device.*/
	FeedBack FeedPosition = 2
)
