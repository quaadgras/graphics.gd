package ReferenceRect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A rectangle box that displays only a colored border around its rectangle. It is used to visualize the extents of a [Control].
*/
type Instance [1]classdb.ReferenceRect
type Any interface {
	gd.IsClass
	AsReferenceRect() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ReferenceRect

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ReferenceRect"))
	return Instance{*(*classdb.ReferenceRect)(unsafe.Pointer(&object))}
}

func (self Instance) BorderColor() Color.RGBA {
	return Color.RGBA(class(self).GetBorderColor())
}

func (self Instance) SetBorderColor(value Color.RGBA) {
	class(self).SetBorderColor(gd.Color(value))
}

func (self Instance) BorderWidth() Float.X {
	return Float.X(Float.X(class(self).GetBorderWidth()))
}

func (self Instance) SetBorderWidth(value Float.X) {
	class(self).SetBorderWidth(gd.Float(value))
}

func (self Instance) EditorOnly() bool {
	return bool(class(self).GetEditorOnly())
}

func (self Instance) SetEditorOnly(value bool) {
	class(self).SetEditorOnly(value)
}

//go:nosplit
func (self class) GetBorderColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_get_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderWidth(width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_set_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEditorOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_get_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditorOnly(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReferenceRect.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsReferenceRect() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsReferenceRect() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced  { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {
	classdb.Register("ReferenceRect", func(ptr gd.Object) any {
		return [1]classdb.ReferenceRect{*(*classdb.ReferenceRect)(unsafe.Pointer(&ptr))}
	})
}
