// Package PopupPanel provides methods for working with PopupPanel object instances.
package PopupPanel

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Popup"
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A popup with a configurable panel background. Any child controls added to this node will be stretched to fit the panel's size (similar to how [PanelContainer] works). If you are making windows, see [Window].
*/
type Instance [1]gdclass.PopupPanel
type Any interface {
	gd.IsClass
	AsPopupPanel() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PopupPanel

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PopupPanel"))
	return Instance{*(*gdclass.PopupPanel)(unsafe.Pointer(&object))}
}

func (self class) AsPopupPanel() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopupPanel() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPopup() Popup.Advanced      { return *((*Popup.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopup() Popup.Instance   { return *((*Popup.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsPopup(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPopup(), name)
	}
}
func init() {
	gdclass.Register("PopupPanel", func(ptr gd.Object) any { return [1]gdclass.PopupPanel{*(*gdclass.PopupPanel)(unsafe.Pointer(&ptr))} })
}
