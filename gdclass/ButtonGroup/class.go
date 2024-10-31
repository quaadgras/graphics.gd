package ButtonGroup

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A group of [BaseButton]-derived buttons. The buttons in a [ButtonGroup] are treated like radio buttons: No more than one button can be pressed at a time. Some types of buttons (such as [CheckBox]) may have a special appearance in this state.
Every member of a [ButtonGroup] should have [member BaseButton.toggle_mode] set to [code]true[/code].
*/
type Instance [1]classdb.ButtonGroup

/*
Returns the current pressed button.
*/
func (self Instance) GetPressedButton() gdclass.BaseButton {
	return gdclass.BaseButton(class(self).GetPressedButton())
}

/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
func (self Instance) GetButtons() gd.Array {
	return gd.Array(class(self).GetButtons())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ButtonGroup

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ButtonGroup"))
	return Instance{classdb.ButtonGroup(object)}
}

func (self Instance) AllowUnpress() bool {
	return bool(class(self).IsAllowUnpress())
}

func (self Instance) SetAllowUnpress(value bool) {
	class(self).SetAllowUnpress(value)
}

/*
Returns the current pressed button.
*/
//go:nosplit
func (self class) GetPressedButton() gdclass.BaseButton {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.BaseButton{classdb.BaseButton(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
//go:nosplit
func (self class) GetButtons() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_get_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowUnpress(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_set_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowUnpress() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_is_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPressed(cb func(button gdclass.BaseButton)) {
	self[0].AsObject().Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsButtonGroup() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButtonGroup() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("ButtonGroup", func(ptr gd.Object) any { return classdb.ButtonGroup(ptr) })
}
