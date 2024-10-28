package EditorInspector

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ScrollContainer"
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
This is the control that implements property editing in the editor's Settings dialogs, the Inspector dock, etc. To get the [EditorInspector] used in the editor's Inspector dock, use [method EditorInterface.get_inspector].
[EditorInspector] will show properties in the same order as the array returned by [method Object.get_property_list].
If a property's name is path-like (i.e. if it contains forward slashes), [EditorInspector] will create nested sections for "directories" along the path. For example, if a property is named [code]highlighting/gdscript/node_path_color[/code], it will be shown as "Node Path Color" inside the "GDScript" section nested inside the "Highlighting" section.
If a property has [constant PROPERTY_USAGE_GROUP] usage, it will group subsequent properties whose name starts with the property's hint string. The group ends when a property does not start with that hint string or when a new group starts. An empty group name effectively ends the current group. [EditorInspector] will create a top-level section for each group. For example, if a property with group usage is named [code]Collide With[/code] and its hint string is [code]collide_with_[/code], a subsequent [code]collide_with_area[/code] property will be shown as "Area" inside the "Collide With" section. There is also a special case: when the hint string contains the name of a property, that property is grouped too. This is mainly to help grouping properties like [code]font[/code], [code]font_color[/code] and [code]font_size[/code] (using the hint string [code]font_[/code]).
If a property has [constant PROPERTY_USAGE_SUBGROUP] usage, a subgroup will be created in the same way as a group, and a second-level section will be created for each subgroup.
[b]Note:[/b] Unlike sections created from path-like property names, [EditorInspector] won't capitalize the name for sections created from groups. So properties with group usage usually use capitalized names instead of snake_cased names.

*/
type Go [1]classdb.EditorInspector

/*
Gets the path of the currently selected property.
*/
func (self Go) GetSelectedPath() string {
	return string(class(self).GetSelectedPath().String())
}

/*
Returns the object currently selected in this inspector.
*/
func (self Go) GetEditedObject() gd.Object {
	return gd.Object(class(self).GetEditedObject())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorInspector
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorInspector"))
	return Go{classdb.EditorInspector(object)}
}

/*
Gets the path of the currently selected property.
*/
//go:nosplit
func (self class) GetSelectedPath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspector.Bind_get_selected_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the object currently selected in this inspector.
*/
//go:nosplit
func (self class) GetEditedObject() gd.Object {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspector.Bind_get_edited_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnPropertySelected(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyKeyed(cb func(property string, value gd.Variant, advance bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyDeleted(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_deleted"), gd.NewCallable(cb), 0)
}


func (self Go) OnResourceSelected(cb func(resource gdclass.Resource, path string)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnObjectIdSelected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("object_id_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyEdited(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_edited"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyToggled(cb func(property string, checked bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_toggled"), gd.NewCallable(cb), 0)
}


func (self Go) OnEditedObjectChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("edited_object_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnRestartRequested(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("restart_requested"), gd.NewCallable(cb), 0)
}


func (self class) AsEditorInspector() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorInspector() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsScrollContainer() ScrollContainer.GD { return *((*ScrollContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsScrollContainer() ScrollContainer.Go { return *((*ScrollContainer.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsScrollContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsScrollContainer(), name)
	}
}
func init() {classdb.Register("EditorInspector", func(ptr gd.Object) any { return classdb.EditorInspector(ptr) })}
