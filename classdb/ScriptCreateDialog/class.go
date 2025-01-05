// Package ScriptCreateDialog provides methods for working with ScriptCreateDialog object instances.
package ScriptCreateDialog

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/ConfirmationDialog"
import "graphics.gd/classdb/AcceptDialog"
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [ScriptCreateDialog] creates script files according to a given template for a given scripting language. The standard use is to configure its fields prior to calling one of the [method Window.popup] methods.
[codeblocks]
[gdscript]
func _ready():

	var dialog = ScriptCreateDialog.new();
	dialog.config("Node", "res://new_node.gd") # For in-engine types.
	dialog.config("\"res://base_node.gd\"", "res://derived_node.gd") # For script types.
	dialog.popup_centered()

[/gdscript]
[csharp]
public override void _Ready()

	{
	    var dialog = new ScriptCreateDialog();
	    dialog.Config("Node", "res://NewNode.cs"); // For in-engine types.
	    dialog.Config("\"res://BaseNode.cs\"", "res://DerivedNode.cs"); // For script types.
	    dialog.PopupCentered();
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.ScriptCreateDialog
type Any interface {
	gd.IsClass
	AsScriptCreateDialog() Instance
}

/*
Prefills required fields to configure the ScriptCreateDialog for use.
*/
func (self Instance) Config(inherits string, path string) {
	class(self).Config(gd.NewString(inherits), gd.NewString(path), true, true)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScriptCreateDialog

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptCreateDialog"))
	return Instance{*(*gdclass.ScriptCreateDialog)(unsafe.Pointer(&object))}
}

/*
Prefills required fields to configure the ScriptCreateDialog for use.
*/
//go:nosplit
func (self class) Config(inherits gd.String, path gd.String, built_in_enabled bool, load_enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(inherits))
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, built_in_enabled)
	callframe.Arg(frame, load_enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptCreateDialog.Bind_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnScriptCreated(cb func(script [1]gdclass.Script)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("script_created"), gd.NewCallable(cb), 0)
}

func (self class) AsScriptCreateDialog() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptCreateDialog() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsConfirmationDialog() ConfirmationDialog.Advanced {
	return *((*ConfirmationDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsConfirmationDialog() ConfirmationDialog.Instance {
	return *((*ConfirmationDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAcceptDialog() AcceptDialog.Advanced {
	return *((*AcceptDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAcceptDialog() AcceptDialog.Instance {
	return *((*AcceptDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Advanced(self.AsConfirmationDialog()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Instance(self.AsConfirmationDialog()), name)
	}
}
func init() {
	gdclass.Register("ScriptCreateDialog", func(ptr gd.Object) any {
		return [1]gdclass.ScriptCreateDialog{*(*gdclass.ScriptCreateDialog)(unsafe.Pointer(&ptr))}
	})
}
