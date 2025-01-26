// Package InputEventScreenDrag provides methods for working with InputEventScreenDrag object instances.
package InputEventScreenDrag

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/InputEventFromWindow"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
Stores information about screen drag events. See [method Node._input].
*/
type Instance [1]gdclass.InputEventScreenDrag

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventScreenDrag() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventScreenDrag

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventScreenDrag"))
	casted := Instance{*(*gdclass.InputEventScreenDrag)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Index() int {
	return int(int(class(self).GetIndex()))
}

func (self Instance) SetIndex(value int) {
	class(self).SetIndex(gd.Int(value))
}

func (self Instance) Tilt() Vector2.XY {
	return Vector2.XY(class(self).GetTilt())
}

func (self Instance) SetTilt(value Vector2.XY) {
	class(self).SetTilt(gd.Vector2(value))
}

func (self Instance) Pressure() Float.X {
	return Float.X(Float.X(class(self).GetPressure()))
}

func (self Instance) SetPressure(value Float.X) {
	class(self).SetPressure(gd.Float(value))
}

func (self Instance) PenInverted() bool {
	return bool(class(self).GetPenInverted())
}

func (self Instance) SetPenInverted(value bool) {
	class(self).SetPenInverted(value)
}

func (self Instance) Position() Vector2.XY {
	return Vector2.XY(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector2.XY) {
	class(self).SetPosition(gd.Vector2(value))
}

func (self Instance) Relative() Vector2.XY {
	return Vector2.XY(class(self).GetRelative())
}

func (self Instance) SetRelative(value Vector2.XY) {
	class(self).SetRelative(gd.Vector2(value))
}

func (self Instance) ScreenRelative() Vector2.XY {
	return Vector2.XY(class(self).GetScreenRelative())
}

func (self Instance) SetScreenRelative(value Vector2.XY) {
	class(self).SetScreenRelative(gd.Vector2(value))
}

func (self Instance) Velocity() Vector2.XY {
	return Vector2.XY(class(self).GetVelocity())
}

func (self Instance) SetVelocity(value Vector2.XY) {
	class(self).SetVelocity(gd.Vector2(value))
}

func (self Instance) ScreenVelocity() Vector2.XY {
	return Vector2.XY(class(self).GetScreenVelocity())
}

func (self Instance) SetScreenVelocity(value Vector2.XY) {
	class(self).SetScreenVelocity(gd.Vector2(value))
}

//go:nosplit
func (self class) SetIndex(index gd.Int) { //gd:InputEventScreenDrag.set_index
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIndex() gd.Int { //gd:InputEventScreenDrag.get_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTilt(tilt gd.Vector2) { //gd:InputEventScreenDrag.set_tilt
	var frame = callframe.New()
	callframe.Arg(frame, tilt)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_tilt, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTilt() gd.Vector2 { //gd:InputEventScreenDrag.get_tilt
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_tilt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressure(pressure gd.Float) { //gd:InputEventScreenDrag.set_pressure
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressure() gd.Float { //gd:InputEventScreenDrag.get_pressure
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPenInverted(pen_inverted bool) { //gd:InputEventScreenDrag.set_pen_inverted
	var frame = callframe.New()
	callframe.Arg(frame, pen_inverted)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPenInverted() bool { //gd:InputEventScreenDrag.get_pen_inverted
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector2) { //gd:InputEventScreenDrag.set_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector2 { //gd:InputEventScreenDrag.get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRelative(relative gd.Vector2) { //gd:InputEventScreenDrag.set_relative
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRelative() gd.Vector2 { //gd:InputEventScreenDrag.get_relative
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenRelative(relative gd.Vector2) { //gd:InputEventScreenDrag.set_screen_relative
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_screen_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenRelative() gd.Vector2 { //gd:InputEventScreenDrag.get_screen_relative
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_screen_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2) { //gd:InputEventScreenDrag.set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Vector2 { //gd:InputEventScreenDrag.get_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenVelocity(velocity gd.Vector2) { //gd:InputEventScreenDrag.set_screen_velocity
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenVelocity() gd.Vector2 { //gd:InputEventScreenDrag.get_screen_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventScreenDrag() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventScreenDrag() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEventFromWindow() InputEventFromWindow.Advanced {
	return *((*InputEventFromWindow.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventFromWindow() InputEventFromWindow.Instance {
	return *((*InputEventFromWindow.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventFromWindow.Advanced(self.AsInputEventFromWindow()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventFromWindow.Instance(self.AsInputEventFromWindow()), name)
	}
}
func init() {
	gdclass.Register("InputEventScreenDrag", func(ptr gd.Object) any {
		return [1]gdclass.InputEventScreenDrag{*(*gdclass.InputEventScreenDrag)(unsafe.Pointer(&ptr))}
	})
}
