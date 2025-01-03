package EditorScriptPicker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/EditorResourcePicker"
import "graphics.gd/objects/HBoxContainer"
import "graphics.gd/objects/BoxContainer"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Similar to [EditorResourcePicker] this [Control] node is used in the editor's Inspector dock, but only to edit the [code]script[/code] property of a [Node]. Default options for creating new resources of all possible subtypes are replaced with dedicated buttons that open the "Attach Node Script" dialog. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] You must set the [member script_owner] for the custom context menu items to work.
*/
type Instance [1]classdb.EditorScriptPicker
type Any interface {
	gd.IsClass
	AsEditorScriptPicker() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorScriptPicker

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScriptPicker"))
	return Instance{*(*classdb.EditorScriptPicker)(unsafe.Pointer(&object))}
}

func (self Instance) ScriptOwner() objects.Node {
	return objects.Node(class(self).GetScriptOwner())
}

func (self Instance) SetScriptOwner(value objects.Node) {
	class(self).SetScriptOwner(value)
}

//go:nosplit
func (self class) SetScriptOwner(owner_node objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(owner_node[0].AsObject()))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScriptPicker.Bind_set_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScriptOwner() objects.Node {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScriptPicker.Bind_get_script_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Node{gd.PointerMustAssertInstanceID[classdb.Node](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsEditorScriptPicker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorScriptPicker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsEditorResourcePicker() EditorResourcePicker.Advanced {
	return *((*EditorResourcePicker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorResourcePicker() EditorResourcePicker.Instance {
	return *((*EditorResourcePicker.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsHBoxContainer() HBoxContainer.Advanced {
	return *((*HBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsHBoxContainer() HBoxContainer.Instance {
	return *((*HBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsEditorResourcePicker(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorResourcePicker(), name)
	}
}
func init() {
	classdb.Register("EditorScriptPicker", func(ptr gd.Object) any {
		return [1]classdb.EditorScriptPicker{*(*classdb.EditorScriptPicker)(unsafe.Pointer(&ptr))}
	})
}
