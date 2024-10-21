package EditorScriptPicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/EditorResourcePicker"
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
Similar to [EditorResourcePicker] this [Control] node is used in the editor's Inspector dock, but only to edit the [code]script[/code] property of a [Node]. Default options for creating new resources of all possible subtypes are replaced with dedicated buttons that open the "Attach Node Script" dialog. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] You must set the [member script_owner] for the custom context menu items to work.

*/
type Simple [1]classdb.EditorScriptPicker
func (self Simple) SetScriptOwner(owner_node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScriptOwner(owner_node)
}
func (self Simple) GetScriptOwner() [1]classdb.Node {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Node(Expert(self).GetScriptOwner(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorScriptPicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetScriptOwner(owner_node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(owner_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScriptPicker.Bind_set_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScriptOwner(ctx gd.Lifetime) object.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScriptPicker.Bind_get_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorScriptPicker() Expert { return self[0].AsEditorScriptPicker() }


//go:nosplit
func (self Simple) AsEditorScriptPicker() Simple { return self[0].AsEditorScriptPicker() }


//go:nosplit
func (self class) AsEditorResourcePicker() EditorResourcePicker.Expert { return self[0].AsEditorResourcePicker() }


//go:nosplit
func (self Simple) AsEditorResourcePicker() EditorResourcePicker.Simple { return self[0].AsEditorResourcePicker() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorScriptPicker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
