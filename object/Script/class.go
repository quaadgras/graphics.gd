package Script

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A class stored as a resource. A script extends the functionality of all objects that instantiate it.
This is the base class for all scripts and should not be used directly. Trying to create a new script with this class will result in an error.
The [code]new[/code] method of a script subclass creates a new instance. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.

*/
type Simple [1]classdb.Script
func (self Simple) CanInstantiate() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanInstantiate())
}
func (self Simple) InstanceHas(base_object gd.Object) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).InstanceHas(base_object))
}
func (self Simple) HasSourceCode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSourceCode())
}
func (self Simple) GetSourceCode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSourceCode(gc).String())
}
func (self Simple) SetSourceCode(source string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSourceCode(gc.String(source))
}
func (self Simple) Reload(keep_state bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Reload(keep_state))
}
func (self Simple) GetBaseScript() [1]classdb.Script {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Script(Expert(self).GetBaseScript(gc))
}
func (self Simple) GetInstanceBaseType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetInstanceBaseType(gc).String())
}
func (self Simple) GetGlobalName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetGlobalName(gc).String())
}
func (self Simple) HasScriptSignal(signal_name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasScriptSignal(gc.StringName(signal_name)))
}
func (self Simple) GetScriptPropertyList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetScriptPropertyList(gc))
}
func (self Simple) GetScriptMethodList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetScriptMethodList(gc))
}
func (self Simple) GetScriptSignalList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetScriptSignalList(gc))
}
func (self Simple) GetScriptConstantMap() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetScriptConstantMap(gc))
}
func (self Simple) GetPropertyDefaultValue(property string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetPropertyDefaultValue(gc, gc.StringName(property)))
}
func (self Simple) IsTool() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTool())
}
func (self Simple) IsAbstract() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAbstract())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Script
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if the script can be instantiated.
*/
//go:nosplit
func (self class) CanInstantiate() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param base_object] is an instance of this script.
*/
//go:nosplit
func (self class) InstanceHas(base_object gd.Object) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(base_object.AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_instance_has, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_has_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSourceCode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSourceCode(source gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(source))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_set_source_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Reloads the script's class implementation. Returns an error code.
*/
//go:nosplit
func (self class) Reload(keep_state bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keep_state)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_reload, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the script directly inherited by this script.
*/
//go:nosplit
func (self class) GetBaseScript(ctx gd.Lifetime) object.Script {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_base_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Script
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the script's base type.
*/
//go:nosplit
func (self class) GetInstanceBaseType(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_instance_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) GetGlobalName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_global_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
*/
//go:nosplit
func (self class) HasScriptSignal(signal_name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(signal_name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_has_script_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of properties in this [Script].
*/
//go:nosplit
func (self class) GetScriptPropertyList(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_script_property_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns the list of methods in this [Script].
*/
//go:nosplit
func (self class) GetScriptMethodList(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_script_method_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns the list of user signals defined in this [Script].
*/
//go:nosplit
func (self class) GetScriptSignalList(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_script_signal_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns a dictionary containing constant names and their values.
*/
//go:nosplit
func (self class) GetScriptConstantMap(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_script_constant_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the default value of the specified property.
*/
//go:nosplit
func (self class) GetPropertyDefaultValue(ctx gd.Lifetime, property gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
*/
//go:nosplit
func (self class) IsTool() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_is_tool, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
//go:nosplit
func (self class) IsAbstract() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Script.Bind_is_abstract, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsScript() Expert { return self[0].AsScript() }


//go:nosplit
func (self Simple) AsScript() Simple { return self[0].AsScript() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("Script", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
