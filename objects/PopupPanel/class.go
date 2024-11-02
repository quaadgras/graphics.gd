package PopupPanel

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Popup"
import "grow.graphics/gd/objects/Window"
import "grow.graphics/gd/objects/Viewport"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A popup with a configurable panel background. Any child controls added to this node will be stretched to fit the panel's size (similar to how [PanelContainer] works). If you are making windows, see [Window].
*/
type Instance [1]classdb.PopupPanel

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PopupPanel

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PopupPanel"))
	return Instance{classdb.PopupPanel(object)}
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
	classdb.Register("PopupPanel", func(ptr gd.Object) any { return classdb.PopupPanel(ptr) })
}
