package EditorResourcePicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/HBoxContainer"
import "grow.graphics/gd/gdclass/BoxContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This [Control] node is used in the editor's Inspector dock to allow editing of [Resource] type properties. It provides options for creating, loading, saving and converting resources. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] This [Control] does not include any editor for the resource, as editing is controlled by the Inspector dock itself or sub-Inspectors.
	// EditorResourcePicker methods that can be overridden by a [Class] that extends it.
	type EditorResourcePicker interface {
		//This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
		//[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
		SetCreateOptions(menu_node gd.Object) 
		//This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
		HandleMenuSelected(id int) bool
	}

*/
type Go [1]classdb.EditorResourcePicker

/*
This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
*/
func (Go) _set_create_options(impl func(ptr unsafe.Pointer, menu_node gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var menu_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(menu_node)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, menu_node)
	}
}

/*
This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
*/
func (Go) _handle_menu_selected(impl func(ptr unsafe.Pointer, id int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a list of all allowed types and subtypes corresponding to the [member base_type]. If the [member base_type] is empty, an empty list is returned.
*/
func (self Go) GetAllowedTypes() []string {
	return []string(class(self).GetAllowedTypes().Strings())
}

/*
Sets the toggle mode state for the main button. Works only if [member toggle_mode] is set to [code]true[/code].
*/
func (self Go) SetTogglePressed(pressed bool) {
	class(self).SetTogglePressed(pressed)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorResourcePicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourcePicker"))
	return Go{classdb.EditorResourcePicker(object)}
}

func (self Go) BaseType() string {
		return string(class(self).GetBaseType().String())
}

func (self Go) SetBaseType(value string) {
	class(self).SetBaseType(gd.NewString(value))
}

func (self Go) EditedResource() gdclass.Resource {
		return gdclass.Resource(class(self).GetEditedResource())
}

func (self Go) SetEditedResource(value gdclass.Resource) {
	class(self).SetEditedResource(value)
}

func (self Go) Editable() bool {
		return bool(class(self).IsEditable())
}

func (self Go) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Go) ToggleMode() bool {
		return bool(class(self).IsToggleMode())
}

func (self Go) SetToggleMode(value bool) {
	class(self).SetToggleMode(value)
}

/*
This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
*/
func (class) _set_create_options(impl func(ptr unsafe.Pointer, menu_node gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var menu_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(menu_node)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, menu_node)
	}
}

/*
This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
*/
func (class) _handle_menu_selected(impl func(ptr unsafe.Pointer, id gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetBaseType(base_type gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(base_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaseType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all allowed types and subtypes corresponding to the [member base_type]. If the [member base_type] is empty, an empty list is returned.
*/
//go:nosplit
func (self class) GetAllowedTypes() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_allowed_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditedResource(resource gdclass.Resource)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(resource[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEditedResource() gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetToggleMode(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsToggleMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_is_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the toggle mode state for the main button. Works only if [member toggle_mode] is set to [code]true[/code].
*/
//go:nosplit
func (self class) SetTogglePressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_toggle_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEditable(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnResourceSelected(cb func(resource gdclass.Resource, inspect bool)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnResourceChanged(cb func(resource gdclass.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsEditorResourcePicker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorResourcePicker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsHBoxContainer() HBoxContainer.GD { return *((*HBoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsHBoxContainer() HBoxContainer.Go { return *((*HBoxContainer.Go)(unsafe.Pointer(&self))) }
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
	case "_set_create_options": return reflect.ValueOf(self._set_create_options);
	case "_handle_menu_selected": return reflect.ValueOf(self._handle_menu_selected);
	default: return gd.VirtualByName(self.AsHBoxContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_set_create_options": return reflect.ValueOf(self._set_create_options);
	case "_handle_menu_selected": return reflect.ValueOf(self._handle_menu_selected);
	default: return gd.VirtualByName(self.AsHBoxContainer(), name)
	}
}
func init() {classdb.Register("EditorResourcePicker", func(ptr gd.Object) any { return classdb.EditorResourcePicker(ptr) })}
