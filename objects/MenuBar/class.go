package MenuBar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A horizontal menu bar that creates a [MenuButton] for each [PopupMenu] child. New items are created by adding [PopupMenu]s to this node.
*/
type Instance [1]classdb.MenuBar

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Instance) SetDisableShortcuts(disabled bool) {
	class(self).SetDisableShortcuts(disabled)
}

/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
func (self Instance) IsNativeMenu() bool {
	return bool(class(self).IsNativeMenu())
}

/*
Returns number of menu items.
*/
func (self Instance) GetMenuCount() int {
	return int(int(class(self).GetMenuCount()))
}

/*
Sets menu item title.
*/
func (self Instance) SetMenuTitle(menu int, title string) {
	class(self).SetMenuTitle(gd.Int(menu), gd.NewString(title))
}

/*
Returns menu item title.
*/
func (self Instance) GetMenuTitle(menu int) string {
	return string(class(self).GetMenuTitle(gd.Int(menu)).String())
}

/*
Sets menu item tooltip.
*/
func (self Instance) SetMenuTooltip(menu int, tooltip string) {
	class(self).SetMenuTooltip(gd.Int(menu), gd.NewString(tooltip))
}

/*
Returns menu item tooltip.
*/
func (self Instance) GetMenuTooltip(menu int) string {
	return string(class(self).GetMenuTooltip(gd.Int(menu)).String())
}

/*
If [code]true[/code], menu item is disabled.
*/
func (self Instance) SetMenuDisabled(menu int, disabled bool) {
	class(self).SetMenuDisabled(gd.Int(menu), disabled)
}

/*
Returns [code]true[/code], if menu item is disabled.
*/
func (self Instance) IsMenuDisabled(menu int) bool {
	return bool(class(self).IsMenuDisabled(gd.Int(menu)))
}

/*
If [code]true[/code], menu item is hidden.
*/
func (self Instance) SetMenuHidden(menu int, hidden bool) {
	class(self).SetMenuHidden(gd.Int(menu), hidden)
}

/*
Returns [code]true[/code], if menu item is hidden.
*/
func (self Instance) IsMenuHidden(menu int) bool {
	return bool(class(self).IsMenuHidden(gd.Int(menu)))
}

/*
Returns [PopupMenu] associated with menu item.
*/
func (self Instance) GetMenuPopup(menu int) objects.PopupMenu {
	return objects.PopupMenu(class(self).GetMenuPopup(gd.Int(menu)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MenuBar

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MenuBar"))
	return Instance{classdb.MenuBar(object)}
}

func (self Instance) Flat() bool {
	return bool(class(self).IsFlat())
}

func (self Instance) SetFlat(value bool) {
	class(self).SetFlat(value)
}

func (self Instance) StartIndex() int {
	return int(int(class(self).GetStartIndex()))
}

func (self Instance) SetStartIndex(value int) {
	class(self).SetStartIndex(gd.Int(value))
}

func (self Instance) SwitchOnHover() bool {
	return bool(class(self).IsSwitchOnHover())
}

func (self Instance) SetSwitchOnHover(value bool) {
	class(self).SetSwitchOnHover(value)
}

func (self Instance) PreferGlobalMenu() bool {
	return bool(class(self).IsPreferGlobalMenu())
}

func (self Instance) SetPreferGlobalMenu(value bool) {
	class(self).SetPreferGlobalMenu(value)
}

func (self Instance) TextDirection() classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value classdb.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

//go:nosplit
func (self class) SetSwitchOnHover(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSwitchOnHover() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
//go:nosplit
func (self class) SetDisableShortcuts(disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreferGlobalMenu(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPreferGlobalMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
//go:nosplit
func (self class) IsNativeMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of menu items.
*/
//go:nosplit
func (self class) GetMenuCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartIndex(enabled gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_start_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_start_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets menu item title.
*/
//go:nosplit
func (self class) SetMenuTitle(menu gd.Int, title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, pointers.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns menu item title.
*/
//go:nosplit
func (self class) GetMenuTitle(menu gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets menu item tooltip.
*/
//go:nosplit
func (self class) SetMenuTooltip(menu gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns menu item tooltip.
*/
//go:nosplit
func (self class) GetMenuTooltip(menu gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
If [code]true[/code], menu item is disabled.
*/
//go:nosplit
func (self class) SetMenuDisabled(menu gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code], if menu item is disabled.
*/
//go:nosplit
func (self class) IsMenuDisabled(menu gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], menu item is hidden.
*/
//go:nosplit
func (self class) SetMenuHidden(menu gd.Int, hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code], if menu item is hidden.
*/
//go:nosplit
func (self class) IsMenuHidden(menu gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [PopupMenu] associated with menu item.
*/
//go:nosplit
func (self class) GetMenuPopup(menu gd.Int) objects.PopupMenu {
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PopupMenu{classdb.PopupMenu(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsMenuBar() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMenuBar() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() { classdb.Register("MenuBar", func(ptr gd.Object) any { return classdb.MenuBar(ptr) }) }
