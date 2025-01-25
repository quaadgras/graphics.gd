// Package XRTracker provides methods for working with XRTracker object instances.
package XRTracker

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

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This object is the base of all XR trackers.
*/
type Instance [1]gdclass.XRTracker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRTracker() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRTracker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRTracker"))
	casted := Instance{*(*gdclass.XRTracker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Type() gdclass.XRServerTrackerType {
	return gdclass.XRServerTrackerType(class(self).GetTrackerType())
}

func (self Instance) SetType(value gdclass.XRServerTrackerType) {
	class(self).SetTrackerType(value)
}

func (self Instance) Name() string {
	return string(class(self).GetTrackerName().String())
}

func (self Instance) SetName(value string) {
	class(self).SetTrackerName(gd.NewStringName(value))
}

func (self Instance) Description() string {
	return string(class(self).GetTrackerDesc().String())
}

func (self Instance) SetDescription(value string) {
	class(self).SetTrackerDesc(gd.NewString(value))
}

//go:nosplit
func (self class) GetTrackerType() gdclass.XRServerTrackerType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRServerTrackerType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerType(atype gdclass.XRServerTrackerType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrackerName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerName(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrackerDesc() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_desc, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerDesc(description gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(description))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_desc, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsXRTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("XRTracker", func(ptr gd.Object) any { return [1]gdclass.XRTracker{*(*gdclass.XRTracker)(unsafe.Pointer(&ptr))} })
}
