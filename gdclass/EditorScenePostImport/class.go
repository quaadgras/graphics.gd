package EditorScenePostImport

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
Imported scenes can be automatically modified right after import by setting their [b]Custom Script[/b] Import property to a [code]tool[/code] script that inherits from this class.
The [method _post_import] callback receives the imported scene's root node and returns the modified version of the scene. Usage example:
[codeblocks]
[gdscript]
@tool # Needed so it runs in editor.
extends EditorScenePostImport

# This sample changes all node names.
# Called right after the scene is imported and gets the root node.
func _post_import(scene):
    # Change all node names to "modified_[oldnodename]"
    iterate(scene)
    return scene # Remember to return the imported scene

func iterate(node):
    if node != null:
        node.name = "modified_" + node.name
        for child in node.get_children():
            iterate(child)
[/gdscript]
[csharp]
using Godot;

// This sample changes all node names.
// Called right after the scene is imported and gets the root node.
[Tool]
public partial class NodeRenamer : EditorScenePostImport
{
    public override GodotObject _PostImport(Node scene)
    {
        // Change all node names to "modified_[oldnodename]"
        Iterate(scene);
        return scene; // Remember to return the imported scene
    }

    public void Iterate(Node node)
    {
        if (node != null)
        {
            node.Name = $"modified_{node.Name}";
            foreach (Node child in node.GetChildren())
            {
                Iterate(child);
            }
        }
    }
}
[/csharp]
[/codeblocks]
	// EditorScenePostImport methods that can be overridden by a [Class] that extends it.
	type EditorScenePostImport interface {
		//Called after the scene was imported. This method must return the modified version of the scene.
		PostImport(scene gdclass.Node) gd.Object
	}

*/
type Go [1]classdb.EditorScenePostImport

/*
Called after the scene was imported. This method must return the modified version of the scene.
*/
func (Go) _post_import(impl func(ptr unsafe.Pointer, scene gdclass.Node) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var scene gdclass.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		gc.End()
	}
}

/*
Returns the source file path which got imported (e.g. [code]res://scene.dae[/code]).
*/
func (self Go) GetSourceFile() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSourceFile(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorScenePostImport
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorScenePostImport"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called after the scene was imported. This method must return the modified version of the scene.
*/
func (class) _post_import(impl func(ptr unsafe.Pointer, scene gdclass.Node) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var scene gdclass.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

/*
Returns the source file path which got imported (e.g. [code]res://scene.dae[/code]).
*/
//go:nosplit
func (self class) GetSourceFile(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScenePostImport.Bind_get_source_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorScenePostImport() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorScenePostImport() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_post_import": return reflect.ValueOf(self._post_import);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_post_import": return reflect.ValueOf(self._post_import);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorScenePostImport", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
