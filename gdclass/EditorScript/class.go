package EditorScript

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.EditorScript

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (Go) _run(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
func (self Go) AddRootNode(node gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddRootNode(node)
}

/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
func (self Go) GetScene() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).GetScene(gc))
}

/*
Returns the [EditorInterface] singleton instance.
*/
func (self Go) GetEditorInterface() gdclass.EditorInterface {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.EditorInterface(class(self).GetEditorInterface(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorScript
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorScript"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
This method is executed by the Editor when [b]File > Run[/b] is used.
*/
func (class) _run(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Makes [param node] root of the currently opened scene. Only works if the scene is empty. If the [param node] is a scene instance, an inheriting scene will be created.
*/
//go:nosplit
func (self class) AddRootNode(node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScript.Bind_add_root_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the edited (current) scene's root [Node]. Equivalent of [method EditorInterface.get_edited_scene_root].
*/
//go:nosplit
func (self class) GetScene(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScript.Bind_get_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [EditorInterface] singleton instance.
*/
//go:nosplit
func (self class) GetEditorInterface(ctx gd.Lifetime) gdclass.EditorInterface {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScript.Bind_get_editor_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorInterface
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsEditorScript() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorScript() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_run": return reflect.ValueOf(self._run);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_run": return reflect.ValueOf(self._run);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorScript", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
