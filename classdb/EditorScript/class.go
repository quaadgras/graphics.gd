// Package EditorScript provides methods for working with EditorScript object instances.
package EditorScript

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

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorScript)
*/
type Instance [1]gdclass.EditorScript

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorScript() Instance
}
type Interface interface {
	//This method is executed by the Editor when [b]File > Run[/b] is used.
	Run()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Run() { return }

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (Instance) _run(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
func (self Instance) AddRootNode(node [1]gdclass.Node) { //gd:EditorScript.add_root_node
	class(self).AddRootNode(node)
}

/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
func (self Instance) GetScene() [1]gdclass.Node { //gd:EditorScript.get_scene
	return [1]gdclass.Node(class(self).GetScene())
}

/*
Returns the [EditorInterface] singleton instance.
*/
func (self Instance) GetEditorInterface() [1]gdclass.EditorInterface { //gd:EditorScript.get_editor_interface
	return [1]gdclass.EditorInterface(class(self).GetEditorInterface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorScript

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScript"))
	casted := Instance{*(*gdclass.EditorScript)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (class) _run(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
//go:nosplit
func (self class) AddRootNode(node [1]gdclass.Node) { //gd:EditorScript.add_root_node
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(node[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_add_root_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
//go:nosplit
func (self class) GetScene() [1]gdclass.Node { //gd:EditorScript.get_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_get_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [EditorInterface] singleton instance.
*/
//go:nosplit
func (self class) GetEditorInterface() [1]gdclass.EditorInterface { //gd:EditorScript.get_editor_interface
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScript.Bind_get_editor_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorInterface{gd.PointerLifetimeBoundTo[gdclass.EditorInterface](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsEditorScript() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorScript() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_run":
		return reflect.ValueOf(self._run)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_run":
		return reflect.ValueOf(self._run)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorScript", func(ptr gd.Object) any {
		return [1]gdclass.EditorScript{*(*gdclass.EditorScript)(unsafe.Pointer(&ptr))}
	})
}
