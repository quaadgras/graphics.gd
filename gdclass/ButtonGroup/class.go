package ButtonGroup

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A group of [BaseButton]-derived buttons. The buttons in a [ButtonGroup] are treated like radio buttons: No more than one button can be pressed at a time. Some types of buttons (such as [CheckBox]) may have a special appearance in this state.
Every member of a [ButtonGroup] should have [member BaseButton.toggle_mode] set to [code]true[/code].

*/
type Go [1]classdb.ButtonGroup

/*
Returns the current pressed button.
*/
func (self Go) GetPressedButton() gdclass.BaseButton {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.BaseButton(class(self).GetPressedButton(gc))
}

/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
func (self Go) GetButtons() gd.ArrayOf[gdclass.BaseButton] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.BaseButton](class(self).GetButtons(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ButtonGroup
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ButtonGroup"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) AllowUnpress() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAllowUnpress())
}

func (self Go) SetAllowUnpress(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAllowUnpress(value)
}

/*
Returns the current pressed button.
*/
//go:nosplit
func (self class) GetPressedButton(ctx gd.Lifetime) gdclass.BaseButton {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.BaseButton
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
//go:nosplit
func (self class) GetButtons(ctx gd.Lifetime) gd.ArrayOf[gdclass.BaseButton] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_get_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.BaseButton](ret)
}
//go:nosplit
func (self class) SetAllowUnpress(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_set_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowUnpress() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_is_allow_unpress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnPressed(cb func(button gdclass.BaseButton)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pressed"), gc.Callable(cb), 0)
}


func (self class) AsButtonGroup() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsButtonGroup() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("ButtonGroup", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
