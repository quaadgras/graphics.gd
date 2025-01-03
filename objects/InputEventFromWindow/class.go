package InputEventFromWindow

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/InputEvent"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
InputEventFromWindow represents events specifically received by windows. This includes mouse events, keyboard events in focused windows or touch screen actions.
*/
type Instance [1]classdb.InputEventFromWindow
type Any interface {
	gd.IsClass
	AsInputEventFromWindow() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventFromWindow

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventFromWindow"))
	return Instance{*(*classdb.InputEventFromWindow)(unsafe.Pointer(&object))}
}

func (self Instance) WindowId() int {
	return int(int(class(self).GetWindowId()))
}

func (self Instance) SetWindowId(value int) {
	class(self).SetWindowId(gd.Int(value))
}

//go:nosplit
func (self class) SetWindowId(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventFromWindow.Bind_set_window_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWindowId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventFromWindow.Bind_get_window_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventFromWindow() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventFromWindow() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}
func init() {
	classdb.Register("InputEventFromWindow", func(ptr gd.Object) any {
		return [1]classdb.InputEventFromWindow{*(*classdb.InputEventFromWindow)(unsafe.Pointer(&ptr))}
	})
}
