package VBoxContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/BoxContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A variant of [BoxContainer] that can only arrange its child controls vertically. Child controls are rearranged automatically when their minimum size changes.

*/
type Go [1]classdb.VBoxContainer
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VBoxContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VBoxContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsVBoxContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVBoxContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsBoxContainer() BoxContainer.GD { return *((*BoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxContainer() BoxContainer.Go { return *((*BoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBoxContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBoxContainer(), name)
	}
}
func init() {classdb.Register("VBoxContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
