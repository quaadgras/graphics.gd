// Package MenuBar provides methods for working with MenuBar object instances.
package MenuBar

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A horizontal menu bar that creates a menu for each [PopupMenu] child. New items are created by adding [PopupMenu]s to this node. Item title is determined by [member Window.title], or node name if [member Window.title] is empty. Item title can be overridden using [method set_menu_title].
*/
type Instance [1]gdclass.MenuBar

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMenuBar() Instance
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Instance) SetDisableShortcuts(disabled bool) { //gd:MenuBar.set_disable_shortcuts
	class(self).SetDisableShortcuts(disabled)
}

/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
func (self Instance) IsNativeMenu() bool { //gd:MenuBar.is_native_menu
	return bool(class(self).IsNativeMenu())
}

/*
Returns number of menu items.
*/
func (self Instance) GetMenuCount() int { //gd:MenuBar.get_menu_count
	return int(int(class(self).GetMenuCount()))
}

/*
Sets menu item title.
*/
func (self Instance) SetMenuTitle(menu int, title string) { //gd:MenuBar.set_menu_title
	class(self).SetMenuTitle(int64(menu), String.New(title))
}

/*
Returns menu item title.
*/
func (self Instance) GetMenuTitle(menu int) string { //gd:MenuBar.get_menu_title
	return string(class(self).GetMenuTitle(int64(menu)).String())
}

/*
Sets menu item tooltip.
*/
func (self Instance) SetMenuTooltip(menu int, tooltip string) { //gd:MenuBar.set_menu_tooltip
	class(self).SetMenuTooltip(int64(menu), String.New(tooltip))
}

/*
Returns menu item tooltip.
*/
func (self Instance) GetMenuTooltip(menu int) string { //gd:MenuBar.get_menu_tooltip
	return string(class(self).GetMenuTooltip(int64(menu)).String())
}

/*
If [code]true[/code], menu item is disabled.
*/
func (self Instance) SetMenuDisabled(menu int, disabled bool) { //gd:MenuBar.set_menu_disabled
	class(self).SetMenuDisabled(int64(menu), disabled)
}

/*
Returns [code]true[/code], if menu item is disabled.
*/
func (self Instance) IsMenuDisabled(menu int) bool { //gd:MenuBar.is_menu_disabled
	return bool(class(self).IsMenuDisabled(int64(menu)))
}

/*
If [code]true[/code], menu item is hidden.
*/
func (self Instance) SetMenuHidden(menu int, hidden bool) { //gd:MenuBar.set_menu_hidden
	class(self).SetMenuHidden(int64(menu), hidden)
}

/*
Returns [code]true[/code], if menu item is hidden.
*/
func (self Instance) IsMenuHidden(menu int) bool { //gd:MenuBar.is_menu_hidden
	return bool(class(self).IsMenuHidden(int64(menu)))
}

/*
Returns [PopupMenu] associated with menu item.
*/
func (self Instance) GetMenuPopup(menu int) [1]gdclass.PopupMenu { //gd:MenuBar.get_menu_popup
	return [1]gdclass.PopupMenu(class(self).GetMenuPopup(int64(menu)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MenuBar

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MenuBar"))
	casted := Instance{*(*gdclass.MenuBar)(unsafe.Pointer(&object))}
	return casted
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
	class(self).SetStartIndex(int64(value))
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

func (self Instance) TextDirection() gdclass.ControlTextDirection {
	return gdclass.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(String.New(value))
}

//go:nosplit
func (self class) SetSwitchOnHover(enable bool) { //gd:MenuBar.set_switch_on_hover
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSwitchOnHover() bool { //gd:MenuBar.is_switch_on_hover
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
//go:nosplit
func (self class) SetDisableShortcuts(disabled bool) { //gd:MenuBar.set_disable_shortcuts
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreferGlobalMenu(enabled bool) { //gd:MenuBar.set_prefer_global_menu
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPreferGlobalMenu() bool { //gd:MenuBar.is_prefer_global_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
//go:nosplit
func (self class) IsNativeMenu() bool { //gd:MenuBar.is_native_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of menu items.
*/
//go:nosplit
func (self class) GetMenuCount() int64 { //gd:MenuBar.get_menu_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.ControlTextDirection) { //gd:MenuBar.set_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.ControlTextDirection { //gd:MenuBar.get_text_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language String.Readable) { //gd:MenuBar.set_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() String.Readable { //gd:MenuBar.get_language
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(enabled bool) { //gd:MenuBar.set_flat
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool { //gd:MenuBar.is_flat
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartIndex(enabled int64) { //gd:MenuBar.set_start_index
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_start_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartIndex() int64 { //gd:MenuBar.get_start_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_start_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets menu item title.
*/
//go:nosplit
func (self class) SetMenuTitle(menu int64, title String.Readable) { //gd:MenuBar.set_menu_title
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns menu item title.
*/
//go:nosplit
func (self class) GetMenuTitle(menu int64) String.Readable { //gd:MenuBar.get_menu_title
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets menu item tooltip.
*/
//go:nosplit
func (self class) SetMenuTooltip(menu int64, tooltip String.Readable) { //gd:MenuBar.set_menu_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns menu item tooltip.
*/
//go:nosplit
func (self class) GetMenuTooltip(menu int64) String.Readable { //gd:MenuBar.get_menu_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
If [code]true[/code], menu item is disabled.
*/
//go:nosplit
func (self class) SetMenuDisabled(menu int64, disabled bool) { //gd:MenuBar.set_menu_disabled
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code], if menu item is disabled.
*/
//go:nosplit
func (self class) IsMenuDisabled(menu int64) bool { //gd:MenuBar.is_menu_disabled
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], menu item is hidden.
*/
//go:nosplit
func (self class) SetMenuHidden(menu int64, hidden bool) { //gd:MenuBar.set_menu_hidden
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_set_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code], if menu item is hidden.
*/
//go:nosplit
func (self class) IsMenuHidden(menu int64) bool { //gd:MenuBar.is_menu_hidden
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_is_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [PopupMenu] associated with menu item.
*/
//go:nosplit
func (self class) GetMenuPopup(menu int64) [1]gdclass.PopupMenu { //gd:MenuBar.get_menu_popup
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuBar.Bind_get_menu_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
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
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("MenuBar", func(ptr gd.Object) any { return [1]gdclass.MenuBar{*(*gdclass.MenuBar)(unsafe.Pointer(&ptr))} })
}
