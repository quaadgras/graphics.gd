package EditorResourcePicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/HBoxContainer"
import "grow.graphics/gd/object/BoxContainer"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This [Control] node is used in the editor's Inspector dock to allow editing of [Resource] type properties. It provides options for creating, loading, saving and converting resources. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] This [Control] does not include any editor for the resource, as editing is controlled by the Inspector dock itself or sub-Inspectors.
	// EditorResourcePicker methods that can be overridden by a [Class] that extends it.
	type EditorResourcePicker interface {
		//This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
		//[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
		SetCreateOptions(menu_node gd.Object) 
		//This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
		HandleMenuSelected(id gd.Int) bool
	}

*/
type Simple [1]classdb.EditorResourcePicker
func (Simple) _set_create_options(impl func(ptr unsafe.Pointer, menu_node gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var menu_node gd.Object
		menu_node.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, menu_node)
		gc.End()
	}
}
func (Simple) _handle_menu_selected(impl func(ptr unsafe.Pointer, id int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Simple) SetBaseType(base_type string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBaseType(gc.String(base_type))
}
func (self Simple) GetBaseType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBaseType(gc).String())
}
func (self Simple) GetAllowedTypes() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetAllowedTypes(gc))
}
func (self Simple) SetEditedResource(resource [1]classdb.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditedResource(resource)
}
func (self Simple) GetEditedResource() [1]classdb.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Resource(Expert(self).GetEditedResource(gc))
}
func (self Simple) SetToggleMode(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetToggleMode(enable)
}
func (self Simple) IsToggleMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsToggleMode())
}
func (self Simple) SetTogglePressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTogglePressed(pressed)
}
func (self Simple) SetEditable(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditable(enable)
}
func (self Simple) IsEditable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditable())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorResourcePicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
*/
func (class) _set_create_options(impl func(ptr unsafe.Pointer, menu_node gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var menu_node gd.Object
		menu_node.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, menu_node)
		ctx.End()
	}
}

/*
This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
*/
func (class) _handle_menu_selected(impl func(ptr unsafe.Pointer, id gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetBaseType(base_type gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_set_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaseType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_get_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all allowed types and subtypes corresponding to the [member base_type]. If the [member base_type] is empty, an empty list is returned.
*/
//go:nosplit
func (self class) GetAllowedTypes(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_get_allowed_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditedResource(resource object.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_set_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEditedResource(ctx gd.Lifetime) object.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_get_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetToggleMode(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_set_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsToggleMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_is_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the toggle mode state for the main button. Works only if [member toggle_mode] is set to [code]true[/code].
*/
//go:nosplit
func (self class) SetTogglePressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_set_toggle_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEditable(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePicker.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorResourcePicker() Expert { return self[0].AsEditorResourcePicker() }


//go:nosplit
func (self Simple) AsEditorResourcePicker() Simple { return self[0].AsEditorResourcePicker() }


//go:nosplit
func (self class) AsHBoxContainer() HBoxContainer.Expert { return self[0].AsHBoxContainer() }


//go:nosplit
func (self Simple) AsHBoxContainer() HBoxContainer.Simple { return self[0].AsHBoxContainer() }


//go:nosplit
func (self class) AsBoxContainer() BoxContainer.Expert { return self[0].AsBoxContainer() }


//go:nosplit
func (self Simple) AsBoxContainer() BoxContainer.Simple { return self[0].AsBoxContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


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
	case "_set_create_options": return reflect.ValueOf(self._set_create_options);
	case "_handle_menu_selected": return reflect.ValueOf(self._handle_menu_selected);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_set_create_options": return reflect.ValueOf(self._set_create_options);
	case "_handle_menu_selected": return reflect.ValueOf(self._handle_menu_selected);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorResourcePicker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
