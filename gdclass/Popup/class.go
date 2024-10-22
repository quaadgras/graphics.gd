package Popup

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Popup] is a base class for contextual windows and panels with fixed position. It's a modal by default (see [member Window.popup_window]) and provides methods for implementing custom popup behavior.

*/
type Go [1]classdb.Popup
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Popup
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Popup"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) OnPopupHide(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("popup_hide"), gc.Callable(cb), 0)
}


func (self class) AsPopup() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPopup() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.GD { return *((*Window.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWindow() Window.Go { return *((*Window.Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsWindow(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsWindow(), name)
	}
}
func init() {classdb.Register("Popup", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
