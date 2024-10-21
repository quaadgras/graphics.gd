package InputEventMagnifyGesture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/InputEventGesture"
import "grow.graphics/gd/object/InputEventWithModifiers"
import "grow.graphics/gd/object/InputEventFromWindow"
import "grow.graphics/gd/object/InputEvent"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Stores the factor of a magnifying touch gesture. This is usually performed when the user pinches the touch screen and used for zooming in/out.
[b]Note:[/b] On Android, this requires the [member ProjectSettings.input_devices/pointing/android/enable_pan_and_scale_gestures] project setting to be enabled.

*/
type Simple [1]classdb.InputEventMagnifyGesture
func (self Simple) SetFactor(factor float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFactor(gd.Float(factor))
}
func (self Simple) GetFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFactor()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputEventMagnifyGesture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFactor(factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMagnifyGesture.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMagnifyGesture.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsInputEventMagnifyGesture() Expert { return self[0].AsInputEventMagnifyGesture() }


//go:nosplit
func (self Simple) AsInputEventMagnifyGesture() Simple { return self[0].AsInputEventMagnifyGesture() }


//go:nosplit
func (self class) AsInputEventGesture() InputEventGesture.Expert { return self[0].AsInputEventGesture() }


//go:nosplit
func (self Simple) AsInputEventGesture() InputEventGesture.Simple { return self[0].AsInputEventGesture() }


//go:nosplit
func (self class) AsInputEventWithModifiers() InputEventWithModifiers.Expert { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self Simple) AsInputEventWithModifiers() InputEventWithModifiers.Simple { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self class) AsInputEventFromWindow() InputEventFromWindow.Expert { return self[0].AsInputEventFromWindow() }


//go:nosplit
func (self Simple) AsInputEventFromWindow() InputEventFromWindow.Simple { return self[0].AsInputEventFromWindow() }


//go:nosplit
func (self class) AsInputEvent() InputEvent.Expert { return self[0].AsInputEvent() }


//go:nosplit
func (self Simple) AsInputEvent() InputEvent.Simple { return self[0].AsInputEvent() }


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
func init() {classdb.Register("InputEventMagnifyGesture", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
