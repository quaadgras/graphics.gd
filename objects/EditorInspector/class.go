package EditorInspector

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/ScrollContainer"
import "grow.graphics/gd/objects/Container"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This is the control that implements property editing in the editor's Settings dialogs, the Inspector dock, etc. To get the [EditorInspector] used in the editor's Inspector dock, use [method EditorInterface.get_inspector].
[EditorInspector] will show properties in the same order as the array returned by [method Object.get_property_list].
If a property's name is path-like (i.e. if it contains forward slashes), [EditorInspector] will create nested sections for "directories" along the path. For example, if a property is named [code]highlighting/gdscript/node_path_color[/code], it will be shown as "Node Path Color" inside the "GDScript" section nested inside the "Highlighting" section.
If a property has [constant PROPERTY_USAGE_GROUP] usage, it will group subsequent properties whose name starts with the property's hint string. The group ends when a property does not start with that hint string or when a new group starts. An empty group name effectively ends the current group. [EditorInspector] will create a top-level section for each group. For example, if a property with group usage is named [code]Collide With[/code] and its hint string is [code]collide_with_[/code], a subsequent [code]collide_with_area[/code] property will be shown as "Area" inside the "Collide With" section. There is also a special case: when the hint string contains the name of a property, that property is grouped too. This is mainly to help grouping properties like [code]font[/code], [code]font_color[/code] and [code]font_size[/code] (using the hint string [code]font_[/code]).
If a property has [constant PROPERTY_USAGE_SUBGROUP] usage, a subgroup will be created in the same way as a group, and a second-level section will be created for each subgroup.
[b]Note:[/b] Unlike sections created from path-like property names, [EditorInspector] won't capitalize the name for sections created from groups. So properties with group usage usually use capitalized names instead of snake_cased names.
*/
type Instance [1]classdb.EditorInspector

/*
Gets the path of the currently selected property.
*/
func (self Instance) GetSelectedPath() string {
	return string(class(self).GetSelectedPath().String())
}

/*
Returns the object currently selected in this inspector.
*/
func (self Instance) GetEditedObject() gd.Object {
	return gd.Object(class(self).GetEditedObject())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorInspector

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorInspector"))
	return Instance{classdb.EditorInspector(object)}
}

/*
Gets the path of the currently selected property.
*/
//go:nosplit
func (self class) GetSelectedPath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspector.Bind_get_selected_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self Instance) OnPropertySelected(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyKeyed(cb func(property string, value gd.Variant, advance bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyDeleted(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_deleted"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceSelected(cb func(resource objects.Resource, path string)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnObjectIdSelected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("object_id_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyEdited(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_edited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyToggled(cb func(property string, checked bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_toggled"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEditedObjectChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("edited_object_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRestartRequested(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("restart_requested"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorInspector() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorInspector() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScrollContainer() ScrollContainer.Advanced {
	return *((*ScrollContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsScrollContainer() ScrollContainer.Instance {
	return *((*ScrollContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsScrollContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsScrollContainer(), name)
	}
}
func init() {
	classdb.Register("EditorInspector", func(ptr gd.Object) any { return classdb.EditorInspector(ptr) })
}
