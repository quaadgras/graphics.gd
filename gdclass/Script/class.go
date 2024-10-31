package Script

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A class stored as a resource. A script extends the functionality of all objects that instantiate it.
This is the base class for all scripts and should not be used directly. Trying to create a new script with this class will result in an error.
The [code]new[/code] method of a script subclass creates a new instance. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
*/
type Instance [1]classdb.Script

/*
Returns [code]true[/code] if the script can be instantiated.
*/
func (self Instance) CanInstantiate() bool {
	return bool(class(self).CanInstantiate())
}

/*
Returns [code]true[/code] if [param base_object] is an instance of this script.
*/
func (self Instance) InstanceHas(base_object gd.Object) bool {
	return bool(class(self).InstanceHas(base_object))
}

/*
Returns [code]true[/code] if the script contains non-empty source code.
[b]Note:[/b] If a script does not have source code, this does not mean that it is invalid or unusable. For example, a [GDScript] that was exported with binary tokenization has no source code, but still behaves as expected and could be instantiated. This can be checked with [method can_instantiate].
*/
func (self Instance) HasSourceCode() bool {
	return bool(class(self).HasSourceCode())
}

/*
Reloads the script's class implementation. Returns an error code.
*/
func (self Instance) Reload() gd.Error {
	return gd.Error(class(self).Reload(false))
}

/*
Returns the script directly inherited by this script.
*/
func (self Instance) GetBaseScript() gdclass.Script {
	return gdclass.Script(class(self).GetBaseScript())
}

/*
Returns the script's base type.
*/
func (self Instance) GetInstanceBaseType() string {
	return string(class(self).GetInstanceBaseType().String())
}

/*
Returns the class name associated with the script, if there is one. Returns an empty string otherwise.
To give the script a global name, you can use the [code]class_name[/code] keyword in GDScript and the [code][GlobalClass][/code] attribute in C#.
[codeblocks]
[gdscript]
class_name MyNode
extends Node
[/gdscript]
[csharp]
using Godot;

[GlobalClass]
public partial class MyNode : Node
{
}
[/csharp]
[/codeblocks]
*/
func (self Instance) GetGlobalName() string {
	return string(class(self).GetGlobalName().String())
}

/*
Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
*/
func (self Instance) HasScriptSignal(signal_name string) bool {
	return bool(class(self).HasScriptSignal(gd.NewStringName(signal_name)))
}

/*
Returns the list of properties in this [Script].
*/
func (self Instance) GetScriptPropertyList() gd.Array {
	return gd.Array(class(self).GetScriptPropertyList())
}

/*
Returns the list of methods in this [Script].
*/
func (self Instance) GetScriptMethodList() gd.Array {
	return gd.Array(class(self).GetScriptMethodList())
}

/*
Returns the list of user signals defined in this [Script].
*/
func (self Instance) GetScriptSignalList() gd.Array {
	return gd.Array(class(self).GetScriptSignalList())
}

/*
Returns a dictionary containing constant names and their values.
*/
func (self Instance) GetScriptConstantMap() gd.Dictionary {
	return gd.Dictionary(class(self).GetScriptConstantMap())
}

/*
Returns the default value of the specified property.
*/
func (self Instance) GetPropertyDefaultValue(property string) gd.Variant {
	return gd.Variant(class(self).GetPropertyDefaultValue(gd.NewStringName(property)))
}

/*
Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
*/
func (self Instance) IsTool() bool {
	return bool(class(self).IsTool())
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (self Instance) IsAbstract() bool {
	return bool(class(self).IsAbstract())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Script

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Script"))
	return Instance{classdb.Script(object)}
}

func (self Instance) SourceCode() string {
	return string(class(self).GetSourceCode().String())
}

func (self Instance) SetSourceCode(value string) {
	class(self).SetSourceCode(gd.NewString(value))
}

/*
Returns [code]true[/code] if the script can be instantiated.
*/
//go:nosplit
func (self class) CanInstantiate() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param base_object] is an instance of this script.
*/
//go:nosplit
func (self class) InstanceHas(base_object gd.Object) bool {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(base_object))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_instance_has, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script contains non-empty source code.
[b]Note:[/b] If a script does not have source code, this does not mean that it is invalid or unusable. For example, a [GDScript] that was exported with binary tokenization has no source code, but still behaves as expected and could be instantiated. This can be checked with [method can_instantiate].
*/
//go:nosplit
func (self class) HasSourceCode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_has_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSourceCode() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceCode(source gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(source))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_set_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Reloads the script's class implementation. Returns an error code.
*/
//go:nosplit
func (self class) Reload(keep_state bool) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, keep_state)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_reload, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the script directly inherited by this script.
*/
//go:nosplit
func (self class) GetBaseScript() gdclass.Script {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_base_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Script{classdb.Script(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the script's base type.
*/
//go:nosplit
func (self class) GetInstanceBaseType() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_instance_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the class name associated with the script, if there is one. Returns an empty string otherwise.
To give the script a global name, you can use the [code]class_name[/code] keyword in GDScript and the [code][GlobalClass][/code] attribute in C#.
[codeblocks]
[gdscript]
class_name MyNode
extends Node
[/gdscript]
[csharp]
using Godot;

[GlobalClass]
public partial class MyNode : Node
{
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetGlobalName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_global_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
*/
//go:nosplit
func (self class) HasScriptSignal(signal_name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(signal_name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_has_script_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of properties in this [Script].
*/
//go:nosplit
func (self class) GetScriptPropertyList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_property_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of methods in this [Script].
*/
//go:nosplit
func (self class) GetScriptMethodList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_method_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of user signals defined in this [Script].
*/
//go:nosplit
func (self class) GetScriptSignalList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_signal_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a dictionary containing constant names and their values.
*/
//go:nosplit
func (self class) GetScriptConstantMap() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_constant_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the default value of the specified property.
*/
//go:nosplit
func (self class) GetPropertyDefaultValue(property gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
*/
//go:nosplit
func (self class) IsTool() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_is_tool, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
//go:nosplit
func (self class) IsAbstract() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_is_abstract, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsScript() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScript() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Script", func(ptr gd.Object) any { return classdb.Script(ptr) }) }
