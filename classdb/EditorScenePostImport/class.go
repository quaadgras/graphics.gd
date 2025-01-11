// Package EditorScenePostImport provides methods for working with EditorScenePostImport object instances.
package EditorScenePostImport

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

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

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorScenePostImport)
*/
type Instance [1]gdclass.EditorScenePostImport

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorScenePostImport() Instance
}
type Interface interface {
	//Called after the scene was imported. This method must return the modified version of the scene.
	PostImport(scene [1]gdclass.Node) Object.Instance
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) PostImport(scene [1]gdclass.Node) (_ Object.Instance) { return }

/*
Called after the scene was imported. This method must return the modified version of the scene.
*/
func (Instance) _post_import(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node) Object.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the source file path which got imported (e.g. [code]res://scene.dae[/code]).
*/
func (self Instance) GetSourceFile() string {
	return string(class(self).GetSourceFile().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorScenePostImport

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScenePostImport"))
	casted := Instance{*(*gdclass.EditorScenePostImport)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Called after the scene was imported. This method must return the modified version of the scene.
*/
func (class) _post_import(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node) [1]gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the source file path which got imported (e.g. [code]res://scene.dae[/code]).
*/
//go:nosplit
func (self class) GetSourceFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImport.Bind_get_source_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorScenePostImport() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorScenePostImport() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_post_import":
		return reflect.ValueOf(self._post_import)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_post_import":
		return reflect.ValueOf(self._post_import)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorScenePostImport", func(ptr gd.Object) any {
		return [1]gdclass.EditorScenePostImport{*(*gdclass.EditorScenePostImport)(unsafe.Pointer(&ptr))}
	})
}
