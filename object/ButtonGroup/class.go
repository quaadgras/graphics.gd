package ButtonGroup

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A group of [BaseButton]-derived buttons. The buttons in a [ButtonGroup] are treated like radio buttons: No more than one button can be pressed at a time. Some types of buttons (such as [CheckBox]) may have a special appearance in this state.
Every member of a [ButtonGroup] should have [member BaseButton.toggle_mode] set to [code]true[/code].

*/
type Simple [1]classdb.ButtonGroup
func (self Simple) GetPressedButton() [1]classdb.BaseButton {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.BaseButton(Expert(self).GetPressedButton(gc))
}
func (self Simple) GetButtons() gd.ArrayOf[[1]classdb.BaseButton] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.BaseButton](Expert(self).GetButtons(gc))
}
func (self Simple) SetAllowUnpress(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowUnpress(enabled)
}
func (self Simple) IsAllowUnpress() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAllowUnpress())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ButtonGroup
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the current pressed button.
*/
//go:nosplit
func (self class) GetPressedButton(ctx gd.Lifetime) object.BaseButton {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.BaseButton
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Array] of [Button]s who have this as their [ButtonGroup] (see [member BaseButton.button_group]).
*/
//go:nosplit
func (self class) GetButtons(ctx gd.Lifetime) gd.ArrayOf[object.BaseButton] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ButtonGroup.Bind_get_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.BaseButton](ret)
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

//go:nosplit
func (self class) AsButtonGroup() Expert { return self[0].AsButtonGroup() }


//go:nosplit
func (self Simple) AsButtonGroup() Simple { return self[0].AsButtonGroup() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ButtonGroup", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
