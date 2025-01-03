package ScrollBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Range"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class for scrollbars, typically used to navigate through content that extends beyond the visible area of a control. Scrollbars are [Range]-based controls.
*/
type Instance [1]classdb.ScrollBar
type Any interface {
	gd.IsClass
	AsScrollBar() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ScrollBar

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScrollBar"))
	return Instance{*(*classdb.ScrollBar)(unsafe.Pointer(&object))}
}

func (self Instance) CustomStep() Float.X {
	return Float.X(Float.X(class(self).GetCustomStep()))
}

func (self Instance) SetCustomStep(value Float.X) {
	class(self).SetCustomStep(gd.Float(value))
}

//go:nosplit
func (self class) SetCustomStep(step gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, step)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollBar.Bind_set_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollBar.Bind_get_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnScrolling(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("scrolling"), gd.NewCallable(cb), 0)
}

func (self class) AsScrollBar() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScrollBar() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced     { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance  { return *((*Range.Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRange(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRange(), name)
	}
}
func init() {
	classdb.Register("ScrollBar", func(ptr gd.Object) any { return [1]classdb.ScrollBar{*(*classdb.ScrollBar)(unsafe.Pointer(&ptr))} })
}
