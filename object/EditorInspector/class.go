package EditorInspector

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ScrollContainer"
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
This is the control that implements property editing in the editor's Settings dialogs, the Inspector dock, etc. To get the [EditorInspector] used in the editor's Inspector dock, use [method EditorInterface.get_inspector].
[EditorInspector] will show properties in the same order as the array returned by [method Object.get_property_list].
If a property's name is path-like (i.e. if it contains forward slashes), [EditorInspector] will create nested sections for "directories" along the path. For example, if a property is named [code]highlighting/gdscript/node_path_color[/code], it will be shown as "Node Path Color" inside the "GDScript" section nested inside the "Highlighting" section.
If a property has [constant PROPERTY_USAGE_GROUP] usage, it will group subsequent properties whose name starts with the property's hint string. The group ends when a property does not start with that hint string or when a new group starts. An empty group name effectively ends the current group. [EditorInspector] will create a top-level section for each group. For example, if a property with group usage is named [code]Collide With[/code] and its hint string is [code]collide_with_[/code], a subsequent [code]collide_with_area[/code] property will be shown as "Area" inside the "Collide With" section. There is also a special case: when the hint string contains the name of a property, that property is grouped too. This is mainly to help grouping properties like [code]font[/code], [code]font_color[/code] and [code]font_size[/code] (using the hint string [code]font_[/code]).
If a property has [constant PROPERTY_USAGE_SUBGROUP] usage, a subgroup will be created in the same way as a group, and a second-level section will be created for each subgroup.
[b]Note:[/b] Unlike sections created from path-like property names, [EditorInspector] won't capitalize the name for sections created from groups. So properties with group usage usually use capitalized names instead of snake_cased names.

*/
type Simple [1]classdb.EditorInspector
func (self Simple) GetSelectedPath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSelectedPath(gc).String())
}
func (self Simple) GetEditedObject() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(Expert(self).GetEditedObject(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorInspector
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Gets the path of the currently selected property.
*/
//go:nosplit
func (self class) GetSelectedPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInspector.Bind_get_selected_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the object currently selected in this inspector.
*/
//go:nosplit
func (self class) GetEditedObject(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInspector.Bind_get_edited_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorInspector() Expert { return self[0].AsEditorInspector() }


//go:nosplit
func (self Simple) AsEditorInspector() Simple { return self[0].AsEditorInspector() }


//go:nosplit
func (self class) AsScrollContainer() ScrollContainer.Expert { return self[0].AsScrollContainer() }


//go:nosplit
func (self Simple) AsScrollContainer() ScrollContainer.Simple { return self[0].AsScrollContainer() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorInspector", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
