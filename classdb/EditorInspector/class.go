// Package EditorInspector provides methods for working with EditorInspector object instances.
package EditorInspector

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/ScrollContainer"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

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

/*
This is the control that implements property editing in the editor's Settings dialogs, the Inspector dock, etc. To get the [EditorInspector] used in the editor's Inspector dock, use [method EditorInterface.get_inspector].
[EditorInspector] will show properties in the same order as the array returned by [method Object.get_property_list].
If a property's name is path-like (i.e. if it contains forward slashes), [EditorInspector] will create nested sections for "directories" along the path. For example, if a property is named [code]highlighting/gdscript/node_path_color[/code], it will be shown as "Node Path Color" inside the "GDScript" section nested inside the "Highlighting" section.
If a property has [constant PROPERTY_USAGE_GROUP] usage, it will group subsequent properties whose name starts with the property's hint string. The group ends when a property does not start with that hint string or when a new group starts. An empty group name effectively ends the current group. [EditorInspector] will create a top-level section for each group. For example, if a property with group usage is named [code]Collide With[/code] and its hint string is [code]collide_with_[/code], a subsequent [code]collide_with_area[/code] property will be shown as "Area" inside the "Collide With" section. There is also a special case: when the hint string contains the name of a property, that property is grouped too. This is mainly to help grouping properties like [code]font[/code], [code]font_color[/code] and [code]font_size[/code] (using the hint string [code]font_[/code]).
If a property has [constant PROPERTY_USAGE_SUBGROUP] usage, a subgroup will be created in the same way as a group, and a second-level section will be created for each subgroup.
[b]Note:[/b] Unlike sections created from path-like property names, [EditorInspector] won't capitalize the name for sections created from groups. So properties with group usage usually use capitalized names instead of snake_cased names.
*/
type Instance [1]gdclass.EditorInspector

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorInspector() Instance
}

/*
Gets the path of the currently selected property.
*/
func (self Instance) GetSelectedPath() string { //gd:EditorInspector.get_selected_path
	return string(class(self).GetSelectedPath().String())
}

/*
Returns the object currently selected in this inspector.
*/
func (self Instance) GetEditedObject() Object.Instance { //gd:EditorInspector.get_edited_object
	return Object.Instance(class(self).GetEditedObject())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorInspector

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorInspector"))
	casted := Instance{*(*gdclass.EditorInspector)(unsafe.Pointer(&object))}
	return casted
}

/*
Gets the path of the currently selected property.
*/
//go:nosplit
func (self class) GetSelectedPath() String.Readable { //gd:EditorInspector.get_selected_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspector.Bind_get_selected_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the object currently selected in this inspector.
*/
//go:nosplit
func (self class) GetEditedObject() [1]gd.Object { //gd:EditorInspector.get_edited_object
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspector.Bind_get_edited_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}
func (self Instance) OnPropertySelected(cb func(property string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("property_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyKeyed(cb func(property string, value any, advance bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("property_keyed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyDeleted(cb func(property string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("property_deleted"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceSelected(cb func(resource [1]gdclass.Resource, path string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnObjectIdSelected(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("object_id_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyEdited(cb func(property string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("property_edited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyToggled(cb func(property string, checked bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("property_toggled"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEditedObjectChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("edited_object_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRestartRequested(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("restart_requested"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(ScrollContainer.Advanced(self.AsScrollContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ScrollContainer.Instance(self.AsScrollContainer()), name)
	}
}
func init() {
	gdclass.Register("EditorInspector", func(ptr gd.Object) any {
		return [1]gdclass.EditorInspector{*(*gdclass.EditorInspector)(unsafe.Pointer(&ptr))}
	})
}
