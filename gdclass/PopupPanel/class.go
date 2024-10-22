package PopupPanel

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Popup"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A popup with a configurable panel background. Any child controls added to this node will be stretched to fit the panel's size (similar to how [PanelContainer] works). If you are making windows, see [Window].

*/
type Go [1]classdb.PopupPanel
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PopupPanel
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PopupPanel"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsPopupPanel() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPopupPanel() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPopup() Popup.GD { return *((*Popup.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPopup() Popup.Go { return *((*Popup.Go)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.GD { return *((*Window.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWindow() Window.Go { return *((*Window.Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPopup(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPopup(), name)
	}
}
func init() {classdb.Register("PopupPanel", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
