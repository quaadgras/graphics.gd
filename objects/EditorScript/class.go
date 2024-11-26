package EditorScript

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Scripts extending this class and implementing its [method _run] method can be executed from the Script Editor's [b]File > Run[/b] menu option (or by pressing [kbd]Ctrl + Shift + X[/kbd]) while the editor is running. This is useful for adding custom in-editor functionality to Godot. For more complex additions, consider using [EditorPlugin]s instead.
[b]Note:[/b] Extending scripts need to have [code]tool[/code] mode enabled.
[b]Example script:[/b]
[codeblocks]
[gdscript]
@tool
extends EditorScript

func _run():

	print("Hello from the Godot Editor!")

[/gdscript]
[csharp]
using Godot;

[Tool]
public partial class HelloEditor : EditorScript

	{
	    public override void _Run()
	    {
	        GD.Print("Hello from the Godot Editor!");
	    }
	}

[/csharp]
[/codeblocks]
[b]Note:[/b] The script is run in the Editor context, which means the output is visible in the console window started with the Editor (stdout) instead of the usual Godot [b]Output[/b] dock.
[b]Note:[/b] EditorScript is [RefCounted], meaning it is destroyed when nothing references it. This can cause errors during asynchronous operations if there are no references to the script.

	// EditorScript methods that can be overridden by a [Class] that extends it.
	type EditorScript interface {
		//This method is executed by the Editor when [b]File > Run[/b] is used.
		Run()
	}
*/
type Instance [1]classdb.EditorScript

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (Instance) _run(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
func (self Instance) AddRootNode(node objects.Node) {
	class(self).AddRootNode(node)
}

/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
func (self Instance) GetScene() objects.Node {
	return objects.Node(class(self).GetScene())
}

/*
Returns the [EditorInterface] singleton instance.
*/
func (self Instance) GetEditorInterface() objects.EditorInterface {
	return objects.EditorInterface(class(self).GetEditorInterface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorScript

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScript"))
	return Instance{classdb.EditorScript(object)}
}

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (class) _run(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
//go:nosplit
func (self class) AddRootNode(node objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(node[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_add_root_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
//go:nosplit
func (self class) GetScene() objects.Node {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_get_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Node{classdb.Node(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [EditorInterface] singleton instance.
*/
//go:nosplit
func (self class) GetEditorInterface() objects.EditorInterface {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_get_editor_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorInterface{classdb.EditorInterface(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsEditorScript() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorScript() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_run":
		return reflect.ValueOf(self._run)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_run":
		return reflect.ValueOf(self._run)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorScript", func(ptr gd.Object) any { return classdb.EditorScript(ptr) })
}
