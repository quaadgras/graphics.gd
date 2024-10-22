package EditorScriptPicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorResourcePicker"
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
var _ mmm.Lifetime

/*
Similar to [EditorResourcePicker] this [Control] node is used in the editor's Inspector dock, but only to edit the [code]script[/code] property of a [Node]. Default options for creating new resources of all possible subtypes are replaced with dedicated buttons that open the "Attach Node Script" dialog. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] You must set the [member script_owner] for the custom context menu items to work.

*/
type Go [1]classdb.EditorScriptPicker
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorScriptPicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorScriptPicker"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ScriptOwner() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Node(class(self).GetScriptOwner(gc))
}

func (self Go) SetScriptOwner(value gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScriptOwner(value)
}

//go:nosplit
func (self class) SetScriptOwner(owner_node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(owner_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScriptPicker.Bind_set_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScriptOwner(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScriptPicker.Bind_get_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsEditorScriptPicker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorScriptPicker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorResourcePicker() EditorResourcePicker.GD { return *((*EditorResourcePicker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorResourcePicker() EditorResourcePicker.Go { return *((*EditorResourcePicker.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsEditorResourcePicker(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorResourcePicker(), name)
	}
}
func init() {classdb.Register("EditorScriptPicker", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
