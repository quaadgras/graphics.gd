// Package Script provides methods for working with Script object instances.
package Script

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
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Resource"

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
var _ Path.ToNode

/*
A class stored as a resource. A script extends the functionality of all objects that instantiate it.
This is the base class for all scripts and should not be used directly. Trying to create a new script with this class will result in an error.
The [code]new[/code] method of a script subclass creates a new instance. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
*/
type Instance [1]gdclass.Script

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsScript() Instance
}

/*
Returns [code]true[/code] if the script can be instantiated.
*/
func (self Instance) CanInstantiate() bool { //gd:Script.can_instantiate
	return bool(class(self).CanInstantiate())
}

/*
Returns [code]true[/code] if [param base_object] is an instance of this script.
*/
func (self Instance) InstanceHas(base_object Object.Instance) bool { //gd:Script.instance_has
	return bool(class(self).InstanceHas(base_object))
}

/*
Returns [code]true[/code] if the script contains non-empty source code.
[b]Note:[/b] If a script does not have source code, this does not mean that it is invalid or unusable. For example, a [GDScript] that was exported with binary tokenization has no source code, but still behaves as expected and could be instantiated. This can be checked with [method can_instantiate].
*/
func (self Instance) HasSourceCode() bool { //gd:Script.has_source_code
	return bool(class(self).HasSourceCode())
}

/*
Reloads the script's class implementation. Returns an error code.
*/
func (self Instance) Reload() error { //gd:Script.reload
	return error(gd.ToError(class(self).Reload(false)))
}

/*
Returns the script directly inherited by this script.
*/
func (self Instance) GetBaseScript() [1]gdclass.Script { //gd:Script.get_base_script
	return [1]gdclass.Script(class(self).GetBaseScript())
}

/*
Returns the script's base type.
*/
func (self Instance) GetInstanceBaseType() string { //gd:Script.get_instance_base_type
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
func (self Instance) GetGlobalName() string { //gd:Script.get_global_name
	return string(class(self).GetGlobalName().String())
}

/*
Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
*/
func (self Instance) HasScriptSignal(signal_name string) bool { //gd:Script.has_script_signal
	return bool(class(self).HasScriptSignal(String.Name(String.New(signal_name))))
}

/*
Returns the list of properties in this [Script].
*/
func (self Instance) GetScriptPropertyList() []PropertyInfo { //gd:Script.get_script_property_list
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(class(self).GetScriptPropertyList())))
}

/*
Returns the list of methods in this [Script].
*/
func (self Instance) GetScriptMethodList() []PropertyInfo { //gd:Script.get_script_method_list
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(class(self).GetScriptMethodList())))
}

/*
Returns the list of user signals defined in this [Script].
*/
func (self Instance) GetScriptSignalList() []SignalInfo { //gd:Script.get_script_signal_list
	return []SignalInfo(gd.ArrayAs[[]SignalInfo](gd.InternalArray(class(self).GetScriptSignalList())))
}

/*
Returns a dictionary containing constant names and their values.
*/
func (self Instance) GetScriptConstantMap() map[string]interface{} { //gd:Script.get_script_constant_map
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](class(self).GetScriptConstantMap()))
}

/*
Returns the default value of the specified property.
*/
func (self Instance) GetPropertyDefaultValue(property string) any { //gd:Script.get_property_default_value
	return any(class(self).GetPropertyDefaultValue(String.Name(String.New(property))).Interface())
}

/*
Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
*/
func (self Instance) IsTool() bool { //gd:Script.is_tool
	return bool(class(self).IsTool())
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (self Instance) IsAbstract() bool { //gd:Script.is_abstract
	return bool(class(self).IsAbstract())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Script

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Script"))
	casted := Instance{*(*gdclass.Script)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SourceCode() string {
	return string(class(self).GetSourceCode().String())
}

func (self Instance) SetSourceCode(value string) {
	class(self).SetSourceCode(String.New(value))
}

/*
Returns [code]true[/code] if the script can be instantiated.
*/
//go:nosplit
func (self class) CanInstantiate() bool { //gd:Script.can_instantiate
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param base_object] is an instance of this script.
*/
//go:nosplit
func (self class) InstanceHas(base_object [1]gd.Object) bool { //gd:Script.instance_has
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_object[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_instance_has, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script contains non-empty source code.
[b]Note:[/b] If a script does not have source code, this does not mean that it is invalid or unusable. For example, a [GDScript] that was exported with binary tokenization has no source code, but still behaves as expected and could be instantiated. This can be checked with [method can_instantiate].
*/
//go:nosplit
func (self class) HasSourceCode() bool { //gd:Script.has_source_code
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_has_source_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSourceCode() String.Readable { //gd:Script.get_source_code
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_source_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceCode(source String.Readable) { //gd:Script.set_source_code
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(source)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_set_source_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Reloads the script's class implementation. Returns an error code.
*/
//go:nosplit
func (self class) Reload(keep_state bool) gd.Error { //gd:Script.reload
	var frame = callframe.New()
	callframe.Arg(frame, keep_state)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_reload, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the script directly inherited by this script.
*/
//go:nosplit
func (self class) GetBaseScript() [1]gdclass.Script { //gd:Script.get_base_script
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_base_script, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Script{gd.PointerWithOwnershipTransferredToGo[gdclass.Script](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the script's base type.
*/
//go:nosplit
func (self class) GetInstanceBaseType() String.Name { //gd:Script.get_instance_base_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_instance_base_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
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
func (self class) GetGlobalName() String.Name { //gd:Script.get_global_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_global_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
*/
//go:nosplit
func (self class) HasScriptSignal(signal_name String.Name) bool { //gd:Script.has_script_signal
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(signal_name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_has_script_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of properties in this [Script].
*/
//go:nosplit
func (self class) GetScriptPropertyList() Array.Contains[Dictionary.Any] { //gd:Script.get_script_property_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_property_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the list of methods in this [Script].
*/
//go:nosplit
func (self class) GetScriptMethodList() Array.Contains[Dictionary.Any] { //gd:Script.get_script_method_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_method_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the list of user signals defined in this [Script].
*/
//go:nosplit
func (self class) GetScriptSignalList() Array.Contains[Dictionary.Any] { //gd:Script.get_script_signal_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_signal_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a dictionary containing constant names and their values.
*/
//go:nosplit
func (self class) GetScriptConstantMap() Dictionary.Any { //gd:Script.get_script_constant_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_script_constant_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the default value of the specified property.
*/
//go:nosplit
func (self class) GetPropertyDefaultValue(property String.Name) gd.Variant { //gd:Script.get_property_default_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
*/
//go:nosplit
func (self class) IsTool() bool { //gd:Script.is_tool
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_is_tool, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
//go:nosplit
func (self class) IsAbstract() bool { //gd:Script.is_abstract
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Script.Bind_is_abstract, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Script", func(ptr gd.Object) any { return [1]gdclass.Script{*(*gdclass.Script)(unsafe.Pointer(&ptr))} })
}

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)

type PropertyInfo struct {
	ClassName  string       `gd:"class_name"`
	Name       string       `gd:"name"`
	Hint       int          `gd:"hint"`
	HintString string       `gd:"hint_string"`
	Type       reflect.Type `gd:"type"`
	Usage      int          `gd:"usage"`
}
type SignalInfo struct {
	Name        string         `gd:"name"`
	Flags       int            `gd:"flags"`
	ID          int            `gd:"id"`
	DefaultArgs []interface{}  `gd:"default_args"`
	Args        []PropertyInfo `gd:"args"`
}
