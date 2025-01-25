// Package ButtonGroup provides methods for working with ButtonGroup object instances.
package ButtonGroup

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
import "graphics.gd/classdb/Resource"

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
A group of [BaseButton]-derived buttons. The buttons in a [ButtonGroup] are treated like radio buttons: No more than one button can be pressed at a time. Some types of buttons (such as [CheckBox]) may have a special appearance in this state.
Every member of a [ButtonGroup] should have [member BaseButton.toggle_mode] set to [code]true[/code].
*/
type Instance [1]gdclass.ButtonGroup

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsButtonGroup() Instance
}

/*
Returns the current pressed button.
*/
func (self Instance) GetPressedButton() [1]gdclass.BaseButton { //gd:ButtonGroup.get_pressed_button
	return [1]gdclass.BaseButton(class(self).GetPressedButton())
}

/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
func (self Instance) GetButtons() [][1]gdclass.BaseButton { //gd:ButtonGroup.get_buttons
	return [][1]gdclass.BaseButton(gd.ArrayAs[[][1]gdclass.BaseButton](gd.InternalArray(class(self).GetButtons())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ButtonGroup

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ButtonGroup"))
	casted := Instance{*(*gdclass.ButtonGroup)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) GetPressedButton() [1]gdclass.BaseButton { //gd:ButtonGroup.get_pressed_button
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.BaseButton{gd.PointerMustAssertInstanceID[gdclass.BaseButton](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
//go:nosplit
func (self class) GetButtons() Array.Contains[[1]gdclass.BaseButton] { //gd:ButtonGroup.get_buttons
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_get_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.BaseButton]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowUnpress(enabled bool) { //gd:ButtonGroup.set_allow_unpress
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_set_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowUnpress() bool { //gd:ButtonGroup.is_allow_unpress
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ButtonGroup.Bind_is_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPressed(cb func(button [1]gdclass.BaseButton)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsButtonGroup() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButtonGroup() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("ButtonGroup", func(ptr gd.Object) any { return [1]gdclass.ButtonGroup{*(*gdclass.ButtonGroup)(unsafe.Pointer(&ptr))} })
}
