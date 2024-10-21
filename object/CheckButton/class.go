package CheckButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Button"
import "grow.graphics/gd/object/BaseButton"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[CheckButton] is a toggle button displayed as a check field. It's similar to [CheckBox] in functionality, but it has a different appearance. To follow established UX patterns, it's recommended to use [CheckButton] when toggling it has an [b]immediate[/b] effect on something. For example, it can be used when pressing it shows or hides advanced settings, without asking the user to confirm this action.
See also [BaseButton] which contains common properties and methods associated with this node.

*/
type Simple [1]classdb.CheckButton
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CheckButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsCheckButton() Expert { return self[0].AsCheckButton() }


//go:nosplit
func (self Simple) AsCheckButton() Simple { return self[0].AsCheckButton() }


//go:nosplit
func (self class) AsButton() Button.Expert { return self[0].AsButton() }


//go:nosplit
func (self Simple) AsButton() Button.Simple { return self[0].AsButton() }


//go:nosplit
func (self class) AsBaseButton() BaseButton.Expert { return self[0].AsBaseButton() }


//go:nosplit
func (self Simple) AsBaseButton() BaseButton.Simple { return self[0].AsBaseButton() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CheckButton", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
