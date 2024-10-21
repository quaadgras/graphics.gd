package ConfirmationDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AcceptDialog"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A dialog used for confirmation of actions. This window is similar to [AcceptDialog], but pressing its Cancel button can have a different outcome from pressing the OK button. The order of the two buttons varies depending on the host OS.
To get cancel action, you can use:
[codeblocks]
[gdscript]
get_cancel_button().pressed.connect(_on_canceled)
[/gdscript]
[csharp]
GetCancelButton().Pressed += OnCanceled;
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.ConfirmationDialog
func (self Simple) GetCancelButton() [1]classdb.Button {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Button(Expert(self).GetCancelButton(gc))
}
func (self Simple) SetCancelButtonText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCancelButtonText(gc.String(text))
}
func (self Simple) GetCancelButtonText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCancelButtonText(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ConfirmationDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the cancel button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetCancelButton(ctx gd.Lifetime) object.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfirmationDialog.Bind_get_cancel_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCancelButtonText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfirmationDialog.Bind_set_cancel_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCancelButtonText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfirmationDialog.Bind_get_cancel_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsConfirmationDialog() Expert { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self Simple) AsConfirmationDialog() Simple { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self class) AsAcceptDialog() AcceptDialog.Expert { return self[0].AsAcceptDialog() }


//go:nosplit
func (self Simple) AsAcceptDialog() AcceptDialog.Simple { return self[0].AsAcceptDialog() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


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
func init() {classdb.Register("ConfirmationDialog", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
