// Package Popup provides methods for working with Popup object instances.
package Popup

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
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

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
[Popup] is a base class for contextual windows and panels with fixed position. It's a modal by default (see [member Window.popup_window]) and provides methods for implementing custom popup behavior.
*/
type Instance [1]gdclass.Popup

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPopup() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Popup

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Popup"))
	casted := Instance{*(*gdclass.Popup)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) OnPopupHide(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("popup_hide"), gd.NewCallable(cb), 0)
}

func (self class) AsPopup() Advanced            { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopup() Instance         { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Window.Advanced(self.AsWindow()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Window.Instance(self.AsWindow()), name)
	}
}
func init() {
	gdclass.Register("Popup", func(ptr gd.Object) any { return [1]gdclass.Popup{*(*gdclass.Popup)(unsafe.Pointer(&ptr))} })
}
