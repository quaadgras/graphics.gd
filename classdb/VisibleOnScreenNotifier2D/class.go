// Package VisibleOnScreenNotifier2D provides methods for working with VisibleOnScreenNotifier2D object instances.
package VisibleOnScreenNotifier2D

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
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Rect2"

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
var _ String.Readable
var _ Path.ToNode

/*
[VisibleOnScreenNotifier2D] represents a rectangular region of 2D space. When any part of this region becomes visible on screen or in a viewport, it will emit a [signal screen_entered] signal, and likewise it will emit a [signal screen_exited] signal when no part of it remains visible.
If you want a node to be enabled automatically when this region is visible on screen, use [VisibleOnScreenEnabler2D].
[b]Note:[/b] [VisibleOnScreenNotifier2D] uses the render culling code to determine whether it's visible on screen, so it won't function unless [member CanvasItem.visible] is set to [code]true[/code].
*/
type Instance [1]gdclass.VisibleOnScreenNotifier2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisibleOnScreenNotifier2D() Instance
}

/*
If [code]true[/code], the bounding rectangle is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier2D]'s visibility to be determined once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated, before the draw pass.
*/
func (self Instance) IsOnScreen() bool { //gd:VisibleOnScreenNotifier2D.is_on_screen
	return bool(class(self).IsOnScreen())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisibleOnScreenNotifier2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisibleOnScreenNotifier2D"))
	casted := Instance{*(*gdclass.VisibleOnScreenNotifier2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Rect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRect())
}

func (self Instance) SetRect(value Rect2.PositionSize) {
	class(self).SetRect(gd.Rect2(value))
}

//go:nosplit
func (self class) SetRect(rect gd.Rect2) { //gd:VisibleOnScreenNotifier2D.set_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenNotifier2D.Bind_set_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRect() gd.Rect2 { //gd:VisibleOnScreenNotifier2D.get_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenNotifier2D.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the bounding rectangle is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier2D]'s visibility to be determined once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated, before the draw pass.
*/
//go:nosplit
func (self class) IsOnScreen() bool { //gd:VisibleOnScreenNotifier2D.is_on_screen
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenNotifier2D.Bind_is_on_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnScreenEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("screen_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScreenExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("screen_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsVisibleOnScreenNotifier2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisibleOnScreenNotifier2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("VisibleOnScreenNotifier2D", func(ptr gd.Object) any {
		return [1]gdclass.VisibleOnScreenNotifier2D{*(*gdclass.VisibleOnScreenNotifier2D)(unsafe.Pointer(&ptr))}
	})
}
