package ScriptCreateDialog

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ConfirmationDialog"
import "grow.graphics/gd/object/AcceptDialog"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.ScriptCreateDialog
func (self Simple) Config(inherits string, path string, built_in_enabled bool, load_enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Config(gc.String(inherits), gc.String(path), built_in_enabled, load_enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ScriptCreateDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Prefills required fields to configure the ScriptCreateDialog for use.
*/
//go:nosplit
func (self class) Config(inherits gd.String, path gd.String, built_in_enabled bool, load_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(inherits))
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, built_in_enabled)
	callframe.Arg(frame, load_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptCreateDialog.Bind_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsScriptCreateDialog() Expert { return self[0].AsScriptCreateDialog() }


//go:nosplit
func (self Simple) AsScriptCreateDialog() Simple { return self[0].AsScriptCreateDialog() }


//go:nosplit
func (self class) AsConfirmationDialog() ConfirmationDialog.Expert { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self Simple) AsConfirmationDialog() ConfirmationDialog.Simple { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self class) AsAcceptDialog() AcceptDialog.Expert { return self[0].AsAcceptDialog() }


//go:nosplit
func (self Simple) AsAcceptDialog() AcceptDialog.Simple { return self[0].AsAcceptDialog() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ScriptCreateDialog", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
