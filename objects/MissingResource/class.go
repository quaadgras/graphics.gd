package MissingResource

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This is an internal editor class intended for keeping data of resources of unknown type (most likely this type was supplied by an extension that is no longer loaded). It can't be manually instantiated or placed in a scene.
[b]Warning:[/b] Ignore missing resources unless you know what you are doing. Existing properties on a missing resource can be freely modified in code, regardless of the type they are intended to be.
*/
type Instance [1]classdb.MissingResource
type Any interface {
	gd.IsClass
	AsMissingResource() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MissingResource

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MissingResource"))
	return Instance{*(*classdb.MissingResource)(unsafe.Pointer(&object))}
}

func (self Instance) OriginalClass() string {
	return string(class(self).GetOriginalClass().String())
}

func (self Instance) SetOriginalClass(value string) {
	class(self).SetOriginalClass(gd.NewString(value))
}

func (self Instance) RecordingProperties() bool {
	return bool(class(self).IsRecordingProperties())
}

func (self Instance) SetRecordingProperties(value bool) {
	class(self).SetRecordingProperties(value)
}

//go:nosplit
func (self class) SetOriginalClass(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingResource.Bind_set_original_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOriginalClass() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingResource.Bind_get_original_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRecordingProperties(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingResource.Bind_set_recording_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRecordingProperties() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingResource.Bind_is_recording_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsMissingResource() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMissingResource() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("MissingResource", func(ptr gd.Object) any {
		return [1]classdb.MissingResource{*(*classdb.MissingResource)(unsafe.Pointer(&ptr))}
	})
}
