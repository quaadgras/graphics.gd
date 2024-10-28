package XRTracker

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
This object is the base of all XR trackers.

*/
type Go [1]classdb.XRTracker
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRTracker"))
	return Go{classdb.XRTracker(object)}
}

func (self Go) Type() classdb.XRServerTrackerType {
		return classdb.XRServerTrackerType(class(self).GetTrackerType())
}

func (self Go) SetType(value classdb.XRServerTrackerType) {
	class(self).SetTrackerType(value)
}

func (self Go) Name() string {
		return string(class(self).GetTrackerName().String())
}

func (self Go) SetName(value string) {
	class(self).SetTrackerName(gd.NewStringName(value))
}

func (self Go) Description() string {
		return string(class(self).GetTrackerDesc().String())
}

func (self Go) SetDescription(value string) {
	class(self).SetTrackerDesc(gd.NewString(value))
}

//go:nosplit
func (self class) GetTrackerType() classdb.XRServerTrackerType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRServerTrackerType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackerType(atype classdb.XRServerTrackerType)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrackerName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackerName(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrackerDesc() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_get_tracker_desc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackerDesc(description gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(description))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRTracker.Bind_set_tracker_desc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsXRTracker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRTracker() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("XRTracker", func(ptr gd.Object) any { return classdb.XRTracker(ptr) })}
