package CameraFeed

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A camera feed gives you access to a single physical camera attached to your device. When enabled, Godot will start capturing frames from the camera which can then be used. See also [CameraServer].
[b]Note:[/b] Many cameras will return YCbCr images which are split into two textures and need to be combined in a shader. Godot does this automatically for you if you set the environment to show the camera image in the background.

*/
type Go [1]classdb.CameraFeed

/*
Returns the unique ID for this feed.
*/
func (self Go) GetId() int {
	return int(int(class(self).GetId()))
}

/*
Returns the camera's name.
*/
func (self Go) GetName() string {
	return string(class(self).GetName().String())
}

/*
Returns the position of camera on the device.
*/
func (self Go) GetPosition() classdb.CameraFeedFeedPosition {
	return classdb.CameraFeedFeedPosition(class(self).GetPosition())
}

/*
Returns feed image data type.
*/
func (self Go) GetDatatype() classdb.CameraFeedFeedDataType {
	return classdb.CameraFeedFeedDataType(class(self).GetDatatype())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CameraFeed
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraFeed"))
	return Go{classdb.CameraFeed(object)}
}

func (self Go) FeedIsActive() bool {
		return bool(class(self).IsActive())
}

func (self Go) SetFeedIsActive(value bool) {
	class(self).SetActive(value)
}

func (self Go) FeedTransform() gd.Transform2D {
		return gd.Transform2D(class(self).GetTransform())
}

func (self Go) SetFeedTransform(value gd.Transform2D) {
	class(self).SetTransform(value)
}

/*
Returns the unique ID for this feed.
*/
//go:nosplit
func (self class) GetId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActive(active bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the camera's name.
*/
//go:nosplit
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the position of camera on the device.
*/
//go:nosplit
func (self class) GetPosition() classdb.CameraFeedFeedPosition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraFeedFeedPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransform(transform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns feed image data type.
*/
//go:nosplit
func (self class) GetDatatype() classdb.CameraFeedFeedDataType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraFeedFeedDataType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_datatype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraFeed() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCameraFeed() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("CameraFeed", func(ptr gd.Object) any { return classdb.CameraFeed(ptr) })}
type FeedDataType = classdb.CameraFeedFeedDataType

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
type FeedPosition = classdb.CameraFeedFeedPosition

const (
/*Unspecified position.*/
	FeedUnspecified FeedPosition = 0
/*Camera is mounted at the front of the device.*/
	FeedFront FeedPosition = 1
/*Camera is mounted at the back of the device.*/
	FeedBack FeedPosition = 2
)
